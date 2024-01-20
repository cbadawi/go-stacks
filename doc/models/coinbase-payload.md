# Coinbase Payload

## Structure

`CoinbasePayload`

## Fields

| Name           | Type               | Tags     | Description                                                                                                                                                                                                                |
| -------------- | ------------------ | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Data`         | `string`           | Required | Hex encoded 32-byte scratch space for block leader's use                                                                                                                                                                   |
| `AltRecipient` | `Optional[string]` | Optional | A principal that will receive the miner rewards for this coinbase transaction. Can be either a standard principal or contract principal. Only specified for `coinbase-to-alt-recipient` transaction types, otherwise null. |
| `VrfProof`     | `Optional[string]` | Optional | Hex encoded 80-byte VRF proof                                                                                                                                                                                              |

## Example (as JSON)

```json
{
  "data": "data0",
  "alt_recipient": "alt_recipient2",
  "vrf_proof": "vrf_proof0"
}
```
