
# Rosetta Account Balance Request

An AccountBalanceRequest is utilized to make a balance request on the /account/balance endpoint. If the block_identifier is populated, a historical balance query should be performed.

## Structure

`RosettaAccountBalanceRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NetworkIdentifier` | [`models.NetworkIdentifier`](../../doc/models/network-identifier.md) | Required | The network_identifier specifies which network a particular object is associated with. |
| `AccountIdentifier` | [`models.RosettaAccount`](../../doc/models/rosetta-account.md) | Required | The account_identifier uniquely identifies an account within a network. All fields in the account_identifier are utilized to determine this uniqueness (including the metadata field, if populated). |
| `BlockIdentifier` | [`*models.RosettaPartialBlockIdentifier`](../../doc/models/rosetta-partial-block-identifier.md) | Optional | When fetching data by BlockIdentifier, it may be possible to only specify the index or hash. If neither property is specified, it is assumed that the client is making a request at the current block. |

## Example (as JSON)

```json
{
  "network_identifier": {
    "blockchain": "blockchain2",
    "network": "network6",
    "sub_network_identifier": {
      "network": "network6",
      "metadata": {
        "producer": "producer4"
      }
    }
  },
  "account_identifier": {
    "address": "address6",
    "sub_account": {
      "address": "address4",
      "metadata": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "metadata": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "block_identifier": {
    "hash": "hash6",
    "index": 54
  }
}
```

