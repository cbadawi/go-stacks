# RosettaConstructionPreprocessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkIdentifier** | [**NetworkIdentifier**](NetworkIdentifier.md) |  | 
**Operations** | [**[]RosettaOperation**](RosettaOperation.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**MaxFee** | Pointer to [**[]RosettaMaxFeeAmount**](RosettaMaxFeeAmount.md) |  | [optional] 
**SuggestedFeeMultiplier** | Pointer to **int32** |  The caller can also provide a suggested fee multiplier to indicate that the suggested fee should be scaled. This may be used to set higher fees for urgent transactions or to pay lower fees when there is less urgency. It is assumed that providing a very low multiplier (like 0.0001) will never lead to a transaction being created with a fee less than the minimum network fee (if applicable). In the case that the caller provides both a max fee and a suggested fee multiplier, the max fee will set an upper bound on the suggested fee (regardless of the multiplier provided). | [optional] 

## Methods

### NewRosettaConstructionPreprocessRequest

`func NewRosettaConstructionPreprocessRequest(networkIdentifier NetworkIdentifier, operations []RosettaOperation, ) *RosettaConstructionPreprocessRequest`

NewRosettaConstructionPreprocessRequest instantiates a new RosettaConstructionPreprocessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaConstructionPreprocessRequestWithDefaults

`func NewRosettaConstructionPreprocessRequestWithDefaults() *RosettaConstructionPreprocessRequest`

NewRosettaConstructionPreprocessRequestWithDefaults instantiates a new RosettaConstructionPreprocessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkIdentifier

`func (o *RosettaConstructionPreprocessRequest) GetNetworkIdentifier() NetworkIdentifier`

GetNetworkIdentifier returns the NetworkIdentifier field if non-nil, zero value otherwise.

### GetNetworkIdentifierOk

`func (o *RosettaConstructionPreprocessRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool)`

GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIdentifier

`func (o *RosettaConstructionPreprocessRequest) SetNetworkIdentifier(v NetworkIdentifier)`

SetNetworkIdentifier sets NetworkIdentifier field to given value.


### GetOperations

`func (o *RosettaConstructionPreprocessRequest) GetOperations() []RosettaOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *RosettaConstructionPreprocessRequest) GetOperationsOk() (*[]RosettaOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *RosettaConstructionPreprocessRequest) SetOperations(v []RosettaOperation)`

SetOperations sets Operations field to given value.


### GetMetadata

`func (o *RosettaConstructionPreprocessRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RosettaConstructionPreprocessRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RosettaConstructionPreprocessRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RosettaConstructionPreprocessRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMaxFee

`func (o *RosettaConstructionPreprocessRequest) GetMaxFee() []RosettaMaxFeeAmount`

GetMaxFee returns the MaxFee field if non-nil, zero value otherwise.

### GetMaxFeeOk

`func (o *RosettaConstructionPreprocessRequest) GetMaxFeeOk() (*[]RosettaMaxFeeAmount, bool)`

GetMaxFeeOk returns a tuple with the MaxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFee

`func (o *RosettaConstructionPreprocessRequest) SetMaxFee(v []RosettaMaxFeeAmount)`

SetMaxFee sets MaxFee field to given value.

### HasMaxFee

`func (o *RosettaConstructionPreprocessRequest) HasMaxFee() bool`

HasMaxFee returns a boolean if a field has been set.

### GetSuggestedFeeMultiplier

`func (o *RosettaConstructionPreprocessRequest) GetSuggestedFeeMultiplier() int32`

GetSuggestedFeeMultiplier returns the SuggestedFeeMultiplier field if non-nil, zero value otherwise.

### GetSuggestedFeeMultiplierOk

`func (o *RosettaConstructionPreprocessRequest) GetSuggestedFeeMultiplierOk() (*int32, bool)`

GetSuggestedFeeMultiplierOk returns a tuple with the SuggestedFeeMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedFeeMultiplier

`func (o *RosettaConstructionPreprocessRequest) SetSuggestedFeeMultiplier(v int32)`

SetSuggestedFeeMultiplier sets SuggestedFeeMultiplier field to given value.

### HasSuggestedFeeMultiplier

`func (o *RosettaConstructionPreprocessRequest) HasSuggestedFeeMultiplier() bool`

HasSuggestedFeeMultiplier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


