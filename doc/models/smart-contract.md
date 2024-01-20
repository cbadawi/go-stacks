# Smart Contract

## Structure

`SmartContract`

## Fields

| Name             | Type                | Tags     | Description                                                                                             |
| ---------------- | ------------------- | -------- | ------------------------------------------------------------------------------------------------------- |
| `ClarityVersion` | `Optional[float64]` | Optional | The Clarity version of the contract, only specified for versioned contract transactions, otherwise null |
| `ContractId`     | `string`            | Required | Contract identifier formatted as `<principaladdress>.<contract_name>`                                   |
| `SourceCode`     | `string`            | Required | Clarity code of the smart contract being deployed                                                       |

## Example (as JSON)

```json
{
  "clarity_version": 178.16,
  "contract_id": "contract_id4",
  "source_code": "source_code0"
}
```
