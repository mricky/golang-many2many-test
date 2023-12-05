package post

type Service interface {
	CreatePost(input PostInput) (Post, error)
	Update(postId int, input PostInput) (Post, error)
	Delete(postId int) int64
	FindAll() ([]Post, error)
	FindById(postId int) ([]Post, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Delete(postId int) int64 {

	result := s.repository.Delete(postId)

	return result
}
func (s *service) Update(postId int, input PostInput) (Post, error) {
	post := Post{}
	post.TITLE = input.TITLE
	post.CONTENT = input.CONTENT
	post.Tags = input.TAGS
	posts, err := s.repository.Update(postId, post)

	if err != nil {
		return posts, err
	}

	return posts, nil
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
	post.Tags = input.TAGS

	newPost, err := s.repository.Save(post)

	if err != nil {
		return newPost, err
	}

	return newPost, nil

}
