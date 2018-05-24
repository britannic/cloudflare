package main

import (
	"fmt"
	"os"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/pkg/errors"
)

// Cloudflare  interface wraps cloudflare.API functions that can be overriden
type Cloudflare interface {
	DNSRecords(*cloudflare.API, string, cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error)
	UpdateDNSRecord(*cloudflare.API, string, string, cloudflare.DNSRecord) error
	ZoneIDByName(*cloudflare.API, string) (string, error)
}

// type Dnsrecords interface {
// 	GetDNSRecord(*cloudflare.API, string, string, cloudflare.DNSRecord) (cloudflare.DNSRecord, error)
// }

type cfAPI struct {
	api *cloudflare.API
}

// Ensure cfAPI implements the Cloudflare and Dnsrecords interfaces, compile will fail otherwise
var (
	_ Cloudflare = (*cfAPI)(nil)
	// _ Dnsrecords = (*cfAPI)(nil)
)

func setBaseURL(url string) cloudflare.Option {
	return func(api *cloudflare.API) error {
		api.BaseURL = url
		return nil
	}
}

// See Cloudflare instructions here: https://api.cloudflare.com/#getting-started-endpoints
func newcfAPI(o *opts) (*cfAPI, error) {
	var (
		email = os.Getenv("CF_API_EMAIL")
		key   = os.Getenv("CF_API_KEY")
		opt   cloudflare.Option
	)

	if is(os.Getenv("CF_API_URL")) {
		opt = setBaseURL(os.Getenv("CF_API_URL"))
	}

	if is(*o.apiURL) {
		opt = setBaseURL(*o.apiURL)
	}

	if is(*o.apiKey) {
		key = *o.apiKey
	}

	if is(*o.email) {
		email = *o.email
	}

	if key == "" || email == "" {
		return nil, errors.New("invalid credentials: key & email must not be empty")
	}

	if opt != nil {
		cf, err := cloudflare.New(key, email, opt)
		return &cfAPI{api: cf}, err
	}

	cf, err := cloudflare.New(key, email)
	return &cfAPI{api: cf}, err
}

// GetDNSRecord returns a single cloudflare.DNSRecord
func (c *cfAPI) GetDNSRecord(cf Cloudflare, zoneID, fqdn string, r cloudflare.DNSRecord) (cloudflare.DNSRecord, error) {
	recs, err := cf.DNSRecords(c.api, zoneID, r)
	if err != nil {
		return cloudflare.DNSRecord{}, fmt.Errorf("function GetDNSRecord() failed: %v", err)
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

func is(s string) bool {
	return s != ""
}

// DNSRecords returns a slice of DNS records for the given zone identifier.
//
// This takes a DNSRecord to allow filtering of the results returned.
func (c *cfAPI) DNSRecords(api *cloudflare.API, zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return api.DNSRecords(zoneID, rr)
}

func (c *cfAPI) UpdateDNSRecord(api *cloudflare.API, zoneID string, recordID string, rr cloudflare.DNSRecord) error {
	return api.UpdateDNSRecord(zoneID, recordID, rr)
}

// ZoneIDByName retrieves a zone's ID from the name.
func (c *cfAPI) ZoneIDByName(api *cloudflare.API, zoneName string) (string, error) {
	return api.ZoneIDByName(zoneName)
}
