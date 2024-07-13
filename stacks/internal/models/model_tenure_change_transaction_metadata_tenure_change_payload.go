package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TenureChangeTransactionMetadataTenureChangePayload type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TenureChangeTransactionMetadataTenureChangePayload{}

// TenureChangeTransactionMetadataTenureChangePayload struct for TenureChangeTransactionMetadataTenureChangePayload
type TenureChangeTransactionMetadataTenureChangePayload struct {
	// Consensus hash of this tenure. Corresponds to the sortition in which the miner of this block was chosen.
	TenureConsensusHash string `json:"tenure_consensus_hash"`
	// Consensus hash of the previous tenure. Corresponds to the sortition of the previous winning block-commit.
	PrevTenureConsensusHash string `json:"prev_tenure_consensus_hash"`
	// Current consensus hash on the underlying burnchain. Corresponds to the last-seen sortition.
	BurnViewConsensusHash string `json:"burn_view_consensus_hash"`
	// (Hex string) Stacks Block hash
	PreviousTenureEnd string `json:"previous_tenure_end"`
	// The number of blocks produced in the previous tenure.
	PreviousTenureBlocks float32 `json:"previous_tenure_blocks"`
	// Cause of change in mining tenure. Depending on cause, tenure can be ended or extended.
	Cause string `json:"cause"`
	// (Hex string) The ECDSA public key hash of the current tenure.
	PubkeyHash string `json:"pubkey_hash"`
}

type _TenureChangeTransactionMetadataTenureChangePayload TenureChangeTransactionMetadataTenureChangePayload

// NewTenureChangeTransactionMetadataTenureChangePayload instantiates a new TenureChangeTransactionMetadataTenureChangePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenureChangeTransactionMetadataTenureChangePayload(tenureConsensusHash string, prevTenureConsensusHash string, burnViewConsensusHash string, previousTenureEnd string, previousTenureBlocks float32, cause string, pubkeyHash string) *TenureChangeTransactionMetadataTenureChangePayload {
	this := TenureChangeTransactionMetadataTenureChangePayload{}
	this.TenureConsensusHash = tenureConsensusHash
	this.PrevTenureConsensusHash = prevTenureConsensusHash
	this.BurnViewConsensusHash = burnViewConsensusHash
	this.PreviousTenureEnd = previousTenureEnd
	this.PreviousTenureBlocks = previousTenureBlocks
	this.Cause = cause
	this.PubkeyHash = pubkeyHash
	return &this
}

// NewTenureChangeTransactionMetadataTenureChangePayloadWithDefaults instantiates a new TenureChangeTransactionMetadataTenureChangePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenureChangeTransactionMetadataTenureChangePayloadWithDefaults() *TenureChangeTransactionMetadataTenureChangePayload {
	this := TenureChangeTransactionMetadataTenureChangePayload{}
	return &this
}

// GetTenureConsensusHash returns the TenureConsensusHash field value
func (o *TenureChangeTransactionMetadataTenureChangePayload) GetTenureConsensusHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenureConsensusHash
}

// GetTenureConsensusHashOk returns a tuple with the TenureConsensusHash field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransactionMetadataTenureChangePayload) GetTenureConsensusHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenureConsensusHash, true
}

// SetTenureConsensusHash sets field value
func (o *TenureChangeTransactionMetadataTenureChangePayload) SetTenureConsensusHash(v string) {
	o.TenureConsensusHash = v
}

// GetPrevTenureConsensusHash returns the PrevTenureConsensusHash field value
func (o *TenureChangeTransactionMetadataTenureChangePayload) GetPrevTenureConsensusHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrevTenureConsensusHash
}

// GetPrevTenureConsensusHashOk returns a tuple with the PrevTenureConsensusHash field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransactionMetadataTenureChangePayload) GetPrevTenureConsensusHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrevTenureConsensusHash, true
}

// SetPrevTenureConsensusHash sets field value
func (o *TenureChangeTransactionMetadataTenureChangePayload) SetPrevTenureConsensusHash(v string) {
	o.PrevTenureConsensusHash = v
}

// GetBurnViewConsensusHash returns the BurnViewConsensusHash field value
func (o *TenureChangeTransactionMetadataTenureChangePayload) GetBurnViewConsensusHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BurnViewConsensusHash
}

// GetBurnViewConsensusHashOk returns a tuple with the BurnViewConsensusHash field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransactionMetadataTenureChangePayload) GetBurnViewConsensusHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnViewConsensusHash, true
}

