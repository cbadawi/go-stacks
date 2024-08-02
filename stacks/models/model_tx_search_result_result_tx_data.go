package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the TxSearchResultResultTxData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TxSearchResultResultTxData{}

// TxSearchResultResultTxData Returns basic search result information about the requested id
type TxSearchResultResultTxData struct {
	// If the transaction lies within the canonical chain
	Canonical bool `json:"canonical"`
	// Refers to the hash of the block for searched transaction
	BlockHash     string `json:"block_hash"`
	BurnBlockTime int32  `json:"burn_block_time"`
	BlockHeight   int32  `json:"block_height"`
	TxType        string `json:"tx_type"`
}

type _TxSearchResultResultTxData TxSearchResultResultTxData

// NewTxSearchResultResultTxData instantiates a new TxSearchResultResultTxData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTxSearchResultResultTxData(canonical bool, blockHash string, burnBlockTime int32, blockHeight int32, txType string) *TxSearchResultResultTxData {
	this := TxSearchResultResultTxData{}
	this.Canonical = canonical
	this.BlockHash = blockHash
	this.BurnBlockTime = burnBlockTime
	this.BlockHeight = blockHeight
	this.TxType = txType
	return &this
}

// NewTxSearchResultResultTxDataWithDefaults instantiates a new TxSearchResultResultTxData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxSearchResultResultTxDataWithDefaults() *TxSearchResultResultTxData {
	this := TxSearchResultResultTxData{}
	return &this
}

// GetCanonical returns the Canonical field value
func (o *TxSearchResultResultTxData) GetCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value
// and a boolean to check if the value has been set.
func (o *TxSearchResultResultTxData) GetCanonicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canonical, true
}

// SetCanonical sets field value
func (o *TxSearchResultResultTxData) SetCanonical(v bool) {
	o.Canonical = v
}

// GetBlockHash returns the BlockHash field value
func (o *TxSearchResultResultTxData) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *TxSearchResultResultTxData) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *TxSearchResultResultTxData) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetBurnBlockTime returns the BurnBlockTime field value
func (o *TxSearchResultResultTxData) GetBurnBlockTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnBlockTime
}

// GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field value
// and a boolean to check if the value has been set.
func (o *TxSearchResultResultTxData) GetBurnBlockTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockTime, true
}

// SetBurnBlockTime sets field value
func (o *TxSearchResultResultTxData) SetBurnBlockTime(v int32) {
	o.BurnBlockTime = v
}

// GetBlockHeight returns the BlockHeight field value
func (o *TxSearchResultResultTxData) GetBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *TxSearchResultResultTxData) GetBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *TxSearchResultResultTxData) SetBlockHeight(v int32) {
	o.BlockHeight = v
}

// GetTxType returns the TxType field value
func (o *TxSearchResultResultTxData) GetTxType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxType
}

// GetTxTypeOk returns a tuple with the TxType field value
// and a boolean to check if the value has been set.
func (o *TxSearchResultResultTxData) GetTxTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxType, true
}

// SetTxType sets field value
func (o *TxSearchResultResultTxData) SetTxType(v string) {
	o.TxType = v
}

func (o TxSearchResultResultTxData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TxSearchResultResultTxData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["canonical"] = o.Canonical
	toSerialize["block_hash"] = o.BlockHash
	toSerialize["burn_block_time"] = o.BurnBlockTime
	toSerialize["block_height"] = o.BlockHeight
	toSerialize["tx_type"] = o.TxType
	return toSerialize, nil
}

func (o *TxSearchResultResultTxData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"canonical",
		"block_hash",
		"burn_block_time",
		"block_height",
		"tx_type",
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

	varTxSearchResultResultTxData := _TxSearchResultResultTxData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTxSearchResultResultTxData)

	if err != nil {
		return err
	}

	*o = TxSearchResultResultTxData(varTxSearchResultResultTxData)

	return err
}

type NullableTxSearchResultResultTxData struct {
	value *TxSearchResultResultTxData
	isSet bool
}

func (v NullableTxSearchResultResultTxData) Get() *TxSearchResultResultTxData {
	return v.value
}

func (v *NullableTxSearchResultResultTxData) Set(val *TxSearchResultResultTxData) {
	v.value = val
	v.isSet = true
}

func (v NullableTxSearchResultResultTxData) IsSet() bool {
	return v.isSet
}

func (v *NullableTxSearchResultResultTxData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxSearchResultResultTxData(val *TxSearchResultResultTxData) *NullableTxSearchResultResultTxData {
	return &NullableTxSearchResultResultTxData{value: val, isSet: true}
}

func (v NullableTxSearchResultResultTxData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxSearchResultResultTxData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
