package models

import (
	"encoding/json"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BnsFetchHistoricalZoneFileResponseAnyOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BnsFetchHistoricalZoneFileResponseAnyOf{}

// BnsFetchHistoricalZoneFileResponseAnyOf struct for BnsFetchHistoricalZoneFileResponseAnyOf
type BnsFetchHistoricalZoneFileResponseAnyOf struct {
	Zonefile *string `json:"zonefile,omitempty"`
}

// NewBnsFetchHistoricalZoneFileResponseAnyOf instantiates a new BnsFetchHistoricalZoneFileResponseAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBnsFetchHistoricalZoneFileResponseAnyOf() *BnsFetchHistoricalZoneFileResponseAnyOf {
	this := BnsFetchHistoricalZoneFileResponseAnyOf{}
	return &this
}

// NewBnsFetchHistoricalZoneFileResponseAnyOfWithDefaults instantiates a new BnsFetchHistoricalZoneFileResponseAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBnsFetchHistoricalZoneFileResponseAnyOfWithDefaults() *BnsFetchHistoricalZoneFileResponseAnyOf {
	this := BnsFetchHistoricalZoneFileResponseAnyOf{}
	return &this
}

// GetZonefile returns the Zonefile field value if set, zero value otherwise.
func (o *BnsFetchHistoricalZoneFileResponseAnyOf) GetZonefile() string {
	if o == nil || utils.IsNil(o.Zonefile) {
		var ret string
		return ret
	}
	return *o.Zonefile
}

// GetZonefileOk returns a tuple with the Zonefile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BnsFetchHistoricalZoneFileResponseAnyOf) GetZonefileOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Zonefile) {
		return nil, false
	}
	return o.Zonefile, true
}

// HasZonefile returns a boolean if a field has been set.
func (o *BnsFetchHistoricalZoneFileResponseAnyOf) HasZonefile() bool {
	if o != nil && !utils.IsNil(o.Zonefile) {
		return true
	}

	return false
}

// SetZonefile gets a reference to the given string and assigns it to the Zonefile field.
func (o *BnsFetchHistoricalZoneFileResponseAnyOf) SetZonefile(v string) {
	o.Zonefile = &v
}

func (o BnsFetchHistoricalZoneFileResponseAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BnsFetchHistoricalZoneFileResponseAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Zonefile) {
		toSerialize["zonefile"] = o.Zonefile
	}
	return toSerialize, nil
}

type NullableBnsFetchHistoricalZoneFileResponseAnyOf struct {
	value *BnsFetchHistoricalZoneFileResponseAnyOf
	isSet bool
}

func (v NullableBnsFetchHistoricalZoneFileResponseAnyOf) Get() *BnsFetchHistoricalZoneFileResponseAnyOf {
	return v.value
}

func (v *NullableBnsFetchHistoricalZoneFileResponseAnyOf) Set(val *BnsFetchHistoricalZoneFileResponseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBnsFetchHistoricalZoneFileResponseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBnsFetchHistoricalZoneFileResponseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBnsFetchHistoricalZoneFileResponseAnyOf(val *BnsFetchHistoricalZoneFileResponseAnyOf) *NullableBnsFetchHistoricalZoneFileResponseAnyOf {
	return &NullableBnsFetchHistoricalZoneFileResponseAnyOf{value: val, isSet: true}
}

func (v NullableBnsFetchHistoricalZoneFileResponseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBnsFetchHistoricalZoneFileResponseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
