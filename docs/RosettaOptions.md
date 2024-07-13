# RosettaOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SenderAddress** | Pointer to **string** | sender&#39;s address  | [optional] 
**Type** | Pointer to **string** | Type of operation e.g transfer | [optional] 
**Status** | Pointer to **string** | This value indicates the state of the operations | [optional] 
**TokenTransferRecipientAddress** | Pointer to **string** | Recipient&#39;s address | [optional] 
**Amount** | Pointer to **string** | Amount to be transfered. | [optional] 
**Symbol** | Pointer to **string** | Currency symbol e.g STX | [optional] 
**Decimals** | Pointer to **int32** | Number of decimal places | [optional] 
**GasLimit** | Pointer to **float32** | Maximum price a user is willing to pay. | [optional] 
**GasPrice** | Pointer to **float32** | Cost necessary to perform a transaction on the network | [optional] 
**SuggestedFeeMultiplier** | Pointer to **float32** |  A suggested fee multiplier to indicate that the suggested fee should be scaled. This may be used to set higher fees for urgent transactions or to pay lower fees when there is less urgency. | [optional] 
**MaxFee** | Pointer to **string** | Maximum fee user is willing to pay | [optional] 
**Fee** | Pointer to **string** | Fee for this transaction | [optional] 
**Size** | Pointer to **int32** | Transaction approximative size (used to calculate total fee). | [optional] 
**Memo** | Pointer to **string** | STX token transfer memo. | [optional] 
**NumberOfCycles** | Pointer to **int32** | Number of cycles when stacking. | [optional] 
**ContractAddress** | Pointer to **string** | Address of the contract to call. | [optional] 
**ContractName** | Pointer to **string** | Name of the contract to call. | [optional] 
**BurnBlockHeight** | Pointer to **int32** | Set the burnchain (BTC) block for stacking lock to start. | [optional] 
**RewardCycleId** | Pointer to **int32** | The reward cycle ID for stacking transaction. | [optional] 
**DelegateTo** | Pointer to **string** | Delegator address for when calling &#x60;delegate-stacking&#x60;. | [optional] 
**PoxAddr** | Pointer to **string** | The reward address for stacking transaction. It should be a valid Bitcoin address | [optional] 
**SignerKey** | Pointer to **string** | The hex-encoded signer key (buff 33) for PoX. | [optional] 
**SignerPrivateKey** | Pointer to **string** | The hex-encoded signer private key for PoX. Specify either this or &#x60;signer_signature&#x60;, otherwise the PoX transaction requires allow-listing from the signer. | [optional] 
**SignerSignature** | Pointer to **string** | The hex-encoded signer signature for PoX. Specify either this or &#x60;signer_private_key&#x60;, otherwise the PoX transaction requires allow-listing from the signer. | [optional] 
**PoxMaxAmount** | Pointer to **string** | The maximum amount of STX to stack for PoX. If not specified, the &#x60;amount&#x60; will be used as the &#x60;max-amount&#x60; for the PoX transaction. | [optional] 
**PoxAuthId** | Pointer to **string** | The auth ID for the PoX transaction. If not specified, a random value will be generated. | [optional] 

## Methods

### NewRosettaOptions

`func NewRosettaOptions() *RosettaOptions`

NewRosettaOptions instantiates a new RosettaOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaOptionsWithDefaults

`func NewRosettaOptionsWithDefaults() *RosettaOptions`

NewRosettaOptionsWithDefaults instantiates a new RosettaOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSenderAddress

`func (o *RosettaOptions) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *RosettaOptions) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *RosettaOptions) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.

### HasSenderAddress

`func (o *RosettaOptions) HasSenderAddress() bool`

HasSenderAddress returns a boolean if a field has been set.

### GetType

`func (o *RosettaOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RosettaOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RosettaOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RosettaOptions) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *RosettaOptions) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RosettaOptions) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RosettaOptions) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RosettaOptions) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTokenTransferRecipientAddress

`func (o *RosettaOptions) GetTokenTransferRecipientAddress() string`

GetTokenTransferRecipientAddress returns the TokenTransferRecipientAddress field if non-nil, zero value otherwise.

### GetTokenTransferRecipientAddressOk

`func (o *RosettaOptions) GetTokenTransferRecipientAddressOk() (*string, bool)`

GetTokenTransferRecipientAddressOk returns a tuple with the TokenTransferRecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTransferRecipientAddress

`func (o *RosettaOptions) SetTokenTransferRecipientAddress(v string)`

SetTokenTransferRecipientAddress sets TokenTransferRecipientAddress field to given value.

### HasTokenTransferRecipientAddress

`func (o *RosettaOptions) HasTokenTransferRecipientAddress() bool`

HasTokenTransferRecipientAddress returns a boolean if a field has been set.

### GetAmount

`func (o *RosettaOptions) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RosettaOptions) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RosettaOptions) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RosettaOptions) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetSymbol

`func (o *RosettaOptions) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *RosettaOptions) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *RosettaOptions) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *RosettaOptions) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDecimals

