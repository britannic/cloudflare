package main

import (
	"errors"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_getCFAPI(t *testing.T) {
	tests := []struct {
		name      string
		osEnv     bool
		wantEmail string
		wantKey   string
		wantErr   error
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
			osEnv:     true,
			wantErr:   nil,
		},
		{
			name:      "With args, but without OS env variables",
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
	}

	for _, tt := range tests {
		Convey(tt.name, t, func() {
			So(os.Unsetenv("CF_API_EMAIL"), ShouldBeNil)
			So(os.Unsetenv("CF_API_KEY"), ShouldBeNil)
			*env.email = ""
			*env.apiKey = ""

			if tt.osEnv {
				So(os.Setenv("CF_API_EMAIL", tt.wantEmail), ShouldBeNil)
				So(os.Setenv("CF_API_KEY", tt.wantKey), ShouldBeNil)
			}

			if !tt.osEnv && tt.wantErr == nil {
				*env.email = tt.wantEmail
				*env.apiKey = tt.wantKey
			}

			got, err := env.getCFAPI()

			switch {
			case tt.wantErr == nil:
				So(err, ShouldBeNil)
				So(got.APIEmail, ShouldEqual, tt.wantEmail)
				So(got.APIKey, ShouldEqual, tt.wantKey)
			case tt.wantErr != nil:
				So(err.Error(), ShouldResemble, tt.wantErr.Error())
			}
		})
	}
}
