package main

import (
	"fmt"
	"github.com/gosnmp/gosnmp"
	"log"
	"time"
)

func testGet() {
	gosnmp.Default.Target = "10.24.16.69"

	err := gosnmp.Default.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}

	defer gosnmp.Default.Conn.Close()

	oids := []string{
		"1.3.6.1.2.1.1.5.0",
	}

	result, err2 := gosnmp.Default.Get(oids)
	if err != nil {
		log.Fatalf("Get() err: %v", err2)
	}

	for _, variable := range result.Variables {
		err = printPdu(variable)
		if err != nil {
			log.Fatalf("PrintPdu() err: %v", err)
		}
	}
}

func testWalk() {
	snmpClient := &gosnmp.GoSNMP{
		Target:             "10.24.16.199",
		Port:               161,
		Community:          "public",
		Version:            gosnmp.Version2c,
		Timeout:            time.Duration(5) * time.Second,
		Retries:            0,
		ExponentialTimeout: false,
	}

	err := snmpClient.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}

	defer snmpClient.Conn.Close()

	oid := "1.3.6.1.2.1.2.2.1.3" // ifType

	err = snmpClient.BulkWalk(oid, printPdu)
	if err != nil {
		log.Fatalf("SNMP WALK err: %v", err)
	}
}

func printPdu(pdu gosnmp.SnmpPDU) error {
	fmt.Printf("%s = %v : ", pdu.Name, pdu.Type)

	switch pdu.Type {
	case gosnmp.OctetString:
		fmt.Printf("%s\n", string(pdu.Value.([]byte)))
	default:
		fmt.Printf("%d\n", gosnmp.ToBigInt(pdu.Value))
	}

	return nil
}
