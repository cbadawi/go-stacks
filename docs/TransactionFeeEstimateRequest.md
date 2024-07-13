# TransactionFeeEstimateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionPayload** | **string** |  | 
**EstimatedLen** | Pointer to **int32** |  | [optional] 

## Methods

### NewTransactionFeeEstimateRequest

`func NewTransactionFeeEstimateRequest(transactionPayload string, ) *TransactionFeeEstimateRequest`

NewTransactionFeeEstimateRequest instantiates a new TransactionFeeEstimateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionFeeEstimateRequestWithDefaults

`func NewTransactionFeeEstimateRequestWithDefaults() *TransactionFeeEstimateRequest`

NewTransactionFeeEstimateRequestWithDefaults instantiates a new TransactionFeeEstimateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionPayload

`func (o *TransactionFeeEstimateRequest) GetTransactionPayload() string`

GetTransactionPayload returns the TransactionPayload field if non-nil, zero value otherwise.

### GetTransactionPayloadOk

`func (o *TransactionFeeEstimateRequest) GetTransactionPayloadOk() (*string, bool)`

GetTransactionPayloadOk returns a tuple with the TransactionPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionPayload

`func (o *TransactionFeeEstimateRequest) SetTransactionPayload(v string)`

SetTransactionPayload sets TransactionPayload field to given value.


### GetEstimatedLen

`func (o *TransactionFeeEstimateRequest) GetEstimatedLen() int32`

GetEstimatedLen returns the EstimatedLen field if non-nil, zero value otherwise.

### GetEstimatedLenOk

`func (o *TransactionFeeEstimateRequest) GetEstimatedLenOk() (*int32, bool)`

GetEstimatedLenOk returns a tuple with the EstimatedLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedLen

`func (o *TransactionFeeEstimateRequest) SetEstimatedLen(v int32)`

SetEstimatedLen sets EstimatedLen field to given value.

### HasEstimatedLen

`func (o *TransactionFeeEstimateRequest) HasEstimatedLen() bool`

HasEstimatedLen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


