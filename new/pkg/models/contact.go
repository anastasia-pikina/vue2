package models

type Contact struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	Address string `json:"address"`
}
