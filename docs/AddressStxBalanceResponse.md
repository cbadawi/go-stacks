# AddressStxBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **string** |  | 
**TotalSent** | **string** |  | 
**TotalReceived** | **string** |  | 
**TotalFeesSent** | **string** |  | 
**TotalMinerRewardsReceived** | **string** |  | 
**LockTxId** | **string** | The transaction where the lock event occurred. Empty if no tokens are locked. | 
**Locked** | **string** | The amount of locked STX, as string quoted micro-STX. Zero if no tokens are locked. | 
**LockHeight** | **int32** | The STX chain block height of when the lock event occurred. Zero if no tokens are locked. | 
**BurnchainLockHeight** | **int32** | The burnchain block height of when the lock event occurred. Zero if no tokens are locked. | 
**BurnchainUnlockHeight** | **int32** | The burnchain block height of when the tokens unlock. Zero if no tokens are locked. | 
**TokenOfferingLocked** | Pointer to [**AddressTokenOfferingLocked**](AddressTokenOfferingLocked.md) |  | [optional] 

## Methods

### NewAddressStxBalanceResponse

`func NewAddressStxBalanceResponse(balance string, totalSent string, totalReceived string, totalFeesSent string, totalMinerRewardsReceived string, lockTxId string, locked string, lockHeight int32, burnchainLockHeight int32, burnchainUnlockHeight int32, ) *AddressStxBalanceResponse`

NewAddressStxBalanceResponse instantiates a new AddressStxBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressStxBalanceResponseWithDefaults

`func NewAddressStxBalanceResponseWithDefaults() *AddressStxBalanceResponse`

NewAddressStxBalanceResponseWithDefaults instantiates a new AddressStxBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *AddressStxBalanceResponse) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AddressStxBalanceResponse) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AddressStxBalanceResponse) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetTotalSent

`func (o *AddressStxBalanceResponse) GetTotalSent() string`

GetTotalSent returns the TotalSent field if non-nil, zero value otherwise.

### GetTotalSentOk

`func (o *AddressStxBalanceResponse) GetTotalSentOk() (*string, bool)`

GetTotalSentOk returns a tuple with the TotalSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSent

`func (o *AddressStxBalanceResponse) SetTotalSent(v string)`

SetTotalSent sets TotalSent field to given value.


### GetTotalReceived

`func (o *AddressStxBalanceResponse) GetTotalReceived() string`

GetTotalReceived returns the TotalReceived field if non-nil, zero value otherwise.

### GetTotalReceivedOk

`func (o *AddressStxBalanceResponse) GetTotalReceivedOk() (*string, bool)`

GetTotalReceivedOk returns a tuple with the TotalReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReceived

`func (o *AddressStxBalanceResponse) SetTotalReceived(v string)`

SetTotalReceived sets TotalReceived field to given value.


### GetTotalFeesSent

`func (o *AddressStxBalanceResponse) GetTotalFeesSent() string`

GetTotalFeesSent returns the TotalFeesSent field if non-nil, zero value otherwise.

### GetTotalFeesSentOk

`func (o *AddressStxBalanceResponse) GetTotalFeesSentOk() (*string, bool)`

GetTotalFeesSentOk returns a tuple with the TotalFeesSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFeesSent

`func (o *AddressStxBalanceResponse) SetTotalFeesSent(v string)`

SetTotalFeesSent sets TotalFeesSent field to given value.


### GetTotalMinerRewardsReceived

`func (o *AddressStxBalanceResponse) GetTotalMinerRewardsReceived() string`

GetTotalMinerRewardsReceived returns the TotalMinerRewardsReceived field if non-nil, zero value otherwise.

### GetTotalMinerRewardsReceivedOk

`func (o *AddressStxBalanceResponse) GetTotalMinerRewardsReceivedOk() (*string, bool)`

GetTotalMinerRewardsReceivedOk returns a tuple with the TotalMinerRewardsReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMinerRewardsReceived

`func (o *AddressStxBalanceResponse) SetTotalMinerRewardsReceived(v string)`

SetTotalMinerRewardsReceived sets TotalMinerRewardsReceived field to given value.


### GetLockTxId

`func (o *AddressStxBalanceResponse) GetLockTxId() string`

GetLockTxId returns the LockTxId field if non-nil, zero value otherwise.

### GetLockTxIdOk

`func (o *AddressStxBalanceResponse) GetLockTxIdOk() (*string, bool)`

GetLockTxIdOk returns a tuple with the LockTxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTxId

`func (o *AddressStxBalanceResponse) SetLockTxId(v string)`

SetLockTxId sets LockTxId field to given value.


### GetLocked

`func (o *AddressStxBalanceResponse) GetLocked() string`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *AddressStxBalanceResponse) GetLockedOk() (*string, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *AddressStxBalanceResponse) SetLocked(v string)`

SetLocked sets Locked field to given value.


### GetLockHeight

`func (o *AddressStxBalanceResponse) GetLockHeight() int32`

GetLockHeight returns the LockHeight field if non-nil, zero value otherwise.

### GetLockHeightOk

`func (o *AddressStxBalanceResponse) GetLockHeightOk() (*int32, bool)`

GetLockHeightOk returns a tuple with the LockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockHeight

`func (o *AddressStxBalanceResponse) SetLockHeight(v int32)`

SetLockHeight sets LockHeight field to given value.


### GetBurnchainLockHeight

`func (o *AddressStxBalanceResponse) GetBurnchainLockHeight() int32`

GetBurnchainLockHeight returns the BurnchainLockHeight field if non-nil, zero value otherwise.

### GetBurnchainLockHeightOk

`func (o *AddressStxBalanceResponse) GetBurnchainLockHeightOk() (*int32, bool)`

GetBurnchainLockHeightOk returns a tuple with the BurnchainLockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnchainLockHeight

`func (o *AddressStxBalanceResponse) SetBurnchainLockHeight(v int32)`

SetBurnchainLockHeight sets BurnchainLockHeight field to given value.


### GetBurnchainUnlockHeight

`func (o *AddressStxBalanceResponse) GetBurnchainUnlockHeight() int32`

GetBurnchainUnlockHeight returns the BurnchainUnlockHeight field if non-nil, zero value otherwise.

### GetBurnchainUnlockHeightOk

`func (o *AddressStxBalanceResponse) GetBurnchainUnlockHeightOk() (*int32, bool)`

GetBurnchainUnlockHeightOk returns a tuple with the BurnchainUnlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnchainUnlockHeight

`func (o *AddressStxBalanceResponse) SetBurnchainUnlockHeight(v int32)`

SetBurnchainUnlockHeight sets BurnchainUnlockHeight field to given value.


### GetTokenOfferingLocked

`func (o *AddressStxBalanceResponse) GetTokenOfferingLocked() AddressTokenOfferingLocked`

GetTokenOfferingLocked returns the TokenOfferingLocked field if non-nil, zero value otherwise.

### GetTokenOfferingLockedOk

`func (o *AddressStxBalanceResponse) GetTokenOfferingLockedOk() (*AddressTokenOfferingLocked, bool)`

GetTokenOfferingLockedOk returns a tuple with the TokenOfferingLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenOfferingLocked

`func (o *AddressStxBalanceResponse) SetTokenOfferingLocked(v AddressTokenOfferingLocked)`

SetTokenOfferingLocked sets TokenOfferingLocked field to given value.

### HasTokenOfferingLocked

`func (o *AddressStxBalanceResponse) HasTokenOfferingLocked() bool`

HasTokenOfferingLocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


