package response


type ResponseUserLogin struct {
     Email        string     		`json:"email"`
	 Token        interface{}		`json:"token"`
}
