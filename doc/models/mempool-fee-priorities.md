# Mempool Fee Priorities

GET request that returns fee priorities from mempool transactions

## Structure

`MempoolFeePriorities`

## Fields

| Name            | Type                                                             | Tags     | Description |
| --------------- | ---------------------------------------------------------------- | -------- | ----------- |
| `All`           | [`models.All`](../../doc/models/all.md)                          | Required | -           |
| `TokenTransfer` | [`*models.TokenTransfer1`](../../doc/models/token-transfer-1.md) | Optional | -           |
| `SmartContract` | [`*models.SmartContract1`](../../doc/models/smart-contract-1.md) | Optional | -           |
| `ContractCall`  | [`*models.ContractCall1`](../../doc/models/contract-call-1.md)   | Optional | -           |

## Example (as JSON)

```json
{
  "all": {
    "no_priority": 180,
    "low_priority": 220,
    "medium_priority": 90,
    "high_priority": 2
  },
  "token_transfer": {
    "no_priority": 20,
    "low_priority": 60,
    "medium_priority": 186,
    "high_priority": 98
  },
  "smart_contract": {
    "no_priority": 78,
    "low_priority": 118,
    "medium_priority": 244,
    "high_priority": 156
  },
  "contract_call": {
    "no_priority": 96,
    "low_priority": 136,
    "medium_priority": 6,
    "high_priority": 174
  }
}
```
