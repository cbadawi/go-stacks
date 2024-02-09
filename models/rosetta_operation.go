package models

import (
	"encoding/json"
)

// RosettaOperation represents a RosettaOperation struct.
// Operations contain all balance-changing information within a transaction. They are always one-sided (only affect 1 AccountIdentifier) and can succeed or fail independently from a Transaction.
type RosettaOperation struct {
	// The operation_identifier uniquely identifies an operation within a transaction.
	OperationIdentifier RosettaOperationIdentifier `json:"operation_identifier"`
	// Restrict referenced related_operations to identifier indexes < the current operation_identifier.index. This ensures there exists a clear DAG-structure of relations. Since operations are one-sided, one could imagine relating operations in a single transfer or linking operations in a call tree.
	RelatedOperations []RosettaRelatedOperation `json:"related_operations,omitempty"`
	// The network-specific type of the operation. Ensure that any type that can be returned here is also specified in the NetworkStatus. This can be very useful to downstream consumers that parse all block data.
	Type string `json:"type"`
	// The network-specific status of the operation. Status is not defined on the transaction object because blockchains with smart contracts may have transactions that partially apply. Blockchains with atomic transactions (all operations succeed or all operations fail) will have the same status for each operation.
	Status *string `json:"status,omitempty"`
	// The account_identifier uniquely identifies an account within a network. All fields in the account_identifier are utilized to determine this uniqueness (including the metadata field, if populated).
	Account *RosettaAccount `json:"account,omitempty"`
	// Amount is some Value of a Currency. It is considered invalid to specify a Value without a Currency.
	Amount *RosettaAmount `json:"amount,omitempty"`
	// CoinChange is used to represent a change in state of a some coin identified by a coin_identifier. This object is part of the Operation model and must be populated for UTXO-based blockchains. Coincidentally, this abstraction of UTXOs allows for supporting both account-based transfers and UTXO-based transfers on the same blockchain (when a transfer is account-based, don't populate this model).
	CoinChange *RosettaCoinChange `json:"coin_change,omitempty"`
	// Operations Meta Data
	Metadata *interface{} `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaOperation.
// It customizes the JSON marshaling process for RosettaOperation objects.
func (r *RosettaOperation) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaOperation object to a map representation for JSON marshaling.
func (r *RosettaOperation) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["operation_identifier"] = r.OperationIdentifier
	if r.RelatedOperations != nil {
		structMap["related_operations"] = r.RelatedOperations
	}
	structMap["type"] = r.Type
	if r.Status != nil {
		structMap["status"] = r.Status
	}
	if r.Account != nil {
		structMap["account"] = r.Account
	}
	if r.Amount != nil {
		structMap["amount"] = r.Amount
	}
	if r.CoinChange != nil {
		structMap["coin_change"] = r.CoinChange
	}
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaOperation.
// It customizes the JSON unmarshaling process for RosettaOperation objects.
func (r *RosettaOperation) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		OperationIdentifier RosettaOperationIdentifier `json:"operation_identifier"`
		RelatedOperations   []RosettaRelatedOperation  `json:"related_operations,omitempty"`
		Type                string                     `json:"type"`
		Status              *string                    `json:"status,omitempty"`
		Account             *RosettaAccount            `json:"account,omitempty"`
		Amount              *RosettaAmount             `json:"amount,omitempty"`
		CoinChange          *RosettaCoinChange         `json:"coin_change,omitempty"`
		Metadata            *interface{}               `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.OperationIdentifier = temp.OperationIdentifier
	r.RelatedOperations = temp.RelatedOperations
	r.Type = temp.Type
	r.Status = temp.Status
	r.Account = temp.Account
	r.Amount = temp.Amount
	r.CoinChange = temp.CoinChange
	r.Metadata = temp.Metadata
	return nil
}
