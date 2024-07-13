# TxSearchResultResultTxData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Canonical** | **bool** | If the transaction lies within the canonical chain | 
**BlockHash** | **string** | Refers to the hash of the block for searched transaction | 
**BurnBlockTime** | **int32** |  | 
**BlockHeight** | **int32** |  | 
**TxType** | **string** |  | 

## Methods

### NewTxSearchResultResultTxData

`func NewTxSearchResultResultTxData(canonical bool, blockHash string, burnBlockTime int32, blockHeight int32, txType string, ) *TxSearchResultResultTxData`

NewTxSearchResultResultTxData instantiates a new TxSearchResultResultTxData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTxSearchResultResultTxDataWithDefaults

`func NewTxSearchResultResultTxDataWithDefaults() *TxSearchResultResultTxData`

NewTxSearchResultResultTxDataWithDefaults instantiates a new TxSearchResultResultTxData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonical

`func (o *TxSearchResultResultTxData) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *TxSearchResultResultTxData) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *TxSearchResultResultTxData) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.


### GetBlockHash

`func (o *TxSearchResultResultTxData) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TxSearchResultResultTxData) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TxSearchResultResultTxData) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetBurnBlockTime

`func (o *TxSearchResultResultTxData) GetBurnBlockTime() int32`

GetBurnBlockTime returns the BurnBlockTime field if non-nil, zero value otherwise.

### GetBurnBlockTimeOk

`func (o *TxSearchResultResultTxData) GetBurnBlockTimeOk() (*int32, bool)`

GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTime

`func (o *TxSearchResultResultTxData) SetBurnBlockTime(v int32)`

SetBurnBlockTime sets BurnBlockTime field to given value.


### GetBlockHeight

`func (o *TxSearchResultResultTxData) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *TxSearchResultResultTxData) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *TxSearchResultResultTxData) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetTxType

`func (o *TxSearchResultResultTxData) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *TxSearchResultResultTxData) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *TxSearchResultResultTxData) SetTxType(v string)`

SetTxType sets TxType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


