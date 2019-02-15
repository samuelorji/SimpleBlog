package service

import (
	"database/sql"
	"encoding/json"
	"github.com/samuelorji/simpleblog/database"
	"github.com/samuelorji/simpleblog/database/authorRepo"
	"github.com/samuelorji/simpleblog/model"
	"net/http"
)


type authorService struct {
	service database.AuthorInterface
}
func NewAuthorServiceObject(cn *sql.DB) *authorService {

	return &authorService{
		authorRepo.NewAuthorRepoObject(cn),
	}

}

func (s *authorService) FetchAllAuthors(res http.ResponseWriter , req *http.Request){
	res.Header().Set("Content-Type", "application/json")

	authors, err := s.service.FetchAllAuthors()

	if(err != nil){
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(res).Encode(authors)

}

func (s *authorService) FetchAuthorById(res http.ResponseWriter , req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	req.ParseForm()
	id := req.FormValue("id")
	author, err := s.service.FetchAuthor(id)
	if(err != nil){
		if(err.Error() == "Blog does not exist"){
			http.Error(res, err.Error(), http.StatusNotFound)
		}else {
			http.Error(res, "", http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(res).Encode(author)

}

func (s *authorService) UpdateAuthorPhoto(res http.ResponseWriter , req *http.Request) {

	var author = new(model.Author)
	err := json.NewDecoder(req.Body).Decode(&author)

	if(err != nil){
		http.Error(res,err.Error(), http.StatusBadRequest)
		return
	}

	updateError := s.service.UpdateAuthorPhoto(author)
	if(updateError != nil){
		if(updateError.Error() == "Author does not exist"){
			http.Error(res, err.Error(), http.StatusNotFound)
		}else {
			http.Error(res, "", http.StatusInternalServerError)
		}
		return
	}
}
func (s *authorService) UpdateAuthorUsername(res http.ResponseWriter , req *http.Request) {

	var author = new(model.Author)
	err := json.NewDecoder(req.Body).Decode(&author)

	if(err != nil){
		http.Error(res,err.Error(), http.StatusBadRequest)
		return
	}

	updateError := s.service.UpdateAuthorUsername(author)
	if(updateError != nil){
		if(updateError.Error() == "Author does not exist"){
			http.Error(res, updateError.Error(), http.StatusNotFound)
		}else {
			http.Error(res, "", http.StatusInternalServerError)
		}
		return
	}
}

func (s *authorService) RemoveAuthor(res http.ResponseWriter , req *http.Request) {


	req.ParseForm()

	id := req.FormValue("id")
	err := s.service.RemoveAuthor(id)
	if(err != nil){
		if(err.Error() == "Author does not exist"){
			http.Error(res, err.Error(), http.StatusNotFound)
		}else {
			http.Error(res, "", http.StatusInternalServerError)
		}
		return
	}
}

func (s *authorService) AddAuthor(res http.ResponseWriter , req *http.Request) {

	var author = new(model.Author)
	err := json.NewDecoder(req.Body).Decode(&author)

	if(err != nil){
		http.Error(res,err.Error(), http.StatusBadRequest)
		return
	}

	insertError := s.service.AddAuthor(author)
	if(insertError != nil){
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}




