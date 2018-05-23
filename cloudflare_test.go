package main

import (
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

// Ensure cfAPI implements the Cloudflare & Dnsrecords interfaces, compile will fail otherwise
var (
	_ Cloudflare = (*CfAPI)(nil)
	// _ Dnsrecords = (*mocks.CfAPI)(nil)
)

func TestGetDNSRecord(t *testing.T) {
	tests := []struct {
		name       string
		api        Cloudflare
		z          *cfAPI
		cf         *CfAPI
		dns        cloudflare.DNSRecord
		fqdn       string
		message    string
		wantError  bool
		wantFQDN   string
		wantResult string
		zoneID     string
	}{
		{
			name:       "GetDNSRecord() should fail",
			message:    "system panic",
			wantError:  true,
			wantResult: "function GetDNSRecord() failed: error from makeRequest",
		},
		{
			name:      "GetDNSRecord() should work",
			fqdn:      "nelly.topper.local",
			wantError: false,
			wantFQDN:  "nelly.topper.local",
			zoneID:    "deadbeef",
		},
		{
			name:       "GetDNSRecord() should fail",
			fqdn:       "nelly.topper.local",
			wantError:  true,
			wantResult: "nelly.topper.local was not found",
			zoneID:     "deadbeef",
		},
	}

	for _, tt := range tests {
		Convey(tt.name, t, func() {
			tt.cf = &CfAPI{}
			tt.api = &CfAPI{
				Fqdn:    tt.wantFQDN,
				Message: tt.message,
				WantErr: tt.wantError,
				ZoneID:  tt.zoneID,
			}
			tt.dns = cloudflare.DNSRecord{}
			tt.z = &cfAPI{}

			act, err := tt.z.GetDNSRecord(tt.api, tt.zoneID, tt.fqdn, tt.dns)
			exp := cloudflare.DNSRecord{Name: tt.fqdn}

			switch tt.wantError {
			case true:
				So(act, ShouldResemble, cloudflare.DNSRecord{})
				So(err.Error(), ShouldResemble, tt.wantResult)
			default:
				So(act.Name, ShouldResemble, exp.Name)
				So(err, ShouldBeNil)
			}
		})
	}
}

func TestNewcfAPI(t *testing.T) {
	tests := []struct {
		name      string
		osEnv     bool
		wantEmail string
		wantErr   error
		wantKey   string
		wantURL   string
	}{
		{
			name:    "No args or OS env variables",
			osEnv:   false,
			wantErr: errors.New("invalid credentials: key & email must not be empty"),
		},
		{
			name:      "No args, but with OS env variables",
			wantEmail: "test@testing.com",
			wantKey:   "1a234ef12d0b57a",
			wantURL:   "https://api.cf.local",
			osEnv:     true,
			wantErr:   nil,
		},
		{
			name:      "With args, but without OS env variables",
			wantEmail: "test@testing.com",
			wantKey:   "1a234ef12d0b57a",
			wantURL:   "https://api.cf.local",
			osEnv:     false,
			wantErr:   nil,
		},
		{
			name:      "With email and key args, but without OS env variables",
			wantEmail: "test@testing.com",
			wantKey:   "1a234ef12d0b57a",
			osEnv:     false,
			wantErr:   nil,
		},
		{
			name:      "Without args, but only OS env email variable set",
			wantEmail: "test@testing.com",
			osEnv:     true,
			wantErr:   errors.New("invalid credentials: key & email must not be empty"),
		},
		{
			name:    "Without args, but only OS env API key variable set",
			wantKey: "1a234ef12d0b57a",
			osEnv:   true,
			wantErr: errors.New("invalid credentials: key & email must not be empty"),
		},
		{
			name:      "Without args: OS env API email and key variables set",
			wantEmail: "test@testing.com",
			wantKey:   "1a234ef12d0b57a",
			osEnv:     true,
			wantErr:   nil,
		},
		{
			name:    "Without args, but only OS env API URL variable set",
			wantURL: "https://api.cf.local",
			osEnv:   true,
			wantErr: errors.New("invalid credentials: key & email must not be empty"),
		},
	}

	for _, tt := range tests {
		Convey(tt.name, t, func() {
			env := newOpts()
			So(os.Unsetenv("CF_API_EMAIL"), ShouldBeNil)
			So(os.Unsetenv("CF_API_KEY"), ShouldBeNil)
			So(os.Unsetenv("CF_API_URL"), ShouldBeNil)
			*env.email = ""
			*env.apiKey = ""
			*env.apiURL = ""

			if tt.osEnv {
				So(os.Setenv("CF_API_EMAIL", tt.wantEmail), ShouldBeNil)
				So(os.Setenv("CF_API_KEY", tt.wantKey), ShouldBeNil)
				So(os.Setenv("CF_API_URL", tt.wantURL), ShouldBeNil)
			}

			if !tt.osEnv && tt.wantErr == nil {
				*env.email = tt.wantEmail
				*env.apiKey = tt.wantKey
			}

			if tt.wantURL != "" {
				*env.apiURL = tt.wantURL
			} else {
				tt.wantURL = "https://api.cloudflare.com/client/v4"
			}

			c, err := newcfAPI(env)

			switch {
			case tt.wantErr == nil:
				So(err, ShouldBeNil)
				So(c.api.APIEmail, ShouldEqual, tt.wantEmail)
				So(c.api.APIKey, ShouldEqual, tt.wantKey)
				So(c.api.BaseURL, ShouldEqual, tt.wantURL)
				So(c, ShouldResemble, &cfAPI{})
			default:
				So(err.Error(), ShouldResemble, tt.wantErr.Error())
				So(c, ShouldBeNil)
			}
		})
	}
}
