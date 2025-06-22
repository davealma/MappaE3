package main

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
)

type SwapiResponse struct {
	Count   int            `json:"Count"`
	Next    *string        `json:"Next,omitempty"`
	Results []PersonResult `json:"Results"`
}

type PlanetSwapiResponse struct {
	Name string	`json:"Name"`
}

func GetHolocronSwapi(url string) SwapiResponse {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var people SwapiResponse

	err = json.Unmarshal(body, &people)
	if err != nil {
		panic(err)
	}

	return people
}

func GetPlanetSwapi(url string) PlanetSwapiResponse {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var planet PlanetSwapiResponse

	err = json.Unmarshal(body, &planet)
	if err != nil {
		panic(err)
	}

	return planet
}