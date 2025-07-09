package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/eduufreire/desafio-codecon/internal/model"
)

func (h *Handler) GetSuperusers(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	allUsers := h.db.UserRepository.GetAllUsers()

	filteredUsers := make([]model.User, 0)

	for _, user := range *allUsers {
		if user.Score >= 900 && user.Active {
			filteredUsers = append(filteredUsers, user)
		}
	}

	response := ApiResponse{
		Timestamp:     time.Now().Format(time.RFC3339),
		ExecutionTime: int(time.Since(start).Milliseconds()),
		Data:          filteredUsers,
	}

	result, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
