package models

import (
	"encoding/json"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the ContractSearchResultResultTxData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ContractSearchResultResultTxData{}

// ContractSearchResultResultTxData Returns basic search result information about the requested id
type ContractSearchResultResultTxData struct {
	// If the transaction lies within the canonical chain
	Canonical *bool `json:"canonical,omitempty"`
	// Refers to the hash of the block for searched transaction
	BlockHash     *string `json:"block_hash,omitempty"`
	BurnBlockTime *int32  `json:"burn_block_time,omitempty"`
	BlockHeight   *int32  `json:"block_height,omitempty"`
	TxType        *string `json:"tx_type,omitempty"`
	// Corresponding tx_id for smart_contract
	TxId *string `json:"tx_id,omitempty"`
}

// NewContractSearchResultResultTxData instantiates a new ContractSearchResultResultTxData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractSearchResultResultTxData() *ContractSearchResultResultTxData {
	this := ContractSearchResultResultTxData{}
	return &this
}

// NewContractSearchResultResultTxDataWithDefaults instantiates a new ContractSearchResultResultTxData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractSearchResultResultTxDataWithDefaults() *ContractSearchResultResultTxData {
	this := ContractSearchResultResultTxData{}
	return &this
}

// GetCanonical returns the Canonical field value if set, zero value otherwise.
func (o *ContractSearchResultResultTxData) GetCanonical() bool {
	if o == nil || utils.IsNil(o.Canonical) {
		var ret bool
		return ret
	}
	return *o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractSearchResultResultTxData) GetCanonicalOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Canonical) {
		return nil, false
	}
	return o.Canonical, true
}

// HasCanonical returns a boolean if a field has been set.
func (o *ContractSearchResultResultTxData) HasCanonical() bool {
	if o != nil && !utils.IsNil(o.Canonical) {
		return true
	}

	return false
}

// SetCanonical gets a reference to the given bool and assigns it to the Canonical field.
func (o *ContractSearchResultResultTxData) SetCanonical(v bool) {
	o.Canonical = &v
}

// GetBlockHash returns the BlockHash field value if set, zero value otherwise.
func (o *ContractSearchResultResultTxData) GetBlockHash() string {
	if o == nil || utils.IsNil(o.BlockHash) {
		var ret string
		return ret
	}
	return *o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractSearchResultResultTxData) GetBlockHashOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BlockHash) {
		return nil, false
	}
	return o.BlockHash, true
}

// HasBlockHash returns a boolean if a field has been set.
func (o *ContractSearchResultResultTxData) HasBlockHash() bool {
	if o != nil && !utils.IsNil(o.BlockHash) {
		return true
	}

	return false
}

// SetBlockHash gets a reference to the given string and assigns it to the BlockHash field.
func (o *ContractSearchResultResultTxData) SetBlockHash(v string) {
	o.BlockHash = &v
}

// GetBurnBlockTime returns the BurnBlockTime field value if set, zero value otherwise.
func (o *ContractSearchResultResultTxData) GetBurnBlockTime() int32 {
	if o == nil || utils.IsNil(o.BurnBlockTime) {
		var ret int32
		return ret
	}
	return *o.BurnBlockTime
}

// GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractSearchResultResultTxData) GetBurnBlockTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.BurnBlockTime) {
		return nil, false
	}
	return o.BurnBlockTime, true
}

// HasBurnBlockTime returns a boolean if a field has been set.
func (o *ContractSearchResultResultTxData) HasBurnBlockTime() bool {
	if o != nil && !utils.IsNil(o.BurnBlockTime) {
		return true
	}

	return false
}

// SetBurnBlockTime gets a reference to the given int32 and assigns it to the BurnBlockTime field.
func (o *ContractSearchResultResultTxData) SetBurnBlockTime(v int32) {
	o.BurnBlockTime = &v
}

// GetBlockHeight returns the BlockHeight field value if set, zero value otherwise.
func (o *ContractSearchResultResultTxData) GetBlockHeight() int32 {
	if o == nil || utils.IsNil(o.BlockHeight) {
		var ret int32
		return ret
	}
	return *o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractSearchResultResultTxData) GetBlockHeightOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.BlockHeight) {
		return nil, false
	}
	return o.BlockHeight, true
}

// HasBlockHeight returns a boolean if a field has been set.
func (o *ContractSearchResultResultTxData) HasBlockHeight() bool {
	if o != nil && !utils.IsNil(o.BlockHeight) {
		return true
	}

	return false
}

// SetBlockHeight gets a reference to the given int32 and assigns it to the BlockHeight field.
func (o *ContractSearchResultResultTxData) SetBlockHeight(v int32) {
	o.BlockHeight = &v
}

// GetTxType returns the TxType field value if set, zero value otherwise.
func (o *ContractSearchResultResultTxData) GetTxType() string {
	if o == nil || utils.IsNil(o.TxType) {
		var ret string
		return ret
	}
	return *o.TxType
}

// GetTxTypeOk returns a tuple with the TxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractSearchResultResultTxData) GetTxTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TxType) {
		return nil, false
	}
	return o.TxType, true
}

// HasTxType returns a boolean if a field has been set.
func (o *ContractSearchResultResultTxData) HasTxType() bool {
	if o != nil && !utils.IsNil(o.TxType) {
		return true
	}

	return false
}

// SetTxType gets a reference to the given string and assigns it to the TxType field.
func (o *ContractSearchResultResultTxData) SetTxType(v string) {
	o.TxType = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *ContractSearchResultResultTxData) GetTxId() string {
	if o == nil || utils.IsNil(o.TxId) {
		var ret string
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractSearchResultResultTxData) GetTxIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *ContractSearchResultResultTxData) HasTxId() bool {
	if o != nil && !utils.IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given string and assigns it to the TxId field.
func (o *ContractSearchResultResultTxData) SetTxId(v string) {
	o.TxId = &v
}

func (o ContractSearchResultResultTxData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractSearchResultResultTxData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Canonical) {
		toSerialize["canonical"] = o.Canonical
	}
	if !utils.IsNil(o.BlockHash) {
		toSerialize["block_hash"] = o.BlockHash
	}
	if !utils.IsNil(o.BurnBlockTime) {
		toSerialize["burn_block_time"] = o.BurnBlockTime
	}
	if !utils.IsNil(o.BlockHeight) {
		toSerialize["block_height"] = o.BlockHeight
	}
	if !utils.IsNil(o.TxType) {
		toSerialize["tx_type"] = o.TxType
	}
	if !utils.IsNil(o.TxId) {
		toSerialize["tx_id"] = o.TxId
	}
	return toSerialize, nil
}

type NullableContractSearchResultResultTxData struct {
	value *ContractSearchResultResultTxData
	isSet bool
}

func (v NullableContractSearchResultResultTxData) Get() *ContractSearchResultResultTxData {
	return v.value
}

func (v *NullableContractSearchResultResultTxData) Set(val *ContractSearchResultResultTxData) {
	v.value = val
	v.isSet = true
}

func (v NullableContractSearchResultResultTxData) IsSet() bool {
	return v.isSet
}

func (v *NullableContractSearchResultResultTxData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractSearchResultResultTxData(val *ContractSearchResultResultTxData) *NullableContractSearchResultResultTxData {
	return &NullableContractSearchResultResultTxData{value: val, isSet: true}
}

func (v NullableContractSearchResultResultTxData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractSearchResultResultTxData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
