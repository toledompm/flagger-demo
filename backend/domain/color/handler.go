package color

import (
	"encoding/json"
	"net/http"
)

func GetColorHandler(color string) http.HandlerFunc {
	colorHandler := func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"color": color}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(response)
	}

	return colorHandler
}
