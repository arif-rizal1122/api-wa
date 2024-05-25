package request



type RequestCreateStatus struct {
     Picture               string         	`json:"picture"`
	 Caption               string			`json:"caption"`
	 UserId                int				`json:"user_id"`
}



type RequestUpdateStatus struct {
	 Picture                   string
	 Caption                   string
}


