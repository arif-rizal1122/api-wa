package request


type RequestCreateStatus struct {
    Caption string `form:"caption" binding:"required"`
    Picture string // No need to bind or validate this field directly in the struct
}



type RequestUpdateStatus struct {
    Caption string `form:"caption" binding:"required"`
    Picture string // No need to bind or validate this field directly in the struct
}

