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

// checks if the GetSolMarketplaceTraitDataRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSolMarketplaceTraitDataRequest{}

// GetSolMarketplaceTraitDataRequest struct for GetSolMarketplaceTraitDataRequest
type GetSolMarketplaceTraitDataRequest struct {
	// Collection_tag
	CollectionTag string `json:"collection_tag"`
}

// NewGetSolMarketplaceTraitDataRequest instantiates a new GetSolMarketplaceTraitDataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolMarketplaceTraitDataRequest(collectionTag string) *GetSolMarketplaceTraitDataRequest {
	this := GetSolMarketplaceTraitDataRequest{}
	this.CollectionTag = collectionTag
	return &this
}

// NewGetSolMarketplaceTraitDataRequestWithDefaults instantiates a new GetSolMarketplaceTraitDataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolMarketplaceTraitDataRequestWithDefaults() *GetSolMarketplaceTraitDataRequest {
	this := GetSolMarketplaceTraitDataRequest{}
	return &this
}

// GetCollectionTag returns the CollectionTag field value
func (o *GetSolMarketplaceTraitDataRequest) GetCollectionTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionTag
}

// GetCollectionTagOk returns a tuple with the CollectionTag field value
// and a boolean to check if the value has been set.
func (o *GetSolMarketplaceTraitDataRequest) GetCollectionTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionTag, true
}

// SetCollectionTag sets field value
func (o *GetSolMarketplaceTraitDataRequest) SetCollectionTag(v string) {
	o.CollectionTag = v
}

func (o GetSolMarketplaceTraitDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSolMarketplaceTraitDataRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_tag"] = o.CollectionTag
	return toSerialize, nil
}

type NullableGetSolMarketplaceTraitDataRequest struct {
	value *GetSolMarketplaceTraitDataRequest
	isSet bool
}

func (v NullableGetSolMarketplaceTraitDataRequest) Get() *GetSolMarketplaceTraitDataRequest {
	return v.value
}

func (v *NullableGetSolMarketplaceTraitDataRequest) Set(val *GetSolMarketplaceTraitDataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolMarketplaceTraitDataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolMarketplaceTraitDataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolMarketplaceTraitDataRequest(val *GetSolMarketplaceTraitDataRequest) *NullableGetSolMarketplaceTraitDataRequest {
	return &NullableGetSolMarketplaceTraitDataRequest{value: val, isSet: true}
}

func (v NullableGetSolMarketplaceTraitDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSolMarketplaceTraitDataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


