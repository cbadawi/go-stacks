package models

import (
	"encoding/json"
)

// CoreNodeInfoResponse represents a CoreNodeInfoResponse struct.
// GET request that core node information
type CoreNodeInfoResponse struct {
	// identifies the version number for the networking communication, this should not change while a node is running, and will only change if there's an upgrade
	PeerVersion int `json:"peer_version"`
	// is a hash used to identify the burnchain view for a node. it incorporates bitcoin chain information and PoX information. nodes that disagree on this value will appear to each other as forks. this value will change after every block
	PoxConsensus string `json:"pox_consensus"`
	// latest bitcoin chain height
	BurnBlockHeight int `json:"burn_block_height"`
	// same as burn_consensus, but evaluated at stable_burn_block_height
	StablePoxConsensus string `json:"stable_pox_consensus"`
	// leftover from stacks 1.0, basically always burn_block_height - 1
	StableBurnBlockHeight int `json:"stable_burn_block_height"`
	// is a version descriptor
	ServerVersion string `json:"server_version"`
	// is similar to peer_version and will be used to differentiate between different testnets. this value will be different between mainnet and testnet. once launched, this value will not change
	NetworkId int `json:"network_id"`
	// same as network_id, but for bitcoin
	ParentNetworkId int `json:"parent_network_id"`
	// the latest Stacks chain height. Stacks forks can occur independent of the Bitcoin chain, that height doesn't increase 1-to-1 with the Bitcoin height
	StacksTipHeight int `json:"stacks_tip_height"`
	// the best known block hash for the Stack chain (not including any pending microblocks)
	StacksTip string `json:"stacks_tip"`
	// the burn chain (i.e., bitcoin) consensus hash at the time that stacks_tip was mined
	StacksTipConsensusHash string `json:"stacks_tip_consensus_hash"`
	// the latest microblock hash if any microblocks were processed. if no microblock has been processed for the current block, a 000.., hex array is returned
	UnanchoredTip string `json:"unanchored_tip"`
	// the block height at which the testnet network will be reset. not applicable for mainnet
	ExitAtBlockHeight int `json:"exit_at_block_height"`
}

// MarshalJSON implements the json.Marshaler interface for CoreNodeInfoResponse.
// It customizes the JSON marshaling process for CoreNodeInfoResponse objects.
func (c *CoreNodeInfoResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CoreNodeInfoResponse object to a map representation for JSON marshaling.
func (c *CoreNodeInfoResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["peer_version"] = c.PeerVersion
	structMap["pox_consensus"] = c.PoxConsensus
	structMap["burn_block_height"] = c.BurnBlockHeight
	structMap["stable_pox_consensus"] = c.StablePoxConsensus
	structMap["stable_burn_block_height"] = c.StableBurnBlockHeight
	structMap["server_version"] = c.ServerVersion
	structMap["network_id"] = c.NetworkId
	structMap["parent_network_id"] = c.ParentNetworkId
	structMap["stacks_tip_height"] = c.StacksTipHeight
	structMap["stacks_tip"] = c.StacksTip
	structMap["stacks_tip_consensus_hash"] = c.StacksTipConsensusHash
	structMap["unanchored_tip"] = c.UnanchoredTip
	structMap["exit_at_block_height"] = c.ExitAtBlockHeight
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CoreNodeInfoResponse.
// It customizes the JSON unmarshaling process for CoreNodeInfoResponse objects.
func (c *CoreNodeInfoResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		PeerVersion            int    `json:"peer_version"`
		PoxConsensus           string `json:"pox_consensus"`
		BurnBlockHeight        int    `json:"burn_block_height"`
		StablePoxConsensus     string `json:"stable_pox_consensus"`
		StableBurnBlockHeight  int    `json:"stable_burn_block_height"`
		ServerVersion          string `json:"server_version"`
		NetworkId              int    `json:"network_id"`
		ParentNetworkId        int    `json:"parent_network_id"`
		StacksTipHeight        int    `json:"stacks_tip_height"`
		StacksTip              string `json:"stacks_tip"`
		StacksTipConsensusHash string `json:"stacks_tip_consensus_hash"`
		UnanchoredTip          string `json:"unanchored_tip"`
		ExitAtBlockHeight      int    `json:"exit_at_block_height"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.PeerVersion = temp.PeerVersion
	c.PoxConsensus = temp.PoxConsensus
	c.BurnBlockHeight = temp.BurnBlockHeight
	c.StablePoxConsensus = temp.StablePoxConsensus
	c.StableBurnBlockHeight = temp.StableBurnBlockHeight
	c.ServerVersion = temp.ServerVersion
	c.NetworkId = temp.NetworkId
	c.ParentNetworkId = temp.ParentNetworkId
	c.StacksTipHeight = temp.StacksTipHeight
	c.StacksTip = temp.StacksTip
	c.StacksTipConsensusHash = temp.StacksTipConsensusHash
	c.UnanchoredTip = temp.UnanchoredTip
	c.ExitAtBlockHeight = temp.ExitAtBlockHeight
	return nil
}
