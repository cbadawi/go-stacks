# RosettaNetworkOptionsResponseAllow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationStatuses** | [**[]RosettaOperationStatus**](RosettaOperationStatus.md) | All Operation.Status this implementation supports. Any status that is returned during parsing that is not listed here will cause client validation to error. | 
**OperationTypes** | [**[]RosettaNetworkOptionsResponseAllowOperationTypesInner**](RosettaNetworkOptionsResponseAllowOperationTypesInner.md) | All Operation.Type this implementation supports. Any type that is returned during parsing that is not listed here will cause client validation to error. | 
**Errors** | [**[]RosettaErrorNoDetails**](RosettaErrorNoDetails.md) | All Errors that this implementation could return. Any error that is returned during parsing that is not listed here will cause client validation to error. | 
**HistoricalBalanceLookup** | **bool** | Any Rosetta implementation that supports querying the balance of an account at any height in the past should set this to true. | 

## Methods

### NewRosettaNetworkOptionsResponseAllow

`func NewRosettaNetworkOptionsResponseAllow(operationStatuses []RosettaOperationStatus, operationTypes []RosettaNetworkOptionsResponseAllowOperationTypesInner, errors []RosettaErrorNoDetails, historicalBalanceLookup bool, ) *RosettaNetworkOptionsResponseAllow`

NewRosettaNetworkOptionsResponseAllow instantiates a new RosettaNetworkOptionsResponseAllow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaNetworkOptionsResponseAllowWithDefaults

`func NewRosettaNetworkOptionsResponseAllowWithDefaults() *RosettaNetworkOptionsResponseAllow`

NewRosettaNetworkOptionsResponseAllowWithDefaults instantiates a new RosettaNetworkOptionsResponseAllow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationStatuses

`func (o *RosettaNetworkOptionsResponseAllow) GetOperationStatuses() []RosettaOperationStatus`

GetOperationStatuses returns the OperationStatuses field if non-nil, zero value otherwise.

### GetOperationStatusesOk

`func (o *RosettaNetworkOptionsResponseAllow) GetOperationStatusesOk() (*[]RosettaOperationStatus, bool)`

GetOperationStatusesOk returns a tuple with the OperationStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationStatuses

`func (o *RosettaNetworkOptionsResponseAllow) SetOperationStatuses(v []RosettaOperationStatus)`

SetOperationStatuses sets OperationStatuses field to given value.


### GetOperationTypes

`func (o *RosettaNetworkOptionsResponseAllow) GetOperationTypes() []RosettaNetworkOptionsResponseAllowOperationTypesInner`

GetOperationTypes returns the OperationTypes field if non-nil, zero value otherwise.

### GetOperationTypesOk

`func (o *RosettaNetworkOptionsResponseAllow) GetOperationTypesOk() (*[]RosettaNetworkOptionsResponseAllowOperationTypesInner, bool)`

GetOperationTypesOk returns a tuple with the OperationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationTypes

`func (o *RosettaNetworkOptionsResponseAllow) SetOperationTypes(v []RosettaNetworkOptionsResponseAllowOperationTypesInner)`

SetOperationTypes sets OperationTypes field to given value.


### GetErrors

`func (o *RosettaNetworkOptionsResponseAllow) GetErrors() []RosettaErrorNoDetails`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *RosettaNetworkOptionsResponseAllow) GetErrorsOk() (*[]RosettaErrorNoDetails, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *RosettaNetworkOptionsResponseAllow) SetErrors(v []RosettaErrorNoDetails)`

SetErrors sets Errors field to given value.


### GetHistoricalBalanceLookup

`func (o *RosettaNetworkOptionsResponseAllow) GetHistoricalBalanceLookup() bool`

GetHistoricalBalanceLookup returns the HistoricalBalanceLookup field if non-nil, zero value otherwise.

### GetHistoricalBalanceLookupOk

`func (o *RosettaNetworkOptionsResponseAllow) GetHistoricalBalanceLookupOk() (*bool, bool)`

GetHistoricalBalanceLookupOk returns a tuple with the HistoricalBalanceLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricalBalanceLookup

`func (o *RosettaNetworkOptionsResponseAllow) SetHistoricalBalanceLookup(v bool)`

SetHistoricalBalanceLookup sets HistoricalBalanceLookup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


