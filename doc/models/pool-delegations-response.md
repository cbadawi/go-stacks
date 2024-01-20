# Pool Delegations Response

GET request that returns stacking pool member details for a given pool (delegator) principal

## Structure

`PoolDelegationsResponse`

## Fields

| Name      | Type                                                             | Tags     | Description                                                          |
| --------- | ---------------------------------------------------------------- | -------- | -------------------------------------------------------------------- |
| `Limit`   | `int`                                                            | Required | The number of Stackers to return<br>**Constraints**: `<= 200`        |
| `Offset`  | `int`                                                            | Required | The number to Stackers to skip (starting at `0`)<br>**Default**: `0` |
| `Total`   | `int`                                                            | Required | The total number of Stackers                                         |
| `Results` | [`[]models.PoolDelegation`](../../doc/models/pool-delegation.md) | Required | -                                                                    |

## Example (as JSON)

```json
{
  "limit": 250,
  "offset": 0,
  "total": 88,
  "results": [
    {
      "stacker": "stacker2",
      "pox_addr": "pox_addr2",
      "amount_ustx": "amount_ustx0",
      "burn_block_unlock_height": 198,
      "block_height": 52,
      "tx_id": "tx_id0"
    }
  ]
}
```
