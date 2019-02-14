package database

import (
	"github.com/samuelorji/simpleblog/model"
)

type BlogInterface interface {
	FetchAllBlogs() ([]*model.Blog, error)
	FetchBlog(id int) (*model.Blog, error)
	AddBlog(blog model.Blog) error
	UpdateBlog(blog model.Blog) error
	DeleteBlog(id int) error
}
