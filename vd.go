package validator

type StructValidator interface {
	ValidateStruct(any) error
	Engine() any
	SetLanguage(lang Language) error
}
