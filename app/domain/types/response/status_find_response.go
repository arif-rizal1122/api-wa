package response


type PayloadStatusFind struct {
     Message          string  `json:"message"`
	 Status           int	  `json:"status"`
	 StatusResponseFind    StatusResponseFind  
}

type PayloadStatusFinds struct {
	Message          string  `json:"message"`
	Status           int	  `json:"status"`
	StatusResponseFinds    []StatusResponseFinds  
}


type StatusResponseFind struct {
	Picture         string
	Caption         string
}


type StatusResponseFinds struct {
	Picture         string
	Caption         string
}



func NewStatusResponseFind(message string, status int, rqs StatusResponseFind) *PayloadStatusFind {
	 return &PayloadStatusFind{
		Message: message,
		Status: status,
		StatusResponseFind: rqs,
	 }
}


func NewStatusResponseFinds(message string, status int, rqs []StatusResponseFinds) *PayloadStatusFinds {
	return &PayloadStatusFinds{
		Message:              message,
		Status:               status,
		StatusResponseFinds:  rqs,
	}
}