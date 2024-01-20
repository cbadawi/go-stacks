
# Block 1

A block

## Structure

`Block1`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Canonical` | `bool` | Required | Set to `true` if block corresponds to the canonical chain tip |
| `Height` | `int` | Required | Height of the block |
| `Hash` | `string` | Required | Hash representing the block |
| `IndexBlockHash` | `string` | Required | The only hash that can uniquely identify an anchored block or an unconfirmed state trie |
| `ParentBlockHash` | `string` | Required | Hash of the parent block |
| `BurnBlockTime` | `float64` | Required | Unix timestamp (in seconds) indicating when this block was mined. |
| `BurnBlockTimeIso` | `string` | Required | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined. |
| `BurnBlockHash` | `string` | Required | Hash of the anchor chain block |
| `BurnBlockHeight` | `int` | Required | Height of the anchor chain block |
| `MinerTxid` | `string` | Required | Anchor chain transaction ID |
| `ParentMicroblockHash` | `string` | Required | The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1. |
| `ParentMicroblockSequence` | `int` | Required | The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1. |
| `Txs` | `[]string` | Required | List of transactions included in the block |
| `MicroblocksAccepted` | `[]string` | Required | List of microblocks that were accepted in this anchor block. Not every anchored block will have a accepted all (or any) of the previously streamed microblocks. Microblocks that were orphaned are not included in this list. |
| `MicroblocksStreamed` | `[]string` | Required | List of microblocks that were streamed/produced by this anchor block's miner. This list only includes microblocks that were accepted in the following anchor block. Microblocks that were orphaned are not included in this list. |
| `ExecutionCostReadCount` | `int` | Required | Execution cost read count. |
| `ExecutionCostReadLength` | `int` | Required | Execution cost read length. |
| `ExecutionCostRuntime` | `int` | Required | Execution cost runtime. |
| `ExecutionCostWriteCount` | `int` | Required | Execution cost write count. |
| `ExecutionCostWriteLength` | `int` | Required | Execution cost write length. |
| `MicroblockTxCount` | `map[string]float64` | Required | List of txs counts in each accepted microblock |
| `TxId` | `string` | Required | Transaction ID |
| `Nonce` | `int` | Required | Used for ordering the transactions originating from and paying from an account. The nonce ensures that a transaction is processed at most once. The nonce counts the number of times an account's owner(s) have authorized a transaction. The first transaction from an account will have a nonce value equal to 0, the second will have a nonce value equal to 1, and so on. |
| `FeeRate` | `string` | Required | Transaction fee as Integer string (64-bit unsigned integer). |
| `SenderAddress` | `string` | Required | Address of the transaction initiator |
| `SponsorNonce` | `*int` | Optional | - |
| `Sponsored` | `bool` | Required | Denotes whether the originating account is the same as the paying account |
| `SponsorAddress` | `*string` | Optional | - |
| `PostConditionMode` | [`models.PostConditionModeEnum`](../../doc/models/post-condition-mode-enum.md) | Required | - |
| `PostConditions` | [`[]models.PostCondition1`](../../doc/models/post-condition-1.md) | Required | - |
| `AnchorMode` | [`models.TransactionAnchorModeTypeEnum`](../../doc/models/transaction-anchor-mode-type-enum.md) | Required | `on_chain_only`: the transaction MUST be included in an anchored block, `off_chain_only`: the transaction MUST be included in a microblock, `any`: the leader can choose where to include the transaction. |
| `TxStatus` | [`models.MempoolTransactionStatusEnum`](../../doc/models/mempool-transaction-status-enum.md) | Required | Status of the transaction |
| `ReceiptTime` | `float64` | Required | A unix timestamp (in seconds) indicating when the transaction broadcast was received by the node. |
| `ReceiptTimeIso` | `string` | Required | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when the transaction broadcast was received by the node. |
| `TxType` | `string` | Required, Constant | **Default**: `"token_transfer"` |
| `TokenTransfer` | [`models.TokenTransfer`](../../doc/models/token-transfer.md) | Required | - |
| `SmartContract` | [`models.SmartContract`](../../doc/models/smart-contract.md) | Required | - |
| `ContractCall` | [`models.ContractCall`](../../doc/models/contract-call.md) | Required | - |
| `PoisonMicroblock` | [`models.PoisonMicroblock`](../../doc/models/poison-microblock.md) | Required | - |
| `CoinbasePayload` | [`models.CoinbasePayload`](../../doc/models/coinbase-payload.md) | Required | - |
| `TenureChangePayload` | [`*models.TenureChangePayload`](../../doc/models/tenure-change-payload.md) | Optional | - |
| `BlockHash` | `string` | Required | Hash of the blocked this transactions was associated with |
| `BlockHeight` | `int` | Required | Height of the block this transactions was associated with |
| `ParentBurnBlockTime` | `int` | Required | Unix timestamp (in seconds) indicating when this parent block was mined |
| `ParentBurnBlockTimeIso` | `string` | Required | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when this parent block was mined. |
| `TxIndex` | `int` | Required | Index of the transaction, indicating the order. Starts at `0` and increases with each transaction |
| `TxResult` | [`models.TxResult`](../../doc/models/tx-result.md) | Required | Result of the transaction. For contract calls, this will show the value returned by the call. For other transaction types, this will return a boolean indicating the success of the transaction. |
| `EventCount` | `int` | Required | Number of transaction events |
| `IsUnanchored` | `bool` | Required | True if the transaction is included in a microblock that has not been confirmed by an anchor block. |
| `MicroblockHash` | `string` | Required | The microblock hash that this transaction was streamed in. If the transaction was batched in an anchor block (not included within a microblock) then this value will be an empty string. |
| `MicroblockSequence` | `int` | Required | The microblock sequence number that this transaction was streamed in. If the transaction was batched in an anchor block (not included within a microblock) then this value will be 2147483647 (0x7fffffff, the max int32 value), this value preserves logical transaction ordering on (block_height, microblock_sequence, tx_index). |
| `MicroblockCanonical` | `bool` | Required | Set to `true` if microblock is anchored in the canonical chain tip, `false` if the transaction was orphaned in a micro-fork. |
| `Events` | [`[]models.TransactionEvent1`](../../doc/models/transaction-event-1.md) | Required | List of transaction events |

## Example (as JSON)

```json
{
  "canonical": false,
  "height": 2,
  "hash": "hash8",
  "index_block_hash": "index_block_hash2",
  "parent_block_hash": "parent_block_hash6",
  "burn_block_time": 70.18,
  "burn_block_time_iso": "burn_block_time_iso0",
  "burn_block_hash": "burn_block_hash6",
  "burn_block_height": 168,
  "miner_txid": "miner_txid4",
  "parent_microblock_hash": "parent_microblock_hash4",
  "parent_microblock_sequence": 32,
  "txs": [
    "txs9"
  ],
  "microblocks_accepted": [
    "microblocks_accepted4",
    "microblocks_accepted5"
  ],
  "microblocks_streamed": [
    "microblocks_streamed4"
  ],
  "execution_cost_read_count": 74,
  "execution_cost_read_length": 182,
  "execution_cost_runtime": 218,
  "execution_cost_write_count": 194,
  "execution_cost_write_length": 130,
  "microblock_tx_count": {
    "key0": 5.6,
    "key1": 5.61
  },
  "tx_id": "tx_id6",
  "nonce": 48,
  "fee_rate": "fee_rate4",
  "sender_address": "sender_address4",
  "sponsored": false,
  "post_condition_mode": "allow",
  "post_conditions": [
    {
      "principal": {
        "type_id": "principal_origin",
        "address": "address6",
        "contract_name": "contract_name2"
      },
      "condition_code": "sent_less_than",
      "amount": "amount6",
      "type": "stx",
      "asset": {
        "asset_name": "asset_name8",
        "contract_address": "contract_address4",
        "contract_name": "contract_name2"
      },
      "asset_value": {
        "hex": "hex2",
        "repr": "repr2"
      }
    }
  ],
  "anchor_mode": "off_chain_only",
  "tx_status": "dropped_replace_across_fork",
  "receipt_time": 38.56,
  "receipt_time_iso": "receipt_time_iso8",
  "tx_type": "token_transfer",
  "token_transfer": {
    "recipient_address": "recipient_address6",
    "amount": "amount2",
    "memo": "memo4"
  },
  "smart_contract": {
    "clarity_version": 187.5,
    "contract_id": "contract_id8",
    "source_code": "source_code4"
  },
  "contract_call": {
    "contract_id": "contract_id6",
    "function_name": "function_name6",
    "function_signature": "function_signature2",
    "function_args": [
      {
        "hex": "hex8",
        "repr": "repr2",
        "name": "name8",
        "type": "type2"
      },
      {
        "hex": "hex8",
        "repr": "repr2",
        "name": "name8",
        "type": "type2"
      }
    ]
  },
  "poison_microblock": {
    "microblock_header_1": "microblock_header_18",
    "microblock_header_2": "microblock_header_26"
  },
  "coinbase_payload": {
    "data": "data0",
    "alt_recipient": "alt_recipient8",
    "vrf_proof": "vrf_proof0"
  },
  "block_hash": "block_hash4",
  "block_height": 116,
  "parent_burn_block_time": 74,
  "parent_burn_block_time_iso": "parent_burn_block_time_iso6",
  "tx_index": 150,
  "tx_result": {
    "hex": "hex2",
    "repr": "repr8"
  },
  "event_count": 82,
  "is_unanchored": false,
  "microblock_hash": "microblock_hash0",
  "microblock_sequence": 170,
  "microblock_canonical": false,
  "events": [
    {
      "event_index": 246,
      "event_type": "smart_contract_log",
      "tx_id": "tx_id4",
      "contract_log": {
        "contract_id": "contract_id0",
        "topic": "topic4",
        "value": {
          "hex": "hex6",
          "repr": "repr2"
        }
      },
      "stx_lock_event": {
        "locked_amount": "locked_amount0",
        "unlock_height": 234,
        "locked_address": "locked_address2"
      },
      "asset": {
        "asset_event_type": "burn",
        "asset_id": "asset_id6",
        "sender": "sender2",
        "recipient": "recipient2",
        "amount": "amount2",
        "value": {
          "hex": "hex6",
          "repr": "repr2"
        },
        "memo": "memo4"
      }
    }
  ],
  "sponsor_nonce": 164,
  "sponsor_address": "sponsor_address0",
  "tenure_change_payload": {
    "tenure_consensus_hash": "tenure_consensus_hash8",
    "prev_tenure_consensus_hash": "prev_tenure_consensus_hash2",
    "burn_view_consensus_hash": "burn_view_consensus_hash8",
    "previous_tenure_end": "previous_tenure_end4",
    "previous_tenure_blocks": 199.94,
    "cause": "block_found",
    "pubkey_hash": "pubkey_hash0",
    "signature": "signature8",
    "signers": "signers8"
  }
}
```

