package models

type Post struct {
	ID     int          `json:"id" gorm:"primaryKey"`
	Title  string       `json:"title" form:"title" gorm:"not null"`
	Body   string       `json:"body" form:"body" gorm:"not null"`
	UserID int          `json:"user_id" form:"user_id"`
	User   UserResponse `json:"user"`
	Tags   []Tag        `json:"tags" gorm:"many2many:post_tags"`
	TagsID []int        `json:"tags_id" form:"tags_id" gorm:"-"`
}

type PostResponse struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"-"`
}

type PostResponseWithTag struct {
	ID     int          `json:"id"`
	Title  string       `json:"title"`
	Body   string       `json:"body"`
	User   UserResponse `json:"user"`
	UserID int          `json:"-"`
	Tags   []Tag        `json:"tags" gorm:"many2many:post_tags;ForeignKey:ID;joinForeignKey:PostID;References:ID;joinReferences:TagID"`
}

func (PostResponse) TableName() string {
	return "posts"
}

func (PostResponseWithTag) TableName() string {
	return "posts"
}
