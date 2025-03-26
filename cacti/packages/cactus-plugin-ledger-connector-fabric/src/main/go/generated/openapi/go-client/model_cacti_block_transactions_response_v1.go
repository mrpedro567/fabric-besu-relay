/*
Hyperledger Cactus Plugin - Connector Fabric

Can perform basic tasks on a fabric ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-fabric

import (
	"encoding/json"
)

// checks if the CactiBlockTransactionsResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CactiBlockTransactionsResponseV1{}

// CactiBlockTransactionsResponseV1 Custom response containing block transactions summary. Compatible with legacy fabric-socketio connector monitoring.
type CactiBlockTransactionsResponseV1 struct {
	// List of transactions summary
	CactiTransactionsEvents []CactiBlockTransactionEventV1 `json:"cactiTransactionsEvents"`
}

// NewCactiBlockTransactionsResponseV1 instantiates a new CactiBlockTransactionsResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCactiBlockTransactionsResponseV1(cactiTransactionsEvents []CactiBlockTransactionEventV1) *CactiBlockTransactionsResponseV1 {
	this := CactiBlockTransactionsResponseV1{}
	this.CactiTransactionsEvents = cactiTransactionsEvents
	return &this
}

// NewCactiBlockTransactionsResponseV1WithDefaults instantiates a new CactiBlockTransactionsResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCactiBlockTransactionsResponseV1WithDefaults() *CactiBlockTransactionsResponseV1 {
	this := CactiBlockTransactionsResponseV1{}
	return &this
}

// GetCactiTransactionsEvents returns the CactiTransactionsEvents field value
func (o *CactiBlockTransactionsResponseV1) GetCactiTransactionsEvents() []CactiBlockTransactionEventV1 {
	if o == nil {
		var ret []CactiBlockTransactionEventV1
		return ret
	}

	return o.CactiTransactionsEvents
}

// GetCactiTransactionsEventsOk returns a tuple with the CactiTransactionsEvents field value
// and a boolean to check if the value has been set.
func (o *CactiBlockTransactionsResponseV1) GetCactiTransactionsEventsOk() ([]CactiBlockTransactionEventV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.CactiTransactionsEvents, true
}

// SetCactiTransactionsEvents sets field value
func (o *CactiBlockTransactionsResponseV1) SetCactiTransactionsEvents(v []CactiBlockTransactionEventV1) {
	o.CactiTransactionsEvents = v
}

func (o CactiBlockTransactionsResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CactiBlockTransactionsResponseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cactiTransactionsEvents"] = o.CactiTransactionsEvents
	return toSerialize, nil
}

type NullableCactiBlockTransactionsResponseV1 struct {
	value *CactiBlockTransactionsResponseV1
	isSet bool
}

func (v NullableCactiBlockTransactionsResponseV1) Get() *CactiBlockTransactionsResponseV1 {
	return v.value
}

func (v *NullableCactiBlockTransactionsResponseV1) Set(val *CactiBlockTransactionsResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableCactiBlockTransactionsResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableCactiBlockTransactionsResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCactiBlockTransactionsResponseV1(val *CactiBlockTransactionsResponseV1) *NullableCactiBlockTransactionsResponseV1 {
	return &NullableCactiBlockTransactionsResponseV1{value: val, isSet: true}
}

func (v NullableCactiBlockTransactionsResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCactiBlockTransactionsResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


