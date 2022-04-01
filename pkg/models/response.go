package models

type Credits struct {
	Consumed         interface{} `json:"consumed"`
	RemainingBalance interface{} `json:"remaining_balance"`
}

type Response struct {
	Message interface{} `json:"message"`
	Credits Credits     `json:"credits"`
}
