package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaAccountBalanceResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaAccountBalanceResponse{}

// RosettaAccountBalanceResponse An AccountBalanceResponse is returned on the /account/balance endpoint. If an account has a balance for each AccountIdentifier describing it (ex: an ERC-20 token balance on a few smart contracts), an account balance request must be made with each AccountIdentifier.
type RosettaAccountBalanceResponse struct {
	BlockIdentifier RosettaBlockIdentifier `json:"block_identifier"`
	// A single account balance may have multiple currencies
	Balances []RosettaAmount `json:"balances"`
	// If a blockchain is UTXO-based, all unspent Coins owned by an account_identifier should be returned alongside the balance. It is highly recommended to populate this field so that users of the Rosetta API implementation don't need to maintain their own indexer to track their UTXOs.
	Coins    []RosettaCoin                          `json:"coins,omitempty"`
	Metadata *RosettaAccountBalanceResponseMetadata `json:"metadata,omitempty"`
}

type _RosettaAccountBalanceResponse RosettaAccountBalanceResponse

// NewRosettaAccountBalanceResponse instantiates a new RosettaAccountBalanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaAccountBalanceResponse(blockIdentifier RosettaBlockIdentifier, balances []RosettaAmount) *RosettaAccountBalanceResponse {
	this := RosettaAccountBalanceResponse{}
	this.BlockIdentifier = blockIdentifier
	this.Balances = balances
	return &this
}

// NewRosettaAccountBalanceResponseWithDefaults instantiates a new RosettaAccountBalanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaAccountBalanceResponseWithDefaults() *RosettaAccountBalanceResponse {
	this := RosettaAccountBalanceResponse{}
	return &this
}

// GetBlockIdentifier returns the BlockIdentifier field value
func (o *RosettaAccountBalanceResponse) GetBlockIdentifier() RosettaBlockIdentifier {
	if o == nil {
		var ret RosettaBlockIdentifier
		return ret
	}

	return o.BlockIdentifier
}

// GetBlockIdentifierOk returns a tuple with the BlockIdentifier field value
// and a boolean to check if the value has been set.
func (o *RosettaAccountBalanceResponse) GetBlockIdentifierOk() (*RosettaBlockIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockIdentifier, true
}

// SetBlockIdentifier sets field value
func (o *RosettaAccountBalanceResponse) SetBlockIdentifier(v RosettaBlockIdentifier) {
	o.BlockIdentifier = v
}

// GetBalances returns the Balances field value
func (o *RosettaAccountBalanceResponse) GetBalances() []RosettaAmount {
	if o == nil {
		var ret []RosettaAmount
		return ret
	}

	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value
// and a boolean to check if the value has been set.
func (o *RosettaAccountBalanceResponse) GetBalancesOk() ([]RosettaAmount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Balances, true
}

// SetBalances sets field value
func (o *RosettaAccountBalanceResponse) SetBalances(v []RosettaAmount) {
	o.Balances = v
}

// GetCoins returns the Coins field value if set, zero value otherwise.
func (o *RosettaAccountBalanceResponse) GetCoins() []RosettaCoin {
	if o == nil || utils.IsNil(o.Coins) {
		var ret []RosettaCoin
		return ret
	}
	return o.Coins
}

// GetCoinsOk returns a tuple with the Coins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaAccountBalanceResponse) GetCoinsOk() ([]RosettaCoin, bool) {
	if o == nil || utils.IsNil(o.Coins) {
		return nil, false
	}
	return o.Coins, true
}

// HasCoins returns a boolean if a field has been set.
func (o *RosettaAccountBalanceResponse) HasCoins() bool {
	if o != nil && !utils.IsNil(o.Coins) {
		return true
	}

	return false
}

// SetCoins gets a reference to the given []RosettaCoin and assigns it to the Coins field.
func (o *RosettaAccountBalanceResponse) SetCoins(v []RosettaCoin) {
	o.Coins = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RosettaAccountBalanceResponse) GetMetadata() RosettaAccountBalanceResponseMetadata {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret RosettaAccountBalanceResponseMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RosettaAccountBalanceResponse) GetMetadataOk() (*RosettaAccountBalanceResponseMetadata, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RosettaAccountBalanceResponse) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given RosettaAccountBalanceResponseMetadata and assigns it to the Metadata field.
func (o *RosettaAccountBalanceResponse) SetMetadata(v RosettaAccountBalanceResponseMetadata) {
	o.Metadata = &v
}

func (o RosettaAccountBalanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaAccountBalanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["block_identifier"] = o.BlockIdentifier
	toSerialize["balances"] = o.Balances
	if !utils.IsNil(o.Coins) {
		toSerialize["coins"] = o.Coins
	}
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *RosettaAccountBalanceResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"block_identifier",
		"balances",
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

	varRosettaAccountBalanceResponse := _RosettaAccountBalanceResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaAccountBalanceResponse)

	if err != nil {
		return err
	}

	*o = RosettaAccountBalanceResponse(varRosettaAccountBalanceResponse)

	return err
}

type NullableRosettaAccountBalanceResponse struct {
	value *RosettaAccountBalanceResponse
	isSet bool
}

func (v NullableRosettaAccountBalanceResponse) Get() *RosettaAccountBalanceResponse {
	return v.value
}

func (v *NullableRosettaAccountBalanceResponse) Set(val *RosettaAccountBalanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaAccountBalanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaAccountBalanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaAccountBalanceResponse(val *RosettaAccountBalanceResponse) *NullableRosettaAccountBalanceResponse {
	return &NullableRosettaAccountBalanceResponse{value: val, isSet: true}
}

func (v NullableRosettaAccountBalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaAccountBalanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
