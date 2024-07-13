# RunFaucetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Indicates if the faucet call was successful | 
**TxId** | Pointer to **string** | The transaction ID for the faucet call | [optional] 
**TxRaw** | Pointer to **string** | Raw transaction in hex string representation | [optional] 

## Methods

### NewRunFaucetResponse

`func NewRunFaucetResponse(success bool, ) *RunFaucetResponse`

NewRunFaucetResponse instantiates a new RunFaucetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunFaucetResponseWithDefaults

`func NewRunFaucetResponseWithDefaults() *RunFaucetResponse`

NewRunFaucetResponseWithDefaults instantiates a new RunFaucetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *RunFaucetResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RunFaucetResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RunFaucetResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetTxId

`func (o *RunFaucetResponse) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *RunFaucetResponse) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *RunFaucetResponse) SetTxId(v string)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *RunFaucetResponse) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetTxRaw

`func (o *RunFaucetResponse) GetTxRaw() string`

GetTxRaw returns the TxRaw field if non-nil, zero value otherwise.

### GetTxRawOk

`func (o *RunFaucetResponse) GetTxRawOk() (*string, bool)`

GetTxRawOk returns a tuple with the TxRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRaw

`func (o *RunFaucetResponse) SetTxRaw(v string)`

SetTxRaw sets TxRaw field to given value.

### HasTxRaw

`func (o *RunFaucetResponse) HasTxRaw() bool`

HasTxRaw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


