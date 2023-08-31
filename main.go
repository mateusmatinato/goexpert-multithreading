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
	ViaCEPURL = "https://viacep.com.br/ws/%s/json/"
	ApiCEPUrl = "https://cdn.apicep.com/file/apicep/%s.json"
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

	go GetApiCep(cepStr, respApiCep)
	go GetViaCep(cepStr, respViaCep)

	var fastestResp CepAPIComparation
	select {
	case resp := <-respApiCep:
		fastestResp = resp
	case resp := <-respViaCep:
		fastestResp = resp
	case <-time.After(5 * time.Second):
		log.Println("Timeout after 5 seconds")
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
	}
}

func GetApiCep(cepStr string, resp chan CepAPIComparation) {
	formattedUrl := fmt.Sprintf(ApiCEPUrl, cepStr)
	apiResp, err := CallExternalAPI(formattedUrl)
	if err == nil {
		resp <- CepAPIComparation{
			APIName:  "APICEP",
			APIURL:   formattedUrl,
			Duration: apiResp.Duration,
			Response: apiResp.Response,
		}
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
