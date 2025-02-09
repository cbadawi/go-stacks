package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaNetworkStatusResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaNetworkStatusResponse{}

// RosettaNetworkStatusResponse NetworkStatusResponse contains basic information about the node's view of a blockchain network. It is assumed that any BlockIdentifier.Index less than or equal to CurrentBlockIdentifier.Index can be queried. If a Rosetta implementation prunes historical state, it should populate the optional oldest_block_identifier field with the oldest block available to query. If this is not populated, it is assumed that the genesis_block_identifier is the oldest queryable block. If a Rosetta implementation performs some pre-sync before it is possible to query blocks, sync_status should be populated so that clients can still monitor healthiness. Without this field, it may appear that the implementation is stuck syncing and needs to be terminated.
type RosettaNetworkStatusResponse struct {
	CurrentBlockIdentifier RosettaBlockIdentifier `json:"current_block_identifier"`
	// The timestamp of the block in milliseconds since the Unix Epoch. The timestamp is stored in milliseconds because some blockchains produce blocks more often than once a second.
	CurrentBlockTimestamp  int32                         `json:"current_block_timestamp"`
	GenesisBlockIdentifier RosettaGenesisBlockIdentifier `json:"genesis_block_identifier"`
	OldestBlockIdentifier  *RosettaOldestBlockIdentifier `json:"oldest_block_identifier,omitempty"`
	SyncStatus             *RosettaSyncStatus            `json:"sync_status,omitempty"`
	// Peers information
	Peers []RosettaPeers `json:"peers"`
	// The latest burn block height
	CurrentBurnBlockHeight int32 `json:"current_burn_block_height"`
}

type _RosettaNetworkStatusResponse RosettaNetworkStatusResponse

// NewRosettaNetworkStatusResponse instantiates a new RosettaNetworkStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaNetworkStatusResponse(currentBlockIdentifier RosettaBlockIdentifier, currentBlockTimestamp int32, genesisBlockIdentifier RosettaGenesisBlockIdentifier, peers []RosettaPeers, currentBurnBlockHeight int32) *RosettaNetworkStatusResponse {
	this := RosettaNetworkStatusResponse{}
	this.CurrentBlockIdentifier = currentBlockIdentifier
	this.CurrentBlockTimestamp = currentBlockTimestamp
	this.GenesisBlockIdentifier = genesisBlockIdentifier
	this.Peers = peers
	this.CurrentBurnBlockHeight = currentBurnBlockHeight
	return &this
}

// NewRosettaNetworkStatusResponseWithDefaults instantiates a new RosettaNetworkStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaNetworkStatusResponseWithDefaults() *RosettaNetworkStatusResponse {
	this := RosettaNetworkStatusResponse{}
	return &this
}

// GetCurrentBlockIdentifier returns the CurrentBlockIdentifier field value
func (o *RosettaNetworkStatusResponse) GetCurrentBlockIdentifier() RosettaBlockIdentifier {
	if o == nil {
		var ret RosettaBlockIdentifier
		return ret
	}

	return o.CurrentBlockIdentifier
}

// GetCurrentBlockIdentifierOk returns a tuple with the CurrentBlockIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkStatusResponse) GetCurrentBlockIdentifierOk() (*RosettaBlockIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentBlockIdentifier, true
}

// SetCurrentBlockIdentifier sets field value
func (o *RosettaNetworkStatusResponse) SetCurrentBlockIdentifier(v RosettaBlockIdentifier) {
	o.CurrentBlockIdentifier = v
}

// GetCurrentBlockTimestamp returns the CurrentBlockTimestamp field value
func (o *RosettaNetworkStatusResponse) GetCurrentBlockTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentBlockTimestamp
}

// GetCurrentBlockTimestampOk returns a tuple with the CurrentBlockTimestamp field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkStatusResponse) GetCurrentBlockTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentBlockTimestamp, true
}

// SetCurrentBlockTimestamp sets field value
func (o *RosettaNetworkStatusResponse) SetCurrentBlockTimestamp(v int32) {
	o.CurrentBlockTimestamp = v
}

// GetGenesisBlockIdentifier returns the GenesisBlockIdentifier field value
func (o *RosettaNetworkStatusResponse) GetGenesisBlockIdentifier() RosettaGenesisBlockIdentifier {
	if o == nil {
		var ret RosettaGenesisBlockIdentifier
		return ret
	}

	return o.GenesisBlockIdentifier
}

// GetGenesisBlockIdentifierOk returns a tuple with the GenesisBlockIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkStatusResponse) GetGenesisBlockIdentifierOk() (*RosettaGenesisBlockIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GenesisBlockIdentifier, true
}

// SetGenesisBlockIdentifier sets field value
func (o *RosettaNetworkStatusResponse) SetGenesisBlockIdentifier(v RosettaGenesisBlockIdentifier) {
	o.GenesisBlockIdentifier = v
}

// GetOldestBlockIdentifier returns the OldestBlockIdentifier field value if set, zero value otherwise.
func (o *RosettaNetworkStatusResponse) GetOldestBlockIdentifier() RosettaOldestBlockIdentifier {
	if o == nil || utils.IsNil(o.OldestBlockIdentifier) {
		var ret RosettaOldestBlockIdentifier
		return ret
	}
	return *o.OldestBlockIdentifier
}

