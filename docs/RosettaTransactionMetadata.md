# RosettaTransactionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memo** | Pointer to **string** | STX token transfer memo. | [optional] 
**Size** | Pointer to **int32** | The Size | [optional] 
**LockTime** | Pointer to **int32** | The locktime | [optional] 

## Methods

### NewRosettaTransactionMetadata

`func NewRosettaTransactionMetadata() *RosettaTransactionMetadata`

NewRosettaTransactionMetadata instantiates a new RosettaTransactionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaTransactionMetadataWithDefaults

`func NewRosettaTransactionMetadataWithDefaults() *RosettaTransactionMetadata`

NewRosettaTransactionMetadataWithDefaults instantiates a new RosettaTransactionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemo

`func (o *RosettaTransactionMetadata) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *RosettaTransactionMetadata) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *RosettaTransactionMetadata) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *RosettaTransactionMetadata) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetSize

`func (o *RosettaTransactionMetadata) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RosettaTransactionMetadata) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RosettaTransactionMetadata) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *RosettaTransactionMetadata) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetLockTime

`func (o *RosettaTransactionMetadata) GetLockTime() int32`

GetLockTime returns the LockTime field if non-nil, zero value otherwise.

### GetLockTimeOk

`func (o *RosettaTransactionMetadata) GetLockTimeOk() (*int32, bool)`

GetLockTimeOk returns a tuple with the LockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTime

`func (o *RosettaTransactionMetadata) SetLockTime(v int32)`

SetLockTime sets LockTime field to given value.

### HasLockTime

`func (o *RosettaTransactionMetadata) HasLockTime() bool`

HasLockTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


