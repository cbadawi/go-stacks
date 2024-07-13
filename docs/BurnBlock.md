# BurnBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BurnBlockTime** | **float32** | Unix timestamp (in seconds) indicating when this block was mined. | 
**BurnBlockTimeIso** | **string** | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined. | 
**BurnBlockHash** | **string** | Hash of the anchor chain block | 
**BurnBlockHeight** | **int32** | Height of the anchor chain block | 
**StacksBlocks** | **[]string** | Hashes of the Stacks blocks included in the burn block | 
**AvgBlockTime** | **float32** | Average time between blocks in seconds. Returns 0 if there is only one block in the burn block. | 
**TotalTxCount** | **int32** | Total number of transactions in the Stacks blocks associated with this burn block | 

## Methods

### NewBurnBlock

`func NewBurnBlock(burnBlockTime float32, burnBlockTimeIso string, burnBlockHash string, burnBlockHeight int32, stacksBlocks []string, avgBlockTime float32, totalTxCount int32, ) *BurnBlock`

NewBurnBlock instantiates a new BurnBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBurnBlockWithDefaults

`func NewBurnBlockWithDefaults() *BurnBlock`

NewBurnBlockWithDefaults instantiates a new BurnBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBurnBlockTime

`func (o *BurnBlock) GetBurnBlockTime() float32`

GetBurnBlockTime returns the BurnBlockTime field if non-nil, zero value otherwise.

### GetBurnBlockTimeOk

`func (o *BurnBlock) GetBurnBlockTimeOk() (*float32, bool)`

GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTime

`func (o *BurnBlock) SetBurnBlockTime(v float32)`

SetBurnBlockTime sets BurnBlockTime field to given value.


### GetBurnBlockTimeIso

`func (o *BurnBlock) GetBurnBlockTimeIso() string`

GetBurnBlockTimeIso returns the BurnBlockTimeIso field if non-nil, zero value otherwise.

### GetBurnBlockTimeIsoOk

`func (o *BurnBlock) GetBurnBlockTimeIsoOk() (*string, bool)`

GetBurnBlockTimeIsoOk returns a tuple with the BurnBlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTimeIso

`func (o *BurnBlock) SetBurnBlockTimeIso(v string)`

SetBurnBlockTimeIso sets BurnBlockTimeIso field to given value.


### GetBurnBlockHash

`func (o *BurnBlock) GetBurnBlockHash() string`

GetBurnBlockHash returns the BurnBlockHash field if non-nil, zero value otherwise.

### GetBurnBlockHashOk

`func (o *BurnBlock) GetBurnBlockHashOk() (*string, bool)`

GetBurnBlockHashOk returns a tuple with the BurnBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHash

`func (o *BurnBlock) SetBurnBlockHash(v string)`

SetBurnBlockHash sets BurnBlockHash field to given value.


### GetBurnBlockHeight

`func (o *BurnBlock) GetBurnBlockHeight() int32`

GetBurnBlockHeight returns the BurnBlockHeight field if non-nil, zero value otherwise.

### GetBurnBlockHeightOk

`func (o *BurnBlock) GetBurnBlockHeightOk() (*int32, bool)`

GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHeight

`func (o *BurnBlock) SetBurnBlockHeight(v int32)`

SetBurnBlockHeight sets BurnBlockHeight field to given value.


### GetStacksBlocks

`func (o *BurnBlock) GetStacksBlocks() []string`

GetStacksBlocks returns the StacksBlocks field if non-nil, zero value otherwise.

### GetStacksBlocksOk

`func (o *BurnBlock) GetStacksBlocksOk() (*[]string, bool)`

GetStacksBlocksOk returns a tuple with the StacksBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacksBlocks

`func (o *BurnBlock) SetStacksBlocks(v []string)`

SetStacksBlocks sets StacksBlocks field to given value.


### GetAvgBlockTime

`func (o *BurnBlock) GetAvgBlockTime() float32`

GetAvgBlockTime returns the AvgBlockTime field if non-nil, zero value otherwise.

### GetAvgBlockTimeOk

`func (o *BurnBlock) GetAvgBlockTimeOk() (*float32, bool)`

GetAvgBlockTimeOk returns a tuple with the AvgBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgBlockTime

`func (o *BurnBlock) SetAvgBlockTime(v float32)`

SetAvgBlockTime sets AvgBlockTime field to given value.


### GetTotalTxCount

`func (o *BurnBlock) GetTotalTxCount() int32`

GetTotalTxCount returns the TotalTxCount field if non-nil, zero value otherwise.

### GetTotalTxCountOk

`func (o *BurnBlock) GetTotalTxCountOk() (*int32, bool)`

GetTotalTxCountOk returns a tuple with the TotalTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTxCount

`func (o *BurnBlock) SetTotalTxCount(v int32)`

SetTotalTxCount sets TotalTxCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


