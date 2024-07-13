# PostConditionNonFungible

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principal** | [**PostConditionPrincipal**](PostConditionPrincipal.md) |  | 
**ConditionCode** | **string** | A non-fungible condition code encodes a statement being made about a non-fungible token, with respect to whether or not the particular non-fungible token is owned by the account. | 
**Type** | **string** |  | 
**AssetValue** | [**PostConditionNonFungibleAllOfAssetValue**](PostConditionNonFungibleAllOfAssetValue.md) |  | 
**Asset** | [**PostConditionFungibleAllOfAsset**](PostConditionFungibleAllOfAsset.md) |  | 

## Methods

### NewPostConditionNonFungible

`func NewPostConditionNonFungible(principal PostConditionPrincipal, conditionCode string, type_ string, assetValue PostConditionNonFungibleAllOfAssetValue, asset PostConditionFungibleAllOfAsset, ) *PostConditionNonFungible`

NewPostConditionNonFungible instantiates a new PostConditionNonFungible object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostConditionNonFungibleWithDefaults

`func NewPostConditionNonFungibleWithDefaults() *PostConditionNonFungible`

NewPostConditionNonFungibleWithDefaults instantiates a new PostConditionNonFungible object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipal

`func (o *PostConditionNonFungible) GetPrincipal() PostConditionPrincipal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *PostConditionNonFungible) GetPrincipalOk() (*PostConditionPrincipal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *PostConditionNonFungible) SetPrincipal(v PostConditionPrincipal)`

SetPrincipal sets Principal field to given value.


### GetConditionCode

`func (o *PostConditionNonFungible) GetConditionCode() string`

GetConditionCode returns the ConditionCode field if non-nil, zero value otherwise.

### GetConditionCodeOk

`func (o *PostConditionNonFungible) GetConditionCodeOk() (*string, bool)`

GetConditionCodeOk returns a tuple with the ConditionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionCode

`func (o *PostConditionNonFungible) SetConditionCode(v string)`

SetConditionCode sets ConditionCode field to given value.


### GetType

`func (o *PostConditionNonFungible) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostConditionNonFungible) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostConditionNonFungible) SetType(v string)`

SetType sets Type field to given value.


### GetAssetValue

`func (o *PostConditionNonFungible) GetAssetValue() PostConditionNonFungibleAllOfAssetValue`

GetAssetValue returns the AssetValue field if non-nil, zero value otherwise.

### GetAssetValueOk

`func (o *PostConditionNonFungible) GetAssetValueOk() (*PostConditionNonFungibleAllOfAssetValue, bool)`

GetAssetValueOk returns a tuple with the AssetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetValue

`func (o *PostConditionNonFungible) SetAssetValue(v PostConditionNonFungibleAllOfAssetValue)`

SetAssetValue sets AssetValue field to given value.


### GetAsset

`func (o *PostConditionNonFungible) GetAsset() PostConditionFungibleAllOfAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *PostConditionNonFungible) GetAssetOk() (*PostConditionFungibleAllOfAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *PostConditionNonFungible) SetAsset(v PostConditionFungibleAllOfAsset)`

SetAsset sets Asset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


