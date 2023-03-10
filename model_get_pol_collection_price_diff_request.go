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

// checks if the GetPolCollectionPriceDiffRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPolCollectionPriceDiffRequest{}

// GetPolCollectionPriceDiffRequest struct for GetPolCollectionPriceDiffRequest
type GetPolCollectionPriceDiffRequest struct {
	// The Polygon contract address to identify the collection.
	CollectionAddress string `json:"collection_address"`
	// The start date to pull data for calculations
	StartDate *string `json:"start_date,omitempty"`
	// The end date to pull data for calculations
	EndDate *string `json:"end_date,omitempty"`
	// Exclude suspected wash transactions?
	ExcludeWash *bool `json:"exclude_wash,omitempty"`
}

// NewGetPolCollectionPriceDiffRequest instantiates a new GetPolCollectionPriceDiffRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPolCollectionPriceDiffRequest(collectionAddress string) *GetPolCollectionPriceDiffRequest {
	this := GetPolCollectionPriceDiffRequest{}
	this.CollectionAddress = collectionAddress
	return &this
}

// NewGetPolCollectionPriceDiffRequestWithDefaults instantiates a new GetPolCollectionPriceDiffRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPolCollectionPriceDiffRequestWithDefaults() *GetPolCollectionPriceDiffRequest {
	this := GetPolCollectionPriceDiffRequest{}
	return &this
}

// GetCollectionAddress returns the CollectionAddress field value
func (o *GetPolCollectionPriceDiffRequest) GetCollectionAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value
// and a boolean to check if the value has been set.
func (o *GetPolCollectionPriceDiffRequest) GetCollectionAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionAddress, true
}

// SetCollectionAddress sets field value
func (o *GetPolCollectionPriceDiffRequest) SetCollectionAddress(v string) {
	o.CollectionAddress = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *GetPolCollectionPriceDiffRequest) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionPriceDiffRequest) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *GetPolCollectionPriceDiffRequest) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *GetPolCollectionPriceDiffRequest) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GetPolCollectionPriceDiffRequest) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionPriceDiffRequest) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GetPolCollectionPriceDiffRequest) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *GetPolCollectionPriceDiffRequest) SetEndDate(v string) {
	o.EndDate = &v
}

// GetExcludeWash returns the ExcludeWash field value if set, zero value otherwise.
func (o *GetPolCollectionPriceDiffRequest) GetExcludeWash() bool {
	if o == nil || IsNil(o.ExcludeWash) {
		var ret bool
		return ret
	}
	return *o.ExcludeWash
}

// GetExcludeWashOk returns a tuple with the ExcludeWash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionPriceDiffRequest) GetExcludeWashOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeWash) {
		return nil, false
	}
	return o.ExcludeWash, true
}

// HasExcludeWash returns a boolean if a field has been set.
func (o *GetPolCollectionPriceDiffRequest) HasExcludeWash() bool {
	if o != nil && !IsNil(o.ExcludeWash) {
		return true
	}

	return false
}

// SetExcludeWash gets a reference to the given bool and assigns it to the ExcludeWash field.
func (o *GetPolCollectionPriceDiffRequest) SetExcludeWash(v bool) {
	o.ExcludeWash = &v
}

func (o GetPolCollectionPriceDiffRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPolCollectionPriceDiffRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_address"] = o.CollectionAddress
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !IsNil(o.ExcludeWash) {
		toSerialize["exclude_wash"] = o.ExcludeWash
	}
	return toSerialize, nil
}

type NullableGetPolCollectionPriceDiffRequest struct {
	value *GetPolCollectionPriceDiffRequest
	isSet bool
}

func (v NullableGetPolCollectionPriceDiffRequest) Get() *GetPolCollectionPriceDiffRequest {
	return v.value
}

func (v *NullableGetPolCollectionPriceDiffRequest) Set(val *GetPolCollectionPriceDiffRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPolCollectionPriceDiffRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPolCollectionPriceDiffRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPolCollectionPriceDiffRequest(val *GetPolCollectionPriceDiffRequest) *NullableGetPolCollectionPriceDiffRequest {
	return &NullableGetPolCollectionPriceDiffRequest{value: val, isSet: true}
}

func (v NullableGetPolCollectionPriceDiffRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPolCollectionPriceDiffRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


