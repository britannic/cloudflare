package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/britannic/mflag"
)

var (
	// updated by go build -ldflags
	architecture = "UNKNOWN"
	build        = "UNKNOWN"
	cftoken      = "UNKNOWN"
	githash      = "UNKNOWN"
	hostOS       = "UNKNOWN"
	version      = "UNKNOWN"
	env          = newOpts()
	exitCmd      = os.Exit
	prog         = basename(os.Args[0])

	// logFatalln = func(args ...interface{}) { log.Fatalln(args...); exitCmd(1) }

	// prefix  = fmt.Sprintf("%s: ", prog)
)

// logger implements an interface of logging functions that can be overriden
type logger interface {
	Fatalln(v ...interface{})
}

// opts struct for command line options and setting initial variables
type opts struct {
	*mflag.FlagSet
	// Cloudflare vars
	domain     *string
	email      *string
	apiKey     *string
	url        *string
	userSrvKey *string
	arch       *string
	dbug       *bool
	dryrun     *bool
	file       *string
	help       *bool
	hostOS     *string
	logger     logger
	mips64     *string
	mipsle     *string
	showVer    *bool
	verbose    *bool
}

func main() {
	env.Init(prog, mflag.ExitOnError)
	env.setArgs()
	api, err := env.getCFAPI()
	if err != nil {
		env.logger.Fatalln(err)
	}

	if *env.url != "" {
		api.BaseURL = *env.url
	}

	fmt.Println(api, err)
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

func newOpts() *opts {
	var (
		flags mflag.FlagSet
	)

	return &opts{
		FlagSet: &flags,
		logger:  log.New(os.Stderr, "", log.Ltime),
		// Cloudflare settings
		domain:     flags.String("domain", "", "domain registered with Cloudflare to update", true),
		email:      flags.String("email", "", "email address registered with Cloudflare", true),
		apiKey:     flags.String("apiKey", "", "Cloudflare API key", true),
		url:        flags.String("url", "", "Cloudflare API v4 URI", true),
		userSrvKey: flags.String("userSrvKey", "", "restricted endpoints Cloudflare API key, prefix \"v1.0-\", variable length", true),
		//----------------
		arch:    flags.String("arch", runtime.GOARCH, "set EdgeOS CPU architecture", false),
		dbug:    flags.Bool("debug", false, "enable Debug mode", false),
		dryrun:  flags.Bool("dryrun", false, "run config and data validation tests", true),
		file:    flags.String("f", "", "`<file>` # load a config.boot file", true),
		help:    flags.Bool("h", false, "display help", true),
		hostOS:  flags.String("os", runtime.GOOS, "override native EdgeOS OS", false),
		mips64:  flags.String("mips64", "mips64", "override target EdgeOS CPU architecture", false),
		mipsle:  flags.String("mipsle", "mipsle", "override target EdgeOS CPU architecture", false),
		showVer: flags.Bool("version", false, "show version", true),
		verbose: flags.Bool("v", false, "verbose display", true),
	}
}

// setArgs retrieves arguments entered on the command line
func (env *opts) setArgs() {
	if env.Parse(cleanArgs((os.Args[1:]))) != nil {
		exitCmd(0)
	}

	if *env.apiKey != "" {
		cftoken = *env.apiKey
	}

	if *env.dbug {
		fmt.Println("screenLog()")
	}

	if *env.help {
		env.PrintDefaults()
		exitCmd(0)
	}

	if *env.dryrun {
		fmt.Println("dry run only, no actions will be executed!")
		exitCmd(0)
	}

	if *env.verbose {
		fmt.Println("screenLog()")
	}

	if *env.showVer {
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
