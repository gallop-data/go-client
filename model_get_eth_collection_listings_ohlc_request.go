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

// checks if the GetEthCollectionListingsOHLCRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEthCollectionListingsOHLCRequest{}

// GetEthCollectionListingsOHLCRequest struct for GetEthCollectionListingsOHLCRequest
type GetEthCollectionListingsOHLCRequest struct {
	// The Ethereum contract address to identify the collection.
	CollectionAddress string `json:"collection_address"`
	// If `true`, report only historical floor prices. Otherwise, report OHFC candlesticks, number of active listings, number of unique owners and the average age of open listings.
	FloorOnly *bool `json:"floor_only,omitempty"`
	// The interval at which to return Floor prices / OHLF, e.g. `1D` for daily, `1M` for monthly etc. Must be >= `6H`
	Frequency *string `json:"frequency,omitempty"`
	// The currency to report results in
	ReptCurr *string `json:"rept_curr,omitempty"`
	// The ISO 8601 start date/datetime to return results for
	ReportStartDate *string `json:"report_start_date,omitempty"`
	// The ISO 8601 end date/datetime to return results for
	ReportEndDate *string `json:"report_end_date,omitempty"`
}

// NewGetEthCollectionListingsOHLCRequest instantiates a new GetEthCollectionListingsOHLCRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEthCollectionListingsOHLCRequest(collectionAddress string) *GetEthCollectionListingsOHLCRequest {
	this := GetEthCollectionListingsOHLCRequest{}
	this.CollectionAddress = collectionAddress
	return &this
}

// NewGetEthCollectionListingsOHLCRequestWithDefaults instantiates a new GetEthCollectionListingsOHLCRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEthCollectionListingsOHLCRequestWithDefaults() *GetEthCollectionListingsOHLCRequest {
	this := GetEthCollectionListingsOHLCRequest{}
	return &this
}

// GetCollectionAddress returns the CollectionAddress field value
func (o *GetEthCollectionListingsOHLCRequest) GetCollectionAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value
// and a boolean to check if the value has been set.
func (o *GetEthCollectionListingsOHLCRequest) GetCollectionAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionAddress, true
}

// SetCollectionAddress sets field value
func (o *GetEthCollectionListingsOHLCRequest) SetCollectionAddress(v string) {
	o.CollectionAddress = v
}

// GetFloorOnly returns the FloorOnly field value if set, zero value otherwise.
func (o *GetEthCollectionListingsOHLCRequest) GetFloorOnly() bool {
	if o == nil || IsNil(o.FloorOnly) {
		var ret bool
		return ret
	}
	return *o.FloorOnly
}

// GetFloorOnlyOk returns a tuple with the FloorOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionListingsOHLCRequest) GetFloorOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.FloorOnly) {
		return nil, false
	}
	return o.FloorOnly, true
}

// HasFloorOnly returns a boolean if a field has been set.
func (o *GetEthCollectionListingsOHLCRequest) HasFloorOnly() bool {
	if o != nil && !IsNil(o.FloorOnly) {
		return true
	}

	return false
}

// SetFloorOnly gets a reference to the given bool and assigns it to the FloorOnly field.
func (o *GetEthCollectionListingsOHLCRequest) SetFloorOnly(v bool) {
	o.FloorOnly = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *GetEthCollectionListingsOHLCRequest) GetFrequency() string {
	if o == nil || IsNil(o.Frequency) {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionListingsOHLCRequest) GetFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *GetEthCollectionListingsOHLCRequest) HasFrequency() bool {
	if o != nil && !IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *GetEthCollectionListingsOHLCRequest) SetFrequency(v string) {
	o.Frequency = &v
}

// GetReptCurr returns the ReptCurr field value if set, zero value otherwise.
func (o *GetEthCollectionListingsOHLCRequest) GetReptCurr() string {
	if o == nil || IsNil(o.ReptCurr) {
		var ret string
		return ret
	}
	return *o.ReptCurr
}

// GetReptCurrOk returns a tuple with the ReptCurr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionListingsOHLCRequest) GetReptCurrOk() (*string, bool) {
	if o == nil || IsNil(o.ReptCurr) {
		return nil, false
	}
	return o.ReptCurr, true
}

