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

// cmd/tests/geoip._test.go
package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/iamvladw/GeoIP2-go/cmd/geoip2"
)

func TestCloudflare(t *testing.T) {
	fmt.Println("Running Cloudflare test...")

	// Create a new GeoIP2 API client.
	api := geoip2.New(os.Getenv("ACCOUNTID"), os.Getenv("LICENSEKEY"), "geolite.info")

	// Make a request to the GeoIP2 API.
	res, err := api.City("1.1.1.1")
	if err != nil {
		t.Error(err)
		return
	}

	// Checks the response.
	if res.RegisteredCountry.Names.En != "Australia" {
		t.Error("Expected Australia, got", res.RegisteredCountry.Names.En)
		return
	}

	// Checks the IsoCode.
	if res.RegisteredCountry.IsoCode != "AU" {
		t.Error("Expected AU, got", res.RegisteredCountry.IsoCode)
		return
	}

	// Checks the GeoNameID.
	if res.RegisteredCountry.GeonameID != 2077456 {
		t.Error("Expected 2077456, got", res.RegisteredCountry.GeonameID)
		return
	}

	// Checks the AutonomousSystemNumber.
	if res.Traits.AutonomousSystemNumber != 13335 {
		t.Error("Expected 13335, got", res.RegisteredCountry.GeonameID)
		return
	}

	// Checks the AutonomousSystemOrganization.
	if res.Traits.AutonomousSystemOrganization != "CLOUDFLARENET" {
		t.Error("Expected CLOUDFLARENET, got", res.Traits.AutonomousSystemOrganization)
		return
	}

	// Checks the AnonymousProxy.
	if res.Traits.IsAnonymousProxy != false {
		t.Error("Expected false, got", res.Traits.IsAnonymousProxy)
		return
	}

	// Checks the SatelliteProvider.
	if res.Traits.IsSatelliteProvider != false {
		t.Error("Expected false, got", res.Traits.IsSatelliteProvider)
		return
	}

	// Finished the test.
	fmt.Println("Passed BadCredentials test!")
}
