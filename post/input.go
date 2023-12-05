package post

type PostInput struct {
	TITLE   string
	CONTENT string
	TAGS    []*Tag
}
