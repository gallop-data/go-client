/*
Gallop API

Data and insights APIs, webooks, and dashboards enabling businesses to launch tokenized products in seconds.

API version: 1.0.0
Contact: support@higallop.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GetEthWalletLabelsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEthWalletLabelsRequest{}

// GetEthWalletLabelsRequest struct for GetEthWalletLabelsRequest
type GetEthWalletLabelsRequest struct {
	// The EVM compatible wallet address
	WalletAddress string `json:"wallet_address"`
}

// NewGetEthWalletLabelsRequest instantiates a new GetEthWalletLabelsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEthWalletLabelsRequest(walletAddress string) *GetEthWalletLabelsRequest {
	this := GetEthWalletLabelsRequest{}
	this.WalletAddress = walletAddress
	return &this
}

// NewGetEthWalletLabelsRequestWithDefaults instantiates a new GetEthWalletLabelsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEthWalletLabelsRequestWithDefaults() *GetEthWalletLabelsRequest {
	this := GetEthWalletLabelsRequest{}
	return &this
}

// GetWalletAddress returns the WalletAddress field value
func (o *GetEthWalletLabelsRequest) GetWalletAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletAddress
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value
// and a boolean to check if the value has been set.
func (o *GetEthWalletLabelsRequest) GetWalletAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletAddress, true
}

// SetWalletAddress sets field value
func (o *GetEthWalletLabelsRequest) SetWalletAddress(v string) {
	o.WalletAddress = v
}

func (o GetEthWalletLabelsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEthWalletLabelsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wallet_address"] = o.WalletAddress
	return toSerialize, nil
}

type NullableGetEthWalletLabelsRequest struct {
	value *GetEthWalletLabelsRequest
	isSet bool
}

func (v NullableGetEthWalletLabelsRequest) Get() *GetEthWalletLabelsRequest {
	return v.value
}

func (v *NullableGetEthWalletLabelsRequest) Set(val *GetEthWalletLabelsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEthWalletLabelsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEthWalletLabelsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEthWalletLabelsRequest(val *GetEthWalletLabelsRequest) *NullableGetEthWalletLabelsRequest {
	return &NullableGetEthWalletLabelsRequest{value: val, isSet: true}
}

func (v NullableGetEthWalletLabelsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEthWalletLabelsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


