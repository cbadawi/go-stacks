# Bns Fetch File Zone Response 2

Fetch a user's raw zone file. This only works for RFC-compliant zone files. This method returns an error for names that have non-standard zone files.

## Structure

`BnsFetchFileZoneResponse2`

## Fields

| Name       | Type      | Tags     | Description                      |
| ---------- | --------- | -------- | -------------------------------- |
| `Zonefile` | `*string` | Optional | **Constraints**: _Pattern_: `.+` |
| `Error`    | `*string` | Optional | **Constraints**: _Pattern_: `.+` |

## Example (as JSON)

```json
{
  "zonefile": "zonefile4",
  "error": "error0"
}
```
