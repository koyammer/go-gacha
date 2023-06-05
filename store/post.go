package store

import (
	"errors"
	"go-echo-starter/model"

	"gorm.io/gorm"
)

type PostStore struct {
	db *gorm.DB
}

func NewPostStore(db *gorm.DB) *PostStore {
	return &PostStore{db: db}
}

func (us *UserStore) AllPosts() ([]model.Post, error) {
	var p []model.Post
	err := us.db.Preload("User").Find(&p).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return p, nil
		}
		return nil, err
	}
	return p, nil
}
func (ps *PostStore) AllPostsA() ([]model.Post, error) {
	var p []model.Post
	err := ps.db.Find(&p).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return p, nil
		}
		return nil, err
	}
	return p, nil
}

func (us *UserStore) UpdatePost(post *model.Post) error {
	return us.db.Model(post).Updates(post).Error
}
