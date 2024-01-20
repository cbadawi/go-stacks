# Network Identifier

The network_identifier specifies which network a particular object is associated with.

## Structure

`NetworkIdentifier`

## Fields

| Name                   | Type                                                                         | Tags     | Description                                                                                                                                                                             |
| ---------------------- | ---------------------------------------------------------------------------- | -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Blockchain`           | `string`                                                                     | Required | Blockchain name                                                                                                                                                                         |
| `Network`              | `string`                                                                     | Required | If a blockchain has a specific chain-id or network identifier, it should go in this field. It is up to the client to determine which network-specific identifier is mainnet or testnet. |
| `SubNetworkIdentifier` | [`*models.SubNetworkIdentifier`](../../doc/models/sub-network-identifier.md) | Optional | In blockchains with sharded state, the SubNetworkIdentifier is required to query some object on a specific shard. This identifier is optional for all non-sharded blockchains.          |

## Example (as JSON)

```json
{
  "blockchain": "blockchain0",
  "network": "network4",
  "sub_network_identifier": {
    "network": "network6",
    "metadata": {
      "producer": "producer4"
    }
  }
}
```
