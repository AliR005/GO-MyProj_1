package postgres

import (
	"NewProj1/internal/domain/models"
	"context"
	"fmt"
)


func (p *postgres) CreateTable(ctx context.Context) error{
	_, err := p.db.Exec(`create table if not exists limitchat (
	ID integer primary key,
	chatID VARCHAR(20),
	timeMessage time
	)`)
	
	if err != nil {
		return err
	}

	return nil
}

func (p *postgres) AppendMessage(ctx context.Context, user *models.User) error{
    _, err := p.db.Exec(`insert into limitchat (id, chatid, timemessage)
    values ($1, $2, $3)`, user.ID, user.ChatID, user.MessageTime)
    
	if err != nil {
		return err
	}

	return nil
}


func (p *postgres) ReturnIdMin(ctx context.Context, user *models.User) int{
    res := p.db.QueryRow(fmt.Sprintf(`select %s(%s) from limitchat`, "min", "id"))
    var id int
    res.Scan(&id)
    return id
}

func (p *postgres) ReturnIdMax(ctx context.Context, user *models.User) int{
    res := p.db.QueryRow(fmt.Sprintf(`select %s(%s) from limitchat`, "max", "id"))
    var id int
    res.Scan(&id)
    return id
}
