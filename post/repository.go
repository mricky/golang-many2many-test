package post

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(post Post) (Post, error)
	FindAll() ([]Post, error)
	FindById(postID int) ([]Post, error)
	Update(postID int, post Post) (Post, error)
	Delete(postID int) int64
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(post Post) (Post, error) {
	err := r.db.Create(&post).Error
	if err != nil {
		return post, err
	}

	return post, nil
}

func (r *repository) Update(postId int, post Post) (Post, error) {

	err := r.db.Where("id = ?", postId).Updates(&post).Error

	if err != nil {
		return post, err
	}

	return post, nil
}

func (r *repository) Delete(postId int) int64 {
	var post Post
	result := r.db.Where("id = ?", postId).Delete(&post)

	return result.RowsAffected
}
func (r *repository) FindAll() ([]Post, error) {
	var posts []Post

	err := r.db.Find(&posts).Error

	if err != nil {
		return posts, err
	}

	return posts, nil

}

func (r *repository) FindById(postId int) ([]Post, error) {
	var posts []Post

	err := r.db.Where("id = ?", postId).Find(&posts).Error

	if err != nil {
		return posts, err
	}

	return posts, nil
}
