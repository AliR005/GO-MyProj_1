package models

type User struct {
	ID          int64
	MessageText string
	MessageTime string
}

func New(id int64, messText, messTime string) *User {
	return &User{
		ID:          id,
		MessageText: messText,
		MessageTime: messTime,
	}
}
