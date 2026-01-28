package application

type Validator interface {
	ValidateStruct(s interface{}) error
	FormatValidationErrors(err error) map[string]interface{}
}