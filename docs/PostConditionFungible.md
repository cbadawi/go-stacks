# PostConditionFungible

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principal** | [**PostConditionPrincipal**](PostConditionPrincipal.md) |  | 
**ConditionCode** | **string** | A fungible condition code encodes a statement being made for either STX or a fungible token, with respect to the originating account. | 
**Type** | **string** |  | 
**Amount** | **string** |  | 
**Asset** | [**PostConditionFungibleAllOfAsset**](PostConditionFungibleAllOfAsset.md) |  | 

## Methods

### NewPostConditionFungible

`func NewPostConditionFungible(principal PostConditionPrincipal, conditionCode string, type_ string, amount string, asset PostConditionFungibleAllOfAsset, ) *PostConditionFungible`

NewPostConditionFungible instantiates a new PostConditionFungible object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostConditionFungibleWithDefaults

`func NewPostConditionFungibleWithDefaults() *PostConditionFungible`

NewPostConditionFungibleWithDefaults instantiates a new PostConditionFungible object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipal

`func (o *PostConditionFungible) GetPrincipal() PostConditionPrincipal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *PostConditionFungible) GetPrincipalOk() (*PostConditionPrincipal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *PostConditionFungible) SetPrincipal(v PostConditionPrincipal)`

SetPrincipal sets Principal field to given value.


### GetConditionCode

`func (o *PostConditionFungible) GetConditionCode() string`

GetConditionCode returns the ConditionCode field if non-nil, zero value otherwise.

### GetConditionCodeOk

`func (o *PostConditionFungible) GetConditionCodeOk() (*string, bool)`

GetConditionCodeOk returns a tuple with the ConditionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionCode

`func (o *PostConditionFungible) SetConditionCode(v string)`

SetConditionCode sets ConditionCode field to given value.


### GetType

`func (o *PostConditionFungible) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostConditionFungible) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostConditionFungible) SetType(v string)`

SetType sets Type field to given value.


### GetAmount

`func (o *PostConditionFungible) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PostConditionFungible) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PostConditionFungible) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAsset

`func (o *PostConditionFungible) GetAsset() PostConditionFungibleAllOfAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *PostConditionFungible) GetAssetOk() (*PostConditionFungibleAllOfAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *PostConditionFungible) SetAsset(v PostConditionFungibleAllOfAsset)`

SetAsset sets Asset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


