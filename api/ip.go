package ip

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	IP string `json:"ip,omitempty"`
}

// Handler is the endpoint that creates the serverless function for Zeit
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Print("IP microservice started")
	ip := header(r)
	resp := response{IP: ip}

	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func header(r *http.Request) string {
	switch r.Header.Get("X-Forwarded-For") {
	case "":
		return r.RemoteAddr
	}

	return r.Header.Get("X-Forwarded-For")
}
