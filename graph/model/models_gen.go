// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Event struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	Source    string `json:"source"`
	Category  string `json:"category"`
	Message   string `json:"message"`
	Severity  string `json:"severity"`
}

type NewEvent struct {
	Source   string `json:"source"`
	Category string `json:"category"`
	Message  string `json:"message"`
	Severity string `json:"severity"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
