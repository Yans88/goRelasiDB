package models

type Post struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title" form:"title" gorm:"not null"`
	Body   string `json:"body" form:"body" gorm:"not null"`
	UserID int    `json:"user_id" form:"user_id"`
}
