package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BnsGetAllNamespacesResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BnsGetAllNamespacesResponse{}

// BnsGetAllNamespacesResponse Fetch a list of all namespaces known to the node.
type BnsGetAllNamespacesResponse struct {
	Namespaces []string `json:"namespaces"`
}

type _BnsGetAllNamespacesResponse BnsGetAllNamespacesResponse

// NewBnsGetAllNamespacesResponse instantiates a new BnsGetAllNamespacesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBnsGetAllNamespacesResponse(namespaces []string) *BnsGetAllNamespacesResponse {
	this := BnsGetAllNamespacesResponse{}
	this.Namespaces = namespaces
	return &this
}

// NewBnsGetAllNamespacesResponseWithDefaults instantiates a new BnsGetAllNamespacesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBnsGetAllNamespacesResponseWithDefaults() *BnsGetAllNamespacesResponse {
	this := BnsGetAllNamespacesResponse{}
	return &this
}

// GetNamespaces returns the Namespaces field value
func (o *BnsGetAllNamespacesResponse) GetNamespaces() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value
// and a boolean to check if the value has been set.
func (o *BnsGetAllNamespacesResponse) GetNamespacesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespaces, true
}

// SetNamespaces sets field value
func (o *BnsGetAllNamespacesResponse) SetNamespaces(v []string) {
	o.Namespaces = v
}

func (o BnsGetAllNamespacesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BnsGetAllNamespacesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["namespaces"] = o.Namespaces
	return toSerialize, nil
}

func (o *BnsGetAllNamespacesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"namespaces",
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

	varBnsGetAllNamespacesResponse := _BnsGetAllNamespacesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBnsGetAllNamespacesResponse)

	if err != nil {
		return err
	}

	*o = BnsGetAllNamespacesResponse(varBnsGetAllNamespacesResponse)

	return err
}

type NullableBnsGetAllNamespacesResponse struct {
	value *BnsGetAllNamespacesResponse
	isSet bool
}

func (v NullableBnsGetAllNamespacesResponse) Get() *BnsGetAllNamespacesResponse {
	return v.value
}

func (v *NullableBnsGetAllNamespacesResponse) Set(val *BnsGetAllNamespacesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBnsGetAllNamespacesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBnsGetAllNamespacesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBnsGetAllNamespacesResponse(val *BnsGetAllNamespacesResponse) *NullableBnsGetAllNamespacesResponse {
	return &NullableBnsGetAllNamespacesResponse{value: val, isSet: true}
}

func (v NullableBnsGetAllNamespacesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBnsGetAllNamespacesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
