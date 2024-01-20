# Sub Network Identifier

In blockchains with sharded state, the SubNetworkIdentifier is required to query some object on a specific shard. This identifier is optional for all non-sharded blockchains.

## Structure

`SubNetworkIdentifier`

## Fields

| Name       | Type                                                  | Tags     | Description                          |
| ---------- | ----------------------------------------------------- | -------- | ------------------------------------ |
| `Network`  | `string`                                              | Required | Network name                         |
| `Metadata` | [`*models.Metadata1`](../../doc/models/metadata-1.md) | Optional | Meta data from subnetwork identifier |

## Example (as JSON)

```json
{
  "network": "network2",
  "metadata": {
    "producer": "producer4"
  }
}
```
