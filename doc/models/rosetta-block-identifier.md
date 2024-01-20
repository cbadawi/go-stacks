
# Rosetta Block Identifier

The block_identifier uniquely identifies a block in a particular network.

## Structure

`RosettaBlockIdentifier`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Hash` | `string` | Required | This is also known as the block hash. |
| `Index` | `int` | Required | This is also known as the block height. |

## Example (as JSON)

```json
{
  "hash": "hash0",
  "index": 248
}
```

