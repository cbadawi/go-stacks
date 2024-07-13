# TransactionFeeEstimateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstimatedCostScalar** | **int32** |  | 
**CostScalarChangeByByte** | Pointer to **float32** |  | [optional] 
**EstimatedCost** | [**TransactionFeeEstimateResponseEstimatedCost**](TransactionFeeEstimateResponseEstimatedCost.md) |  | 
**Estimations** | Pointer to [**[]TransactionFeeEstimateResponseEstimationsInner**](TransactionFeeEstimateResponseEstimationsInner.md) |  | [optional] 

## Methods

### NewTransactionFeeEstimateResponse

`func NewTransactionFeeEstimateResponse(estimatedCostScalar int32, estimatedCost TransactionFeeEstimateResponseEstimatedCost, ) *TransactionFeeEstimateResponse`

NewTransactionFeeEstimateResponse instantiates a new TransactionFeeEstimateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionFeeEstimateResponseWithDefaults

`func NewTransactionFeeEstimateResponseWithDefaults() *TransactionFeeEstimateResponse`

NewTransactionFeeEstimateResponseWithDefaults instantiates a new TransactionFeeEstimateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimatedCostScalar

`func (o *TransactionFeeEstimateResponse) GetEstimatedCostScalar() int32`

GetEstimatedCostScalar returns the EstimatedCostScalar field if non-nil, zero value otherwise.

### GetEstimatedCostScalarOk

`func (o *TransactionFeeEstimateResponse) GetEstimatedCostScalarOk() (*int32, bool)`

GetEstimatedCostScalarOk returns a tuple with the EstimatedCostScalar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCostScalar

`func (o *TransactionFeeEstimateResponse) SetEstimatedCostScalar(v int32)`

SetEstimatedCostScalar sets EstimatedCostScalar field to given value.


### GetCostScalarChangeByByte

`func (o *TransactionFeeEstimateResponse) GetCostScalarChangeByByte() float32`

GetCostScalarChangeByByte returns the CostScalarChangeByByte field if non-nil, zero value otherwise.

### GetCostScalarChangeByByteOk

`func (o *TransactionFeeEstimateResponse) GetCostScalarChangeByByteOk() (*float32, bool)`

GetCostScalarChangeByByteOk returns a tuple with the CostScalarChangeByByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScalarChangeByByte

`func (o *TransactionFeeEstimateResponse) SetCostScalarChangeByByte(v float32)`

SetCostScalarChangeByByte sets CostScalarChangeByByte field to given value.

### HasCostScalarChangeByByte

`func (o *TransactionFeeEstimateResponse) HasCostScalarChangeByByte() bool`

HasCostScalarChangeByByte returns a boolean if a field has been set.

### GetEstimatedCost

`func (o *TransactionFeeEstimateResponse) GetEstimatedCost() TransactionFeeEstimateResponseEstimatedCost`

GetEstimatedCost returns the EstimatedCost field if non-nil, zero value otherwise.

### GetEstimatedCostOk

`func (o *TransactionFeeEstimateResponse) GetEstimatedCostOk() (*TransactionFeeEstimateResponseEstimatedCost, bool)`

GetEstimatedCostOk returns a tuple with the EstimatedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCost

`func (o *TransactionFeeEstimateResponse) SetEstimatedCost(v TransactionFeeEstimateResponseEstimatedCost)`

SetEstimatedCost sets EstimatedCost field to given value.


### GetEstimations

`func (o *TransactionFeeEstimateResponse) GetEstimations() []TransactionFeeEstimateResponseEstimationsInner`

GetEstimations returns the Estimations field if non-nil, zero value otherwise.

### GetEstimationsOk

`func (o *TransactionFeeEstimateResponse) GetEstimationsOk() (*[]TransactionFeeEstimateResponseEstimationsInner, bool)`

GetEstimationsOk returns a tuple with the Estimations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimations

`func (o *TransactionFeeEstimateResponse) SetEstimations(v []TransactionFeeEstimateResponseEstimationsInner)`

SetEstimations sets Estimations field to given value.

### HasEstimations

`func (o *TransactionFeeEstimateResponse) HasEstimations() bool`

HasEstimations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


