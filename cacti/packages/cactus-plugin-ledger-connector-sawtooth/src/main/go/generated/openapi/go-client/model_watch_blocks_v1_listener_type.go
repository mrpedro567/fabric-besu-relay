/*
Hyperledger Cacti Plugin - Connector Sawtooth

Can perform basic tasks on a Sawtooth ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-sawtooth

import (
	"encoding/json"
	"fmt"
)

// WatchBlocksV1ListenerType Response type from WatchBlocks. 'Cacti*' are custom views, others correspond to plain sawtooth data.
type WatchBlocksV1ListenerType string

// List of WatchBlocksV1ListenerType
const (
	Full WatchBlocksV1ListenerType = "full"
	CactiTransactions WatchBlocksV1ListenerType = "cacti:transactions"
)

// All allowed values of WatchBlocksV1ListenerType enum
var AllowedWatchBlocksV1ListenerTypeEnumValues = []WatchBlocksV1ListenerType{
	"full",
	"cacti:transactions",
}

func (v *WatchBlocksV1ListenerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WatchBlocksV1ListenerType(value)
	for _, existing := range AllowedWatchBlocksV1ListenerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WatchBlocksV1ListenerType", value)
}

// NewWatchBlocksV1ListenerTypeFromValue returns a pointer to a valid WatchBlocksV1ListenerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWatchBlocksV1ListenerTypeFromValue(v string) (*WatchBlocksV1ListenerType, error) {
	ev := WatchBlocksV1ListenerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WatchBlocksV1ListenerType: valid values are %v", v, AllowedWatchBlocksV1ListenerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WatchBlocksV1ListenerType) IsValid() bool {
	for _, existing := range AllowedWatchBlocksV1ListenerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WatchBlocksV1ListenerType value
func (v WatchBlocksV1ListenerType) Ptr() *WatchBlocksV1ListenerType {
	return &v
}

type NullableWatchBlocksV1ListenerType struct {
	value *WatchBlocksV1ListenerType
	isSet bool
}

func (v NullableWatchBlocksV1ListenerType) Get() *WatchBlocksV1ListenerType {
	return v.value
}

func (v *NullableWatchBlocksV1ListenerType) Set(val *WatchBlocksV1ListenerType) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchBlocksV1ListenerType) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchBlocksV1ListenerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchBlocksV1ListenerType(val *WatchBlocksV1ListenerType) *NullableWatchBlocksV1ListenerType {
	return &NullableWatchBlocksV1ListenerType{value: val, isSet: true}
}

func (v NullableWatchBlocksV1ListenerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchBlocksV1ListenerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

