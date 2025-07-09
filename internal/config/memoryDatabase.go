package config

import "github.com/eduufreire/desafio-codecon/internal/repository"

type MemoryDatabase struct {
	UserRepository repository.UserRepository
	TeamRepository repository.TeamRepository
	LogsRepository repository.LogsRepository
}

func (md *MemoryDatabase) Init() {
	md.TeamRepository.Init()
	md.UserRepository.Init()
	md.LogsRepository.Init()
}
