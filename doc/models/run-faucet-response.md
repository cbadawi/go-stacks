# Run Faucet Response

POST request that initiates a transfer of tokens to a specified testnet address

## Structure

`RunFaucetResponse`

## Fields

| Name      | Type      | Tags     | Description                                  |
| --------- | --------- | -------- | -------------------------------------------- |
| `Success` | `bool`    | Required | Indicates if the faucet call was successful  |
| `TxId`    | `*string` | Optional | The transaction ID for the faucet call       |
| `TxRaw`   | `*string` | Optional | Raw transaction in hex string representation |

## Example (as JSON)

```json
{
  "success": false,
  "txId": "txId4",
  "txRaw": "txRaw4"
}
```
