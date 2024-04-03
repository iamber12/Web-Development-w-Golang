package models

// CAPS FIRST CHARACTER OF TYPE MEANS ITS PUBLIC
type User struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	Id     string `json:"id"`
}
