package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	ViaCEPURL    = "https://viacep.com.br/ws/%s/json/"
	BrasilAPIUrl = "https://brasilapi.com.br/api/cep/v2/%s"
)

type CepAPIComparation struct {
	APIName  string
	APIURL   string
	Duration time.Duration
	Response map[string]any
}

type APIResponse struct {
	Duration time.Duration
	Response map[string]any
}

func main() {
	args := os.Args

	params := args[1:]
	if len(params) != 1 {
		log.Fatalln("Invalid number of parameters - expected 1 but got", len(params))
	}
	cepStr := params[0]

	respApiCep := make(chan CepAPIComparation)
	respViaCep := make(chan CepAPIComparation)

	go GetBrasilAPI(cepStr, respApiCep)
	go GetViaCep(cepStr, respViaCep)

	var fastestResp CepAPIComparation
	select {
	case resp := <-respApiCep:
		fastestResp = resp
	case resp := <-respViaCep:
		fastestResp = resp
	case <-time.After(1 * time.Second):
		log.Println("Timeout after 1 second")
		return
	}

	log.Println("API Name:", fastestResp.APIName)
	log.Println("API URL:", fastestResp.APIURL)
	log.Println("Duration:", fastestResp.Duration)
	log.Println("Response:", fastestResp.Response)
}

func GetViaCep(cepStr string, resp chan CepAPIComparation) {
	formattedUrl := fmt.Sprintf(ViaCEPURL, cepStr)
	apiResp, err := CallExternalAPI(formattedUrl)
	if err == nil {
		resp <- CepAPIComparation{
			APIName:  "VIACEP",
			APIURL:   formattedUrl,
			Duration: apiResp.Duration,
			Response: apiResp.Response,
		}
	} else {
		log.Printf("Error on VIA CEP - URL: %s - Error: %s\n", formattedUrl, err.Error())
	}
}

func GetBrasilAPI(cepStr string, resp chan CepAPIComparation) {
	formattedUrl := fmt.Sprintf(BrasilAPIUrl, cepStr)
	apiResp, err := CallExternalAPI(formattedUrl)
	if err == nil {
		resp <- CepAPIComparation{
			APIName:  "BrasilAPI",
			APIURL:   formattedUrl,
			Duration: apiResp.Duration,
			Response: apiResp.Response,
		}
	} else {
		log.Printf("Error on BrasilAPI - URL: %s - Error: %s\n", formattedUrl, err.Error())
	}
}

func CallExternalAPI(url string) (*APIResponse, error) {
	startTime := time.Now()
	httpResp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("invalid status code %d", httpResp.StatusCode)
	}

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	cepResp := make(map[string]any)
	err = json.Unmarshal(body, &cepResp)
	if err != nil {
		return nil, err
	}
	endTime := time.Now()

	return &APIResponse{
		Duration: endTime.Sub(startTime),
		Response: cepResp,
	}, nil
}