// GetOldestBlockIdentifierOk returns a tuple with the OldestBlockIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaNetworkStatusResponse) GetOldestBlockIdentifierOk() (*RosettaOldestBlockIdentifier, bool) {
	if o == nil || utils.IsNil(o.OldestBlockIdentifier) {
		return nil, false
	}
	return o.OldestBlockIdentifier, true
}

// HasOldestBlockIdentifier returns a boolean if a field has been set.
func (o *RosettaNetworkStatusResponse) HasOldestBlockIdentifier() bool {
	if o != nil && !utils.IsNil(o.OldestBlockIdentifier) {
		return true
	}

	return false
}

// SetOldestBlockIdentifier gets a reference to the given RosettaOldestBlockIdentifier and assigns it to the OldestBlockIdentifier field.
func (o *RosettaNetworkStatusResponse) SetOldestBlockIdentifier(v RosettaOldestBlockIdentifier) {
	o.OldestBlockIdentifier = &v
}

// GetSyncStatus returns the SyncStatus field value if set, zero value otherwise.
func (o *RosettaNetworkStatusResponse) GetSyncStatus() RosettaSyncStatus {
	if o == nil || utils.IsNil(o.SyncStatus) {
		var ret RosettaSyncStatus
		return ret
	}
	return *o.SyncStatus
}

// GetSyncStatusOk returns a tuple with the SyncStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaNetworkStatusResponse) GetSyncStatusOk() (*RosettaSyncStatus, bool) {
	if o == nil || utils.IsNil(o.SyncStatus) {
		return nil, false
	}
	return o.SyncStatus, true
}

// HasSyncStatus returns a boolean if a field has been set.
func (o *RosettaNetworkStatusResponse) HasSyncStatus() bool {
	if o != nil && !utils.IsNil(o.SyncStatus) {
		return true
	}

	return false
}

// SetSyncStatus gets a reference to the given RosettaSyncStatus and assigns it to the SyncStatus field.
func (o *RosettaNetworkStatusResponse) SetSyncStatus(v RosettaSyncStatus) {
	o.SyncStatus = &v
}

// GetPeers returns the Peers field value
func (o *RosettaNetworkStatusResponse) GetPeers() []RosettaPeers {
	if o == nil {
		var ret []RosettaPeers
		return ret
	}

	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkStatusResponse) GetPeersOk() ([]RosettaPeers, bool) {
	if o == nil {
		return nil, false
	}
	return o.Peers, true
}

// SetPeers sets field value
func (o *RosettaNetworkStatusResponse) SetPeers(v []RosettaPeers) {
	o.Peers = v
}

// GetCurrentBurnBlockHeight returns the CurrentBurnBlockHeight field value
func (o *RosettaNetworkStatusResponse) GetCurrentBurnBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentBurnBlockHeight
}

// GetCurrentBurnBlockHeightOk returns a tuple with the CurrentBurnBlockHeight field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkStatusResponse) GetCurrentBurnBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentBurnBlockHeight, true
}

// SetCurrentBurnBlockHeight sets field value
func (o *RosettaNetworkStatusResponse) SetCurrentBurnBlockHeight(v int32) {
	o.CurrentBurnBlockHeight = v
}

func (o RosettaNetworkStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaNetworkStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["current_block_identifier"] = o.CurrentBlockIdentifier
	toSerialize["current_block_timestamp"] = o.CurrentBlockTimestamp
	toSerialize["genesis_block_identifier"] = o.GenesisBlockIdentifier
	if !utils.IsNil(o.OldestBlockIdentifier) {
		toSerialize["oldest_block_identifier"] = o.OldestBlockIdentifier
	}
	if !utils.IsNil(o.SyncStatus) {
		toSerialize["sync_status"] = o.SyncStatus
	}
	toSerialize["peers"] = o.Peers
	toSerialize["current_burn_block_height"] = o.CurrentBurnBlockHeight
	return toSerialize, nil
}

func (o *RosettaNetworkStatusResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"current_block_identifier",
		"current_block_timestamp",
		"genesis_block_identifier",
		"peers",
		"current_burn_block_height",
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

	varRosettaNetworkStatusResponse := _RosettaNetworkStatusResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaNetworkStatusResponse)

	if err != nil {
		return err
	}

	*o = RosettaNetworkStatusResponse(varRosettaNetworkStatusResponse)

	return err
}

type NullableRosettaNetworkStatusResponse struct {
	value *RosettaNetworkStatusResponse
	isSet bool
}

func (v NullableRosettaNetworkStatusResponse) Get() *RosettaNetworkStatusResponse {
	return v.value
}

func (v *NullableRosettaNetworkStatusResponse) Set(val *RosettaNetworkStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaNetworkStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaNetworkStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaNetworkStatusResponse(val *RosettaNetworkStatusResponse) *NullableRosettaNetworkStatusResponse {
	return &NullableRosettaNetworkStatusResponse{value: val, isSet: true}
}

func (v NullableRosettaNetworkStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaNetworkStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
