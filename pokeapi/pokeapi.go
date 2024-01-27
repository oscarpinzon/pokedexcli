package pokeapi

import (
	"encoding/json"
	"net/http"
)

type LocationArea struct {
	Name string `json:"name"`
}

func GetLocationAreas(url string) ([]LocationArea, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	var data struct {
		Results []LocationArea `json:"results"`
		Next    string         `json:"next"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, "", err
	}

	return data.Results, data.Next, nil
}
