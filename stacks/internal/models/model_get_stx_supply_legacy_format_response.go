package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the GetStxSupplyLegacyFormatResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GetStxSupplyLegacyFormatResponse{}

// GetStxSupplyLegacyFormatResponse GET request that returns network target block times
type GetStxSupplyLegacyFormatResponse struct {
	// String quoted decimal number of the percentage of STX that have unlocked
	UnlockedPercent string `json:"unlockedPercent"`
	// String quoted decimal number of the total possible number of STX
	TotalStacks string `json:"totalStacks"`
	// Same as `totalStacks` but formatted with comma thousands separators
	TotalStacksFormatted string `json:"totalStacksFormatted"`
	// String quoted decimal number of the STX that have been mined or unlocked
	UnlockedSupply string `json:"unlockedSupply"`
	// Same as `unlockedSupply` but formatted with comma thousands separators
	UnlockedSupplyFormatted string `json:"unlockedSupplyFormatted"`
	// The block height at which this information was queried
	BlockHeight string `json:"blockHeight"`
}

type _GetStxSupplyLegacyFormatResponse GetStxSupplyLegacyFormatResponse

// NewGetStxSupplyLegacyFormatResponse instantiates a new GetStxSupplyLegacyFormatResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStxSupplyLegacyFormatResponse(unlockedPercent string, totalStacks string, totalStacksFormatted string, unlockedSupply string, unlockedSupplyFormatted string, blockHeight string) *GetStxSupplyLegacyFormatResponse {
	this := GetStxSupplyLegacyFormatResponse{}
	this.UnlockedPercent = unlockedPercent
	this.TotalStacks = totalStacks
	this.TotalStacksFormatted = totalStacksFormatted
	this.UnlockedSupply = unlockedSupply
	this.UnlockedSupplyFormatted = unlockedSupplyFormatted
	this.BlockHeight = blockHeight
	return &this
}

// NewGetStxSupplyLegacyFormatResponseWithDefaults instantiates a new GetStxSupplyLegacyFormatResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStxSupplyLegacyFormatResponseWithDefaults() *GetStxSupplyLegacyFormatResponse {
	this := GetStxSupplyLegacyFormatResponse{}
	return &this
}

// GetUnlockedPercent returns the UnlockedPercent field value
func (o *GetStxSupplyLegacyFormatResponse) GetUnlockedPercent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnlockedPercent
}

// GetUnlockedPercentOk returns a tuple with the UnlockedPercent field value
// and a boolean to check if the value has been set.
func (o *GetStxSupplyLegacyFormatResponse) GetUnlockedPercentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnlockedPercent, true
}

// SetUnlockedPercent sets field value
func (o *GetStxSupplyLegacyFormatResponse) SetUnlockedPercent(v string) {
	o.UnlockedPercent = v
}

// GetTotalStacks returns the TotalStacks field value
func (o *GetStxSupplyLegacyFormatResponse) GetTotalStacks() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalStacks
}

// GetTotalStacksOk returns a tuple with the TotalStacks field value
// and a boolean to check if the value has been set.
func (o *GetStxSupplyLegacyFormatResponse) GetTotalStacksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalStacks, true
}

// SetTotalStacks sets field value
func (o *GetStxSupplyLegacyFormatResponse) SetTotalStacks(v string) {
	o.TotalStacks = v
}

// GetTotalStacksFormatted returns the TotalStacksFormatted field value
func (o *GetStxSupplyLegacyFormatResponse) GetTotalStacksFormatted() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalStacksFormatted
}

// GetTotalStacksFormattedOk returns a tuple with the TotalStacksFormatted field value
// and a boolean to check if the value has been set.
func (o *GetStxSupplyLegacyFormatResponse) GetTotalStacksFormattedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalStacksFormatted, true
}

// SetTotalStacksFormatted sets field value
func (o *GetStxSupplyLegacyFormatResponse) SetTotalStacksFormatted(v string) {
	o.TotalStacksFormatted = v
}

// GetUnlockedSupply returns the UnlockedSupply field value
func (o *GetStxSupplyLegacyFormatResponse) GetUnlockedSupply() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnlockedSupply
}

// GetUnlockedSupplyOk returns a tuple with the UnlockedSupply field value
// and a boolean to check if the value has been set.
func (o *GetStxSupplyLegacyFormatResponse) GetUnlockedSupplyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnlockedSupply, true
}

// SetUnlockedSupply sets field value
func (o *GetStxSupplyLegacyFormatResponse) SetUnlockedSupply(v string) {
	o.UnlockedSupply = v
}

// GetUnlockedSupplyFormatted returns the UnlockedSupplyFormatted field value
func (o *GetStxSupplyLegacyFormatResponse) GetUnlockedSupplyFormatted() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnlockedSupplyFormatted
}

// GetUnlockedSupplyFormattedOk returns a tuple with the UnlockedSupplyFormatted field value
// and a boolean to check if the value has been set.
func (o *GetStxSupplyLegacyFormatResponse) GetUnlockedSupplyFormattedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnlockedSupplyFormatted, true
}

// SetUnlockedSupplyFormatted sets field value
func (o *GetStxSupplyLegacyFormatResponse) SetUnlockedSupplyFormatted(v string) {
	o.UnlockedSupplyFormatted = v
}

// GetBlockHeight returns the BlockHeight field value
func (o *GetStxSupplyLegacyFormatResponse) GetBlockHeight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *GetStxSupplyLegacyFormatResponse) GetBlockHeightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *GetStxSupplyLegacyFormatResponse) SetBlockHeight(v string) {
	o.BlockHeight = v
}

func (o GetStxSupplyLegacyFormatResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStxSupplyLegacyFormatResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["unlockedPercent"] = o.UnlockedPercent
	toSerialize["totalStacks"] = o.TotalStacks
	toSerialize["totalStacksFormatted"] = o.TotalStacksFormatted
	toSerialize["unlockedSupply"] = o.UnlockedSupply
	toSerialize["unlockedSupplyFormatted"] = o.UnlockedSupplyFormatted
	toSerialize["blockHeight"] = o.BlockHeight
	return toSerialize, nil
}

func (o *GetStxSupplyLegacyFormatResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"unlockedPercent",
		"totalStacks",
		"totalStacksFormatted",
		"unlockedSupply",
		"unlockedSupplyFormatted",
		"blockHeight",
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

	varGetStxSupplyLegacyFormatResponse := _GetStxSupplyLegacyFormatResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetStxSupplyLegacyFormatResponse)

	if err != nil {
		return err
	}

	*o = GetStxSupplyLegacyFormatResponse(varGetStxSupplyLegacyFormatResponse)

	return err
}

type NullableGetStxSupplyLegacyFormatResponse struct {
	value *GetStxSupplyLegacyFormatResponse
	isSet bool
}

func (v NullableGetStxSupplyLegacyFormatResponse) Get() *GetStxSupplyLegacyFormatResponse {
	return v.value
}

func (v *NullableGetStxSupplyLegacyFormatResponse) Set(val *GetStxSupplyLegacyFormatResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStxSupplyLegacyFormatResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStxSupplyLegacyFormatResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStxSupplyLegacyFormatResponse(val *GetStxSupplyLegacyFormatResponse) *NullableGetStxSupplyLegacyFormatResponse {
	return &NullableGetStxSupplyLegacyFormatResponse{value: val, isSet: true}
}

func (v NullableGetStxSupplyLegacyFormatResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStxSupplyLegacyFormatResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
