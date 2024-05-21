package helper



type Payload struct {
	Message string 					 `json:"message"`
	Status  int                		 `json:"status"`
	Data    ResponseUserRegister	 `json:"data"`

}



type ResponseUserRegister struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}



func NewAPIregisterResponse(status int, message string, rqs ResponseUserRegister) Payload {
	return Payload{
		Message: message,
		Status:  status,
		Data:    rqs,
	}
}




