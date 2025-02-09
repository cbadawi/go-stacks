package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the CoreNodeInfoResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CoreNodeInfoResponse{}

// CoreNodeInfoResponse GET request that core node information
type CoreNodeInfoResponse struct {
	// identifies the version number for the networking communication, this should not change while a node is running, and will only change if there's an upgrade
	PeerVersion int32 `json:"peer_version"`
	// is a hash used to identify the burnchain view for a node. it incorporates bitcoin chain information and PoX information. nodes that disagree on this value will appear to each other as forks. this value will change after every block
	PoxConsensus string `json:"pox_consensus"`
	// latest bitcoin chain height
	BurnBlockHeight int32 `json:"burn_block_height"`
	// same as burn_consensus, but evaluated at stable_burn_block_height
	StablePoxConsensus string `json:"stable_pox_consensus"`
	// leftover from stacks 1.0, basically always burn_block_height - 1
	StableBurnBlockHeight int32 `json:"stable_burn_block_height"`
	// is a version descriptor
	ServerVersion string `json:"server_version"`
	// is similar to peer_version and will be used to differentiate between different testnets. this value will be different between mainnet and testnet. once launched, this value will not change
	NetworkId int32 `json:"network_id"`
	// same as network_id, but for bitcoin
	ParentNetworkId int32 `json:"parent_network_id"`
	// the latest Stacks chain height. Stacks forks can occur independent of the Bitcoin chain, that height doesn't increase 1-to-1 with the Bitcoin height
	StacksTipHeight int32 `json:"stacks_tip_height"`
	// the best known block hash for the Stack chain (not including any pending microblocks)
	StacksTip string `json:"stacks_tip"`
	// the burn chain (i.e., bitcoin) consensus hash at the time that stacks_tip was mined
	StacksTipConsensusHash string `json:"stacks_tip_consensus_hash"`
	// the latest microblock hash if any microblocks were processed. if no microblock has been processed for the current block, a 000.., hex array is returned
	UnanchoredTip string `json:"unanchored_tip"`
	// the block height at which the testnet network will be reset. not applicable for mainnet
	ExitAtBlockHeight int32 `json:"exit_at_block_height"`
}

type _CoreNodeInfoResponse CoreNodeInfoResponse

// NewCoreNodeInfoResponse instantiates a new CoreNodeInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreNodeInfoResponse(peerVersion int32, poxConsensus string, burnBlockHeight int32, stablePoxConsensus string, stableBurnBlockHeight int32, serverVersion string, networkId int32, parentNetworkId int32, stacksTipHeight int32, stacksTip string, stacksTipConsensusHash string, unanchoredTip string, exitAtBlockHeight int32) *CoreNodeInfoResponse {
	this := CoreNodeInfoResponse{}
	this.PeerVersion = peerVersion
	this.PoxConsensus = poxConsensus
	this.BurnBlockHeight = burnBlockHeight
	this.StablePoxConsensus = stablePoxConsensus
	this.StableBurnBlockHeight = stableBurnBlockHeight
	this.ServerVersion = serverVersion
	this.NetworkId = networkId
	this.ParentNetworkId = parentNetworkId
	this.StacksTipHeight = stacksTipHeight
	this.StacksTip = stacksTip
	this.StacksTipConsensusHash = stacksTipConsensusHash
	this.UnanchoredTip = unanchoredTip
	this.ExitAtBlockHeight = exitAtBlockHeight
	return &this
}

// NewCoreNodeInfoResponseWithDefaults instantiates a new CoreNodeInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreNodeInfoResponseWithDefaults() *CoreNodeInfoResponse {
	this := CoreNodeInfoResponse{}
	return &this
}

// GetPeerVersion returns the PeerVersion field value
func (o *CoreNodeInfoResponse) GetPeerVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PeerVersion
}

// GetPeerVersionOk returns a tuple with the PeerVersion field value
// and a boolean to check if the value has been set.
func (o *CoreNodeInfoResponse) GetPeerVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeerVersion, true
}

// SetPeerVersion sets field value
func (o *CoreNodeInfoResponse) SetPeerVersion(v int32) {
	o.PeerVersion = v
}

