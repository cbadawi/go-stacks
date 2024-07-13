# TransactionEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 
**Results** | [**[]TransactionEvent**](TransactionEvent.md) |  | 

## Methods

### NewTransactionEventsResponse

`func NewTransactionEventsResponse(limit int32, offset int32, results []TransactionEvent, ) *TransactionEventsResponse`

NewTransactionEventsResponse instantiates a new TransactionEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionEventsResponseWithDefaults

`func NewTransactionEventsResponseWithDefaults() *TransactionEventsResponse`

NewTransactionEventsResponseWithDefaults instantiates a new TransactionEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *TransactionEventsResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TransactionEventsResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TransactionEventsResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *TransactionEventsResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TransactionEventsResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TransactionEventsResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetResults

`func (o *TransactionEventsResponse) GetResults() []TransactionEvent`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TransactionEventsResponse) GetResultsOk() (*[]TransactionEvent, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TransactionEventsResponse) SetResults(v []TransactionEvent)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


