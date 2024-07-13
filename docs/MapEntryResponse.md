# MapEntryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | Hex-encoded string of clarity value. It is always an optional tuple. | 
**Proof** | Pointer to **string** | Hex-encoded string of the MARF proof for the data | [optional] 

## Methods

### NewMapEntryResponse

`func NewMapEntryResponse(data string, ) *MapEntryResponse`

NewMapEntryResponse instantiates a new MapEntryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapEntryResponseWithDefaults

`func NewMapEntryResponseWithDefaults() *MapEntryResponse`

NewMapEntryResponseWithDefaults instantiates a new MapEntryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MapEntryResponse) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MapEntryResponse) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MapEntryResponse) SetData(v string)`

SetData sets Data field to given value.


### GetProof

`func (o *MapEntryResponse) GetProof() string`

GetProof returns the Proof field if non-nil, zero value otherwise.

### GetProofOk

`func (o *MapEntryResponse) GetProofOk() (*string, bool)`

GetProofOk returns a tuple with the Proof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProof

`func (o *MapEntryResponse) SetProof(v string)`

SetProof sets Proof field to given value.

### HasProof

`func (o *MapEntryResponse) HasProof() bool`

HasProof returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


