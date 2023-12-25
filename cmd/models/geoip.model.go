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

// cmd/models/geoip.model.go
package models

type GeoIPInfo struct {
	Continent struct {
		Code      string `json:"code,omitempty"`
		Names     struct {
			De   string `json:"de,omitempty"`
			En   string `json:"en,omitempty"`
			Es   string `json:"es,omitempty"`
			Fr   string `json:"fr,omitempty"`
			Ja   string `json:"ja,omitempty"`
			PtBR string `json:"pt-BR,omitempty"`
			Ru   string `json:"ru,omitempty"`
			ZhCN string `json:"zh-CN,omitempty"`
		} `json:"names,omitempty"`
		GeonameID int    `json:"geoname_id,omitempty"`
	} `json:"continent,omitempty"`
	Country struct {
		IsoCode           string `json:"iso_code,omitempty"`
		Names             struct {
			De   string `json:"de,omitempty"`
			En   string `json:"en,omitempty"`
			Es   string `json:"es,omitempty"`
			Fr   string `json:"fr,omitempty"`
			Ja   string `json:"ja,omitempty"`
			PtBR string `json:"pt-BR,omitempty"`
			Ru   string `json:"ru,omitempty"`
			ZhCN string `json:"zh-CN,omitempty"`
		} `json:"names,omitempty"`
		GeonameID         int    `json:"geoname_id,omitempty"`
		IsInEuropeanUnion bool   `json:"is_in_european_union,omitempty"`
	} `json:"country,omitempty"`
	RegisteredCountry struct {
		IsoCode   string `json:"iso_code,omitempty"`
		Names     struct {
			De   string `json:"de,omitempty"`
			En   string `json:"en,omitempty"`
			Es   string `json:"es,omitempty"`
			Fr   string `json:"fr,omitempty"`
			Ja   string `json:"ja,omitempty"`
			PtBR string `json:"pt-BR,omitempty"`
			Ru   string `json:"ru,omitempty"`
			ZhCN string `json:"zh-CN,omitempty"`
		} `json:"names,omitempty"`
		GeonameID int    `json:"geoname_id,omitempty"`
		IsInEuropeanUnion bool   `json:"is_in_european_union,omitempty"`
	} `json:"registered_country,omitempty"`
	City struct {
		IsoCode   string `json:"iso_code,omitempty"`
		Names     struct {
			De   string `json:"de,omitempty"`
			En   string `json:"en,omitempty"`
			Es   string `json:"es,omitempty"`
			Fr   string `json:"fr,omitempty"`
			Ja   string `json:"ja,omitempty"`
			PtBR string `json:"pt-BR,omitempty"`
			Ru   string `json:"ru,omitempty"`
			ZhCN string `json:"zh-CN,omitempty"`
		} `json:"names,omitempty"`
		GeonameID int `json:"geoname_id,omitempty"`
	} `json:"city,omitempty"`
	Location struct {
		AccuracyRadius int     `json:"accuracy_radius,omitempty"`
		Latitude       float64 `json:"latitude,omitempty"`
		Longitude      float64 `json:"longitude,omitempty"`
		TimeZone       string  `json:"time_zone,omitempty"`
	} `json:"location,omitempty"`
	Postal struct {
		Code string `json:"code,omitempty"`
	} `json:"postal,omitempty"`
	Traits struct {
		AutonomousSystemNumber       int    `json:"autonomous_system_number,omitempty"`
		AutonomousSystemOrganization string `json:"autonomous_system_organization,omitempty"`
		IPAddress                    string `json:"ip_address,omitempty"`
		Network                      string `json:"network,omitempty"`
		IsAnonymous                  bool   `json:"is_anonymous,omitempty"`
		IsAnonymousProxy             bool   `json:"is_anonymous_proxy,omitempty"`
		IsAnonymousVpn               bool   `json:"is_anonymous_vpn,omitempty"`
		IsHostingProvider            bool   `json:"is_hosting_provider,omitempty"`
		IsLegitimateProxy            bool   `json:"is_legitimate_proxy,omitempty"`
		IsPublicProxy                bool   `json:"is_public_proxy,omitempty"`
		IsResidentialProxy           bool   `json:"is_residential_proxy,omitempty"`
		IsSatelliteProvider          bool   `json:"is_satellite_provider,omitempty"`
		IsTorExitNode                bool   `json:"is_tor_exit_node,omitempty"`
		ISP                          string `json:"isp,omitempty"`
		MobileCountryCode            string `json:"mobile_country_code,omitempty"`
		MobileNetworkCode            string `json:"mobile_network_code,omitempty"`
		Organization                 string `json:"organization,omitempty"`
		StaticIPScore                int    `json:"static_ip_score,omitempty"`
		UserCount                    int    `json:"user_count,omitempty"`
		UserType                     string `json:"user_type,omitempty"`
	} `json:"traits,omitempty"`
	Subdivisions []struct {
		IsoCode   string `json:"iso_code,omitempty"`
		GeonameID int    `json:"geoname_id,omitempty"`
		Names     struct {
			De   string `json:"de,omitempty"`
			En   string `json:"en,omitempty"`
			Es   string `json:"es,omitempty"`
			Fr   string `json:"fr,omitempty"`
			Ja   string `json:"ja,omitempty"`
			PtBR string `json:"pt-BR,omitempty"`
			Ru   string `json:"ru,omitempty"`
			ZhCN string `json:"zh-CN,omitempty"`
		} `json:"names,omitempty"`
	} `json:"subdivisions,omitempty"`
}
