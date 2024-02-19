package models

import (
	"encoding/json"
)

// BnsFetchFileZoneResponse2 represents a BnsFetchFileZoneResponse2 struct.
// Fetch a user's raw zone file. This only works for RFC-compliant zone files. This method returns an error for names that have non-standard zone files.
type BnsFetchFileZoneResponse2 struct {
	Zonefile *string `json:"zonefile,omitempty"`
	Error    *string `json:"error,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for BnsFetchFileZoneResponse2.
// It customizes the JSON marshaling process for BnsFetchFileZoneResponse2 objects.
func (b *BnsFetchFileZoneResponse2) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BnsFetchFileZoneResponse2 object to a map representation for JSON marshaling.
func (b *BnsFetchFileZoneResponse2) toMap() map[string]any {
	structMap := make(map[string]any)
	if b.Zonefile != nil {
		structMap["zonefile"] = b.Zonefile
	}
	if b.Error != nil {
		structMap["error"] = b.Error
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BnsFetchFileZoneResponse2.
// It customizes the JSON unmarshaling process for BnsFetchFileZoneResponse2 objects.
func (b *BnsFetchFileZoneResponse2) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Zonefile *string `json:"zonefile,omitempty"`
		Error    *string `json:"error,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	b.Zonefile = temp.Zonefile
	b.Error = temp.Error
	return nil
}
