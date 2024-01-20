# Inbound Stx Transfer

A inbound STX transfer with a memo

## Structure

`InboundStxTransfer`

## Fields

| Name           | Type                                                                | Tags     | Description                                                                                 |
| -------------- | ------------------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------- |
| `Sender`       | `string`                                                            | Required | Principal that sent this transfer                                                           |
| `Amount`       | `string`                                                            | Required | Transfer amount in micro-STX as integer string                                              |
| `Memo`         | `string`                                                            | Required | Hex encoded memo bytes associated with the transfer                                         |
| `BlockHeight`  | `float64`                                                           | Required | Block height at which this transfer occurred                                                |
| `TxId`         | `string`                                                            | Required | The transaction ID in which this transfer occurred                                          |
| `TransferType` | [`models.TransferTypeEnum`](../../doc/models/transfer-type-enum.md) | Required | Indicates if the transfer is from a stx-transfer transaction or a contract-call transaction |
| `TxIndex`      | `float64`                                                           | Required | Index of the transaction within a block                                                     |

## Example (as JSON)

```json
{
  "sender": "sender6",
  "amount": "amount4",
  "memo": "memo6",
  "block_height": 7.88,
  "tx_id": "tx_id6",
  "transfer_type": "bulk-send",
  "tx_index": 92.7
}
```
