package common

import (
	"flag"
	"fmt"
)

func PrintHelp() {
	fmt.Println("Help:")
	flag.PrintDefaults()
	fmt.Println("Examples:")
	fmt.Println("CLI:")
	fmt.Println("\tsnmp-browser -get -host 1.1.1.1 -oids 1.3.6.1.2.1.1.5.0")
	fmt.Println("\tsnmp-browser -get -host 1.1.1.1 -oids 1.3.6.1.2.1.1.1.0,1.3.6.1.2.1.1.4.0,1.3.6.1.2.1.1.3.0")
	fmt.Println("\tsnmp-browser -get -host 1.1.1.1 -oids 1.3.6.1.2.1.1.5.0 -port 161 -retries 3 -timeout 10 -verbose -version 2c")
	fmt.Println()
	fmt.Println("\tsnmp-browser -walk -host 10.24.16.69 -oids 1.3.6.1.2.1.2.2.1.5")
	fmt.Println("\tsnmp-browser -walk -host 10.24.16.69 -oids 1.3.6.1.2.1.2.2.1.5,1.3.6.1.2.1.31.1.1.1.15")
}

func PrintError(err error) {
	fmt.Printf("Error: %v\n", err)
}
