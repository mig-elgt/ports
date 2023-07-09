package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"net/http"
	"os"

	"github.com/bcicen/jstream"
	"github.com/mitchellh/mapstructure"
)

type Port struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

func main() {
	filePath := flag.String("file", "ports.json", "json file path")
	flag.Parse()

	f, err := os.Open(*filePath)
	if err != nil {
		panic(err)
	}
	// Create new JSON decoder as stream
	decoder := jstream.NewDecoder(f, 1)
	// For each stream input, read stream, store it on memory and
	// send a request to our Port API Service to store the port.
	for stream := range decoder.Stream() {
		var port Port
		// Convert stream value to Port Object
		mapstructure.Decode(stream.Value, &port)
		// Send a request to our API Server
		if err := sendRequest(port); err != nil {
			panic(err)
		}
	}
}

const requestURL = "http://localhost:8082/v1/ports"

func sendRequest(port Port) error {
	jsonBody, err := json.Marshal(port)
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(jsonBody)
	// Create new Request
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return err
	}
	// Send POST Request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	// Validate response status code
	if res.StatusCode != http.StatusOK {
		if err != nil {
			return err
		}
	}
	return nil
}
