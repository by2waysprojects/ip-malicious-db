package services

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	url = "https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/ipip_country/ipip_country_%s.netset"
)

var countryCodes = []string{
	"af", "ax", "al", "dz", "as", "ad", "ao", "ai", "aq", "ag", "ar", "am", "aw", "au", "at", "az",
	"bs", "bh", "bd", "bb", "by", "be", "bz", "bj", "bm", "bt", "bo", "bq", "ba", "bw", "bv", "br",
	"io", "bn", "bg", "bf", "bi", "cv", "kh", "cm", "ca", "ky", "cf", "td", "cl", "cn", "cx", "cc",
	"co", "km", "cg", "cd", "ck", "cr", "ci", "hr", "cu", "cw", "cy", "cz", "dk", "dj", "dm", "do",
	"ec", "eg", "sv", "gq", "er", "ee", "sz", "et", "fk", "fo", "fj", "fi", "fr", "gf", "pf", "tf",
	"ga", "gm", "ge", "de", "gh", "gi", "gr", "gl", "gd", "gp", "gu", "gt", "gg", "gn", "gw", "gy",
	"ht", "hm", "va", "hn", "hk", "hu", "is", "in", "id", "ir", "iq", "ie", "im", "il", "it", "jm",
	"jp", "je", "jo", "kz", "ke", "ki", "kp", "kr", "kw", "kg", "la", "lv", "lb", "ls", "lr", "ly",
	"li", "lt", "lu", "mo", "mg", "mw", "my", "mv", "ml", "mt", "mh", "mq", "mr", "mu", "yt", "mx",
	"fm", "md", "mc", "mn", "me", "ms", "ma", "mz", "mm", "na", "nr", "np", "nl", "nc", "nz", "ni",
	"ne", "ng", "nu", "nf", "mk", "mp", "no", "om", "pk", "pw", "ps", "pa", "pg", "py", "pe", "ph",
	"pn", "pl", "pt", "pr", "qa", "re", "ro", "ru", "rw", "bl", "sh", "kn", "lc", "mf", "pm", "vc",
	"ws", "sm", "st", "sa", "sn", "rs", "sc", "sl", "sg", "sx", "sk", "si", "sb", "so", "za", "gs",
	"ss", "es", "lk", "sd", "sr", "sj", "se", "ch", "sy", "tw", "tj", "tz", "th", "tl", "tg", "tk",
	"to", "tt", "tn", "tr", "tm", "tc", "tv", "ug", "ua", "ae", "gb", "us", "um", "uy", "uz", "vu",
	"ve", "vn", "vg", "vi", "wf", "eh", "ye", "zm", "zw"}

type GithubIpService struct{}

func NewGithubIpService() *GithubIpService {
	return &GithubIpService{}
}

func fetchCountryIPs(countryCode string) ([]string, error) {
	url := fmt.Sprintf(url, countryCode)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data for %s: %w", countryCode, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	var ips []string
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading data: %w", err)
		}
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") { // Ignore comments or empty lines
			ips = append(ips, line)
		}
	}
	return ips, nil
}

// fetchAllCountryIPs fetches IPs for all countries.
func (g *GithubIpService) FetchAllCountryIPs(limit int) (map[string][]string, error) {
	allIPs := make(map[string][]string)
	processed := 0

	for _, countryCode := range countryCodes {
		ips, err := fetchCountryIPs(countryCode)
		if err != nil {
			fmt.Printf("Error fetching IPs for %s: %v\n", countryCode, err)
			continue
		}
		allIPs[countryCode] = ips

		processed += len(ips)
		if processed >= limit {
			break
		}
	}
	return allIPs, nil
}
