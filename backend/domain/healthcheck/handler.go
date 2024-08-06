package healthcheck

import (
	"encoding/json"
	"net/http"
)

func GetHealthcheckHandler() http.HandlerFunc {
	healthcheckHandler := func(w http.ResponseWriter, r *http.Request) {
		response := map[string]bool{"ok": true}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(response)
	}

	return healthcheckHandler
}
