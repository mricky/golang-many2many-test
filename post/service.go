package post

type Service interface {
	CreatePost(input PostInput) (Post, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreatePost(input PostInput) (Post, error) {
	post := Post{}
	post.ID = 6
	post.TITLE = input.TITLE
	post.CONTENT = input.CONTENT
	newPost, err := s.repository.Save(post)

	if err != nil {
		return newPost, err
	}

	return newPost, nil

}
