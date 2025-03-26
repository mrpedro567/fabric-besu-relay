/*
Hyperledger Cacti Plugin - Connector Corda

Can perform basic tasks on a Corda ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-corda

import (
	"encoding/json"
	"time"
)

// checks if the ListCpiV1ResponseCpisInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCpiV1ResponseCpisInner{}

// ListCpiV1ResponseCpisInner struct for ListCpiV1ResponseCpisInner
type ListCpiV1ResponseCpisInner struct {
	CpiFileChecksum *string `json:"cpiFileChecksum,omitempty"`
	CpiFileFullChecksum *string `json:"cpiFileFullChecksum,omitempty"`
	Cpks []ListCpiV1ResponseCpisInnerCpksInner `json:"cpks,omitempty"`
	GroupPolicy NullableString `json:"groupPolicy,omitempty"`
	Id *CPIIDV1 `json:"id,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewListCpiV1ResponseCpisInner instantiates a new ListCpiV1ResponseCpisInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCpiV1ResponseCpisInner() *ListCpiV1ResponseCpisInner {
	this := ListCpiV1ResponseCpisInner{}
	return &this
}

// NewListCpiV1ResponseCpisInnerWithDefaults instantiates a new ListCpiV1ResponseCpisInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCpiV1ResponseCpisInnerWithDefaults() *ListCpiV1ResponseCpisInner {
	this := ListCpiV1ResponseCpisInner{}
	return &this
}

// GetCpiFileChecksum returns the CpiFileChecksum field value if set, zero value otherwise.
func (o *ListCpiV1ResponseCpisInner) GetCpiFileChecksum() string {
	if o == nil || IsNil(o.CpiFileChecksum) {
		var ret string
		return ret
	}
	return *o.CpiFileChecksum
}

// GetCpiFileChecksumOk returns a tuple with the CpiFileChecksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCpiV1ResponseCpisInner) GetCpiFileChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.CpiFileChecksum) {
		return nil, false
	}
	return o.CpiFileChecksum, true
}

// HasCpiFileChecksum returns a boolean if a field has been set.
func (o *ListCpiV1ResponseCpisInner) HasCpiFileChecksum() bool {
	if o != nil && !IsNil(o.CpiFileChecksum) {
		return true
	}

	return false
}

// SetCpiFileChecksum gets a reference to the given string and assigns it to the CpiFileChecksum field.
func (o *ListCpiV1ResponseCpisInner) SetCpiFileChecksum(v string) {
	o.CpiFileChecksum = &v
}

// GetCpiFileFullChecksum returns the CpiFileFullChecksum field value if set, zero value otherwise.
func (o *ListCpiV1ResponseCpisInner) GetCpiFileFullChecksum() string {
	if o == nil || IsNil(o.CpiFileFullChecksum) {
		var ret string
		return ret
	}
	return *o.CpiFileFullChecksum
}

// GetCpiFileFullChecksumOk returns a tuple with the CpiFileFullChecksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCpiV1ResponseCpisInner) GetCpiFileFullChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.CpiFileFullChecksum) {
		return nil, false
	}
	return o.CpiFileFullChecksum, true
}

// HasCpiFileFullChecksum returns a boolean if a field has been set.
func (o *ListCpiV1ResponseCpisInner) HasCpiFileFullChecksum() bool {
	if o != nil && !IsNil(o.CpiFileFullChecksum) {
		return true
	}

	return false
}

// SetCpiFileFullChecksum gets a reference to the given string and assigns it to the CpiFileFullChecksum field.
func (o *ListCpiV1ResponseCpisInner) SetCpiFileFullChecksum(v string) {
	o.CpiFileFullChecksum = &v
}

// GetCpks returns the Cpks field value if set, zero value otherwise.
func (o *ListCpiV1ResponseCpisInner) GetCpks() []ListCpiV1ResponseCpisInnerCpksInner {
	if o == nil || IsNil(o.Cpks) {
		var ret []ListCpiV1ResponseCpisInnerCpksInner
		return ret
	}
	return o.Cpks
}

// GetCpksOk returns a tuple with the Cpks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCpiV1ResponseCpisInner) GetCpksOk() ([]ListCpiV1ResponseCpisInnerCpksInner, bool) {
	if o == nil || IsNil(o.Cpks) {
		return nil, false
	}
	return o.Cpks, true
}

// HasCpks returns a boolean if a field has been set.
func (o *ListCpiV1ResponseCpisInner) HasCpks() bool {
	if o != nil && !IsNil(o.Cpks) {
		return true
	}

	return false
}

// SetCpks gets a reference to the given []ListCpiV1ResponseCpisInnerCpksInner and assigns it to the Cpks field.
func (o *ListCpiV1ResponseCpisInner) SetCpks(v []ListCpiV1ResponseCpisInnerCpksInner) {
	o.Cpks = v
}

// GetGroupPolicy returns the GroupPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListCpiV1ResponseCpisInner) GetGroupPolicy() string {
	if o == nil || IsNil(o.GroupPolicy.Get()) {
		var ret string
		return ret
	}
	return *o.GroupPolicy.Get()
}

// GetGroupPolicyOk returns a tuple with the GroupPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListCpiV1ResponseCpisInner) GetGroupPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupPolicy.Get(), o.GroupPolicy.IsSet()
}

// HasGroupPolicy returns a boolean if a field has been set.
func (o *ListCpiV1ResponseCpisInner) HasGroupPolicy() bool {
	if o != nil && o.GroupPolicy.IsSet() {
		return true
	}

	return false
}

// SetGroupPolicy gets a reference to the given NullableString and assigns it to the GroupPolicy field.
func (o *ListCpiV1ResponseCpisInner) SetGroupPolicy(v string) {
	o.GroupPolicy.Set(&v)
}
// SetGroupPolicyNil sets the value for GroupPolicy to be an explicit nil
func (o *ListCpiV1ResponseCpisInner) SetGroupPolicyNil() {
	o.GroupPolicy.Set(nil)
}

// UnsetGroupPolicy ensures that no value is present for GroupPolicy, not even an explicit nil
func (o *ListCpiV1ResponseCpisInner) UnsetGroupPolicy() {
	o.GroupPolicy.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListCpiV1ResponseCpisInner) GetId() CPIIDV1 {
	if o == nil || IsNil(o.Id) {
		var ret CPIIDV1
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCpiV1ResponseCpisInner) GetIdOk() (*CPIIDV1, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListCpiV1ResponseCpisInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given CPIIDV1 and assigns it to the Id field.
func (o *ListCpiV1ResponseCpisInner) SetId(v CPIIDV1) {
	o.Id = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ListCpiV1ResponseCpisInner) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCpiV1ResponseCpisInner) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ListCpiV1ResponseCpisInner) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *ListCpiV1ResponseCpisInner) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o ListCpiV1ResponseCpisInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCpiV1ResponseCpisInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CpiFileChecksum) {
		toSerialize["cpiFileChecksum"] = o.CpiFileChecksum
	}
	if !IsNil(o.CpiFileFullChecksum) {
		toSerialize["cpiFileFullChecksum"] = o.CpiFileFullChecksum
	}
	if !IsNil(o.Cpks) {
		toSerialize["cpks"] = o.Cpks
	}
	if o.GroupPolicy.IsSet() {
		toSerialize["groupPolicy"] = o.GroupPolicy.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableListCpiV1ResponseCpisInner struct {
	value *ListCpiV1ResponseCpisInner
	isSet bool
}

func (v NullableListCpiV1ResponseCpisInner) Get() *ListCpiV1ResponseCpisInner {
	return v.value
}

func (v *NullableListCpiV1ResponseCpisInner) Set(val *ListCpiV1ResponseCpisInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListCpiV1ResponseCpisInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListCpiV1ResponseCpisInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCpiV1ResponseCpisInner(val *ListCpiV1ResponseCpisInner) *NullableListCpiV1ResponseCpisInner {
	return &NullableListCpiV1ResponseCpisInner{value: val, isSet: true}
}

func (v NullableListCpiV1ResponseCpisInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCpiV1ResponseCpisInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


