package handler

type BodyParser struct {
	Name  string
	Grade int
}
type User struct {
	ID 	int `json:"id"`
	Name  string `json:"name" binding:"required"`
	Grade int	`json:"grade" binding:"required"`
}
