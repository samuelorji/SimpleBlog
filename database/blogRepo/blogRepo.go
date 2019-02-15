package blogRepo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/samuelorji/simpleblog/database"
	"github.com/samuelorji/simpleblog/model"
	"strconv"
)

type mysqlBlogRepo struct {
	db *sql.DB
}

func MySqlBlogRepoObject(conn *sql.DB) database.BlogInterface {
	return &mysqlBlogRepo{conn}
}

func check(err error) error {
	fmt.Println(err)
	return err
}

func (m *mysqlBlogRepo) FetchAllBlogs() ([]*model.Blog, error) {

	queryString := "select * from blog"
	rows, err := m.db.Query(queryString)
	if err != nil {
		return nil, check(err)
	}

	defer rows.Close()

	blogs := []*model.Blog{}

	for rows.Next() {

		b := new(model.Blog)
		err := rows.Scan(
			&b.Id,
			&b.Title,
			&b.Content,
			&b.Image,
			&b.Created,
		)
		if err != nil {
			return nil, check(err)
		}

		blogs = append(blogs, b)
	}
	return blogs, nil

}

func (m *mysqlBlogRepo) FetchBlog(id int) (*model.Blog, error) {

	//idString := strconv.Itoa(id)
	queryString := fmt.Sprintf("SELECT * from blog where id = %d", id)
	rows, err := m.db.Query(queryString)
	if err != nil {
		return nil, check(err)
	}
	defer rows.Close()

	for rows.Next() {
		b := new(model.Blog)
		err := rows.Scan(
			&b.Id,
			&b.Title,
			&b.Content,
			&b.Created,
			&b.Image,
		)

		if err != nil {
			return nil, check(err)
		}
		return b,nil
	}

	return nil, nil

}

func (m *mysqlBlogRepo) AddBlog(blog model.Blog) error {
	queryString := "INSERT INTO blog(title,content,image) VALUES(?,?,?)"
	stmt, err := m.db.Prepare(queryString)
	if err != nil {
		check(err)
	}

	defer stmt.Close()

	rows, er := stmt.Exec(blog.Title, blog.Content, blog.Image)
	if er != nil {
		check(er)
	}
	num_rows_affected , _  := rows.RowsAffected()
	if(num_rows_affected == 0){
		return errors.New("Unable to Insert Blog")
	}
	return nil

}

func (m *mysqlBlogRepo) UpdateBlog(blog model.Blog) error {
	idString := strconv.Itoa(blog.Id)
	queryString := "UPDATE blog SET title=?,content=?,image=? WHERE id=?"
	stmt, err := m.db.Prepare(queryString)
	if err != nil {
		check(err)
	}

	defer stmt.Close()

	rows, er := stmt.Exec(blog.Title, blog.Content, blog.Image, idString)

	if er != nil {
		check(er)
	}
	num_rows_affected , _  := rows.RowsAffected()
	if(num_rows_affected == 0){
		return errors.New("Blog does not exist")
	}
	return nil

}

func (m *mysqlBlogRepo) DeleteBlog(id int) error {

	queryString := "delete from blog where id=?"
	stmt, err := m.db.Prepare(queryString)
	if err != nil {
		check(err)
	}

	defer stmt.Close()

	_, er := stmt.Exec(id)
	if err != nil {
		check(er)
	}
	return nil
}
