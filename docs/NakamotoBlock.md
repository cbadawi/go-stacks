# NakamotoBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Canonical** | **bool** | Set to &#x60;true&#x60; if block corresponds to the canonical chain tip | 
**Height** | **int32** | Height of the block | 
**Hash** | **string** | Hash representing the block | 
**BlockTime** | **float32** | Unix timestamp (in seconds) indicating when this block was mined. | 
**BlockTimeIso** | **string** | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined. | 
**IndexBlockHash** | **string** | The only hash that can uniquely identify an anchored block or an unconfirmed state trie | 
**ParentBlockHash** | **string** | Hash of the parent block | 
**ParentIndexBlockHash** | **string** | Index block hash of the parent block | 
**BurnBlockTime** | **float32** | Unix timestamp (in seconds) indicating when this block was mined. | 
**BurnBlockTimeIso** | **string** | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined. | 
**BurnBlockHash** | **string** | Hash of the anchor chain block | 
**BurnBlockHeight** | **int32** | Height of the anchor chain block | 
**MinerTxid** | **string** | Anchor chain transaction ID | 
**TxCount** | **int32** | Number of transactions included in the block | 
**ExecutionCostReadCount** | **int32** | Execution cost read count. | 
**ExecutionCostReadLength** | **int32** | Execution cost read length. | 
**ExecutionCostRuntime** | **int32** | Execution cost runtime. | 
**ExecutionCostWriteCount** | **int32** | Execution cost write count. | 
**ExecutionCostWriteLength** | **int32** | Execution cost write length. | 

## Methods

### NewNakamotoBlock

`func NewNakamotoBlock(canonical bool, height int32, hash string, blockTime float32, blockTimeIso string, indexBlockHash string, parentBlockHash string, parentIndexBlockHash string, burnBlockTime float32, burnBlockTimeIso string, burnBlockHash string, burnBlockHeight int32, minerTxid string, txCount int32, executionCostReadCount int32, executionCostReadLength int32, executionCostRuntime int32, executionCostWriteCount int32, executionCostWriteLength int32, ) *NakamotoBlock`

NewNakamotoBlock instantiates a new NakamotoBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNakamotoBlockWithDefaults

`func NewNakamotoBlockWithDefaults() *NakamotoBlock`

NewNakamotoBlockWithDefaults instantiates a new NakamotoBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonical

`func (o *NakamotoBlock) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *NakamotoBlock) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *NakamotoBlock) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.


### GetHeight

`func (o *NakamotoBlock) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *NakamotoBlock) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *NakamotoBlock) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetHash

`func (o *NakamotoBlock) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *NakamotoBlock) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *NakamotoBlock) SetHash(v string)`

SetHash sets Hash field to given value.


### GetBlockTime

