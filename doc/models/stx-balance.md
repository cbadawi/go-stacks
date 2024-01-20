# Stx Balance

## Structure

`StxBalance`

## Fields

| Name                        | Type     | Tags     | Description                                                                               |
| --------------------------- | -------- | -------- | ----------------------------------------------------------------------------------------- |
| `Balance`                   | `string` | Required | -                                                                                         |
| `TotalSent`                 | `string` | Required | -                                                                                         |
| `TotalReceived`             | `string` | Required | -                                                                                         |
| `TotalFeesSent`             | `string` | Required | -                                                                                         |
| `TotalMinerRewardsReceived` | `string` | Required | -                                                                                         |
| `LockTxId`                  | `string` | Required | The transaction where the lock event occurred. Empty if no tokens are locked.             |
| `Locked`                    | `string` | Required | The amount of locked STX, as string quoted micro-STX. Zero if no tokens are locked.       |
| `LockHeight`                | `int`    | Required | The STX chain block height of when the lock event occurred. Zero if no tokens are locked. |
| `BurnchainLockHeight`       | `int`    | Required | The burnchain block height of when the lock event occurred. Zero if no tokens are locked. |
| `BurnchainUnlockHeight`     | `int`    | Required | The burnchain block height of when the tokens unlock. Zero if no tokens are locked.       |

## Example (as JSON)

```json
{
  "balance": "balance6",
  "total_sent": "total_sent8",
  "total_received": "total_received6",
  "total_fees_sent": "total_fees_sent8",
  "total_miner_rewards_received": "total_miner_rewards_received2",
  "lock_tx_id": "lock_tx_id0",
  "locked": "locked0",
  "lock_height": 16,
  "burnchain_lock_height": 252,
  "burnchain_unlock_height": 156
}
```
