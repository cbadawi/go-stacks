package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the GetStxSupplyResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GetStxSupplyResponse{}

// GetStxSupplyResponse GET request that returns network target block times
type GetStxSupplyResponse struct {
	// String quoted decimal number of the percentage of STX that have unlocked
	UnlockedPercent string `json:"unlocked_percent"`
	// String quoted decimal number of the total possible number of STX
	TotalStx string `json:"total_stx"`
	// String quoted decimal number of the STX that have been mined or unlocked
	UnlockedStx string `json:"unlocked_stx"`
	// The block height at which this information was queried
	BlockHeight int32 `json:"block_height"`
}

type _GetStxSupplyResponse GetStxSupplyResponse

// NewGetStxSupplyResponse instantiates a new GetStxSupplyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStxSupplyResponse(unlockedPercent string, totalStx string, unlockedStx string, blockHeight int32) *GetStxSupplyResponse {
	this := GetStxSupplyResponse{}
	this.UnlockedPercent = unlockedPercent
	this.TotalStx = totalStx
	this.UnlockedStx = unlockedStx
	this.BlockHeight = blockHeight
	return &this
}

// NewGetStxSupplyResponseWithDefaults instantiates a new GetStxSupplyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStxSupplyResponseWithDefaults() *GetStxSupplyResponse {
	this := GetStxSupplyResponse{}
	return &this
}

// GetUnlockedPercent returns the UnlockedPercent field value
func (o *GetStxSupplyResponse) GetUnlockedPercent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnlockedPercent
}

// GetUnlockedPercentOk returns a tuple with the UnlockedPercent field value
// and a boolean to check if the value has been set.
func (o *GetStxSupplyResponse) GetUnlockedPercentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnlockedPercent, true
}

// SetUnlockedPercent sets field value
func (o *GetStxSupplyResponse) SetUnlockedPercent(v string) {
	o.UnlockedPercent = v
}

// GetTotalStx returns the TotalStx field value
func (o *GetStxSupplyResponse) GetTotalStx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalStx
}

// GetTotalStxOk returns a tuple with the TotalStx field value
// and a boolean to check if the value has been set.
func (o *GetStxSupplyResponse) GetTotalStxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalStx, true
}

// SetTotalStx sets field value
func (o *GetStxSupplyResponse) SetTotalStx(v string) {
	o.TotalStx = v
}

// GetUnlockedStx returns the UnlockedStx field value
func (o *GetStxSupplyResponse) GetUnlockedStx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnlockedStx
}

// GetUnlockedStxOk returns a tuple with the UnlockedStx field value
// and a boolean to check if the value has been set.
func (o *GetStxSupplyResponse) GetUnlockedStxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnlockedStx, true
}

// SetUnlockedStx sets field value
func (o *GetStxSupplyResponse) SetUnlockedStx(v string) {
	o.UnlockedStx = v
}

// GetBlockHeight returns the BlockHeight field value
func (o *GetStxSupplyResponse) GetBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *GetStxSupplyResponse) GetBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *GetStxSupplyResponse) SetBlockHeight(v int32) {
	o.BlockHeight = v
}

func (o GetStxSupplyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStxSupplyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["unlocked_percent"] = o.UnlockedPercent
	toSerialize["total_stx"] = o.TotalStx
	toSerialize["unlocked_stx"] = o.UnlockedStx
	toSerialize["block_height"] = o.BlockHeight
	return toSerialize, nil
}

func (o *GetStxSupplyResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"unlocked_percent",
		"total_stx",
		"unlocked_stx",
		"block_height",
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

	varGetStxSupplyResponse := _GetStxSupplyResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetStxSupplyResponse)

	if err != nil {
		return err
	}

	*o = GetStxSupplyResponse(varGetStxSupplyResponse)

	return err
}

type NullableGetStxSupplyResponse struct {
	value *GetStxSupplyResponse
	isSet bool
}

func (v NullableGetStxSupplyResponse) Get() *GetStxSupplyResponse {
	return v.value
}

func (v *NullableGetStxSupplyResponse) Set(val *GetStxSupplyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStxSupplyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStxSupplyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStxSupplyResponse(val *GetStxSupplyResponse) *NullableGetStxSupplyResponse {
	return &NullableGetStxSupplyResponse{value: val, isSet: true}
}

func (v NullableGetStxSupplyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStxSupplyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
