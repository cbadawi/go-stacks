# ContractSourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  | 
**PublishHeight** | **int32** |  | 
**Proof** | **string** |  | 

## Methods

### NewContractSourceResponse

`func NewContractSourceResponse(source string, publishHeight int32, proof string, ) *ContractSourceResponse`

NewContractSourceResponse instantiates a new ContractSourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractSourceResponseWithDefaults

`func NewContractSourceResponseWithDefaults() *ContractSourceResponse`

NewContractSourceResponseWithDefaults instantiates a new ContractSourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ContractSourceResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ContractSourceResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ContractSourceResponse) SetSource(v string)`

SetSource sets Source field to given value.


### GetPublishHeight

`func (o *ContractSourceResponse) GetPublishHeight() int32`

GetPublishHeight returns the PublishHeight field if non-nil, zero value otherwise.

### GetPublishHeightOk

`func (o *ContractSourceResponse) GetPublishHeightOk() (*int32, bool)`

GetPublishHeightOk returns a tuple with the PublishHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishHeight

`func (o *ContractSourceResponse) SetPublishHeight(v int32)`

SetPublishHeight sets PublishHeight field to given value.


### GetProof

`func (o *ContractSourceResponse) GetProof() string`

GetProof returns the Proof field if non-nil, zero value otherwise.

### GetProofOk

`func (o *ContractSourceResponse) GetProofOk() (*string, bool)`

GetProofOk returns a tuple with the Proof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProof

`func (o *ContractSourceResponse) SetProof(v string)`

SetProof sets Proof field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