`func (o *NakamotoBlock) GetBlockTime() float32`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *NakamotoBlock) GetBlockTimeOk() (*float32, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *NakamotoBlock) SetBlockTime(v float32)`

SetBlockTime sets BlockTime field to given value.


### GetBlockTimeIso

`func (o *NakamotoBlock) GetBlockTimeIso() string`

GetBlockTimeIso returns the BlockTimeIso field if non-nil, zero value otherwise.

### GetBlockTimeIsoOk

`func (o *NakamotoBlock) GetBlockTimeIsoOk() (*string, bool)`

GetBlockTimeIsoOk returns a tuple with the BlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTimeIso

`func (o *NakamotoBlock) SetBlockTimeIso(v string)`

SetBlockTimeIso sets BlockTimeIso field to given value.


### GetIndexBlockHash

`func (o *NakamotoBlock) GetIndexBlockHash() string`

GetIndexBlockHash returns the IndexBlockHash field if non-nil, zero value otherwise.

### GetIndexBlockHashOk

`func (o *NakamotoBlock) GetIndexBlockHashOk() (*string, bool)`

GetIndexBlockHashOk returns a tuple with the IndexBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexBlockHash

`func (o *NakamotoBlock) SetIndexBlockHash(v string)`

SetIndexBlockHash sets IndexBlockHash field to given value.


### GetParentBlockHash

`func (o *NakamotoBlock) GetParentBlockHash() string`

GetParentBlockHash returns the ParentBlockHash field if non-nil, zero value otherwise.

### GetParentBlockHashOk

`func (o *NakamotoBlock) GetParentBlockHashOk() (*string, bool)`

GetParentBlockHashOk returns a tuple with the ParentBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBlockHash

`func (o *NakamotoBlock) SetParentBlockHash(v string)`

SetParentBlockHash sets ParentBlockHash field to given value.


### GetParentIndexBlockHash

`func (o *NakamotoBlock) GetParentIndexBlockHash() string`

GetParentIndexBlockHash returns the ParentIndexBlockHash field if non-nil, zero value otherwise.

### GetParentIndexBlockHashOk

`func (o *NakamotoBlock) GetParentIndexBlockHashOk() (*string, bool)`

GetParentIndexBlockHashOk returns a tuple with the ParentIndexBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIndexBlockHash

`func (o *NakamotoBlock) SetParentIndexBlockHash(v string)`

SetParentIndexBlockHash sets ParentIndexBlockHash field to given value.


### GetBurnBlockTime

`func (o *NakamotoBlock) GetBurnBlockTime() float32`

GetBurnBlockTime returns the BurnBlockTime field if non-nil, zero value otherwise.

### GetBurnBlockTimeOk

`func (o *NakamotoBlock) GetBurnBlockTimeOk() (*float32, bool)`

GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTime

`func (o *NakamotoBlock) SetBurnBlockTime(v float32)`

SetBurnBlockTime sets BurnBlockTime field to given value.


### GetBurnBlockTimeIso

`func (o *NakamotoBlock) GetBurnBlockTimeIso() string`

GetBurnBlockTimeIso returns the BurnBlockTimeIso field if non-nil, zero value otherwise.

### GetBurnBlockTimeIsoOk

`func (o *NakamotoBlock) GetBurnBlockTimeIsoOk() (*string, bool)`

GetBurnBlockTimeIsoOk returns a tuple with the BurnBlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTimeIso

`func (o *NakamotoBlock) SetBurnBlockTimeIso(v string)`

SetBurnBlockTimeIso sets BurnBlockTimeIso field to given value.


### GetBurnBlockHash

`func (o *NakamotoBlock) GetBurnBlockHash() string`

GetBurnBlockHash returns the BurnBlockHash field if non-nil, zero value otherwise.

### GetBurnBlockHashOk

`func (o *NakamotoBlock) GetBurnBlockHashOk() (*string, bool)`

GetBurnBlockHashOk returns a tuple with the BurnBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHash

`func (o *NakamotoBlock) SetBurnBlockHash(v string)`

SetBurnBlockHash sets BurnBlockHash field to given value.


### GetBurnBlockHeight

`func (o *NakamotoBlock) GetBurnBlockHeight() int32`

GetBurnBlockHeight returns the BurnBlockHeight field if non-nil, zero value otherwise.

### GetBurnBlockHeightOk

`func (o *NakamotoBlock) GetBurnBlockHeightOk() (*int32, bool)`

GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHeight

`func (o *NakamotoBlock) SetBurnBlockHeight(v int32)`

SetBurnBlockHeight sets BurnBlockHeight field to given value.


### GetMinerTxid

`func (o *NakamotoBlock) GetMinerTxid() string`

GetMinerTxid returns the MinerTxid field if non-nil, zero value otherwise.

### GetMinerTxidOk

`func (o *NakamotoBlock) GetMinerTxidOk() (*string, bool)`

GetMinerTxidOk returns a tuple with the MinerTxid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinerTxid

`func (o *NakamotoBlock) SetMinerTxid(v string)`

SetMinerTxid sets MinerTxid field to given value.


### GetTxCount

`func (o *NakamotoBlock) GetTxCount() int32`

GetTxCount returns the TxCount field if non-nil, zero value otherwise.

### GetTxCountOk

`func (o *NakamotoBlock) GetTxCountOk() (*int32, bool)`

GetTxCountOk returns a tuple with the TxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxCount

`func (o *NakamotoBlock) SetTxCount(v int32)`

SetTxCount sets TxCount field to given value.


### GetExecutionCostReadCount

`func (o *NakamotoBlock) GetExecutionCostReadCount() int32`

GetExecutionCostReadCount returns the ExecutionCostReadCount field if non-nil, zero value otherwise.

### GetExecutionCostReadCountOk

`func (o *NakamotoBlock) GetExecutionCostReadCountOk() (*int32, bool)`

GetExecutionCostReadCountOk returns a tuple with the ExecutionCostReadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostReadCount

`func (o *NakamotoBlock) SetExecutionCostReadCount(v int32)`

SetExecutionCostReadCount sets ExecutionCostReadCount field to given value.


### GetExecutionCostReadLength

`func (o *NakamotoBlock) GetExecutionCostReadLength() int32`

GetExecutionCostReadLength returns the ExecutionCostReadLength field if non-nil, zero value otherwise.

### GetExecutionCostReadLengthOk

`func (o *NakamotoBlock) GetExecutionCostReadLengthOk() (*int32, bool)`

GetExecutionCostReadLengthOk returns a tuple with the ExecutionCostReadLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostReadLength

`func (o *NakamotoBlock) SetExecutionCostReadLength(v int32)`

SetExecutionCostReadLength sets ExecutionCostReadLength field to given value.


### GetExecutionCostRuntime

`func (o *NakamotoBlock) GetExecutionCostRuntime() int32`

GetExecutionCostRuntime returns the ExecutionCostRuntime field if non-nil, zero value otherwise.

### GetExecutionCostRuntimeOk

`func (o *NakamotoBlock) GetExecutionCostRuntimeOk() (*int32, bool)`

GetExecutionCostRuntimeOk returns a tuple with the ExecutionCostRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostRuntime

`func (o *NakamotoBlock) SetExecutionCostRuntime(v int32)`

SetExecutionCostRuntime sets ExecutionCostRuntime field to given value.


### GetExecutionCostWriteCount

`func (o *NakamotoBlock) GetExecutionCostWriteCount() int32`

GetExecutionCostWriteCount returns the ExecutionCostWriteCount field if non-nil, zero value otherwise.

### GetExecutionCostWriteCountOk

`func (o *NakamotoBlock) GetExecutionCostWriteCountOk() (*int32, bool)`

GetExecutionCostWriteCountOk returns a tuple with the ExecutionCostWriteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostWriteCount

`func (o *NakamotoBlock) SetExecutionCostWriteCount(v int32)`

SetExecutionCostWriteCount sets ExecutionCostWriteCount field to given value.


### GetExecutionCostWriteLength

`func (o *NakamotoBlock) GetExecutionCostWriteLength() int32`

GetExecutionCostWriteLength returns the ExecutionCostWriteLength field if non-nil, zero value otherwise.

### GetExecutionCostWriteLengthOk

`func (o *NakamotoBlock) GetExecutionCostWriteLengthOk() (*int32, bool)`

GetExecutionCostWriteLengthOk returns a tuple with the ExecutionCostWriteLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostWriteLength

`func (o *NakamotoBlock) SetExecutionCostWriteLength(v int32)`

SetExecutionCostWriteLength sets ExecutionCostWriteLength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


