package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type LogsResponse struct {
	Date  string `json:"date"`
	Total int    `json:"total"`
}

func (h *Handler) GetActiveUsersPerDay(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	allLogs := h.db.LogsRepository.GetAllLogs()

	logsPerDay := make(map[string]int, 0)
	for _, log := range *allLogs {
		_, exist := logsPerDay[log.Date]
		if !exist {
			logsPerDay[log.Date] = 0
		}

		if log.Action == "login" {
			logsPerDay[log.Date] += 1
		}
	}

	parsedResponse := make([]LogsResponse, 0)
	for key, value := range logsPerDay {
		parsedResponse = append(parsedResponse, LogsResponse{Date: key, Total: value})
	}

	response := ApiResponse{
		Timestamp:     time.Now().Format(time.RFC3339),
		ExecutionTime: int(time.Since(start).Milliseconds()),
		Data:          parsedResponse,
	}

	result, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)

}
