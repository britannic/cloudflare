package main

import (
	"fmt"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

// Cloudflare  interface wraps cloudflare.API functions that can be overriden
// type Cloudflare interface {
// 	DNSRecords(*cloudflare.API, string, cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error)
// 	ZoneIDByName(*cloudflare.API, string) (string, error)
// }

// CfAPI is a mock type for the cfAPI type
type CfAPI struct {
	Fqdn    string
	Message string
	WantErr bool
	ZoneID  string
}

// getDNSRecord provides a mock function with given fields: api *cloudflare.API, zoneID, fqdn string
// func (c *CfAPI) GetDNSRecord(api *cloudflare.API, zoneID, fqdn string, r cloudflare.DNSRecord) (cloudflare.DNSRecord, error) {
// 	recs, err := c.DNSRecords(api, zoneID, cloudflare.DNSRecord{})
// 	if err != nil {
// 		return cloudflare.DNSRecord{}, fmt.Errorf("function GetDNSRecord() failed: %v", c.Message)
// 	}

// 	for _, rec := range recs {
// 		// fmt.Printf("%s: %s\n", r.Name, r.Content)
// 		// fmt.Printf("Name: %s\nID: %s\nProxied: %v\n", r.Name, r.ID, r.Proxied)
// 		if rec.Name == fqdn {
// 			return rec, nil
// 		}
// 	}

// 	return cloudflare.DNSRecord{}, errors.New(fqdn + " was not found")
// }

// DNSRecords returns a slice of DNS records for the given zone identifier.
//
// This takes a DNSRecord to allow filtering of the results returned.
func (c *CfAPI) DNSRecords(_ *cloudflare.API, zoneID string, _ cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	var r []cloudflare.DNSRecord
	if zoneID == "" {
		return r, fmt.Errorf("%v", "error from makeRequest")
	}

	r = append(r, cloudflare.DNSRecord{Name: c.Fqdn})
	return r, nil
}

// ZoneIDByName retrieves a zone's ID from the name.
func (c *CfAPI) ZoneIDByName(_ *cloudflare.API, zoneName string) (s string, err error) {
	return s, nil
}
