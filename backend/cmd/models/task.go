package models

type Task struct {
	Id          int    `json:"id"`
	UserId      int    `json:"userid"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
