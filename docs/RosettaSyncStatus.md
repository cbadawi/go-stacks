# RosettaSyncStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentIndex** | **int32** | CurrentIndex is the index of the last synced block in the current stage. | 
**TargetIndex** | Pointer to **int32** | TargetIndex is the index of the block that the implementation is attempting to sync to in the current stage. | [optional] 
**Stage** | Pointer to **string** | Stage is the phase of the sync process. | [optional] 
**Synced** | Pointer to **bool** | Synced indicates if an implementation has synced up to the most recent block. | [optional] 

## Methods

### NewRosettaSyncStatus

`func NewRosettaSyncStatus(currentIndex int32, ) *RosettaSyncStatus`

NewRosettaSyncStatus instantiates a new RosettaSyncStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaSyncStatusWithDefaults

`func NewRosettaSyncStatusWithDefaults() *RosettaSyncStatus`

NewRosettaSyncStatusWithDefaults instantiates a new RosettaSyncStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentIndex

`func (o *RosettaSyncStatus) GetCurrentIndex() int32`

GetCurrentIndex returns the CurrentIndex field if non-nil, zero value otherwise.

### GetCurrentIndexOk

`func (o *RosettaSyncStatus) GetCurrentIndexOk() (*int32, bool)`

GetCurrentIndexOk returns a tuple with the CurrentIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentIndex

`func (o *RosettaSyncStatus) SetCurrentIndex(v int32)`

SetCurrentIndex sets CurrentIndex field to given value.


### GetTargetIndex

`func (o *RosettaSyncStatus) GetTargetIndex() int32`

GetTargetIndex returns the TargetIndex field if non-nil, zero value otherwise.

### GetTargetIndexOk

`func (o *RosettaSyncStatus) GetTargetIndexOk() (*int32, bool)`

GetTargetIndexOk returns a tuple with the TargetIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIndex

`func (o *RosettaSyncStatus) SetTargetIndex(v int32)`

SetTargetIndex sets TargetIndex field to given value.

### HasTargetIndex

`func (o *RosettaSyncStatus) HasTargetIndex() bool`

HasTargetIndex returns a boolean if a field has been set.

### GetStage

`func (o *RosettaSyncStatus) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *RosettaSyncStatus) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *RosettaSyncStatus) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *RosettaSyncStatus) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetSynced

`func (o *RosettaSyncStatus) GetSynced() bool`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *RosettaSyncStatus) GetSyncedOk() (*bool, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *RosettaSyncStatus) SetSynced(v bool)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *RosettaSyncStatus) HasSynced() bool`

HasSynced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


