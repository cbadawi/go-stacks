package models

import (
	"encoding/json"
)

// RosettaOptions represents a RosettaOptions struct.
// The options that will be sent directly to /construction/metadata by the caller.
type RosettaOptions struct {
	// sender's address
	SenderAddress *string `json:"sender_address,omitempty"`
	// Type of operation e.g transfer
	Type *string `json:"type,omitempty"`
	// This value indicates the state of the operations
	Status *string `json:"status,omitempty"`
	// Recipient's address
	TokenTransferRecipientAddress *string `json:"token_transfer_recipient_address,omitempty"`
	// Amount to be transfered.
	Amount *string `json:"amount,omitempty"`
	// Currency symbol e.g STX
	Symbol *string `json:"symbol,omitempty"`
	// Number of decimal places
	Decimals *int `json:"decimals,omitempty"`
	// Maximum price a user is willing to pay.
	GasLimit *float64 `json:"gas_limit,omitempty"`
	// Cost necessary to perform a transaction on the network
	GasPrice *float64 `json:"gas_price,omitempty"`
	// A suggested fee multiplier to indicate that the suggested fee should be scaled. This may be used to set higher fees for urgent transactions or to pay lower fees when there is less urgency.
	SuggestedFeeMultiplier *float64 `json:"suggested_fee_multiplier,omitempty"`
	// Maximum fee user is willing to pay
	MaxFee *string `json:"max_fee,omitempty"`
	// Fee for this transaction
	Fee *string `json:"fee,omitempty"`
	// Transaction approximative size (used to calculate total fee).
	Size *int `json:"size,omitempty"`
	// STX token transfer memo.
	Memo *string `json:"memo,omitempty"`
	// Number of cycles when stacking.
	NumberOfCycles *int `json:"number_of_cycles,omitempty"`
	// Address of the contract to call.
	ContractAddress *string `json:"contract_address,omitempty"`
	// Name of the contract to call.
	ContractName *string `json:"contract_name,omitempty"`
	// Set the burnchain (BTC) block for stacking lock to start.
	BurnBlockHeight *int `json:"burn_block_height,omitempty"`
	// Delegator address for when calling `delegate-stacking`.
	DelegateTo *string `json:"delegate_to,omitempty"`
	// The reward address for stacking transaction. It should be a valid Bitcoin address
	PoxAddr *string `json:"pox_addr,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaOptions.
// It customizes the JSON marshaling process for RosettaOptions objects.
func (r *RosettaOptions) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaOptions object to a map representation for JSON marshaling.
func (r *RosettaOptions) toMap() map[string]any {
	structMap := make(map[string]any)
	if r.SenderAddress != nil {
		structMap["sender_address"] = r.SenderAddress
	}
	if r.Type != nil {
		structMap["type"] = r.Type
	}
	if r.Status != nil {
		structMap["status"] = r.Status
	}
	if r.TokenTransferRecipientAddress != nil {
		structMap["token_transfer_recipient_address"] = r.TokenTransferRecipientAddress
	}
	if r.Amount != nil {
		structMap["amount"] = r.Amount
	}
	if r.Symbol != nil {
		structMap["symbol"] = r.Symbol
	}
	if r.Decimals != nil {
		structMap["decimals"] = r.Decimals
	}
	if r.GasLimit != nil {
		structMap["gas_limit"] = r.GasLimit
	}
	if r.GasPrice != nil {
		structMap["gas_price"] = r.GasPrice
	}
	if r.SuggestedFeeMultiplier != nil {
		structMap["suggested_fee_multiplier"] = r.SuggestedFeeMultiplier
	}
	if r.MaxFee != nil {
		structMap["max_fee"] = r.MaxFee
	}
	if r.Fee != nil {
		structMap["fee"] = r.Fee
	}
	if r.Size != nil {
		structMap["size"] = r.Size
	}
	if r.Memo != nil {
		structMap["memo"] = r.Memo
	}
	if r.NumberOfCycles != nil {
		structMap["number_of_cycles"] = r.NumberOfCycles
	}
	if r.ContractAddress != nil {
		structMap["contract_address"] = r.ContractAddress
	}
	if r.ContractName != nil {
		structMap["contract_name"] = r.ContractName
	}
	if r.BurnBlockHeight != nil {
		structMap["burn_block_height"] = r.BurnBlockHeight
	}
	if r.DelegateTo != nil {
		structMap["delegate_to"] = r.DelegateTo
	}
	if r.PoxAddr != nil {
		structMap["pox_addr"] = r.PoxAddr
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaOptions.
// It customizes the JSON unmarshaling process for RosettaOptions objects.
func (r *RosettaOptions) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SenderAddress                 *string  `json:"sender_address,omitempty"`
		Type                          *string  `json:"type,omitempty"`
		Status                        *string  `json:"status,omitempty"`
		TokenTransferRecipientAddress *string  `json:"token_transfer_recipient_address,omitempty"`
		Amount                        *string  `json:"amount,omitempty"`
		Symbol                        *string  `json:"symbol,omitempty"`
		Decimals                      *int     `json:"decimals,omitempty"`
		GasLimit                      *float64 `json:"gas_limit,omitempty"`
		GasPrice                      *float64 `json:"gas_price,omitempty"`
		SuggestedFeeMultiplier        *float64 `json:"suggested_fee_multiplier,omitempty"`
		MaxFee                        *string  `json:"max_fee,omitempty"`
		Fee                           *string  `json:"fee,omitempty"`
		Size                          *int     `json:"size,omitempty"`
		Memo                          *string  `json:"memo,omitempty"`
		NumberOfCycles                *int     `json:"number_of_cycles,omitempty"`
		ContractAddress               *string  `json:"contract_address,omitempty"`
		ContractName                  *string  `json:"contract_name,omitempty"`
		BurnBlockHeight               *int     `json:"burn_block_height,omitempty"`
		DelegateTo                    *string  `json:"delegate_to,omitempty"`
		PoxAddr                       *string  `json:"pox_addr,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.SenderAddress = temp.SenderAddress
	r.Type = temp.Type
	r.Status = temp.Status
	r.TokenTransferRecipientAddress = temp.TokenTransferRecipientAddress
	r.Amount = temp.Amount
	r.Symbol = temp.Symbol
	r.Decimals = temp.Decimals
	r.GasLimit = temp.GasLimit
	r.GasPrice = temp.GasPrice
	r.SuggestedFeeMultiplier = temp.SuggestedFeeMultiplier
	r.MaxFee = temp.MaxFee
	r.Fee = temp.Fee
	r.Size = temp.Size
	r.Memo = temp.Memo
	r.NumberOfCycles = temp.NumberOfCycles
	r.ContractAddress = temp.ContractAddress
	r.ContractName = temp.ContractName
	r.BurnBlockHeight = temp.BurnBlockHeight
	r.DelegateTo = temp.DelegateTo
	r.PoxAddr = temp.PoxAddr
	return nil
}
