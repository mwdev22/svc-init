package models

// models package simplifies communication between repository, service and transport layers
// it contains all the models used in the application, we can easily pass them around

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
