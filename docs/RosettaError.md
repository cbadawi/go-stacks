# RosettaError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | Code is a network-specific error code. If desired, this code can be equivalent to an HTTP status code. | 
**Message** | **string** | Message is a network-specific error message. The message MUST NOT change for a given code. In particular, this means that any contextual information should be included in the details field. | 
**Retriable** | **bool** | An error is retriable if the same request may succeed if submitted again. | 
**Details** | Pointer to [**RosettaErrorDetails**](RosettaErrorDetails.md) |  | [optional] 

## Methods

### NewRosettaError

`func NewRosettaError(code int32, message string, retriable bool, ) *RosettaError`

NewRosettaError instantiates a new RosettaError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaErrorWithDefaults

`func NewRosettaErrorWithDefaults() *RosettaError`

NewRosettaErrorWithDefaults instantiates a new RosettaError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *RosettaError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RosettaError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RosettaError) SetCode(v int32)`

SetCode sets Code field to given value.


### GetMessage

`func (o *RosettaError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RosettaError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RosettaError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRetriable

`func (o *RosettaError) GetRetriable() bool`

GetRetriable returns the Retriable field if non-nil, zero value otherwise.

### GetRetriableOk

`func (o *RosettaError) GetRetriableOk() (*bool, bool)`

GetRetriableOk returns a tuple with the Retriable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetriable

`func (o *RosettaError) SetRetriable(v bool)`

SetRetriable sets Retriable field to given value.


### GetDetails

`func (o *RosettaError) GetDetails() RosettaErrorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RosettaError) GetDetailsOk() (*RosettaErrorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RosettaError) SetDetails(v RosettaErrorDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RosettaError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


