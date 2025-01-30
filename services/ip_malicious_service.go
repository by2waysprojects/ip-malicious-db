package services

type MaliciousIpService struct {
	Neo4jService    *Neo4jService
	GithubIpService *GithubIpService
}

func NewMaliciousIpService(dbService *Neo4jService, githubIpService *GithubIpService) *MaliciousIpService {
	return &MaliciousIpService{Neo4jService: dbService, GithubIpService: githubIpService}
}

func (s *MaliciousIpService) SaveGithubMaliciousIp(limit int) error {
	countryIps, err := s.GithubIpService.FetchAllCountryIPs()
	if err != nil {
		return err
	}

	if err := s.Neo4jService.SaveMaliciousIps(countryIps, limit); err != nil {
		return err
	}
	return nil
}
