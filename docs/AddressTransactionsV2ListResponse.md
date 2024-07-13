# AddressTransactionsV2ListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 
**Total** | **int32** |  | 
**Results** | [**[]AddressTransaction**](AddressTransaction.md) |  | 

## Methods

### NewAddressTransactionsV2ListResponse

`func NewAddressTransactionsV2ListResponse(limit int32, offset int32, total int32, results []AddressTransaction, ) *AddressTransactionsV2ListResponse`

NewAddressTransactionsV2ListResponse instantiates a new AddressTransactionsV2ListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransactionsV2ListResponseWithDefaults

`func NewAddressTransactionsV2ListResponseWithDefaults() *AddressTransactionsV2ListResponse`

NewAddressTransactionsV2ListResponseWithDefaults instantiates a new AddressTransactionsV2ListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *AddressTransactionsV2ListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AddressTransactionsV2ListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AddressTransactionsV2ListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *AddressTransactionsV2ListResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AddressTransactionsV2ListResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AddressTransactionsV2ListResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *AddressTransactionsV2ListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AddressTransactionsV2ListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AddressTransactionsV2ListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *AddressTransactionsV2ListResponse) GetResults() []AddressTransaction`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AddressTransactionsV2ListResponse) GetResultsOk() (*[]AddressTransaction, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AddressTransactionsV2ListResponse) SetResults(v []AddressTransaction)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


