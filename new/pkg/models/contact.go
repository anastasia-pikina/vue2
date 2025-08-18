package models

type Contact struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Address  string `json:"address"`
	IsActive bool   `json:"is_active" gorm:"default:1"`
}
