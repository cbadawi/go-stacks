package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the MempoolTransactionListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MempoolTransactionListResponse{}

// MempoolTransactionListResponse GET request that returns transactions
type MempoolTransactionListResponse struct {
	Limit   int32                `json:"limit"`
	Offset  int32                `json:"offset"`
	Total   int32                `json:"total"`
	Results []MempoolTransaction `json:"results"`
}

type _MempoolTransactionListResponse MempoolTransactionListResponse

// NewMempoolTransactionListResponse instantiates a new MempoolTransactionListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMempoolTransactionListResponse(limit int32, offset int32, total int32, results []MempoolTransaction) *MempoolTransactionListResponse {
	this := MempoolTransactionListResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewMempoolTransactionListResponseWithDefaults instantiates a new MempoolTransactionListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMempoolTransactionListResponseWithDefaults() *MempoolTransactionListResponse {
	this := MempoolTransactionListResponse{}
	return &this
}

// GetLimit returns the Limit field value
func (o *MempoolTransactionListResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionListResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *MempoolTransactionListResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *MempoolTransactionListResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionListResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *MempoolTransactionListResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *MempoolTransactionListResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionListResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *MempoolTransactionListResponse) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *MempoolTransactionListResponse) GetResults() []MempoolTransaction {
	if o == nil {
		var ret []MempoolTransaction
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionListResponse) GetResultsOk() ([]MempoolTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *MempoolTransactionListResponse) SetResults(v []MempoolTransaction) {
	o.Results = v
}

func (o MempoolTransactionListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MempoolTransactionListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *MempoolTransactionListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"limit",
		"offset",
		"total",
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

	varMempoolTransactionListResponse := _MempoolTransactionListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMempoolTransactionListResponse)

	if err != nil {
		return err
	}

	*o = MempoolTransactionListResponse(varMempoolTransactionListResponse)

	return err
}

type NullableMempoolTransactionListResponse struct {
	value *MempoolTransactionListResponse
	isSet bool
}

func (v NullableMempoolTransactionListResponse) Get() *MempoolTransactionListResponse {
	return v.value
}

func (v *NullableMempoolTransactionListResponse) Set(val *MempoolTransactionListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMempoolTransactionListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMempoolTransactionListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMempoolTransactionListResponse(val *MempoolTransactionListResponse) *NullableMempoolTransactionListResponse {
	return &NullableMempoolTransactionListResponse{value: val, isSet: true}
}

func (v NullableMempoolTransactionListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMempoolTransactionListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
