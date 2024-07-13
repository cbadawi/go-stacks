# PoxSigner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SigningKey** | **string** |  | 
**SignerAddress** | **string** | The Stacks address derived from the signing_key. | 
**Weight** | **int32** |  | 
**StackedAmount** | **string** |  | 
**WeightPercent** | **float32** |  | 
**StackedAmountPercent** | **float32** |  | 
**SoloStackerCount** | **int32** | The number of solo stackers associated with this signer. | 
**PooledStackerCount** | **int32** | The number of pooled stackers associated with this signer. | 

## Methods

### NewPoxSigner

`func NewPoxSigner(signingKey string, signerAddress string, weight int32, stackedAmount string, weightPercent float32, stackedAmountPercent float32, soloStackerCount int32, pooledStackerCount int32, ) *PoxSigner`

NewPoxSigner instantiates a new PoxSigner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoxSignerWithDefaults

`func NewPoxSignerWithDefaults() *PoxSigner`

NewPoxSignerWithDefaults instantiates a new PoxSigner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigningKey

`func (o *PoxSigner) GetSigningKey() string`

GetSigningKey returns the SigningKey field if non-nil, zero value otherwise.

### GetSigningKeyOk

`func (o *PoxSigner) GetSigningKeyOk() (*string, bool)`

GetSigningKeyOk returns a tuple with the SigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKey

`func (o *PoxSigner) SetSigningKey(v string)`

SetSigningKey sets SigningKey field to given value.


### GetSignerAddress

`func (o *PoxSigner) GetSignerAddress() string`

GetSignerAddress returns the SignerAddress field if non-nil, zero value otherwise.

### GetSignerAddressOk

`func (o *PoxSigner) GetSignerAddressOk() (*string, bool)`

GetSignerAddressOk returns a tuple with the SignerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerAddress

`func (o *PoxSigner) SetSignerAddress(v string)`

SetSignerAddress sets SignerAddress field to given value.


### GetWeight

`func (o *PoxSigner) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PoxSigner) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PoxSigner) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetStackedAmount

`func (o *PoxSigner) GetStackedAmount() string`

GetStackedAmount returns the StackedAmount field if non-nil, zero value otherwise.

### GetStackedAmountOk

`func (o *PoxSigner) GetStackedAmountOk() (*string, bool)`

GetStackedAmountOk returns a tuple with the StackedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackedAmount

`func (o *PoxSigner) SetStackedAmount(v string)`

SetStackedAmount sets StackedAmount field to given value.


### GetWeightPercent

`func (o *PoxSigner) GetWeightPercent() float32`

GetWeightPercent returns the WeightPercent field if non-nil, zero value otherwise.

### GetWeightPercentOk

`func (o *PoxSigner) GetWeightPercentOk() (*float32, bool)`

GetWeightPercentOk returns a tuple with the WeightPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightPercent

`func (o *PoxSigner) SetWeightPercent(v float32)`

SetWeightPercent sets WeightPercent field to given value.


### GetStackedAmountPercent

`func (o *PoxSigner) GetStackedAmountPercent() float32`

GetStackedAmountPercent returns the StackedAmountPercent field if non-nil, zero value otherwise.

### GetStackedAmountPercentOk

`func (o *PoxSigner) GetStackedAmountPercentOk() (*float32, bool)`

GetStackedAmountPercentOk returns a tuple with the StackedAmountPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackedAmountPercent

`func (o *PoxSigner) SetStackedAmountPercent(v float32)`

SetStackedAmountPercent sets StackedAmountPercent field to given value.


### GetSoloStackerCount

`func (o *PoxSigner) GetSoloStackerCount() int32`

GetSoloStackerCount returns the SoloStackerCount field if non-nil, zero value otherwise.

### GetSoloStackerCountOk

`func (o *PoxSigner) GetSoloStackerCountOk() (*int32, bool)`

GetSoloStackerCountOk returns a tuple with the SoloStackerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoloStackerCount

`func (o *PoxSigner) SetSoloStackerCount(v int32)`

SetSoloStackerCount sets SoloStackerCount field to given value.


### GetPooledStackerCount

`func (o *PoxSigner) GetPooledStackerCount() int32`

GetPooledStackerCount returns the PooledStackerCount field if non-nil, zero value otherwise.

### GetPooledStackerCountOk

`func (o *PoxSigner) GetPooledStackerCountOk() (*int32, bool)`

GetPooledStackerCountOk returns a tuple with the PooledStackerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPooledStackerCount

`func (o *PoxSigner) SetPooledStackerCount(v int32)`

SetPooledStackerCount sets PooledStackerCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


