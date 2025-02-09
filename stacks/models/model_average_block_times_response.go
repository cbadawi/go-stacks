package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the AverageBlockTimesResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AverageBlockTimesResponse{}

// AverageBlockTimesResponse Request to fetch average block times (in seconds)
type AverageBlockTimesResponse struct {
	// Average block times over the last hour (in seconds)
	Last1h float32 `json:"last_1h"`
	// Average block times over the last 24 hours (in seconds)
	Last24h float32 `json:"last_24h"`
	// Average block times over the last 7 days (in seconds)
	Last7d float32 `json:"last_7d"`
	// Average block times over the last 30 days (in seconds)
	Last30d float32 `json:"last_30d"`
}

type _AverageBlockTimesResponse AverageBlockTimesResponse

// NewAverageBlockTimesResponse instantiates a new AverageBlockTimesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAverageBlockTimesResponse(last1h float32, last24h float32, last7d float32, last30d float32) *AverageBlockTimesResponse {
	this := AverageBlockTimesResponse{}
	this.Last1h = last1h
	this.Last24h = last24h
	this.Last7d = last7d
	this.Last30d = last30d
	return &this
}

// NewAverageBlockTimesResponseWithDefaults instantiates a new AverageBlockTimesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAverageBlockTimesResponseWithDefaults() *AverageBlockTimesResponse {
	this := AverageBlockTimesResponse{}
	return &this
}

// GetLast1h returns the Last1h field value
func (o *AverageBlockTimesResponse) GetLast1h() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Last1h
}

// GetLast1hOk returns a tuple with the Last1h field value
// and a boolean to check if the value has been set.
func (o *AverageBlockTimesResponse) GetLast1hOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Last1h, true
}

// SetLast1h sets field value
func (o *AverageBlockTimesResponse) SetLast1h(v float32) {
	o.Last1h = v
}

// GetLast24h returns the Last24h field value
func (o *AverageBlockTimesResponse) GetLast24h() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Last24h
}

// GetLast24hOk returns a tuple with the Last24h field value
// and a boolean to check if the value has been set.
func (o *AverageBlockTimesResponse) GetLast24hOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Last24h, true
}

// SetLast24h sets field value
func (o *AverageBlockTimesResponse) SetLast24h(v float32) {
	o.Last24h = v
}

// GetLast7d returns the Last7d field value
func (o *AverageBlockTimesResponse) GetLast7d() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Last7d
}

// GetLast7dOk returns a tuple with the Last7d field value
// and a boolean to check if the value has been set.
func (o *AverageBlockTimesResponse) GetLast7dOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Last7d, true
}

// SetLast7d sets field value
func (o *AverageBlockTimesResponse) SetLast7d(v float32) {
	o.Last7d = v
}

// GetLast30d returns the Last30d field value
func (o *AverageBlockTimesResponse) GetLast30d() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Last30d
}

// GetLast30dOk returns a tuple with the Last30d field value
// and a boolean to check if the value has been set.
func (o *AverageBlockTimesResponse) GetLast30dOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Last30d, true
}

// SetLast30d sets field value
func (o *AverageBlockTimesResponse) SetLast30d(v float32) {
	o.Last30d = v
}

func (o AverageBlockTimesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AverageBlockTimesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["last_1h"] = o.Last1h
	toSerialize["last_24h"] = o.Last24h
	toSerialize["last_7d"] = o.Last7d
	toSerialize["last_30d"] = o.Last30d
	return toSerialize, nil
}

func (o *AverageBlockTimesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"last_1h",
		"last_24h",
		"last_7d",
		"last_30d",
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

	varAverageBlockTimesResponse := _AverageBlockTimesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAverageBlockTimesResponse)

	if err != nil {
		return err
	}

	*o = AverageBlockTimesResponse(varAverageBlockTimesResponse)

	return err
}

type NullableAverageBlockTimesResponse struct {
	value *AverageBlockTimesResponse
	isSet bool
}

func (v NullableAverageBlockTimesResponse) Get() *AverageBlockTimesResponse {
	return v.value
}

func (v *NullableAverageBlockTimesResponse) Set(val *AverageBlockTimesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAverageBlockTimesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAverageBlockTimesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAverageBlockTimesResponse(val *AverageBlockTimesResponse) *NullableAverageBlockTimesResponse {
	return &NullableAverageBlockTimesResponse{value: val, isSet: true}
}

func (v NullableAverageBlockTimesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAverageBlockTimesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
