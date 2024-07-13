# RosettaNetworkStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentBlockIdentifier** | [**RosettaBlockIdentifier**](RosettaBlockIdentifier.md) |  | 
**CurrentBlockTimestamp** | **int32** | The timestamp of the block in milliseconds since the Unix Epoch. The timestamp is stored in milliseconds because some blockchains produce blocks more often than once a second. | 
**GenesisBlockIdentifier** | [**RosettaGenesisBlockIdentifier**](RosettaGenesisBlockIdentifier.md) |  | 
**OldestBlockIdentifier** | Pointer to [**RosettaOldestBlockIdentifier**](RosettaOldestBlockIdentifier.md) |  | [optional] 
**SyncStatus** | Pointer to [**RosettaSyncStatus**](RosettaSyncStatus.md) |  | [optional] 
**Peers** | [**[]RosettaPeers**](RosettaPeers.md) | Peers information | 
**CurrentBurnBlockHeight** | **int32** | The latest burn block height | 

## Methods

### NewRosettaNetworkStatusResponse

`func NewRosettaNetworkStatusResponse(currentBlockIdentifier RosettaBlockIdentifier, currentBlockTimestamp int32, genesisBlockIdentifier RosettaGenesisBlockIdentifier, peers []RosettaPeers, currentBurnBlockHeight int32, ) *RosettaNetworkStatusResponse`

NewRosettaNetworkStatusResponse instantiates a new RosettaNetworkStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaNetworkStatusResponseWithDefaults

`func NewRosettaNetworkStatusResponseWithDefaults() *RosettaNetworkStatusResponse`

NewRosettaNetworkStatusResponseWithDefaults instantiates a new RosettaNetworkStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentBlockIdentifier

`func (o *RosettaNetworkStatusResponse) GetCurrentBlockIdentifier() RosettaBlockIdentifier`

GetCurrentBlockIdentifier returns the CurrentBlockIdentifier field if non-nil, zero value otherwise.

### GetCurrentBlockIdentifierOk

`func (o *RosettaNetworkStatusResponse) GetCurrentBlockIdentifierOk() (*RosettaBlockIdentifier, bool)`

GetCurrentBlockIdentifierOk returns a tuple with the CurrentBlockIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBlockIdentifier

`func (o *RosettaNetworkStatusResponse) SetCurrentBlockIdentifier(v RosettaBlockIdentifier)`

SetCurrentBlockIdentifier sets CurrentBlockIdentifier field to given value.


### GetCurrentBlockTimestamp

`func (o *RosettaNetworkStatusResponse) GetCurrentBlockTimestamp() int32`

GetCurrentBlockTimestamp returns the CurrentBlockTimestamp field if non-nil, zero value otherwise.

### GetCurrentBlockTimestampOk

`func (o *RosettaNetworkStatusResponse) GetCurrentBlockTimestampOk() (*int32, bool)`

GetCurrentBlockTimestampOk returns a tuple with the CurrentBlockTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBlockTimestamp

`func (o *RosettaNetworkStatusResponse) SetCurrentBlockTimestamp(v int32)`

SetCurrentBlockTimestamp sets CurrentBlockTimestamp field to given value.


### GetGenesisBlockIdentifier

`func (o *RosettaNetworkStatusResponse) GetGenesisBlockIdentifier() RosettaGenesisBlockIdentifier`

GetGenesisBlockIdentifier returns the GenesisBlockIdentifier field if non-nil, zero value otherwise.

### GetGenesisBlockIdentifierOk

`func (o *RosettaNetworkStatusResponse) GetGenesisBlockIdentifierOk() (*RosettaGenesisBlockIdentifier, bool)`

GetGenesisBlockIdentifierOk returns a tuple with the GenesisBlockIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenesisBlockIdentifier

`func (o *RosettaNetworkStatusResponse) SetGenesisBlockIdentifier(v RosettaGenesisBlockIdentifier)`

SetGenesisBlockIdentifier sets GenesisBlockIdentifier field to given value.


### GetOldestBlockIdentifier

`func (o *RosettaNetworkStatusResponse) GetOldestBlockIdentifier() RosettaOldestBlockIdentifier`

GetOldestBlockIdentifier returns the OldestBlockIdentifier field if non-nil, zero value otherwise.

### GetOldestBlockIdentifierOk

`func (o *RosettaNetworkStatusResponse) GetOldestBlockIdentifierOk() (*RosettaOldestBlockIdentifier, bool)`

GetOldestBlockIdentifierOk returns a tuple with the OldestBlockIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestBlockIdentifier

`func (o *RosettaNetworkStatusResponse) SetOldestBlockIdentifier(v RosettaOldestBlockIdentifier)`

SetOldestBlockIdentifier sets OldestBlockIdentifier field to given value.

### HasOldestBlockIdentifier

`func (o *RosettaNetworkStatusResponse) HasOldestBlockIdentifier() bool`

HasOldestBlockIdentifier returns a boolean if a field has been set.

### GetSyncStatus

`func (o *RosettaNetworkStatusResponse) GetSyncStatus() RosettaSyncStatus`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *RosettaNetworkStatusResponse) GetSyncStatusOk() (*RosettaSyncStatus, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *RosettaNetworkStatusResponse) SetSyncStatus(v RosettaSyncStatus)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *RosettaNetworkStatusResponse) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.

### GetPeers

`func (o *RosettaNetworkStatusResponse) GetPeers() []RosettaPeers`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *RosettaNetworkStatusResponse) GetPeersOk() (*[]RosettaPeers, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *RosettaNetworkStatusResponse) SetPeers(v []RosettaPeers)`

SetPeers sets Peers field to given value.


### GetCurrentBurnBlockHeight

`func (o *RosettaNetworkStatusResponse) GetCurrentBurnBlockHeight() int32`

GetCurrentBurnBlockHeight returns the CurrentBurnBlockHeight field if non-nil, zero value otherwise.

### GetCurrentBurnBlockHeightOk

`func (o *RosettaNetworkStatusResponse) GetCurrentBurnBlockHeightOk() (*int32, bool)`

GetCurrentBurnBlockHeightOk returns a tuple with the CurrentBurnBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBurnBlockHeight

`func (o *RosettaNetworkStatusResponse) SetCurrentBurnBlockHeight(v int32)`

SetCurrentBurnBlockHeight sets CurrentBurnBlockHeight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


