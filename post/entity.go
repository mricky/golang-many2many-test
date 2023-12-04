package post

import "time"

type Post struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	TITLE     string `json:"title"`
	CONTENT   string `json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Tags      []*Tag `gorm:"many2many:post_tags;"`
}

type Tag struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	LABEL     string `json:"label"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Posts     []*Post `gorm:"many2many:post_tags;"`
}
