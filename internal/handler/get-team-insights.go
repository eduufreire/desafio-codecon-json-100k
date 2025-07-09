package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type teamInsights struct {
	Name              string
	Leaders           int
	TotalMembers      int
	ActiveMembers     int
	CompletedProjects int
	ActivePercentage  int
}

type TeamInsightsResponse struct {
	Name              string  `json:"name"`
	Leaders           int     `json:"leaders"`
	TotalMembers      int     `json:"total_members"`
	CompletedProjects int     `json:"completed_projects"`
	ActivePercentage  float64 `json:"active_percentage"`
}

func (h *Handler) GetTeamInsights(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	allUsers := h.db.UserRepository.GetAllUsers()
	teams := h.db.TeamRepository.GetAllTeams()

	infoTeams := make(map[int]teamInsights)
	for _, team := range *teams {
		id := h.db.TeamRepository.GetTeamId(team.Name)
		infoTeams[id] = teamInsights{
			Name:              team.Name,
			Leaders:           0,
			TotalMembers:      0,
			ActiveMembers:     0,
			CompletedProjects: 0,
			ActivePercentage:  0,
		}

		for _, project := range team.Projects {
			if project.Completed {
				t := infoTeams[id]
				t.CompletedProjects += 1
				infoTeams[id] = t
			}
		}
	}

	for _, user := range *allUsers {
		t := infoTeams[user.TeamId]
		t.TotalMembers += 1

		if user.Active {
			t.ActiveMembers += 1
		}

		if user.IsLeader {
			t.Leaders += 1
		}
		infoTeams[user.TeamId] = t
	}

	parsedTeamInsights := make([]TeamInsightsResponse, 0)
	for _, team := range infoTeams {
		parsedTeamInsights = append(parsedTeamInsights, TeamInsightsResponse{
			Name:              team.Name,
			Leaders:           team.Leaders,
			TotalMembers:      team.TotalMembers,
			CompletedProjects: team.CompletedProjects,
			ActivePercentage:  float64(team.ActiveMembers * 100 / team.TotalMembers),
		})
	}

	response := ApiResponse{
		Timestamp:     time.Now().Format(time.RFC3339),
		ExecutionTime: int(time.Since(start).Milliseconds()),
		Data:          parsedTeamInsights,
	}

	result, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)

}
