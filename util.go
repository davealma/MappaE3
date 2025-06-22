package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RoloDexResponse struct {
	OracleNotes	string `json:"oracle_notes"`
}

func DecodedOracleResponse(encoded string) string {
	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}

	return string(decodedBytes)
}

func GetRoloDex(name string) string {
	req, err := http.NewRequest("Get", os.Getenv("API_URL")+"/v1/s1/e3/resources/oracle-rolodex", nil)
	if err != nil {
		panic(err)
	}	
	req.Header.Add("API-KEY", os.Getenv("API_KEY"))

	q := req.URL.Query()
    q.Add("name", name)
    req.URL.RawQuery = q.Encode()
	// fmt.Println("url string: ",req.URL.String())

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)	
	}
	// fmt.Println("Resp: ", string(body))

	var oracleResponse RoloDexResponse

	err = json.Unmarshal(body, &oracleResponse)
	if err != nil {
		panic(err)
	}

	return oracleResponse.OracleNotes
}

func GetBalancePlanet(m map[string]int) string {
	for key, value := range m {
		if value == 0 {
			return key
		}
	}
	return ""
}

func PostSolution(planet string) {
	payload := []byte (`{
		"planet": "`+planet+`"
	}`)

	req, err := http.NewRequest("POST", os.Getenv("API_URL") + "/v1/s1/e3/solution", bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("API-KEY", os.Getenv("API_KEY"))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response: ", string(bodyResp))
}