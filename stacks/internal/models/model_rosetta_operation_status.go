package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaOperationStatus type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaOperationStatus{}

// RosettaOperationStatus OperationStatus is utilized to indicate which Operation status are considered successful.
type RosettaOperationStatus struct {
	// The status is the network-specific status of the operation.
	Status string `json:"status"`
	// An Operation is considered successful if the Operation.Amount should affect the Operation.Account. Some blockchains (like Bitcoin) only include successful operations in blocks but other blockchains (like Ethereum) include unsuccessful operations that incur a fee. To reconcile the computed balance from the stream of Operations, it is critical to understand which Operation.Status indicate an Operation is successful and should affect an Account.
	Successful bool `json:"successful"`
}

type _RosettaOperationStatus RosettaOperationStatus

// NewRosettaOperationStatus instantiates a new RosettaOperationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaOperationStatus(status string, successful bool) *RosettaOperationStatus {
	this := RosettaOperationStatus{}
	this.Status = status
	this.Successful = successful
	return &this
}

// NewRosettaOperationStatusWithDefaults instantiates a new RosettaOperationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaOperationStatusWithDefaults() *RosettaOperationStatus {
	this := RosettaOperationStatus{}
	return &this
}

// GetStatus returns the Status field value
func (o *RosettaOperationStatus) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RosettaOperationStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RosettaOperationStatus) SetStatus(v string) {
	o.Status = v
}

// GetSuccessful returns the Successful field value
func (o *RosettaOperationStatus) GetSuccessful() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Successful
}

// GetSuccessfulOk returns a tuple with the Successful field value
// and a boolean to check if the value has been set.
func (o *RosettaOperationStatus) GetSuccessfulOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Successful, true
}

// SetSuccessful sets field value
func (o *RosettaOperationStatus) SetSuccessful(v bool) {
	o.Successful = v
}

func (o RosettaOperationStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaOperationStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["successful"] = o.Successful
	return toSerialize, nil
}

func (o *RosettaOperationStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"successful",
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

	varRosettaOperationStatus := _RosettaOperationStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaOperationStatus)

	if err != nil {
		return err
	}

	*o = RosettaOperationStatus(varRosettaOperationStatus)

	return err
}

type NullableRosettaOperationStatus struct {
	value *RosettaOperationStatus
	isSet bool
}

func (v NullableRosettaOperationStatus) Get() *RosettaOperationStatus {
	return v.value
}

func (v *NullableRosettaOperationStatus) Set(val *RosettaOperationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaOperationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaOperationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaOperationStatus(val *RosettaOperationStatus) *NullableRosettaOperationStatus {
	return &NullableRosettaOperationStatus{value: val, isSet: true}
}

func (v NullableRosettaOperationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaOperationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
