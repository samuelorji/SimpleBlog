package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/samuelorji/simpleblog/service"
	"log"
	"net/http"
)

var db *sql.DB
func init() {

	database, err := sql.Open("mysql", "root:awesome0@tcp(127.0.0.1:3306)/blog?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	db = database
}

func main() {

	router := mux.NewRouter()
	articleService  := service.NewArticleUseCase(db)
	authorService   := service.NewAuthorServiceObject(db)
	router.HandleFunc("/api/articles",articleService.FetchAll).Methods("GET")
	router.HandleFunc("/api/article/{id}",articleService.FetchBlogByID).Methods("GET")
	router.HandleFunc("/api/article/add",articleService.AddNewBlog).Methods("POST")
	router.HandleFunc("/api/article/update",articleService.UpdateBlog).Methods("PUT")
	router.HandleFunc("/api/article/delete",articleService.DeleteBlog).Methods("DELETE")

	//author routes
	router.HandleFunc("/api/authors",authorService.FetchAllAuthors).Methods("GET")
	router.HandleFunc("/api/authors",authorService.FetchAuthorById).Methods("GET")
	router.HandleFunc("/api/authors/add",authorService.AddAuthor).Methods("POST")
	router.HandleFunc("/api/authors/update",authorService.UpdateAuthorUsername).Methods("PUT")
	router.HandleFunc("/api/authors/delete",authorService.RemoveAuthor).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9090", router))


}
