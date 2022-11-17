package models

type Locker struct {
	ID     int          `json:"id" form:"id" gorm:"primaryKey"`
	Code   string       `json:"code" form:"code" gorm:"unique;not null"`
	UserID int          `json:"user_id" form:"user_id" gorm:"not null"`
	User   UserResponse `json:"user"`
}

type LockerResponse struct {
	ID     int    `json:"id""`
	Code   string `json:"code"`
	UserID int    `json:"-"`
}

func (LockerResponse) TableName() string {
	return "lockers"
}
