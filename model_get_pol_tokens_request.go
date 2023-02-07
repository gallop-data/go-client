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

// checks if the GetPolTokensRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPolTokensRequest{}

// GetPolTokensRequest struct for GetPolTokensRequest
type GetPolTokensRequest struct {
	// The contract address of the token collection.
	CollectionAddress string `json:"collection_address"`
	// A list of token ids.
	TokenId []string `json:"token_id,omitempty"`
	// The pagination cursor.
	Page *int32 `json:"page,omitempty"`
	// The number of records returned per page.
	PageSize *int32 `json:"page_size,omitempty"`
}

// NewGetPolTokensRequest instantiates a new GetPolTokensRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPolTokensRequest(collectionAddress string) *GetPolTokensRequest {
	this := GetPolTokensRequest{}
	this.CollectionAddress = collectionAddress
	return &this
}

// NewGetPolTokensRequestWithDefaults instantiates a new GetPolTokensRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPolTokensRequestWithDefaults() *GetPolTokensRequest {
	this := GetPolTokensRequest{}
	return &this
}

// GetCollectionAddress returns the CollectionAddress field value
func (o *GetPolTokensRequest) GetCollectionAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value
// and a boolean to check if the value has been set.
func (o *GetPolTokensRequest) GetCollectionAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionAddress, true
}

// SetCollectionAddress sets field value
func (o *GetPolTokensRequest) SetCollectionAddress(v string) {
	o.CollectionAddress = v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *GetPolTokensRequest) GetTokenId() []string {
	if o == nil || isNil(o.TokenId) {
		var ret []string
		return ret
	}
	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolTokensRequest) GetTokenIdOk() ([]string, bool) {
	if o == nil || isNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *GetPolTokensRequest) HasTokenId() bool {
	if o != nil && !isNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given []string and assigns it to the TokenId field.
func (o *GetPolTokensRequest) SetTokenId(v []string) {
	o.TokenId = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetPolTokensRequest) GetPage() int32 {
	if o == nil || isNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolTokensRequest) GetPageOk() (*int32, bool) {
	if o == nil || isNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetPolTokensRequest) HasPage() bool {
	if o != nil && !isNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *GetPolTokensRequest) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *GetPolTokensRequest) GetPageSize() int32 {
	if o == nil || isNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolTokensRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || isNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *GetPolTokensRequest) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *GetPolTokensRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o GetPolTokensRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPolTokensRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_address"] = o.CollectionAddress
	if !isNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	if !isNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !isNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	return toSerialize, nil
}

type NullableGetPolTokensRequest struct {
	value *GetPolTokensRequest
	isSet bool
}

func (v NullableGetPolTokensRequest) Get() *GetPolTokensRequest {
	return v.value
}

func (v *NullableGetPolTokensRequest) Set(val *GetPolTokensRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPolTokensRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPolTokensRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPolTokensRequest(val *GetPolTokensRequest) *NullableGetPolTokensRequest {
	return &NullableGetPolTokensRequest{value: val, isSet: true}
}

func (v NullableGetPolTokensRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPolTokensRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

