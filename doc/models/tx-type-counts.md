# Tx Type Counts

Number of tranasction in the mempool, broken down by transaction type.

## Structure

`TxTypeCounts`

## Fields

| Name               | Type      | Tags     | Description |
| ------------------ | --------- | -------- | ----------- |
| `TokenTransfer`    | `float64` | Required | -           |
| `SmartContract`    | `float64` | Required | -           |
| `ContractCall`     | `float64` | Required | -           |
| `PoisonMicroblock` | `float64` | Required | -           |

## Example (as JSON)

```json
{
  "token_transfer": 192.32,
  "smart_contract": 134.02,
  "contract_call": 70.2,
  "poison_microblock": 216.72
}
```
