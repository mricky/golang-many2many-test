package post

type Service interface {
	CreatePost(input PostInput) (Post, error)
	FindAll() ([]Post, error)
	FindById(postId int) ([]Post, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindById(postId int) ([]Post, error) {
	posts, err := s.repository.FindById(postId)

	if err != nil {
		return posts, err
	}

	return posts, nil
}
func (s *service) FindAll() ([]Post, error) {
	posts, err := s.repository.FindAll()

	if err != nil {
		return posts, err
	}

	return posts, nil
}
func (s *service) CreatePost(input PostInput) (Post, error) {
	post := Post{}
	post.TITLE = input.TITLE
	post.CONTENT = input.CONTENT
	newPost, err := s.repository.Save(post)

	if err != nil {
		return newPost, err
	}

	return newPost, nil

}
