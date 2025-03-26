/*
Hyperledger Cactus Plugin - HTLC-ETH Besu

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-htlc-eth-besu

import (
	"encoding/json"
)

// checks if the Web3TransactionReceipt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3TransactionReceipt{}

// Web3TransactionReceipt struct for Web3TransactionReceipt
type Web3TransactionReceipt struct {
	Status bool `json:"status"`
	TransactionHash string `json:"transactionHash"`
	TransactionIndex float32 `json:"transactionIndex"`
	BlockHash string `json:"blockHash"`
	BlockNumber float32 `json:"blockNumber"`
	GasUsed float32 `json:"gasUsed"`
	ContractAddress NullableString `json:"contractAddress,omitempty"`
	From string `json:"from"`
	To string `json:"to"`
	AdditionalProperties map[string]interface{}
}

type _Web3TransactionReceipt Web3TransactionReceipt

// NewWeb3TransactionReceipt instantiates a new Web3TransactionReceipt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3TransactionReceipt(status bool, transactionHash string, transactionIndex float32, blockHash string, blockNumber float32, gasUsed float32, from string, to string) *Web3TransactionReceipt {
	this := Web3TransactionReceipt{}
	this.Status = status
	this.TransactionHash = transactionHash
	this.TransactionIndex = transactionIndex
	this.BlockHash = blockHash
	this.BlockNumber = blockNumber
	this.GasUsed = gasUsed
	this.From = from
	this.To = to
	return &this
}

// NewWeb3TransactionReceiptWithDefaults instantiates a new Web3TransactionReceipt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3TransactionReceiptWithDefaults() *Web3TransactionReceipt {
	this := Web3TransactionReceipt{}
	return &this
}

// GetStatus returns the Status field value
func (o *Web3TransactionReceipt) GetStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Web3TransactionReceipt) SetStatus(v bool) {
	o.Status = v
}

// GetTransactionHash returns the TransactionHash field value
func (o *Web3TransactionReceipt) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *Web3TransactionReceipt) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetTransactionIndex returns the TransactionIndex field value
func (o *Web3TransactionReceipt) GetTransactionIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TransactionIndex
}

// GetTransactionIndexOk returns a tuple with the TransactionIndex field value
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetTransactionIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionIndex, true
}

// SetTransactionIndex sets field value
func (o *Web3TransactionReceipt) SetTransactionIndex(v float32) {
	o.TransactionIndex = v
}

// GetBlockHash returns the BlockHash field value
func (o *Web3TransactionReceipt) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *Web3TransactionReceipt) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetBlockNumber returns the BlockNumber field value
func (o *Web3TransactionReceipt) GetBlockNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BlockNumber
}

// GetBlockNumberOk returns a tuple with the BlockNumber field value
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetBlockNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockNumber, true
}

// SetBlockNumber sets field value
func (o *Web3TransactionReceipt) SetBlockNumber(v float32) {
	o.BlockNumber = v
}

// GetGasUsed returns the GasUsed field value
func (o *Web3TransactionReceipt) GetGasUsed() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetGasUsedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *Web3TransactionReceipt) SetGasUsed(v float32) {
	o.GasUsed = v
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TransactionReceipt) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress.Get()) {
		var ret string
		return ret
	}
	return *o.ContractAddress.Get()
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TransactionReceipt) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractAddress.Get(), o.ContractAddress.IsSet()
}

// HasContractAddress returns a boolean if a field has been set.
func (o *Web3TransactionReceipt) HasContractAddress() bool {
	if o != nil && o.ContractAddress.IsSet() {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given NullableString and assigns it to the ContractAddress field.
func (o *Web3TransactionReceipt) SetContractAddress(v string) {
	o.ContractAddress.Set(&v)
}
// SetContractAddressNil sets the value for ContractAddress to be an explicit nil
func (o *Web3TransactionReceipt) SetContractAddressNil() {
	o.ContractAddress.Set(nil)
}

// UnsetContractAddress ensures that no value is present for ContractAddress, not even an explicit nil
func (o *Web3TransactionReceipt) UnsetContractAddress() {
	o.ContractAddress.Unset()
}

// GetFrom returns the From field value
func (o *Web3TransactionReceipt) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *Web3TransactionReceipt) SetFrom(v string) {
	o.From = v
}

// GetTo returns the To field value
func (o *Web3TransactionReceipt) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *Web3TransactionReceipt) SetTo(v string) {
	o.To = v
}

func (o Web3TransactionReceipt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3TransactionReceipt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["transactionHash"] = o.TransactionHash
	toSerialize["transactionIndex"] = o.TransactionIndex
	toSerialize["blockHash"] = o.BlockHash
	toSerialize["blockNumber"] = o.BlockNumber
	toSerialize["gasUsed"] = o.GasUsed
	if o.ContractAddress.IsSet() {
		toSerialize["contractAddress"] = o.ContractAddress.Get()
	}
	toSerialize["from"] = o.From
	toSerialize["to"] = o.To

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3TransactionReceipt) UnmarshalJSON(bytes []byte) (err error) {
	varWeb3TransactionReceipt := _Web3TransactionReceipt{}

	if err = json.Unmarshal(bytes, &varWeb3TransactionReceipt); err == nil {
		*o = Web3TransactionReceipt(varWeb3TransactionReceipt)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "transactionHash")
		delete(additionalProperties, "transactionIndex")
		delete(additionalProperties, "blockHash")
		delete(additionalProperties, "blockNumber")
		delete(additionalProperties, "gasUsed")
		delete(additionalProperties, "contractAddress")
		delete(additionalProperties, "from")
		delete(additionalProperties, "to")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3TransactionReceipt struct {
	value *Web3TransactionReceipt
	isSet bool
}

func (v NullableWeb3TransactionReceipt) Get() *Web3TransactionReceipt {
	return v.value
}

func (v *NullableWeb3TransactionReceipt) Set(val *Web3TransactionReceipt) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3TransactionReceipt) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3TransactionReceipt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3TransactionReceipt(val *Web3TransactionReceipt) *NullableWeb3TransactionReceipt {
	return &NullableWeb3TransactionReceipt{value: val, isSet: true}
}

func (v NullableWeb3TransactionReceipt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3TransactionReceipt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


