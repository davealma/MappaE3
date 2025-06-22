package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type PersonResult struct {
	Name	string	`json:"Name"`
	Planet string `json:"homeworld"`
}

func main() {
	var persons []PersonResult
	planetBalance := make(map[string]int)
	godotenv.Load()
	fmt.Println("Measuring the balance force...")
	swapiResponse := GetHolocronSwapi(os.Getenv("API_HOLOCRON_CHARACTERS"))
	persons = append(persons, swapiResponse.Results...)
	for {		
		if swapiResponse.Next != nil {
			swapiResponse = GetHolocronSwapi(*swapiResponse.Next)
			persons = append(persons, swapiResponse.Results...)
		}else {
			break
		}		
	}
	
	for i := 0; i < len(persons); i++ {
		p := persons[i]
		codedResponse := GetRoloDex(p.Name)
		decodedResponse := DecodedOracleResponse(codedResponse)
		if strings.Contains(decodedResponse, "Dark") {
			if count, ok := planetBalance[p.Planet]; ok {
				planetBalance[p.Planet] = count - 1
			}else {
				planetBalance[p.Planet] = -1
			}
		}else {
			if count, ok := planetBalance[p.Planet]; ok {
				planetBalance[p.Planet] = count + 1
			}else {
				planetBalance[p.Planet] = 1
			}
		}
	}

	var planetUrl = GetBalancePlanet(planetBalance)
	fmt.Println("BalancedWorld", planetUrl)
	var planet = GetPlanetSwapi(planetUrl)
	fmt.Println("Planet Name: ", planet)
	PostSolution(planet.Name)
}