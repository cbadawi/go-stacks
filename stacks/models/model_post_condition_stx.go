package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the PostConditionStx type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PostConditionStx{}

// PostConditionStx struct for PostConditionStx
type PostConditionStx struct {
	Principal PostConditionPrincipal `json:"principal"`
	// A fungible condition code encodes a statement being made for either STX or a fungible token, with respect to the originating account.
	ConditionCode string `json:"condition_code"`
	Amount        string `json:"amount"`
	Type          string `json:"type"`
}

type _PostConditionStx PostConditionStx

// NewPostConditionStx instantiates a new PostConditionStx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostConditionStx(principal PostConditionPrincipal, conditionCode string, amount string, type_ string) *PostConditionStx {
	this := PostConditionStx{}
	this.Principal = principal
	this.ConditionCode = conditionCode
	this.Amount = amount
	this.Type = type_
	return &this
}

// NewPostConditionStxWithDefaults instantiates a new PostConditionStx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostConditionStxWithDefaults() *PostConditionStx {
	this := PostConditionStx{}
	return &this
}

// GetPrincipal returns the Principal field value
func (o *PostConditionStx) GetPrincipal() PostConditionPrincipal {
	if o == nil {
		var ret PostConditionPrincipal
		return ret
	}

	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value
// and a boolean to check if the value has been set.
func (o *PostConditionStx) GetPrincipalOk() (*PostConditionPrincipal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principal, true
}

// SetPrincipal sets field value
func (o *PostConditionStx) SetPrincipal(v PostConditionPrincipal) {
	o.Principal = v
}

// GetConditionCode returns the ConditionCode field value
func (o *PostConditionStx) GetConditionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionCode
}

// GetConditionCodeOk returns a tuple with the ConditionCode field value
// and a boolean to check if the value has been set.
func (o *PostConditionStx) GetConditionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionCode, true
}

// SetConditionCode sets field value
func (o *PostConditionStx) SetConditionCode(v string) {
	o.ConditionCode = v
}

// GetAmount returns the Amount field value
func (o *PostConditionStx) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PostConditionStx) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PostConditionStx) SetAmount(v string) {
	o.Amount = v
}

// GetType returns the Type field value
func (o *PostConditionStx) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostConditionStx) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostConditionStx) SetType(v string) {
	o.Type = v
}

func (o PostConditionStx) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostConditionStx) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principal"] = o.Principal
	toSerialize["condition_code"] = o.ConditionCode
	toSerialize["amount"] = o.Amount
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *PostConditionStx) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principal",
		"condition_code",
		"amount",
		"type",
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

	varPostConditionStx := _PostConditionStx{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostConditionStx)

	if err != nil {
		return err
	}

	*o = PostConditionStx(varPostConditionStx)

	return err
}

type NullablePostConditionStx struct {
	value *PostConditionStx
	isSet bool
}

func (v NullablePostConditionStx) Get() *PostConditionStx {
	return v.value
}

func (v *NullablePostConditionStx) Set(val *PostConditionStx) {
	v.value = val
	v.isSet = true
}

func (v NullablePostConditionStx) IsSet() bool {
	return v.isSet
}

func (v *NullablePostConditionStx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostConditionStx(val *PostConditionStx) *NullablePostConditionStx {
	return &NullablePostConditionStx{value: val, isSet: true}
}

func (v NullablePostConditionStx) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostConditionStx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
