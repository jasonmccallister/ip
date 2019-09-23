package ip

import (
	"encoding/json"
	"net/http"

	"github.com/jasonmccallister/ip/internal"
)

// Handler is the endpoint that creates the serverless function for Zeit
func Handler(w http.ResponseWriter, r *http.Request) {
	resp := internal.IPResponse{
		IP: internal.GetIP(r),
	}

	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
