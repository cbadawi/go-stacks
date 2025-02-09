package models

import (
	"encoding/json"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BnsFetchFileZoneResponseAnyOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BnsFetchFileZoneResponseAnyOf{}

// BnsFetchFileZoneResponseAnyOf struct for BnsFetchFileZoneResponseAnyOf
type BnsFetchFileZoneResponseAnyOf struct {
	Zonefile *string `json:"zonefile,omitempty" validate:"regexp=.+"`
}

// NewBnsFetchFileZoneResponseAnyOf instantiates a new BnsFetchFileZoneResponseAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBnsFetchFileZoneResponseAnyOf() *BnsFetchFileZoneResponseAnyOf {
	this := BnsFetchFileZoneResponseAnyOf{}
	return &this
}

// NewBnsFetchFileZoneResponseAnyOfWithDefaults instantiates a new BnsFetchFileZoneResponseAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBnsFetchFileZoneResponseAnyOfWithDefaults() *BnsFetchFileZoneResponseAnyOf {
	this := BnsFetchFileZoneResponseAnyOf{}
	return &this
}

// GetZonefile returns the Zonefile field value if set, zero value otherwise.
func (o *BnsFetchFileZoneResponseAnyOf) GetZonefile() string {
	if o == nil || utils.IsNil(o.Zonefile) {
		var ret string
		return ret
	}
	return *o.Zonefile
}

// GetZonefileOk returns a tuple with the Zonefile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BnsFetchFileZoneResponseAnyOf) GetZonefileOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Zonefile) {
		return nil, false
	}
	return o.Zonefile, true
}

// HasZonefile returns a boolean if a field has been set.
func (o *BnsFetchFileZoneResponseAnyOf) HasZonefile() bool {
	if o != nil && !utils.IsNil(o.Zonefile) {
		return true
	}

	return false
}

// SetZonefile gets a reference to the given string and assigns it to the Zonefile field.
func (o *BnsFetchFileZoneResponseAnyOf) SetZonefile(v string) {
	o.Zonefile = &v
}

func (o BnsFetchFileZoneResponseAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BnsFetchFileZoneResponseAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Zonefile) {
		toSerialize["zonefile"] = o.Zonefile
	}
	return toSerialize, nil
}

type NullableBnsFetchFileZoneResponseAnyOf struct {
	value *BnsFetchFileZoneResponseAnyOf
	isSet bool
}

func (v NullableBnsFetchFileZoneResponseAnyOf) Get() *BnsFetchFileZoneResponseAnyOf {
	return v.value
}

func (v *NullableBnsFetchFileZoneResponseAnyOf) Set(val *BnsFetchFileZoneResponseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBnsFetchFileZoneResponseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBnsFetchFileZoneResponseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBnsFetchFileZoneResponseAnyOf(val *BnsFetchFileZoneResponseAnyOf) *NullableBnsFetchFileZoneResponseAnyOf {
	return &NullableBnsFetchFileZoneResponseAnyOf{value: val, isSet: true}
}

func (v NullableBnsFetchFileZoneResponseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBnsFetchFileZoneResponseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
