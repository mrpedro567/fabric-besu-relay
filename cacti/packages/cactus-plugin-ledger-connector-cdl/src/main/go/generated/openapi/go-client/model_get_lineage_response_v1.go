/*
Hyperledger Cacti Plugin - Connector CDL

Can perform basic tasks on Fujitsu CDL service.

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-cdl

import (
	"encoding/json"
)

// checks if the GetLineageResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLineageResponseV1{}

// GetLineageResponseV1 struct for GetLineageResponseV1
type GetLineageResponseV1 struct {
	Detail []TrailEventDetailsV1 `json:"detail"`
	Result String `json:"result"`
}

// NewGetLineageResponseV1 instantiates a new GetLineageResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLineageResponseV1(detail []TrailEventDetailsV1, result String) *GetLineageResponseV1 {
	this := GetLineageResponseV1{}
	this.Detail = detail
	this.Result = result
	return &this
}

// NewGetLineageResponseV1WithDefaults instantiates a new GetLineageResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLineageResponseV1WithDefaults() *GetLineageResponseV1 {
	this := GetLineageResponseV1{}
	return &this
}

// GetDetail returns the Detail field value
func (o *GetLineageResponseV1) GetDetail() []TrailEventDetailsV1 {
	if o == nil {
		var ret []TrailEventDetailsV1
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *GetLineageResponseV1) GetDetailOk() ([]TrailEventDetailsV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detail, true
}

// SetDetail sets field value
func (o *GetLineageResponseV1) SetDetail(v []TrailEventDetailsV1) {
	o.Detail = v
}

// GetResult returns the Result field value
func (o *GetLineageResponseV1) GetResult() String {
	if o == nil {
		var ret String
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *GetLineageResponseV1) GetResultOk() (*String, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *GetLineageResponseV1) SetResult(v String) {
	o.Result = v
}

func (o GetLineageResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLineageResponseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["detail"] = o.Detail
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableGetLineageResponseV1 struct {
	value *GetLineageResponseV1
	isSet bool
}

func (v NullableGetLineageResponseV1) Get() *GetLineageResponseV1 {
	return v.value
}

func (v *NullableGetLineageResponseV1) Set(val *GetLineageResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLineageResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLineageResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLineageResponseV1(val *GetLineageResponseV1) *NullableGetLineageResponseV1 {
	return &NullableGetLineageResponseV1{value: val, isSet: true}
}

func (v NullableGetLineageResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLineageResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


