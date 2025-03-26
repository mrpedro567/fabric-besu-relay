/*
Hyperledger Cactus Plugin - Persistence Ethereum

Synchronizes state of an ethereum ledger into a DB that can later be viewed in GUI

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-persistence-ethereum

import (
	"encoding/json"
	"fmt"
)

// TokenTypeV1 the model 'TokenTypeV1'
type TokenTypeV1 string

// List of TokenTypeV1
const (
	ERC20 TokenTypeV1 = "erc20"
	ERC721 TokenTypeV1 = "erc721"
)

// All allowed values of TokenTypeV1 enum
var AllowedTokenTypeV1EnumValues = []TokenTypeV1{
	"erc20",
	"erc721",
}

func (v *TokenTypeV1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TokenTypeV1(value)
	for _, existing := range AllowedTokenTypeV1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TokenTypeV1", value)
}

// NewTokenTypeV1FromValue returns a pointer to a valid TokenTypeV1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTokenTypeV1FromValue(v string) (*TokenTypeV1, error) {
	ev := TokenTypeV1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TokenTypeV1: valid values are %v", v, AllowedTokenTypeV1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TokenTypeV1) IsValid() bool {
	for _, existing := range AllowedTokenTypeV1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TokenTypeV1 value
func (v TokenTypeV1) Ptr() *TokenTypeV1 {
	return &v
}

type NullableTokenTypeV1 struct {
	value *TokenTypeV1
	isSet bool
}

func (v NullableTokenTypeV1) Get() *TokenTypeV1 {
	return v.value
}

func (v *NullableTokenTypeV1) Set(val *TokenTypeV1) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenTypeV1) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenTypeV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenTypeV1(val *TokenTypeV1) *NullableTokenTypeV1 {
	return &NullableTokenTypeV1{value: val, isSet: true}
}

func (v NullableTokenTypeV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenTypeV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

