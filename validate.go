package validator

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	tw_translations "github.com/go-playground/validator/v10/translations/zh_tw"
	"reflect"
	"strings"
	"sync"
)

type Language byte

const (
	LangEn Language = iota
	LangZh
	LangTw
)

var v StructValidator = &defaultValidator{}

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
	trans    ut.Translator
}

func (v *defaultValidator) ValidateStruct(obj any) error {
	if obj == nil {
		return nil
	}

	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Ptr:
		err := v.ValidateStruct(value.Elem().Interface())
		return v.translate(err)
	case reflect.Struct:
		err := v.validateStruct(obj)
		return v.translate(err)
	case reflect.Slice, reflect.Array:
		count := value.Len()
		validateRet := make(sliceValidationError, 0)
		for i := 0; i < count; i++ {
			if err := v.ValidateStruct(value.Index(i).Interface()); err != nil {
				validateRet = append(validateRet, v.translate(err))
			}
		}
		if len(validateRet) == 0 {
			return nil
		}
		return validateRet
	default:
		return nil
	}
}
func (v *defaultValidator) validateStruct(obj any) error {
	v.lazyinit()
	return v.validate.Struct(obj)
}
func (v *defaultValidator) Engine() any {
	v.lazyinit()
	return v.validate
}
func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
	})
}
func (v *defaultValidator) translate(err error) error {
	if v.trans != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			validateRet := make(sliceValidationError, 0)
			for _, e := range errs {
				validateRet = append(validateRet, fmt.Errorf("%s", e.Translate(v.trans)))
			}
			if len(validateRet) == 0 {
				return nil
			}
			return validateRet
		} else {
			return err
		}
	}
	return err
}
func (v *defaultValidator) SetLanguage(lang Language) error {
	v.lazyinit()
	switch lang {
	case LangEn:
		v.trans = nil
		return nil
	case LangZh:
		z := zh.New()
		uni := ut.New(z, z)
		trans, ok := uni.GetTranslator("zh")
		if !ok {
			return fmt.Errorf("get translator fail: %v", lang)
		}
		if err := zh_translations.RegisterDefaultTranslations(v.validate, trans); err != nil {
			return err
		}
		v.trans = trans
		return nil
	case LangTw:
		z := zh_Hant_TW.New()
		uni := ut.New(z, z)
		trans, ok := uni.GetTranslator("zh_Hant_TW")
		if !ok {
			return fmt.Errorf("get translator fail: %v", lang)
		}
		if err := tw_translations.RegisterDefaultTranslations(v.validate, trans); err != nil {
			return err
		}
		v.trans = trans
		return nil
	}
	return fmt.Errorf("unsupported language: %v", lang)
}

func ValidateStruct(obj any) error {
	return v.ValidateStruct(obj)
}
func SetLanguage(lang Language) error {
	return v.SetLanguage(lang)
}

type sliceValidationError []error

func (err sliceValidationError) Error() string {
	n := len(err)
	switch n {
	case 0:
		return ""
	default:
		var b strings.Builder
		if err[0] != nil {
			fmt.Fprintf(&b, "%s.", err[0].Error())
		}
		if n > 1 {
			for i := 1; i < n; i++ {
				if err[i] != nil {
					b.WriteString("\n")
					fmt.Fprintf(&b, "%s.", err[i].Error())
				}
			}
		}
		return b.String()
	}
}
