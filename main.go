package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/britannic/mflag"
)

var (
	// updated by go build -ldflags
	architecture = "UNKNOWN"
	build        = "UNKNOWN"
	githash      = "UNKNOWN"
	hostOS       = "UNKNOWN"
	version      = "UNKNOWN"
	// ----------------------------

	exitCmd = os.Exit
	prog    = basename(os.Args[0])
	prefix  = fmt.Sprintf("%s: ", prog)
)

// opts struct for command line options and setting initial variables
type opts struct {
	*mflag.FlagSet
	arch    *string
	dbug    *bool
	dryrun  *bool
	file    *string
	help    *bool
	hostOS  *string
	mips64  *string
	mipsle  *string
	showVer *bool
	verbose *bool
}

var (
	flags mflag.FlagSet
	env   = &opts{
		FlagSet: &flags,
		arch:    flags.String("arch", runtime.GOARCH, "Set EdgeOS CPU architecture", false),
		dbug:    flags.Bool("debug", false, "Enable Debug mode", false),
		dryrun:  flags.Bool("dryrun", false, "Run config and data validation tests", false),
		file:    flags.String("f", "", "`<file>` # Load a config.boot file", true),
		help:    flags.Bool("h", false, "Display help", true),
		hostOS:  flags.String("os", runtime.GOOS, "Override native EdgeOS OS", false),
		mips64:  flags.String("mips64", "mips64", "Override target EdgeOS CPU architecture", false),
		mipsle:  flags.String("mipsle", "mipsle", "Override target EdgeOS CPU architecture", false),
		showVer: flags.Bool("version", false, "Show version", true),
		verbose: flags.Bool("v", false, "Verbose display", true),
	}
)

func main() {
	// initialization
	setArgs()
}

// basename removes directory components and file extensions.
func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Preserve everything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// cleanArgs removes flags when code is being tested
func cleanArgs(args []string) (r []string) {
	for _, a := range args {
		switch {
		case strings.HasPrefix(a, "-test"), strings.HasPrefix(a, "-convey"):
			continue
		default:
			r = append(r, a)
		}
	}
	return r
}

// setArgs retrieves arguments entered on the command line
func setArgs() {
	flags.Init(prog, mflag.ExitOnError)
	flags.Usage = env.PrintDefaults
	if env.Parse(cleanArgs((os.Args[1:]))) != nil {
		env.Usage()
		exitCmd(0)
	}

	switch {
	case *env.dbug:
		// screenLog("")
		// e.Dbug(*o.Dbug)
	case *env.help:
		env.Usage()
		exitCmd(0)
	case *env.dryrun:
		fmt.Println("dry run only, no actions will be executed!")
		exitCmd(0)
	case *env.verbose:
		// screenLog("")
	case *env.showVer:
		fmt.Printf(
			" Build Information:\n"+
				"   Version:\t\t\t%s\n"+
				"   Date:\t\t\t%s\n"+
				"   CPU:\t\t\t\t%v\n"+
				"   OS:\t\t\t\t%v\n"+
				"   Git hash:\t\t\t%v\n\n"+
				" This software comes with ABSOLUTELY NO WARRANTY.\n"+
				" %s is free software, and you are\n"+
				" welcome to redistribute it under the terms of\n"+
				" the Simplified BSD License.\n",
			version,
			build,
			architecture,
			hostOS,
			githash,
			prog,
		)
		exitCmd(0)
	}
}
