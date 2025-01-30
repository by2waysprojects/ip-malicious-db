package controllers

import (
	"fmt"
	"ip-malicious-db/services"
	"log"
	"net/http"
	"strconv"
)

type MaliciousIpController struct {
	MaliciousIpService *services.MaliciousIpService
}

func NewMaliciousIpController(maliciousIpService *services.MaliciousIpService) *MaliciousIpController {
	return &MaliciousIpController{MaliciousIpService: maliciousIpService}
}

func (mc *MaliciousIpController) LoadMaliciousIps(w http.ResponseWriter, r *http.Request) error {
	log.Println("Saving all malicious ips...")

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 {
		limit = 10000
	}

	err := mc.MaliciousIpService.SaveGithubMaliciousIp(limit)
	if err != nil {
		http.Error(w, "Failed saving ips", http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "All ips are correctly saved")
	return nil
}
