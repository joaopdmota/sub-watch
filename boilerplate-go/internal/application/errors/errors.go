package app_errors

const (
	ERROR_UNKNOW               = "UNKNOW"
	ERROR_BAD_REQUEST          = "BAD_REQUEST"
	ERROR_UNPROCESSABLE_ENTITY = "UNPROCESSABLE_ENTITY"
	ERROR_NOT_FOUND            = "NOT_FOUND"
)

type Error struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Path    string `json:"path,omitempty"`
	Message string `json:"message,omitempty"`
}

type Errors []Error

type ErrorsResponseDTO struct {
	Errors Errors `json:"errors"`
}

func CreateErrors(errs ...Error) Errors {
	var errors Errors

	for _, err := range errs {
		errors = append(errors, err)
	}

	return errors
}
