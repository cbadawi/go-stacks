# RosettaErrorNoDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | Code is a network-specific error code. If desired, this code can be equivalent to an HTTP status code. | 
**Message** | **string** | Message is a network-specific error message. The message MUST NOT change for a given code. In particular, this means that any contextual information should be included in the details field. | 
**Retriable** | **bool** | An error is retriable if the same request may succeed if submitted again. | 

## Methods

### NewRosettaErrorNoDetails

`func NewRosettaErrorNoDetails(code int32, message string, retriable bool, ) *RosettaErrorNoDetails`

NewRosettaErrorNoDetails instantiates a new RosettaErrorNoDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaErrorNoDetailsWithDefaults

`func NewRosettaErrorNoDetailsWithDefaults() *RosettaErrorNoDetails`

NewRosettaErrorNoDetailsWithDefaults instantiates a new RosettaErrorNoDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *RosettaErrorNoDetails) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RosettaErrorNoDetails) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RosettaErrorNoDetails) SetCode(v int32)`

SetCode sets Code field to given value.


### GetMessage

`func (o *RosettaErrorNoDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RosettaErrorNoDetails) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RosettaErrorNoDetails) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRetriable

`func (o *RosettaErrorNoDetails) GetRetriable() bool`

GetRetriable returns the Retriable field if non-nil, zero value otherwise.

### GetRetriableOk

`func (o *RosettaErrorNoDetails) GetRetriableOk() (*bool, bool)`

GetRetriableOk returns a tuple with the Retriable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetriable

`func (o *RosettaErrorNoDetails) SetRetriable(v bool)`

SetRetriable sets Retriable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


