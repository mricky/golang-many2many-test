package post

import "time"

type Post struct {
	ID        int
	TITLE     string
	CONTENT   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
