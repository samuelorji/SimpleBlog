package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/samuelorji/simpleblog/database"
	"github.com/samuelorji/simpleblog/database/blogRepo"
	"github.com/samuelorji/simpleblog/model"
	"io"
	"net/http"
	"strconv"
)

// we should get the database from the main and supply here

type ArticleService struct {
	service database.BlogInterface
}



func NewArticleUseCase(cn *sql.DB) *ArticleService {

	return &ArticleService{
		blogRepo.MySqlBlogRepoObject(cn),
	}
}

func (a *ArticleService) FetchAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type","application/json")
	blogs, err := a.service.FetchAllBlogs()
	if err != nil {
		fmt.Println(err.Error())
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(res).Encode(blogs)
}

func (a *ArticleService) FetchBlogByID(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type","application/json")
	params := mux.Vars(req)
	stringId := params["id"]
	value , _  := strconv.Atoi(stringId)
	blog, err := a.service.FetchBlog(value)
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if blog != nil {
		json.NewEncoder(res).Encode(blog)
	}else {
		http.Error(res,"Resource Not Found",http.StatusNotFound)
	}

}

func (a *ArticleService) AddNewBlog (res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type","application/json")
	var blog  model.Blog
	decodingError := json.NewDecoder(req.Body).Decode(&blog)
	defer req.Body.Close()
	if (decodingError != nil) {
	http.Error(res,decodingError.Error(),http.StatusBadRequest)
	return
	}
	err := a.service.AddBlog(blog)
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	io.WriteString(res, `{"Status" :"Added"}`)

}

func (a *ArticleService) UpdateBlog (res http.ResponseWriter, req *http.Request) {
	var blog  model.Blog
	decodingError := json.NewDecoder(req.Body).Decode(&blog)
	defer req.Body.Close()
	if (decodingError != nil) {
		http.Error(res,decodingError.Error(),http.StatusBadRequest)
		return
	}
	err := a.service.UpdateBlog(blog)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}
	io.WriteString(res,"Blog UPDATED")
}

func (a *ArticleService) DeleteBlog (res http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	id := req.FormValue("id")
	idd, _ := strconv.Atoi(id)
	err := a.service.DeleteBlog(idd)
	if(err != nil){
		http.Error(res,err.Error(),http.StatusInternalServerError)
		return
	}
	io.WriteString(res,"Blog Deleted")

}


