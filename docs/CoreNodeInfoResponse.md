# CoreNodeInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeerVersion** | **int32** | identifies the version number for the networking communication, this should not change while a node is running, and will only change if there&#39;s an upgrade | 
**PoxConsensus** | **string** | is a hash used to identify the burnchain view for a node. it incorporates bitcoin chain information and PoX information. nodes that disagree on this value will appear to each other as forks. this value will change after every block | 
**BurnBlockHeight** | **int32** | latest bitcoin chain height | 
**StablePoxConsensus** | **string** | same as burn_consensus, but evaluated at stable_burn_block_height | 
**StableBurnBlockHeight** | **int32** | leftover from stacks 1.0, basically always burn_block_height - 1 | 
**ServerVersion** | **string** | is a version descriptor | 
**NetworkId** | **int32** | is similar to peer_version and will be used to differentiate between different testnets. this value will be different between mainnet and testnet. once launched, this value will not change | 
**ParentNetworkId** | **int32** | same as network_id, but for bitcoin | 
**StacksTipHeight** | **int32** | the latest Stacks chain height. Stacks forks can occur independent of the Bitcoin chain, that height doesn&#39;t increase 1-to-1 with the Bitcoin height | 
**StacksTip** | **string** | the best known block hash for the Stack chain (not including any pending microblocks) | 
**StacksTipConsensusHash** | **string** | the burn chain (i.e., bitcoin) consensus hash at the time that stacks_tip was mined | 
**UnanchoredTip** | **string** | the latest microblock hash if any microblocks were processed. if no microblock has been processed for the current block, a 000.., hex array is returned | 
**ExitAtBlockHeight** | **int32** | the block height at which the testnet network will be reset. not applicable for mainnet | 

## Methods

### NewCoreNodeInfoResponse

`func NewCoreNodeInfoResponse(peerVersion int32, poxConsensus string, burnBlockHeight int32, stablePoxConsensus string, stableBurnBlockHeight int32, serverVersion string, networkId int32, parentNetworkId int32, stacksTipHeight int32, stacksTip string, stacksTipConsensusHash string, unanchoredTip string, exitAtBlockHeight int32, ) *CoreNodeInfoResponse`

NewCoreNodeInfoResponse instantiates a new CoreNodeInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreNodeInfoResponseWithDefaults

`func NewCoreNodeInfoResponseWithDefaults() *CoreNodeInfoResponse`

NewCoreNodeInfoResponseWithDefaults instantiates a new CoreNodeInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeerVersion

`func (o *CoreNodeInfoResponse) GetPeerVersion() int32`

GetPeerVersion returns the PeerVersion field if non-nil, zero value otherwise.

### GetPeerVersionOk

`func (o *CoreNodeInfoResponse) GetPeerVersionOk() (*int32, bool)`

GetPeerVersionOk returns a tuple with the PeerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerVersion

`func (o *CoreNodeInfoResponse) SetPeerVersion(v int32)`

SetPeerVersion sets PeerVersion field to given value.


### GetPoxConsensus

`func (o *CoreNodeInfoResponse) GetPoxConsensus() string`

GetPoxConsensus returns the PoxConsensus field if non-nil, zero value otherwise.

### GetPoxConsensusOk

`func (o *CoreNodeInfoResponse) GetPoxConsensusOk() (*string, bool)`

GetPoxConsensusOk returns a tuple with the PoxConsensus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoxConsensus

`func (o *CoreNodeInfoResponse) SetPoxConsensus(v string)`

SetPoxConsensus sets PoxConsensus field to given value.


### GetBurnBlockHeight

`func (o *CoreNodeInfoResponse) GetBurnBlockHeight() int32`

GetBurnBlockHeight returns the BurnBlockHeight field if non-nil, zero value otherwise.

### GetBurnBlockHeightOk

`func (o *CoreNodeInfoResponse) GetBurnBlockHeightOk() (*int32, bool)`

GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHeight

`func (o *CoreNodeInfoResponse) SetBurnBlockHeight(v int32)`

SetBurnBlockHeight sets BurnBlockHeight field to given value.


### GetStablePoxConsensus

`func (o *CoreNodeInfoResponse) GetStablePoxConsensus() string`

GetStablePoxConsensus returns the StablePoxConsensus field if non-nil, zero value otherwise.

### GetStablePoxConsensusOk

`func (o *CoreNodeInfoResponse) GetStablePoxConsensusOk() (*string, bool)`

GetStablePoxConsensusOk returns a tuple with the StablePoxConsensus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStablePoxConsensus

`func (o *CoreNodeInfoResponse) SetStablePoxConsensus(v string)`

SetStablePoxConsensus sets StablePoxConsensus field to given value.


