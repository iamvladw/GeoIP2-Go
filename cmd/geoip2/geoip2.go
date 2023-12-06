/* Copyright 2023 Vlad White & Alexandru Suciu

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. */

// cmd/geoip2/geoip2.go
package geoip2

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/iamvladw/GeoIP2-go/cmd/models"
)

// Api is a client for the GeoIP2 API.
type Api struct {
	accountID  string
	licenseKey string
	host       string
}

// New creates a new GeoIP2 API client.
func New(accountID, licenseKey, host string) *Api {
	return &Api{
		accountID:  accountID,
		licenseKey: licenseKey,
		host:       host,
	}
}

// Country returns GeoIP2 country information for the given IP address.
func (a *Api) Country(ipAddress string) (models.GeoIPInfo, error) {
	return a.request("/geoip/v2.1/country/", ipAddress)
}

// City returns GeoIP2 city information for the given IP address.
func (a *Api) City(ipAddress string) (models.GeoIPInfo, error) {
	return a.request("/geoip/v2.1/city/", ipAddress)
}

// Insights returns GeoIP2 insights information for the given IP address.
func (a *Api) Insights(ipAddress string) (models.GeoIPInfo, error) {
	return a.request("/geoip/v2.1/insights/", ipAddress)
}

// request makes a request to the GeoIP2 API.
func (a *Api) request(prefix, ipAddress string) (models.GeoIPInfo, error) {
	// Build the request URL.
	request := fmt.Sprintf("https://%s%s%s", a.host, prefix, ipAddress)

	// Create a models.GeoIPInfo struct.
	var data models.GeoIPInfo

	// Create the HTTP request.
	req, err := http.NewRequest("GET", request, nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return models.GeoIPInfo{}, err
	}

	// Set the request headers.
	req.SetBasicAuth(a.accountID, a.licenseKey)

	// Send the HTTP request.
	client := &http.Client{}

	// Get the HTTP response.
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return models.GeoIPInfo{}, err
	}

	// Close the response body.
	defer response.Body.Close()

	// Check the HTTP status code.
	if response.StatusCode != http.StatusOK {
		fmt.Println("HTTP request failed with status code", response.StatusCode)
		return models.GeoIPInfo{}, err
	}

	// Read the response body.
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return models.GeoIPInfo{}, err
	}

	// Unmarshal the JSON response.
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return models.GeoIPInfo{}, err
	}

	// Return the models.GeoIPInfo struct.
	return data, nil
}
