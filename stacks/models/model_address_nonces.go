package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AddressNonces type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressNonces{}

// AddressNonces The latest nonce values used by an account by inspecting the mempool, microblock transactions, and anchored transactions
type AddressNonces struct {
	// The latest nonce found within mempool transactions sent by this address. Will be null if there are no current mempool transactions for this address.
	LastMempoolTxNonce utils.NullableInt32 `json:"last_mempool_tx_nonce"`
	// The latest nonce found within transactions sent by this address, including unanchored microblock transactions. Will be null if there are no current transactions for this address.
	LastExecutedTxNonce utils.NullableInt32 `json:"last_executed_tx_nonce"`
	// The likely nonce required for creating the next transaction, based on the last nonces seen by the API. This can be incorrect if the API's mempool or transactions aren't fully synchronized, even by a small amount, or if a previous transaction is still propagating through the Stacks blockchain network when this endpoint is called.
	PossibleNextNonce int32 `json:"possible_next_nonce"`
	// Nonces that appear to be missing and likely causing a mempool transaction to be stuck.
	DetectedMissingNonces []int32 `json:"detected_missing_nonces"`
	// Nonces currently in mempool for this address.
	DetectedMempoolNonces []int32 `json:"detected_mempool_nonces,omitempty"`
}

type _AddressNonces AddressNonces

// NewAddressNonces instantiates a new AddressNonces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressNonces(lastMempoolTxNonce utils.NullableInt32, lastExecutedTxNonce utils.NullableInt32, possibleNextNonce int32, detectedMissingNonces []int32) *AddressNonces {
	this := AddressNonces{}
	this.LastMempoolTxNonce = lastMempoolTxNonce
	this.LastExecutedTxNonce = lastExecutedTxNonce
	this.PossibleNextNonce = possibleNextNonce
	this.DetectedMissingNonces = detectedMissingNonces
	return &this
}

// NewAddressNoncesWithDefaults instantiates a new AddressNonces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressNoncesWithDefaults() *AddressNonces {
	this := AddressNonces{}
	return &this
}

// GetLastMempoolTxNonce returns the LastMempoolTxNonce field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *AddressNonces) GetLastMempoolTxNonce() int32 {
	if o == nil || o.LastMempoolTxNonce.Get() == nil {
		var ret int32
		return ret
	}

	return *o.LastMempoolTxNonce.Get()
}

// GetLastMempoolTxNonceOk returns a tuple with the LastMempoolTxNonce field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressNonces) GetLastMempoolTxNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastMempoolTxNonce.Get(), o.LastMempoolTxNonce.IsSet()
}

// SetLastMempoolTxNonce sets field value
func (o *AddressNonces) SetLastMempoolTxNonce(v int32) {
	o.LastMempoolTxNonce.Set(&v)
}

// GetLastExecutedTxNonce returns the LastExecutedTxNonce field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *AddressNonces) GetLastExecutedTxNonce() int32 {
	if o == nil || o.LastExecutedTxNonce.Get() == nil {
		var ret int32
		return ret
	}

	return *o.LastExecutedTxNonce.Get()
}

// GetLastExecutedTxNonceOk returns a tuple with the LastExecutedTxNonce field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressNonces) GetLastExecutedTxNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastExecutedTxNonce.Get(), o.LastExecutedTxNonce.IsSet()
}

// SetLastExecutedTxNonce sets field value
func (o *AddressNonces) SetLastExecutedTxNonce(v int32) {
	o.LastExecutedTxNonce.Set(&v)
}

// GetPossibleNextNonce returns the PossibleNextNonce field value
func (o *AddressNonces) GetPossibleNextNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PossibleNextNonce
}

// GetPossibleNextNonceOk returns a tuple with the PossibleNextNonce field value
// and a boolean to check if the value has been set.
func (o *AddressNonces) GetPossibleNextNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PossibleNextNonce, true
}

// SetPossibleNextNonce sets field value
func (o *AddressNonces) SetPossibleNextNonce(v int32) {
	o.PossibleNextNonce = v
}

// GetDetectedMissingNonces returns the DetectedMissingNonces field value
func (o *AddressNonces) GetDetectedMissingNonces() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.DetectedMissingNonces
}

// GetDetectedMissingNoncesOk returns a tuple with the DetectedMissingNonces field value
// and a boolean to check if the value has been set.
func (o *AddressNonces) GetDetectedMissingNoncesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DetectedMissingNonces, true
}

// SetDetectedMissingNonces sets field value
func (o *AddressNonces) SetDetectedMissingNonces(v []int32) {
	o.DetectedMissingNonces = v
}

// GetDetectedMempoolNonces returns the DetectedMempoolNonces field value if set, zero value otherwise.
func (o *AddressNonces) GetDetectedMempoolNonces() []int32 {
	if o == nil || utils.IsNil(o.DetectedMempoolNonces) {
		var ret []int32
		return ret
	}
	return o.DetectedMempoolNonces
}

// GetDetectedMempoolNoncesOk returns a tuple with the DetectedMempoolNonces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressNonces) GetDetectedMempoolNoncesOk() ([]int32, bool) {
	if o == nil || utils.IsNil(o.DetectedMempoolNonces) {
		return nil, false
	}
	return o.DetectedMempoolNonces, true
}

// HasDetectedMempoolNonces returns a boolean if a field has been set.
func (o *AddressNonces) HasDetectedMempoolNonces() bool {
	if o != nil && !utils.IsNil(o.DetectedMempoolNonces) {
		return true
	}

	return false
}

// SetDetectedMempoolNonces gets a reference to the given []int32 and assigns it to the DetectedMempoolNonces field.
func (o *AddressNonces) SetDetectedMempoolNonces(v []int32) {
	o.DetectedMempoolNonces = v
}

func (o AddressNonces) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressNonces) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["last_mempool_tx_nonce"] = o.LastMempoolTxNonce.Get()
	toSerialize["last_executed_tx_nonce"] = o.LastExecutedTxNonce.Get()
	toSerialize["possible_next_nonce"] = o.PossibleNextNonce
	toSerialize["detected_missing_nonces"] = o.DetectedMissingNonces
	if !utils.IsNil(o.DetectedMempoolNonces) {
		toSerialize["detected_mempool_nonces"] = o.DetectedMempoolNonces
	}
	return toSerialize, nil
}

func (o *AddressNonces) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"last_mempool_tx_nonce",
		"last_executed_tx_nonce",
		"possible_next_nonce",
		"detected_missing_nonces",
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

	varAddressNonces := _AddressNonces{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressNonces)

	if err != nil {
		return err
	}

	*o = AddressNonces(varAddressNonces)

	return err
}

type NullableAddressNonces struct {
	value *AddressNonces
	isSet bool
}

func (v NullableAddressNonces) Get() *AddressNonces {
	return v.value
}

func (v *NullableAddressNonces) Set(val *AddressNonces) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressNonces) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressNonces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressNonces(val *AddressNonces) *NullableAddressNonces {
	return &NullableAddressNonces{value: val, isSet: true}
}

func (v NullableAddressNonces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressNonces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
