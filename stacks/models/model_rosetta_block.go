package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaBlock type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaBlock{}

// RosettaBlock Blocks contain an array of Transactions that occurred at a particular BlockIdentifier. A hard requirement for blocks returned by Rosetta implementations is that they MUST be inalterable: once a client has requested and received a block identified by a specific BlockIndentifier, all future calls for that same BlockIdentifier must return the same block contents.
type RosettaBlock struct {
	BlockIdentifier       RosettaBlockIdentifier       `json:"block_identifier"`
	ParentBlockIdentifier RosettaParentBlockIdentifier `json:"parent_block_identifier"`
	// The timestamp of the block in milliseconds since the Unix Epoch. The timestamp is stored in milliseconds because some blockchains produce blocks more often than once a second.
	Timestamp int32 `json:"timestamp"`
	// All the transactions in the block
	Transactions []RosettaTransaction `json:"transactions"`
	Metadata     RosettaBlockMetadata `json:"metadata"`
}

type _RosettaBlock RosettaBlock

// NewRosettaBlock instantiates a new RosettaBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaBlock(blockIdentifier RosettaBlockIdentifier, parentBlockIdentifier RosettaParentBlockIdentifier, timestamp int32, transactions []RosettaTransaction, metadata RosettaBlockMetadata) *RosettaBlock {
	this := RosettaBlock{}
	this.BlockIdentifier = blockIdentifier
	this.ParentBlockIdentifier = parentBlockIdentifier
	this.Timestamp = timestamp
	this.Transactions = transactions
	this.Metadata = metadata
	return &this
}

// NewRosettaBlockWithDefaults instantiates a new RosettaBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaBlockWithDefaults() *RosettaBlock {
	this := RosettaBlock{}
	return &this
}

// GetBlockIdentifier returns the BlockIdentifier field value
func (o *RosettaBlock) GetBlockIdentifier() RosettaBlockIdentifier {
	if o == nil {
		var ret RosettaBlockIdentifier
		return ret
	}

	return o.BlockIdentifier
}

// GetBlockIdentifierOk returns a tuple with the BlockIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaBlock) GetBlockIdentifierOk() (*RosettaBlockIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockIdentifier, true
}

// SetBlockIdentifier sets field value
func (o *RosettaBlock) SetBlockIdentifier(v RosettaBlockIdentifier) {
	o.BlockIdentifier = v
}

// GetParentBlockIdentifier returns the ParentBlockIdentifier field value
func (o *RosettaBlock) GetParentBlockIdentifier() RosettaParentBlockIdentifier {
	if o == nil {
		var ret RosettaParentBlockIdentifier
		return ret
	}

	return o.ParentBlockIdentifier
}

// GetParentBlockIdentifierOk returns a tuple with the ParentBlockIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaBlock) GetParentBlockIdentifierOk() (*RosettaParentBlockIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBlockIdentifier, true
}

// SetParentBlockIdentifier sets field value
func (o *RosettaBlock) SetParentBlockIdentifier(v RosettaParentBlockIdentifier) {
	o.ParentBlockIdentifier = v
}

// GetTimestamp returns the Timestamp field value
func (o *RosettaBlock) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *RosettaBlock) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *RosettaBlock) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetTransactions returns the Transactions field value
func (o *RosettaBlock) GetTransactions() []RosettaTransaction {
	if o == nil {
		var ret []RosettaTransaction
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *RosettaBlock) GetTransactionsOk() ([]RosettaTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transactions, true
}

// SetTransactions sets field value
func (o *RosettaBlock) SetTransactions(v []RosettaTransaction) {
	o.Transactions = v
}

// GetMetadata returns the Metadata field value
func (o *RosettaBlock) GetMetadata() RosettaBlockMetadata {
	if o == nil {
		var ret RosettaBlockMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *RosettaBlock) GetMetadataOk() (*RosettaBlockMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *RosettaBlock) SetMetadata(v RosettaBlockMetadata) {
	o.Metadata = v
}

func (o RosettaBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["block_identifier"] = o.BlockIdentifier
	toSerialize["parent_block_identifier"] = o.ParentBlockIdentifier
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["transactions"] = o.Transactions
	toSerialize["metadata"] = o.Metadata
	return toSerialize, nil
}

func (o *RosettaBlock) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"block_identifier",
		"parent_block_identifier",
		"timestamp",
		"transactions",
		"metadata",
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

	varRosettaBlock := _RosettaBlock{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaBlock)

	if err != nil {
		return err
	}

	*o = RosettaBlock(varRosettaBlock)

	return err
}

type NullableRosettaBlock struct {
	value *RosettaBlock
	isSet bool
}

func (v NullableRosettaBlock) Get() *RosettaBlock {
	return v.value
}

func (v *NullableRosettaBlock) Set(val *RosettaBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaBlock(val *RosettaBlock) *NullableRosettaBlock {
	return &NullableRosettaBlock{value: val, isSet: true}
}

func (v NullableRosettaBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
