package models

type Tag struct {
	ID   int    `json:"id" gorm:"primarykey"`
	Name string `json:"name" gorm:"not null"`
}

type TagResponseWithPost struct {
	ID    int            `json:"id"`
	Name  string         `json:"name"`
	Posts []PostResponse `json:"posts" gorm:"many2many:post_tags;ForeignKey:ID;joinForeignKey:TagID;References:ID;joinReferences:PostID"`
}

func (TagResponseWithPost) TableName() string {
	return "tags"
}
