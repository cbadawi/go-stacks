# RosettaPublicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HexBytes** | **string** | Hex-encoded public key bytes in the format specified by the CurveType. | 
**CurveType** | **string** | CurveType is the type of cryptographic curve associated with a PublicKey. | 

## Methods

### NewRosettaPublicKey

`func NewRosettaPublicKey(hexBytes string, curveType string, ) *RosettaPublicKey`

NewRosettaPublicKey instantiates a new RosettaPublicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaPublicKeyWithDefaults

`func NewRosettaPublicKeyWithDefaults() *RosettaPublicKey`

NewRosettaPublicKeyWithDefaults instantiates a new RosettaPublicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHexBytes

`func (o *RosettaPublicKey) GetHexBytes() string`

GetHexBytes returns the HexBytes field if non-nil, zero value otherwise.

### GetHexBytesOk

`func (o *RosettaPublicKey) GetHexBytesOk() (*string, bool)`

GetHexBytesOk returns a tuple with the HexBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexBytes

`func (o *RosettaPublicKey) SetHexBytes(v string)`

SetHexBytes sets HexBytes field to given value.


### GetCurveType

`func (o *RosettaPublicKey) GetCurveType() string`

GetCurveType returns the CurveType field if non-nil, zero value otherwise.

### GetCurveTypeOk

`func (o *RosettaPublicKey) GetCurveTypeOk() (*string, bool)`

GetCurveTypeOk returns a tuple with the CurveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurveType

`func (o *RosettaPublicKey) SetCurveType(v string)`

SetCurveType sets CurveType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


