# Account Data Response

GET request for account data

## Structure

`AccountDataResponse`

## Fields

| Name           | Type     | Tags     | Description |
| -------------- | -------- | -------- | ----------- |
| `Balance`      | `string` | Required | -           |
| `Locked`       | `string` | Required | -           |
| `UnlockHeight` | `int`    | Required | -           |
| `Nonce`        | `int`    | Required | -           |
| `BalanceProof` | `string` | Required | -           |
| `NonceProof`   | `string` | Required | -           |

## Example (as JSON)

```json
{
  "balance": "balance8",
  "locked": "locked2",
  "unlock_height": 116,
  "nonce": 170,
  "balance_proof": "balance_proof4",
  "nonce_proof": "nonce_proof4"
}
```
