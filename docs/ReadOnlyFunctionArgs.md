# ReadOnlyFunctionArgs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | **string** | The simulated tx-sender | 
**Arguments** | **[]string** | An array of hex serialized Clarity values | 

## Methods

### NewReadOnlyFunctionArgs

`func NewReadOnlyFunctionArgs(sender string, arguments []string, ) *ReadOnlyFunctionArgs`

NewReadOnlyFunctionArgs instantiates a new ReadOnlyFunctionArgs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadOnlyFunctionArgsWithDefaults

`func NewReadOnlyFunctionArgsWithDefaults() *ReadOnlyFunctionArgs`

NewReadOnlyFunctionArgsWithDefaults instantiates a new ReadOnlyFunctionArgs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *ReadOnlyFunctionArgs) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *ReadOnlyFunctionArgs) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *ReadOnlyFunctionArgs) SetSender(v string)`

SetSender sets Sender field to given value.


### GetArguments

`func (o *ReadOnlyFunctionArgs) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ReadOnlyFunctionArgs) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ReadOnlyFunctionArgs) SetArguments(v []string)`

SetArguments sets Arguments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


