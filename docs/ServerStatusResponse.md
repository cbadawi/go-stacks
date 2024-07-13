# ServerStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerVersion** | Pointer to **string** | the server version that is currently running | [optional] 
**Status** | **string** | the current server status | 
**PoxV1UnlockHeight** | Pointer to **NullableInt32** |  | [optional] 
**PoxV2UnlockHeight** | Pointer to **NullableInt32** |  | [optional] 
**PoxV3UnlockHeight** | Pointer to **NullableInt32** |  | [optional] 
**ChainTip** | Pointer to [**ChainTip**](ChainTip.md) |  | [optional] 

## Methods

### NewServerStatusResponse

`func NewServerStatusResponse(status string, ) *ServerStatusResponse`

NewServerStatusResponse instantiates a new ServerStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerStatusResponseWithDefaults

`func NewServerStatusResponseWithDefaults() *ServerStatusResponse`

NewServerStatusResponseWithDefaults instantiates a new ServerStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerVersion

`func (o *ServerStatusResponse) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *ServerStatusResponse) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *ServerStatusResponse) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *ServerStatusResponse) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetStatus

`func (o *ServerStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPoxV1UnlockHeight

`func (o *ServerStatusResponse) GetPoxV1UnlockHeight() int32`

GetPoxV1UnlockHeight returns the PoxV1UnlockHeight field if non-nil, zero value otherwise.

### GetPoxV1UnlockHeightOk

`func (o *ServerStatusResponse) GetPoxV1UnlockHeightOk() (*int32, bool)`

GetPoxV1UnlockHeightOk returns a tuple with the PoxV1UnlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoxV1UnlockHeight

`func (o *ServerStatusResponse) SetPoxV1UnlockHeight(v int32)`

SetPoxV1UnlockHeight sets PoxV1UnlockHeight field to given value.

### HasPoxV1UnlockHeight

`func (o *ServerStatusResponse) HasPoxV1UnlockHeight() bool`

HasPoxV1UnlockHeight returns a boolean if a field has been set.

### SetPoxV1UnlockHeightNil

`func (o *ServerStatusResponse) SetPoxV1UnlockHeightNil(b bool)`

 SetPoxV1UnlockHeightNil sets the value for PoxV1UnlockHeight to be an explicit nil

### UnsetPoxV1UnlockHeight
`func (o *ServerStatusResponse) UnsetPoxV1UnlockHeight()`

UnsetPoxV1UnlockHeight ensures that no value is present for PoxV1UnlockHeight, not even an explicit nil
### GetPoxV2UnlockHeight

`func (o *ServerStatusResponse) GetPoxV2UnlockHeight() int32`

GetPoxV2UnlockHeight returns the PoxV2UnlockHeight field if non-nil, zero value otherwise.

### GetPoxV2UnlockHeightOk

`func (o *ServerStatusResponse) GetPoxV2UnlockHeightOk() (*int32, bool)`

GetPoxV2UnlockHeightOk returns a tuple with the PoxV2UnlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoxV2UnlockHeight

`func (o *ServerStatusResponse) SetPoxV2UnlockHeight(v int32)`

SetPoxV2UnlockHeight sets PoxV2UnlockHeight field to given value.

### HasPoxV2UnlockHeight

`func (o *ServerStatusResponse) HasPoxV2UnlockHeight() bool`

HasPoxV2UnlockHeight returns a boolean if a field has been set.

### SetPoxV2UnlockHeightNil

`func (o *ServerStatusResponse) SetPoxV2UnlockHeightNil(b bool)`

 SetPoxV2UnlockHeightNil sets the value for PoxV2UnlockHeight to be an explicit nil

### UnsetPoxV2UnlockHeight
`func (o *ServerStatusResponse) UnsetPoxV2UnlockHeight()`

UnsetPoxV2UnlockHeight ensures that no value is present for PoxV2UnlockHeight, not even an explicit nil
### GetPoxV3UnlockHeight

`func (o *ServerStatusResponse) GetPoxV3UnlockHeight() int32`

GetPoxV3UnlockHeight returns the PoxV3UnlockHeight field if non-nil, zero value otherwise.

### GetPoxV3UnlockHeightOk

`func (o *ServerStatusResponse) GetPoxV3UnlockHeightOk() (*int32, bool)`

GetPoxV3UnlockHeightOk returns a tuple with the PoxV3UnlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoxV3UnlockHeight

`func (o *ServerStatusResponse) SetPoxV3UnlockHeight(v int32)`

SetPoxV3UnlockHeight sets PoxV3UnlockHeight field to given value.

### HasPoxV3UnlockHeight

`func (o *ServerStatusResponse) HasPoxV3UnlockHeight() bool`

HasPoxV3UnlockHeight returns a boolean if a field has been set.

### SetPoxV3UnlockHeightNil

`func (o *ServerStatusResponse) SetPoxV3UnlockHeightNil(b bool)`

 SetPoxV3UnlockHeightNil sets the value for PoxV3UnlockHeight to be an explicit nil

### UnsetPoxV3UnlockHeight
`func (o *ServerStatusResponse) UnsetPoxV3UnlockHeight()`

UnsetPoxV3UnlockHeight ensures that no value is present for PoxV3UnlockHeight, not even an explicit nil
### GetChainTip

`func (o *ServerStatusResponse) GetChainTip() ChainTip`

GetChainTip returns the ChainTip field if non-nil, zero value otherwise.

### GetChainTipOk

`func (o *ServerStatusResponse) GetChainTipOk() (*ChainTip, bool)`

GetChainTipOk returns a tuple with the ChainTip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainTip

`func (o *ServerStatusResponse) SetChainTip(v ChainTip)`

SetChainTip sets ChainTip field to given value.

### HasChainTip

`func (o *ServerStatusResponse) HasChainTip() bool`

HasChainTip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


