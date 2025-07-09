package repository

import "github.com/eduufreire/desafio-codecon/internal/model"

type LogsRepository struct {
	logs []model.Log
}

func (lr *LogsRepository) Init() {
	lr.logs = make([]model.Log, 0)
}

func (lr *LogsRepository) SaveLogs(log *model.Log) {
	lr.logs = append(lr.logs, *log)
}

func (lr *LogsRepository) GetAllLogs() *[]model.Log {
	return &lr.logs
}
