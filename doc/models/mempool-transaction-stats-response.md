# Mempool Transaction Stats Response

GET request that returns stats on mempool transactions

## Structure

`MempoolTransactionStatsResponse`

## Fields

| Name                  | Type                                                                       | Tags     | Description                                                                                                                                                                                                                                                                                                            |
| --------------------- | -------------------------------------------------------------------------- | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `TxTypeCounts`        | [`models.TxTypeCounts`](../../doc/models/tx-type-counts.md)                | Required | Number of tranasction in the mempool, broken down by transaction type.                                                                                                                                                                                                                                                 |
| `TxSimpleFeeAverages` | [`models.TxSimpleFeeAverages`](../../doc/models/tx-simple-fee-averages.md) | Required | The simple mean (average) transaction fee, broken down by transaction type. Note that this does not factor in actual execution costs. The average fee is not a reliable metric for calculating a fee for a new transaction.                                                                                            |
| `TxAges`              | [`models.TxAges`](../../doc/models/tx-ages.md)                             | Required | The average time (in blocks) that transactions have lived in the mempool. The start block height is simply the current chain-tip of when the attached Stacks node receives the transaction. This timing can be different across Stacks nodes / API instances due to propagation timing differences in the p2p network. |
| `TxByteSizes`         | [`models.TxByteSizes`](../../doc/models/tx-byte-sizes.md)                  | Required | The average byte size of transactions in the mempool, broken down by transaction type.                                                                                                                                                                                                                                 |

## Example (as JSON)

```json
{
  "tx_type_counts": {
    "token_transfer": 31.86,
    "smart_contract": 229.56,
    "contract_call": 165.74,
    "poison_microblock": 56.26
  },
  "tx_simple_fee_averages": {
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
  },
  "tx_ages": {
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
  },
  "tx_byte_sizes": {
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
}
```
