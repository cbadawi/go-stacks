# TenureChangeTransactionMetadataTenureChangePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenureConsensusHash** | **string** | Consensus hash of this tenure. Corresponds to the sortition in which the miner of this block was chosen. | 
**PrevTenureConsensusHash** | **string** | Consensus hash of the previous tenure. Corresponds to the sortition of the previous winning block-commit. | 
**BurnViewConsensusHash** | **string** | Current consensus hash on the underlying burnchain. Corresponds to the last-seen sortition. | 
**PreviousTenureEnd** | **string** | (Hex string) Stacks Block hash | 
**PreviousTenureBlocks** | **float32** | The number of blocks produced in the previous tenure. | 
**Cause** | **string** | Cause of change in mining tenure. Depending on cause, tenure can be ended or extended. | 
**PubkeyHash** | **string** | (Hex string) The ECDSA public key hash of the current tenure. | 

## Methods

### NewTenureChangeTransactionMetadataTenureChangePayload

`func NewTenureChangeTransactionMetadataTenureChangePayload(tenureConsensusHash string, prevTenureConsensusHash string, burnViewConsensusHash string, previousTenureEnd string, previousTenureBlocks float32, cause string, pubkeyHash string, ) *TenureChangeTransactionMetadataTenureChangePayload`

NewTenureChangeTransactionMetadataTenureChangePayload instantiates a new TenureChangeTransactionMetadataTenureChangePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenureChangeTransactionMetadataTenureChangePayloadWithDefaults

`func NewTenureChangeTransactionMetadataTenureChangePayloadWithDefaults() *TenureChangeTransactionMetadataTenureChangePayload`

NewTenureChangeTransactionMetadataTenureChangePayloadWithDefaults instantiates a new TenureChangeTransactionMetadataTenureChangePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenureConsensusHash

`func (o *TenureChangeTransactionMetadataTenureChangePayload) GetTenureConsensusHash() string`

GetTenureConsensusHash returns the TenureConsensusHash field if non-nil, zero value otherwise.

### GetTenureConsensusHashOk

`func (o *TenureChangeTransactionMetadataTenureChangePayload) GetTenureConsensusHashOk() (*string, bool)`

GetTenureConsensusHashOk returns a tuple with the TenureConsensusHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenureConsensusHash

`func (o *TenureChangeTransactionMetadataTenureChangePayload) SetTenureConsensusHash(v string)`

SetTenureConsensusHash sets TenureConsensusHash field to given value.


### GetPrevTenureConsensusHash

`func (o *TenureChangeTransactionMetadataTenureChangePayload) GetPrevTenureConsensusHash() string`

GetPrevTenureConsensusHash returns the PrevTenureConsensusHash field if non-nil, zero value otherwise.

### GetPrevTenureConsensusHashOk

`func (o *TenureChangeTransactionMetadataTenureChangePayload) GetPrevTenureConsensusHashOk() (*string, bool)`

GetPrevTenureConsensusHashOk returns a tuple with the PrevTenureConsensusHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevTenureConsensusHash

`func (o *TenureChangeTransactionMetadataTenureChangePayload) SetPrevTenureConsensusHash(v string)`

SetPrevTenureConsensusHash sets PrevTenureConsensusHash field to given value.


### GetBurnViewConsensusHash

`func (o *TenureChangeTransactionMetadataTenureChangePayload) GetBurnViewConsensusHash() string`

GetBurnViewConsensusHash returns the BurnViewConsensusHash field if non-nil, zero value otherwise.

### GetBurnViewConsensusHashOk

`func (o *TenureChangeTransactionMetadataTenureChangePayload) GetBurnViewConsensusHashOk() (*string, bool)`

GetBurnViewConsensusHashOk returns a tuple with the BurnViewConsensusHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnViewConsensusHash

`func (o *TenureChangeTransactionMetadataTenureChangePayload) SetBurnViewConsensusHash(v string)`

SetBurnViewConsensusHash sets BurnViewConsensusHash field to given value.


### GetPreviousTenureEnd

`func (o *TenureChangeTransactionMetadataTenureChangePayload) GetPreviousTenureEnd() string`

GetPreviousTenureEnd returns the PreviousTenureEnd field if non-nil, zero value otherwise.

### GetPreviousTenureEndOk

`func (o *TenureChangeTransactionMetadataTenureChangePayload) GetPreviousTenureEndOk() (*string, bool)`

GetPreviousTenureEndOk returns a tuple with the PreviousTenureEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTenureEnd

`func (o *TenureChangeTransactionMetadataTenureChangePayload) SetPreviousTenureEnd(v string)`

SetPreviousTenureEnd sets PreviousTenureEnd field to given value.


### GetPreviousTenureBlocks

`func (o *TenureChangeTransactionMetadataTenureChangePayload) GetPreviousTenureBlocks() float32`

GetPreviousTenureBlocks returns the PreviousTenureBlocks field if non-nil, zero value otherwise.

### GetPreviousTenureBlocksOk

`func (o *TenureChangeTransactionMetadataTenureChangePayload) GetPreviousTenureBlocksOk() (*float32, bool)`

GetPreviousTenureBlocksOk returns a tuple with the PreviousTenureBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTenureBlocks

`func (o *TenureChangeTransactionMetadataTenureChangePayload) SetPreviousTenureBlocks(v float32)`

SetPreviousTenureBlocks sets PreviousTenureBlocks field to given value.


### GetCause

`func (o *TenureChangeTransactionMetadataTenureChangePayload) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *TenureChangeTransactionMetadataTenureChangePayload) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *TenureChangeTransactionMetadataTenureChangePayload) SetCause(v string)`

SetCause sets Cause field to given value.


### GetPubkeyHash

`func (o *TenureChangeTransactionMetadataTenureChangePayload) GetPubkeyHash() string`

GetPubkeyHash returns the PubkeyHash field if non-nil, zero value otherwise.

### GetPubkeyHashOk

`func (o *TenureChangeTransactionMetadataTenureChangePayload) GetPubkeyHashOk() (*string, bool)`

GetPubkeyHashOk returns a tuple with the PubkeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkeyHash

`func (o *TenureChangeTransactionMetadataTenureChangePayload) SetPubkeyHash(v string)`

SetPubkeyHash sets PubkeyHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


