# PostConditionStx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principal** | [**PostConditionPrincipal**](PostConditionPrincipal.md) |  | 
**ConditionCode** | **string** | A fungible condition code encodes a statement being made for either STX or a fungible token, with respect to the originating account. | 
**Amount** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewPostConditionStx

`func NewPostConditionStx(principal PostConditionPrincipal, conditionCode string, amount string, type_ string, ) *PostConditionStx`

NewPostConditionStx instantiates a new PostConditionStx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostConditionStxWithDefaults

`func NewPostConditionStxWithDefaults() *PostConditionStx`

NewPostConditionStxWithDefaults instantiates a new PostConditionStx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipal

`func (o *PostConditionStx) GetPrincipal() PostConditionPrincipal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *PostConditionStx) GetPrincipalOk() (*PostConditionPrincipal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *PostConditionStx) SetPrincipal(v PostConditionPrincipal)`

SetPrincipal sets Principal field to given value.


### GetConditionCode

`func (o *PostConditionStx) GetConditionCode() string`

GetConditionCode returns the ConditionCode field if non-nil, zero value otherwise.

### GetConditionCodeOk

`func (o *PostConditionStx) GetConditionCodeOk() (*string, bool)`

GetConditionCodeOk returns a tuple with the ConditionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionCode

`func (o *PostConditionStx) SetConditionCode(v string)`

SetConditionCode sets ConditionCode field to given value.


### GetAmount

`func (o *PostConditionStx) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PostConditionStx) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PostConditionStx) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetType

`func (o *PostConditionStx) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostConditionStx) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostConditionStx) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


