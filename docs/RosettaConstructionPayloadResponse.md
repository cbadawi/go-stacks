# RosettaConstructionPayloadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnsignedTransaction** | **string** | This is an unsigned transaction blob (that is usually needed to construct the a network transaction from a collection of signatures) | 
**Payloads** | [**[]SigningPayload**](SigningPayload.md) | An array of payloads that must be signed by the caller | 

## Methods

### NewRosettaConstructionPayloadResponse

`func NewRosettaConstructionPayloadResponse(unsignedTransaction string, payloads []SigningPayload, ) *RosettaConstructionPayloadResponse`

NewRosettaConstructionPayloadResponse instantiates a new RosettaConstructionPayloadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaConstructionPayloadResponseWithDefaults

`func NewRosettaConstructionPayloadResponseWithDefaults() *RosettaConstructionPayloadResponse`

NewRosettaConstructionPayloadResponseWithDefaults instantiates a new RosettaConstructionPayloadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnsignedTransaction

`func (o *RosettaConstructionPayloadResponse) GetUnsignedTransaction() string`

GetUnsignedTransaction returns the UnsignedTransaction field if non-nil, zero value otherwise.

### GetUnsignedTransactionOk

`func (o *RosettaConstructionPayloadResponse) GetUnsignedTransactionOk() (*string, bool)`

GetUnsignedTransactionOk returns a tuple with the UnsignedTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsignedTransaction

`func (o *RosettaConstructionPayloadResponse) SetUnsignedTransaction(v string)`

SetUnsignedTransaction sets UnsignedTransaction field to given value.


### GetPayloads

`func (o *RosettaConstructionPayloadResponse) GetPayloads() []SigningPayload`

GetPayloads returns the Payloads field if non-nil, zero value otherwise.

### GetPayloadsOk

`func (o *RosettaConstructionPayloadResponse) GetPayloadsOk() (*[]SigningPayload, bool)`

GetPayloadsOk returns a tuple with the Payloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloads

`func (o *RosettaConstructionPayloadResponse) SetPayloads(v []SigningPayload)`

SetPayloads sets Payloads field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


