package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the MempoolTransactionStatsResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MempoolTransactionStatsResponse{}

// MempoolTransactionStatsResponse GET request that returns stats on mempool transactions
type MempoolTransactionStatsResponse struct {
	TxTypeCounts        MempoolTransactionStatsResponseTxTypeCounts        `json:"tx_type_counts"`
	TxSimpleFeeAverages MempoolTransactionStatsResponseTxSimpleFeeAverages `json:"tx_simple_fee_averages"`
	TxAges              MempoolTransactionStatsResponseTxAges              `json:"tx_ages"`
	TxByteSizes         MempoolTransactionStatsResponseTxByteSizes         `json:"tx_byte_sizes"`
}

type _MempoolTransactionStatsResponse MempoolTransactionStatsResponse

// NewMempoolTransactionStatsResponse instantiates a new MempoolTransactionStatsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMempoolTransactionStatsResponse(txTypeCounts MempoolTransactionStatsResponseTxTypeCounts, txSimpleFeeAverages MempoolTransactionStatsResponseTxSimpleFeeAverages, txAges MempoolTransactionStatsResponseTxAges, txByteSizes MempoolTransactionStatsResponseTxByteSizes) *MempoolTransactionStatsResponse {
	this := MempoolTransactionStatsResponse{}
	this.TxTypeCounts = txTypeCounts
	this.TxSimpleFeeAverages = txSimpleFeeAverages
	this.TxAges = txAges
	this.TxByteSizes = txByteSizes
	return &this
}

// NewMempoolTransactionStatsResponseWithDefaults instantiates a new MempoolTransactionStatsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMempoolTransactionStatsResponseWithDefaults() *MempoolTransactionStatsResponse {
	this := MempoolTransactionStatsResponse{}
	return &this
}

// GetTxTypeCounts returns the TxTypeCounts field value
func (o *MempoolTransactionStatsResponse) GetTxTypeCounts() MempoolTransactionStatsResponseTxTypeCounts {
	if o == nil {
		var ret MempoolTransactionStatsResponseTxTypeCounts
		return ret
	}

	return o.TxTypeCounts
}

// GetTxTypeCountsOk returns a tuple with the TxTypeCounts field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponse) GetTxTypeCountsOk() (*MempoolTransactionStatsResponseTxTypeCounts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxTypeCounts, true
}

// SetTxTypeCounts sets field value
func (o *MempoolTransactionStatsResponse) SetTxTypeCounts(v MempoolTransactionStatsResponseTxTypeCounts) {
	o.TxTypeCounts = v
}

// GetTxSimpleFeeAverages returns the TxSimpleFeeAverages field value
func (o *MempoolTransactionStatsResponse) GetTxSimpleFeeAverages() MempoolTransactionStatsResponseTxSimpleFeeAverages {
	if o == nil {
		var ret MempoolTransactionStatsResponseTxSimpleFeeAverages
		return ret
	}

	return o.TxSimpleFeeAverages
}

// GetTxSimpleFeeAveragesOk returns a tuple with the TxSimpleFeeAverages field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponse) GetTxSimpleFeeAveragesOk() (*MempoolTransactionStatsResponseTxSimpleFeeAverages, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxSimpleFeeAverages, true
}

// SetTxSimpleFeeAverages sets field value
func (o *MempoolTransactionStatsResponse) SetTxSimpleFeeAverages(v MempoolTransactionStatsResponseTxSimpleFeeAverages) {
	o.TxSimpleFeeAverages = v
}

// GetTxAges returns the TxAges field value
func (o *MempoolTransactionStatsResponse) GetTxAges() MempoolTransactionStatsResponseTxAges {
	if o == nil {
		var ret MempoolTransactionStatsResponseTxAges
		return ret
	}

	return o.TxAges
}

// GetTxAgesOk returns a tuple with the TxAges field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponse) GetTxAgesOk() (*MempoolTransactionStatsResponseTxAges, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxAges, true
}

// SetTxAges sets field value
func (o *MempoolTransactionStatsResponse) SetTxAges(v MempoolTransactionStatsResponseTxAges) {
	o.TxAges = v
}

// GetTxByteSizes returns the TxByteSizes field value
func (o *MempoolTransactionStatsResponse) GetTxByteSizes() MempoolTransactionStatsResponseTxByteSizes {
	if o == nil {
		var ret MempoolTransactionStatsResponseTxByteSizes
		return ret
	}

	return o.TxByteSizes
}

// GetTxByteSizesOk returns a tuple with the TxByteSizes field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponse) GetTxByteSizesOk() (*MempoolTransactionStatsResponseTxByteSizes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxByteSizes, true
}

// SetTxByteSizes sets field value
func (o *MempoolTransactionStatsResponse) SetTxByteSizes(v MempoolTransactionStatsResponseTxByteSizes) {
	o.TxByteSizes = v
}

func (o MempoolTransactionStatsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MempoolTransactionStatsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tx_type_counts"] = o.TxTypeCounts
	toSerialize["tx_simple_fee_averages"] = o.TxSimpleFeeAverages
	toSerialize["tx_ages"] = o.TxAges
	toSerialize["tx_byte_sizes"] = o.TxByteSizes
	return toSerialize, nil
}

func (o *MempoolTransactionStatsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tx_type_counts",
		"tx_simple_fee_averages",
		"tx_ages",
		"tx_byte_sizes",
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

	varMempoolTransactionStatsResponse := _MempoolTransactionStatsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMempoolTransactionStatsResponse)

	if err != nil {
		return err
	}

	*o = MempoolTransactionStatsResponse(varMempoolTransactionStatsResponse)

	return err
}

type NullableMempoolTransactionStatsResponse struct {
	value *MempoolTransactionStatsResponse
	isSet bool
}

func (v NullableMempoolTransactionStatsResponse) Get() *MempoolTransactionStatsResponse {
	return v.value
}

func (v *NullableMempoolTransactionStatsResponse) Set(val *MempoolTransactionStatsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMempoolTransactionStatsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMempoolTransactionStatsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMempoolTransactionStatsResponse(val *MempoolTransactionStatsResponse) *NullableMempoolTransactionStatsResponse {
	return &NullableMempoolTransactionStatsResponse{value: val, isSet: true}
}

func (v NullableMempoolTransactionStatsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMempoolTransactionStatsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
