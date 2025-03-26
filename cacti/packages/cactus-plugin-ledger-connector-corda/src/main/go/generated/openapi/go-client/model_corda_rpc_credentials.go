/*
Hyperledger Cacti Plugin - Connector Corda

Can perform basic tasks on a Corda ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-corda

import (
	"encoding/json"
)

// checks if the CordaRpcCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CordaRpcCredentials{}

// CordaRpcCredentials struct for CordaRpcCredentials
type CordaRpcCredentials struct {
	Hostname string `json:"hostname"`
	Port int32 `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewCordaRpcCredentials instantiates a new CordaRpcCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCordaRpcCredentials(hostname string, port int32, username string, password string) *CordaRpcCredentials {
	this := CordaRpcCredentials{}
	this.Hostname = hostname
	this.Port = port
	this.Username = username
	this.Password = password
	return &this
}

// NewCordaRpcCredentialsWithDefaults instantiates a new CordaRpcCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCordaRpcCredentialsWithDefaults() *CordaRpcCredentials {
	this := CordaRpcCredentials{}
	return &this
}

// GetHostname returns the Hostname field value
func (o *CordaRpcCredentials) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *CordaRpcCredentials) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *CordaRpcCredentials) SetHostname(v string) {
	o.Hostname = v
}

// GetPort returns the Port field value
func (o *CordaRpcCredentials) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *CordaRpcCredentials) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *CordaRpcCredentials) SetPort(v int32) {
	o.Port = v
}

// GetUsername returns the Username field value
func (o *CordaRpcCredentials) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CordaRpcCredentials) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CordaRpcCredentials) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *CordaRpcCredentials) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CordaRpcCredentials) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CordaRpcCredentials) SetPassword(v string) {
	o.Password = v
}

func (o CordaRpcCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CordaRpcCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hostname"] = o.Hostname
	toSerialize["port"] = o.Port
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

type NullableCordaRpcCredentials struct {
	value *CordaRpcCredentials
	isSet bool
}

func (v NullableCordaRpcCredentials) Get() *CordaRpcCredentials {
	return v.value
}

func (v *NullableCordaRpcCredentials) Set(val *CordaRpcCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableCordaRpcCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableCordaRpcCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCordaRpcCredentials(val *CordaRpcCredentials) *NullableCordaRpcCredentials {
	return &NullableCordaRpcCredentials{value: val, isSet: true}
}

func (v NullableCordaRpcCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCordaRpcCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


