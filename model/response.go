package model

type Response[T any] struct {
	Ok      int    `json:"ok"`
	Message string `json:"message,omitempty"`
	Data    *T     `json:"data"`
}

type ResponseList[T any] struct {
	Ok      int    `json:"ok"`
	Message string `json:"message,omitempty"`
	Total   int    `json:"total"`
	Data    []*T   `json:"data"`
}