### GetStableBurnBlockHeight

`func (o *CoreNodeInfoResponse) GetStableBurnBlockHeight() int32`

GetStableBurnBlockHeight returns the StableBurnBlockHeight field if non-nil, zero value otherwise.

### GetStableBurnBlockHeightOk

`func (o *CoreNodeInfoResponse) GetStableBurnBlockHeightOk() (*int32, bool)`

GetStableBurnBlockHeightOk returns a tuple with the StableBurnBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStableBurnBlockHeight

`func (o *CoreNodeInfoResponse) SetStableBurnBlockHeight(v int32)`

SetStableBurnBlockHeight sets StableBurnBlockHeight field to given value.


### GetServerVersion

`func (o *CoreNodeInfoResponse) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *CoreNodeInfoResponse) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *CoreNodeInfoResponse) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.


### GetNetworkId

`func (o *CoreNodeInfoResponse) GetNetworkId() int32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *CoreNodeInfoResponse) GetNetworkIdOk() (*int32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *CoreNodeInfoResponse) SetNetworkId(v int32)`

SetNetworkId sets NetworkId field to given value.


### GetParentNetworkId

`func (o *CoreNodeInfoResponse) GetParentNetworkId() int32`

GetParentNetworkId returns the ParentNetworkId field if non-nil, zero value otherwise.

### GetParentNetworkIdOk

`func (o *CoreNodeInfoResponse) GetParentNetworkIdOk() (*int32, bool)`

GetParentNetworkIdOk returns a tuple with the ParentNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkId

`func (o *CoreNodeInfoResponse) SetParentNetworkId(v int32)`

SetParentNetworkId sets ParentNetworkId field to given value.


### GetStacksTipHeight

`func (o *CoreNodeInfoResponse) GetStacksTipHeight() int32`

GetStacksTipHeight returns the StacksTipHeight field if non-nil, zero value otherwise.

### GetStacksTipHeightOk

`func (o *CoreNodeInfoResponse) GetStacksTipHeightOk() (*int32, bool)`

GetStacksTipHeightOk returns a tuple with the StacksTipHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacksTipHeight

`func (o *CoreNodeInfoResponse) SetStacksTipHeight(v int32)`

SetStacksTipHeight sets StacksTipHeight field to given value.


### GetStacksTip

`func (o *CoreNodeInfoResponse) GetStacksTip() string`

GetStacksTip returns the StacksTip field if non-nil, zero value otherwise.

### GetStacksTipOk

`func (o *CoreNodeInfoResponse) GetStacksTipOk() (*string, bool)`

GetStacksTipOk returns a tuple with the StacksTip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacksTip

`func (o *CoreNodeInfoResponse) SetStacksTip(v string)`

SetStacksTip sets StacksTip field to given value.


### GetStacksTipConsensusHash

`func (o *CoreNodeInfoResponse) GetStacksTipConsensusHash() string`

GetStacksTipConsensusHash returns the StacksTipConsensusHash field if non-nil, zero value otherwise.

### GetStacksTipConsensusHashOk

`func (o *CoreNodeInfoResponse) GetStacksTipConsensusHashOk() (*string, bool)`

GetStacksTipConsensusHashOk returns a tuple with the StacksTipConsensusHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacksTipConsensusHash

`func (o *CoreNodeInfoResponse) SetStacksTipConsensusHash(v string)`

SetStacksTipConsensusHash sets StacksTipConsensusHash field to given value.


### GetUnanchoredTip

`func (o *CoreNodeInfoResponse) GetUnanchoredTip() string`

GetUnanchoredTip returns the UnanchoredTip field if non-nil, zero value otherwise.

### GetUnanchoredTipOk

`func (o *CoreNodeInfoResponse) GetUnanchoredTipOk() (*string, bool)`

GetUnanchoredTipOk returns a tuple with the UnanchoredTip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnanchoredTip

`func (o *CoreNodeInfoResponse) SetUnanchoredTip(v string)`

SetUnanchoredTip sets UnanchoredTip field to given value.


### GetExitAtBlockHeight

`func (o *CoreNodeInfoResponse) GetExitAtBlockHeight() int32`

GetExitAtBlockHeight returns the ExitAtBlockHeight field if non-nil, zero value otherwise.

### GetExitAtBlockHeightOk

`func (o *CoreNodeInfoResponse) GetExitAtBlockHeightOk() (*int32, bool)`

GetExitAtBlockHeightOk returns a tuple with the ExitAtBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitAtBlockHeight

`func (o *CoreNodeInfoResponse) SetExitAtBlockHeight(v int32)`

SetExitAtBlockHeight sets ExitAtBlockHeight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


