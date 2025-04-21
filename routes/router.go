package routes

import (
	"ip-malicious-db/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, maliciousIpController *controllers.MaliciousIpController) {
	router.HandleFunc("/save-malicious-ip", func(w http.ResponseWriter, r *http.Request) {
		err := maliciousIpController.LoadMaliciousIps(w, r)
		if err != nil {
			http.Error(w, "Failed to save malicious ips", http.StatusInternalServerError)
		}
	}).Methods("GET")
}

func Alive(router *mux.Router, maliciousIpController *controllers.MaliciousIpController) {
	router.HandleFunc("/health-module", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")
}
