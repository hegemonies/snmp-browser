package dto

type SnmpRequest struct {
	Method         string   `json:"method"`
	TargetHostname string   `json:"targetHostname"`
	Oids           []string `json:"oids"`
	Communities    []string `json:"communities"`
	Port           int      `json:"port"`
	Version        string   `json:"version"`
	Timeout        int      `json:"timeout"`
	Retries        int      `json:"retries"`
}