// SetBurnViewConsensusHash sets field value
func (o *TenureChangeTransactionMetadataTenureChangePayload) SetBurnViewConsensusHash(v string) {
	o.BurnViewConsensusHash = v
}

// GetPreviousTenureEnd returns the PreviousTenureEnd field value
func (o *TenureChangeTransactionMetadataTenureChangePayload) GetPreviousTenureEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousTenureEnd
}

// GetPreviousTenureEndOk returns a tuple with the PreviousTenureEnd field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransactionMetadataTenureChangePayload) GetPreviousTenureEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousTenureEnd, true
}

// SetPreviousTenureEnd sets field value
func (o *TenureChangeTransactionMetadataTenureChangePayload) SetPreviousTenureEnd(v string) {
	o.PreviousTenureEnd = v
}

// GetPreviousTenureBlocks returns the PreviousTenureBlocks field value
func (o *TenureChangeTransactionMetadataTenureChangePayload) GetPreviousTenureBlocks() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PreviousTenureBlocks
}

// GetPreviousTenureBlocksOk returns a tuple with the PreviousTenureBlocks field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransactionMetadataTenureChangePayload) GetPreviousTenureBlocksOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousTenureBlocks, true
}

// SetPreviousTenureBlocks sets field value
func (o *TenureChangeTransactionMetadataTenureChangePayload) SetPreviousTenureBlocks(v float32) {
	o.PreviousTenureBlocks = v
}

// GetCause returns the Cause field value
func (o *TenureChangeTransactionMetadataTenureChangePayload) GetCause() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransactionMetadataTenureChangePayload) GetCauseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *TenureChangeTransactionMetadataTenureChangePayload) SetCause(v string) {
	o.Cause = v
}

// GetPubkeyHash returns the PubkeyHash field value
func (o *TenureChangeTransactionMetadataTenureChangePayload) GetPubkeyHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PubkeyHash
}

// GetPubkeyHashOk returns a tuple with the PubkeyHash field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransactionMetadataTenureChangePayload) GetPubkeyHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PubkeyHash, true
}

// SetPubkeyHash sets field value
func (o *TenureChangeTransactionMetadataTenureChangePayload) SetPubkeyHash(v string) {
	o.PubkeyHash = v
}

func (o TenureChangeTransactionMetadataTenureChangePayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenureChangeTransactionMetadataTenureChangePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenure_consensus_hash"] = o.TenureConsensusHash
	toSerialize["prev_tenure_consensus_hash"] = o.PrevTenureConsensusHash
	toSerialize["burn_view_consensus_hash"] = o.BurnViewConsensusHash
	toSerialize["previous_tenure_end"] = o.PreviousTenureEnd
	toSerialize["previous_tenure_blocks"] = o.PreviousTenureBlocks
	toSerialize["cause"] = o.Cause
	toSerialize["pubkey_hash"] = o.PubkeyHash
	return toSerialize, nil
}

func (o *TenureChangeTransactionMetadataTenureChangePayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenure_consensus_hash",
		"prev_tenure_consensus_hash",
		"burn_view_consensus_hash",
		"previous_tenure_end",
		"previous_tenure_blocks",
		"cause",
		"pubkey_hash",
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

	varTenureChangeTransactionMetadataTenureChangePayload := _TenureChangeTransactionMetadataTenureChangePayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTenureChangeTransactionMetadataTenureChangePayload)

	if err != nil {
		return err
	}

	*o = TenureChangeTransactionMetadataTenureChangePayload(varTenureChangeTransactionMetadataTenureChangePayload)

	return err
}

type NullableTenureChangeTransactionMetadataTenureChangePayload struct {
	value *TenureChangeTransactionMetadataTenureChangePayload
	isSet bool
}

func (v NullableTenureChangeTransactionMetadataTenureChangePayload) Get() *TenureChangeTransactionMetadataTenureChangePayload {
	return v.value
}

func (v *NullableTenureChangeTransactionMetadataTenureChangePayload) Set(val *TenureChangeTransactionMetadataTenureChangePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableTenureChangeTransactionMetadataTenureChangePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableTenureChangeTransactionMetadataTenureChangePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenureChangeTransactionMetadataTenureChangePayload(val *TenureChangeTransactionMetadataTenureChangePayload) *NullableTenureChangeTransactionMetadataTenureChangePayload {
	return &NullableTenureChangeTransactionMetadataTenureChangePayload{value: val, isSet: true}
}

func (v NullableTenureChangeTransactionMetadataTenureChangePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenureChangeTransactionMetadataTenureChangePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
