# RosettaConstructionMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**RosettaConstructionMetadataResponseMetadata**](RosettaConstructionMetadataResponseMetadata.md) |  | 
**SuggestedFee** | Pointer to [**[]RosettaAmount**](RosettaAmount.md) |  | [optional] 

## Methods

### NewRosettaConstructionMetadataResponse

`func NewRosettaConstructionMetadataResponse(metadata RosettaConstructionMetadataResponseMetadata, ) *RosettaConstructionMetadataResponse`

NewRosettaConstructionMetadataResponse instantiates a new RosettaConstructionMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaConstructionMetadataResponseWithDefaults

`func NewRosettaConstructionMetadataResponseWithDefaults() *RosettaConstructionMetadataResponse`

NewRosettaConstructionMetadataResponseWithDefaults instantiates a new RosettaConstructionMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *RosettaConstructionMetadataResponse) GetMetadata() RosettaConstructionMetadataResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RosettaConstructionMetadataResponse) GetMetadataOk() (*RosettaConstructionMetadataResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RosettaConstructionMetadataResponse) SetMetadata(v RosettaConstructionMetadataResponseMetadata)`

SetMetadata sets Metadata field to given value.


### GetSuggestedFee

`func (o *RosettaConstructionMetadataResponse) GetSuggestedFee() []RosettaAmount`

GetSuggestedFee returns the SuggestedFee field if non-nil, zero value otherwise.

### GetSuggestedFeeOk

`func (o *RosettaConstructionMetadataResponse) GetSuggestedFeeOk() (*[]RosettaAmount, bool)`

GetSuggestedFeeOk returns a tuple with the SuggestedFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedFee

`func (o *RosettaConstructionMetadataResponse) SetSuggestedFee(v []RosettaAmount)`

SetSuggestedFee sets SuggestedFee field to given value.

### HasSuggestedFee

`func (o *RosettaConstructionMetadataResponse) HasSuggestedFee() bool`

HasSuggestedFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


