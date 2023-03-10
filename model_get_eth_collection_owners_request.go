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

// checks if the GetEthCollectionOwnersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEthCollectionOwnersRequest{}

// GetEthCollectionOwnersRequest struct for GetEthCollectionOwnersRequest
type GetEthCollectionOwnersRequest struct {
	// The contract address of the collection.
	CollectionAddress string `json:"collection_address"`
	// The pagination cursor.
	Page *int32 `json:"page,omitempty"`
	// The number of records returned per page.
	PageSize *int32 `json:"page_size,omitempty"`
}

// NewGetEthCollectionOwnersRequest instantiates a new GetEthCollectionOwnersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEthCollectionOwnersRequest(collectionAddress string) *GetEthCollectionOwnersRequest {
	this := GetEthCollectionOwnersRequest{}
	this.CollectionAddress = collectionAddress
	return &this
}

// NewGetEthCollectionOwnersRequestWithDefaults instantiates a new GetEthCollectionOwnersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEthCollectionOwnersRequestWithDefaults() *GetEthCollectionOwnersRequest {
	this := GetEthCollectionOwnersRequest{}
	return &this
}

// GetCollectionAddress returns the CollectionAddress field value
func (o *GetEthCollectionOwnersRequest) GetCollectionAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value
// and a boolean to check if the value has been set.
func (o *GetEthCollectionOwnersRequest) GetCollectionAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionAddress, true
}

// SetCollectionAddress sets field value
func (o *GetEthCollectionOwnersRequest) SetCollectionAddress(v string) {
	o.CollectionAddress = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetEthCollectionOwnersRequest) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionOwnersRequest) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetEthCollectionOwnersRequest) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *GetEthCollectionOwnersRequest) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *GetEthCollectionOwnersRequest) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionOwnersRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *GetEthCollectionOwnersRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *GetEthCollectionOwnersRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o GetEthCollectionOwnersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEthCollectionOwnersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_address"] = o.CollectionAddress
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	return toSerialize, nil
}

type NullableGetEthCollectionOwnersRequest struct {
	value *GetEthCollectionOwnersRequest
	isSet bool
}

func (v NullableGetEthCollectionOwnersRequest) Get() *GetEthCollectionOwnersRequest {
	return v.value
}

func (v *NullableGetEthCollectionOwnersRequest) Set(val *GetEthCollectionOwnersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEthCollectionOwnersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEthCollectionOwnersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEthCollectionOwnersRequest(val *GetEthCollectionOwnersRequest) *NullableGetEthCollectionOwnersRequest {
	return &NullableGetEthCollectionOwnersRequest{value: val, isSet: true}
}

func (v NullableGetEthCollectionOwnersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEthCollectionOwnersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


