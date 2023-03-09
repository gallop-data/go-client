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

// checks if the GetEthLiveListingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEthLiveListingsRequest{}

// GetEthLiveListingsRequest struct for GetEthLiveListingsRequest
type GetEthLiveListingsRequest struct {
	// The contract address of a collection.
	CollectionAddress string `json:"collection_address"`
}

// NewGetEthLiveListingsRequest instantiates a new GetEthLiveListingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEthLiveListingsRequest(collectionAddress string) *GetEthLiveListingsRequest {
	this := GetEthLiveListingsRequest{}
	this.CollectionAddress = collectionAddress
	return &this
}

// NewGetEthLiveListingsRequestWithDefaults instantiates a new GetEthLiveListingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEthLiveListingsRequestWithDefaults() *GetEthLiveListingsRequest {
	this := GetEthLiveListingsRequest{}
	return &this
}

// GetCollectionAddress returns the CollectionAddress field value
func (o *GetEthLiveListingsRequest) GetCollectionAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value
// and a boolean to check if the value has been set.
func (o *GetEthLiveListingsRequest) GetCollectionAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionAddress, true
}

// SetCollectionAddress sets field value
func (o *GetEthLiveListingsRequest) SetCollectionAddress(v string) {
	o.CollectionAddress = v
}

func (o GetEthLiveListingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEthLiveListingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_address"] = o.CollectionAddress
	return toSerialize, nil
}

type NullableGetEthLiveListingsRequest struct {
	value *GetEthLiveListingsRequest
	isSet bool
}

func (v NullableGetEthLiveListingsRequest) Get() *GetEthLiveListingsRequest {
	return v.value
}

func (v *NullableGetEthLiveListingsRequest) Set(val *GetEthLiveListingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEthLiveListingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEthLiveListingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEthLiveListingsRequest(val *GetEthLiveListingsRequest) *NullableGetEthLiveListingsRequest {
	return &NullableGetEthLiveListingsRequest{value: val, isSet: true}
}

func (v NullableGetEthLiveListingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEthLiveListingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


