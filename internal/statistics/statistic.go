package statistics

import (
	"fmt"
	"log"
	"net/http"
)

var requestCount int

func StartStatisticsService() {
	http.HandleFunc("/stats", handleStats)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleStats(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Total requests: %d", requestCount)))
}
func UpdateStatistics() {
	requestCount++
}