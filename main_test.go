package main

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/britannic/mflag"
	. "github.com/smartystreets/goconvey/convey"
)

func (o *opts) String() string {
	var s string
	o.VisitAll(func(f *mflag.Flag) {
		s += fmt.Sprintf("  -%s", f.Name) // Two spaces before -; see next two comments.

		name, usage := mflag.UnquoteUsage(f)
		if len(name) > 0 {
			s += " " + name
		}
		// Boolean flags of one ASCII letter are so common we
		// treat them specially, putting their usage on the same line.
		if len(s) <= 4 { // space, space, '-', 'x'.
			s += "\t"
		} else {
			// Four spaces before the tab triggers good alignment
			// for both 4- and 8-space tab stops.
			s += "\n    \t"
		}
		s += usage
		if !mflag.IsZeroValue(f, f.DefValue) {
			if _, ok := f.Value.(*mflag.StringValue); ok {
				// put quotes on the value
				s += fmt.Sprintf(" (default %q)", f.DefValue)
			} else {
				s += fmt.Sprintf(" (default %v)", f.DefValue)
			}
		}
		s = fmt.Sprint(s, "\n")

	})

	return s
}

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

func TestExitCmd(t *testing.T) {
	Convey("Testing exitCmd", t, func() {
		var (
			act int
		)
		exitCmd = func(i int) {
			act = i
		}

		exitCmd(0)
		So(act, ShouldEqual, 0)
	})
}

func TestSetArgs(t *testing.T) {
	var (
		origArgs = os.Args
		prog     = path.Base(os.Args[0])
	)

	exitCmd = func(int) {}
	defer func() { os.Args = origArgs }()

	tests := []struct {
		name string
		args []string
		exp  interface{}
	}{
		// {
		// 	name: "h",
		// 	args: []string{prog, "-convey-json", "-h"},
		// 	exp:  true,
		// },
		// {
		// 	name: "debug",
		// 	args: []string{prog, "-debug"},
		// 	exp:  true,
		// },
		// {
		// 	name: "dryrun",
		// 	args: []string{prog, "-dryrun"},
		// 	exp:  true,
		// },
		// {
		// 	name: "token",
		// 	args: []string{prog, "-token", "a12s23984d32e123f432"},
		// 	exp:  "a12s23984d32e123f432",
		// },
		// {
		// 	name: "version",
		// 	args: []string{prog, "-version"},
		// 	exp:  true,
		// },
		// {
		// 	name: "v",
		// 	args: []string{prog, "-v"},
		// 	exp:  true,
		// },
		{
			name: "invalid flag",
			args: []string{prog, "-z"},
			exp: `flag provided but not defined: -z
Usage of ` + prog + `:
  -dryrun
    	Run config and data validation tests
  -f <file>
    	<file> # Load a config.boot file
  -h	Display help
  -token string
    	Cloudflare API token
  -v	Verbose display
  -version
    	Show version
`,
		},
	}

	// check := func(flg interface{}){}

	for _, tt := range tests {
		os.Args = nil
		if tt.args != nil {
			os.Args = tt.args
		}

		env := newOpts()
		env.Init(prog, mflag.ContinueOnError)

		Convey("Testing commandline output", t, func() {
			Convey("Testing setArgs() with "+tt.name+"\n", func() {
				switch {
				case tt.name == "token":
					env.setArgs()
					So(cftoken, ShouldEqual, tt.exp.(string))
				case tt.name == "invalid flag":
					act := new(bytes.Buffer)
					env.SetOutput(act)
					env.setArgs()
					So(act.String(), ShouldEqual, tt.exp.(string))
				default:
					env.setArgs()
					So(fmt.Sprint(env.Lookup(tt.name).Value.String()), ShouldEqual, fmt.Sprint(tt.exp))
				}
			})
		})
	}
}
