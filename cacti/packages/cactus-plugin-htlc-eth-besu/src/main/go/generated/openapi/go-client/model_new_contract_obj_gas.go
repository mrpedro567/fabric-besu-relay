/*
Hyperledger Cactus Plugin - HTLC-ETH Besu

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-htlc-eth-besu

import (
	"encoding/json"
	"fmt"
)

// NewContractObjGas - struct for NewContractObjGas
type NewContractObjGas struct {
	Float32 *float32
	String *string
}

// float32AsNewContractObjGas is a convenience function that returns float32 wrapped in NewContractObjGas
func Float32AsNewContractObjGas(v *float32) NewContractObjGas {
	return NewContractObjGas{
		Float32: v,
	}
}

// stringAsNewContractObjGas is a convenience function that returns string wrapped in NewContractObjGas
func StringAsNewContractObjGas(v *string) NewContractObjGas {
	return NewContractObjGas{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *NewContractObjGas) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Float32
	err = newStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			match++
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Float32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(NewContractObjGas)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(NewContractObjGas)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NewContractObjGas) MarshalJSON() ([]byte, error) {
	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NewContractObjGas) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Float32 != nil {
		return obj.Float32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableNewContractObjGas struct {
	value *NewContractObjGas
	isSet bool
}

func (v NullableNewContractObjGas) Get() *NewContractObjGas {
	return v.value
}

func (v *NullableNewContractObjGas) Set(val *NewContractObjGas) {
	v.value = val
	v.isSet = true
}

func (v NullableNewContractObjGas) IsSet() bool {
	return v.isSet
}

func (v *NullableNewContractObjGas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewContractObjGas(val *NewContractObjGas) *NullableNewContractObjGas {
	return &NullableNewContractObjGas{value: val, isSet: true}
}

func (v NullableNewContractObjGas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewContractObjGas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


