
# Result

## Structure

`Result`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TxId` | `*string` | Optional | Transaction ID |
| `Nonce` | `*int` | Optional | Used for ordering the transactions originating from and paying from an account. The nonce ensures that a transaction is processed at most once. The nonce counts the number of times an account's owner(s) have authorized a transaction. The first transaction from an account will have a nonce value equal to 0, the second will have a nonce value equal to 1, and so on. |
| `FeeRate` | `*string` | Optional | Transaction fee as Integer string (64-bit unsigned integer). |
| `SenderAddress` | `*string` | Optional | Address of the transaction initiator |
| `SponsorNonce` | `*int` | Optional | - |
| `Sponsored` | `*bool` | Optional | Denotes whether the originating account is the same as the paying account |
| `SponsorAddress` | `*string` | Optional | - |
| `PostConditionMode` | [`*models.PostConditionModeEnum`](../../doc/models/post-condition-mode-enum.md) | Optional | - |
| `PostConditions` | [`[]models.PostCondition1`](../../doc/models/post-condition-1.md) | Optional | - |
| `AnchorMode` | [`*models.TransactionAnchorModeTypeEnum`](../../doc/models/transaction-anchor-mode-type-enum.md) | Optional | `on_chain_only`: the transaction MUST be included in an anchored block, `off_chain_only`: the transaction MUST be included in a microblock, `any`: the leader can choose where to include the transaction. |
| `TxStatus` | [`*models.MempoolTransactionStatusEnum`](../../doc/models/mempool-transaction-status-enum.md) | Optional | Status of the transaction |
| `ReceiptTime` | `*float64` | Optional | A unix timestamp (in seconds) indicating when the transaction broadcast was received by the node. |
| `ReceiptTimeIso` | `*string` | Optional | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when the transaction broadcast was received by the node. |
| `TxType` | [`*models.TxTypeEnum`](../../doc/models/tx-type-enum.md) | Optional | - |
| `TokenTransfer` | [`*models.TokenTransfer`](../../doc/models/token-transfer.md) | Optional | - |
| `SmartContract` | [`*models.SmartContract`](../../doc/models/smart-contract.md) | Optional | - |
| `ContractCall` | [`*models.ContractCall`](../../doc/models/contract-call.md) | Optional | - |
| `PoisonMicroblock` | [`*models.PoisonMicroblock`](../../doc/models/poison-microblock.md) | Optional | - |
| `CoinbasePayload` | [`*models.CoinbasePayload`](../../doc/models/coinbase-payload.md) | Optional | - |
| `TenureChangePayload` | [`*models.TenureChangePayload`](../../doc/models/tenure-change-payload.md) | Optional | - |
| `BlockHash` | `*string` | Optional | Hash of the blocked this transactions was associated with |
| `BlockHeight` | `*int` | Optional | Height of the block this transactions was associated with |
| `BurnBlockTime` | `*int` | Optional | Unix timestamp (in seconds) indicating when this block was mined |
| `BurnBlockTimeIso` | `*string` | Optional | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when this block was mined. |
| `ParentBurnBlockTime` | `*int` | Optional | Unix timestamp (in seconds) indicating when this parent block was mined |
| `ParentBurnBlockTimeIso` | `*string` | Optional | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when this parent block was mined. |
| `Canonical` | `*bool` | Optional | Set to `true` if block corresponds to the canonical chain tip |
| `TxIndex` | `*int` | Optional | Index of the transaction, indicating the order. Starts at `0` and increases with each transaction |
| `TxResult` | [`*models.TxResult`](../../doc/models/tx-result.md) | Optional | Result of the transaction. For contract calls, this will show the value returned by the call. For other transaction types, this will return a boolean indicating the success of the transaction. |
| `EventCount` | `*int` | Optional | Number of transaction events |
| `ParentBlockHash` | `*string` | Optional | Hash of the previous block. |
| `IsUnanchored` | `*bool` | Optional | True if the transaction is included in a microblock that has not been confirmed by an anchor block. |
| `MicroblockHash` | `*string` | Optional | The microblock hash that this transaction was streamed in. If the transaction was batched in an anchor block (not included within a microblock) then this value will be an empty string. |
| `MicroblockSequence` | `*int` | Optional | The microblock sequence number that this transaction was streamed in. If the transaction was batched in an anchor block (not included within a microblock) then this value will be 2147483647 (0x7fffffff, the max int32 value), this value preserves logical transaction ordering on (block_height, microblock_sequence, tx_index). |
| `MicroblockCanonical` | `*bool` | Optional | Set to `true` if microblock is anchored in the canonical chain tip, `false` if the transaction was orphaned in a micro-fork. |
| `ExecutionCostReadCount` | `*int` | Optional | Execution cost read count. |
| `ExecutionCostReadLength` | `*int` | Optional | Execution cost read length. |
| `ExecutionCostRuntime` | `*int` | Optional | Execution cost runtime. |
| `ExecutionCostWriteCount` | `*int` | Optional | Execution cost write count. |
| `ExecutionCostWriteLength` | `*int` | Optional | Execution cost write length. |
| `Events` | [`[]models.TransactionEvent1`](../../doc/models/transaction-event-1.md) | Optional | List of transaction events |

## Example (as JSON)

```json
{
  "tx_id": "tx_id4",
  "nonce": 122,
  "fee_rate": "fee_rate2",
  "sender_address": "sender_address2",
  "sponsor_nonce": 90
}
```

