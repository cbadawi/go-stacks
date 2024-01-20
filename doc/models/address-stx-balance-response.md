# Address Stx Balance Response

GET request that returns address balances

## Structure

`AddressStxBalanceResponse`

## Fields

| Name                        | Type                                                                                      | Tags     | Description                                                                               |
| --------------------------- | ----------------------------------------------------------------------------------------- | -------- | ----------------------------------------------------------------------------------------- |
| `Balance`                   | `string`                                                                                  | Required | -                                                                                         |
| `TotalSent`                 | `string`                                                                                  | Required | -                                                                                         |
| `TotalReceived`             | `string`                                                                                  | Required | -                                                                                         |
| `TotalFeesSent`             | `string`                                                                                  | Required | -                                                                                         |
| `TotalMinerRewardsReceived` | `string`                                                                                  | Required | -                                                                                         |
| `LockTxId`                  | `string`                                                                                  | Required | The transaction where the lock event occurred. Empty if no tokens are locked.             |
| `Locked`                    | `string`                                                                                  | Required | The amount of locked STX, as string quoted micro-STX. Zero if no tokens are locked.       |
| `LockHeight`                | `int`                                                                                     | Required | The STX chain block height of when the lock event occurred. Zero if no tokens are locked. |
| `BurnchainLockHeight`       | `int`                                                                                     | Required | The burnchain block height of when the lock event occurred. Zero if no tokens are locked. |
| `BurnchainUnlockHeight`     | `int`                                                                                     | Required | The burnchain block height of when the tokens unlock. Zero if no tokens are locked.       |
| `TokenOfferingLocked`       | [`*models.AddressTokenOfferingLocked`](../../doc/models/address-token-offering-locked.md) | Optional | Token Offering Locked                                                                     |

## Example (as JSON)

```json
{
  "balance": "balance4",
  "total_sent": "total_sent0",
  "total_received": "total_received4",
  "total_fees_sent": "total_fees_sent0",
  "total_miner_rewards_received": "total_miner_rewards_received0",
  "lock_tx_id": "lock_tx_id8",
  "locked": "locked8",
  "lock_height": 88,
  "burnchain_lock_height": 68,
  "burnchain_unlock_height": 84,
  "token_offering_locked": {
    "total_locked": "total_locked4",
    "total_unlocked": "total_unlocked6",
    "unlock_schedule": [
      {
        "amount": "amount4",
        "block_height": 184.68
      }
    ]
  }
}
```
