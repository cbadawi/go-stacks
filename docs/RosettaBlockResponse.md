# RosettaBlockResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Block** | Pointer to [**RosettaBlock**](RosettaBlock.md) |  | [optional] 
**OtherTransactions** | Pointer to [**[]OtherTransactionIdentifier**](OtherTransactionIdentifier.md) | Some blockchains may require additional transactions to be fetched that weren&#39;t returned in the block response (ex: block only returns transaction hashes). For blockchains with a lot of transactions in each block, this can be very useful as consumers can concurrently fetch all transactions returned. | [optional] 

## Methods

### NewRosettaBlockResponse

`func NewRosettaBlockResponse() *RosettaBlockResponse`

NewRosettaBlockResponse instantiates a new RosettaBlockResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaBlockResponseWithDefaults

`func NewRosettaBlockResponseWithDefaults() *RosettaBlockResponse`

NewRosettaBlockResponseWithDefaults instantiates a new RosettaBlockResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlock

`func (o *RosettaBlockResponse) GetBlock() RosettaBlock`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *RosettaBlockResponse) GetBlockOk() (*RosettaBlock, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *RosettaBlockResponse) SetBlock(v RosettaBlock)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *RosettaBlockResponse) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetOtherTransactions

`func (o *RosettaBlockResponse) GetOtherTransactions() []OtherTransactionIdentifier`

GetOtherTransactions returns the OtherTransactions field if non-nil, zero value otherwise.

### GetOtherTransactionsOk

`func (o *RosettaBlockResponse) GetOtherTransactionsOk() (*[]OtherTransactionIdentifier, bool)`

GetOtherTransactionsOk returns a tuple with the OtherTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherTransactions

`func (o *RosettaBlockResponse) SetOtherTransactions(v []OtherTransactionIdentifier)`

SetOtherTransactions sets OtherTransactions field to given value.

### HasOtherTransactions

`func (o *RosettaBlockResponse) HasOtherTransactions() bool`

HasOtherTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


