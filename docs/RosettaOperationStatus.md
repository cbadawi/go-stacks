# RosettaOperationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The status is the network-specific status of the operation. | 
**Successful** | **bool** | An Operation is considered successful if the Operation.Amount should affect the Operation.Account. Some blockchains (like Bitcoin) only include successful operations in blocks but other blockchains (like Ethereum) include unsuccessful operations that incur a fee. To reconcile the computed balance from the stream of Operations, it is critical to understand which Operation.Status indicate an Operation is successful and should affect an Account. | 

## Methods

### NewRosettaOperationStatus

`func NewRosettaOperationStatus(status string, successful bool, ) *RosettaOperationStatus`

NewRosettaOperationStatus instantiates a new RosettaOperationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaOperationStatusWithDefaults

`func NewRosettaOperationStatusWithDefaults() *RosettaOperationStatus`

NewRosettaOperationStatusWithDefaults instantiates a new RosettaOperationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RosettaOperationStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RosettaOperationStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RosettaOperationStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSuccessful

`func (o *RosettaOperationStatus) GetSuccessful() bool`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *RosettaOperationStatus) GetSuccessfulOk() (*bool, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *RosettaOperationStatus) SetSuccessful(v bool)`

SetSuccessful sets Successful field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