`func (o *RosettaOptions) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *RosettaOptions) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *RosettaOptions) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *RosettaOptions) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetGasLimit

`func (o *RosettaOptions) GetGasLimit() float32`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *RosettaOptions) GetGasLimitOk() (*float32, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *RosettaOptions) SetGasLimit(v float32)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *RosettaOptions) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.

### GetGasPrice

`func (o *RosettaOptions) GetGasPrice() float32`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *RosettaOptions) GetGasPriceOk() (*float32, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *RosettaOptions) SetGasPrice(v float32)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *RosettaOptions) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetSuggestedFeeMultiplier

`func (o *RosettaOptions) GetSuggestedFeeMultiplier() float32`

GetSuggestedFeeMultiplier returns the SuggestedFeeMultiplier field if non-nil, zero value otherwise.

### GetSuggestedFeeMultiplierOk

`func (o *RosettaOptions) GetSuggestedFeeMultiplierOk() (*float32, bool)`

GetSuggestedFeeMultiplierOk returns a tuple with the SuggestedFeeMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedFeeMultiplier

`func (o *RosettaOptions) SetSuggestedFeeMultiplier(v float32)`

SetSuggestedFeeMultiplier sets SuggestedFeeMultiplier field to given value.

### HasSuggestedFeeMultiplier

`func (o *RosettaOptions) HasSuggestedFeeMultiplier() bool`

HasSuggestedFeeMultiplier returns a boolean if a field has been set.

### GetMaxFee

`func (o *RosettaOptions) GetMaxFee() string`

GetMaxFee returns the MaxFee field if non-nil, zero value otherwise.

### GetMaxFeeOk

`func (o *RosettaOptions) GetMaxFeeOk() (*string, bool)`

GetMaxFeeOk returns a tuple with the MaxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFee

`func (o *RosettaOptions) SetMaxFee(v string)`

SetMaxFee sets MaxFee field to given value.

### HasMaxFee

`func (o *RosettaOptions) HasMaxFee() bool`

HasMaxFee returns a boolean if a field has been set.

### GetFee

`func (o *RosettaOptions) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *RosettaOptions) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *RosettaOptions) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *RosettaOptions) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetSize

`func (o *RosettaOptions) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RosettaOptions) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RosettaOptions) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *RosettaOptions) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMemo

`func (o *RosettaOptions) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *RosettaOptions) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *RosettaOptions) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *RosettaOptions) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetNumberOfCycles

`func (o *RosettaOptions) GetNumberOfCycles() int32`

GetNumberOfCycles returns the NumberOfCycles field if non-nil, zero value otherwise.

### GetNumberOfCyclesOk

`func (o *RosettaOptions) GetNumberOfCyclesOk() (*int32, bool)`

GetNumberOfCyclesOk returns a tuple with the NumberOfCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCycles

`func (o *RosettaOptions) SetNumberOfCycles(v int32)`

SetNumberOfCycles sets NumberOfCycles field to given value.

### HasNumberOfCycles

`func (o *RosettaOptions) HasNumberOfCycles() bool`

HasNumberOfCycles returns a boolean if a field has been set.

### GetContractAddress

