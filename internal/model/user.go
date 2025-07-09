package model

type User struct {
	Id      string
	Name    string
	Age     int16
	Score   int
	Active  bool
	Country string
	IsLeader bool
	TeamId  int
}
