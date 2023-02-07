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

// checks if the GetSolAccountNFTsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSolAccountNFTsRequest{}

// GetSolAccountNFTsRequest struct for GetSolAccountNFTsRequest
type GetSolAccountNFTsRequest struct {
	// The account address to query.
	AccountAddress string `json:"account_address"`
}

// NewGetSolAccountNFTsRequest instantiates a new GetSolAccountNFTsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolAccountNFTsRequest(accountAddress string) *GetSolAccountNFTsRequest {
	this := GetSolAccountNFTsRequest{}
	this.AccountAddress = accountAddress
	return &this
}

// NewGetSolAccountNFTsRequestWithDefaults instantiates a new GetSolAccountNFTsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolAccountNFTsRequestWithDefaults() *GetSolAccountNFTsRequest {
	this := GetSolAccountNFTsRequest{}
	return &this
}

// GetAccountAddress returns the AccountAddress field value
func (o *GetSolAccountNFTsRequest) GetAccountAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountAddress
}

// GetAccountAddressOk returns a tuple with the AccountAddress field value
// and a boolean to check if the value has been set.
func (o *GetSolAccountNFTsRequest) GetAccountAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountAddress, true
}

// SetAccountAddress sets field value
func (o *GetSolAccountNFTsRequest) SetAccountAddress(v string) {
	o.AccountAddress = v
}

func (o GetSolAccountNFTsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSolAccountNFTsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_address"] = o.AccountAddress
	return toSerialize, nil
}

type NullableGetSolAccountNFTsRequest struct {
	value *GetSolAccountNFTsRequest
	isSet bool
}

func (v NullableGetSolAccountNFTsRequest) Get() *GetSolAccountNFTsRequest {
	return v.value
}

func (v *NullableGetSolAccountNFTsRequest) Set(val *GetSolAccountNFTsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolAccountNFTsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolAccountNFTsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolAccountNFTsRequest(val *GetSolAccountNFTsRequest) *NullableGetSolAccountNFTsRequest {
	return &NullableGetSolAccountNFTsRequest{value: val, isSet: true}
}

func (v NullableGetSolAccountNFTsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSolAccountNFTsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

