package models

import (
	"encoding/json"
)

// RosettaNetworkStatusResponse represents a RosettaNetworkStatusResponse struct.
// NetworkStatusResponse contains basic information about the node's view of a blockchain network. It is assumed that any BlockIdentifier.Index less than or equal to CurrentBlockIdentifier.Index can be queried. If a Rosetta implementation prunes historical state, it should populate the optional oldest_block_identifier field with the oldest block available to query. If this is not populated, it is assumed that the genesis_block_identifier is the oldest queryable block. If a Rosetta implementation performs some pre-sync before it is possible to query blocks, sync_status should be populated so that clients can still monitor healthiness. Without this field, it may appear that the implementation is stuck syncing and needs to be terminated.
type RosettaNetworkStatusResponse struct {
	// The block_identifier uniquely identifies a block in a particular network.
	CurrentBlockIdentifier RosettaBlockIdentifier `json:"current_block_identifier"`
	// The timestamp of the block in milliseconds since the Unix Epoch. The timestamp is stored in milliseconds because some blockchains produce blocks more often than once a second.
	CurrentBlockTimestamp int `json:"current_block_timestamp"`
	// The block_identifier uniquely identifies a block in a particular network.
	GenesisBlockIdentifier RosettaGenesisBlockIdentifier `json:"genesis_block_identifier"`
	// The block_identifier uniquely identifies a block in a particular network.
	OldestBlockIdentifier *RosettaOldestBlockIdentifier `json:"oldest_block_identifier,omitempty"`
	// SyncStatus is used to provide additional context about an implementation's sync status. It is often used to indicate that an implementation is healthy when it cannot be queried until some sync phase occurs. If an implementation is immediately queryable, this model is often not populated.
	SyncStatus *RosettaSyncStatus `json:"sync_status,omitempty"`
	// Peers information
	Peers []RosettaPeers `json:"peers"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaNetworkStatusResponse.
// It customizes the JSON marshaling process for RosettaNetworkStatusResponse objects.
func (r *RosettaNetworkStatusResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaNetworkStatusResponse object to a map representation for JSON marshaling.
func (r *RosettaNetworkStatusResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["current_block_identifier"] = r.CurrentBlockIdentifier
	structMap["current_block_timestamp"] = r.CurrentBlockTimestamp
	structMap["genesis_block_identifier"] = r.GenesisBlockIdentifier
	if r.OldestBlockIdentifier != nil {
		structMap["oldest_block_identifier"] = r.OldestBlockIdentifier
	}
	if r.SyncStatus != nil {
		structMap["sync_status"] = r.SyncStatus
	}
	structMap["peers"] = r.Peers
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaNetworkStatusResponse.
// It customizes the JSON unmarshaling process for RosettaNetworkStatusResponse objects.
func (r *RosettaNetworkStatusResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		CurrentBlockIdentifier RosettaBlockIdentifier        `json:"current_block_identifier"`
		CurrentBlockTimestamp  int                           `json:"current_block_timestamp"`
		GenesisBlockIdentifier RosettaGenesisBlockIdentifier `json:"genesis_block_identifier"`
		OldestBlockIdentifier  *RosettaOldestBlockIdentifier `json:"oldest_block_identifier,omitempty"`
		SyncStatus             *RosettaSyncStatus            `json:"sync_status,omitempty"`
		Peers                  []RosettaPeers                `json:"peers"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.CurrentBlockIdentifier = temp.CurrentBlockIdentifier
	r.CurrentBlockTimestamp = temp.CurrentBlockTimestamp
	r.GenesisBlockIdentifier = temp.GenesisBlockIdentifier
	r.OldestBlockIdentifier = temp.OldestBlockIdentifier
	r.SyncStatus = temp.SyncStatus
	r.Peers = temp.Peers
	return nil
}
