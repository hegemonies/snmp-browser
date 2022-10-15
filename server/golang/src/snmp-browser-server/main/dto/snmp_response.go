package dto

type SnmpResponse struct {
	Results []SnmpResult `json:"results"`
}

type SnmpPingResponse struct {
	Available bool `json:"available"`
}
