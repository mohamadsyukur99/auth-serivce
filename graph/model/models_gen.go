// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Status   *bool  `json:"status"`
}
