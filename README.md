# MaxMind GeoIP2 Go API

[![Go Reference](https://pkg.go.dev/badge/github.com/iamvladw/GeoIP2-go.svg)](https://pkg.go.dev/github.com/iamvladw/GeoIP2-go)

## Description

This package provides a API for the [GeoIP2 web services and GeoLite2 web
services](https://dev.maxmind.com/geoip/docs/web-services?lang=en).


## Installation

```sh
go get github.com/iamvladw/GeoIP2-go
```

## IP Geolocation Usage

IP geolocation is inherently imprecise. Locations are often near the center of
the population. Any location provided by a GeoIP2 database or web service
should not be used to identify a particular address or household.

## Web Service Usage

To use the web service API, you must create a new `API Instance`, using
your MaxMind `accountID` and `licenseKey` as parameters. The third argument is
the `host` option. Set `host` to `geolite.info`
to use the GeoLite2 web service. Set `host` to `geoip.maxmind.com`. Set `host` to
`sandbox.maxmind.com` to use the Sandbox environment.

You may then call the function corresponding to a specific end point, passing it
the IP address you want to lookup.

If the request succeeds, the function's response will resolve with the model
for the end point you called. This model in turn contains multiple
records, each of which represents part of the data returned by the web service.

If the request fails, the function's response will reject with an error object.

## Web Service Example

### Country Service

```go
import (
    "fmt"

    "github.com/iamvladw/GeoIP2-go"
)

// To use the GeoIP2 web service instead of the GeoLite2 web service, set
// the host to geoip.maxmind.com, e.g.:
// api := geoip2.New("1234", "licenseKey", "geoip.maxmind.com")
//
// To use the GeoLite2 web service instead of the GeoIP2 web service, set
// the host to geolite.info, e.g.:
// api := geoip2.New("1234", "licenseKey", "geolite.info")
//
// To use the Sandbox GeoIP2 web service instead of the production GeoIP2
// web service, set the host to sandbox.maxmind.com, e.g.:
// api := geoip2.New("1234", "licenseKey", "sandbox.maxmind.com")
api := geoip2.New("1234", "licenseKey", "geoip.maxmind.com")

res, err := api.Country("142.1.1.1")
if err != nil {
	fmt.Println(err)
	return err
}

fmt.Println(res.Country.IsoCode) // 'CA'
```

### City Plus Service

```go
import (
    "fmt"

    "github.com/iamvladw/GeoIP2-go"
)

// To use the GeoIP2 web service instead of the GeoLite2 web service, set
// the host to geoip.maxmind.com, e.g.:
// api := geoip2.New("1234", "licenseKey", "geoip.maxmind.com")
//
// To use the GeoLite2 web service instead of the GeoIP2 web service, set
// the host to geolite.info, e.g.:
// api := geoip2.New("1234", "licenseKey", "geolite.info")
//
// To use the Sandbox GeoIP2 web service instead of the production GeoIP2
// web service, set the host to sandbox.maxmind.com, e.g.:
// api := geoip2.New("1234", "licenseKey", "sandbox.maxmind.com")
api := geoip2.New("1234", "licenseKey", "geoip.maxmind.com")

res, err := api.City("142.1.1.1")
if err != nil {
	fmt.Println(err)
	return err
}

fmt.Println(res.Country.IsoCode) // 'CA'
fmt.Println(res.Postal.Code) // 'M5S'
```

### Insights Service

```go
import (
    "fmt"

    "github.com/iamvladw/GeoIP2-go"
)

// To use the GeoIP2 web service instead of the GeoLite2 web service, set
// the host to geoip.maxmind.com, e.g.:
// api := geoip2.New("1234", "licenseKey", "geoip.maxmind.com")
//
// To use the GeoLite2 web service instead of the GeoIP2 web service, set
// the host to geolite.info, e.g.:
// api := geoip2.New("1234", "licenseKey", "geolite.info")
//
// To use the Sandbox GeoIP2 web service instead of the production GeoIP2
// web service, set the host to sandbox.maxmind.com, e.g.:
// api := geoip2.New("1234", "licenseKey", "sandbox.maxmind.com")
api := geoip2.New("1234", "licenseKey", "geoip.maxmind.com")

res, err := api.City("142.1.1.1")
if err != nil {
	fmt.Println(err)
	return err
}

fmt.Println(res.Country.IsoCode) // 'CA'
fmt.Println(res.Postal.Code) // 'M5S'
fmt.Println(res.Traits.UserType) // 'school'
```

## Web Service Errors

For details on the possible errors returned by the web service itself, [see
the GeoIP2 web service documentation](https://dev.maxmind.com/geoip2/geoip/web-services).

If the web service returns an explicit error document, the promise will be rejected
with the following object structure:

```go
"some human readable error"
```

In addition to the possible errors returned by the web service, the following error
codes are provided:

* `Error creating HTTP request` or `Error sending HTTP request` for errors creating or sending the HTTP request to the web service
* `HTTP request failed with status code` for 4xx & 5xx level errors
* `Error reading response body` or `Error unmarshaling JSON` for invalid JSON responses or unmarshable response bodies
* [General Go error codes](https://go.dev/blog/error-handling-and-go)

## Integration with GeoNames

[GeoNames](https://www.geonames.org/) offers web services and downloadable
databases with data on geographical features around the world, including
populated places. They offer both free and paid premium data. Each feature
is uniquely identified by a `geonameId`, which is an integer.

Many of the records returned by the GeoIP web services and databases include a
`geonameId` field. This is the ID of a geographical feature (city, region,
country, etc.) in the GeoNames database.

Some of the data that MaxMind provides is also sourced from GeoNames. We source
things like place names, ISO codes, and other similar data from the GeoNames
premium data set.

## Reporting Data Problems

If the problem you find is that an IP address is incorrectly mapped, please
[submit your correction to MaxMind](https://www.maxmind.com/en/geoip-data-correction-request).

If you find some other sort of mistake, like an incorrect spelling,
please check [the GeoNames site](https://www.geonames.org/) first. Once
you've searched for a place and found it on the GeoNames map view, there
are a number of links you can use to correct data ("move", "edit",
"alternate names", etc.). Once the correction is part of the GeoNames
data set, it will be automatically incorporated into future MaxMind
releases.

If you are a paying MaxMind customer and you're not sure where to submit a
correction, please [contact MaxMind support for help](https://support.maxmind.com/hc/en-us/requests/new).

## Requirements

[@Vlad](https://github.com/iamvladw) & [@Alex](https://github.com/Alexandru200509) have tested this API with Go versions 21. We aim to support
active and maintained versions of Go.

MaxMind did not test this API (yet, maybe in the future) on any version of Go.

## Contributing

Patches and pull requests are encouraged. Please include unit tests
whenever possible, as we strive to maintain 100% code coverage.

## Support

Please report all issues with this code using the [GitHub issue
tracker](https://github.com/iamvladw/GeoIP2-go/issues)

If you are having an issue with a MaxMind service that is not specific to the
client API, please contact [MaxMind support for assistance](https://support.maxmind.com/hc/en-us/requests/new).

## Copyright and License

This is free software, licensed under the Apache License, Version 2.0.