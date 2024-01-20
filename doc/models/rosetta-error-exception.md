# Rosetta Error Exception

Instead of utilizing HTTP status codes to describe node errors (which often do not have a good analog), rich errors are returned using this object. Both the code and message fields can be individually used to correctly identify an error. Implementations MUST use unique values for both fields.

## Structure

`RosettaErrorException`

## Fields

| Name        | Type                                             | Tags     | Description                                                                                                                                                                                   |
| ----------- | ------------------------------------------------ | -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Code`      | `int`                                            | Required | Code is a network-specific error code. If desired, this code can be equivalent to an HTTP status code.                                                                                        |
| `Message`   | `string`                                         | Required | Message is a network-specific error message. The message MUST NOT change for a given code. In particular, this means that any contextual information should be included in the details field. |
| `Retriable` | `bool`                                           | Required | An error is retriable if the same request may succeed if submitted again.                                                                                                                     |
| `Details`   | [`*models.Details`](../../doc/models/details.md) | Optional | Often times it is useful to return context specific to the request that caused the error (i.e. a sample of the stack trace or impacted account) in addition to the standard error message.    |

## Example (as JSON)

```json
{
  "code": 36,
  "message": "message2",
  "retriable": false,
  "details": {
    "address": "address6",
    "error": "error4"
  }
}
```
