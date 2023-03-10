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

// checks if the GetSolWashTransactionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSolWashTransactionsRequest{}

// GetSolWashTransactionsRequest struct for GetSolWashTransactionsRequest
type GetSolWashTransactionsRequest struct {
	// The Gallop slug for the collection. Please see sol/getCollections endpoint.
	CollectionTag string `json:"collection_tag"`
	// An optional list of token addresses.
	MintAddress []string `json:"mint_address,omitempty"`
	// The pagination cursor.
	Page *int32 `json:"page,omitempty"`
	// The number of records returned per page.
	PageSize *int32 `json:"page_size,omitempty"`
}

// NewGetSolWashTransactionsRequest instantiates a new GetSolWashTransactionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolWashTransactionsRequest(collectionTag string) *GetSolWashTransactionsRequest {
	this := GetSolWashTransactionsRequest{}
	this.CollectionTag = collectionTag
	return &this
}

// NewGetSolWashTransactionsRequestWithDefaults instantiates a new GetSolWashTransactionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolWashTransactionsRequestWithDefaults() *GetSolWashTransactionsRequest {
	this := GetSolWashTransactionsRequest{}
	return &this
}

// GetCollectionTag returns the CollectionTag field value
func (o *GetSolWashTransactionsRequest) GetCollectionTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionTag
}

// GetCollectionTagOk returns a tuple with the CollectionTag field value
// and a boolean to check if the value has been set.
func (o *GetSolWashTransactionsRequest) GetCollectionTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionTag, true
}

// SetCollectionTag sets field value
func (o *GetSolWashTransactionsRequest) SetCollectionTag(v string) {
	o.CollectionTag = v
}

// GetMintAddress returns the MintAddress field value if set, zero value otherwise.
func (o *GetSolWashTransactionsRequest) GetMintAddress() []string {
	if o == nil || IsNil(o.MintAddress) {
		var ret []string
		return ret
	}
	return o.MintAddress
}

// GetMintAddressOk returns a tuple with the MintAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolWashTransactionsRequest) GetMintAddressOk() ([]string, bool) {
	if o == nil || IsNil(o.MintAddress) {
		return nil, false
	}
	return o.MintAddress, true
}

// HasMintAddress returns a boolean if a field has been set.
func (o *GetSolWashTransactionsRequest) HasMintAddress() bool {
	if o != nil && !IsNil(o.MintAddress) {
		return true
	}

	return false
}

// SetMintAddress gets a reference to the given []string and assigns it to the MintAddress field.
func (o *GetSolWashTransactionsRequest) SetMintAddress(v []string) {
	o.MintAddress = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetSolWashTransactionsRequest) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolWashTransactionsRequest) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetSolWashTransactionsRequest) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *GetSolWashTransactionsRequest) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *GetSolWashTransactionsRequest) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolWashTransactionsRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *GetSolWashTransactionsRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *GetSolWashTransactionsRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o GetSolWashTransactionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSolWashTransactionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_tag"] = o.CollectionTag
	if !IsNil(o.MintAddress) {
		toSerialize["mint_address"] = o.MintAddress
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	return toSerialize, nil
}

type NullableGetSolWashTransactionsRequest struct {
	value *GetSolWashTransactionsRequest
	isSet bool
}

func (v NullableGetSolWashTransactionsRequest) Get() *GetSolWashTransactionsRequest {
	return v.value
}

func (v *NullableGetSolWashTransactionsRequest) Set(val *GetSolWashTransactionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolWashTransactionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolWashTransactionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolWashTransactionsRequest(val *GetSolWashTransactionsRequest) *NullableGetSolWashTransactionsRequest {
	return &NullableGetSolWashTransactionsRequest{value: val, isSet: true}
}

func (v NullableGetSolWashTransactionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSolWashTransactionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


