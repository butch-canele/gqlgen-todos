// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type DeleteTodo struct {
	TodoID string `json:"todoId"`
}

type GetTodo struct {
	UserID string `json:"userId"`
}

type Mutation struct {
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Query struct {
}

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
