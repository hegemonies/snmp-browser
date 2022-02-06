package snmp

import "github.com/gosnmp/gosnmp"

const (
	snmpVersion1  = "1"
	snmpVersion2c = "2c"
)

// Конвертирует строку strVersion в gosnmp.SnmpVersion.
// strVersion может быть пустой строкой - то есть, значением по умолчанию, или "1", или "2c".
func ConvertSnmpVersion(strVersion string) gosnmp.SnmpVersion {
	if strVersion == snmpVersion1 {
		return gosnmp.Version1
	} else if strVersion == snmpVersion2c {
		return gosnmp.Version2c
	} else {
		return gosnmp.Version2c
	}
}
