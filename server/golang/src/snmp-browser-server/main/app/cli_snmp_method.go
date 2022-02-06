package app

type CliSnmpMethod uint8

const (
	CliSnmpUnknownMethod CliSnmpMethod = 0x0
	CliSnmpGetMethod     CliSnmpMethod = 0x1
	CliSnmpWalkMethod    CliSnmpMethod = iota
)
