package postgres

import (
	"NewProj1/internal/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func returnERR(err error){
    if err != nil{
        panic(err)
    }
}


func connectDB (cfg config.Postgres) string{
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSL,
	)
}

func createTable(db *sql.DB){
    _, err := db.Exec(`create table if not exists limitchat (
    ID smallserial primary key,
    chatID VARCHAR(20),
    timeMessage time
    )`)

    returnERR(err)
}

func editTable(db *sql.DB, id int, chatID string, timeMess string){
    _, err := db.Exec(`insert into limitchat (id, chatid, timemessage)
    values ($1, $2, $3)`, id, chatID, timeMess)
    returnERR(err)
}

func ReturnId(method, value string, cfg config.Postgres) int{
    connStr := connectDB(cfg)
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    } 
    defer db.Close()

    res := db.QueryRow(fmt.Sprintf(`select %s(%s) from limitchat`, method, value))
    var id int
    res.Scan(&id)
    return id
}

func ReturnVal(method, value string, cfg config.Postgres) string{
    connStr := connectDB(cfg)
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    } 
    defer db.Close()

    res := db.QueryRow(fmt.Sprintf(`select %s(%s) from limitchat`, method, value))
    var timeMess string
    res.Scan(&timeMess)
    return timeMess
}

func ClearTable(cfg config.Postgres){
    connStr := connectDB(cfg)
    db, err := sql.Open("postgres", connStr)
    returnERR(err)
    defer db.Close()

    _, err = db.Exec("delete from limitchat")
    returnERR(err)
}

func InitDB(id int, chatID string, timeMess string, cfg config.Postgres){
    connStr := connectDB(cfg)
    db, err := sql.Open("postgres", connStr)
    returnERR(err)
    defer db.Close()

    createTable(db)
    editTable(db, id, chatID, timeMess)
}
