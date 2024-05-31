package request



type RequestCreateStatus struct {
	Picture string `json:"picture" binding:"required"`
	Caption string `json:"caption" binding:"required"`
	UserId  int    `json:"user_id"`
}




type RequestUpdateStatus struct {
	 Picture                   string	`json:"picture"`
	 Caption                   string	`json:"caption"`
}


