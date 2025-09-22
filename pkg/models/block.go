package models

type Block struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Content  string `json:"content"`
	Code     string `json:"code"`
	Title    string `json:"title"`
	Image    string `json:"image"`
	Link     string `json:"link"`
	IsActive bool   `json:"is_active" gorm:"default:1"`
}
