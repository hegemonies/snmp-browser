package dto

import (
	"github.com/gosnmp/gosnmp"
	"strconv"
)

type SnmpResult struct {
	Name  string `json:"name"`
	Type_ string `json:"type"`
	Value string `json:"value"`
}

func MakeFromPdu(variable gosnmp.SnmpPDU) *SnmpResult {
	name := variable.Name
	_type := convertType(variable.Type)
	value := convertValue(variable.Value, variable.Type)

	return &SnmpResult{
		Name:  name,
		Type_: _type,
		Value: value,
	}
}

func convertType(_type gosnmp.Asn1BER) string {
	var convertedType string

	switch _type {
	case gosnmp.UnknownType:
		convertedType = "UnknownType"
	case gosnmp.Boolean:
		convertedType = "Boolean"
	case gosnmp.Integer:
		convertedType = "Integer"
	case gosnmp.BitString:
		convertedType = "BitString"
	case gosnmp.OctetString:
		convertedType = "OctetString"
	case gosnmp.Null:
		convertedType = "Null"
	case gosnmp.ObjectIdentifier:
		convertedType = "ObjectIdentifier"
	case gosnmp.ObjectDescription:
		convertedType = "ObjectDescription"
	case gosnmp.IPAddress:
		convertedType = "IPAddress"
	case gosnmp.Counter32:
		convertedType = "Counter32"
	case gosnmp.Gauge32:
		convertedType = "Gauge32"
	case gosnmp.TimeTicks:
		convertedType = "TimeTicks"
	case gosnmp.Opaque:
		convertedType = "Opaque"
	case gosnmp.NsapAddress:
		convertedType = "NsapAddress"
	case gosnmp.Counter64:
		convertedType = "Counter64"
	case gosnmp.Uinteger32:
		convertedType = "Uinteger32"
	case gosnmp.OpaqueFloat:
		convertedType = "OpaqueFloat"
	case gosnmp.OpaqueDouble:
		convertedType = "OpaqueDouble"
	case gosnmp.NoSuchObject:
		convertedType = "NoSuchObject"
	case gosnmp.NoSuchInstance:
		convertedType = "NoSuchInstance"
	}

	return convertedType
}

func convertValue(value interface{}, _type gosnmp.Asn1BER) string {
	var convertedValue string

	switch _type {
	case gosnmp.OctetString, gosnmp.BitString:
		convertedValue = string(value.([]byte))
	case gosnmp.Boolean:
		convertedValue = strconv.FormatBool(value.(bool))
	case gosnmp.Integer, gosnmp.Counter32, gosnmp.Gauge32, gosnmp.Counter64, gosnmp.Uinteger32:
		convertedValue = gosnmp.ToBigInt(value).String()
	case gosnmp.Null:
		convertedValue = "Null"
	case gosnmp.UnknownType:
		convertedValue = "UnknownType"
	case gosnmp.NoSuchObject:
		convertedValue = "NoSuchObject"
	case gosnmp.NoSuchInstance:
		convertedValue = "NoSuchInstance"
	default:
		convertedValue = ""
	}

	return convertedValue
}
