package snmp

import (
	"github.com/gosnmp/gosnmp"
	"golang/src/snmp-browser-server/main/dto"
	"time"
)

func Ping(snmpClient *gosnmp.GoSNMP) (bool, error) {
	sysObjectIdOid := "1.3.6.1.2.1.1.2.0"
	oids := []string{sysObjectIdOid}

	_, err := snmpClient.Get(oids)
	if err != nil {
		return false, err
	}

	return true, nil
}

func Get(
	hostname string,
	snmpOids []string,
	snmpCommunities []string,
	snmpPort uint16,
	snmpVersion gosnmp.SnmpVersion,
	timeout time.Duration,
	numberRetry int,
) ([]dto.SnmpResult, error) {

	var convertedResult []dto.SnmpResult

	for _, community := range snmpCommunities {
		snmpClient := &gosnmp.GoSNMP{
			Target:             hostname,
			Port:               snmpPort,
			Community:          community,
			Version:            snmpVersion,
			Timeout:            timeout,
			Retries:            numberRetry,
			ExponentialTimeout: false,
		}

		err := snmpClient.Connect()
		if err != nil {
			return nil, err
		}
		defer snmpClient.Conn.Close()

		communityIsOk, err2 := Ping(snmpClient)
		if err2 != nil {
			return nil, err2
		}

		if communityIsOk {
			snmpResult, err3 := snmpClient.Get(snmpOids)
			if err3 != nil {
				return nil, err3
			}

			for _, variable := range snmpResult.Variables {
				convertedResult = append(convertedResult, *dto.MakeFromPdu(variable))
			}

			break
		}
	}

	return convertedResult, nil
}

func Walk(
	hostname string,
	snmpOids []string,
	snmpCommunities []string,
	snmpPort uint16,
	snmpVersion gosnmp.SnmpVersion,
	timeout time.Duration,
	numberRetry int,
	resultChannel *chan dto.SnmpResult,
) error {

	for _, community := range snmpCommunities {
		snmpClient := &gosnmp.GoSNMP{
			Target:             hostname,
			Port:               snmpPort,
			Community:          community,
			Version:            snmpVersion,
			Timeout:            timeout * time.Second,
			Retries:            numberRetry,
			ExponentialTimeout: false,
		}

		err := snmpClient.Connect()
		if err != nil {
			return err
		}
		defer snmpClient.Conn.Close()

		communityIsOk, err2 := Ping(snmpClient)
		if err2 != nil {
			return err2
		}

		if communityIsOk {

			for _, oid := range snmpOids {

				snmpResultReceiver := func(pdu gosnmp.SnmpPDU) error {
					*resultChannel <- *dto.MakeFromPdu(pdu)
					return nil
				}

				err3 := snmpClient.BulkWalk(oid, snmpResultReceiver)
				if err3 != nil {
					return err3
				}
			}

			break
		}
	}

	close(*resultChannel)

	return nil
}
