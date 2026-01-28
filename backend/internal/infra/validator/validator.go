package validator

import (
	"fmt"

	validatorLib "github.com/go-playground/validator/v10"
)

type GoPlaygroundValidator struct {
    validator *validatorLib.Validate
}

func NewGoPlaygroundValidator() *GoPlaygroundValidator {
    return &GoPlaygroundValidator{
        validator: validatorLib.New(),
    }
}

func (g *GoPlaygroundValidator) ValidateStruct(s interface{}) error {
	return g.validator.Struct(s)
}

func (g *GoPlaygroundValidator) FormatValidationErrors(err error) map[string]interface{} {
    errors := make(map[string]interface{})
    
    if validationErrors, ok := err.(validatorLib.ValidationErrors); ok {
        for _, ve := range validationErrors {
            errors[ve.Field()] = fmt.Sprintf("Field '%s' failed on the '%s' tag.", 
                                              ve.Field(), 
                                              ve.Tag())
        }
    } else {
        errors["general"] = "Unknown validation error."
    }
    return errors
}