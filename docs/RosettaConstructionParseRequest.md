# RosettaConstructionParseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkIdentifier** | [**NetworkIdentifier**](NetworkIdentifier.md) |  | 
**Signed** | **bool** | Signed is a boolean indicating whether the transaction is signed. | 
**Transaction** | **string** | This must be either the unsigned transaction blob returned by /construction/payloads or the signed transaction blob returned by /construction/combine. | 

## Methods

### NewRosettaConstructionParseRequest

`func NewRosettaConstructionParseRequest(networkIdentifier NetworkIdentifier, signed bool, transaction string, ) *RosettaConstructionParseRequest`

NewRosettaConstructionParseRequest instantiates a new RosettaConstructionParseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaConstructionParseRequestWithDefaults

`func NewRosettaConstructionParseRequestWithDefaults() *RosettaConstructionParseRequest`

NewRosettaConstructionParseRequestWithDefaults instantiates a new RosettaConstructionParseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkIdentifier

`func (o *RosettaConstructionParseRequest) GetNetworkIdentifier() NetworkIdentifier`

GetNetworkIdentifier returns the NetworkIdentifier field if non-nil, zero value otherwise.

### GetNetworkIdentifierOk

`func (o *RosettaConstructionParseRequest) GetNetworkIdentifierOk() (*NetworkIdentifier, bool)`

GetNetworkIdentifierOk returns a tuple with the NetworkIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIdentifier

`func (o *RosettaConstructionParseRequest) SetNetworkIdentifier(v NetworkIdentifier)`

SetNetworkIdentifier sets NetworkIdentifier field to given value.


### GetSigned

`func (o *RosettaConstructionParseRequest) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *RosettaConstructionParseRequest) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *RosettaConstructionParseRequest) SetSigned(v bool)`

SetSigned sets Signed field to given value.


### GetTransaction

`func (o *RosettaConstructionParseRequest) GetTransaction() string`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *RosettaConstructionParseRequest) GetTransactionOk() (*string, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *RosettaConstructionParseRequest) SetTransaction(v string)`

SetTransaction sets Transaction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


