package models

import (
	"encoding/json"
)

// AddressTransactionWithTransfers represents a AddressTransactionWithTransfers struct.
// Transaction with STX transfers for a given address
type AddressTransactionWithTransfers struct {
	// Describes all transaction types on Stacks 2.0 blockchain
	Tx Transaction `json:"tx"`
	// Total sent from the given address, including the tx fee, in micro-STX as an integer string.
	StxSent string `json:"stx_sent"`
	// Total received by the given address in micro-STX as an integer string.
	StxReceived  string        `json:"stx_received"`
	StxTransfers []StxTransfer `json:"stx_transfers"`
	FtTransfers  []FtTransfer  `json:"ft_transfers,omitempty"`
	NftTransfers []NftTransfer `json:"nft_transfers,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for AddressTransactionWithTransfers.
// It customizes the JSON marshaling process for AddressTransactionWithTransfers objects.
func (a *AddressTransactionWithTransfers) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AddressTransactionWithTransfers object to a map representation for JSON marshaling.
func (a *AddressTransactionWithTransfers) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["tx"] = a.Tx
	structMap["stx_sent"] = a.StxSent
	structMap["stx_received"] = a.StxReceived
	structMap["stx_transfers"] = a.StxTransfers
	if a.FtTransfers != nil {
		structMap["ft_transfers"] = a.FtTransfers
	}
	if a.NftTransfers != nil {
		structMap["nft_transfers"] = a.NftTransfers
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddressTransactionWithTransfers.
// It customizes the JSON unmarshaling process for AddressTransactionWithTransfers objects.
func (a *AddressTransactionWithTransfers) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Tx           Transaction   `json:"tx"`
		StxSent      string        `json:"stx_sent"`
		StxReceived  string        `json:"stx_received"`
		StxTransfers []StxTransfer `json:"stx_transfers"`
		FtTransfers  []FtTransfer  `json:"ft_transfers,omitempty"`
		NftTransfers []NftTransfer `json:"nft_transfers,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.Tx = temp.Tx
	a.StxSent = temp.StxSent
	a.StxReceived = temp.StxReceived
	a.StxTransfers = temp.StxTransfers
	a.FtTransfers = temp.FtTransfers
	a.NftTransfers = temp.NftTransfers
	return nil
}
