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

// checks if the GetSolNFTAccountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSolNFTAccountRequest{}

// GetSolNFTAccountRequest struct for GetSolNFTAccountRequest
type GetSolNFTAccountRequest struct {
	// The Solana token mint address.
	MintAddress string `json:"mint_address"`
}

// NewGetSolNFTAccountRequest instantiates a new GetSolNFTAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolNFTAccountRequest(mintAddress string) *GetSolNFTAccountRequest {
	this := GetSolNFTAccountRequest{}
	this.MintAddress = mintAddress
	return &this
}

// NewGetSolNFTAccountRequestWithDefaults instantiates a new GetSolNFTAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolNFTAccountRequestWithDefaults() *GetSolNFTAccountRequest {
	this := GetSolNFTAccountRequest{}
	return &this
}

// GetMintAddress returns the MintAddress field value
func (o *GetSolNFTAccountRequest) GetMintAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MintAddress
}

// GetMintAddressOk returns a tuple with the MintAddress field value
// and a boolean to check if the value has been set.
func (o *GetSolNFTAccountRequest) GetMintAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MintAddress, true
}

// SetMintAddress sets field value
func (o *GetSolNFTAccountRequest) SetMintAddress(v string) {
	o.MintAddress = v
}

func (o GetSolNFTAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSolNFTAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mint_address"] = o.MintAddress
	return toSerialize, nil
}

type NullableGetSolNFTAccountRequest struct {
	value *GetSolNFTAccountRequest
	isSet bool
}

func (v NullableGetSolNFTAccountRequest) Get() *GetSolNFTAccountRequest {
	return v.value
}

func (v *NullableGetSolNFTAccountRequest) Set(val *GetSolNFTAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolNFTAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolNFTAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolNFTAccountRequest(val *GetSolNFTAccountRequest) *NullableGetSolNFTAccountRequest {
	return &NullableGetSolNFTAccountRequest{value: val, isSet: true}
}

func (v NullableGetSolNFTAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSolNFTAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


