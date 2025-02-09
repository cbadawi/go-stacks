package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the MempoolTxSearchResultResultTxData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MempoolTxSearchResultResultTxData{}

// MempoolTxSearchResultResultTxData Returns basic search result information about the requested id
type MempoolTxSearchResultResultTxData struct {
	TxType string `json:"tx_type"`
}

type _MempoolTxSearchResultResultTxData MempoolTxSearchResultResultTxData

// NewMempoolTxSearchResultResultTxData instantiates a new MempoolTxSearchResultResultTxData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMempoolTxSearchResultResultTxData(txType string) *MempoolTxSearchResultResultTxData {
	this := MempoolTxSearchResultResultTxData{}
	this.TxType = txType
	return &this
}

// NewMempoolTxSearchResultResultTxDataWithDefaults instantiates a new MempoolTxSearchResultResultTxData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMempoolTxSearchResultResultTxDataWithDefaults() *MempoolTxSearchResultResultTxData {
	this := MempoolTxSearchResultResultTxData{}
	return &this
}

// GetTxType returns the TxType field value
func (o *MempoolTxSearchResultResultTxData) GetTxType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxType
}

// GetTxTypeOk returns a tuple with the TxType field value
// and a boolean to check if the value has been set.
func (o *MempoolTxSearchResultResultTxData) GetTxTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxType, true
}

// SetTxType sets field value
func (o *MempoolTxSearchResultResultTxData) SetTxType(v string) {
	o.TxType = v
}

func (o MempoolTxSearchResultResultTxData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MempoolTxSearchResultResultTxData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tx_type"] = o.TxType
	return toSerialize, nil
}

func (o *MempoolTxSearchResultResultTxData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tx_type",
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

	varMempoolTxSearchResultResultTxData := _MempoolTxSearchResultResultTxData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMempoolTxSearchResultResultTxData)

	if err != nil {
		return err
	}

	*o = MempoolTxSearchResultResultTxData(varMempoolTxSearchResultResultTxData)

	return err
}

type NullableMempoolTxSearchResultResultTxData struct {
	value *MempoolTxSearchResultResultTxData
	isSet bool
}

func (v NullableMempoolTxSearchResultResultTxData) Get() *MempoolTxSearchResultResultTxData {
	return v.value
}

func (v *NullableMempoolTxSearchResultResultTxData) Set(val *MempoolTxSearchResultResultTxData) {
	v.value = val
	v.isSet = true
}

func (v NullableMempoolTxSearchResultResultTxData) IsSet() bool {
	return v.isSet
}

func (v *NullableMempoolTxSearchResultResultTxData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMempoolTxSearchResultResultTxData(val *MempoolTxSearchResultResultTxData) *NullableMempoolTxSearchResultResultTxData {
	return &NullableMempoolTxSearchResultResultTxData{value: val, isSet: true}
}

func (v NullableMempoolTxSearchResultResultTxData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMempoolTxSearchResultResultTxData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
