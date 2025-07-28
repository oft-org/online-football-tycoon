package player

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func GetRandomNameByCountry(country string) (string, string, error) {
	baseURL := "https://randomuser.me/api/"
	params := url.Values{}

	apiKey := os.Getenv("RANDOMAPI_KEY")
	apiRef := os.Getenv("RANDOMAPI_REF")

	params.Add("nat", country)
	params.Add("gender", "male")
	params.Add("key", apiKey)
	params.Add("ref", apiRef)
	params.Add("results", "1")
	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		log.Printf("Error creating HTTP request: %v", err)
		return "", "", err
	}
	req.Header.Set("User-Agent", "Go-http-client/1.1")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error making HTTP request: %v", err)
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code: %d", resp.StatusCode)
		return "", "", fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return "", "", err
	}

	if len(body) == 0 || body[0] != '{' {
		log.Printf("Non-JSON response (first 200 chars): %s", string(body[:min(200, len(body))]))
		return "", "", fmt.Errorf("non-JSON response")
	}

	log.Printf("API response: %s", string(body))

	var userName RandomUserName
	if err := json.Unmarshal(body, &userName); err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return "", "", err
	}

	if len(userName.Results) == 0 {
		return "", "", fmt.Errorf("no results in API response")
	}

	firstName := userName.Results[0].Name.First
	lastName := userName.Results[0].Name.Last

	return firstName, lastName, nil
}

type RandomUserName struct {
	Results []struct {
		Name struct {
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Nat string `json:"nat"`
	} `json:"results"`
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
