package response


type UpdateResponse struct {
	// fields here
	Status       int
	Message      string
}


func ApiUpdateResponse(status int, message string) *UpdateResponse {
	return &UpdateResponse{
		Status: status,
		Message: message,
	}
}