package app

import (
	"flag"
	"fmt"
	"golang/src/snmp-browser-server/main/common"
	"golang/src/snmp-browser-server/main/constants"
	"golang/src/snmp-browser-server/main/snmp"
	"os"
	"strings"
	"time"

	"github.com/gosnmp/gosnmp"
)

type Options struct {
	Mode          Mode
	CliSnmpMethod CliSnmpMethod

	TargetHostname     string // or ip
	SnmpOids           []string
	SnmpCommunities    []string
	SnmpPort           int
	SnmpVersion        gosnmp.SnmpVersion
	SnmpRequestTimeout time.Duration // in seconds
	SnmpRequestRetries int

	HttpHostname string
	HttpPort     int
}

func NewOptions() *Options {
	return &Options{}
}

func (options *Options) FillFromCommandLine() {

	var isCliMode bool
	flag.BoolVar(&isCliMode, constants.AppModeCli, false, "set cli mode; "+constants.AppModeHelpDescription)

	var isHttpMode bool
	flag.BoolVar(&isHttpMode, constants.AppModeHttp, false, "set http mode; "+constants.AppModeHelpDescription)

	var hostname string
	flag.StringVar(
		&hostname,
		constants.SnmpOptionHostname,
		constants.SnmpOptionHostnameDefault,
		constants.SnmpOptionHostnameHelp,
	)

	var oidsOption string
	flag.StringVar(
		&oidsOption,
		constants.SnmpOptionOids,
		constants.SnmpOptionOidDefault,
		constants.SnmpOptionOidHelp,
	)

	var communitiesOption string
	flag.StringVar(
		&communitiesOption,
		constants.SnmpOptionCommunities,
		constants.SnmpOptionCommunityDefault,
		constants.SnmpOptionCommunityHelp,
	)

	var port int
	flag.IntVar(
		&port,
		constants.SnmpOptionPort,
		constants.SnmpOptionDefaultPort,
		constants.SnmpOptionPortHelp,
	)

	var versionOption string
	flag.StringVar(
		&versionOption,
		constants.SnmpOptionVersion,
		constants.SnmpOptionVersionDefault,
		constants.SnmpOptionVersionHelp,
	)

	var timeout int
	flag.IntVar(
		&timeout,
		constants.SnmpOptionTimeout,
		constants.SnmpOptionTimeoutDefault,
		constants.SnmpOptionTimeoutHelp,
	)

	var numberRetries int
	flag.IntVar(
		&numberRetries,
		constants.SnmpOptionNumberRetries,
		constants.SnmpOptionNumberRetriesDefault,
		constants.SnmpOptionNumberRetriesHelp,
	)

	var isSnmpGetMethod bool
	flag.BoolVar(&isSnmpGetMethod, constants.SnmpMethodGet, false, constants.SnmpMethodGetHelp)

	var isSnmpWalkMethod bool
	flag.BoolVar(&isSnmpWalkMethod, constants.SnmpMethodWalk, false, constants.SnmpMethodWalkHelp)

	var help bool
	flag.BoolVar(&help, constants.Help, false, constants.HelpDescription)

	var verbose bool
	flag.BoolVar(&verbose, constants.Verbose, false, constants.VerboseHelp)

	var httpHostname string
	flag.StringVar(&httpHostname, constants.HttpServerHostname, "127.0.0.1", constants.HttpServerHostnameHelp)

	var httpPort int
	flag.IntVar(&httpPort, constants.HttpServerPort, 7000, constants.HttpServerPortHelp)

	flag.Parse()

	if help {
		common.PrintHelp()
		os.Exit(0)
	}

	var mode Mode
	if isCliMode {
		mode = CliAppMode
	} else if isHttpMode {
		mode = HttpAppMode
	} else {
		mode = CliAppMode
	}

	if mode == constants.AppModeHttp {
		options.Mode = HttpAppMode
		options.HttpHostname = httpHostname
		options.HttpPort = httpPort
	} else {
		if hostname == constants.SnmpOptionHostnameDefault {
			fmt.Printf("Error: hostname must be set\n\n")
			common.PrintHelp()
			os.Exit(1)
		}

		oids := strings.Split(oidsOption, ",")

		if len(oids) == 0 {
			fmt.Printf("Error: oids must be set\n\n")
			common.PrintHelp()
			os.Exit(1)
		}

		communities := strings.Split(communitiesOption, ",")
		if len(communities) == 0 {
			communities = []string{"public"}
		}
		version := snmp.ConvertSnmpVersion(versionOption)

		options.Mode = CliAppMode
		if isSnmpGetMethod {
			options.CliSnmpMethod = CliSnmpGetMethod
		} else if isSnmpWalkMethod {
			options.CliSnmpMethod = CliSnmpWalkMethod
		} else {
			options.CliSnmpMethod = CliSnmpUnknownMethod
		}
		options.TargetHostname = hostname
		options.SnmpOids = oids
		options.SnmpCommunities = communities
		options.SnmpPort = port
		options.SnmpVersion = version
		options.SnmpRequestTimeout = time.Duration(timeout) * time.Second
		options.SnmpRequestRetries = numberRetries

		if verbose {
			options.print()
		}
	}
}

func (options *Options) print() {
	fmt.Printf("Mode=%v\n", options.Mode)
	fmt.Printf("TargetHostname=%v\n", options.TargetHostname)
	fmt.Printf("SnmpOids=%v\n", options.SnmpOids)
	fmt.Printf("SnmpCommunities=%v\n", options.SnmpCommunities)
	fmt.Printf("SnmpVersion=%v\n", options.SnmpVersion)
	fmt.Printf("SnmpPort=%v\n", options.SnmpPort)
	fmt.Printf("SnmpRequestTimeout=%v\n", options.SnmpRequestTimeout)
	fmt.Printf("SnmpRequestRetries=%v\n", options.SnmpRequestRetries)
	fmt.Printf("CliSnmpMethod=%v\n", options.CliSnmpMethod)
	fmt.Printf("HttpHostname=%v\n", options.HttpHostname)
	fmt.Printf("HttpPort=%v\n", options.HttpPort)
}
