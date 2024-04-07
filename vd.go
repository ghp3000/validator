package validator

type StructValidator interface {
	ValidateStruct(any) error
	Engine() any
}
