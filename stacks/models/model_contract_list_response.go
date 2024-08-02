package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the ContractListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ContractListResponse{}

// ContractListResponse GET list of contracts
type ContractListResponse struct {
	// The number of contracts to return
	Limit int32 `json:"limit"`
	// The number to contracts to skip (starting at `0`)
	Offset  int32           `json:"offset"`
	Results []SmartContract `json:"results"`
}

type _ContractListResponse ContractListResponse

// NewContractListResponse instantiates a new ContractListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractListResponse(limit int32, offset int32, results []SmartContract) *ContractListResponse {
	this := ContractListResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Results = results
	return &this
}

// NewContractListResponseWithDefaults instantiates a new ContractListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractListResponseWithDefaults() *ContractListResponse {
	this := ContractListResponse{}
	var offset int32 = 0
	this.Offset = offset
	return &this
}

// GetLimit returns the Limit field value
func (o *ContractListResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ContractListResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ContractListResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *ContractListResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *ContractListResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *ContractListResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetResults returns the Results field value
func (o *ContractListResponse) GetResults() []SmartContract {
	if o == nil {
		var ret []SmartContract
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ContractListResponse) GetResultsOk() ([]SmartContract, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ContractListResponse) SetResults(v []SmartContract) {
	o.Results = v
}

func (o ContractListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *ContractListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"limit",
		"offset",
		"results",
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

	varContractListResponse := _ContractListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContractListResponse)

	if err != nil {
		return err
	}

	*o = ContractListResponse(varContractListResponse)

	return err
}

type NullableContractListResponse struct {
	value *ContractListResponse
	isSet bool
}

func (v NullableContractListResponse) Get() *ContractListResponse {
	return v.value
}

func (v *NullableContractListResponse) Set(val *ContractListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContractListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContractListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractListResponse(val *ContractListResponse) *NullableContractListResponse {
	return &NullableContractListResponse{value: val, isSet: true}
}

func (v NullableContractListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
