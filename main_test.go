package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBasename(t *testing.T) {
	Convey("Testing basename()", t, func() {
		tests := []struct {
			s   string
			exp string
		}{
			{s: "e.txt", exp: "e"},
			{s: "/github.com/britannic/blacklist/internal/edgeos", exp: "edgeos"},
		}

		for _, tt := range tests {
			So(basename(tt.s), ShouldEqual, tt.exp)
		}
	})
}

func TestCleanArgs(t *testing.T) {
	Convey("Testing cleanArgs()", t, func() {
		tests := []struct {
			s   []string
			exp []string
		}{
			{s: []string{"-convey", "-test", "-file", "-h"}, exp: []string{"-file", "-h"}},
			{s: []string{"-file", "-h"}, exp: []string{"-file", "-h"}},
		}

		for _, tt := range tests {
			So(cleanArgs(tt.s), ShouldResemble, tt.exp)
		}
	})
}
