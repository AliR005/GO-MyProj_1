package service

import (
	"NewProj1/internal/config"
	"NewProj1/internal/domain/models"
	"NewProj1/repositories/postgres"
	"strconv"
	"time"

	"NewProj1/internal/service/utils"
	"fmt"
)

var userNow *models.User
var now time.Time

func TextProcessing(chatID, pass string, cfgDB config.Postgres) string {
	now = time.Now()
	userNow = models.New(postgres.ReturnId("max", "id", cfgDB), chatID, fmt.Sprintf("%v:%v", now.Hour(), now.Minute()))

	postgres.InitDB(userNow.ID+1, userNow.ChatID, userNow.MessageTime, cfgDB)
	sendMess := timeMess(now, cfgDB)
	
	if sendMess {
		for key, value := range utils.MapHash {
			if utils.GetMd5(pass) == key {
				return value
			} else if pass == value {
				return key
			}
		}
		return fmt.Sprint("Password/hash \"", pass, "\" does not exist")
	} else {
		return "Query time has expired"
	}
}

func timeMess(now time.Time, cfgDB config.Postgres) bool {
	var id_max int = postgres.ReturnId("max", "id", cfgDB)
	parsedTimeMin, err_1 := time.Parse(time.RFC3339, postgres.ReturnVal("min", "timemessage", cfgDB))
	if err_1 != nil{
		fmt.Println("Error in process time!!!")
		postgres.ClearTable(cfgDB)
	}

	hour_min := parsedTimeMin.Hour()
	minute_min := parsedTimeMin.Minute()


	hour_now, _ := strconv.Atoi(fmt.Sprint(now.Hour()))
	minute_now, _ := strconv.Atoi(fmt.Sprint(now.Minute()))

	if id_max > 10 || (hour_now - hour_min > 1 || (hour_now - hour_min == 1 && minute_now >= minute_min)){
		postgres.ClearTable(cfgDB)
	} else if id_max <= 10 {
		return true
	} else {
		return false
	}

	return true
}
