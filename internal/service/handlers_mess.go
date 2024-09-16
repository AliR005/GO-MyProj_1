package service

import (
	"NewProj1/internal/domain/models"
	"NewProj1/internal/service/utils"
	"fmt"
	"time"
)

func TextProcessing(id int64, pass string) string {
	now := time.Now()
	datetime := fmt.Sprintf("%s:%s", fmt.Sprint(now.Hour()), fmt.Sprint(now.Minute()))

	models.New(id, pass, datetime)
	for key, val := range utils.MapHash {
		if utils.GetMd5(pass) == key {
			return val
		} else if pass == val {
			return key
		}
	}
	return fmt.Sprint("Password/hash \"", pass, "\" does not exist")
}

//func SettingsBot(password, chatID string) string {
//	sendMess := handlersDataForTable(chatID)
//	if sendMess {
//		for key, value := range utils.MapHash {
//			if utils.GetMd5(password) == key {
//				return value
//			} else if password == value {
//				return key
//			}
//		}
//		return fmt.Sprint("Password/hash \"", password, "\" does not exist")
//	} else {
//		return "Query time has expired"
//	}
//}
//
//func handlersDataForTable(chatID string) bool {
//	var id int = db.ReturnId("max", "id")
//	now := time.Now()
//	hour, minute := fmt.Sprint(now.Hour()), fmt.Sprint(now.Minute())
//	go db.InitDB(id+1, chatID, fmt.Sprintf("%s:%s", hour, minute))
//
//	var sendMess bool = testTime(now)
//	if sendMess {
//		return true
//	} else {
//		return false
//	}
//}
//
//func testTime(now time.Time) bool {
//	var id_min int = db.ReturnId("min", "id")
//	var id_max int = db.ReturnId("max", "id")
//	var time_min []string = strings.Split((db.ReturnVal("min", "timemessage")), ":")
//	var time_max []string = strings.Split((db.ReturnVal("max", "timemessage")), ":")
//
//	hour_min, _ := strconv.Atoi(time_min[0])
//	minute_min, _ := strconv.Atoi(time_min[1])
//	hour_max, _ := strconv.Atoi(time_max[0])
//	minute_max, _ := strconv.Atoi(time_max[1])
//
//	now_hour, _ := strconv.Atoi(fmt.Sprint(now.Hour()))
//	if id_max-id_min >= 10 && (hour_max-hour_min >= 1 || (hour_max == hour_min-1 && minute_max == minute_min)) {
//		return true
//	} else if id_max-id_min >= 10 {
//		fmt.Println(":::::::::::::::::", id_max, id_min)
//		return false
//	}
//	if ((now_hour-hour_max == 1 && minute_max == minute_min) || now_hour-hour_max >= 1) && id_max > 10 {
//		db.ClearTable()
//	}
//	fmt.Println("В конец")
//	return true
//}
