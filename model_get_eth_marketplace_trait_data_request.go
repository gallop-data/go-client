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

// checks if the GetEthMarketplaceTraitDataRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEthMarketplaceTraitDataRequest{}

// GetEthMarketplaceTraitDataRequest struct for GetEthMarketplaceTraitDataRequest
type GetEthMarketplaceTraitDataRequest struct {
	// The contract address of the collection.
	CollectionAddress string `json:"collection_address"`
}

// NewGetEthMarketplaceTraitDataRequest instantiates a new GetEthMarketplaceTraitDataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEthMarketplaceTraitDataRequest(collectionAddress string) *GetEthMarketplaceTraitDataRequest {
	this := GetEthMarketplaceTraitDataRequest{}
	this.CollectionAddress = collectionAddress
	return &this
}

// NewGetEthMarketplaceTraitDataRequestWithDefaults instantiates a new GetEthMarketplaceTraitDataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEthMarketplaceTraitDataRequestWithDefaults() *GetEthMarketplaceTraitDataRequest {
	this := GetEthMarketplaceTraitDataRequest{}
	return &this
}

// GetCollectionAddress returns the CollectionAddress field value
func (o *GetEthMarketplaceTraitDataRequest) GetCollectionAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value
// and a boolean to check if the value has been set.
func (o *GetEthMarketplaceTraitDataRequest) GetCollectionAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionAddress, true
}

// SetCollectionAddress sets field value
func (o *GetEthMarketplaceTraitDataRequest) SetCollectionAddress(v string) {
	o.CollectionAddress = v
}

func (o GetEthMarketplaceTraitDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEthMarketplaceTraitDataRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_address"] = o.CollectionAddress
	return toSerialize, nil
}

type NullableGetEthMarketplaceTraitDataRequest struct {
	value *GetEthMarketplaceTraitDataRequest
	isSet bool
}

func (v NullableGetEthMarketplaceTraitDataRequest) Get() *GetEthMarketplaceTraitDataRequest {
	return v.value
}

func (v *NullableGetEthMarketplaceTraitDataRequest) Set(val *GetEthMarketplaceTraitDataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEthMarketplaceTraitDataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEthMarketplaceTraitDataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEthMarketplaceTraitDataRequest(val *GetEthMarketplaceTraitDataRequest) *NullableGetEthMarketplaceTraitDataRequest {
	return &NullableGetEthMarketplaceTraitDataRequest{value: val, isSet: true}
}

func (v NullableGetEthMarketplaceTraitDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEthMarketplaceTraitDataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


