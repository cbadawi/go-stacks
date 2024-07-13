# RosettaAccountBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockIdentifier** | [**RosettaBlockIdentifier**](RosettaBlockIdentifier.md) |  | 
**Balances** | [**[]RosettaAmount**](RosettaAmount.md) | A single account balance may have multiple currencies | 
**Coins** | Pointer to [**[]RosettaCoin**](RosettaCoin.md) | If a blockchain is UTXO-based, all unspent Coins owned by an account_identifier should be returned alongside the balance. It is highly recommended to populate this field so that users of the Rosetta API implementation don&#39;t need to maintain their own indexer to track their UTXOs. | [optional] 
**Metadata** | Pointer to [**RosettaAccountBalanceResponseMetadata**](RosettaAccountBalanceResponseMetadata.md) |  | [optional] 

## Methods

### NewRosettaAccountBalanceResponse

`func NewRosettaAccountBalanceResponse(blockIdentifier RosettaBlockIdentifier, balances []RosettaAmount, ) *RosettaAccountBalanceResponse`

NewRosettaAccountBalanceResponse instantiates a new RosettaAccountBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaAccountBalanceResponseWithDefaults

`func NewRosettaAccountBalanceResponseWithDefaults() *RosettaAccountBalanceResponse`

NewRosettaAccountBalanceResponseWithDefaults instantiates a new RosettaAccountBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockIdentifier

`func (o *RosettaAccountBalanceResponse) GetBlockIdentifier() RosettaBlockIdentifier`

GetBlockIdentifier returns the BlockIdentifier field if non-nil, zero value otherwise.

### GetBlockIdentifierOk

`func (o *RosettaAccountBalanceResponse) GetBlockIdentifierOk() (*RosettaBlockIdentifier, bool)`

GetBlockIdentifierOk returns a tuple with the BlockIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockIdentifier

`func (o *RosettaAccountBalanceResponse) SetBlockIdentifier(v RosettaBlockIdentifier)`

SetBlockIdentifier sets BlockIdentifier field to given value.


### GetBalances

`func (o *RosettaAccountBalanceResponse) GetBalances() []RosettaAmount`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *RosettaAccountBalanceResponse) GetBalancesOk() (*[]RosettaAmount, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *RosettaAccountBalanceResponse) SetBalances(v []RosettaAmount)`

SetBalances sets Balances field to given value.


### GetCoins

`func (o *RosettaAccountBalanceResponse) GetCoins() []RosettaCoin`

GetCoins returns the Coins field if non-nil, zero value otherwise.

### GetCoinsOk

`func (o *RosettaAccountBalanceResponse) GetCoinsOk() (*[]RosettaCoin, bool)`

GetCoinsOk returns a tuple with the Coins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoins

`func (o *RosettaAccountBalanceResponse) SetCoins(v []RosettaCoin)`

SetCoins sets Coins field to given value.

### HasCoins

`func (o *RosettaAccountBalanceResponse) HasCoins() bool`

HasCoins returns a boolean if a field has been set.

### GetMetadata

`func (o *RosettaAccountBalanceResponse) GetMetadata() RosettaAccountBalanceResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RosettaAccountBalanceResponse) GetMetadataOk() (*RosettaAccountBalanceResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RosettaAccountBalanceResponse) SetMetadata(v RosettaAccountBalanceResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RosettaAccountBalanceResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


