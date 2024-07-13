package models

import (
	"encoding/json"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaBlockResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaBlockResponse{}

// RosettaBlockResponse A BlockResponse includes a fully-populated block or a partially-populated block with a list of other transactions to fetch (other_transactions). As a result of the consensus algorithm of some blockchains, blocks can be omitted (i.e. certain block indexes can be skipped). If a query for one of these omitted indexes is made, the response should not include a Block object. It is VERY important to note that blocks MUST still form a canonical, connected chain of blocks where each block has a unique index. In other words, the PartialBlockIdentifier of a block after an omitted block should reference the last non-omitted block.
type RosettaBlockResponse struct {
	Block *RosettaBlock `json:"block,omitempty"`
	// Some blockchains may require additional transactions to be fetched that weren't returned in the block response (ex: block only returns transaction hashes). For blockchains with a lot of transactions in each block, this can be very useful as consumers can concurrently fetch all transactions returned.
	OtherTransactions []OtherTransactionIdentifier `json:"other_transactions,omitempty"`
}

// NewRosettaBlockResponse instantiates a new RosettaBlockResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaBlockResponse() *RosettaBlockResponse {
	this := RosettaBlockResponse{}
	return &this
}

// NewRosettaBlockResponseWithDefaults instantiates a new RosettaBlockResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaBlockResponseWithDefaults() *RosettaBlockResponse {
	this := RosettaBlockResponse{}
	return &this
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *RosettaBlockResponse) GetBlock() RosettaBlock {
	if o == nil || utils.IsNil(o.Block) {
		var ret RosettaBlock
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaBlockResponse) GetBlockOk() (*RosettaBlock, bool) {
	if o == nil || utils.IsNil(o.Block) {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *RosettaBlockResponse) HasBlock() bool {
	if o != nil && !utils.IsNil(o.Block) {
		return true
	}

	return false
}

// SetBlock gets a reference to the given RosettaBlock and assigns it to the Block field.
func (o *RosettaBlockResponse) SetBlock(v RosettaBlock) {
	o.Block = &v
}

// GetOtherTransactions returns the OtherTransactions field value if set, zero value otherwise.
func (o *RosettaBlockResponse) GetOtherTransactions() []OtherTransactionIdentifier {
	if o == nil || utils.IsNil(o.OtherTransactions) {
		var ret []OtherTransactionIdentifier
		return ret
	}
	return o.OtherTransactions
}

// GetOtherTransactionsOk returns a tuple with the OtherTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaBlockResponse) GetOtherTransactionsOk() ([]OtherTransactionIdentifier, bool) {
	if o == nil || utils.IsNil(o.OtherTransactions) {
		return nil, false
	}
	return o.OtherTransactions, true
}

// HasOtherTransactions returns a boolean if a field has been set.
func (o *RosettaBlockResponse) HasOtherTransactions() bool {
	if o != nil && !utils.IsNil(o.OtherTransactions) {
		return true
	}

	return false
}

// SetOtherTransactions gets a reference to the given []OtherTransactionIdentifier and assigns it to the OtherTransactions field.
func (o *RosettaBlockResponse) SetOtherTransactions(v []OtherTransactionIdentifier) {
	o.OtherTransactions = v
}

func (o RosettaBlockResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaBlockResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Block) {
		toSerialize["block"] = o.Block
	}
	if !utils.IsNil(o.OtherTransactions) {
		toSerialize["other_transactions"] = o.OtherTransactions
	}
	return toSerialize, nil
}

type NullableRosettaBlockResponse struct {
	value *RosettaBlockResponse
	isSet bool
}

func (v NullableRosettaBlockResponse) Get() *RosettaBlockResponse {
	return v.value
}

func (v *NullableRosettaBlockResponse) Set(val *RosettaBlockResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaBlockResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaBlockResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaBlockResponse(val *RosettaBlockResponse) *NullableRosettaBlockResponse {
	return &NullableRosettaBlockResponse{value: val, isSet: true}
}

func (v NullableRosettaBlockResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaBlockResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
