package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"strings"

	"github.com/britannic/mflag"
	cloudflare "github.com/cloudflare/cloudflare-go"
)

var (
	// updated by go build -ldflags
	architecture = "UNKNOWN"
	build        = "UNKNOWN"
	cftoken      = "UNKNOWN"
	githash      = "UNKNOWN"
	hostOS       = "UNKNOWN"
	version      = "UNKNOWN"
	exitCmd      = os.Exit
	env          = newOpts()
	prog         = basename(os.Args[0])

	// prefix  = fmt.Sprintf("%s: ", prog)
)

// logger interface wraps logging functions that can be overriden
type logger interface {
	Fatalf(s string, v ...interface{})
	Fatalln(v ...interface{})
	// Panic(v ...interface{})
}

// opts struct for command line options and setting initial variables
type opts struct {
	*mflag.FlagSet
	// Cloudflare vars
	cf         cfAPI
	api        *cloudflare.API
	domain     *string
	email      *string
	apiKey     *string
	apiURL     *string
	userSrvKey *string
	arch       *string
	dbug       *bool
	dryrun     *bool
	file       *string
	help       *bool
	hostOS     *string
	log        logger
	mips64     *string
	mipsle     *string
	ok         bool
	showVer    *bool
	verbose    *bool
}

func main() {
	env.Init(prog, mflag.ExitOnError)
	env.setArgs()

	if err := env.getCFAPI(); err != nil {
		env.log.Fatalln(err)
	}

	fmt.Println(env.routableIP("udp", "8.8.8.8:80"))

	// Fetch a slice of all zones available to this account.
	// zones, err := api.ListZones()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, z := range zones {
	// 	fmt.Println(z.Name)
	// }

	if env.ok {
		zoneID, err := env.cf.ZoneIDByName("orbc2.org")
		if err != nil {
			env.log.Fatalln(err)
		}

		r, err := env.getDNSRecord(zoneID, "sylvania.kh.orbc2.org")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Name: %s\nID: %s\nProxied: %v\n", r.Name, r.ID, r.Proxied)
	}

	// api.UpdateDNSRecord(zoneID string, recordID string, rr cloudflare.DNSRecord)

}

// routableIP returns the WAN address
func (o *opts) routableIP(network, address string) (string, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return "", fmt.Errorf("net.Dial: %v", err)
	}

	ip := conn.LocalAddr().(*net.UDPAddr).String()
	ndx := strings.LastIndex(ip, ":")
	conn.Close()
	// return ip
	return ip[0:ndx], nil
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
		cf:      &cloudflare.API{},
		FlagSet: &flags,
		log:     log.New(os.Stderr, "", log.Ltime),
		// Cloudflare settings
		domain:     flags.String("domain", "", "domain registered with Cloudflare to update", true),
		email:      flags.String("email", "", "email address registered with Cloudflare", true),
		apiKey:     flags.String("apikey", "", "Cloudflare API key", true),
		apiURL:     flags.String("url", "", "Cloudflare API v4 URI", true),
		userSrvKey: flags.String("userkey", "", "restricted endpoints Cloudflare API key, prefix \"v1.0-\", variable length", true),
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
func (o *opts) setArgs() {
	if o.Parse(cleanArgs((os.Args[1:]))) != nil {
		exitCmd(0)
	}

	if *o.apiKey != "" {
		cftoken = *o.apiKey
	}

	if *o.dbug {
		fmt.Println("screenLog()")
	}

	if *o.help {
		o.PrintDefaults()
		exitCmd(0)
	}

	if *o.dryrun {
		fmt.Println("dry run only, no actions will be executed!")
		exitCmd(0)
	}

	if *o.verbose {
		fmt.Println("screenLog()")
	}

	if *o.showVer {
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
