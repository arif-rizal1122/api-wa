package response



type PayloadStatusCreate struct {
	Message            string       `json:"message"`
	Status             int			`json:"status"`
	Data            StatusCreateResponse
}

type StatusCreateResponse struct {
	Picture               string
	Caption               string
	UserId                int
}


func NewStatusCreateResponse(message string, status int, rqs StatusCreateResponse) *PayloadStatusCreate {
	return &PayloadStatusCreate{
		Message: message,
		Status: status,
	    Data: rqs,
	}
}