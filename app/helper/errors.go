package helper

type Meta struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Errors  interface{} `json:"errors"`
}

func NewErrorsResponse(message string, status int, errors interface{}) Meta {
	return Meta{
		Message: message,
		Status:  status,
		Errors:  errors,
	}
}
