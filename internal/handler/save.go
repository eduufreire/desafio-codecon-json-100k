package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eduufreire/desafio-codecon/internal/config"
	"github.com/eduufreire/desafio-codecon/internal/model"
)

type projectFileSchema struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type teamFileSchema struct {
	Name     string              `json:"name"`
	Leader   bool                `json:"leader"`
	Projects []projectFileSchema `json:"projects"`
}

type logsFileSchema struct {
	Date   string `json:"date"`
	Action string `json:"action"`
}

type UserFileSchema struct {
	Id      string           `json:"id"`
	Name    string           `json:"name"`
	Age     int16            `json:"age"`
	Score   int              `json:"score"`
	Active  bool             `json:"active"`
	Country string           `json:"country"`
	Team    model.Team       `json:"team"`
	Logs    []logsFileSchema `json:"logs"`
}

type Handler struct {
	db *config.MemoryDatabase
}

func (h *Handler) Init(db *config.MemoryDatabase) {
	h.db = db
}

type SaveResponse struct {
	Message     string `json:"message"`
	UserCounter int    `json:"user_count"`
}

type ApiResponse struct {
	Timestamp     string `json:"timestamp"`
	ExecutionTime int    `json:"execution_time_ms"`
	Data          any    `json:"data"`
}

func (h *Handler) Save(w http.ResponseWriter, r *http.Request) {
	file, _, _ := r.FormFile("jsonFile")

	allUsers := make([]UserFileSchema, 0)
	err := json.NewDecoder(file).Decode(&allUsers)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range allUsers {
		teamId := h.db.TeamRepository.GetTeamId(user.Team.Name)

		if teamId == -1 {
			teamId = h.db.TeamRepository.SaveTeam(user.Team)
		}

		userToModel := model.User{
			Id:       user.Id,
			Name:     user.Name,
			Age:      user.Age,
			Score:    user.Score,
			Active:   user.Active,
			Country:  user.Country,
			TeamId:   teamId,
			IsLeader: user.Team.Leader,
		}
		h.db.UserRepository.SaveUser(&userToModel)

		for _, log := range user.Logs {

			logUser := model.Log{
				Date:   log.Date,
				Action: log.Action,
				UserId: user.Id,
			}
			h.db.LogsRepository.SaveLogs(&logUser)
		}
	}

	result, _ := json.Marshal(SaveResponse{Message: "Arquivo recebico com sucesso", UserCounter: len(allUsers)})
	w.WriteHeader(200)
	w.Write(result)
}
