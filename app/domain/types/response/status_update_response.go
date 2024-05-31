package response



type PayloadUpdateStatus struct {
	Message        string   `json:"message"`   
	Status         int      `json:"status"`
    StatusUpdateResponse    StatusUpdateResponse
}


type StatusUpdateResponse struct {
	Picture         string  `json:"picture"`
	Caption         string	`json:"caption"`
}



func NewStatusUpdateResponse(message string, status int, rqs StatusUpdateResponse) *PayloadUpdateStatus {
	return &PayloadUpdateStatus{
		Message: message,
		Status: status,
		StatusUpdateResponse: rqs,
	}
}