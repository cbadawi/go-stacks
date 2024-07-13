# RosettaConstructionCombineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkIdentifier** | [**NetworkIdentifier**](NetworkIdentifier.md) |  | 
**UnsignedTransaction** | **string** |  | 
**Signatures** | [**[]RosettaSignature**](RosettaSignature.md) |  | 

## Methods

### NewRosettaConstructionCombineRequest

`func NewRosettaConstructionCombineRequest(networkIdentifier NetworkIdentifier, unsignedTransaction string, signatures []RosettaSignature, ) *RosettaConstructionCombineRequest`

NewRosettaConstructionCombineRequest instantiates a new RosettaConstructionCombineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaConstructionCombineRequestWithDefaults

`func NewRosettaConstructionCombineRequestWithDefaults() *RosettaConstructionCombineRequest`

NewRosettaConstructionCombineRequestWithDefaults instantiates a new RosettaConstructionCombineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkIdentifier

`func (o *RosettaConstructionCombineRequest) GetNetworkIdentifier() NetworkIdentifier`

GetNetworkIdentifier returns the NetworkIdentifier field if non-nil, zero value otherwise.

### GetNetworkIdentifierOk

`func (o *RosettaConstructionCombineRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool)`

GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIdentifier

`func (o *RosettaConstructionCombineRequest) SetNetworkIdentifier(v NetworkIdentifier)`

SetNetworkIdentifier sets NetworkIdentifier field to given value.


### GetUnsignedTransaction

`func (o *RosettaConstructionCombineRequest) GetUnsignedTransaction() string`

GetUnsignedTransaction returns the UnsignedTransaction field if non-nil, zero value otherwise.

### GetUnsignedTransactionOk

`func (o *RosettaConstructionCombineRequest) GetUnsignedTransactionOk() (*string, bool)`

GetUnsignedTransactionOk returns a tuple with the UnsignedTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsignedTransaction

`func (o *RosettaConstructionCombineRequest) SetUnsignedTransaction(v string)`

SetUnsignedTransaction sets UnsignedTransaction field to given value.


### GetSignatures

`func (o *RosettaConstructionCombineRequest) GetSignatures() []RosettaSignature`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *RosettaConstructionCombineRequest) GetSignaturesOk() (*[]RosettaSignature, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *RosettaConstructionCombineRequest) SetSignatures(v []RosettaSignature)`

SetSignatures sets Signatures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


