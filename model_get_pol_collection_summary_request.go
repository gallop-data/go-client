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

// checks if the GetPolCollectionSummaryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPolCollectionSummaryRequest{}

// GetPolCollectionSummaryRequest struct for GetPolCollectionSummaryRequest
type GetPolCollectionSummaryRequest struct {
	// The contract address of the collection.
	CollectionAddress string `json:"collection_address"`
	// An attribute of the NFT.
	GroupBy *string `json:"group_by,omitempty"`
	// The start date to pull data for calculations
	StartDate *string `json:"start_date,omitempty"`
	// The end date to pull data for calculations
	EndDate *string `json:"end_date,omitempty"`
	// The currency to report results in
	ReptCurr *string `json:"rept_curr,omitempty"`
	// Exclude suspected wash transactions?
	ExcludeWash *bool `json:"exclude_wash,omitempty"`
}

// NewGetPolCollectionSummaryRequest instantiates a new GetPolCollectionSummaryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPolCollectionSummaryRequest(collectionAddress string) *GetPolCollectionSummaryRequest {
	this := GetPolCollectionSummaryRequest{}
	this.CollectionAddress = collectionAddress
	return &this
}

// NewGetPolCollectionSummaryRequestWithDefaults instantiates a new GetPolCollectionSummaryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPolCollectionSummaryRequestWithDefaults() *GetPolCollectionSummaryRequest {
	this := GetPolCollectionSummaryRequest{}
	return &this
}

// GetCollectionAddress returns the CollectionAddress field value
func (o *GetPolCollectionSummaryRequest) GetCollectionAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value
// and a boolean to check if the value has been set.
func (o *GetPolCollectionSummaryRequest) GetCollectionAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionAddress, true
}

// SetCollectionAddress sets field value
func (o *GetPolCollectionSummaryRequest) SetCollectionAddress(v string) {
	o.CollectionAddress = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *GetPolCollectionSummaryRequest) GetGroupBy() string {
	if o == nil || isNil(o.GroupBy) {
		var ret string
		return ret
	}
	return *o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionSummaryRequest) GetGroupByOk() (*string, bool) {
	if o == nil || isNil(o.GroupBy) {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *GetPolCollectionSummaryRequest) HasGroupBy() bool {
	if o != nil && !isNil(o.GroupBy) {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given string and assigns it to the GroupBy field.
func (o *GetPolCollectionSummaryRequest) SetGroupBy(v string) {
	o.GroupBy = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *GetPolCollectionSummaryRequest) GetStartDate() string {
	if o == nil || isNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionSummaryRequest) GetStartDateOk() (*string, bool) {
	if o == nil || isNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *GetPolCollectionSummaryRequest) HasStartDate() bool {
	if o != nil && !isNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *GetPolCollectionSummaryRequest) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GetPolCollectionSummaryRequest) GetEndDate() string {
	if o == nil || isNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionSummaryRequest) GetEndDateOk() (*string, bool) {
	if o == nil || isNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GetPolCollectionSummaryRequest) HasEndDate() bool {
	if o != nil && !isNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *GetPolCollectionSummaryRequest) SetEndDate(v string) {
	o.EndDate = &v
}

// GetReptCurr returns the ReptCurr field value if set, zero value otherwise.
func (o *GetPolCollectionSummaryRequest) GetReptCurr() string {
	if o == nil || isNil(o.ReptCurr) {
		var ret string
		return ret
	}
	return *o.ReptCurr
}

// GetReptCurrOk returns a tuple with the ReptCurr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionSummaryRequest) GetReptCurrOk() (*string, bool) {
	if o == nil || isNil(o.ReptCurr) {
		return nil, false
	}
	return o.ReptCurr, true
}

// HasReptCurr returns a boolean if a field has been set.
func (o *GetPolCollectionSummaryRequest) HasReptCurr() bool {
	if o != nil && !isNil(o.ReptCurr) {
		return true
	}

	return false
}

// SetReptCurr gets a reference to the given string and assigns it to the ReptCurr field.
func (o *GetPolCollectionSummaryRequest) SetReptCurr(v string) {
	o.ReptCurr = &v
}

// GetExcludeWash returns the ExcludeWash field value if set, zero value otherwise.
func (o *GetPolCollectionSummaryRequest) GetExcludeWash() bool {
	if o == nil || isNil(o.ExcludeWash) {
		var ret bool
		return ret
	}
	return *o.ExcludeWash
}

// GetExcludeWashOk returns a tuple with the ExcludeWash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionSummaryRequest) GetExcludeWashOk() (*bool, bool) {
	if o == nil || isNil(o.ExcludeWash) {
		return nil, false
	}
	return o.ExcludeWash, true
}

// HasExcludeWash returns a boolean if a field has been set.
func (o *GetPolCollectionSummaryRequest) HasExcludeWash() bool {
	if o != nil && !isNil(o.ExcludeWash) {
		return true
	}

	return false
}

// SetExcludeWash gets a reference to the given bool and assigns it to the ExcludeWash field.
func (o *GetPolCollectionSummaryRequest) SetExcludeWash(v bool) {
	o.ExcludeWash = &v
}

func (o GetPolCollectionSummaryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPolCollectionSummaryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_address"] = o.CollectionAddress
	if !isNil(o.GroupBy) {
		toSerialize["group_by"] = o.GroupBy
	}
	if !isNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !isNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !isNil(o.ReptCurr) {
		toSerialize["rept_curr"] = o.ReptCurr
	}
	if !isNil(o.ExcludeWash) {
		toSerialize["exclude_wash"] = o.ExcludeWash
	}
	return toSerialize, nil
}

type NullableGetPolCollectionSummaryRequest struct {
	value *GetPolCollectionSummaryRequest
	isSet bool
}

func (v NullableGetPolCollectionSummaryRequest) Get() *GetPolCollectionSummaryRequest {
	return v.value
}

func (v *NullableGetPolCollectionSummaryRequest) Set(val *GetPolCollectionSummaryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPolCollectionSummaryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPolCollectionSummaryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPolCollectionSummaryRequest(val *GetPolCollectionSummaryRequest) *NullableGetPolCollectionSummaryRequest {
	return &NullableGetPolCollectionSummaryRequest{value: val, isSet: true}
}

func (v NullableGetPolCollectionSummaryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPolCollectionSummaryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

