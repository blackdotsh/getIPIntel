package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"os"
)

const (
	getIPIntelURL = "http://check.getipintel.net/check.php?" //change URL here if you want to add more arguments
	timeout        = 5 * time.Second // 5 seconds
	maxProbability = 0.99
	contactEmail = "yourContactEmail" //Replace with your contact email
)

func checkIP(ip string) (bool, error) {
	client := &http.Client{Timeout: timeout}

	// Construct the request URL
	url := fmt.Sprintf("%sip=%s&contact=%s", getIPIntelURL, ip, contactEmail)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, fmt.Errorf("error creating request: %w", err)
	}

	// Send the request and handle potential errors
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body and convert to float
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("error reading response body: %w", err)
	}
	probability, err := strconv.ParseFloat(string(body), 64) // Use ParseFloat for precision

	if err != nil {
		return false, fmt.Errorf("error parsing probability: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("GetIPIntel API returned error: %s, GetIPIntel error code: %.0f", resp.Status, probability)
	}

	return probability > maxProbability, nil
}

func main() {
	// Get the IP address from command-line arguments or user input
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <IP address>")
		return
	}
	ip := os.Args[1]

	// Check the IP and handle errors
	isSuspicious, err := checkIP(ip)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if isSuspicious {
		fmt.Printf("%s is probably a proxy / VPN / bad IP.\n", ip)
	} else {
		fmt.Printf("%s is probably NOT a proxy / VPN / bad IP.\n", ip)
	}
}
