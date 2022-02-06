package constants

const (
	AppModeCli             = "cli"
	AppModeHttp            = "http"
	AppModeHelpDescription = "app mode may be are cli or http; by default is cli"

	SnmpMethodGet      = "get"
	SnmpMethodGetHelp  = "set SNMP get method"
	SnmpMethodWalk     = "walk"
	SnmpMethodWalkHelp = "set SNMP walk method"

	SnmpOptionHostname        = "host"
	SnmpOptionHostnameDefault = ""
	SnmpOptionHostnameHelp    = "set ip/hostname"

	SnmpOptionPort        = "port"
	SnmpOptionDefaultPort = 161
	SnmpOptionPortHelp    = "set SNMP port; default is 161"

	SnmpOptionCommunities      = "communities"
	SnmpOptionCommunityDefault = "public"
	SnmpOptionCommunityHelp    = "set SNMP communities; example: 'public,private'; default is public"

	SnmpOptionVersion        = "version"
	SnmpOptionVersionDefault = "2c"
	SnmpOptionVersionHelp    = "set SNMP version; example: 1 or 2c; default is 2c"

	SnmpOptionOids       = "oids"
	SnmpOptionOidDefault = ""
	SnmpOptionOidHelp    = "set SNMP oids; example: '1.3.6.1.2.1.2.2.1.10,1.3.6.1.2.1.2.2.1.15'"

	SnmpOptionTimeout        = "timeout"
	SnmpOptionTimeoutDefault = 5
	SnmpOptionTimeoutHelp    = "set timeout request in sec; default is 5 sec"

	SnmpOptionNumberRetries        = "retries"
	SnmpOptionNumberRetriesDefault = 0
	SnmpOptionNumberRetriesHelp    = "set number of retries request; default is 3"

	Help            = "help"
	HelpDescription = "print help"

	Verbose     = "verbose"
	VerboseHelp = "print arguments"

	HttpServerHostname     = "serverhost"
	HttpServerHostnameHelp = "set up of http server hostname"

	HttpServerPort     = "serverport"
	HttpServerPortHelp = "set up of http server port"
)
