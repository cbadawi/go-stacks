# NetworkBlockTimesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mainnet** | [**TargetBlockTime**](TargetBlockTime.md) |  | 
**Testnet** | [**TargetBlockTime**](TargetBlockTime.md) |  | 

## Methods

### NewNetworkBlockTimesResponse

`func NewNetworkBlockTimesResponse(mainnet TargetBlockTime, testnet TargetBlockTime, ) *NetworkBlockTimesResponse`

NewNetworkBlockTimesResponse instantiates a new NetworkBlockTimesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkBlockTimesResponseWithDefaults

`func NewNetworkBlockTimesResponseWithDefaults() *NetworkBlockTimesResponse`

NewNetworkBlockTimesResponseWithDefaults instantiates a new NetworkBlockTimesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMainnet

`func (o *NetworkBlockTimesResponse) GetMainnet() TargetBlockTime`

GetMainnet returns the Mainnet field if non-nil, zero value otherwise.

### GetMainnetOk

`func (o *NetworkBlockTimesResponse) GetMainnetOk() (*TargetBlockTime, bool)`

GetMainnetOk returns a tuple with the Mainnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainnet

`func (o *NetworkBlockTimesResponse) SetMainnet(v TargetBlockTime)`

SetMainnet sets Mainnet field to given value.


### GetTestnet

`func (o *NetworkBlockTimesResponse) GetTestnet() TargetBlockTime`

GetTestnet returns the Testnet field if non-nil, zero value otherwise.

### GetTestnetOk

`func (o *NetworkBlockTimesResponse) GetTestnetOk() (*TargetBlockTime, bool)`

GetTestnetOk returns a tuple with the Testnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestnet

`func (o *NetworkBlockTimesResponse) SetTestnet(v TargetBlockTime)`

SetTestnet sets Testnet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


