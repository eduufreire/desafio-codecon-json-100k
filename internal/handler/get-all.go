package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {

	all := h.db.TeamRepository.GetAllTeams()

	toJson, err := json.Marshal(all)
		if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(toJson)
}
