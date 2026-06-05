package model

type ResponseMsg struct {
	Message string `json:"message"`
}

type DemoSchema struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}