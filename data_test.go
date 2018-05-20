package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_opts_writeFile(t *testing.T) {
	type args struct {
		f string
		r io.Reader
	}

	tests := []struct {
		name    string
		exp     string
		args    args
		wantErr bool
	}{
		{
			name:    "writefile() returns error",
			wantErr: true,
		},
		{
			name: "writefile() without error",
			args: args{
				f: "testfile",
				r: strings.NewReader("test data"),
			},
			exp:     "test data",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		Convey(tt.name, t, func() {
			switch {
			case tt.wantErr:
				So(writeFile(tt.args.f, tt.args.r), ShouldNotBeNil)
			default:
				var (
					dir  = "/tmp"
					ext  = ".delete"
					file = fmt.Sprintf("%v/%v%v", dir, tt.args.f, ext)
				)

				So(writeFile(file, tt.args.r), ShouldBeNil)

				act, err := ioutil.ReadFile(file)
				So(err, ShouldBeNil)
				So(act, ShouldResemble, []byte(tt.exp))

				So(os.Remove(file), ShouldBeNil)
			}
		})
	}
}
