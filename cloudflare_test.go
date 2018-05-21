package main

import (
	"errors"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_getCFAPI(t *testing.T) {
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
			name:    "Without args, but only OS env API URL variable set",
			wantURL: "https://api.cf.local",
			osEnv:   true,
			wantErr: errors.New("invalid credentials: key & email must not be empty"),
		},
	}

	for _, tt := range tests {
		Convey(tt.name, t, func() {
			got := newOpts()
			So(os.Unsetenv("CF_API_EMAIL"), ShouldBeNil)
			So(os.Unsetenv("CF_API_KEY"), ShouldBeNil)
			So(os.Unsetenv("CF_API_URL"), ShouldBeNil)
			*got.email = ""
			*got.apiKey = ""
			*got.apiURL = ""

			if tt.osEnv {
				So(os.Setenv("CF_API_EMAIL", tt.wantEmail), ShouldBeNil)
				So(os.Setenv("CF_API_KEY", tt.wantKey), ShouldBeNil)
				So(os.Setenv("CF_API_URL", tt.wantURL), ShouldBeNil)
			}

			if !tt.osEnv && tt.wantErr == nil {
				*got.email = tt.wantEmail
				*got.apiKey = tt.wantKey
				*got.apiURL = tt.wantURL
			}

			err := got.getCFAPI()

			switch {
			case tt.wantErr == nil:
				So(err, ShouldBeNil)
				So(got.api.APIEmail, ShouldEqual, tt.wantEmail)
				So(got.api.APIKey, ShouldEqual, tt.wantKey)
				So(got.api.BaseURL, ShouldEqual, tt.wantURL)
			case tt.wantErr != nil:
				So(err.Error(), ShouldResemble, tt.wantErr.Error())
			}
		})
	}
}

func TestGetDNSRecord(t *testing.T) {
	exitCmd = func(int) {}

	Convey("Testing getDNSRecord()", t, func() {
		env = newOpts()
		act := &fakeAPI{}
		env.cf = act
		env.log = &fakelogger{}

		*env.email = "user@big.com"
		*env.apiKey = "deadbeef"
		*env.apiURL = "http://testing.home.local"

		_, err := env.getDNSRecord("Zone ID", "fqdn.domains.local")
		So(err.Error(), ShouldEqual, "fqdn.domains.local was not found")
		So(act.message, ShouldEqual, "")
	})
}

// DNSRecords returns an array of DNSRecord
func (f *fakeAPI) DNSRecords(s string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	var r []cloudflare.DNSRecord
	r = append(r, rr)
	return r, nil
}

// ZoneIDByName retrieves a zone's ID from the name.
func (f *fakeAPI) ZoneIDByName(zoneName string) (string, error) {
	if f.message != zoneName {
		return zoneName, errors.New("command ListZones failed: error from makeRequest: HTTP request failed")
	}
	return zoneName, nil
}
