package database

import "github.com/samuelorji/simpleblog/model"

type AuthorInterface interface {
	FetchAllAuthors() ([]*model.Author, error)
	FetchAuthor(id string) (*model.Author,error)
	UpdateAuthorPhoto(author *model.Author) error
	UpdateAuthorUsername(author *model.Author) error
	RemoveAuthor(id string) error
	AddAuthor(author *model.Author) error

}