// GetPoxConsensus returns the PoxConsensus field value
func (o *CoreNodeInfoResponse) GetPoxConsensus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PoxConsensus
}

// GetPoxConsensusOk returns a tuple with the PoxConsensus field value
// and a boolean to check if the value has been set.
func (o *CoreNodeInfoResponse) GetPoxConsensusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoxConsensus, true
}

// SetPoxConsensus sets field value
func (o *CoreNodeInfoResponse) SetPoxConsensus(v string) {
	o.PoxConsensus = v
}

// GetBurnBlockHeight returns the BurnBlockHeight field value
func (o *CoreNodeInfoResponse) GetBurnBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnBlockHeight
}

// GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field value
// and a boolean to check if the value has been set.
func (o *CoreNodeInfoResponse) GetBurnBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockHeight, true
}

// SetBurnBlockHeight sets field value
func (o *CoreNodeInfoResponse) SetBurnBlockHeight(v int32) {
	o.BurnBlockHeight = v
}

// GetStablePoxConsensus returns the StablePoxConsensus field value
func (o *CoreNodeInfoResponse) GetStablePoxConsensus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StablePoxConsensus
}

// GetStablePoxConsensusOk returns a tuple with the StablePoxConsensus field value
// and a boolean to check if the value has been set.
func (o *CoreNodeInfoResponse) GetStablePoxConsensusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StablePoxConsensus, true
}

// SetStablePoxConsensus sets field value
func (o *CoreNodeInfoResponse) SetStablePoxConsensus(v string) {
	o.StablePoxConsensus = v
}

// GetStableBurnBlockHeight returns the StableBurnBlockHeight field value
func (o *CoreNodeInfoResponse) GetStableBurnBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StableBurnBlockHeight
}

// GetStableBurnBlockHeightOk returns a tuple with the StableBurnBlockHeight field value
// and a boolean to check if the value has been set.
func (o *CoreNodeInfoResponse) GetStableBurnBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StableBurnBlockHeight, true
}

// SetStableBurnBlockHeight sets field value
func (o *CoreNodeInfoResponse) SetStableBurnBlockHeight(v int32) {
	o.StableBurnBlockHeight = v
}

// GetServerVersion returns the ServerVersion field value
func (o *CoreNodeInfoResponse) GetServerVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerVersion
}

// GetServerVersionOk returns a tuple with the ServerVersion field value
// and a boolean to check if the value has been set.
func (o *CoreNodeInfoResponse) GetServerVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerVersion, true
}

// SetServerVersion sets field value
func (o *CoreNodeInfoResponse) SetServerVersion(v string) {
	o.ServerVersion = v
}

// GetNetworkId returns the NetworkId field value
func (o *CoreNodeInfoResponse) GetNetworkId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *CoreNodeInfoResponse) GetNetworkIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *CoreNodeInfoResponse) SetNetworkId(v int32) {
	o.NetworkId = v
}

// GetParentNetworkId returns the ParentNetworkId field value
func (o *CoreNodeInfoResponse) GetParentNetworkId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ParentNetworkId
}

// GetParentNetworkIdOk returns a tuple with the ParentNetworkId field value
// and a boolean to check if the value has been set.
func (o *CoreNodeInfoResponse) GetParentNetworkIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentNetworkId, true
}

// SetParentNetworkId sets field value
func (o *CoreNodeInfoResponse) SetParentNetworkId(v int32) {
	o.ParentNetworkId = v
}

// GetStacksTipHeight returns the StacksTipHeight field value
func (o *CoreNodeInfoResponse) GetStacksTipHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StacksTipHeight
}

// GetStacksTipHeightOk returns a tuple with the StacksTipHeight field value
// and a boolean to check if the value has been set.
func (o *CoreNodeInfoResponse) GetStacksTipHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StacksTipHeight, true
}

// SetStacksTipHeight sets field value
func (o *CoreNodeInfoResponse) SetStacksTipHeight(v int32) {
	o.StacksTipHeight = v
}

// GetStacksTip returns the StacksTip field value
func (o *CoreNodeInfoResponse) GetStacksTip() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StacksTip
}

// GetStacksTipOk returns a tuple with the StacksTip field value
// and a boolean to check if the value has been set.
func (o *CoreNodeInfoResponse) GetStacksTipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StacksTip, true
}

