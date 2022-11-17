package models

type User struct {
	ID     int            `json:"id" form:"id" gorm:"primaryKey"`
	Name   string         `json:"name" form:"name" gorm:"not null"`
	Locker LockerResponse `json:"locker"`
	Posts  []Post         `json:"posts"`
}

type UserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (UserResponse) TableName() string {
	return "users"
}
