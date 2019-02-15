package model

type Blog struct {

	Id      int 	`json:"id"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Created string	`json:"created"`
	Image   string  `json:"image"`
	//author  Author

}
