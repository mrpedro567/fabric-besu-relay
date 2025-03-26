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

// checks if the RunTransactionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunTransactionRequest{}

// RunTransactionRequest struct for RunTransactionRequest
type RunTransactionRequest struct {
	// An array of endorsing peers (name or url) for the transaction.
	EndorsingPeers []string `json:"endorsingPeers,omitempty"`
	// An array of endorsing organizations (by mspID or issuer org name on certificate) for the transaction.
	EndorsingOrgs []string `json:"endorsingOrgs,omitempty"`
	TransientData map[string]interface{} `json:"transientData,omitempty"`
	GatewayOptions *GatewayOptions `json:"gatewayOptions,omitempty"`
	SigningCredential FabricSigningCredential `json:"signingCredential"`
	ChannelName string `json:"channelName"`
	ContractName string `json:"contractName"`
	InvocationType FabricContractInvocationType `json:"invocationType"`
	MethodName string `json:"methodName"`
	Params []*string `json:"params"`
	ResponseType *RunTransactionResponseType `json:"responseType,omitempty"`
}

// NewRunTransactionRequest instantiates a new RunTransactionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunTransactionRequest(signingCredential FabricSigningCredential, channelName string, contractName string, invocationType FabricContractInvocationType, methodName string, params []*string) *RunTransactionRequest {
	this := RunTransactionRequest{}
	this.SigningCredential = signingCredential
	this.ChannelName = channelName
	this.ContractName = contractName
	this.InvocationType = invocationType
	this.MethodName = methodName
	this.Params = params
	return &this
}

// NewRunTransactionRequestWithDefaults instantiates a new RunTransactionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunTransactionRequestWithDefaults() *RunTransactionRequest {
	this := RunTransactionRequest{}
	return &this
}

// GetEndorsingPeers returns the EndorsingPeers field value if set, zero value otherwise.
func (o *RunTransactionRequest) GetEndorsingPeers() []string {
	if o == nil || IsNil(o.EndorsingPeers) {
		var ret []string
		return ret
	}
	return o.EndorsingPeers
}

// GetEndorsingPeersOk returns a tuple with the EndorsingPeers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunTransactionRequest) GetEndorsingPeersOk() ([]string, bool) {
	if o == nil || IsNil(o.EndorsingPeers) {
		return nil, false
	}
	return o.EndorsingPeers, true
}

// HasEndorsingPeers returns a boolean if a field has been set.
func (o *RunTransactionRequest) HasEndorsingPeers() bool {
	if o != nil && !IsNil(o.EndorsingPeers) {
		return true
	}

	return false
}

// SetEndorsingPeers gets a reference to the given []string and assigns it to the EndorsingPeers field.
func (o *RunTransactionRequest) SetEndorsingPeers(v []string) {
	o.EndorsingPeers = v
}

// GetEndorsingOrgs returns the EndorsingOrgs field value if set, zero value otherwise.
func (o *RunTransactionRequest) GetEndorsingOrgs() []string {
	if o == nil || IsNil(o.EndorsingOrgs) {
		var ret []string
		return ret
	}
	return o.EndorsingOrgs
}

// GetEndorsingOrgsOk returns a tuple with the EndorsingOrgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunTransactionRequest) GetEndorsingOrgsOk() ([]string, bool) {
	if o == nil || IsNil(o.EndorsingOrgs) {
		return nil, false
	}
	return o.EndorsingOrgs, true
}

// HasEndorsingOrgs returns a boolean if a field has been set.
func (o *RunTransactionRequest) HasEndorsingOrgs() bool {
	if o != nil && !IsNil(o.EndorsingOrgs) {
		return true
	}

	return false
}

// SetEndorsingOrgs gets a reference to the given []string and assigns it to the EndorsingOrgs field.
func (o *RunTransactionRequest) SetEndorsingOrgs(v []string) {
	o.EndorsingOrgs = v
}

// GetTransientData returns the TransientData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunTransactionRequest) GetTransientData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.TransientData
}

// GetTransientDataOk returns a tuple with the TransientData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunTransactionRequest) GetTransientDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TransientData) {
		return map[string]interface{}{}, false
	}
	return o.TransientData, true
}

// HasTransientData returns a boolean if a field has been set.
func (o *RunTransactionRequest) HasTransientData() bool {
	if o != nil && IsNil(o.TransientData) {
		return true
	}

	return false
}

// SetTransientData gets a reference to the given map[string]interface{} and assigns it to the TransientData field.
func (o *RunTransactionRequest) SetTransientData(v map[string]interface{}) {
	o.TransientData = v
}

// GetGatewayOptions returns the GatewayOptions field value if set, zero value otherwise.
func (o *RunTransactionRequest) GetGatewayOptions() GatewayOptions {
	if o == nil || IsNil(o.GatewayOptions) {
		var ret GatewayOptions
		return ret
	}
	return *o.GatewayOptions
}

// GetGatewayOptionsOk returns a tuple with the GatewayOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunTransactionRequest) GetGatewayOptionsOk() (*GatewayOptions, bool) {
	if o == nil || IsNil(o.GatewayOptions) {
		return nil, false
	}
	return o.GatewayOptions, true
}

// HasGatewayOptions returns a boolean if a field has been set.
func (o *RunTransactionRequest) HasGatewayOptions() bool {
	if o != nil && !IsNil(o.GatewayOptions) {
		return true
	}

	return false
}

