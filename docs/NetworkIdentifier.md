# NetworkIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Blockchain name | 
**Network** | **string** | If a blockchain has a specific chain-id or network identifier, it should go in this field. It is up to the client to determine which network-specific identifier is mainnet or testnet. | 
**SubNetworkIdentifier** | Pointer to [**NetworkIdentifierSubNetworkIdentifier**](NetworkIdentifierSubNetworkIdentifier.md) |  | [optional] 

## Methods

### NewNetworkIdentifier

`func NewNetworkIdentifier(blockchain string, network string, ) *NetworkIdentifier`

NewNetworkIdentifier instantiates a new NetworkIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkIdentifierWithDefaults

`func NewNetworkIdentifierWithDefaults() *NetworkIdentifier`

NewNetworkIdentifierWithDefaults instantiates a new NetworkIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *NetworkIdentifier) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *NetworkIdentifier) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *NetworkIdentifier) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *NetworkIdentifier) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkIdentifier) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkIdentifier) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetSubNetworkIdentifier

`func (o *NetworkIdentifier) GetSubNetworkIdentifier() NetworkIdentifierSubNetworkIdentifier`

GetSubNetworkIdentifier returns the SubNetworkIdentifier field if non-nil, zero value otherwise.

### GetSubNetworkIdentifierOk

`func (o *NetworkIdentifier) GetSubNetworkIdentifierOk() (*NetworkIdentifierSubNetworkIdentifier, bool)`

GetSubNetworkIdentifierOk returns a tuple with the SubNetworkIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNetworkIdentifier

`func (o *NetworkIdentifier) SetSubNetworkIdentifier(v NetworkIdentifierSubNetworkIdentifier)`

SetSubNetworkIdentifier sets SubNetworkIdentifier field to given value.

### HasSubNetworkIdentifier

`func (o *NetworkIdentifier) HasSubNetworkIdentifier() bool`

HasSubNetworkIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


