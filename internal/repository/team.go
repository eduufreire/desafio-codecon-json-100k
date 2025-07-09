package repository

import "github.com/eduufreire/desafio-codecon/internal/model"

type TeamRepository struct {
	teams []model.Team
}

func (tr *TeamRepository) Init() {
	tr.teams = make([]model.Team, 0)
}

func (tr *TeamRepository) SaveTeam(teamData model.Team) int {
	tr.teams = append(tr.teams, teamData)
	idTeam := len(tr.teams) - 1
	return idTeam
}

func (tr *TeamRepository) GetAllTeams() *[]model.Team {
	return &tr.teams
}

func (tr *TeamRepository) GetTeamId(name string) int {
	for id, team := range tr.teams {
		if team.Name == name {
			return id
		}
	}
	return -1
}
