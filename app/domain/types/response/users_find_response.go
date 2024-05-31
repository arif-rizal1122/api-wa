package response


type PayloadFind struct {
	Message     string   `json:"message"`
	Status      int		 `json:"status"`
	Data        ResponseFind  `json:"data"`
}

type PayloadFinds struct {
    Message string  `json:"message"`
    Status  int     `json:"status"`
    Datas   []ResponseFinds  `json:"data"`
}

type ResponseFind struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type ResponseFinds struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}


func NewResponseFind(message string, status int, rqs ResponseFind) PayloadFind {
	return PayloadFind{
		Message: message,
		Status: status,
		Data: rqs,
	}
}


func NewResponseFinds(message string, status int, rqs []ResponseFinds) []PayloadFinds {
    return []PayloadFinds{
        {
            Message: message,
            Status:  status,
            Datas:   rqs,
        },
    }
}