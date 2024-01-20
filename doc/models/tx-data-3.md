# Tx Data 3

Returns basic search result information about the requested id

## Structure

`TxData3`

## Fields

| Name            | Type      | Tags     | Description                                              |
| --------------- | --------- | -------- | -------------------------------------------------------- |
| `Canonical`     | `bool`    | Required | If the transaction lies within the canonical chain       |
| `BlockHash`     | `string`  | Required | Refers to the hash of the block for searched transaction |
| `BurnBlockTime` | `int`     | Required | -                                                        |
| `BlockHeight`   | `int`     | Required | -                                                        |
| `TxType`        | `string`  | Required | -                                                        |
| `TxId`          | `*string` | Optional | Corresponding tx_id for smart_contract                   |

## Example (as JSON)

```json
{
  "canonical": false,
  "block_hash": "block_hash6",
  "burn_block_time": 4,
  "block_height": 218,
  "tx_type": "tx_type2",
  "tx_id": "tx_id8"
}
```
