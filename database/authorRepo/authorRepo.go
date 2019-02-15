package authorRepo

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/samuelorji/simpleblog/database"
	"github.com/samuelorji/simpleblog/model"
)

type mysqlAuthorRepo struct {
	conn *sql.DB
}

func check(err error) error {
	fmt.Println(err)
	return err
}

func (repo mysqlAuthorRepo) FetchAllAuthors() ([]*model.Author, error) {

	query := "SELECT * FROM author"
	rows , err := repo.conn.Query(query)
	if(err != nil){
		return nil,check(err)
	}
	authors := []*model.Author{}
	for rows.Next(){
		var author = new(model.Author)
		err := rows.Scan(
			&author.Id,
			&author.Username,
			&author.Photo,
			&author.Created,
			)
		if(err != nil){
			return nil,check(err)
		}

		authors = append(authors,author)
	}
	return authors,nil
}

func (repo mysqlAuthorRepo) FetchAuthor(id string) (*model.Author, error) {
	query := fmt.Sprintf("SELECT FROM author WHERE id = %s",id)
	rows , err := repo.conn.Query(query)
	if(err != nil){
		return nil,check(err)
	}
	var author = new(model.Author)
	for rows.Next(){
		err := rows.Scan(
			&author.Id,
			&author.Username,
			&author.Photo,
			&author.Created,
		)
		if(err != nil){
			return nil,check(err)
		}
	}
	return author,nil



}

func (repo mysqlAuthorRepo) UpdateAuthorPhoto(author *model.Author) error {
	query := "UPDATE author SET photo=? WHERE id=?"

	stmt,err := repo.conn.Prepare(query)
	if(err!= nil){
		check(err)
	}

	rows, execError := stmt.Exec(author.Photo)
	if(execError != nil){
		check(execError)
	}
	num_rows_affected , _  := rows.RowsAffected()
	if(num_rows_affected == 0){
		return errors.New("Author does not exist")
	}
	return nil

}

func (repo mysqlAuthorRepo) UpdateAuthorUsername(author *model.Author) error {
	query := "UPDATE author SET username=? WHERE id=?"

	stmt,err := repo.conn.Prepare(query)
	if(err!= nil){
		check(err)
	}

	rows, execError := stmt.Exec(author.Username,author.Id)
	if(execError != nil){
		check(execError)
	}
	num_rows_affected , _  := rows.RowsAffected()

	if(num_rows_affected == 0){
		return errors.New("Author does not exist")
	}
	return nil
}

func (repo mysqlAuthorRepo) RemoveAuthor(id string)  error {

	query := "DELETE FROM author WHERE id=?"

	stmt,err := repo.conn.Prepare(query)
	if(err!= nil){
		check(err)
	}
	rows, execError := stmt.Exec(id)
	fmt.Println(rows)
	if(execError != nil){
		check(execError)
	}
	num_rows_affected , _  := rows.RowsAffected()
	if(num_rows_affected == 0){
		return errors.New("Author does not exist")
	}
	return nil
}

func (repo mysqlAuthorRepo) AddAuthor(author *model.Author) error {
	query := "INSERT INTO author (username,photo) VALUES (?,?)"

	stmt,err := repo.conn.Prepare(query)
	if(err!= nil){
		check(err)
	}
	rows, execError := stmt.Exec(author.Username,author.Photo)
	if(execError != nil){
		check(execError)
	}
	num_rows_affected , _  := rows.RowsAffected()
	if(num_rows_affected == 0){
		return errors.New("Error inserting into Mysql")
	}
	return nil
}

func NewAuthorRepoObject(conn *sql.DB) database.AuthorInterface{
	return mysqlAuthorRepo{conn}
}

