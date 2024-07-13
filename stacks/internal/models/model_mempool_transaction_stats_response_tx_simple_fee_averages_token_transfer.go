package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer{}

// MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer struct for MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer
type MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer struct {
	P25 utils.NullableFloat32 `json:"p25"`
	P50 utils.NullableFloat32 `json:"p50"`
	P75 utils.NullableFloat32 `json:"p75"`
	P95 utils.NullableFloat32 `json:"p95"`
}

type _MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer

// NewMempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer instantiates a new MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer(p25 utils.NullableFloat32, p50 utils.NullableFloat32, p75 utils.NullableFloat32, p95 utils.NullableFloat32) *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer {
	this := MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer{}
	this.P25 = p25
	this.P50 = p50
	this.P75 = p75
	this.P95 = p95
	return &this
}

// NewMempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransferWithDefaults instantiates a new MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransferWithDefaults() *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer {
	this := MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer{}
	return &this
}

// GetP25 returns the P25 field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) GetP25() float32 {
	if o == nil || o.P25.Get() == nil {
		var ret float32
		return ret
	}

	return *o.P25.Get()
}

// GetP25Ok returns a tuple with the P25 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) GetP25Ok() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.P25.Get(), o.P25.IsSet()
}

// SetP25 sets field value
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) SetP25(v float32) {
	o.P25.Set(&v)
}

// GetP50 returns the P50 field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) GetP50() float32 {
	if o == nil || o.P50.Get() == nil {
		var ret float32
		return ret
	}

	return *o.P50.Get()
}

// GetP50Ok returns a tuple with the P50 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) GetP50Ok() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.P50.Get(), o.P50.IsSet()
}

// SetP50 sets field value
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) SetP50(v float32) {
	o.P50.Set(&v)
}

// GetP75 returns the P75 field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) GetP75() float32 {
	if o == nil || o.P75.Get() == nil {
		var ret float32
		return ret
	}

	return *o.P75.Get()
}

// GetP75Ok returns a tuple with the P75 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) GetP75Ok() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.P75.Get(), o.P75.IsSet()
}

// SetP75 sets field value
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) SetP75(v float32) {
	o.P75.Set(&v)
}

// GetP95 returns the P95 field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) GetP95() float32 {
	if o == nil || o.P95.Get() == nil {
		var ret float32
		return ret
	}

	return *o.P95.Get()
}

// GetP95Ok returns a tuple with the P95 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) GetP95Ok() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.P95.Get(), o.P95.IsSet()
}

// SetP95 sets field value
func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) SetP95(v float32) {
	o.P95.Set(&v)
}

func (o MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["p25"] = o.P25.Get()
	toSerialize["p50"] = o.P50.Get()
	toSerialize["p75"] = o.P75.Get()
	toSerialize["p95"] = o.P95.Get()
	return toSerialize, nil
}

func (o *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"p25",
		"p50",
		"p75",
		"p95",
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

	varMempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer := _MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer)

	if err != nil {
		return err
	}

	*o = MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer(varMempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer)

	return err
}

type NullableMempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer struct {
	value *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer
	isSet bool
}

func (v NullableMempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) Get() *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer {
	return v.value
}

func (v *NullableMempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) Set(val *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableMempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableMempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer(val *MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) *NullableMempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer {
	return &NullableMempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer{value: val, isSet: true}
}

func (v NullableMempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
