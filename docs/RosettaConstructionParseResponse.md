# RosettaConstructionParseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | [**[]RosettaOperation**](RosettaOperation.md) |  | 
**Signers** | Pointer to [**[]RosettaNetworkOptionsResponseAllowOperationTypesInner**](RosettaNetworkOptionsResponseAllowOperationTypesInner.md) | [DEPRECATED by account_identifier_signers in v1.4.4] All signers (addresses) of a particular transaction. If the transaction is unsigned, it should be empty. | [optional] 
**AccountIdentifierSigners** | Pointer to [**[]RosettaAccountIdentifier**](RosettaAccountIdentifier.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRosettaConstructionParseResponse

`func NewRosettaConstructionParseResponse(operations []RosettaOperation, ) *RosettaConstructionParseResponse`

NewRosettaConstructionParseResponse instantiates a new RosettaConstructionParseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaConstructionParseResponseWithDefaults

`func NewRosettaConstructionParseResponseWithDefaults() *RosettaConstructionParseResponse`

NewRosettaConstructionParseResponseWithDefaults instantiates a new RosettaConstructionParseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *RosettaConstructionParseResponse) GetOperations() []RosettaOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *RosettaConstructionParseResponse) GetOperationsOk() (*[]RosettaOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *RosettaConstructionParseResponse) SetOperations(v []RosettaOperation)`

SetOperations sets Operations field to given value.


### GetSigners

`func (o *RosettaConstructionParseResponse) GetSigners() []RosettaNetworkOptionsResponseAllowOperationTypesInner`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *RosettaConstructionParseResponse) GetSignersOk() (*[]RosettaNetworkOptionsResponseAllowOperationTypesInner, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *RosettaConstructionParseResponse) SetSigners(v []RosettaNetworkOptionsResponseAllowOperationTypesInner)`

SetSigners sets Signers field to given value.

### HasSigners

`func (o *RosettaConstructionParseResponse) HasSigners() bool`

HasSigners returns a boolean if a field has been set.

### GetAccountIdentifierSigners

`func (o *RosettaConstructionParseResponse) GetAccountIdentifierSigners() []RosettaAccountIdentifier`

GetAccountIdentifierSigners returns the AccountIdentifierSigners field if non-nil, zero value otherwise.

### GetAccountIdentifierSignersOk

`func (o *RosettaConstructionParseResponse) GetAccountIdentifierSignersOk() (*[]RosettaAccountIdentifier, bool)`

GetAccountIdentifierSignersOk returns a tuple with the AccountIdentifierSigners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentifierSigners

`func (o *RosettaConstructionParseResponse) SetAccountIdentifierSigners(v []RosettaAccountIdentifier)`

SetAccountIdentifierSigners sets AccountIdentifierSigners field to given value.

### HasAccountIdentifierSigners

`func (o *RosettaConstructionParseResponse) HasAccountIdentifierSigners() bool`

HasAccountIdentifierSigners returns a boolean if a field has been set.

### GetMetadata

`func (o *RosettaConstructionParseResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RosettaConstructionParseResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RosettaConstructionParseResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RosettaConstructionParseResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


