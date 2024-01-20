# Rosetta Partial Block Identifier

When fetching data by BlockIdentifier, it may be possible to only specify the index or hash. If neither property is specified, it is assumed that the client is making a request at the current block.

## Structure

`RosettaPartialBlockIdentifier`

## Fields

| Name    | Type      | Tags     | Description                             |
| ------- | --------- | -------- | --------------------------------------- |
| `Hash`  | `*string` | Optional | This is also known as the block hash.   |
| `Index` | `*int`    | Optional | This is also known as the block height. |

## Example (as JSON)

```json
{
  "hash": "hash4",
  "index": 100
}
```
