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

// checks if the GetPolCollectionTraitsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPolCollectionTraitsRequest{}

// GetPolCollectionTraitsRequest struct for GetPolCollectionTraitsRequest
type GetPolCollectionTraitsRequest struct {
	// The contract address of the collection.
	CollectionAddress *string `json:"collection_address,omitempty"`
}

// NewGetPolCollectionTraitsRequest instantiates a new GetPolCollectionTraitsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPolCollectionTraitsRequest() *GetPolCollectionTraitsRequest {
	this := GetPolCollectionTraitsRequest{}
	return &this
}

// NewGetPolCollectionTraitsRequestWithDefaults instantiates a new GetPolCollectionTraitsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPolCollectionTraitsRequestWithDefaults() *GetPolCollectionTraitsRequest {
	this := GetPolCollectionTraitsRequest{}
	return &this
}

// GetCollectionAddress returns the CollectionAddress field value if set, zero value otherwise.
func (o *GetPolCollectionTraitsRequest) GetCollectionAddress() string {
	if o == nil || isNil(o.CollectionAddress) {
		var ret string
		return ret
	}
	return *o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionTraitsRequest) GetCollectionAddressOk() (*string, bool) {
	if o == nil || isNil(o.CollectionAddress) {
		return nil, false
	}
	return o.CollectionAddress, true
}

// HasCollectionAddress returns a boolean if a field has been set.
func (o *GetPolCollectionTraitsRequest) HasCollectionAddress() bool {
	if o != nil && !isNil(o.CollectionAddress) {
		return true
	}

	return false
}

// SetCollectionAddress gets a reference to the given string and assigns it to the CollectionAddress field.
func (o *GetPolCollectionTraitsRequest) SetCollectionAddress(v string) {
	o.CollectionAddress = &v
}

func (o GetPolCollectionTraitsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPolCollectionTraitsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CollectionAddress) {
		toSerialize["collection_address"] = o.CollectionAddress
	}
	return toSerialize, nil
}

type NullableGetPolCollectionTraitsRequest struct {
	value *GetPolCollectionTraitsRequest
	isSet bool
}

func (v NullableGetPolCollectionTraitsRequest) Get() *GetPolCollectionTraitsRequest {
	return v.value
}

func (v *NullableGetPolCollectionTraitsRequest) Set(val *GetPolCollectionTraitsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPolCollectionTraitsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPolCollectionTraitsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPolCollectionTraitsRequest(val *GetPolCollectionTraitsRequest) *NullableGetPolCollectionTraitsRequest {
	return &NullableGetPolCollectionTraitsRequest{value: val, isSet: true}
}

func (v NullableGetPolCollectionTraitsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPolCollectionTraitsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

