package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type countries struct {
	Name  string `json:"country"`
	Total int    `json:"total"`
}

func (h *Handler) GetTopCountries(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	allUsers := h.db.UserRepository.GetAllUsers()

	infoCountries := make(map[string]int)

	for _, user := range *allUsers {
		_, exist := infoCountries[user.Country]
		if exist {
			infoCountries[user.Country] += 1
			continue
		}
		infoCountries[user.Country] = 0
	}

	orderedCountries := make([]countries, 0)
	for key, value := range infoCountries {
		orderedCountries = append(orderedCountries, countries{Name: key, Total: value})
	}

	for i := range orderedCountries {
		for j := i + 1; j < len(orderedCountries); j++ {
			if orderedCountries[j].Total >= orderedCountries[i].Total {
				orderedCountries[i], orderedCountries[j] = orderedCountries[j], orderedCountries[i]
			}
		}
	}
	
	response := ApiResponse{
		Timestamp:     time.Now().Format(time.RFC3339),
		ExecutionTime: int(time.Since(start).Milliseconds()),
		Data:          orderedCountries[0:5],
	}
	result, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
