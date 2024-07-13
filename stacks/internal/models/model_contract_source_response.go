package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the ContractSourceResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ContractSourceResponse{}

// ContractSourceResponse GET request to get contract source
type ContractSourceResponse struct {
	Source        string `json:"source"`
	PublishHeight int32  `json:"publish_height"`
	Proof         string `json:"proof"`
}

type _ContractSourceResponse ContractSourceResponse

// NewContractSourceResponse instantiates a new ContractSourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractSourceResponse(source string, publishHeight int32, proof string) *ContractSourceResponse {
	this := ContractSourceResponse{}
	this.Source = source
	this.PublishHeight = publishHeight
	this.Proof = proof
	return &this
}

// NewContractSourceResponseWithDefaults instantiates a new ContractSourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractSourceResponseWithDefaults() *ContractSourceResponse {
	this := ContractSourceResponse{}
	return &this
}

// GetSource returns the Source field value
func (o *ContractSourceResponse) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ContractSourceResponse) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ContractSourceResponse) SetSource(v string) {
	o.Source = v
}

// GetPublishHeight returns the PublishHeight field value
func (o *ContractSourceResponse) GetPublishHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PublishHeight
}

// GetPublishHeightOk returns a tuple with the PublishHeight field value
// and a boolean to check if the value has been set.
func (o *ContractSourceResponse) GetPublishHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishHeight, true
}

// SetPublishHeight sets field value
func (o *ContractSourceResponse) SetPublishHeight(v int32) {
	o.PublishHeight = v
}

// GetProof returns the Proof field value
func (o *ContractSourceResponse) GetProof() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Proof
}

// GetProofOk returns a tuple with the Proof field value
// and a boolean to check if the value has been set.
func (o *ContractSourceResponse) GetProofOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proof, true
}

// SetProof sets field value
func (o *ContractSourceResponse) SetProof(v string) {
	o.Proof = v
}

func (o ContractSourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractSourceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["publish_height"] = o.PublishHeight
	toSerialize["proof"] = o.Proof
	return toSerialize, nil
}

func (o *ContractSourceResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
		"publish_height",
		"proof",
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

	varContractSourceResponse := _ContractSourceResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContractSourceResponse)

	if err != nil {
		return err
	}

	*o = ContractSourceResponse(varContractSourceResponse)

	return err
}

type NullableContractSourceResponse struct {
	value *ContractSourceResponse
	isSet bool
}

func (v NullableContractSourceResponse) Get() *ContractSourceResponse {
	return v.value
}

func (v *NullableContractSourceResponse) Set(val *ContractSourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContractSourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContractSourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractSourceResponse(val *ContractSourceResponse) *NullableContractSourceResponse {
	return &NullableContractSourceResponse{value: val, isSet: true}
}

func (v NullableContractSourceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractSourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
