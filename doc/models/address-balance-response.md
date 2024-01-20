# Address Balance Response

GET request that returns address balances

## Structure

`AddressBalanceResponse`

## Fields

| Name                  | Type                                                                                      | Tags     | Description           |
| --------------------- | ----------------------------------------------------------------------------------------- | -------- | --------------------- |
| `Stx`                 | [`models.StxBalance`](../../doc/models/stx-balance.md)                                    | Required | -                     |
| `FungibleTokens`      | [`map[string]models.FtBalance`](../../doc/models/ft-balance.md)                           | Required | -                     |
| `NonFungibleTokens`   | [`map[string]models.NftBalance`](../../doc/models/nft-balance.md)                         | Required | -                     |
| `TokenOfferingLocked` | [`*models.AddressTokenOfferingLocked`](../../doc/models/address-token-offering-locked.md) | Optional | Token Offering Locked |

## Example (as JSON)

```json
{
  "stx": {
    "balance": "balance4",
    "total_sent": "total_sent0",
    "total_received": "total_received4",
    "total_fees_sent": "total_fees_sent0",
    "total_miner_rewards_received": "total_miner_rewards_received0",
    "lock_tx_id": "lock_tx_id8",
    "locked": "locked8",
    "lock_height": 238,
    "burnchain_lock_height": 218,
    "burnchain_unlock_height": 190
  },
  "fungible_tokens": {
    "key0": {
      "balance": "balance6",
      "total_sent": "total_sent8",
      "total_received": "total_received6"
    }
  },
  "non_fungible_tokens": {
    "key0": {
      "count": "count6",
      "total_sent": "total_sent2",
      "total_received": "total_received2"
    },
    "key1": {
      "count": "count6",
      "total_sent": "total_sent2",
      "total_received": "total_received2"
    }
  },
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
