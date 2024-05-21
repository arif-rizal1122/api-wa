package helper

type Meta struct {
	Message interface{} `json:"message"`
	Status  int         `json:"status"`
	Errors  interface{} `json:"errors"`
}

func NewErrorsResponse(message interface{}, status int, errors interface{}) Meta {
	return Meta{
		Message: message,
		Status:  status,
		Errors:  errors,
	}
}
