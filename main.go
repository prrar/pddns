/*
	pddns.go: Updates DNS entries on Cloudflare
	          2025-04-30 - prrar

	Example of config.json:

	{
		"zone_id": "abcd1234567890",
		"api_token": "abcd1234567890"
	}

	Usage: pddns --config /path/to/config.json \
	             --dns-entry subdomain.domain.com \
			 	 --ip 123.456.789.012

	If no IP is specified, the new IP is set to current external IP.
*/

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// DNSUpdateRequest represents the payload to update a DNS record.
type DNSUpdateRequest struct {
	Name string `json:"name"`
	IP   string `json:"content"`
	Type string `json:"type"`
}

// DNSRecord represents a DNS record retrieved from Cloudflare.
type DNSRecord struct {
	Name string `json:"name"`
	IP   string `json:"content"`
	Type string `json:"type"`
	ID   string `json:"id"`
}

// DNSRecordsResponse represents the response from Cloudflare's DNS records API.
type DNSRecordsResponse struct {
	Records []DNSRecord `json:"result"`
}

// Config holds the configuration parameters.
type Config struct {
	ZoneID   string `json:"zone_id"`
	APIToken string `json:"api_token"`
}

// httpClient is a reusable HTTP client with a timeout.
var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

// getConfig reads and parses the configuration from the specified file.
func getConfig(filePath string) (Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("reading config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("parsing config file: %w", err)
	}

	if config.ZoneID == "" || config.APIToken == "" {
		return Config{}, errors.New("incomplete config: missing required fields")
	}

	return config, nil
}

// getCurrentIP retrieves the current public IP address.
func getCurrentIP() (string, error) {
	resp, err := httpClient.Get("https://api.ipify.org?format=text")
	if err != nil {
		return "", fmt.Errorf("fetching current IP: %w", err)
	}
	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading IP response: %w", err)
	}

	return string(ip), nil
}

// isValidIPv4 checks if the user specified IP is of valid format.
func isValidIPv4(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil && parsedIP.To4() != nil
}

// setHeaders adds necessary headers to a Cloudflare API request.
func setHeaders(req *http.Request, token string) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
}

// getDNSRecords retrieves DNS records for the specified zone.
func getDNSRecords(zoneID, token string) ([]DNSRecord, error) {
	url := fmt.Sprintf(
		"https://api.cloudflare.com/client/v4/zones/%s/dns_records", zoneID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	setHeaders(req, token)
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %s\nBody: %s", res.Status, body)
	}

	var response DNSRecordsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("parsing response JSON: %w", err)
	}

	return response.Records, nil
}

// updateDNSRecord updates the specified DNS record with the new IP address.
func updateDNSRecord(zoneID, token, recordID string, update DNSUpdateRequest) error {
	jsonData, err := json.Marshal(update)
	if err != nil {
		return fmt.Errorf("marshaling update data: %w", err)
	}

	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s",
		zoneID, recordID)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("creating update request: %w", err)
	}

	setHeaders(req, token)
	res, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("performing update request: %w", err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update DNS record: %s\nBody: %s", res.Status, body)
	}

	return nil
}

func main() {
	configFile := flag.String("config", "", "Path to config file")
	dnsEntry := flag.String("dns-entry", "", "DNS entry to update")
	userDefinedIP := flag.String("ip", "", "User defined IP to DNS entry")
	flag.Parse()

	if *configFile == "" || *dnsEntry == "" {
		progName := filepath.Base(os.Args[0])
		log.Fatalf(
			"Both --config and --dns-entry flags are required, --ip is optional.\n\n"+
				"Usage: %s\t--config /path/to/config.json \\ \n"+
				"\t\t--dns-entry subdomain.domain.com \\ \n"+
				"\t\t--ip 123.456.789.012\n\n"+
				"If no IP is specified, the new IP is set to current external IP.",
			progName,
		)
	}

	config, err := getConfig(*configFile)
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	var newIP string
	if *userDefinedIP == "" {
		newIP, err = getCurrentIP()
		if err != nil {
			log.Fatalf("Error retrieving current IP: %v", err)
		}
	} else {
		if isValidIPv4(*userDefinedIP) {
			newIP = *userDefinedIP
		} else {
			log.Fatalf("User specified IP %s is not a valid IPv4", *userDefinedIP)
		}
	}

	records, err := getDNSRecords(config.ZoneID, config.APIToken)
	if err != nil {
		log.Fatalf("Error retrieving DNS records: %v", err)
	}

	var targetRecord *DNSRecord
	for _, record := range records {
		if record.Name == *dnsEntry {
			targetRecord = &record
			break
		}
	}

	if targetRecord == nil {
		log.Fatalf("DNS entry %s not found.", *dnsEntry)
	}

	if targetRecord.Type != "A" {
		log.Fatalf("DNS entry %s is not of type A.", *dnsEntry)
	}

	if newIP == targetRecord.IP {
		log.Println("IP address has not changed.")
		return
	}

	log.Printf("Updating DNS entry %s: IP changed from %s to %s",
		*dnsEntry, targetRecord.IP, newIP)

	update := DNSUpdateRequest{
		Name: *dnsEntry,
		IP:   newIP,
		Type: "A",
	}

	if err := updateDNSRecord(config.ZoneID, config.APIToken, targetRecord.ID, update); err != nil {
		log.Fatalf("Error updating DNS record: %v", err)
	}

	log.Println("DNS record updated successfully.")
}