// SetGatewayOptions gets a reference to the given GatewayOptions and assigns it to the GatewayOptions field.
func (o *RunTransactionRequest) SetGatewayOptions(v GatewayOptions) {
	o.GatewayOptions = &v
}

// GetSigningCredential returns the SigningCredential field value
func (o *RunTransactionRequest) GetSigningCredential() FabricSigningCredential {
	if o == nil {
		var ret FabricSigningCredential
		return ret
	}

	return o.SigningCredential
}

// GetSigningCredentialOk returns a tuple with the SigningCredential field value
// and a boolean to check if the value has been set.
func (o *RunTransactionRequest) GetSigningCredentialOk() (*FabricSigningCredential, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SigningCredential, true
}

// SetSigningCredential sets field value
func (o *RunTransactionRequest) SetSigningCredential(v FabricSigningCredential) {
	o.SigningCredential = v
}

// GetChannelName returns the ChannelName field value
func (o *RunTransactionRequest) GetChannelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value
// and a boolean to check if the value has been set.
func (o *RunTransactionRequest) GetChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelName, true
}

// SetChannelName sets field value
func (o *RunTransactionRequest) SetChannelName(v string) {
	o.ChannelName = v
}

// GetContractName returns the ContractName field value
func (o *RunTransactionRequest) GetContractName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value
// and a boolean to check if the value has been set.
func (o *RunTransactionRequest) GetContractNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractName, true
}

// SetContractName sets field value
func (o *RunTransactionRequest) SetContractName(v string) {
	o.ContractName = v
}

// GetInvocationType returns the InvocationType field value
func (o *RunTransactionRequest) GetInvocationType() FabricContractInvocationType {
	if o == nil {
		var ret FabricContractInvocationType
		return ret
	}

	return o.InvocationType
}

// GetInvocationTypeOk returns a tuple with the InvocationType field value
// and a boolean to check if the value has been set.
func (o *RunTransactionRequest) GetInvocationTypeOk() (*FabricContractInvocationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvocationType, true
}

// SetInvocationType sets field value
func (o *RunTransactionRequest) SetInvocationType(v FabricContractInvocationType) {
	o.InvocationType = v
}

// GetMethodName returns the MethodName field value
func (o *RunTransactionRequest) GetMethodName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MethodName
}

// GetMethodNameOk returns a tuple with the MethodName field value
// and a boolean to check if the value has been set.
func (o *RunTransactionRequest) GetMethodNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MethodName, true
}

// SetMethodName sets field value
func (o *RunTransactionRequest) SetMethodName(v string) {
	o.MethodName = v
}

// GetParams returns the Params field value
func (o *RunTransactionRequest) GetParams() []*string {
	if o == nil {
		var ret []*string
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *RunTransactionRequest) GetParamsOk() ([]*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *RunTransactionRequest) SetParams(v []*string) {
	o.Params = v
}

// GetResponseType returns the ResponseType field value if set, zero value otherwise.
func (o *RunTransactionRequest) GetResponseType() RunTransactionResponseType {
	if o == nil || IsNil(o.ResponseType) {
		var ret RunTransactionResponseType
		return ret
	}
	return *o.ResponseType
}

// GetResponseTypeOk returns a tuple with the ResponseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunTransactionRequest) GetResponseTypeOk() (*RunTransactionResponseType, bool) {
	if o == nil || IsNil(o.ResponseType) {
		return nil, false
	}
	return o.ResponseType, true
}

// HasResponseType returns a boolean if a field has been set.
func (o *RunTransactionRequest) HasResponseType() bool {
	if o != nil && !IsNil(o.ResponseType) {
		return true
	}

	return false
}

// SetResponseType gets a reference to the given RunTransactionResponseType and assigns it to the ResponseType field.
func (o *RunTransactionRequest) SetResponseType(v RunTransactionResponseType) {
	o.ResponseType = &v
}

func (o RunTransactionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunTransactionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndorsingPeers) {
		toSerialize["endorsingPeers"] = o.EndorsingPeers
	}
	if !IsNil(o.EndorsingOrgs) {
		toSerialize["endorsingOrgs"] = o.EndorsingOrgs
	}
	if o.TransientData != nil {
		toSerialize["transientData"] = o.TransientData
	}
	if !IsNil(o.GatewayOptions) {
		toSerialize["gatewayOptions"] = o.GatewayOptions
	}
	toSerialize["signingCredential"] = o.SigningCredential
	toSerialize["channelName"] = o.ChannelName
	toSerialize["contractName"] = o.ContractName
	toSerialize["invocationType"] = o.InvocationType
	toSerialize["methodName"] = o.MethodName
	toSerialize["params"] = o.Params
	if !IsNil(o.ResponseType) {
		toSerialize["responseType"] = o.ResponseType
	}
	return toSerialize, nil
}

type NullableRunTransactionRequest struct {
	value *RunTransactionRequest
	isSet bool
}

func (v NullableRunTransactionRequest) Get() *RunTransactionRequest {
	return v.value
}

func (v *NullableRunTransactionRequest) Set(val *RunTransactionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRunTransactionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRunTransactionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunTransactionRequest(val *RunTransactionRequest) *NullableRunTransactionRequest {
	return &NullableRunTransactionRequest{value: val, isSet: true}
}

func (v NullableRunTransactionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunTransactionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


