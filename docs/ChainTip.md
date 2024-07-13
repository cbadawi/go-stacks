# ChainTip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockHeight** | **int32** | the current block height | 
**BlockHash** | **string** | the current block hash | 
**IndexBlockHash** | **string** | the current index block hash | 
**MicroblockHash** | Pointer to **string** | the current microblock hash | [optional] 
**MicroblockSequence** | Pointer to **int32** | the current microblock sequence number | [optional] 
**BurnBlockHeight** | **int32** | the current burn chain block height | 

## Methods

### NewChainTip

`func NewChainTip(blockHeight int32, blockHash string, indexBlockHash string, burnBlockHeight int32, ) *ChainTip`

NewChainTip instantiates a new ChainTip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChainTipWithDefaults

`func NewChainTipWithDefaults() *ChainTip`

NewChainTipWithDefaults instantiates a new ChainTip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockHeight

`func (o *ChainTip) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *ChainTip) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *ChainTip) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetBlockHash

`func (o *ChainTip) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *ChainTip) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *ChainTip) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetIndexBlockHash

`func (o *ChainTip) GetIndexBlockHash() string`

GetIndexBlockHash returns the IndexBlockHash field if non-nil, zero value otherwise.

### GetIndexBlockHashOk

`func (o *ChainTip) GetIndexBlockHashOk() (*string, bool)`

GetIndexBlockHashOk returns a tuple with the IndexBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexBlockHash

`func (o *ChainTip) SetIndexBlockHash(v string)`

SetIndexBlockHash sets IndexBlockHash field to given value.


### GetMicroblockHash

`func (o *ChainTip) GetMicroblockHash() string`

GetMicroblockHash returns the MicroblockHash field if non-nil, zero value otherwise.

### GetMicroblockHashOk

`func (o *ChainTip) GetMicroblockHashOk() (*string, bool)`

GetMicroblockHashOk returns a tuple with the MicroblockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockHash

`func (o *ChainTip) SetMicroblockHash(v string)`

SetMicroblockHash sets MicroblockHash field to given value.

### HasMicroblockHash

`func (o *ChainTip) HasMicroblockHash() bool`

HasMicroblockHash returns a boolean if a field has been set.

### GetMicroblockSequence

`func (o *ChainTip) GetMicroblockSequence() int32`

GetMicroblockSequence returns the MicroblockSequence field if non-nil, zero value otherwise.

### GetMicroblockSequenceOk

`func (o *ChainTip) GetMicroblockSequenceOk() (*int32, bool)`

GetMicroblockSequenceOk returns a tuple with the MicroblockSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockSequence

`func (o *ChainTip) SetMicroblockSequence(v int32)`

SetMicroblockSequence sets MicroblockSequence field to given value.

### HasMicroblockSequence

`func (o *ChainTip) HasMicroblockSequence() bool`

HasMicroblockSequence returns a boolean if a field has been set.

### GetBurnBlockHeight

`func (o *ChainTip) GetBurnBlockHeight() int32`

GetBurnBlockHeight returns the BurnBlockHeight field if non-nil, zero value otherwise.

### GetBurnBlockHeightOk

`func (o *ChainTip) GetBurnBlockHeightOk() (*int32, bool)`

GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHeight

`func (o *ChainTip) SetBurnBlockHeight(v int32)`

SetBurnBlockHeight sets BurnBlockHeight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


