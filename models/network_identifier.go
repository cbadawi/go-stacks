package models

import (
	"encoding/json"
)

// NetworkIdentifier represents a NetworkIdentifier struct.
// The network_identifier specifies which network a particular object is associated with.
type NetworkIdentifier struct {
	// Blockchain name
	Blockchain string `json:"blockchain"`
	// If a blockchain has a specific chain-id or network identifier, it should go in this field. It is up to the client to determine which network-specific identifier is mainnet or testnet.
	Network string `json:"network"`
	// In blockchains with sharded state, the SubNetworkIdentifier is required to query some object on a specific shard. This identifier is optional for all non-sharded blockchains.
	SubNetworkIdentifier *SubNetworkIdentifier `json:"sub_network_identifier,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for NetworkIdentifier.
// It customizes the JSON marshaling process for NetworkIdentifier objects.
func (n *NetworkIdentifier) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(n.toMap())
}

// toMap converts the NetworkIdentifier object to a map representation for JSON marshaling.
func (n *NetworkIdentifier) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["blockchain"] = n.Blockchain
	structMap["network"] = n.Network
	if n.SubNetworkIdentifier != nil {
		structMap["sub_network_identifier"] = n.SubNetworkIdentifier
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkIdentifier.
// It customizes the JSON unmarshaling process for NetworkIdentifier objects.
func (n *NetworkIdentifier) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Blockchain           string                `json:"blockchain"`
		Network              string                `json:"network"`
		SubNetworkIdentifier *SubNetworkIdentifier `json:"sub_network_identifier,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	n.Blockchain = temp.Blockchain
	n.Network = temp.Network
	n.SubNetworkIdentifier = temp.SubNetworkIdentifier
	return nil
}
