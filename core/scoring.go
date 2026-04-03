package core

import "strings"

func Score(ua string, method string, path string) int {
	score := 0

	ua = strings.ToLower(ua)

	if strings.Contains(ua, "sqlmap") {
		score += 60
	}
	if strings.Contains(ua, "nmap") {
		score += 40
	}
	if method != "GET" && method != "POST" {
		score += 20
	}
	if strings.Contains(path, ".env") || strings.Contains(path, "wp-admin") {
		score += 30
	}

	return score
}
