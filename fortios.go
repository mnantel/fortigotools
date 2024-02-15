package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

type Fos struct {
	Hostname string `json:"hostname"`
	apikey   string
}

func NewFos(hostname string, apikey string) (*Fos, error) {
	f := &Fos{}
	f.apikey = apikey
	f.Hostname = hostname
	return f, nil
}

func (f *Fos) MakeApiCall(method string, path string, body string) ([]byte, error) {

	url := fmt.Sprintf("https://%v%v", f.Hostname, path)

	// Disable SSL certificate verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Initialize HTTP client with custom transport
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return nil, err
	}

	// Add the API token for authentication
	req.Header.Add("Authorization", "Bearer "+f.apikey)

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	rawbody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil, err
	}

	return rawbody, nil

}
