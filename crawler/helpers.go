package crawler

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func extractBase(rawurl string) (string, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	base := "http://" + url.Hostname()
	return base, nil
}

func findRobots(base string) (string, error) {
	resp, err := http.Get(base + "/robots.txt")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
