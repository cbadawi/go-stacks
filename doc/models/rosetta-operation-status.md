# Rosetta Operation Status

OperationStatus is utilized to indicate which Operation status are considered successful.

## Structure

`RosettaOperationStatus`

## Fields

| Name         | Type     | Tags     | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| ------------ | -------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Status`     | `string` | Required | The status is the network-specific status of the operation.                                                                                                                                                                                                                                                                                                                                                                                                    |
| `Successful` | `bool`   | Required | An Operation is considered successful if the Operation.Amount should affect the Operation.Account. Some blockchains (like Bitcoin) only include successful operations in blocks but other blockchains (like Ethereum) include unsuccessful operations that incur a fee. To reconcile the computed balance from the stream of Operations, it is critical to understand which Operation.Status indicate an Operation is successful and should affect an Account. |

## Example (as JSON)

```json
{
  "status": "status0",
  "successful": false
}
```
