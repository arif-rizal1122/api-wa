package response


type UpdateResponse struct {
	// fields here
	Status       int      `json:"status"`
	Message      string	  `json:"message"`
}


func ApiUpdateResponse(status int, message string) *UpdateResponse {
	return &UpdateResponse{
		Status: status, 
		Message: message,
	}
}