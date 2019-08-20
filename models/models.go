package models

type Doc struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Response ...
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    interface{}
}
