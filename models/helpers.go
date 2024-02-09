package models

import (
	"encoding/json"
	"errors"
)

// Handle case when json: cannot unmarshal string into Go value of type struct
// api body likely contains error description
func unmarshalResponseString(input []byte) error {
	var str string
	if err := json.Unmarshal(input, &str); err != nil {
		return err
	}
	return errors.New(str)
}
