package validator

type StructValidator interface {
	ValidateStruct(any) error
	// Var validates a single variable using tag style validation.
	// eg.
	// var i int
	// validate.Var(i, "gt=1,lt=10")
	Var(field interface{}, tag string) error
	Engine() any
	SetLanguage(lang Language) error
}
