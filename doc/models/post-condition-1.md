# Post Condition 1

Post-conditionscan limit the damage done to a user's assets

## Structure

`PostCondition1`

## Fields

| Name            | Type                                                                                                               | Tags               | Description                                                                                                                           |
| --------------- | ------------------------------------------------------------------------------------------------------------------ | ------------------ | ------------------------------------------------------------------------------------------------------------------------------------- |
| `Principal`     | [`models.PostConditionPrincipal4`](../../doc/models/post-condition-principal-4.md)                                 | Required           | -                                                                                                                                     |
| `ConditionCode` | [`models.PostConditionFungibleConditionCodeEnum`](../../doc/models/post-condition-fungible-condition-code-enum.md) | Required           | A fungible condition code encodes a statement being made for either STX or a fungible token, with respect to the originating account. |
| `Amount`        | `string`                                                                                                           | Required           | -                                                                                                                                     |
| `Type`          | `string`                                                                                                           | Required, Constant | **Default**: `"stx"`                                                                                                                  |
| `Asset`         | [`models.Asset`](../../doc/models/asset.md)                                                                        | Required           | -                                                                                                                                     |
| `AssetValue`    | [`models.AssetValue`](../../doc/models/asset-value.md)                                                             | Required           | -                                                                                                                                     |

## Example (as JSON)

```json
{
  "principal": {
    "type_id": "principal_origin",
    "address": "address6",
    "contract_name": "contract_name2"
  },
  "condition_code": "sent_less_than_or_equal_to",
  "amount": "amount2",
  "type": "stx",
  "asset": {
    "asset_name": "asset_name8",
    "contract_address": "contract_address4",
    "contract_name": "contract_name2"
  },
  "asset_value": {
    "hex": "hex2",
    "repr": "repr2"
  }
}
```