// HasReptCurr returns a boolean if a field has been set.
func (o *GetEthCollectionListingsOHLCRequest) HasReptCurr() bool {
	if o != nil && !IsNil(o.ReptCurr) {
		return true
	}

	return false
}

// SetReptCurr gets a reference to the given string and assigns it to the ReptCurr field.
func (o *GetEthCollectionListingsOHLCRequest) SetReptCurr(v string) {
	o.ReptCurr = &v
}

// GetReportStartDate returns the ReportStartDate field value if set, zero value otherwise.
func (o *GetEthCollectionListingsOHLCRequest) GetReportStartDate() string {
	if o == nil || IsNil(o.ReportStartDate) {
		var ret string
		return ret
	}
	return *o.ReportStartDate
}

// GetReportStartDateOk returns a tuple with the ReportStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionListingsOHLCRequest) GetReportStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.ReportStartDate) {
		return nil, false
	}
	return o.ReportStartDate, true
}

// HasReportStartDate returns a boolean if a field has been set.
func (o *GetEthCollectionListingsOHLCRequest) HasReportStartDate() bool {
	if o != nil && !IsNil(o.ReportStartDate) {
		return true
	}

	return false
}

// SetReportStartDate gets a reference to the given string and assigns it to the ReportStartDate field.
func (o *GetEthCollectionListingsOHLCRequest) SetReportStartDate(v string) {
	o.ReportStartDate = &v
}

// GetReportEndDate returns the ReportEndDate field value if set, zero value otherwise.
func (o *GetEthCollectionListingsOHLCRequest) GetReportEndDate() string {
	if o == nil || IsNil(o.ReportEndDate) {
		var ret string
		return ret
	}
	return *o.ReportEndDate
}

// GetReportEndDateOk returns a tuple with the ReportEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionListingsOHLCRequest) GetReportEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.ReportEndDate) {
		return nil, false
	}
	return o.ReportEndDate, true
}

// HasReportEndDate returns a boolean if a field has been set.
func (o *GetEthCollectionListingsOHLCRequest) HasReportEndDate() bool {
	if o != nil && !IsNil(o.ReportEndDate) {
		return true
	}

	return false
}

// SetReportEndDate gets a reference to the given string and assigns it to the ReportEndDate field.
func (o *GetEthCollectionListingsOHLCRequest) SetReportEndDate(v string) {
	o.ReportEndDate = &v
}

func (o GetEthCollectionListingsOHLCRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEthCollectionListingsOHLCRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_address"] = o.CollectionAddress
	if !IsNil(o.FloorOnly) {
		toSerialize["floor_only"] = o.FloorOnly
	}
	if !IsNil(o.Frequency) {
		toSerialize["frequency"] = o.Frequency
	}
	if !IsNil(o.ReptCurr) {
		toSerialize["rept_curr"] = o.ReptCurr
	}
	if !IsNil(o.ReportStartDate) {
		toSerialize["report_start_date"] = o.ReportStartDate
	}
	if !IsNil(o.ReportEndDate) {
		toSerialize["report_end_date"] = o.ReportEndDate
	}
	return toSerialize, nil
}

type NullableGetEthCollectionListingsOHLCRequest struct {
	value *GetEthCollectionListingsOHLCRequest
	isSet bool
}

func (v NullableGetEthCollectionListingsOHLCRequest) Get() *GetEthCollectionListingsOHLCRequest {
	return v.value
}

func (v *NullableGetEthCollectionListingsOHLCRequest) Set(val *GetEthCollectionListingsOHLCRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEthCollectionListingsOHLCRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEthCollectionListingsOHLCRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEthCollectionListingsOHLCRequest(val *GetEthCollectionListingsOHLCRequest) *NullableGetEthCollectionListingsOHLCRequest {
	return &NullableGetEthCollectionListingsOHLCRequest{value: val, isSet: true}
}

func (v NullableGetEthCollectionListingsOHLCRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEthCollectionListingsOHLCRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


