# Contract Interface Response

GET request to get contract interface

## Structure

`ContractInterfaceResponse`

## Fields

| Name                | Type            | Tags     | Description                                 |
| ------------------- | --------------- | -------- | ------------------------------------------- |
| `Functions`         | `[]interface{}` | Required | List of defined methods                     |
| `Variables`         | `[]interface{}` | Required | List of defined variables                   |
| `Maps`              | `[]interface{}` | Required | List of defined data-maps                   |
| `FungibleTokens`    | `[]interface{}` | Required | List of fungible tokens in the contract     |
| `NonFungibleTokens` | `[]interface{}` | Required | List of non-fungible tokens in the contract |

## Example (as JSON)

```json
{
  "functions": [
    {
      "key1": "val1",
      "key2": "val2"
    },
    {
      "key1": "val1",
      "key2": "val2"
    },
    {
      "key1": "val1",
      "key2": "val2"
    }
  ],
  "variables": [
    {
      "key1": "val1",
      "key2": "val2"
    },
    {
      "key1": "val1",
      "key2": "val2"
    },
    {
      "key1": "val1",
      "key2": "val2"
    }
  ],
  "maps": [
    {
      "key1": "val1",
      "key2": "val2"
    },
    {
      "key1": "val1",
      "key2": "val2"
    },
    {
      "key1": "val1",
      "key2": "val2"
    }
  ],
  "fungible_tokens": [
    {
      "key1": "val1",
      "key2": "val2"
    },
    {
      "key1": "val1",
      "key2": "val2"
    },
    {
      "key1": "val1",
      "key2": "val2"
    }
  ],
  "non_fungible_tokens": [
    {
      "key1": "val1",
      "key2": "val2"
    }
  ]
}
```
