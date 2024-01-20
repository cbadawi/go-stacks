
# Result 11

This object carries the search result

## Structure

`Result11`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EntityType` | [`models.EntityTypeEnum`](../../doc/models/entity-type-enum.md) | Required | Shows the currenty category of entity it is searched in. |
| `EntityId` | `string` | Required | The id used to search this query. |
| `Metadata` | [`*models.Block1`](../../doc/models/block-1.md) | Optional | A block |
| `BlockData` | [`models.BlockData`](../../doc/models/block-data.md) | Required | Returns basic search result information about the requested id |
| `TxData` | [`models.TxData3`](../../doc/models/tx-data-3.md) | Required | Returns basic search result information about the requested id |

## Example (as JSON)

```json
{
  "entity_type": "standard_address",
  "entity_id": "entity_id6",
  "metadata": {
    "canonical": false,
    "height": 42,
    "hash": "hash2",
    "index_block_hash": "index_block_hash6",
    "parent_block_hash": "parent_block_hash0",
    "burn_block_time": 170.86,
    "burn_block_time_iso": "burn_block_time_iso4",
    "burn_block_hash": "burn_block_hash0",
    "burn_block_height": 208,
    "miner_txid": "miner_txid8",
    "parent_microblock_hash": "parent_microblock_hash8",
    "parent_microblock_sequence": 72,
    "txs": [
      "txs7"
    ],
    "microblocks_accepted": [
      "microblocks_accepted8",
      "microblocks_accepted9"
    ],
    "microblocks_streamed": [
      "microblocks_streamed8"
    ],
    "execution_cost_read_count": 114,
    "execution_cost_read_length": 222,
    "execution_cost_runtime": 78,
    "execution_cost_write_count": 102,
    "execution_cost_write_length": 166,
    "microblock_tx_count": {
      "key0": 246.64,
      "key1": 246.65
    },
    "tx_id": "tx_id0",
    "nonce": 88,
    "fee_rate": "fee_rate8",
    "sender_address": "sender_address8",
    "sponsor_nonce": 132,
    "sponsored": false,
    "sponsor_address": "sponsor_address4",
    "post_condition_mode": "allow",
    "post_conditions": [
      {
        "principal": {
          "type_id": "type_id0",
          "address": "address6",
          "contract_name": "contract_name2"
        },
        "condition_code": "sent_less_than",
        "amount": "amount6",
        "type": "type6",
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
    "anchor_mode": "any",
    "tx_status": "dropped_replace_across_fork",
    "receipt_time": 23.6,
    "receipt_time_iso": "receipt_time_iso2",
    "tx_type": "tx_type4",
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
    },
    "block_hash": "block_hash8",
    "block_height": 156,
    "parent_burn_block_time": 114,
    "parent_burn_block_time_iso": "parent_burn_block_time_iso8",
    "tx_index": 190,
    "tx_result": {
      "hex": "hex2",
      "repr": "repr8"
    },
    "event_count": 122,
    "is_unanchored": false,
    "microblock_hash": "microblock_hash4",
    "microblock_sequence": 210,
    "microblock_canonical": false,
    "events": [
      {
        "event_index": 246,
        "event_type": "event_type0",
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
      },
      {
        "event_index": 246,
        "event_type": "event_type0",
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
      },
      {
        "event_index": 246,
        "event_type": "event_type0",
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
    ]
  },
  "block_data": {
    "canonical": false,
    "hash": "hash8",
    "parent_block_hash": "parent_block_hash6",
    "burn_block_time": 110,
    "height": 254
  },
  "tx_data": {
    "canonical": false,
    "block_hash": "block_hash8",
    "burn_block_time": 226,
    "block_height": 192,
    "tx_type": "tx_type4",
    "tx_id": "tx_id0"
  }
}
```

