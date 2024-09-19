package service

import (
	"NewProj1/internal/config"
	"NewProj1/repositories/postgres"
	"strconv"
	"strings"
	"time"

	"NewProj1/internal/service/utils"
	"fmt"
)


func TextProcessing(password, chatID string, cfgDB config.Postgres) string {
	sendMess := handlersDataForTable(chatID, cfgDB)
	if sendMess {
		for key, value := range utils.MapHash {
			if utils.GetMd5(password) == key {
				return value
			} else if password == value {
				return key
			}
		}
		return fmt.Sprint("Password/hash \"", password, "\" does not exist")
	} else {
		return "Query time has expired"
	}
}

func handlersDataForTable(chatID string, cfgDB config.Postgres) bool {
	var id int = postgres.ReturnId("max", "id", cfgDB)
	now := time.Now()
	hour, minute := fmt.Sprint(now.Hour()), fmt.Sprint(now.Minute())
	postgres.InitDB(id+1, chatID, fmt.Sprintf("%s:%s", hour, minute), cfgDB)

	var sendMess bool = timeMess(now, cfgDB)
	if sendMess {
		return true
	} else {
		return false
	}
}


func timeMess(now time.Time, cfgDB config.Postgres) bool {
	var id_min int = postgres.ReturnId("min", "id", cfgDB)
	var id_max int = postgres.ReturnId("max", "id", cfgDB)
	var time_min []string = strings.Split((postgres.ReturnVal("min", "timemessage", cfgDB)), ":")
	var time_max []string = strings.Split((postgres.ReturnVal("max", "timemessage", cfgDB)), ":")

	hour_min, _ := strconv.Atoi(time_min[0])
	minute_min, _ := strconv.Atoi(time_min[1])
	hour_max, _ := strconv.Atoi(time_max[0])
	minute_max, _ := strconv.Atoi(time_max[1])

	now_hour, _ := strconv.Atoi(fmt.Sprint(now.Hour()))
	if id_max-id_min >= 10 && (hour_max-hour_min >= 1 || (hour_max == hour_min-1 && minute_max == minute_min)) {
		return true
	} else if id_max-id_min >= 10 {
		return false
	}
	if ((now_hour-hour_max == 1 && minute_max == minute_min) || now_hour-hour_max >= 1) && id_max > 10 {
		postgres.ClearTable(cfgDB)
	}
	return true
}
