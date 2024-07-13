# RosettaAccountBalanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkIdentifier** | [**NetworkIdentifier**](NetworkIdentifier.md) |  | 
**AccountIdentifier** | [**RosettaAccount**](RosettaAccount.md) |  | 
**BlockIdentifier** | Pointer to [**RosettaPartialBlockIdentifier**](RosettaPartialBlockIdentifier.md) |  | [optional] 

## Methods

### NewRosettaAccountBalanceRequest

`func NewRosettaAccountBalanceRequest(networkIdentifier NetworkIdentifier, accountIdentifier RosettaAccount, ) *RosettaAccountBalanceRequest`

NewRosettaAccountBalanceRequest instantiates a new RosettaAccountBalanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaAccountBalanceRequestWithDefaults

`func NewRosettaAccountBalanceRequestWithDefaults() *RosettaAccountBalanceRequest`

NewRosettaAccountBalanceRequestWithDefaults instantiates a new RosettaAccountBalanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkIdentifier

`func (o *RosettaAccountBalanceRequest) GetNetworkIdentifier() NetworkIdentifier`

GetNetworkIdentifier returns the NetworkIdentifier field if non-nil, zero value otherwise.

### GetNetworkIdentifierOk

`func (o *RosettaAccountBalanceRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool)`

GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIdentifier

`func (o *RosettaAccountBalanceRequest) SetNetworkIdentifier(v NetworkIdentifier)`

SetNetworkIdentifier sets NetworkIdentifier field to given value.


### GetAccountIdentifier

`func (o *RosettaAccountBalanceRequest) GetAccountIdentifier() RosettaAccount`

GetAccountIdentifier returns the AccountIdentifier field if non-nil, zero value otherwise.

### GetAccountIdentifierOk

`func (o *RosettaAccountBalanceRequest) GetAccountIdentifierOk() (*RosettaAccount, bool)`

GetAccountIdentifierOk returns a tuple with the AccountIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentifier

`func (o *RosettaAccountBalanceRequest) SetAccountIdentifier(v RosettaAccount)`

SetAccountIdentifier sets AccountIdentifier field to given value.


### GetBlockIdentifier

`func (o *RosettaAccountBalanceRequest) GetBlockIdentifier() RosettaPartialBlockIdentifier`

GetBlockIdentifier returns the BlockIdentifier field if non-nil, zero value otherwise.

### GetBlockIdentifierOk

`func (o *RosettaAccountBalanceRequest) GetBlockIdentifierOk() (*RosettaPartialBlockIdentifier, bool)`

GetBlockIdentifierOk returns a tuple with the BlockIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockIdentifier

`func (o *RosettaAccountBalanceRequest) SetBlockIdentifier(v RosettaPartialBlockIdentifier)`

SetBlockIdentifier sets BlockIdentifier field to given value.

### HasBlockIdentifier

`func (o *RosettaAccountBalanceRequest) HasBlockIdentifier() bool`

HasBlockIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


