# AccountDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **string** |  | 
**Locked** | **string** |  | 
**UnlockHeight** | **int32** |  | 
**Nonce** | **int32** |  | 
**BalanceProof** | **string** |  | 
**NonceProof** | **string** |  | 

## Methods

### NewAccountDataResponse

`func NewAccountDataResponse(balance string, locked string, unlockHeight int32, nonce int32, balanceProof string, nonceProof string, ) *AccountDataResponse`

NewAccountDataResponse instantiates a new AccountDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDataResponseWithDefaults

`func NewAccountDataResponseWithDefaults() *AccountDataResponse`

NewAccountDataResponseWithDefaults instantiates a new AccountDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *AccountDataResponse) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AccountDataResponse) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AccountDataResponse) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetLocked

`func (o *AccountDataResponse) GetLocked() string`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *AccountDataResponse) GetLockedOk() (*string, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *AccountDataResponse) SetLocked(v string)`

SetLocked sets Locked field to given value.


### GetUnlockHeight

`func (o *AccountDataResponse) GetUnlockHeight() int32`

GetUnlockHeight returns the UnlockHeight field if non-nil, zero value otherwise.

### GetUnlockHeightOk

`func (o *AccountDataResponse) GetUnlockHeightOk() (*int32, bool)`

GetUnlockHeightOk returns a tuple with the UnlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockHeight

`func (o *AccountDataResponse) SetUnlockHeight(v int32)`

SetUnlockHeight sets UnlockHeight field to given value.


### GetNonce

`func (o *AccountDataResponse) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *AccountDataResponse) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *AccountDataResponse) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetBalanceProof

`func (o *AccountDataResponse) GetBalanceProof() string`

GetBalanceProof returns the BalanceProof field if non-nil, zero value otherwise.

### GetBalanceProofOk

`func (o *AccountDataResponse) GetBalanceProofOk() (*string, bool)`

GetBalanceProofOk returns a tuple with the BalanceProof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceProof

`func (o *AccountDataResponse) SetBalanceProof(v string)`

SetBalanceProof sets BalanceProof field to given value.


### GetNonceProof

`func (o *AccountDataResponse) GetNonceProof() string`

GetNonceProof returns the NonceProof field if non-nil, zero value otherwise.

### GetNonceProofOk

`func (o *AccountDataResponse) GetNonceProofOk() (*string, bool)`

GetNonceProofOk returns a tuple with the NonceProof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonceProof

`func (o *AccountDataResponse) SetNonceProof(v string)`

SetNonceProof sets NonceProof field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


