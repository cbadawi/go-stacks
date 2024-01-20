# Tx Simple Fee Averages

The simple mean (average) transaction fee, broken down by transaction type. Note that this does not factor in actual execution costs. The average fee is not a reliable metric for calculating a fee for a new transaction.

## Structure

`TxSimpleFeeAverages`

## Fields

| Name               | Type                                                                  | Tags     | Description |
| ------------------ | --------------------------------------------------------------------- | -------- | ----------- |
| `TokenTransfer`    | [`models.TokenTransfer2`](../../doc/models/token-transfer-2.md)       | Required | -           |
| `SmartContract`    | [`models.SmartContract2`](../../doc/models/smart-contract-2.md)       | Required | -           |
| `ContractCall`     | [`models.ContractCall2`](../../doc/models/contract-call-2.md)         | Required | -           |
| `PoisonMicroblock` | [`models.PoisonMicroblock1`](../../doc/models/poison-microblock-1.md) | Required | -           |

## Example (as JSON)

```json
{
  "token_transfer": {
    "p25": 97.48,
    "p50": 144.58,
    "p75": 160.42,
    "p95": 17.1
  },
  "smart_contract": {
    "p25": 39.18,
    "p50": 86.28,
    "p75": 102.12,
    "p95": 214.8
  },
  "contract_call": {
    "p25": 231.36,
    "p50": 22.46,
    "p75": 38.3,
    "p95": 150.98
  },
  "poison_microblock": {
    "p25": 72.48,
    "p50": 119.58,
    "p75": 135.42,
    "p95": 248.1
  }
}
```
