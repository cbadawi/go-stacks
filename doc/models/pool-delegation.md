# Pool Delegation

## Structure

`PoolDelegation`

## Fields

| Name                    | Type      | Tags     | Description                                                                    |
| ----------------------- | --------- | -------- | ------------------------------------------------------------------------------ |
| `Stacker`               | `string`  | Required | The principal of the pool member that issued the delegation                    |
| `PoxAddr`               | `*string` | Optional | The pox-addr value specified by the stacker in the delegation operation        |
| `AmountUstx`            | `string`  | Required | The amount of uSTX delegated by the stacker                                    |
| `BurnBlockUnlockHeight` | `*int`    | Optional | The optional burnchain block unlock height that the stacker may have specified |
| `BlockHeight`           | `int`     | Required | The block height at which the stacker delegation transaction was mined at      |
| `TxId`                  | `string`  | Required | The tx_id of the stacker delegation operation                                  |

## Example (as JSON)

```json
{
  "stacker": "stacker6",
  "pox_addr": "pox_addr6",
  "amount_ustx": "amount_ustx6",
  "burn_block_unlock_height": 48,
  "block_height": 202,
  "tx_id": "tx_id4"
}
```
