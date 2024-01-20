# Rosetta Oldest Block Identifier

The block_identifier uniquely identifies a block in a particular network.

## Structure

`RosettaOldestBlockIdentifier`

## Fields

| Name    | Type     | Tags     | Description                             |
| ------- | -------- | -------- | --------------------------------------- |
| `Index` | `int`    | Required | This is also known as the block height. |
| `Hash`  | `string` | Required | Block hash                              |

## Example (as JSON)

```json
{
  "index": 108,
  "hash": "hash8"
}
```
