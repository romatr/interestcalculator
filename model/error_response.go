package model

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	ID string `json:"id"`
	Code string `json:"code"`
	Title string `json:"title"`
	Detail string `json:"detail"`
	Source string `json:"source"`
	Meta string `json:"meta"`
}
