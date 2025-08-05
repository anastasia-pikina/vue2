package models

type Block struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Code    string `json:"code"`
	Title   string `json:"title"`
	Image   string `json:"image"`
}
