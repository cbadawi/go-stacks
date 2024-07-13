# go-stacks

## Introduction

`go-stacks` is a golang sdk for interacting with the stacks blockchain.
This SDK was generated using the Hiro's OpenAPI specification and modified where needed.

### Requirements

The SDK requires **Go version 1.22 or above**.

## Building

### Install Dependencies

Resolve all the SDK dependencies, using the `go get` command.

## TODOs

- [ ] implement (de)serializing transactions

## Tests
```
go test ./stacks/api
```
## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cbadawi/go-stacks/stacks/api"
	"github.com/cbadawi/go-stacks/stacks/config"
)

func main() {
	principal := "SM3KNVZS30WM7F89SXKVVFY4SN9RMPZZ9FX929N0V" // string | Stacks address or a Contract identifier
	unanchored := false                                      // bool | Include transaction data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
	untilBlock := "400000"                                   //"60000"                                   // string | returned data representing the state up until that point in time, rather than the current block. (optional)

	configuration := config.NewConfiguration()
	configuration.Logger.Verbose = true // defaults to false
	// configuration.Host = "api.testnet.hiro.so" // can define a custom host
	apiClient := api.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountBalance(context.Background(), principal).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountBalance`: AddressBalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountBalance`: %v\n", resp.Stx.Balance)
}
```

## Documentation for API Endpoints



Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsAPI* | [**GetAccountAssets**](docs/AccountsAPI.md#getaccountassets) | **Get** /extended/v1/address/{principal}/assets | Get account assets
*AccountsAPI* | [**GetAccountBalance**](docs/AccountsAPI.md#getaccountbalance) | **Get** /extended/v1/address/{principal}/balances | Get account balances
*AccountsAPI* | [**GetAccountInbound**](docs/AccountsAPI.md#getaccountinbound) | **Get** /extended/v1/address/{principal}/stx_inbound | Get inbound STX transfers
*AccountsAPI* | [**GetAccountInfo**](docs/AccountsAPI.md#getaccountinfo) | **Get** /v2/accounts/{principal} | Get account info
*AccountsAPI* | [**GetAccountNonces**](docs/AccountsAPI.md#getaccountnonces) | **Get** /extended/v1/address/{principal}/nonces | Get the latest nonce used by an account
*AccountsAPI* | [**GetAccountStxBalance**](docs/AccountsAPI.md#getaccountstxbalance) | **Get** /extended/v1/address/{principal}/stx | Get account STX balance
*AccountsAPI* | [**GetAccountTransactions**](docs/AccountsAPI.md#getaccounttransactions) | **Get** /extended/v1/address/{principal}/transactions | Get account transactions
*AccountsAPI* | [**GetAccountTransactionsWithTransfers**](docs/AccountsAPI.md#getaccounttransactionswithtransfers) | **Get** /extended/v1/address/{principal}/transactions_with_transfers | Get account transactions including STX transfers for each transaction.
*AccountsAPI* | [**GetSingleTransactionWithTransfers**](docs/AccountsAPI.md#getsingletransactionwithtransfers) | **Get** /extended/v1/address/{principal}/{tx_id}/with_transfers | Get account transaction information for specific transaction
*BlocksAPI* | [**GetAverageBlockTimes**](docs/BlocksAPI.md#getaverageblocktimes) | **Get** /extended/v2/blocks/average-times | Get average block times
*BlocksAPI* | [**GetBlock**](docs/BlocksAPI.md#getblock) | **Get** /extended/v2/blocks/{height_or_hash} | Get block
*BlocksAPI* | [**GetBlockByBurnBlockHash**](docs/BlocksAPI.md#getblockbyburnblockhash) | **Get** /extended/v1/block/by_burn_block_hash/{burn_block_hash} | Get block by burnchain block hash
*BlocksAPI* | [**GetBlockByBurnBlockHeight**](docs/BlocksAPI.md#getblockbyburnblockheight) | **Get** /extended/v1/block/by_burn_block_height/{burn_block_height} | Get block by burnchain height
*BlocksAPI* | [**GetBlockByHash**](docs/BlocksAPI.md#getblockbyhash) | **Get** /extended/v1/block/{hash} | Get block by hash
*BlocksAPI* | [**GetBlockByHeight**](docs/BlocksAPI.md#getblockbyheight) | **Get** /extended/v1/block/by_height/{height} | Get block by height
*BlocksAPI* | [**GetBlockList**](docs/BlocksAPI.md#getblocklist) | **Get** /extended/v1/block | Get recent blocks
*BlocksAPI* | [**GetBlocks**](docs/BlocksAPI.md#getblocks) | **Get** /extended/v2/blocks | Get blocks
*BlocksAPI* | [**GetBlocksByBurnBlock**](docs/BlocksAPI.md#getblocksbyburnblock) | **Get** /extended/v2/burn-blocks/{height_or_hash}/blocks | Get blocks by burn block
*BurnBlocksAPI* | [**GetBurnBlock**](docs/BurnBlocksAPI.md#getburnblock) | **Get** /extended/v2/burn-blocks/{height_or_hash} | Get burn block
*BurnBlocksAPI* | [**GetBurnBlocks**](docs/BurnBlocksAPI.md#getburnblocks) | **Get** /extended/v2/burn-blocks | Get burn blocks
*FaucetsAPI* | [**RunFaucetBtc**](docs/FaucetsAPI.md#runfaucetbtc) | **Post** /extended/v1/faucets/btc | Add testnet BTC tokens to address
*FaucetsAPI* | [**RunFaucetStx**](docs/FaucetsAPI.md#runfaucetstx) | **Post** /extended/v1/faucets/stx | Get STX testnet tokens
*FeesAPI* | [**FetchFeeRate**](docs/FeesAPI.md#fetchfeerate) | **Post** /extended/v1/fee_rate | Fetch fee rate
*FeesAPI* | [**GetFeeTransfer**](docs/FeesAPI.md#getfeetransfer) | **Get** /v2/fees/transfer | Get estimated fee
*FeesAPI* | [**PostFeeTransaction**](docs/FeesAPI.md#postfeetransaction) | **Post** /v2/fees/transaction | Get approximate fees for a given transaction
*InfoAPI* | [**GetCoreApiInfo**](docs/InfoAPI.md#getcoreapiinfo) | **Get** /v2/info | Get Core API info
*InfoAPI* | [**GetNetworkBlockTimeByNetwork**](docs/InfoAPI.md#getnetworkblocktimebynetwork) | **Get** /extended/v1/info/network_block_time/{network} | Get a given network&#39;s target block time
*InfoAPI* | [**GetNetworkBlockTimes**](docs/InfoAPI.md#getnetworkblocktimes) | **Get** /extended/v1/info/network_block_times | Get the network target block time
*InfoAPI* | [**GetPoxInfo**](docs/InfoAPI.md#getpoxinfo) | **Get** /v2/pox | Get Proof-of-Transfer details
*InfoAPI* | [**GetStatus**](docs/InfoAPI.md#getstatus) | **Get** /extended | API status
*InfoAPI* | [**GetStxSupply**](docs/InfoAPI.md#getstxsupply) | **Get** /extended/v1/stx_supply | Get total and unlocked STX supply
*InfoAPI* | [**GetStxSupplyCirculatingPlain**](docs/InfoAPI.md#getstxsupplycirculatingplain) | **Get** /extended/v1/stx_supply/circulating/plain | Get circulating STX supply in plain text format
*InfoAPI* | [**GetStxSupplyTotalSupplyPlain**](docs/InfoAPI.md#getstxsupplytotalsupplyplain) | **Get** /extended/v1/stx_supply/total/plain | Get total STX supply in plain text format
*InfoAPI* | [**GetTotalStxSupplyLegacyFormat**](docs/InfoAPI.md#gettotalstxsupplylegacyformat) | **Get** /extended/v1/stx_supply/legacy_format | Get total and unlocked STX supply (results formatted the same as the legacy 1.0 API)
*MempoolAPI* | [**GetMempoolFeePriorities**](docs/MempoolAPI.md#getmempoolfeepriorities) | **Get** /extended/v2/mempool/fees | Get mempool transaction fee priorities
*MicroblocksAPI* | [**GetMicroblockByHash**](docs/MicroblocksAPI.md#getmicroblockbyhash) | **Get** /extended/v1/microblock/{hash} | Get microblock
*MicroblocksAPI* | [**GetMicroblockList**](docs/MicroblocksAPI.md#getmicroblocklist) | **Get** /extended/v1/microblock | Get recent microblocks
*MicroblocksAPI* | [**GetUnanchoredTxs**](docs/MicroblocksAPI.md#getunanchoredtxs) | **Get** /extended/v1/microblock/unanchored/txs | Get the list of current transactions that belong to unanchored microblocks
*NamesAPI* | [**FetchSubdomainsListForName**](docs/NamesAPI.md#fetchsubdomainslistforname) | **Get** /v1/names/{name}/subdomains | Get Name Subdomains
*NamesAPI* | [**FetchZoneFile**](docs/NamesAPI.md#fetchzonefile) | **Get** /v1/names/{name}/zonefile | Get Zone File
*NamesAPI* | [**GetAllNames**](docs/NamesAPI.md#getallnames) | **Get** /v1/names | Get All Names
*NamesAPI* | [**GetAllNamespaces**](docs/NamesAPI.md#getallnamespaces) | **Get** /v1/namespaces | Get All Namespaces
*NamesAPI* | [**GetHistoricalZoneFile**](docs/NamesAPI.md#gethistoricalzonefile) | **Get** /v1/names/{name}/zonefile/{zoneFileHash} | Get Historical Zone File
*NamesAPI* | [**GetNameInfo**](docs/NamesAPI.md#getnameinfo) | **Get** /v1/names/{name} | Get Name Details
*NamesAPI* | [**GetNamePrice**](docs/NamesAPI.md#getnameprice) | **Get** /v2/prices/names/{name} | Get Name Price
*NamesAPI* | [**GetNamesOwnedByAddress**](docs/NamesAPI.md#getnamesownedbyaddress) | **Get** /v1/addresses/{blockchain}/{address} | Get Names Owned by Address
*NamesAPI* | [**GetNamespaceNames**](docs/NamesAPI.md#getnamespacenames) | **Get** /v1/namespaces/{tld}/names | Get Namespace Names
*NamesAPI* | [**GetNamespacePrice**](docs/NamesAPI.md#getnamespaceprice) | **Get** /v2/prices/namespaces/{tld} | Get Namespace Price
*NonFungibleTokensAPI* | [**GetNftHistory**](docs/NonFungibleTokensAPI.md#getnfthistory) | **Get** /extended/v1/tokens/nft/history | Non-Fungible Token history
*NonFungibleTokensAPI* | [**GetNftHoldings**](docs/NonFungibleTokensAPI.md#getnftholdings) | **Get** /extended/v1/tokens/nft/holdings | Non-Fungible Token holdings
*NonFungibleTokensAPI* | [**GetNftMints**](docs/NonFungibleTokensAPI.md#getnftmints) | **Get** /extended/v1/tokens/nft/mints | Non-Fungible Token mints
*ProofOfTransferAPI* | [**GetPoxCycle**](docs/ProofOfTransferAPI.md#getpoxcycle) | **Get** /extended/v2/pox/cycles/{cycle_number} | Get PoX cycle
*ProofOfTransferAPI* | [**GetPoxCycleSigner**](docs/ProofOfTransferAPI.md#getpoxcyclesigner) | **Get** /extended/v2/pox/cycles/{cycle_number}/signers/{signer_key} | Get signer in PoX cycle
*ProofOfTransferAPI* | [**GetPoxCycleSignerStackers**](docs/ProofOfTransferAPI.md#getpoxcyclesignerstackers) | **Get** /extended/v2/pox/cycles/{cycle_number}/signers/{signer_key}/stackers | Get stackers for signer in PoX cycle
*ProofOfTransferAPI* | [**GetPoxCycleSigners**](docs/ProofOfTransferAPI.md#getpoxcyclesigners) | **Get** /extended/v2/pox/cycles/{cycle_number}/signers | Get signers in PoX cycle
*ProofOfTransferAPI* | [**GetPoxCycles**](docs/ProofOfTransferAPI.md#getpoxcycles) | **Get** /extended/v2/pox/cycles | Get PoX cycles
*RosettaAPI* | [**RosettaAccountBalance**](docs/RosettaAPI.md#rosettaaccountbalance) | **Post** /rosetta/v1/account/balance | Get an Account Balance
*RosettaAPI* | [**RosettaBlock**](docs/RosettaAPI.md#rosettablock) | **Post** /rosetta/v1/block | Get a Block
*RosettaAPI* | [**RosettaBlockTransaction**](docs/RosettaAPI.md#rosettablocktransaction) | **Post** /rosetta/v1/block/transaction | Get a Block Transaction
*RosettaAPI* | [**RosettaConstructionCombine**](docs/RosettaAPI.md#rosettaconstructioncombine) | **Post** /rosetta/v1/construction/combine | Create Network Transaction from Signatures
*RosettaAPI* | [**RosettaConstructionDerive**](docs/RosettaAPI.md#rosettaconstructionderive) | **Post** /rosetta/v1/construction/derive | Derive an AccountIdentifier from a PublicKey
*RosettaAPI* | [**RosettaConstructionHash**](docs/RosettaAPI.md#rosettaconstructionhash) | **Post** /rosetta/v1/construction/hash | Get the Hash of a Signed Transaction
*RosettaAPI* | [**RosettaConstructionMetadata**](docs/RosettaAPI.md#rosettaconstructionmetadata) | **Post** /rosetta/v1/construction/metadata | Get Metadata for Transaction Construction
*RosettaAPI* | [**RosettaConstructionParse**](docs/RosettaAPI.md#rosettaconstructionparse) | **Post** /rosetta/v1/construction/parse | Parse a Transaction
*RosettaAPI* | [**RosettaConstructionPayloads**](docs/RosettaAPI.md#rosettaconstructionpayloads) | **Post** /rosetta/v1/construction/payloads | Generate an Unsigned Transaction and Signing Payloads
*RosettaAPI* | [**RosettaConstructionPreprocess**](docs/RosettaAPI.md#rosettaconstructionpreprocess) | **Post** /rosetta/v1/construction/preprocess | Create a Request to Fetch Metadata
*RosettaAPI* | [**RosettaConstructionSubmit**](docs/RosettaAPI.md#rosettaconstructionsubmit) | **Post** /rosetta/v1/construction/submit | Submit a Signed Transaction
*RosettaAPI* | [**RosettaMempool**](docs/RosettaAPI.md#rosettamempool) | **Post** /rosetta/v1/mempool | Get All Mempool Transactions
*RosettaAPI* | [**RosettaMempoolTransaction**](docs/RosettaAPI.md#rosettamempooltransaction) | **Post** /rosetta/v1/mempool/transaction | Get a Mempool Transaction
*RosettaAPI* | [**RosettaNetworkList**](docs/RosettaAPI.md#rosettanetworklist) | **Post** /rosetta/v1/network/list | Get List of Available Networks
*RosettaAPI* | [**RosettaNetworkOptions**](docs/RosettaAPI.md#rosettanetworkoptions) | **Post** /rosetta/v1/network/options | Get Network Options
*RosettaAPI* | [**RosettaNetworkStatus**](docs/RosettaAPI.md#rosettanetworkstatus) | **Post** /rosetta/v1/network/status | Get Network Status
*SearchAPI* | [**SearchById**](docs/SearchAPI.md#searchbyid) | **Get** /extended/v1/search/{id} | Search
*SmartContractsAPI* | [**CallReadOnlyFunction**](docs/SmartContractsAPI.md#callreadonlyfunction) | **Post** /v2/contracts/call-read/{contract_address}/{contract_name}/{function_name} | Call read-only function
*SmartContractsAPI* | [**GetContractById**](docs/SmartContractsAPI.md#getcontractbyid) | **Get** /extended/v1/contract/{contract_id} | Get contract info
*SmartContractsAPI* | [**GetContractDataMapEntry**](docs/SmartContractsAPI.md#getcontractdatamapentry) | **Post** /v2/map_entry/{contract_address}/{contract_name}/{map_name} | Get specific data-map inside a contract
*SmartContractsAPI* | [**GetContractEventsById**](docs/SmartContractsAPI.md#getcontracteventsbyid) | **Get** /extended/v1/contract/{contract_id}/events | Get contract events
*SmartContractsAPI* | [**GetContractInterface**](docs/SmartContractsAPI.md#getcontractinterface) | **Get** /v2/contracts/interface/{contract_address}/{contract_name} | Get contract interface
*SmartContractsAPI* | [**GetContractSource**](docs/SmartContractsAPI.md#getcontractsource) | **Get** /v2/contracts/source/{contract_address}/{contract_name} | Get contract source
*SmartContractsAPI* | [**GetContractsByTrait**](docs/SmartContractsAPI.md#getcontractsbytrait) | **Get** /extended/v1/contract/by_trait | Get contracts by trait
*SmartContractsAPI* | [**GetSmartContractsStatus**](docs/SmartContractsAPI.md#getsmartcontractsstatus) | **Get** /extended/v2/smart-contracts/status | Get smart contracts status
*StackingAPI* | [**GetPoolDelegations**](docs/StackingAPI.md#getpooldelegations) | **Get** /extended/beta/stacking/{pool_principal}/delegations | Stacking pool members
*StackingRewardsAPI* | [**GetBurnchainRewardList**](docs/StackingRewardsAPI.md#getburnchainrewardlist) | **Get** /extended/v1/burnchain/rewards | Get recent burnchain reward recipients
*StackingRewardsAPI* | [**GetBurnchainRewardListByAddress**](docs/StackingRewardsAPI.md#getburnchainrewardlistbyaddress) | **Get** /extended/v1/burnchain/rewards/{address} | Get recent burnchain reward for the given recipient
*StackingRewardsAPI* | [**GetBurnchainRewardSlotHolders**](docs/StackingRewardsAPI.md#getburnchainrewardslotholders) | **Get** /extended/v1/burnchain/reward_slot_holders | Get recent reward slot holders
*StackingRewardsAPI* | [**GetBurnchainRewardSlotHoldersByAddress**](docs/StackingRewardsAPI.md#getburnchainrewardslotholdersbyaddress) | **Get** /extended/v1/burnchain/reward_slot_holders/{address} | Get recent reward slot holder entries for the given address
*StackingRewardsAPI* | [**GetBurnchainRewardsTotalByAddress**](docs/StackingRewardsAPI.md#getburnchainrewardstotalbyaddress) | **Get** /extended/v1/burnchain/rewards/{address}/total | Get total burnchain rewards for the given recipient
*TransactionsAPI* | [**GetAddressMempoolTransactions**](docs/TransactionsAPI.md#getaddressmempooltransactions) | **Get** /extended/v1/address/{address}/mempool | Transactions for address
*TransactionsAPI* | [**GetAddressTransactionEvents**](docs/TransactionsAPI.md#getaddresstransactionevents) | **Get** /extended/v2/addresses/{address}/transactions/{tx_id}/events | Get events for an address transaction
*TransactionsAPI* | [**GetAddressTransactions**](docs/TransactionsAPI.md#getaddresstransactions) | **Get** /extended/v2/addresses/{address}/transactions | Get address transactions
*TransactionsAPI* | [**GetDroppedMempoolTransactionList**](docs/TransactionsAPI.md#getdroppedmempooltransactionlist) | **Get** /extended/v1/tx/mempool/dropped | Get dropped mempool transactions
*TransactionsAPI* | [**GetFilteredEvents**](docs/TransactionsAPI.md#getfilteredevents) | **Get** /extended/v1/tx/events | Transaction Events
*TransactionsAPI* | [**GetMempoolTransactionList**](docs/TransactionsAPI.md#getmempooltransactionlist) | **Get** /extended/v1/tx/mempool | Get mempool transactions
*TransactionsAPI* | [**GetMempoolTransactionStats**](docs/TransactionsAPI.md#getmempooltransactionstats) | **Get** /extended/v1/tx/mempool/stats | Get statistics for mempool transactions
*TransactionsAPI* | [**GetRawTransactionById**](docs/TransactionsAPI.md#getrawtransactionbyid) | **Get** /extended/v1/tx/{tx_id}/raw | Get Raw Transaction
*TransactionsAPI* | [**GetTransactionById**](docs/TransactionsAPI.md#gettransactionbyid) | **Get** /extended/v1/tx/{tx_id} | Get transaction
*TransactionsAPI* | [**GetTransactionList**](docs/TransactionsAPI.md#gettransactionlist) | **Get** /extended/v1/tx | Get recent transactions
*TransactionsAPI* | [**GetTransactionsByBlock**](docs/TransactionsAPI.md#gettransactionsbyblock) | **Get** /extended/v2/blocks/{height_or_hash}/transactions | Get transactions by block
*TransactionsAPI* | [**GetTransactionsByBlockHash**](docs/TransactionsAPI.md#gettransactionsbyblockhash) | **Get** /extended/v1/tx/block/{block_hash} | Transactions by block hash
*TransactionsAPI* | [**GetTransactionsByBlockHeight**](docs/TransactionsAPI.md#gettransactionsbyblockheight) | **Get** /extended/v1/tx/block_height/{height} | Transactions by block height
*TransactionsAPI* | [**GetTxListDetails**](docs/TransactionsAPI.md#gettxlistdetails) | **Get** /extended/v1/tx/multiple | Get list of details for transactions
*TransactionsAPI* | [**PostCoreNodeTransactions**](docs/TransactionsAPI.md#postcorenodetransactions) | **Post** /v2/transactions | Broadcast raw transaction


## Documentation For Models

 - [AbstractTransactionAllOfTxResult](docs/AbstractTransactionAllOfTxResult.md)
 - [AccountDataResponse](docs/AccountDataResponse.md)
 - [AddressAssetsListResponse](docs/AddressAssetsListResponse.md)
 - [AddressBalanceResponse](docs/AddressBalanceResponse.md)
 - [AddressBalanceResponseFungibleTokensValue](docs/AddressBalanceResponseFungibleTokensValue.md)
 - [AddressBalanceResponseNonFungibleTokensValue](docs/AddressBalanceResponseNonFungibleTokensValue.md)
 - [AddressNonces](docs/AddressNonces.md)
 - [AddressSearchResult](docs/AddressSearchResult.md)
 - [AddressSearchResultResult](docs/AddressSearchResultResult.md)
 - [AddressStxBalanceResponse](docs/AddressStxBalanceResponse.md)
 - [AddressStxInboundListResponse](docs/AddressStxInboundListResponse.md)
 - [AddressTokenOfferingLocked](docs/AddressTokenOfferingLocked.md)
 - [AddressTransaction](docs/AddressTransaction.md)
 - [AddressTransactionEvent](docs/AddressTransactionEvent.md)
 - [AddressTransactionEventAnyOf](docs/AddressTransactionEventAnyOf.md)
 - [AddressTransactionEventAnyOf1](docs/AddressTransactionEventAnyOf1.md)
 - [AddressTransactionEventAnyOf1Data](docs/AddressTransactionEventAnyOf1Data.md)
 - [AddressTransactionEventAnyOf2](docs/AddressTransactionEventAnyOf2.md)
 - [AddressTransactionEventAnyOf2Data](docs/AddressTransactionEventAnyOf2Data.md)
 - [AddressTransactionEventAnyOf2DataValue](docs/AddressTransactionEventAnyOf2DataValue.md)
 - [AddressTransactionEventAnyOfData](docs/AddressTransactionEventAnyOfData.md)
 - [AddressTransactionEventListResponse](docs/AddressTransactionEventListResponse.md)
 - [AddressTransactionEvents](docs/AddressTransactionEvents.md)
 - [AddressTransactionEventsStx](docs/AddressTransactionEventsStx.md)
 - [AddressTransactionWithTransfers](docs/AddressTransactionWithTransfers.md)
 - [AddressTransactionWithTransfersFtTransfersInner](docs/AddressTransactionWithTransfersFtTransfersInner.md)
 - [AddressTransactionWithTransfersNftTransfersInner](docs/AddressTransactionWithTransfersNftTransfersInner.md)
 - [AddressTransactionWithTransfersStxTransfersInner](docs/AddressTransactionWithTransfersStxTransfersInner.md)
 - [AddressTransactionsListResponse](docs/AddressTransactionsListResponse.md)
 - [AddressTransactionsListResponseResultsInner](docs/AddressTransactionsListResponseResultsInner.md)
 - [AddressTransactionsV2ListResponse](docs/AddressTransactionsV2ListResponse.md)
 - [AddressTransactionsWithTransfersListResponse](docs/AddressTransactionsWithTransfersListResponse.md)
 - [AddressUnlockSchedule](docs/AddressUnlockSchedule.md)
 - [AverageBlockTimesResponse](docs/AverageBlockTimesResponse.md)
 - [Block](docs/Block.md)
 - [BlockListResponse](docs/BlockListResponse.md)
 - [BlockSearchResult](docs/BlockSearchResult.md)
 - [BlockSearchResultResult](docs/BlockSearchResultResult.md)
 - [BlockSearchResultResultBlockData](docs/BlockSearchResultResultBlockData.md)
 - [BnsError](docs/BnsError.md)
 - [BnsFetchFileZoneResponse](docs/BnsFetchFileZoneResponse.md)
 - [BnsFetchFileZoneResponseAnyOf](docs/BnsFetchFileZoneResponseAnyOf.md)
 - [BnsFetchFileZoneResponseAnyOf1](docs/BnsFetchFileZoneResponseAnyOf1.md)
 - [BnsFetchHistoricalZoneFileResponse](docs/BnsFetchHistoricalZoneFileResponse.md)
 - [BnsFetchHistoricalZoneFileResponseAnyOf](docs/BnsFetchHistoricalZoneFileResponseAnyOf.md)
 - [BnsFetchHistoricalZoneFileResponseAnyOf1](docs/BnsFetchHistoricalZoneFileResponseAnyOf1.md)
 - [BnsGetAllNamespacesResponse](docs/BnsGetAllNamespacesResponse.md)
 - [BnsGetNameInfoResponse](docs/BnsGetNameInfoResponse.md)
 - [BnsGetNamePriceResponse](docs/BnsGetNamePriceResponse.md)
 - [BnsGetNamespacePriceResponse](docs/BnsGetNamespacePriceResponse.md)
 - [BnsNamesOwnByAddressResponse](docs/BnsNamesOwnByAddressResponse.md)
 - [BurnBlock](docs/BurnBlock.md)
 - [BurnBlockListResponse](docs/BurnBlockListResponse.md)
 - [BurnchainReward](docs/BurnchainReward.md)
 - [BurnchainRewardListResponse](docs/BurnchainRewardListResponse.md)
 - [BurnchainRewardSlotHolder](docs/BurnchainRewardSlotHolder.md)
 - [BurnchainRewardSlotHolderListResponse](docs/BurnchainRewardSlotHolderListResponse.md)
 - [BurnchainRewardsTotal](docs/BurnchainRewardsTotal.md)
 - [ChainTip](docs/ChainTip.md)
 - [CoinbaseTransaction](docs/CoinbaseTransaction.md)
 - [CoinbaseTransactionMetadataCoinbasePayload](docs/CoinbaseTransactionMetadataCoinbasePayload.md)
 - [ContractCallTransaction](docs/ContractCallTransaction.md)
 - [ContractCallTransactionMetadataContractCall](docs/ContractCallTransactionMetadataContractCall.md)
 - [ContractCallTransactionMetadataContractCallFunctionArgsInner](docs/ContractCallTransactionMetadataContractCallFunctionArgsInner.md)
 - [ContractInterfaceResponse](docs/ContractInterfaceResponse.md)
 - [ContractListResponse](docs/ContractListResponse.md)
 - [ContractSearchResult](docs/ContractSearchResult.md)
 - [ContractSearchResultResult](docs/ContractSearchResultResult.md)
 - [ContractSearchResultResultTxData](docs/ContractSearchResultResultTxData.md)
 - [ContractSourceResponse](docs/ContractSourceResponse.md)
 - [CoreNodeInfoResponse](docs/CoreNodeInfoResponse.md)
 - [CoreNodePoxResponse](docs/CoreNodePoxResponse.md)
 - [FeeRate](docs/FeeRate.md)
 - [FeeRateRequest](docs/FeeRateRequest.md)
 - [FtBalance](docs/FtBalance.md)
 - [GetBurnBlockHeightOrHashParameter](docs/GetBurnBlockHeightOrHashParameter.md)
 - [GetRawTransactionResult](docs/GetRawTransactionResult.md)
 - [GetStxSupplyLegacyFormatResponse](docs/GetStxSupplyLegacyFormatResponse.md)
 - [GetStxSupplyResponse](docs/GetStxSupplyResponse.md)
 - [InboundStxTransfer](docs/InboundStxTransfer.md)
 - [MapEntryResponse](docs/MapEntryResponse.md)
 - [MempoolCoinbaseTransaction](docs/MempoolCoinbaseTransaction.md)
 - [MempoolContractCallTransaction](docs/MempoolContractCallTransaction.md)
 - [MempoolFeePriorities](docs/MempoolFeePriorities.md)
 - [MempoolFeePrioritiesAll](docs/MempoolFeePrioritiesAll.md)
 - [MempoolPoisonMicroblockTransaction](docs/MempoolPoisonMicroblockTransaction.md)
 - [MempoolSmartContractTransaction](docs/MempoolSmartContractTransaction.md)
 - [MempoolTenureChangeTransaction](docs/MempoolTenureChangeTransaction.md)
 - [MempoolTokenTransferTransaction](docs/MempoolTokenTransferTransaction.md)
 - [MempoolTransaction](docs/MempoolTransaction.md)
 - [MempoolTransactionListResponse](docs/MempoolTransactionListResponse.md)
 - [MempoolTransactionStatsResponse](docs/MempoolTransactionStatsResponse.md)
 - [MempoolTransactionStatsResponseTxAges](docs/MempoolTransactionStatsResponseTxAges.md)
 - [MempoolTransactionStatsResponseTxByteSizes](docs/MempoolTransactionStatsResponseTxByteSizes.md)
 - [MempoolTransactionStatsResponseTxSimpleFeeAverages](docs/MempoolTransactionStatsResponseTxSimpleFeeAverages.md)
 - [MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer](docs/MempoolTransactionStatsResponseTxSimpleFeeAveragesTokenTransfer.md)
 - [MempoolTransactionStatsResponseTxTypeCounts](docs/MempoolTransactionStatsResponseTxTypeCounts.md)
 - [MempoolTxSearchResult](docs/MempoolTxSearchResult.md)
 - [MempoolTxSearchResultResult](docs/MempoolTxSearchResultResult.md)
 - [MempoolTxSearchResultResultTxData](docs/MempoolTxSearchResultResultTxData.md)
 - [Microblock](docs/Microblock.md)
 - [MicroblockListResponse](docs/MicroblockListResponse.md)
 - [NakamotoBlock](docs/NakamotoBlock.md)
 - [NakamotoBlockListResponse](docs/NakamotoBlockListResponse.md)
 - [NetworkBlockTimesResponse](docs/NetworkBlockTimesResponse.md)
 - [NetworkIdentifier](docs/NetworkIdentifier.md)
 - [NetworkIdentifierSubNetworkIdentifier](docs/NetworkIdentifierSubNetworkIdentifier.md)
 - [NetworkIdentifierSubNetworkIdentifierMetadata](docs/NetworkIdentifierSubNetworkIdentifierMetadata.md)
 - [NftBalance](docs/NftBalance.md)
 - [NonFungibleTokenHistoryEvent](docs/NonFungibleTokenHistoryEvent.md)
 - [NonFungibleTokenHistoryEventList](docs/NonFungibleTokenHistoryEventList.md)
 - [NonFungibleTokenHistoryEventWithTxId](docs/NonFungibleTokenHistoryEventWithTxId.md)
 - [NonFungibleTokenHistoryEventWithTxMetadata](docs/NonFungibleTokenHistoryEventWithTxMetadata.md)
 - [NonFungibleTokenHolding](docs/NonFungibleTokenHolding.md)
 - [NonFungibleTokenHoldingWithTxId](docs/NonFungibleTokenHoldingWithTxId.md)
 - [NonFungibleTokenHoldingWithTxIdValue](docs/NonFungibleTokenHoldingWithTxIdValue.md)
 - [NonFungibleTokenHoldingWithTxMetadata](docs/NonFungibleTokenHoldingWithTxMetadata.md)
 - [NonFungibleTokenHoldingsList](docs/NonFungibleTokenHoldingsList.md)
 - [NonFungibleTokenMint](docs/NonFungibleTokenMint.md)
 - [NonFungibleTokenMintList](docs/NonFungibleTokenMintList.md)
 - [NonFungibleTokenMintWithTxId](docs/NonFungibleTokenMintWithTxId.md)
 - [NonFungibleTokenMintWithTxMetadata](docs/NonFungibleTokenMintWithTxMetadata.md)
 - [OtherTransactionIdentifier](docs/OtherTransactionIdentifier.md)
 - [PoisonMicroblockTransaction](docs/PoisonMicroblockTransaction.md)
 - [PoisonMicroblockTransactionMetadataPoisonMicroblock](docs/PoisonMicroblockTransactionMetadataPoisonMicroblock.md)
 - [PoolDelegation](docs/PoolDelegation.md)
 - [PoolDelegationsResponse](docs/PoolDelegationsResponse.md)
 - [PostCondition](docs/PostCondition.md)
 - [PostConditionFungible](docs/PostConditionFungible.md)
 - [PostConditionFungibleAllOfAsset](docs/PostConditionFungibleAllOfAsset.md)
 - [PostConditionNonFungible](docs/PostConditionNonFungible.md)
 - [PostConditionNonFungibleAllOfAssetValue](docs/PostConditionNonFungibleAllOfAssetValue.md)
 - [PostConditionPrincipal](docs/PostConditionPrincipal.md)
 - [PostConditionPrincipalAnyOf](docs/PostConditionPrincipalAnyOf.md)
 - [PostConditionPrincipalAnyOf1](docs/PostConditionPrincipalAnyOf1.md)
 - [PostConditionPrincipalAnyOf2](docs/PostConditionPrincipalAnyOf2.md)
 - [PostConditionStx](docs/PostConditionStx.md)
 - [PostCoreNodeTransactionsError](docs/PostCoreNodeTransactionsError.md)
 - [PoxCycle](docs/PoxCycle.md)
 - [PoxCycleListResponse](docs/PoxCycleListResponse.md)
 - [PoxCycleSignerStackersListResponse](docs/PoxCycleSignerStackersListResponse.md)
 - [PoxCycleSignersListResponse](docs/PoxCycleSignersListResponse.md)
 - [PoxSigner](docs/PoxSigner.md)
 - [PoxStacker](docs/PoxStacker.md)
 - [ReadOnlyFunctionArgs](docs/ReadOnlyFunctionArgs.md)
 - [ReadOnlyFunctionSuccessResponse](docs/ReadOnlyFunctionSuccessResponse.md)
 - [RosettaAccount](docs/RosettaAccount.md)
 - [RosettaAccountBalanceRequest](docs/RosettaAccountBalanceRequest.md)
 - [RosettaAccountBalanceResponse](docs/RosettaAccountBalanceResponse.md)
 - [RosettaAccountBalanceResponseMetadata](docs/RosettaAccountBalanceResponseMetadata.md)
 - [RosettaAccountIdentifier](docs/RosettaAccountIdentifier.md)
 - [RosettaAmount](docs/RosettaAmount.md)
 - [RosettaBlock](docs/RosettaBlock.md)
 - [RosettaBlockIdentifier](docs/RosettaBlockIdentifier.md)
 - [RosettaBlockIdentifierHash](docs/RosettaBlockIdentifierHash.md)
 - [RosettaBlockIdentifierHeight](docs/RosettaBlockIdentifierHeight.md)
 - [RosettaBlockMetadata](docs/RosettaBlockMetadata.md)
 - [RosettaBlockRequest](docs/RosettaBlockRequest.md)
 - [RosettaBlockResponse](docs/RosettaBlockResponse.md)
 - [RosettaBlockTransactionRequest](docs/RosettaBlockTransactionRequest.md)
 - [RosettaBlockTransactionResponse](docs/RosettaBlockTransactionResponse.md)
 - [RosettaCoin](docs/RosettaCoin.md)
 - [RosettaCoinChange](docs/RosettaCoinChange.md)
 - [RosettaCoinChangeCoinIdentifier](docs/RosettaCoinChangeCoinIdentifier.md)
 - [RosettaCoinCoinIdentifier](docs/RosettaCoinCoinIdentifier.md)
 - [RosettaConstructionCombineRequest](docs/RosettaConstructionCombineRequest.md)
 - [RosettaConstructionCombineResponse](docs/RosettaConstructionCombineResponse.md)
 - [RosettaConstructionDeriveRequest](docs/RosettaConstructionDeriveRequest.md)
 - [RosettaConstructionDeriveResponse](docs/RosettaConstructionDeriveResponse.md)
 - [RosettaConstructionHashRequest](docs/RosettaConstructionHashRequest.md)
 - [RosettaConstructionHashResponse](docs/RosettaConstructionHashResponse.md)
 - [RosettaConstructionMetadataRequest](docs/RosettaConstructionMetadataRequest.md)
 - [RosettaConstructionMetadataResponse](docs/RosettaConstructionMetadataResponse.md)
 - [RosettaConstructionMetadataResponseMetadata](docs/RosettaConstructionMetadataResponseMetadata.md)
 - [RosettaConstructionParseRequest](docs/RosettaConstructionParseRequest.md)
 - [RosettaConstructionParseResponse](docs/RosettaConstructionParseResponse.md)
 - [RosettaConstructionPayloadResponse](docs/RosettaConstructionPayloadResponse.md)
 - [RosettaConstructionPayloadsRequest](docs/RosettaConstructionPayloadsRequest.md)
 - [RosettaConstructionPreprocessRequest](docs/RosettaConstructionPreprocessRequest.md)
 - [RosettaConstructionPreprocessResponse](docs/RosettaConstructionPreprocessResponse.md)
 - [RosettaConstructionSubmitRequest](docs/RosettaConstructionSubmitRequest.md)
 - [RosettaConstructionSubmitResponse](docs/RosettaConstructionSubmitResponse.md)
 - [RosettaCurrency](docs/RosettaCurrency.md)
 - [RosettaError](docs/RosettaError.md)
 - [RosettaErrorDetails](docs/RosettaErrorDetails.md)
 - [RosettaErrorNoDetails](docs/RosettaErrorNoDetails.md)
 - [RosettaGenesisBlockIdentifier](docs/RosettaGenesisBlockIdentifier.md)
 - [RosettaMaxFeeAmount](docs/RosettaMaxFeeAmount.md)
 - [RosettaMempoolRequest](docs/RosettaMempoolRequest.md)
 - [RosettaMempoolResponse](docs/RosettaMempoolResponse.md)
 - [RosettaMempoolTransactionRequest](docs/RosettaMempoolTransactionRequest.md)
 - [RosettaMempoolTransactionResponse](docs/RosettaMempoolTransactionResponse.md)
 - [RosettaNetworkListResponse](docs/RosettaNetworkListResponse.md)
 - [RosettaNetworkOptionsResponse](docs/RosettaNetworkOptionsResponse.md)
 - [RosettaNetworkOptionsResponseAllow](docs/RosettaNetworkOptionsResponseAllow.md)
 - [RosettaNetworkOptionsResponseAllowOperationTypesInner](docs/RosettaNetworkOptionsResponseAllowOperationTypesInner.md)
 - [RosettaNetworkOptionsResponseVersion](docs/RosettaNetworkOptionsResponseVersion.md)
 - [RosettaNetworkStatusResponse](docs/RosettaNetworkStatusResponse.md)
 - [RosettaOldestBlockIdentifier](docs/RosettaOldestBlockIdentifier.md)
 - [RosettaOperation](docs/RosettaOperation.md)
 - [RosettaOperationIdentifier](docs/RosettaOperationIdentifier.md)
 - [RosettaOperationStatus](docs/RosettaOperationStatus.md)
 - [RosettaOptions](docs/RosettaOptions.md)
 - [RosettaOptionsRequest](docs/RosettaOptionsRequest.md)
 - [RosettaParentBlockIdentifier](docs/RosettaParentBlockIdentifier.md)
 - [RosettaPartialBlockIdentifier](docs/RosettaPartialBlockIdentifier.md)
 - [RosettaPeers](docs/RosettaPeers.md)
 - [RosettaPublicKey](docs/RosettaPublicKey.md)
 - [RosettaRelatedOperation](docs/RosettaRelatedOperation.md)
 - [RosettaSignature](docs/RosettaSignature.md)
 - [RosettaStatusRequest](docs/RosettaStatusRequest.md)
 - [RosettaSubAccount](docs/RosettaSubAccount.md)
 - [RosettaSyncStatus](docs/RosettaSyncStatus.md)
 - [RosettaTransaction](docs/RosettaTransaction.md)
 - [RosettaTransactionMetadata](docs/RosettaTransactionMetadata.md)
 - [RunFaucetBtc403Response](docs/RunFaucetBtc403Response.md)
 - [RunFaucetBtcRequest](docs/RunFaucetBtcRequest.md)
 - [RunFaucetResponse](docs/RunFaucetResponse.md)
 - [SearchErrorResult](docs/SearchErrorResult.md)
 - [SearchErrorResultResult](docs/SearchErrorResultResult.md)
 - [SearchResult](docs/SearchResult.md)
 - [SearchSuccessResult](docs/SearchSuccessResult.md)
 - [ServerStatusResponse](docs/ServerStatusResponse.md)
 - [SigningPayload](docs/SigningPayload.md)
 - [SmartContract](docs/SmartContract.md)
 - [SmartContractFound](docs/SmartContractFound.md)
 - [SmartContractNotFound](docs/SmartContractNotFound.md)
 - [SmartContractStatus](docs/SmartContractStatus.md)
 - [SmartContractTransaction](docs/SmartContractTransaction.md)
 - [SmartContractTransactionMetadataSmartContract](docs/SmartContractTransactionMetadataSmartContract.md)
 - [SmartContractsStatusResponse](docs/SmartContractsStatusResponse.md)
 - [StxBalance](docs/StxBalance.md)
 - [TargetBlockTime](docs/TargetBlockTime.md)
 - [TenureChangeTransaction](docs/TenureChangeTransaction.md)
 - [TenureChangeTransactionMetadataTenureChangePayload](docs/TenureChangeTransactionMetadataTenureChangePayload.md)
 - [TokenTransferTransaction](docs/TokenTransferTransaction.md)
 - [TokenTransferTransactionMetadataTokenTransfer](docs/TokenTransferTransactionMetadataTokenTransfer.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionEvent](docs/TransactionEvent.md)
 - [TransactionEventAsset](docs/TransactionEventAsset.md)
 - [TransactionEventFungibleAsset](docs/TransactionEventFungibleAsset.md)
 - [TransactionEventFungibleAssetAllOfAsset](docs/TransactionEventFungibleAssetAllOfAsset.md)
 - [TransactionEventNonFungibleAsset](docs/TransactionEventNonFungibleAsset.md)
 - [TransactionEventNonFungibleAssetAllOfAsset](docs/TransactionEventNonFungibleAssetAllOfAsset.md)
 - [TransactionEventSmartContractLog](docs/TransactionEventSmartContractLog.md)
 - [TransactionEventSmartContractLogAllOfContractLog](docs/TransactionEventSmartContractLogAllOfContractLog.md)
 - [TransactionEventStxAsset](docs/TransactionEventStxAsset.md)
 - [TransactionEventStxLock](docs/TransactionEventStxLock.md)
 - [TransactionEventStxLockAllOfStxLockEvent](docs/TransactionEventStxLockAllOfStxLockEvent.md)
 - [TransactionEventsResponse](docs/TransactionEventsResponse.md)
 - [TransactionFeeEstimateRequest](docs/TransactionFeeEstimateRequest.md)
 - [TransactionFeeEstimateResponse](docs/TransactionFeeEstimateResponse.md)
 - [TransactionFeeEstimateResponseEstimatedCost](docs/TransactionFeeEstimateResponseEstimatedCost.md)
 - [TransactionFeeEstimateResponseEstimationsInner](docs/TransactionFeeEstimateResponseEstimationsInner.md)
 - [TransactionFound](docs/TransactionFound.md)
 - [TransactionFoundResult](docs/TransactionFoundResult.md)
 - [TransactionIdentifier](docs/TransactionIdentifier.md)
 - [TransactionList](docs/TransactionList.md)
 - [TransactionNotFound](docs/TransactionNotFound.md)
 - [TransactionNotFoundResult](docs/TransactionNotFoundResult.md)
 - [TransactionResults](docs/TransactionResults.md)
 - [TxSearchResult](docs/TxSearchResult.md)
 - [TxSearchResultResult](docs/TxSearchResultResult.md)
 - [TxSearchResultResultTxData](docs/TxSearchResultResultTxData.md)
 - [UnanchoredTransactionListResponse](docs/UnanchoredTransactionListResponse.md)