`func (o *RosettaOptions) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *RosettaOptions) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *RosettaOptions) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *RosettaOptions) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetContractName

`func (o *RosettaOptions) GetContractName() string`

GetContractName returns the ContractName field if non-nil, zero value otherwise.

### GetContractNameOk

`func (o *RosettaOptions) GetContractNameOk() (*string, bool)`

GetContractNameOk returns a tuple with the ContractName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractName

`func (o *RosettaOptions) SetContractName(v string)`

SetContractName sets ContractName field to given value.

### HasContractName

`func (o *RosettaOptions) HasContractName() bool`

HasContractName returns a boolean if a field has been set.

### GetBurnBlockHeight

`func (o *RosettaOptions) GetBurnBlockHeight() int32`

GetBurnBlockHeight returns the BurnBlockHeight field if non-nil, zero value otherwise.

### GetBurnBlockHeightOk

`func (o *RosettaOptions) GetBurnBlockHeightOk() (*int32, bool)`

GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHeight

`func (o *RosettaOptions) SetBurnBlockHeight(v int32)`

SetBurnBlockHeight sets BurnBlockHeight field to given value.

### HasBurnBlockHeight

`func (o *RosettaOptions) HasBurnBlockHeight() bool`

HasBurnBlockHeight returns a boolean if a field has been set.

### GetRewardCycleId

`func (o *RosettaOptions) GetRewardCycleId() int32`

GetRewardCycleId returns the RewardCycleId field if non-nil, zero value otherwise.

### GetRewardCycleIdOk

`func (o *RosettaOptions) GetRewardCycleIdOk() (*int32, bool)`

GetRewardCycleIdOk returns a tuple with the RewardCycleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardCycleId

`func (o *RosettaOptions) SetRewardCycleId(v int32)`

SetRewardCycleId sets RewardCycleId field to given value.

### HasRewardCycleId

`func (o *RosettaOptions) HasRewardCycleId() bool`

HasRewardCycleId returns a boolean if a field has been set.

### GetDelegateTo

`func (o *RosettaOptions) GetDelegateTo() string`

GetDelegateTo returns the DelegateTo field if non-nil, zero value otherwise.

### GetDelegateToOk

`func (o *RosettaOptions) GetDelegateToOk() (*string, bool)`

GetDelegateToOk returns a tuple with the DelegateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateTo

`func (o *RosettaOptions) SetDelegateTo(v string)`

SetDelegateTo sets DelegateTo field to given value.

### HasDelegateTo

`func (o *RosettaOptions) HasDelegateTo() bool`

HasDelegateTo returns a boolean if a field has been set.

### GetPoxAddr

`func (o *RosettaOptions) GetPoxAddr() string`

GetPoxAddr returns the PoxAddr field if non-nil, zero value otherwise.

### GetPoxAddrOk

`func (o *RosettaOptions) GetPoxAddrOk() (*string, bool)`

GetPoxAddrOk returns a tuple with the PoxAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoxAddr

`func (o *RosettaOptions) SetPoxAddr(v string)`

SetPoxAddr sets PoxAddr field to given value.

### HasPoxAddr

`func (o *RosettaOptions) HasPoxAddr() bool`

HasPoxAddr returns a boolean if a field has been set.

### GetSignerKey

`func (o *RosettaOptions) GetSignerKey() string`

GetSignerKey returns the SignerKey field if non-nil, zero value otherwise.

### GetSignerKeyOk

`func (o *RosettaOptions) GetSignerKeyOk() (*string, bool)`

GetSignerKeyOk returns a tuple with the SignerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerKey

`func (o *RosettaOptions) SetSignerKey(v string)`

SetSignerKey sets SignerKey field to given value.

### HasSignerKey

`func (o *RosettaOptions) HasSignerKey() bool`

HasSignerKey returns a boolean if a field has been set.

### GetSignerPrivateKey

`func (o *RosettaOptions) GetSignerPrivateKey() string`

GetSignerPrivateKey returns the SignerPrivateKey field if non-nil, zero value otherwise.

### GetSignerPrivateKeyOk

`func (o *RosettaOptions) GetSignerPrivateKeyOk() (*string, bool)`

GetSignerPrivateKeyOk returns a tuple with the SignerPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerPrivateKey

`func (o *RosettaOptions) SetSignerPrivateKey(v string)`

SetSignerPrivateKey sets SignerPrivateKey field to given value.

### HasSignerPrivateKey

`func (o *RosettaOptions) HasSignerPrivateKey() bool`

HasSignerPrivateKey returns a boolean if a field has been set.

### GetSignerSignature

`func (o *RosettaOptions) GetSignerSignature() string`

GetSignerSignature returns the SignerSignature field if non-nil, zero value otherwise.

### GetSignerSignatureOk

`func (o *RosettaOptions) GetSignerSignatureOk() (*string, bool)`

GetSignerSignatureOk returns a tuple with the SignerSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerSignature

`func (o *RosettaOptions) SetSignerSignature(v string)`

SetSignerSignature sets SignerSignature field to given value.

### HasSignerSignature

`func (o *RosettaOptions) HasSignerSignature() bool`

HasSignerSignature returns a boolean if a field has been set.

### GetPoxMaxAmount

`func (o *RosettaOptions) GetPoxMaxAmount() string`

GetPoxMaxAmount returns the PoxMaxAmount field if non-nil, zero value otherwise.

### GetPoxMaxAmountOk

`func (o *RosettaOptions) GetPoxMaxAmountOk() (*string, bool)`

GetPoxMaxAmountOk returns a tuple with the PoxMaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoxMaxAmount

`func (o *RosettaOptions) SetPoxMaxAmount(v string)`

SetPoxMaxAmount sets PoxMaxAmount field to given value.

### HasPoxMaxAmount

`func (o *RosettaOptions) HasPoxMaxAmount() bool`

HasPoxMaxAmount returns a boolean if a field has been set.

### GetPoxAuthId

`func (o *RosettaOptions) GetPoxAuthId() string`

GetPoxAuthId returns the PoxAuthId field if non-nil, zero value otherwise.

### GetPoxAuthIdOk

`func (o *RosettaOptions) GetPoxAuthIdOk() (*string, bool)`

GetPoxAuthIdOk returns a tuple with the PoxAuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoxAuthId

`func (o *RosettaOptions) SetPoxAuthId(v string)`

SetPoxAuthId sets PoxAuthId field to given value.

### HasPoxAuthId

`func (o *RosettaOptions) HasPoxAuthId() bool`

HasPoxAuthId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


