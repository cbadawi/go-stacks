package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the PoxStacker type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PoxStacker{}

// PoxStacker struct for PoxStacker
type PoxStacker struct {
	StackerAddress string `json:"stacker_address"`
	StackedAmount  string `json:"stacked_amount"`
	PoxAddress     string `json:"pox_address"`
	StackerType    string `json:"stacker_type"`
}

type _PoxStacker PoxStacker

// NewPoxStacker instantiates a new PoxStacker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoxStacker(stackerAddress string, stackedAmount string, poxAddress string, stackerType string) *PoxStacker {
	this := PoxStacker{}
	this.StackerAddress = stackerAddress
	this.StackedAmount = stackedAmount
	this.PoxAddress = poxAddress
	this.StackerType = stackerType
	return &this
}

// NewPoxStackerWithDefaults instantiates a new PoxStacker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoxStackerWithDefaults() *PoxStacker {
	this := PoxStacker{}
	return &this
}

// GetStackerAddress returns the StackerAddress field value
func (o *PoxStacker) GetStackerAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StackerAddress
}

// GetStackerAddressOk returns a tuple with the StackerAddress field value
// and a boolean to check if the value has been set.
func (o *PoxStacker) GetStackerAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StackerAddress, true
}

// SetStackerAddress sets field value
func (o *PoxStacker) SetStackerAddress(v string) {
	o.StackerAddress = v
}

// GetStackedAmount returns the StackedAmount field value
func (o *PoxStacker) GetStackedAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StackedAmount
}

// GetStackedAmountOk returns a tuple with the StackedAmount field value
// and a boolean to check if the value has been set.
func (o *PoxStacker) GetStackedAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StackedAmount, true
}

// SetStackedAmount sets field value
func (o *PoxStacker) SetStackedAmount(v string) {
	o.StackedAmount = v
}

// GetPoxAddress returns the PoxAddress field value
func (o *PoxStacker) GetPoxAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PoxAddress
}

// GetPoxAddressOk returns a tuple with the PoxAddress field value
// and a boolean to check if the value has been set.
func (o *PoxStacker) GetPoxAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoxAddress, true
}

// SetPoxAddress sets field value
func (o *PoxStacker) SetPoxAddress(v string) {
	o.PoxAddress = v
}

// GetStackerType returns the StackerType field value
func (o *PoxStacker) GetStackerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StackerType
}

// GetStackerTypeOk returns a tuple with the StackerType field value
// and a boolean to check if the value has been set.
func (o *PoxStacker) GetStackerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StackerType, true
}

// SetStackerType sets field value
func (o *PoxStacker) SetStackerType(v string) {
	o.StackerType = v
}

func (o PoxStacker) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoxStacker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stacker_address"] = o.StackerAddress
	toSerialize["stacked_amount"] = o.StackedAmount
	toSerialize["pox_address"] = o.PoxAddress
	toSerialize["stacker_type"] = o.StackerType
	return toSerialize, nil
}

func (o *PoxStacker) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"stacker_address",
		"stacked_amount",
		"pox_address",
		"stacker_type",
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

	varPoxStacker := _PoxStacker{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPoxStacker)

	if err != nil {
		return err
	}

	*o = PoxStacker(varPoxStacker)

	return err
}

type NullablePoxStacker struct {
	value *PoxStacker
	isSet bool
}

func (v NullablePoxStacker) Get() *PoxStacker {
	return v.value
}

func (v *NullablePoxStacker) Set(val *PoxStacker) {
	v.value = val
	v.isSet = true
}

func (v NullablePoxStacker) IsSet() bool {
	return v.isSet
}

func (v *NullablePoxStacker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoxStacker(val *PoxStacker) *NullablePoxStacker {
	return &NullablePoxStacker{value: val, isSet: true}
}

func (v NullablePoxStacker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoxStacker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
