package handler

type BodyParser struct {
	Name  string
	Grade int
}
type User struct {
	Name  string `json:"name" binding:"required"`
	Grade int
}
