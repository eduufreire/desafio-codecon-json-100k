package main

import (
	"log"
	"net/http"

	"github.com/eduufreire/desafio-codecon/internal/config"
	"github.com/eduufreire/desafio-codecon/internal/handler"
)

func main() {

	db := config.MemoryDatabase{}
	db.Init()

	h := handler.Handler{}
	h.Init(&db)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", h.Save)
	mux.HandleFunc("GET /superusers", h.GetSuperusers)
	mux.HandleFunc("GET /top-countries", h.GetTopCountries)
	mux.HandleFunc("GET /team-insights", h.GetTeamInsights)
	mux.HandleFunc("GET /active-users-per-day", h.GetActiveUsersPerDay)
	mux.HandleFunc("GET /", h.GetAll)

	log.Fatal(http.ListenAndServe(":3333", mux))
}
