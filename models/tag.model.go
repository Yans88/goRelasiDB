package models

type Tag struct {
	ID   int    `json:"id" gorm:"primarykey"`
	Name string `json:"name" gorm:"not null"`
}
