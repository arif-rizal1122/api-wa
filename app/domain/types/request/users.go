package request

type RequestUserRegister struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}



type RequestUpdateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}


type AuthUserLoginRequest struct {
	Password       string   `json:"password" binding:"required"`
	Email          string	`json:"email" binding:"required,email"`
}





