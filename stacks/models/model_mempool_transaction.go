package models

import (
	"encoding/json"
	"fmt"
)

// MempoolTransaction Describes all transaction types on Stacks 2.0 blockchain
type MempoolTransaction struct {
	MempoolCoinbaseTransaction         *MempoolCoinbaseTransaction
	MempoolContractCallTransaction     *MempoolContractCallTransaction
	MempoolPoisonMicroblockTransaction *MempoolPoisonMicroblockTransaction
	MempoolSmartContractTransaction    *MempoolSmartContractTransaction
	MempoolTenureChangeTransaction     *MempoolTenureChangeTransaction
	MempoolTokenTransferTransaction    *MempoolTokenTransferTransaction
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MempoolTransaction) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MempoolCoinbaseTransaction
	err = json.Unmarshal(data, &dst.MempoolCoinbaseTransaction)
	if err == nil {
		jsonMempoolCoinbaseTransaction, _ := json.Marshal(dst.MempoolCoinbaseTransaction)
		if string(jsonMempoolCoinbaseTransaction) == "{}" { // empty struct
			dst.MempoolCoinbaseTransaction = nil
		} else {
			return nil // data stored in dst.MempoolCoinbaseTransaction, return on the first match
		}
	} else {
		dst.MempoolCoinbaseTransaction = nil
	}

	// try to unmarshal JSON data into MempoolContractCallTransaction
	err = json.Unmarshal(data, &dst.MempoolContractCallTransaction)
	if err == nil {
		jsonMempoolContractCallTransaction, _ := json.Marshal(dst.MempoolContractCallTransaction)
		if string(jsonMempoolContractCallTransaction) == "{}" { // empty struct
			dst.MempoolContractCallTransaction = nil
		} else {
			return nil // data stored in dst.MempoolContractCallTransaction, return on the first match
		}
	} else {
		dst.MempoolContractCallTransaction = nil
	}

	// try to unmarshal JSON data into MempoolPoisonMicroblockTransaction
	err = json.Unmarshal(data, &dst.MempoolPoisonMicroblockTransaction)
	if err == nil {
		jsonMempoolPoisonMicroblockTransaction, _ := json.Marshal(dst.MempoolPoisonMicroblockTransaction)
		if string(jsonMempoolPoisonMicroblockTransaction) == "{}" { // empty struct
			dst.MempoolPoisonMicroblockTransaction = nil
		} else {
			return nil // data stored in dst.MempoolPoisonMicroblockTransaction, return on the first match
		}
	} else {
		dst.MempoolPoisonMicroblockTransaction = nil
	}

	// try to unmarshal JSON data into MempoolSmartContractTransaction
	err = json.Unmarshal(data, &dst.MempoolSmartContractTransaction)
	if err == nil {
		jsonMempoolSmartContractTransaction, _ := json.Marshal(dst.MempoolSmartContractTransaction)
		if string(jsonMempoolSmartContractTransaction) == "{}" { // empty struct
			dst.MempoolSmartContractTransaction = nil
		} else {
			return nil // data stored in dst.MempoolSmartContractTransaction, return on the first match
		}
	} else {
		dst.MempoolSmartContractTransaction = nil
	}

	// try to unmarshal JSON data into MempoolTenureChangeTransaction
	err = json.Unmarshal(data, &dst.MempoolTenureChangeTransaction)
	if err == nil {
		jsonMempoolTenureChangeTransaction, _ := json.Marshal(dst.MempoolTenureChangeTransaction)
		if string(jsonMempoolTenureChangeTransaction) == "{}" { // empty struct
			dst.MempoolTenureChangeTransaction = nil
		} else {
			return nil // data stored in dst.MempoolTenureChangeTransaction, return on the first match
		}
	} else {
		dst.MempoolTenureChangeTransaction = nil
	}

	// try to unmarshal JSON data into MempoolTokenTransferTransaction
	err = json.Unmarshal(data, &dst.MempoolTokenTransferTransaction)
	if err == nil {
		jsonMempoolTokenTransferTransaction, _ := json.Marshal(dst.MempoolTokenTransferTransaction)
		if string(jsonMempoolTokenTransferTransaction) == "{}" { // empty struct
			dst.MempoolTokenTransferTransaction = nil
		} else {
			return nil // data stored in dst.MempoolTokenTransferTransaction, return on the first match
		}
	} else {
		dst.MempoolTokenTransferTransaction = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(MempoolTransaction)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MempoolTransaction) MarshalJSON() ([]byte, error) {
	if src.MempoolCoinbaseTransaction != nil {
		return json.Marshal(&src.MempoolCoinbaseTransaction)
	}

	if src.MempoolContractCallTransaction != nil {
		return json.Marshal(&src.MempoolContractCallTransaction)
	}

	if src.MempoolPoisonMicroblockTransaction != nil {
		return json.Marshal(&src.MempoolPoisonMicroblockTransaction)
	}

	if src.MempoolSmartContractTransaction != nil {
		return json.Marshal(&src.MempoolSmartContractTransaction)
	}

	if src.MempoolTenureChangeTransaction != nil {
		return json.Marshal(&src.MempoolTenureChangeTransaction)
	}

	if src.MempoolTokenTransferTransaction != nil {
		return json.Marshal(&src.MempoolTokenTransferTransaction)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMempoolTransaction struct {
	value *MempoolTransaction
	isSet bool
}

func (v NullableMempoolTransaction) Get() *MempoolTransaction {
	return v.value
}

func (v *NullableMempoolTransaction) Set(val *MempoolTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableMempoolTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableMempoolTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMempoolTransaction(val *MempoolTransaction) *NullableMempoolTransaction {
	return &NullableMempoolTransaction{value: val, isSet: true}
}

func (v NullableMempoolTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMempoolTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
