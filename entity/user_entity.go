// Package entity is describe a table in database
package entity

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Balance  string `json:"balance"`
}
