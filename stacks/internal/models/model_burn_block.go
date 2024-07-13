package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BurnBlock type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BurnBlock{}

// BurnBlock A burn block
type BurnBlock struct {
	// Unix timestamp (in seconds) indicating when this block was mined.
	BurnBlockTime float32 `json:"burn_block_time"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined.
	BurnBlockTimeIso string `json:"burn_block_time_iso"`
	// Hash of the anchor chain block
	BurnBlockHash string `json:"burn_block_hash"`
	// Height of the anchor chain block
	BurnBlockHeight int32 `json:"burn_block_height"`
	// Hashes of the Stacks blocks included in the burn block
	StacksBlocks []string `json:"stacks_blocks"`
	// Average time between blocks in seconds. Returns 0 if there is only one block in the burn block.
	AvgBlockTime float32 `json:"avg_block_time"`
	// Total number of transactions in the Stacks blocks associated with this burn block
	TotalTxCount int32 `json:"total_tx_count"`
}

type _BurnBlock BurnBlock

// NewBurnBlock instantiates a new BurnBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBurnBlock(burnBlockTime float32, burnBlockTimeIso string, burnBlockHash string, burnBlockHeight int32, stacksBlocks []string, avgBlockTime float32, totalTxCount int32) *BurnBlock {
	this := BurnBlock{}
	this.BurnBlockTime = burnBlockTime
	this.BurnBlockTimeIso = burnBlockTimeIso
	this.BurnBlockHash = burnBlockHash
	this.BurnBlockHeight = burnBlockHeight
	this.StacksBlocks = stacksBlocks
	this.AvgBlockTime = avgBlockTime
	this.TotalTxCount = totalTxCount
	return &this
}

// NewBurnBlockWithDefaults instantiates a new BurnBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBurnBlockWithDefaults() *BurnBlock {
	this := BurnBlock{}
	return &this
}

// GetBurnBlockTime returns the BurnBlockTime field value
func (o *BurnBlock) GetBurnBlockTime() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BurnBlockTime
}

// GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field value
// and a boolean to check if the value has been set.
func (o *BurnBlock) GetBurnBlockTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockTime, true
}

// SetBurnBlockTime sets field value
func (o *BurnBlock) SetBurnBlockTime(v float32) {
	o.BurnBlockTime = v
}

// GetBurnBlockTimeIso returns the BurnBlockTimeIso field value
func (o *BurnBlock) GetBurnBlockTimeIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BurnBlockTimeIso
}

// GetBurnBlockTimeIsoOk returns a tuple with the BurnBlockTimeIso field value
// and a boolean to check if the value has been set.
func (o *BurnBlock) GetBurnBlockTimeIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockTimeIso, true
}

// SetBurnBlockTimeIso sets field value
func (o *BurnBlock) SetBurnBlockTimeIso(v string) {
	o.BurnBlockTimeIso = v
}

// GetBurnBlockHash returns the BurnBlockHash field value
func (o *BurnBlock) GetBurnBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BurnBlockHash
}

// GetBurnBlockHashOk returns a tuple with the BurnBlockHash field value
// and a boolean to check if the value has been set.
func (o *BurnBlock) GetBurnBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockHash, true
}

// SetBurnBlockHash sets field value
func (o *BurnBlock) SetBurnBlockHash(v string) {
	o.BurnBlockHash = v
}

// GetBurnBlockHeight returns the BurnBlockHeight field value
func (o *BurnBlock) GetBurnBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnBlockHeight
}

// GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field value
// and a boolean to check if the value has been set.
func (o *BurnBlock) GetBurnBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockHeight, true
}

// SetBurnBlockHeight sets field value
func (o *BurnBlock) SetBurnBlockHeight(v int32) {
	o.BurnBlockHeight = v
}

// GetStacksBlocks returns the StacksBlocks field value
func (o *BurnBlock) GetStacksBlocks() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.StacksBlocks
}

// GetStacksBlocksOk returns a tuple with the StacksBlocks field value
// and a boolean to check if the value has been set.
func (o *BurnBlock) GetStacksBlocksOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StacksBlocks, true
}

// SetStacksBlocks sets field value
func (o *BurnBlock) SetStacksBlocks(v []string) {
	o.StacksBlocks = v
}

// GetAvgBlockTime returns the AvgBlockTime field value
func (o *BurnBlock) GetAvgBlockTime() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AvgBlockTime
}

// GetAvgBlockTimeOk returns a tuple with the AvgBlockTime field value
// and a boolean to check if the value has been set.
func (o *BurnBlock) GetAvgBlockTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgBlockTime, true
}

// SetAvgBlockTime sets field value
func (o *BurnBlock) SetAvgBlockTime(v float32) {
	o.AvgBlockTime = v
}

// GetTotalTxCount returns the TotalTxCount field value
func (o *BurnBlock) GetTotalTxCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalTxCount
}

// GetTotalTxCountOk returns a tuple with the TotalTxCount field value
// and a boolean to check if the value has been set.
func (o *BurnBlock) GetTotalTxCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTxCount, true
}

// SetTotalTxCount sets field value
func (o *BurnBlock) SetTotalTxCount(v int32) {
	o.TotalTxCount = v
}

func (o BurnBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BurnBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["burn_block_time"] = o.BurnBlockTime
	toSerialize["burn_block_time_iso"] = o.BurnBlockTimeIso
	toSerialize["burn_block_hash"] = o.BurnBlockHash
	toSerialize["burn_block_height"] = o.BurnBlockHeight
	toSerialize["stacks_blocks"] = o.StacksBlocks
	toSerialize["avg_block_time"] = o.AvgBlockTime
	toSerialize["total_tx_count"] = o.TotalTxCount
	return toSerialize, nil
}

func (o *BurnBlock) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"burn_block_time",
		"burn_block_time_iso",
		"burn_block_hash",
		"burn_block_height",
		"stacks_blocks",
		"avg_block_time",
		"total_tx_count",
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

	varBurnBlock := _BurnBlock{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBurnBlock)

	if err != nil {
		return err
	}

	*o = BurnBlock(varBurnBlock)

	return err
}

type NullableBurnBlock struct {
	value *BurnBlock
	isSet bool
}

func (v NullableBurnBlock) Get() *BurnBlock {
	return v.value
}

func (v *NullableBurnBlock) Set(val *BurnBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableBurnBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableBurnBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBurnBlock(val *BurnBlock) *NullableBurnBlock {
	return &NullableBurnBlock{value: val, isSet: true}
}

func (v NullableBurnBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBurnBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
