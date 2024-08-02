package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the PoxCycle type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PoxCycle{}

// PoxCycle struct for PoxCycle
type PoxCycle struct {
	BlockHeight        int32  `json:"block_height"`
	IndexBlockHash     string `json:"index_block_hash"`
	CycleNumber        int32  `json:"cycle_number"`
	TotalWeight        int32  `json:"total_weight"`
	TotalStackedAmount string `json:"total_stacked_amount"`
	TotalSigners       int32  `json:"total_signers"`
}

type _PoxCycle PoxCycle

// NewPoxCycle instantiates a new PoxCycle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoxCycle(blockHeight int32, indexBlockHash string, cycleNumber int32, totalWeight int32, totalStackedAmount string, totalSigners int32) *PoxCycle {
	this := PoxCycle{}
	this.BlockHeight = blockHeight
	this.IndexBlockHash = indexBlockHash
	this.CycleNumber = cycleNumber
	this.TotalWeight = totalWeight
	this.TotalStackedAmount = totalStackedAmount
	this.TotalSigners = totalSigners
	return &this
}

// NewPoxCycleWithDefaults instantiates a new PoxCycle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoxCycleWithDefaults() *PoxCycle {
	this := PoxCycle{}
	return &this
}

// GetBlockHeight returns the BlockHeight field value
func (o *PoxCycle) GetBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *PoxCycle) GetBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *PoxCycle) SetBlockHeight(v int32) {
	o.BlockHeight = v
}

// GetIndexBlockHash returns the IndexBlockHash field value
func (o *PoxCycle) GetIndexBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndexBlockHash
}

// GetIndexBlockHashOk returns a tuple with the IndexBlockHash field value
// and a boolean to check if the value has been set.
func (o *PoxCycle) GetIndexBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexBlockHash, true
}

// SetIndexBlockHash sets field value
func (o *PoxCycle) SetIndexBlockHash(v string) {
	o.IndexBlockHash = v
}

// GetCycleNumber returns the CycleNumber field value
func (o *PoxCycle) GetCycleNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CycleNumber
}

// GetCycleNumberOk returns a tuple with the CycleNumber field value
// and a boolean to check if the value has been set.
func (o *PoxCycle) GetCycleNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CycleNumber, true
}

// SetCycleNumber sets field value
func (o *PoxCycle) SetCycleNumber(v int32) {
	o.CycleNumber = v
}

// GetTotalWeight returns the TotalWeight field value
func (o *PoxCycle) GetTotalWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalWeight
}

// GetTotalWeightOk returns a tuple with the TotalWeight field value
// and a boolean to check if the value has been set.
func (o *PoxCycle) GetTotalWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalWeight, true
}

// SetTotalWeight sets field value
func (o *PoxCycle) SetTotalWeight(v int32) {
	o.TotalWeight = v
}

// GetTotalStackedAmount returns the TotalStackedAmount field value
func (o *PoxCycle) GetTotalStackedAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalStackedAmount
}

// GetTotalStackedAmountOk returns a tuple with the TotalStackedAmount field value
// and a boolean to check if the value has been set.
func (o *PoxCycle) GetTotalStackedAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalStackedAmount, true
}

// SetTotalStackedAmount sets field value
func (o *PoxCycle) SetTotalStackedAmount(v string) {
	o.TotalStackedAmount = v
}

// GetTotalSigners returns the TotalSigners field value
func (o *PoxCycle) GetTotalSigners() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalSigners
}

// GetTotalSignersOk returns a tuple with the TotalSigners field value
// and a boolean to check if the value has been set.
func (o *PoxCycle) GetTotalSignersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSigners, true
}

// SetTotalSigners sets field value
func (o *PoxCycle) SetTotalSigners(v int32) {
	o.TotalSigners = v
}

func (o PoxCycle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoxCycle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["block_height"] = o.BlockHeight
	toSerialize["index_block_hash"] = o.IndexBlockHash
	toSerialize["cycle_number"] = o.CycleNumber
	toSerialize["total_weight"] = o.TotalWeight
	toSerialize["total_stacked_amount"] = o.TotalStackedAmount
	toSerialize["total_signers"] = o.TotalSigners
	return toSerialize, nil
}

func (o *PoxCycle) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"block_height",
		"index_block_hash",
		"cycle_number",
		"total_weight",
		"total_stacked_amount",
		"total_signers",
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

	varPoxCycle := _PoxCycle{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPoxCycle)

	if err != nil {
		return err
	}

	*o = PoxCycle(varPoxCycle)

	return err
}

type NullablePoxCycle struct {
	value *PoxCycle
	isSet bool
}

func (v NullablePoxCycle) Get() *PoxCycle {
	return v.value
}

func (v *NullablePoxCycle) Set(val *PoxCycle) {
	v.value = val
	v.isSet = true
}

func (v NullablePoxCycle) IsSet() bool {
	return v.isSet
}

func (v *NullablePoxCycle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoxCycle(val *PoxCycle) *NullablePoxCycle {
	return &NullablePoxCycle{value: val, isSet: true}
}

func (v NullablePoxCycle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoxCycle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
