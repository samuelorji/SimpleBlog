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
	serv := service.NewArticleUseCase(db)
	router.HandleFunc("/api/articles",serv.FetchAll).Methods("GET")
	router.HandleFunc("/api/article/{id}",serv.FetchBlogByID).Methods("GET")
	router.HandleFunc("/api/article",serv.AddNewBlog).Methods("POST")
	router.HandleFunc("/api/article",serv.UpdateBlog).Methods("PUT")
	router.HandleFunc("/api/article",serv.DeleteBlog).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9090", router))


}
