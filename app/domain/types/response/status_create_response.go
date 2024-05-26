package response



type PayloadStatusCreate struct {
	Message            string       `json:"message"`
	Status             int			`json:"status"`
	Data            StatusCreateResponse  `json:"data"`
}

type StatusCreateResponse struct {
	Picture               string   `json:"picture"`
	Caption               string   `json:"caption"`
	UserId                int	   `json:"user_id"`	 
}


func NewStatusCreateResponse(message string, status int, rqs StatusCreateResponse) PayloadStatusCreate {
	return PayloadStatusCreate{
		Message: message,
		Status: status,
	    Data: rqs,
	}
}