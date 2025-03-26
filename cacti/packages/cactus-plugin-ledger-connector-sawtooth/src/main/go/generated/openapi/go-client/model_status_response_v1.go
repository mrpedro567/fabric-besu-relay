/*
Hyperledger Cacti Plugin - Connector Sawtooth

Can perform basic tasks on a Sawtooth ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-sawtooth

import (
	"encoding/json"
)

// checks if the StatusResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusResponseV1{}

// StatusResponseV1 Response with plugin and validator status report.
type StatusResponseV1 struct {
	// Plugin instance id.
	InstanceId string `json:"instanceId"`
	// Version of connectors Open API Spec.
	OpenApiSpecVersion *string `json:"openApiSpecVersion,omitempty"`
	// True if endpoints were created, false otherwise
	Initialized *bool `json:"initialized,omitempty"`
	// Response from sawtooth Rest API status endpoint
	SawtoothStatus interface{} `json:"sawtoothStatus,omitempty"`
}

// NewStatusResponseV1 instantiates a new StatusResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusResponseV1(instanceId string) *StatusResponseV1 {
	this := StatusResponseV1{}
	this.InstanceId = instanceId
	return &this
}

// NewStatusResponseV1WithDefaults instantiates a new StatusResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusResponseV1WithDefaults() *StatusResponseV1 {
	this := StatusResponseV1{}
	return &this
}

// GetInstanceId returns the InstanceId field value
func (o *StatusResponseV1) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *StatusResponseV1) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *StatusResponseV1) SetInstanceId(v string) {
	o.InstanceId = v
}

// GetOpenApiSpecVersion returns the OpenApiSpecVersion field value if set, zero value otherwise.
func (o *StatusResponseV1) GetOpenApiSpecVersion() string {
	if o == nil || IsNil(o.OpenApiSpecVersion) {
		var ret string
		return ret
	}
	return *o.OpenApiSpecVersion
}

// GetOpenApiSpecVersionOk returns a tuple with the OpenApiSpecVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponseV1) GetOpenApiSpecVersionOk() (*string, bool) {
	if o == nil || IsNil(o.OpenApiSpecVersion) {
		return nil, false
	}
	return o.OpenApiSpecVersion, true
}

// HasOpenApiSpecVersion returns a boolean if a field has been set.
func (o *StatusResponseV1) HasOpenApiSpecVersion() bool {
	if o != nil && !IsNil(o.OpenApiSpecVersion) {
		return true
	}

	return false
}

// SetOpenApiSpecVersion gets a reference to the given string and assigns it to the OpenApiSpecVersion field.
func (o *StatusResponseV1) SetOpenApiSpecVersion(v string) {
	o.OpenApiSpecVersion = &v
}

// GetInitialized returns the Initialized field value if set, zero value otherwise.
func (o *StatusResponseV1) GetInitialized() bool {
	if o == nil || IsNil(o.Initialized) {
		var ret bool
		return ret
	}
	return *o.Initialized
}

// GetInitializedOk returns a tuple with the Initialized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponseV1) GetInitializedOk() (*bool, bool) {
	if o == nil || IsNil(o.Initialized) {
		return nil, false
	}
	return o.Initialized, true
}

// HasInitialized returns a boolean if a field has been set.
func (o *StatusResponseV1) HasInitialized() bool {
	if o != nil && !IsNil(o.Initialized) {
		return true
	}

	return false
}

// SetInitialized gets a reference to the given bool and assigns it to the Initialized field.
func (o *StatusResponseV1) SetInitialized(v bool) {
	o.Initialized = &v
}

// GetSawtoothStatus returns the SawtoothStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatusResponseV1) GetSawtoothStatus() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SawtoothStatus
}

// GetSawtoothStatusOk returns a tuple with the SawtoothStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatusResponseV1) GetSawtoothStatusOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SawtoothStatus) {
		return nil, false
	}
	return &o.SawtoothStatus, true
}

// HasSawtoothStatus returns a boolean if a field has been set.
func (o *StatusResponseV1) HasSawtoothStatus() bool {
	if o != nil && IsNil(o.SawtoothStatus) {
		return true
	}

	return false
}

// SetSawtoothStatus gets a reference to the given interface{} and assigns it to the SawtoothStatus field.
func (o *StatusResponseV1) SetSawtoothStatus(v interface{}) {
	o.SawtoothStatus = v
}

func (o StatusResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusResponseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instanceId"] = o.InstanceId
	if !IsNil(o.OpenApiSpecVersion) {
		toSerialize["openApiSpecVersion"] = o.OpenApiSpecVersion
	}
	if !IsNil(o.Initialized) {
		toSerialize["initialized"] = o.Initialized
	}
	if o.SawtoothStatus != nil {
		toSerialize["sawtoothStatus"] = o.SawtoothStatus
	}
	return toSerialize, nil
}

type NullableStatusResponseV1 struct {
	value *StatusResponseV1
	isSet bool
}

func (v NullableStatusResponseV1) Get() *StatusResponseV1 {
	return v.value
}

func (v *NullableStatusResponseV1) Set(val *StatusResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusResponseV1(val *StatusResponseV1) *NullableStatusResponseV1 {
	return &NullableStatusResponseV1{value: val, isSet: true}
}

func (v NullableStatusResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


