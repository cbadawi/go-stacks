# PostCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principal** | [**PostConditionPrincipal**](PostConditionPrincipal.md) |  | 
**ConditionCode** | **string** | A non-fungible condition code encodes a statement being made about a non-fungible token, with respect to whether or not the particular non-fungible token is owned by the account. | 
**Amount** | **string** |  | 
**Type** | **string** |  | 
**Asset** | [**PostConditionFungibleAllOfAsset**](PostConditionFungibleAllOfAsset.md) |  | 
**AssetValue** | [**PostConditionNonFungibleAllOfAssetValue**](PostConditionNonFungibleAllOfAssetValue.md) |  | 

## Methods

### NewPostCondition

`func NewPostCondition(principal PostConditionPrincipal, conditionCode string, amount string, type_ string, asset PostConditionFungibleAllOfAsset, assetValue PostConditionNonFungibleAllOfAssetValue, ) *PostCondition`

NewPostCondition instantiates a new PostCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostConditionWithDefaults

`func NewPostConditionWithDefaults() *PostCondition`

NewPostConditionWithDefaults instantiates a new PostCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipal

`func (o *PostCondition) GetPrincipal() PostConditionPrincipal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *PostCondition) GetPrincipalOk() (*PostConditionPrincipal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *PostCondition) SetPrincipal(v PostConditionPrincipal)`

SetPrincipal sets Principal field to given value.


### GetConditionCode

`func (o *PostCondition) GetConditionCode() string`

GetConditionCode returns the ConditionCode field if non-nil, zero value otherwise.

### GetConditionCodeOk

`func (o *PostCondition) GetConditionCodeOk() (*string, bool)`

GetConditionCodeOk returns a tuple with the ConditionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionCode

`func (o *PostCondition) SetConditionCode(v string)`

SetConditionCode sets ConditionCode field to given value.


### GetAmount

`func (o *PostCondition) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PostCondition) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PostCondition) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetType

`func (o *PostCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCondition) SetType(v string)`

SetType sets Type field to given value.


### GetAsset

`func (o *PostCondition) GetAsset() PostConditionFungibleAllOfAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *PostCondition) GetAssetOk() (*PostConditionFungibleAllOfAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *PostCondition) SetAsset(v PostConditionFungibleAllOfAsset)`

SetAsset sets Asset field to given value.


### GetAssetValue

`func (o *PostCondition) GetAssetValue() PostConditionNonFungibleAllOfAssetValue`

GetAssetValue returns the AssetValue field if non-nil, zero value otherwise.

### GetAssetValueOk

`func (o *PostCondition) GetAssetValueOk() (*PostConditionNonFungibleAllOfAssetValue, bool)`

GetAssetValueOk returns a tuple with the AssetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetValue

`func (o *PostCondition) SetAssetValue(v PostConditionNonFungibleAllOfAssetValue)`

SetAssetValue sets AssetValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


