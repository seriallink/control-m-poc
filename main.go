package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const baseUrl = "https://localhost:8443/automation-api/"

func main() {

	var payload = strings.NewReader("")

	params := map[string]string{
		"username": "workbench",
		"password": "workbench",
	}

	b, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}

	payload = strings.NewReader(string(b))

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := http.Client{
		Transport: transport,
	}

	request, _ := http.NewRequest(http.MethodPost, baseUrl+"session/login", payload)
	request.Header.Add("accept", "application/json")
	request.Header.Add("content-type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(data))

}
