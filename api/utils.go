package api

import "strings"

func extractPathParams(fullURL, baseURL string) []string {
	if strings.Contains(fullURL, "?") {
		fullURL = fullURL[:strings.Index(fullURL, "?")]
	}

	paramStr := strings.Replace(fullURL, baseURL, "", 1)
	if strings.Trim(paramStr, " ") == "" {
		return nil
	}

	params := strings.Split(paramStr, "/")

	return params
}
