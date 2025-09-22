package models

import "gorm.io/gorm"

type Phone struct {
	gorm.Model
	Id        int    `json:"id" gorm:"primaryKey"`
	Addressid int    `json:"addressid"`
	Value     string `json:"value"`
	IsActive  bool   `json:"is_active" gorm:"default:1"`
}
