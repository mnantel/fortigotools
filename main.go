package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// Define a struct to hold the API response
type ApiResponse struct {
	Results []interface{} `json:"results"`
	// Include other response fields as necessary
}

func main() {

	fmt.Println("Opening config file: config.yaml")

	cfg := viper.New()
	cfg.SetConfigType("yaml")
	cfg.SetConfigFile("config.yaml")

	if err := cfg.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return
	}

	// The URL to the FortiOS API endpoint for fetching address objects
	myfw, err := NewFos(cfg.GetString("hostname"), cfg.GetString("apikey"))
	if err != nil {
		fmt.Printf("Error creating Fos object: %v\n", err)
		return
	}
	// Create a new request
	fwaddr, err := myfw.GetFirewallAddress()
	if err != nil {
		fmt.Printf("Error fetching address objects: %v\n", err)
		return
	}
	// Print the address objects
	for _, addressObject := range fwaddr {
		fmt.Printf("Name: %s, Type: %s, Subnet: %s\n", addressObject.Name, addressObject.Type, addressObject.Subnet)
	}
}
