package controllers

import (
	"fmt"
	"ip-malicious-db/services"
	"log"
	"net/http"
)

type MaliciousIpController struct {
	MaliciousIpService *services.MaliciousIpService
}

func NewMaliciousIpController(maliciousIpService *services.MaliciousIpService) *MaliciousIpController {
	return &MaliciousIpController{MaliciousIpService: maliciousIpService}
}

func (mc *MaliciousIpController) LoadMaliciousIps(w http.ResponseWriter, r *http.Request) error {
	log.Println("Saving all malicious ips...")

	err := mc.MaliciousIpService.SaveGithubMaliciousIp()
	if err != nil {
		http.Error(w, "Failed saving ips", http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "All ips are correctly saved")
	return nil
}
