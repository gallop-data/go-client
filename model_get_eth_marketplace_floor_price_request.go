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

// checks if the GetEthMarketplaceFloorPriceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEthMarketplaceFloorPriceRequest{}

// GetEthMarketplaceFloorPriceRequest struct for GetEthMarketplaceFloorPriceRequest
type GetEthMarketplaceFloorPriceRequest struct {
	// The pagination cursor.
	Page *int32 `json:"page,omitempty"`
	// The number of records returned per page.
	PageSize *int32 `json:"page_size,omitempty"`
	// Array of collection addresses
	CollectionAddress []string `json:"collection_address,omitempty"`
}

// NewGetEthMarketplaceFloorPriceRequest instantiates a new GetEthMarketplaceFloorPriceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEthMarketplaceFloorPriceRequest() *GetEthMarketplaceFloorPriceRequest {
	this := GetEthMarketplaceFloorPriceRequest{}
	return &this
}

// NewGetEthMarketplaceFloorPriceRequestWithDefaults instantiates a new GetEthMarketplaceFloorPriceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEthMarketplaceFloorPriceRequestWithDefaults() *GetEthMarketplaceFloorPriceRequest {
	this := GetEthMarketplaceFloorPriceRequest{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetEthMarketplaceFloorPriceRequest) GetPage() int32 {
	if o == nil || isNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthMarketplaceFloorPriceRequest) GetPageOk() (*int32, bool) {
	if o == nil || isNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetEthMarketplaceFloorPriceRequest) HasPage() bool {
	if o != nil && !isNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *GetEthMarketplaceFloorPriceRequest) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *GetEthMarketplaceFloorPriceRequest) GetPageSize() int32 {
	if o == nil || isNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthMarketplaceFloorPriceRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || isNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *GetEthMarketplaceFloorPriceRequest) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *GetEthMarketplaceFloorPriceRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetCollectionAddress returns the CollectionAddress field value if set, zero value otherwise.
func (o *GetEthMarketplaceFloorPriceRequest) GetCollectionAddress() []string {
	if o == nil || isNil(o.CollectionAddress) {
		var ret []string
		return ret
	}
	return o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthMarketplaceFloorPriceRequest) GetCollectionAddressOk() ([]string, bool) {
	if o == nil || isNil(o.CollectionAddress) {
		return nil, false
	}
	return o.CollectionAddress, true
}

// HasCollectionAddress returns a boolean if a field has been set.
func (o *GetEthMarketplaceFloorPriceRequest) HasCollectionAddress() bool {
	if o != nil && !isNil(o.CollectionAddress) {
		return true
	}

	return false
}

// SetCollectionAddress gets a reference to the given []string and assigns it to the CollectionAddress field.
func (o *GetEthMarketplaceFloorPriceRequest) SetCollectionAddress(v []string) {
	o.CollectionAddress = v
}

func (o GetEthMarketplaceFloorPriceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEthMarketplaceFloorPriceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !isNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !isNil(o.CollectionAddress) {
		toSerialize["collection_address"] = o.CollectionAddress
	}
	return toSerialize, nil
}

type NullableGetEthMarketplaceFloorPriceRequest struct {
	value *GetEthMarketplaceFloorPriceRequest
	isSet bool
}

func (v NullableGetEthMarketplaceFloorPriceRequest) Get() *GetEthMarketplaceFloorPriceRequest {
	return v.value
}

func (v *NullableGetEthMarketplaceFloorPriceRequest) Set(val *GetEthMarketplaceFloorPriceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEthMarketplaceFloorPriceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEthMarketplaceFloorPriceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEthMarketplaceFloorPriceRequest(val *GetEthMarketplaceFloorPriceRequest) *NullableGetEthMarketplaceFloorPriceRequest {
	return &NullableGetEthMarketplaceFloorPriceRequest{value: val, isSet: true}
}

func (v NullableGetEthMarketplaceFloorPriceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEthMarketplaceFloorPriceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

