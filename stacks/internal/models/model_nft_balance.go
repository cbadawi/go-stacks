package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the NftBalance type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NftBalance{}

// NftBalance struct for NftBalance
type NftBalance struct {
	Count         string `json:"count"`
	TotalSent     string `json:"total_sent"`
	TotalReceived string `json:"total_received"`
}

type _NftBalance NftBalance

// NewNftBalance instantiates a new NftBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNftBalance(count string, totalSent string, totalReceived string) *NftBalance {
	this := NftBalance{}
	this.Count = count
	this.TotalSent = totalSent
	this.TotalReceived = totalReceived
	return &this
}

// NewNftBalanceWithDefaults instantiates a new NftBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNftBalanceWithDefaults() *NftBalance {
	this := NftBalance{}
	return &this
}

// GetCount returns the Count field value
func (o *NftBalance) GetCount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *NftBalance) GetCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *NftBalance) SetCount(v string) {
	o.Count = v
}

// GetTotalSent returns the TotalSent field value
func (o *NftBalance) GetTotalSent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalSent
}

// GetTotalSentOk returns a tuple with the TotalSent field value
// and a boolean to check if the value has been set.
func (o *NftBalance) GetTotalSentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSent, true
}

// SetTotalSent sets field value
func (o *NftBalance) SetTotalSent(v string) {
	o.TotalSent = v
}

// GetTotalReceived returns the TotalReceived field value
func (o *NftBalance) GetTotalReceived() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalReceived
}

// GetTotalReceivedOk returns a tuple with the TotalReceived field value
// and a boolean to check if the value has been set.
func (o *NftBalance) GetTotalReceivedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalReceived, true
}

// SetTotalReceived sets field value
func (o *NftBalance) SetTotalReceived(v string) {
	o.TotalReceived = v
}

func (o NftBalance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NftBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["total_sent"] = o.TotalSent
	toSerialize["total_received"] = o.TotalReceived
	return toSerialize, nil
}

func (o *NftBalance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"total_sent",
		"total_received",
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

	varNftBalance := _NftBalance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNftBalance)

	if err != nil {
		return err
	}

	*o = NftBalance(varNftBalance)

	return err
}

type NullableNftBalance struct {
	value *NftBalance
	isSet bool
}

func (v NullableNftBalance) Get() *NftBalance {
	return v.value
}

func (v *NullableNftBalance) Set(val *NftBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableNftBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableNftBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNftBalance(val *NftBalance) *NullableNftBalance {
	return &NullableNftBalance{value: val, isSet: true}
}

func (v NullableNftBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNftBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
