package services

import (
	"context"
	services "ip-malicious-db/services/model"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Neo4jService struct {
	Driver neo4j.DriverWithContext
	Limit  int
}

func NewNeo4jService(uri, username, password string) *Neo4jService {
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Fatalf("Failed to create Neo4j driver: %v", err)
	}
	return &Neo4jService{Driver: driver}
}

func (s *Neo4jService) Close() {
	s.Driver.Close(context.Background())
}

func (s *Neo4jService) SaveMaliciousIps(countryIps map[string][]string, limit int) error {
	ctx := context.Background()
	processed := 0

	session := s.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	for country, ips := range countryIps {
		for _, ip := range ips {

			if processed >= limit {
				break
			}

			malicous_ip := services.MaliciousIpNeo4j{ID: ip, Country: country, Action: services.Alert}
			query := `
			CREATE (i:Ip {
				id: $id,
				country: $country
			})
		`
			// Execute the query
			_, err := session.Run(ctx, query, map[string]interface{}{
				"id":      string(malicous_ip.ID),
				"country": string(malicous_ip.Country),
				"action":  string(malicous_ip.Action),
			})
			if err != nil {
				log.Printf("Error inserting record from Ip %s: %v", ip, err)
			}

			processed++
		}

		if processed >= limit {
			break
		}
	}

	return nil
}
