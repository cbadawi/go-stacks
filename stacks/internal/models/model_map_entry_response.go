package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the MapEntryResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MapEntryResponse{}

// MapEntryResponse Response of get data map entry request
type MapEntryResponse struct {
	// Hex-encoded string of clarity value. It is always an optional tuple.
	Data string `json:"data"`
	// Hex-encoded string of the MARF proof for the data
	Proof *string `json:"proof,omitempty"`
}

type _MapEntryResponse MapEntryResponse

// NewMapEntryResponse instantiates a new MapEntryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMapEntryResponse(data string) *MapEntryResponse {
	this := MapEntryResponse{}
	this.Data = data
	return &this
}

// NewMapEntryResponseWithDefaults instantiates a new MapEntryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMapEntryResponseWithDefaults() *MapEntryResponse {
	this := MapEntryResponse{}
	return &this
}

// GetData returns the Data field value
func (o *MapEntryResponse) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MapEntryResponse) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MapEntryResponse) SetData(v string) {
	o.Data = v
}

// GetProof returns the Proof field value if set, zero value otherwise.
func (o *MapEntryResponse) GetProof() string {
	if o == nil || utils.IsNil(o.Proof) {
		var ret string
		return ret
	}
	return *o.Proof
}

// GetProofOk returns a tuple with the Proof field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapEntryResponse) GetProofOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Proof) {
		return nil, false
	}
	return o.Proof, true
}

// HasProof returns a boolean if a field has been set.
func (o *MapEntryResponse) HasProof() bool {
	if o != nil && !utils.IsNil(o.Proof) {
		return true
	}

	return false
}

// SetProof gets a reference to the given string and assigns it to the Proof field.
func (o *MapEntryResponse) SetProof(v string) {
	o.Proof = &v
}

func (o MapEntryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MapEntryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !utils.IsNil(o.Proof) {
		toSerialize["proof"] = o.Proof
	}
	return toSerialize, nil
}

func (o *MapEntryResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varMapEntryResponse := _MapEntryResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMapEntryResponse)

	if err != nil {
		return err
	}

	*o = MapEntryResponse(varMapEntryResponse)

	return err
}

type NullableMapEntryResponse struct {
	value *MapEntryResponse
	isSet bool
}

func (v NullableMapEntryResponse) Get() *MapEntryResponse {
	return v.value
}

func (v *NullableMapEntryResponse) Set(val *MapEntryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMapEntryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMapEntryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapEntryResponse(val *MapEntryResponse) *NullableMapEntryResponse {
	return &NullableMapEntryResponse{value: val, isSet: true}
}

func (v NullableMapEntryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapEntryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
