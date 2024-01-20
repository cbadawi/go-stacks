# Metadata 3

Transactions that are related to other transactions (like a cross-shard transaction) should include the tranaction_identifier of these transactions in the metadata.

## Structure

`Metadata3`

## Fields

| Name       | Type      | Tags     | Description              |
| ---------- | --------- | -------- | ------------------------ |
| `Memo`     | `*string` | Optional | STX token transfer memo. |
| `Size`     | `*int`    | Optional | The Size                 |
| `LockTime` | `*int`    | Optional | The locktime             |

## Example (as JSON)

```json
{
  "memo": "memo0",
  "size": 158,
  "lockTime": 18
}
```
