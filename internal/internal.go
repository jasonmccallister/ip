package internal

import "net/http"

// IPResponse is the struct that contains the IP response
type IPResponse struct {
	IP string `json:"ip,omitempty"`
}

// GetIP will return the IP based on the header or remote address
func GetIP(r *http.Request) string {
	switch r.Header.Get("X-Forwarded-For") {
	case "":
		return r.RemoteAddr
	}

	return r.Header.Get("X-Forwarded-For")
}
