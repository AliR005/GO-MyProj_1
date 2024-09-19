package models

type User struct {
	ID     		int64
	ChatID 		string
	MessageTime string
}

func New(id int64, chatID, messTime string) *User {
	return &User{
		ID:          id,
		ChatID: 	 chatID,
		MessageTime: messTime,
	}
}
