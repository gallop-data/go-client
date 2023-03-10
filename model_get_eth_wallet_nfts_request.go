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

// checks if the GetEthWalletNFTsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEthWalletNFTsRequest{}

// GetEthWalletNFTsRequest struct for GetEthWalletNFTsRequest
type GetEthWalletNFTsRequest struct {
	// The wallet address to search.
	WalletAddress string `json:"wallet_address"`
	// The pagination cursor.
	Page *int32 `json:"page,omitempty"`
	// The number of records returned per page.
	PageSize *int32 `json:"page_size,omitempty"`
}

// NewGetEthWalletNFTsRequest instantiates a new GetEthWalletNFTsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEthWalletNFTsRequest(walletAddress string) *GetEthWalletNFTsRequest {
	this := GetEthWalletNFTsRequest{}
	this.WalletAddress = walletAddress
	return &this
}

// NewGetEthWalletNFTsRequestWithDefaults instantiates a new GetEthWalletNFTsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEthWalletNFTsRequestWithDefaults() *GetEthWalletNFTsRequest {
	this := GetEthWalletNFTsRequest{}
	return &this
}

// GetWalletAddress returns the WalletAddress field value
func (o *GetEthWalletNFTsRequest) GetWalletAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletAddress
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value
// and a boolean to check if the value has been set.
func (o *GetEthWalletNFTsRequest) GetWalletAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletAddress, true
}

// SetWalletAddress sets field value
func (o *GetEthWalletNFTsRequest) SetWalletAddress(v string) {
	o.WalletAddress = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetEthWalletNFTsRequest) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthWalletNFTsRequest) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetEthWalletNFTsRequest) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *GetEthWalletNFTsRequest) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *GetEthWalletNFTsRequest) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthWalletNFTsRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *GetEthWalletNFTsRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *GetEthWalletNFTsRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o GetEthWalletNFTsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEthWalletNFTsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wallet_address"] = o.WalletAddress
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	return toSerialize, nil
}

type NullableGetEthWalletNFTsRequest struct {
	value *GetEthWalletNFTsRequest
	isSet bool
}

func (v NullableGetEthWalletNFTsRequest) Get() *GetEthWalletNFTsRequest {
	return v.value
}

func (v *NullableGetEthWalletNFTsRequest) Set(val *GetEthWalletNFTsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEthWalletNFTsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEthWalletNFTsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEthWalletNFTsRequest(val *GetEthWalletNFTsRequest) *NullableGetEthWalletNFTsRequest {
	return &NullableGetEthWalletNFTsRequest{value: val, isSet: true}
}

func (v NullableGetEthWalletNFTsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEthWalletNFTsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


