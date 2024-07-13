# NonFungibleTokenHolding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetIdentifier** | **string** |  | 
**Value** | [**NonFungibleTokenHoldingWithTxIdValue**](NonFungibleTokenHoldingWithTxIdValue.md) |  | 
**BlockHeight** | **float32** |  | 
**TxId** | **string** |  | 
**Tx** | [**Transaction**](Transaction.md) |  | 

## Methods

### NewNonFungibleTokenHolding

`func NewNonFungibleTokenHolding(assetIdentifier string, value NonFungibleTokenHoldingWithTxIdValue, blockHeight float32, txId string, tx Transaction, ) *NonFungibleTokenHolding`

NewNonFungibleTokenHolding instantiates a new NonFungibleTokenHolding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonFungibleTokenHoldingWithDefaults

`func NewNonFungibleTokenHoldingWithDefaults() *NonFungibleTokenHolding`

NewNonFungibleTokenHoldingWithDefaults instantiates a new NonFungibleTokenHolding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetIdentifier

`func (o *NonFungibleTokenHolding) GetAssetIdentifier() string`

GetAssetIdentifier returns the AssetIdentifier field if non-nil, zero value otherwise.

### GetAssetIdentifierOk

`func (o *NonFungibleTokenHolding) GetAssetIdentifierOk() (*string, bool)`

GetAssetIdentifierOk returns a tuple with the AssetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdentifier

`func (o *NonFungibleTokenHolding) SetAssetIdentifier(v string)`

SetAssetIdentifier sets AssetIdentifier field to given value.


### GetValue

`func (o *NonFungibleTokenHolding) GetValue() NonFungibleTokenHoldingWithTxIdValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NonFungibleTokenHolding) GetValueOk() (*NonFungibleTokenHoldingWithTxIdValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NonFungibleTokenHolding) SetValue(v NonFungibleTokenHoldingWithTxIdValue)`

SetValue sets Value field to given value.


### GetBlockHeight

`func (o *NonFungibleTokenHolding) GetBlockHeight() float32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *NonFungibleTokenHolding) GetBlockHeightOk() (*float32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *NonFungibleTokenHolding) SetBlockHeight(v float32)`

SetBlockHeight sets BlockHeight field to given value.


### GetTxId

`func (o *NonFungibleTokenHolding) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *NonFungibleTokenHolding) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *NonFungibleTokenHolding) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetTx

`func (o *NonFungibleTokenHolding) GetTx() Transaction`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *NonFungibleTokenHolding) GetTxOk() (*Transaction, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *NonFungibleTokenHolding) SetTx(v Transaction)`

SetTx sets Tx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


