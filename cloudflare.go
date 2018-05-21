package main

import (
	"os"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/pkg/errors"
)

// cfAPI interface wraps cloudflare.API functions that can be overriden
type cfAPI interface {
	DNSRecords(string, cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error)
	// getDNSRecord(api *cloudflare.API, zoneID, fqdn string) (cloudflare.DNSRecord, error)
	ZoneIDByName(string) (string, error)
}

// See Cloudflare instructions here: https://api.cloudflare.com/#getting-started-endpoints
func (o *opts) getCFAPI() (err error) {
	var (
		email = os.Getenv("CF_API_EMAIL")
		key   = os.Getenv("CF_API_KEY")
		opt   cloudflare.Option
	)

	setBaseURL := func(url string) cloudflare.Option {
		return func(api *cloudflare.API) error {
			api.BaseURL = url
			return nil
		}
	}

	if os.Getenv("CF_API_URL") != "" {
		opt = setBaseURL(os.Getenv("CF_API_URL"))
	}

	if *o.apiURL != "" {
		opt = setBaseURL(*o.apiURL)
	}

	if *o.apiKey != "" {
		key = *o.apiKey
	}

	if *o.email != "" {
		email = *o.email
	}

	if key != "" && email != "" {
		o.ok = true
	}

	if opt != nil {
		o.api, err = cloudflare.New(key, email, opt)
		// if o.api != nil {
		// 	o.cf = o.api
		// }

		return err
	}

	o.api, err = cloudflare.New(key, email)
	// if o.api != nil {
	// 	o.cf = o.api
	// }

	return err
}

func (o *opts) getDNSRecord(zoneID, fqdn string) (cloudflare.DNSRecord, error) {
	recs, err := o.cf.DNSRecords(zoneID, cloudflare.DNSRecord{})
	if err != nil {
		return cloudflare.DNSRecord{}, errors.Wrap(err, "function GetDNSRecord() failed:")
	}

	for _, r := range recs {
		// fmt.Printf("%s: %s\n", r.Name, r.Content)
		// fmt.Printf("Name: %s\nID: %s\nProxied: %v\n", r.Name, r.ID, r.Proxied)
		if r.Name == fqdn {
			return r, nil
		}
	}

	return cloudflare.DNSRecord{}, errors.New(fqdn + " was not found")
}

// ZoneIDByName retrieves a zone's ID from the name.
// func (o *opts) ZoneIDByName(zoneName string) (string, error) {
// 	return o.cf.ZoneIDByName(zoneName)
// }
