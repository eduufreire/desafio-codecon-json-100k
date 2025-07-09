package repository

import "github.com/eduufreire/desafio-codecon/internal/model"

type UserRepository struct {
	users []model.User
}

func (ur *UserRepository) Init() {
	ur.users = make([]model.User, 0)
}

func (ur *UserRepository) SaveUser(userData *model.User) {
	ur.users = append(ur.users, *userData)
}

func (ur *UserRepository) GetAllUsers() *[]model.User {
	return &ur.users
}