package ip

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type response struct {
	IP string `json:"ip,omitempty"`
}

// Handler is the endpoint that creates the serverless function for Zeit
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Print("IP microservice started")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := header(r)
		resp := response{IP: ip}

		js, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func header(r *http.Request) string {
	switch r.Header.Get("X-Forwarded-For") {
	case "":
		return r.RemoteAddr
	}

	return r.Header.Get("X-Forwarded-For")
}
