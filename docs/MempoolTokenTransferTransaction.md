# MempoolTokenTransferTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxId** | **string** | Transaction ID | 
**Nonce** | **int32** | Used for ordering the transactions originating from and paying from an account. The nonce ensures that a transaction is processed at most once. The nonce counts the number of times an account&#39;s owner(s) have authorized a transaction. The first transaction from an account will have a nonce value equal to 0, the second will have a nonce value equal to 1, and so on. | 
**FeeRate** | **string** | Transaction fee as Integer string (64-bit unsigned integer). | 
**SenderAddress** | **string** | Address of the transaction initiator | 
**SponsorNonce** | Pointer to **int32** |  | [optional] 
**Sponsored** | **bool** | Denotes whether the originating account is the same as the paying account | 
**SponsorAddress** | Pointer to **string** |  | [optional] 
**PostConditionMode** | **string** |  | 
**PostConditions** | [**[]PostCondition**](PostCondition.md) |  | 
**AnchorMode** | **string** | &#x60;on_chain_only&#x60;: the transaction MUST be included in an anchored block, &#x60;off_chain_only&#x60;: the transaction MUST be included in a microblock, &#x60;any&#x60;: the leader can choose where to include the transaction. | 
**TxStatus** | **string** | Status of the transaction | 
**ReceiptTime** | **float32** | A unix timestamp (in seconds) indicating when the transaction broadcast was received by the node. | 
**ReceiptTimeIso** | **string** | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when the transaction broadcast was received by the node. | 
**TxType** | **string** |  | 
**TokenTransfer** | [**TokenTransferTransactionMetadataTokenTransfer**](TokenTransferTransactionMetadataTokenTransfer.md) |  | 

## Methods

### NewMempoolTokenTransferTransaction

`func NewMempoolTokenTransferTransaction(txId string, nonce int32, feeRate string, senderAddress string, sponsored bool, postConditionMode string, postConditions []PostCondition, anchorMode string, txStatus string, receiptTime float32, receiptTimeIso string, txType string, tokenTransfer TokenTransferTransactionMetadataTokenTransfer, ) *MempoolTokenTransferTransaction`

NewMempoolTokenTransferTransaction instantiates a new MempoolTokenTransferTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMempoolTokenTransferTransactionWithDefaults

`func NewMempoolTokenTransferTransactionWithDefaults() *MempoolTokenTransferTransaction`

NewMempoolTokenTransferTransactionWithDefaults instantiates a new MempoolTokenTransferTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxId

`func (o *MempoolTokenTransferTransaction) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *MempoolTokenTransferTransaction) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *MempoolTokenTransferTransaction) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetNonce

`func (o *MempoolTokenTransferTransaction) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *MempoolTokenTransferTransaction) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *MempoolTokenTransferTransaction) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetFeeRate

`func (o *MempoolTokenTransferTransaction) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *MempoolTokenTransferTransaction) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *MempoolTokenTransferTransaction) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.


### GetSenderAddress

`func (o *MempoolTokenTransferTransaction) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *MempoolTokenTransferTransaction) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *MempoolTokenTransferTransaction) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.


### GetSponsorNonce

`func (o *MempoolTokenTransferTransaction) GetSponsorNonce() int32`

GetSponsorNonce returns the SponsorNonce field if non-nil, zero value otherwise.

### GetSponsorNonceOk

`func (o *MempoolTokenTransferTransaction) GetSponsorNonceOk() (*int32, bool)`

GetSponsorNonceOk returns a tuple with the SponsorNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorNonce

`func (o *MempoolTokenTransferTransaction) SetSponsorNonce(v int32)`

SetSponsorNonce sets SponsorNonce field to given value.

### HasSponsorNonce

`func (o *MempoolTokenTransferTransaction) HasSponsorNonce() bool`

HasSponsorNonce returns a boolean if a field has been set.

### GetSponsored

`func (o *MempoolTokenTransferTransaction) GetSponsored() bool`

GetSponsored returns the Sponsored field if non-nil, zero value otherwise.

### GetSponsoredOk

`func (o *MempoolTokenTransferTransaction) GetSponsoredOk() (*bool, bool)`

GetSponsoredOk returns a tuple with the Sponsored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsored

`func (o *MempoolTokenTransferTransaction) SetSponsored(v bool)`

SetSponsored sets Sponsored field to given value.


### GetSponsorAddress

`func (o *MempoolTokenTransferTransaction) GetSponsorAddress() string`

GetSponsorAddress returns the SponsorAddress field if non-nil, zero value otherwise.

### GetSponsorAddressOk

