# ReadOnlyFunctionSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Okay** | **bool** |  | 
**Result** | Pointer to **string** |  | [optional] 
**Cause** | Pointer to **string** |  | [optional] 

## Methods

### NewReadOnlyFunctionSuccessResponse

`func NewReadOnlyFunctionSuccessResponse(okay bool, ) *ReadOnlyFunctionSuccessResponse`

NewReadOnlyFunctionSuccessResponse instantiates a new ReadOnlyFunctionSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadOnlyFunctionSuccessResponseWithDefaults

`func NewReadOnlyFunctionSuccessResponseWithDefaults() *ReadOnlyFunctionSuccessResponse`

NewReadOnlyFunctionSuccessResponseWithDefaults instantiates a new ReadOnlyFunctionSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOkay

`func (o *ReadOnlyFunctionSuccessResponse) GetOkay() bool`

GetOkay returns the Okay field if non-nil, zero value otherwise.

### GetOkayOk

`func (o *ReadOnlyFunctionSuccessResponse) GetOkayOk() (*bool, bool)`

GetOkayOk returns a tuple with the Okay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOkay

`func (o *ReadOnlyFunctionSuccessResponse) SetOkay(v bool)`

SetOkay sets Okay field to given value.


### GetResult

`func (o *ReadOnlyFunctionSuccessResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ReadOnlyFunctionSuccessResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ReadOnlyFunctionSuccessResponse) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ReadOnlyFunctionSuccessResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetCause

`func (o *ReadOnlyFunctionSuccessResponse) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *ReadOnlyFunctionSuccessResponse) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *ReadOnlyFunctionSuccessResponse) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *ReadOnlyFunctionSuccessResponse) HasCause() bool`

HasCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


