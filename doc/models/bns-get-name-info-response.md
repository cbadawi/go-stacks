# Bns Get Name Info Response

Get name details

## Structure

`BnsGetNameInfoResponse`

## Fields

| Name           | Type      | Tags     | Description                            |
| -------------- | --------- | -------- | -------------------------------------- |
| `Address`      | `string`  | Required | -                                      |
| `Blockchain`   | `string`  | Required | **Constraints**: _Pattern_: `^stacks$` |
| `ExpireBlock`  | `*int`    | Optional | **Constraints**: `>= 0`                |
| `GracePeriod`  | `*int`    | Optional | **Constraints**: `>= 0`                |
| `LastTxid`     | `string`  | Required | -                                      |
| `Resolver`     | `*string` | Optional | -                                      |
| `Status`       | `string`  | Required | -                                      |
| `Zonefile`     | `string`  | Required | -                                      |
| `ZonefileHash` | `string`  | Required | -                                      |

## Example (as JSON)

```json
{
  "address": "address2",
  "blockchain": "blockchain8",
  "expire_block": 240,
  "grace_period": 192,
  "last_txid": "last_txid6",
  "resolver": "resolver4",
  "status": "status2",
  "zonefile": "zonefile4",
  "zonefile_hash": "zonefile_hash8"
}
```
