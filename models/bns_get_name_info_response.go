package models

import (
	"encoding/json"
)

// BnsGetNameInfoResponse represents a BnsGetNameInfoResponse struct.
// Get name details
type BnsGetNameInfoResponse struct {
	Address      string  `json:"address"`
	Blockchain   string  `json:"blockchain"`
	ExpireBlock  *int    `json:"expire_block,omitempty"`
	GracePeriod  *int    `json:"grace_period,omitempty"`
	LastTxid     string  `json:"last_txid"`
	Resolver     *string `json:"resolver,omitempty"`
	Status       string  `json:"status"`
	Zonefile     string  `json:"zonefile"`
	ZonefileHash string  `json:"zonefile_hash"`
}

// MarshalJSON implements the json.Marshaler interface for BnsGetNameInfoResponse.
// It customizes the JSON marshaling process for BnsGetNameInfoResponse objects.
func (b *BnsGetNameInfoResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BnsGetNameInfoResponse object to a map representation for JSON marshaling.
func (b *BnsGetNameInfoResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["address"] = b.Address
	structMap["blockchain"] = b.Blockchain
	if b.ExpireBlock != nil {
		structMap["expire_block"] = b.ExpireBlock
	}
	if b.GracePeriod != nil {
		structMap["grace_period"] = b.GracePeriod
	}
	structMap["last_txid"] = b.LastTxid
	if b.Resolver != nil {
		structMap["resolver"] = b.Resolver
	}
	structMap["status"] = b.Status
	structMap["zonefile"] = b.Zonefile
	structMap["zonefile_hash"] = b.ZonefileHash
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BnsGetNameInfoResponse.
// It customizes the JSON unmarshaling process for BnsGetNameInfoResponse objects.
func (b *BnsGetNameInfoResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Address      string  `json:"address"`
		Blockchain   string  `json:"blockchain"`
		ExpireBlock  *int    `json:"expire_block,omitempty"`
		GracePeriod  *int    `json:"grace_period,omitempty"`
		LastTxid     string  `json:"last_txid"`
		Resolver     *string `json:"resolver,omitempty"`
		Status       string  `json:"status"`
		Zonefile     string  `json:"zonefile"`
		ZonefileHash string  `json:"zonefile_hash"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	b.Address = temp.Address
	b.Blockchain = temp.Blockchain
	b.ExpireBlock = temp.ExpireBlock
	b.GracePeriod = temp.GracePeriod
	b.LastTxid = temp.LastTxid
	b.Resolver = temp.Resolver
	b.Status = temp.Status
	b.Zonefile = temp.Zonefile
	b.ZonefileHash = temp.ZonefileHash
	return nil
}
