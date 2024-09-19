package models

type User struct {
	ID     		int
	ChatID 		string
	MessageTime string
}

func New(id int, chatID, messTime string) *User {
	return &User{
		ID:          id,
		ChatID: 	 chatID,
		MessageTime: messTime,
	}
}
