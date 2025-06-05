package player

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func GetRandomNameByCountry(country string) (string, string, error) {
	baseURL := "https://randomuser.me/api/"
	params := url.Values{}
	params.Add("nat", country)
	params.Add("gender", "male")
	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Printf("Error al obtener nombre de la API: %v", err)
		return "", "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error al leer el cuerpo de la respuesta de la API: %v", err)
		return "", "", err
	}

	log.Printf("Respuesta de la API: %s", string(body))

	var userName RandomUserName
	if err := json.Unmarshal(body, &userName); err != nil {
		log.Printf("Error al deserializar la respuesta JSON: %v", err)
		return "", "", err
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
