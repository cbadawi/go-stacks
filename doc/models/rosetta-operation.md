# Rosetta Operation

Operations contain all balance-changing information within a transaction. They are always one-sided (only affect 1 AccountIdentifier) and can succeed or fail independently from a Transaction.

## Structure

`RosettaOperation`

## Fields

| Name                  | Type                                                                                    | Tags     | Description                                                                                                                                                                                                                                                                                                                                                                                                |
| --------------------- | --------------------------------------------------------------------------------------- | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `OperationIdentifier` | [`models.RosettaOperationIdentifier`](../../doc/models/rosetta-operation-identifier.md) | Required | The operation_identifier uniquely identifies an operation within a transaction.                                                                                                                                                                                                                                                                                                                            |
| `RelatedOperations`   | [`[]models.RosettaRelatedOperation`](../../doc/models/rosetta-related-operation.md)     | Optional | Restrict referenced related_operations to identifier indexes < the current operation_identifier.index. This ensures there exists a clear DAG-structure of relations. Since operations are one-sided, one could imagine relating operations in a single transfer or linking operations in a call tree.                                                                                                      |
| `Type`                | `string`                                                                                | Required | The network-specific type of the operation. Ensure that any type that can be returned here is also specified in the NetworkStatus. This can be very useful to downstream consumers that parse all block data.                                                                                                                                                                                              |
| `Status`              | `*string`                                                                               | Optional | The network-specific status of the operation. Status is not defined on the transaction object because blockchains with smart contracts may have transactions that partially apply. Blockchains with atomic transactions (all operations succeed or all operations fail) will have the same status for each operation.                                                                                      |
| `Account`             | [`*models.RosettaAccount`](../../doc/models/rosetta-account.md)                         | Optional | The account_identifier uniquely identifies an account within a network. All fields in the account_identifier are utilized to determine this uniqueness (including the metadata field, if populated).                                                                                                                                                                                                       |
| `Amount`              | [`*models.RosettaAmount`](../../doc/models/rosetta-amount.md)                           | Optional | Amount is some Value of a Currency. It is considered invalid to specify a Value without a Currency.                                                                                                                                                                                                                                                                                                        |
| `CoinChange`          | [`*models.RosettaCoinChange`](../../doc/models/rosetta-coin-change.md)                  | Optional | CoinChange is used to represent a change in state of a some coin identified by a coin_identifier. This object is part of the Operation model and must be populated for UTXO-based blockchains. Coincidentally, this abstraction of UTXOs allows for supporting both account-based transfers and UTXO-based transfers on the same blockchain (when a transfer is account-based, don't populate this model). |
| `Metadata`            | `*interface{}`                                                                          | Optional | Operations Meta Data                                                                                                                                                                                                                                                                                                                                                                                       |

## Example (as JSON)

```json
{
  "operation_identifier": {
    "index": 254,
    "network_index": 122
  },
  "related_operations": [
    {
      "index": 246,
      "network_index": 130
    },
    {
      "index": 246,
      "network_index": 130
    }
  ],
  "type": "type0",
  "status": "status2",
  "account": {
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
  "amount": {
    "value": "value0",
    "currency": {
      "symbol": "symbol8",
      "decimals": 194,
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
  "coin_change": {
    "coin_identifier": {
      "identifier": "identifier8"
    },
    "coin_action": "coin_created"
  }
}
```
