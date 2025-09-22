package models

type New struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	DateCreate  string `json:"date_create"`
	Image       string `json:"image"`
	IsActive    bool   `json:"is_active" gorm:"default:1"`
}
