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

// checks if the GetSolCollectionPriceDiffRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSolCollectionPriceDiffRequest{}

// GetSolCollectionPriceDiffRequest struct for GetSolCollectionPriceDiffRequest
type GetSolCollectionPriceDiffRequest struct {
	// The Gallop tag to identify the collection.
	CollectionTag string `json:"collection_tag"`
	// The start date to pull data for calculations
	StartDate *string `json:"start_date,omitempty"`
	// The end date to pull data for calculations
	EndDate *string `json:"end_date,omitempty"`
	// Exclude suspected wash transactions?
	ExcludeWash *bool `json:"exclude_wash,omitempty"`
}

// NewGetSolCollectionPriceDiffRequest instantiates a new GetSolCollectionPriceDiffRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolCollectionPriceDiffRequest(collectionTag string) *GetSolCollectionPriceDiffRequest {
	this := GetSolCollectionPriceDiffRequest{}
	this.CollectionTag = collectionTag
	return &this
}

// NewGetSolCollectionPriceDiffRequestWithDefaults instantiates a new GetSolCollectionPriceDiffRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolCollectionPriceDiffRequestWithDefaults() *GetSolCollectionPriceDiffRequest {
	this := GetSolCollectionPriceDiffRequest{}
	return &this
}

// GetCollectionTag returns the CollectionTag field value
func (o *GetSolCollectionPriceDiffRequest) GetCollectionTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionTag
}

// GetCollectionTagOk returns a tuple with the CollectionTag field value
// and a boolean to check if the value has been set.
func (o *GetSolCollectionPriceDiffRequest) GetCollectionTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionTag, true
}

// SetCollectionTag sets field value
func (o *GetSolCollectionPriceDiffRequest) SetCollectionTag(v string) {
	o.CollectionTag = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *GetSolCollectionPriceDiffRequest) GetStartDate() string {
	if o == nil || isNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolCollectionPriceDiffRequest) GetStartDateOk() (*string, bool) {
	if o == nil || isNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *GetSolCollectionPriceDiffRequest) HasStartDate() bool {
	if o != nil && !isNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *GetSolCollectionPriceDiffRequest) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GetSolCollectionPriceDiffRequest) GetEndDate() string {
	if o == nil || isNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolCollectionPriceDiffRequest) GetEndDateOk() (*string, bool) {
	if o == nil || isNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GetSolCollectionPriceDiffRequest) HasEndDate() bool {
	if o != nil && !isNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *GetSolCollectionPriceDiffRequest) SetEndDate(v string) {
	o.EndDate = &v
}

// GetExcludeWash returns the ExcludeWash field value if set, zero value otherwise.
func (o *GetSolCollectionPriceDiffRequest) GetExcludeWash() bool {
	if o == nil || isNil(o.ExcludeWash) {
		var ret bool
		return ret
	}
	return *o.ExcludeWash
}

// GetExcludeWashOk returns a tuple with the ExcludeWash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolCollectionPriceDiffRequest) GetExcludeWashOk() (*bool, bool) {
	if o == nil || isNil(o.ExcludeWash) {
		return nil, false
	}
	return o.ExcludeWash, true
}

// HasExcludeWash returns a boolean if a field has been set.
func (o *GetSolCollectionPriceDiffRequest) HasExcludeWash() bool {
	if o != nil && !isNil(o.ExcludeWash) {
		return true
	}

	return false
}

// SetExcludeWash gets a reference to the given bool and assigns it to the ExcludeWash field.
func (o *GetSolCollectionPriceDiffRequest) SetExcludeWash(v bool) {
	o.ExcludeWash = &v
}

func (o GetSolCollectionPriceDiffRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSolCollectionPriceDiffRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_tag"] = o.CollectionTag
	if !isNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !isNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !isNil(o.ExcludeWash) {
		toSerialize["exclude_wash"] = o.ExcludeWash
	}
	return toSerialize, nil
}

type NullableGetSolCollectionPriceDiffRequest struct {
	value *GetSolCollectionPriceDiffRequest
	isSet bool
}

func (v NullableGetSolCollectionPriceDiffRequest) Get() *GetSolCollectionPriceDiffRequest {
	return v.value
}

func (v *NullableGetSolCollectionPriceDiffRequest) Set(val *GetSolCollectionPriceDiffRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolCollectionPriceDiffRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolCollectionPriceDiffRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolCollectionPriceDiffRequest(val *GetSolCollectionPriceDiffRequest) *NullableGetSolCollectionPriceDiffRequest {
	return &NullableGetSolCollectionPriceDiffRequest{value: val, isSet: true}
}

func (v NullableGetSolCollectionPriceDiffRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSolCollectionPriceDiffRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


