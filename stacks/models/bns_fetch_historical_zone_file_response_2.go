package models

import (
	"encoding/json"
)

// BnsFetchHistoricalZoneFileResponse2 represents a BnsFetchHistoricalZoneFileResponse2 struct.
// Fetches the historical zonefile specified by the username and zone hash.
type BnsFetchHistoricalZoneFileResponse2 struct {
	Zonefile *string `json:"zonefile,omitempty"`
	Error    *string `json:"error,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for BnsFetchHistoricalZoneFileResponse2.
// It customizes the JSON marshaling process for BnsFetchHistoricalZoneFileResponse2 objects.
func (b *BnsFetchHistoricalZoneFileResponse2) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BnsFetchHistoricalZoneFileResponse2 object to a map representation for JSON marshaling.
func (b *BnsFetchHistoricalZoneFileResponse2) toMap() map[string]any {
	structMap := make(map[string]any)
	if b.Zonefile != nil {
		structMap["zonefile"] = b.Zonefile
	}
	if b.Error != nil {
		structMap["error"] = b.Error
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BnsFetchHistoricalZoneFileResponse2.
// It customizes the JSON unmarshaling process for BnsFetchHistoricalZoneFileResponse2 objects.
func (b *BnsFetchHistoricalZoneFileResponse2) UnmarshalJSON(input []byte) error {
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
