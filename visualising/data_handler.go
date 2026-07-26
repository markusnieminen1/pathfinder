package visualising

import (
	"encoding/json"
	"net/http"
	"pathfinder/data"
)

func EventsHandler(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(&data.Events)

	if err != nil {
		http.Error(w, "failed to write events", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

}
