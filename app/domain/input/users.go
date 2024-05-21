package input

type RequestUserRegister struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Phone    int    `json:"phone" binding:"required"`
}



type RequestUpdateUser struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"` // Tidak wajib untuk update
	Phone    int    `json:"phone" binding:"required"`
}




type LoginUser struct {
	// fields here
}

type ResponseUserLogin struct {
	// fields here
}

type CheckEmailUser struct {
	Email string
}

type UserResponse struct {
	// fields here
}
