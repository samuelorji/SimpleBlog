package model

type Author struct {
	Username string		`json:"username"`
	Id       int        `json:"id"`
	Photo    string     `json:"photo"`
	Created  string		`json:"created"`

}
