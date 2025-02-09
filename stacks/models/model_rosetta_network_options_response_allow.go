package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the RosettaNetworkOptionsResponseAllow type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RosettaNetworkOptionsResponseAllow{}

// RosettaNetworkOptionsResponseAllow Allow specifies supported Operation status, Operation types, and all possible error statuses. This Allow object is used by clients to validate the correctness of a Rosetta Server implementation. It is expected that these clients will error if they receive some response that contains any of the above information that is not specified here.
type RosettaNetworkOptionsResponseAllow struct {
	// All Operation.Status this implementation supports. Any status that is returned during parsing that is not listed here will cause client validation to error.
	OperationStatuses []RosettaOperationStatus `json:"operation_statuses"`
	// All Operation.Type this implementation supports. Any type that is returned during parsing that is not listed here will cause client validation to error.
	OperationTypes []RosettaNetworkOptionsResponseAllowOperationTypesInner `json:"operation_types"`
	// All Errors that this implementation could return. Any error that is returned during parsing that is not listed here will cause client validation to error.
	Errors []RosettaErrorNoDetails `json:"errors"`
	// Any Rosetta implementation that supports querying the balance of an account at any height in the past should set this to true.
	HistoricalBalanceLookup bool `json:"historical_balance_lookup"`
}

type _RosettaNetworkOptionsResponseAllow RosettaNetworkOptionsResponseAllow

// NewRosettaNetworkOptionsResponseAllow instantiates a new RosettaNetworkOptionsResponseAllow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRosettaNetworkOptionsResponseAllow(operationStatuses []RosettaOperationStatus, operationTypes []RosettaNetworkOptionsResponseAllowOperationTypesInner, errors []RosettaErrorNoDetails, historicalBalanceLookup bool) *RosettaNetworkOptionsResponseAllow {
	this := RosettaNetworkOptionsResponseAllow{}
	this.OperationStatuses = operationStatuses
	this.OperationTypes = operationTypes
	this.Errors = errors
	this.HistoricalBalanceLookup = historicalBalanceLookup
	return &this
}

// NewRosettaNetworkOptionsResponseAllowWithDefaults instantiates a new RosettaNetworkOptionsResponseAllow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRosettaNetworkOptionsResponseAllowWithDefaults() *RosettaNetworkOptionsResponseAllow {
	this := RosettaNetworkOptionsResponseAllow{}
	return &this
}

// GetOperationStatuses returns the OperationStatuses field value
func (o *RosettaNetworkOptionsResponseAllow) GetOperationStatuses() []RosettaOperationStatus {
	if o == nil {
		var ret []RosettaOperationStatus
		return ret
	}

	return o.OperationStatuses
}

// GetOperationStatusesOk returns a tuple with the OperationStatuses field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkOptionsResponseAllow) GetOperationStatusesOk() ([]RosettaOperationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperationStatuses, true
}

// SetOperationStatuses sets field value
func (o *RosettaNetworkOptionsResponseAllow) SetOperationStatuses(v []RosettaOperationStatus) {
	o.OperationStatuses = v
}

// GetOperationTypes returns the OperationTypes field value
func (o *RosettaNetworkOptionsResponseAllow) GetOperationTypes() []RosettaNetworkOptionsResponseAllowOperationTypesInner {
	if o == nil {
		var ret []RosettaNetworkOptionsResponseAllowOperationTypesInner
		return ret
	}

	return o.OperationTypes
}

// GetOperationTypesOk returns a tuple with the OperationTypes field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkOptionsResponseAllow) GetOperationTypesOk() ([]RosettaNetworkOptionsResponseAllowOperationTypesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperationTypes, true
}

// SetOperationTypes sets field value
func (o *RosettaNetworkOptionsResponseAllow) SetOperationTypes(v []RosettaNetworkOptionsResponseAllowOperationTypesInner) {
	o.OperationTypes = v
}

// GetErrors returns the Errors field value
func (o *RosettaNetworkOptionsResponseAllow) GetErrors() []RosettaErrorNoDetails {
	if o == nil {
		var ret []RosettaErrorNoDetails
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkOptionsResponseAllow) GetErrorsOk() ([]RosettaErrorNoDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *RosettaNetworkOptionsResponseAllow) SetErrors(v []RosettaErrorNoDetails) {
	o.Errors = v
}

// GetHistoricalBalanceLookup returns the HistoricalBalanceLookup field value
func (o *RosettaNetworkOptionsResponseAllow) GetHistoricalBalanceLookup() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HistoricalBalanceLookup
}

// GetHistoricalBalanceLookupOk returns a tuple with the HistoricalBalanceLookup field value
// and a boolean to check if the value has been set.
func (o *RosettaNetworkOptionsResponseAllow) GetHistoricalBalanceLookupOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HistoricalBalanceLookup, true
}

// SetHistoricalBalanceLookup sets field value
func (o *RosettaNetworkOptionsResponseAllow) SetHistoricalBalanceLookup(v bool) {
	o.HistoricalBalanceLookup = v
}

func (o RosettaNetworkOptionsResponseAllow) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RosettaNetworkOptionsResponseAllow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operation_statuses"] = o.OperationStatuses
	toSerialize["operation_types"] = o.OperationTypes
	toSerialize["errors"] = o.Errors
	toSerialize["historical_balance_lookup"] = o.HistoricalBalanceLookup
	return toSerialize, nil
}

func (o *RosettaNetworkOptionsResponseAllow) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"operation_statuses",
		"operation_types",
		"errors",
		"historical_balance_lookup",
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

	varRosettaNetworkOptionsResponseAllow := _RosettaNetworkOptionsResponseAllow{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRosettaNetworkOptionsResponseAllow)

	if err != nil {
		return err
	}

	*o = RosettaNetworkOptionsResponseAllow(varRosettaNetworkOptionsResponseAllow)

	return err
}

type NullableRosettaNetworkOptionsResponseAllow struct {
	value *RosettaNetworkOptionsResponseAllow
	isSet bool
}

func (v NullableRosettaNetworkOptionsResponseAllow) Get() *RosettaNetworkOptionsResponseAllow {
	return v.value
}

func (v *NullableRosettaNetworkOptionsResponseAllow) Set(val *RosettaNetworkOptionsResponseAllow) {
	v.value = val
	v.isSet = true
}

func (v NullableRosettaNetworkOptionsResponseAllow) IsSet() bool {
	return v.isSet
}

func (v *NullableRosettaNetworkOptionsResponseAllow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRosettaNetworkOptionsResponseAllow(val *RosettaNetworkOptionsResponseAllow) *NullableRosettaNetworkOptionsResponseAllow {
	return &NullableRosettaNetworkOptionsResponseAllow{value: val, isSet: true}
}

func (v NullableRosettaNetworkOptionsResponseAllow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRosettaNetworkOptionsResponseAllow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
