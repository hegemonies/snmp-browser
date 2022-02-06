package cli

import (
	"fmt"
	"golang/src/snmp-browser-server/main/app"
	"golang/src/snmp-browser-server/main/common"
	"golang/src/snmp-browser-server/main/dto"
	"golang/src/snmp-browser-server/main/snmp"
)

func Handle(options *app.Options) {

	if options.CliSnmpMethod == app.CliSnmpGetMethod {
		results, err := snmp.Get(
			options.TargetHostname,
			options.SnmpOids,
			options.SnmpCommunities,
			uint16(options.SnmpPort),
			options.SnmpVersion,
			options.SnmpRequestTimeout,
			options.SnmpRequestRetries,
		)

		if err != nil {
			common.PrintError(err)
		} else {
			printResults(results)
		}
	} else if options.CliSnmpMethod == app.CliSnmpWalkMethod {
		resultChan := make(chan dto.SnmpResult)

		go func() {
			err := snmp.Walk(
				options.TargetHostname,
				options.SnmpOids,
				options.SnmpCommunities,
				uint16(options.SnmpPort),
				options.SnmpVersion,
				options.SnmpRequestTimeout,
				options.SnmpRequestRetries,
				&resultChan,
			)
			if err != nil {
				common.PrintError(err)
			}
		}()

		printResultsFromChannel(resultChan)
	} else {
		common.PrintHelp()
	}
}

func printResult(result dto.SnmpResult) {
	fmt.Printf("%s : %s = %v\n", result.Name, result.Type_, result.Value)
}

func printResults(results []dto.SnmpResult) {
	for _, result := range results {
		printResult(result)
	}
}

func printResultsFromChannel(results chan dto.SnmpResult) {
	for result := range results {
		go printResult(result)
	}
}
