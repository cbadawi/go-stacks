# Contract List Response

GET list of contracts

## Structure

`ContractListResponse`

## Fields

| Name      | Type                                                              | Tags     | Description                                                           |
| --------- | ----------------------------------------------------------------- | -------- | --------------------------------------------------------------------- |
| `Limit`   | `int`                                                             | Required | The number of contracts to return                                     |
| `Offset`  | `int`                                                             | Required | The number to contracts to skip (starting at `0`)<br>**Default**: `0` |
| `Results` | [`[]models.SmartContract7`](../../doc/models/smart-contract-7.md) | Required | -                                                                     |

## Example (as JSON)

```json
{
  "limit": 166,
  "offset": 0,
  "results": [
    {
      "tx_id": "tx_id0",
      "canonical": false,
      "contract_id": "contract_id4",
      "block_height": 52,
      "source_code": "source_code0",
      "abi": "abi8"
    }
  ]
}
```