`func (o *MempoolTokenTransferTransaction) GetSponsorAddressOk() (*string, bool)`

GetSponsorAddressOk returns a tuple with the SponsorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorAddress

`func (o *MempoolTokenTransferTransaction) SetSponsorAddress(v string)`

SetSponsorAddress sets SponsorAddress field to given value.

### HasSponsorAddress

`func (o *MempoolTokenTransferTransaction) HasSponsorAddress() bool`

HasSponsorAddress returns a boolean if a field has been set.

### GetPostConditionMode

`func (o *MempoolTokenTransferTransaction) GetPostConditionMode() string`

GetPostConditionMode returns the PostConditionMode field if non-nil, zero value otherwise.

### GetPostConditionModeOk

`func (o *MempoolTokenTransferTransaction) GetPostConditionModeOk() (*string, bool)`

GetPostConditionModeOk returns a tuple with the PostConditionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostConditionMode

`func (o *MempoolTokenTransferTransaction) SetPostConditionMode(v string)`

SetPostConditionMode sets PostConditionMode field to given value.


### GetPostConditions

`func (o *MempoolTokenTransferTransaction) GetPostConditions() []PostCondition`

GetPostConditions returns the PostConditions field if non-nil, zero value otherwise.

### GetPostConditionsOk

`func (o *MempoolTokenTransferTransaction) GetPostConditionsOk() (*[]PostCondition, bool)`

GetPostConditionsOk returns a tuple with the PostConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostConditions

`func (o *MempoolTokenTransferTransaction) SetPostConditions(v []PostCondition)`

SetPostConditions sets PostConditions field to given value.


### GetAnchorMode

`func (o *MempoolTokenTransferTransaction) GetAnchorMode() string`

GetAnchorMode returns the AnchorMode field if non-nil, zero value otherwise.

### GetAnchorModeOk

`func (o *MempoolTokenTransferTransaction) GetAnchorModeOk() (*string, bool)`

GetAnchorModeOk returns a tuple with the AnchorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorMode

`func (o *MempoolTokenTransferTransaction) SetAnchorMode(v string)`

SetAnchorMode sets AnchorMode field to given value.


### GetTxStatus

`func (o *MempoolTokenTransferTransaction) GetTxStatus() string`

GetTxStatus returns the TxStatus field if non-nil, zero value otherwise.

### GetTxStatusOk

`func (o *MempoolTokenTransferTransaction) GetTxStatusOk() (*string, bool)`

GetTxStatusOk returns a tuple with the TxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxStatus

`func (o *MempoolTokenTransferTransaction) SetTxStatus(v string)`

SetTxStatus sets TxStatus field to given value.


### GetReceiptTime

`func (o *MempoolTokenTransferTransaction) GetReceiptTime() float32`

GetReceiptTime returns the ReceiptTime field if non-nil, zero value otherwise.

### GetReceiptTimeOk

`func (o *MempoolTokenTransferTransaction) GetReceiptTimeOk() (*float32, bool)`

GetReceiptTimeOk returns a tuple with the ReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptTime

`func (o *MempoolTokenTransferTransaction) SetReceiptTime(v float32)`

SetReceiptTime sets ReceiptTime field to given value.


### GetReceiptTimeIso

`func (o *MempoolTokenTransferTransaction) GetReceiptTimeIso() string`

GetReceiptTimeIso returns the ReceiptTimeIso field if non-nil, zero value otherwise.

### GetReceiptTimeIsoOk

`func (o *MempoolTokenTransferTransaction) GetReceiptTimeIsoOk() (*string, bool)`

GetReceiptTimeIsoOk returns a tuple with the ReceiptTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptTimeIso

`func (o *MempoolTokenTransferTransaction) SetReceiptTimeIso(v string)`

SetReceiptTimeIso sets ReceiptTimeIso field to given value.


### GetTxType

`func (o *MempoolTokenTransferTransaction) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *MempoolTokenTransferTransaction) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *MempoolTokenTransferTransaction) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetTokenTransfer

`func (o *MempoolTokenTransferTransaction) GetTokenTransfer() TokenTransferTransactionMetadataTokenTransfer`

GetTokenTransfer returns the TokenTransfer field if non-nil, zero value otherwise.

### GetTokenTransferOk

`func (o *MempoolTokenTransferTransaction) GetTokenTransferOk() (*TokenTransferTransactionMetadataTokenTransfer, bool)`

GetTokenTransferOk returns a tuple with the TokenTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTransfer

`func (o *MempoolTokenTransferTransaction) SetTokenTransfer(v TokenTransferTransactionMetadataTokenTransfer)`

SetTokenTransfer sets TokenTransfer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


