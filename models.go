package main

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type newItem struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Date_create string `json:"date_create"`
	Image       string `json:"image"`
}

type contactItem struct {
	Id      string `json:"id"`
	Address string `json:"address"`
}

type phone struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

type NullString struct {
	String string
	Valid  bool // Valid is true if String is not NULL
}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//type Newer interface {
//	TableName() string
//}
//type New struct {
//	gorm.Model
//	ID   int `gorm:"size:255"`
//	NAME string
//}