// SetStacksTip sets field value
func (o *CoreNodeInfoResponse) SetStacksTip(v string) {
	o.StacksTip = v
}

// GetStacksTipConsensusHash returns the StacksTipConsensusHash field value
func (o *CoreNodeInfoResponse) GetStacksTipConsensusHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StacksTipConsensusHash
}

// GetStacksTipConsensusHashOk returns a tuple with the StacksTipConsensusHash field value
// and a boolean to check if the value has been set.
func (o *CoreNodeInfoResponse) GetStacksTipConsensusHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StacksTipConsensusHash, true
}

// SetStacksTipConsensusHash sets field value
func (o *CoreNodeInfoResponse) SetStacksTipConsensusHash(v string) {
	o.StacksTipConsensusHash = v
}

// GetUnanchoredTip returns the UnanchoredTip field value
func (o *CoreNodeInfoResponse) GetUnanchoredTip() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnanchoredTip
}

// GetUnanchoredTipOk returns a tuple with the UnanchoredTip field value
// and a boolean to check if the value has been set.
func (o *CoreNodeInfoResponse) GetUnanchoredTipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnanchoredTip, true
}

// SetUnanchoredTip sets field value
func (o *CoreNodeInfoResponse) SetUnanchoredTip(v string) {
	o.UnanchoredTip = v
}

// GetExitAtBlockHeight returns the ExitAtBlockHeight field value
func (o *CoreNodeInfoResponse) GetExitAtBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExitAtBlockHeight
}

// GetExitAtBlockHeightOk returns a tuple with the ExitAtBlockHeight field value
// and a boolean to check if the value has been set.
func (o *CoreNodeInfoResponse) GetExitAtBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExitAtBlockHeight, true
}

// SetExitAtBlockHeight sets field value
func (o *CoreNodeInfoResponse) SetExitAtBlockHeight(v int32) {
	o.ExitAtBlockHeight = v
}

func (o CoreNodeInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreNodeInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["peer_version"] = o.PeerVersion
	toSerialize["pox_consensus"] = o.PoxConsensus
	toSerialize["burn_block_height"] = o.BurnBlockHeight
	toSerialize["stable_pox_consensus"] = o.StablePoxConsensus
	toSerialize["stable_burn_block_height"] = o.StableBurnBlockHeight
	toSerialize["server_version"] = o.ServerVersion
	toSerialize["network_id"] = o.NetworkId
	toSerialize["parent_network_id"] = o.ParentNetworkId
	toSerialize["stacks_tip_height"] = o.StacksTipHeight
	toSerialize["stacks_tip"] = o.StacksTip
	toSerialize["stacks_tip_consensus_hash"] = o.StacksTipConsensusHash
	toSerialize["unanchored_tip"] = o.UnanchoredTip
	toSerialize["exit_at_block_height"] = o.ExitAtBlockHeight
	return toSerialize, nil
}

func (o *CoreNodeInfoResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"peer_version",
		"pox_consensus",
		"burn_block_height",
		"stable_pox_consensus",
		"stable_burn_block_height",
		"server_version",
		"network_id",
		"parent_network_id",
		"stacks_tip_height",
		"stacks_tip",
		"stacks_tip_consensus_hash",
		"unanchored_tip",
		"exit_at_block_height",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCoreNodeInfoResponse := _CoreNodeInfoResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreNodeInfoResponse)

	if err != nil {
		return err
	}

	*o = CoreNodeInfoResponse(varCoreNodeInfoResponse)

	return err
}

type NullableCoreNodeInfoResponse struct {
	value *CoreNodeInfoResponse
	isSet bool
}

func (v NullableCoreNodeInfoResponse) Get() *CoreNodeInfoResponse {
	return v.value
}

func (v *NullableCoreNodeInfoResponse) Set(val *CoreNodeInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreNodeInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreNodeInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreNodeInfoResponse(val *CoreNodeInfoResponse) *NullableCoreNodeInfoResponse {
	return &NullableCoreNodeInfoResponse{value: val, isSet: true}
}

func (v NullableCoreNodeInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreNodeInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
