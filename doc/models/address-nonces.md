# Address Nonces

The latest nonce values used by an account by inspecting the mempool, microblock transactions, and anchored transactions

## Structure

`AddressNonces`

## Fields

| Name                    | Type    | Tags     | Description                                                                                                                                                                                                                                                                                                                                |
| ----------------------- | ------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `LastMempoolTxNonce`    | `*int`  | Required | The latest nonce found within mempool transactions sent by this address. Will be null if there are no current mempool transactions for this address.                                                                                                                                                                                       |
| `LastExecutedTxNonce`   | `*int`  | Required | The latest nonce found within transactions sent by this address, including unanchored microblock transactions. Will be null if there are no current transactions for this address.                                                                                                                                                         |
| `PossibleNextNonce`     | `int`   | Required | The likely nonce required for creating the next transaction, based on the last nonces seen by the API. This can be incorrect if the API's mempool or transactions aren't fully synchronized, even by a small amount, or if a previous transaction is still propagating through the Stacks blockchain network when this endpoint is called. |
| `DetectedMissingNonces` | `[]int` | Required | Nonces that appear to be missing and likely causing a mempool transaction to be stuck.                                                                                                                                                                                                                                                     |
| `DetectedMempoolNonces` | `[]int` | Optional | Nonces currently in mempool for this address.                                                                                                                                                                                                                                                                                              |

## Example (as JSON)

```json
{
  "last_mempool_tx_nonce": 86,
  "last_executed_tx_nonce": 34,
  "possible_next_nonce": 252,
  "detected_missing_nonces": [199, 200],
  "detected_mempool_nonces": [67]
}
```
