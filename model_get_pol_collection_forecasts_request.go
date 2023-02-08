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

// checks if the GetPolCollectionForecastsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPolCollectionForecastsRequest{}

// GetPolCollectionForecastsRequest struct for GetPolCollectionForecastsRequest
type GetPolCollectionForecastsRequest struct {
	// The contract address of the token collection.
	CollectionAddress string `json:"collection_address"`
	// The collection percentile(s)
	Percentiles []int32 `json:"percentiles,omitempty"`
	// Type of statistical forecasting model to be calculated as a 3-char string, e.g. `arc` for ARCH
	Voltype *string `json:"voltype,omitempty"`
	// The forecast horizon (i.e. the number of periods to forecast out)
	Horizon *int32 `json:"horizon,omitempty"`
	// The interval at which to calculate returns to base the forecasts upon, e.g. `1D` for daily, `1M` for monthly etc.
	Frequency *string `json:"frequency,omitempty"`
	// The distribution assumed to calculate parametric risk for.
	Dist *string `json:"dist,omitempty"`
	// The start date to pull data for calculations
	StartDate *string `json:"start_date,omitempty"`
	// The end date to pull data for calculations
	EndDate *string `json:"end_date,omitempty"`
	// Set to true, returns confidencve intervals at alpha significance for price on top of forecasts for returns and volatilities
	ReturnPriceForecasts *bool `json:"return_price_forecasts,omitempty"`
	// The significance level, e.g. 0.05 for 95% confidence
	Alpha *float32 `json:"alpha,omitempty"`
	// The currency to report results in
	ReptCurr *string `json:"rept_curr,omitempty"`
	// Exclude suspected wash transactions?
	ExcludeWash *bool `json:"exclude_wash,omitempty"`
	ArchParams *GetEthCollectionForecastsRequestArchParams `json:"arch_params,omitempty"`
}

// NewGetPolCollectionForecastsRequest instantiates a new GetPolCollectionForecastsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPolCollectionForecastsRequest(collectionAddress string) *GetPolCollectionForecastsRequest {
	this := GetPolCollectionForecastsRequest{}
	this.CollectionAddress = collectionAddress
	return &this
}

// NewGetPolCollectionForecastsRequestWithDefaults instantiates a new GetPolCollectionForecastsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPolCollectionForecastsRequestWithDefaults() *GetPolCollectionForecastsRequest {
	this := GetPolCollectionForecastsRequest{}
	return &this
}

// GetCollectionAddress returns the CollectionAddress field value
func (o *GetPolCollectionForecastsRequest) GetCollectionAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value
// and a boolean to check if the value has been set.
func (o *GetPolCollectionForecastsRequest) GetCollectionAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionAddress, true
}

// SetCollectionAddress sets field value
func (o *GetPolCollectionForecastsRequest) SetCollectionAddress(v string) {
	o.CollectionAddress = v
}

// GetPercentiles returns the Percentiles field value if set, zero value otherwise.
func (o *GetPolCollectionForecastsRequest) GetPercentiles() []int32 {
	if o == nil || IsNil(o.Percentiles) {
		var ret []int32
		return ret
	}
	return o.Percentiles
}

// GetPercentilesOk returns a tuple with the Percentiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionForecastsRequest) GetPercentilesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Percentiles) {
		return nil, false
	}
	return o.Percentiles, true
}

// HasPercentiles returns a boolean if a field has been set.
func (o *GetPolCollectionForecastsRequest) HasPercentiles() bool {
	if o != nil && !IsNil(o.Percentiles) {
		return true
	}

	return false
}

// SetPercentiles gets a reference to the given []int32 and assigns it to the Percentiles field.
func (o *GetPolCollectionForecastsRequest) SetPercentiles(v []int32) {
	o.Percentiles = v
}

// GetVoltype returns the Voltype field value if set, zero value otherwise.
func (o *GetPolCollectionForecastsRequest) GetVoltype() string {
	if o == nil || IsNil(o.Voltype) {
		var ret string
		return ret
	}
	return *o.Voltype
}

// GetVoltypeOk returns a tuple with the Voltype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionForecastsRequest) GetVoltypeOk() (*string, bool) {
	if o == nil || IsNil(o.Voltype) {
		return nil, false
	}
	return o.Voltype, true
}

// HasVoltype returns a boolean if a field has been set.
func (o *GetPolCollectionForecastsRequest) HasVoltype() bool {
	if o != nil && !IsNil(o.Voltype) {
		return true
	}

	return false
}

// SetVoltype gets a reference to the given string and assigns it to the Voltype field.
func (o *GetPolCollectionForecastsRequest) SetVoltype(v string) {
	o.Voltype = &v
}

// GetHorizon returns the Horizon field value if set, zero value otherwise.
func (o *GetPolCollectionForecastsRequest) GetHorizon() int32 {
	if o == nil || IsNil(o.Horizon) {
		var ret int32
		return ret
	}
	return *o.Horizon
}

// GetHorizonOk returns a tuple with the Horizon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionForecastsRequest) GetHorizonOk() (*int32, bool) {
	if o == nil || IsNil(o.Horizon) {
		return nil, false
	}
	return o.Horizon, true
}

// HasHorizon returns a boolean if a field has been set.
func (o *GetPolCollectionForecastsRequest) HasHorizon() bool {
	if o != nil && !IsNil(o.Horizon) {
		return true
	}

	return false
}

// SetHorizon gets a reference to the given int32 and assigns it to the Horizon field.
func (o *GetPolCollectionForecastsRequest) SetHorizon(v int32) {
	o.Horizon = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *GetPolCollectionForecastsRequest) GetFrequency() string {
	if o == nil || IsNil(o.Frequency) {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionForecastsRequest) GetFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *GetPolCollectionForecastsRequest) HasFrequency() bool {
	if o != nil && !IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *GetPolCollectionForecastsRequest) SetFrequency(v string) {
	o.Frequency = &v
}

// GetDist returns the Dist field value if set, zero value otherwise.
func (o *GetPolCollectionForecastsRequest) GetDist() string {
	if o == nil || IsNil(o.Dist) {
		var ret string
		return ret
	}
	return *o.Dist
}

// GetDistOk returns a tuple with the Dist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionForecastsRequest) GetDistOk() (*string, bool) {
	if o == nil || IsNil(o.Dist) {
		return nil, false
	}
	return o.Dist, true
}

// HasDist returns a boolean if a field has been set.
func (o *GetPolCollectionForecastsRequest) HasDist() bool {
	if o != nil && !IsNil(o.Dist) {
		return true
	}

	return false
}

// SetDist gets a reference to the given string and assigns it to the Dist field.
func (o *GetPolCollectionForecastsRequest) SetDist(v string) {
	o.Dist = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *GetPolCollectionForecastsRequest) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionForecastsRequest) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *GetPolCollectionForecastsRequest) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *GetPolCollectionForecastsRequest) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GetPolCollectionForecastsRequest) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionForecastsRequest) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GetPolCollectionForecastsRequest) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *GetPolCollectionForecastsRequest) SetEndDate(v string) {
	o.EndDate = &v
}

// GetReturnPriceForecasts returns the ReturnPriceForecasts field value if set, zero value otherwise.
func (o *GetPolCollectionForecastsRequest) GetReturnPriceForecasts() bool {
	if o == nil || IsNil(o.ReturnPriceForecasts) {
		var ret bool
		return ret
	}
	return *o.ReturnPriceForecasts
}

// GetReturnPriceForecastsOk returns a tuple with the ReturnPriceForecasts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionForecastsRequest) GetReturnPriceForecastsOk() (*bool, bool) {
	if o == nil || IsNil(o.ReturnPriceForecasts) {
		return nil, false
	}
	return o.ReturnPriceForecasts, true
}

// HasReturnPriceForecasts returns a boolean if a field has been set.
func (o *GetPolCollectionForecastsRequest) HasReturnPriceForecasts() bool {
	if o != nil && !IsNil(o.ReturnPriceForecasts) {
		return true
	}

	return false
}

// SetReturnPriceForecasts gets a reference to the given bool and assigns it to the ReturnPriceForecasts field.
func (o *GetPolCollectionForecastsRequest) SetReturnPriceForecasts(v bool) {
	o.ReturnPriceForecasts = &v
}

// GetAlpha returns the Alpha field value if set, zero value otherwise.
func (o *GetPolCollectionForecastsRequest) GetAlpha() float32 {
	if o == nil || IsNil(o.Alpha) {
		var ret float32
		return ret
	}
	return *o.Alpha
}

// GetAlphaOk returns a tuple with the Alpha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionForecastsRequest) GetAlphaOk() (*float32, bool) {
	if o == nil || IsNil(o.Alpha) {
		return nil, false
	}
	return o.Alpha, true
}

// HasAlpha returns a boolean if a field has been set.
func (o *GetPolCollectionForecastsRequest) HasAlpha() bool {
	if o != nil && !IsNil(o.Alpha) {
		return true
	}

	return false
}

// SetAlpha gets a reference to the given float32 and assigns it to the Alpha field.
func (o *GetPolCollectionForecastsRequest) SetAlpha(v float32) {
	o.Alpha = &v
}

// GetReptCurr returns the ReptCurr field value if set, zero value otherwise.
func (o *GetPolCollectionForecastsRequest) GetReptCurr() string {
	if o == nil || IsNil(o.ReptCurr) {
		var ret string
		return ret
	}
	return *o.ReptCurr
}

// GetReptCurrOk returns a tuple with the ReptCurr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionForecastsRequest) GetReptCurrOk() (*string, bool) {
	if o == nil || IsNil(o.ReptCurr) {
		return nil, false
	}
	return o.ReptCurr, true
}

// HasReptCurr returns a boolean if a field has been set.
func (o *GetPolCollectionForecastsRequest) HasReptCurr() bool {
	if o != nil && !IsNil(o.ReptCurr) {
		return true
	}

	return false
}

// SetReptCurr gets a reference to the given string and assigns it to the ReptCurr field.
func (o *GetPolCollectionForecastsRequest) SetReptCurr(v string) {
	o.ReptCurr = &v
}

// GetExcludeWash returns the ExcludeWash field value if set, zero value otherwise.
func (o *GetPolCollectionForecastsRequest) GetExcludeWash() bool {
	if o == nil || IsNil(o.ExcludeWash) {
		var ret bool
		return ret
	}
	return *o.ExcludeWash
}

// GetExcludeWashOk returns a tuple with the ExcludeWash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionForecastsRequest) GetExcludeWashOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeWash) {
		return nil, false
	}
	return o.ExcludeWash, true
}

// HasExcludeWash returns a boolean if a field has been set.
func (o *GetPolCollectionForecastsRequest) HasExcludeWash() bool {
	if o != nil && !IsNil(o.ExcludeWash) {
		return true
	}

	return false
}

// SetExcludeWash gets a reference to the given bool and assigns it to the ExcludeWash field.
func (o *GetPolCollectionForecastsRequest) SetExcludeWash(v bool) {
	o.ExcludeWash = &v
}

// GetArchParams returns the ArchParams field value if set, zero value otherwise.
func (o *GetPolCollectionForecastsRequest) GetArchParams() GetEthCollectionForecastsRequestArchParams {
	if o == nil || IsNil(o.ArchParams) {
		var ret GetEthCollectionForecastsRequestArchParams
		return ret
	}
	return *o.ArchParams
}

// GetArchParamsOk returns a tuple with the ArchParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCollectionForecastsRequest) GetArchParamsOk() (*GetEthCollectionForecastsRequestArchParams, bool) {
	if o == nil || IsNil(o.ArchParams) {
		return nil, false
	}
	return o.ArchParams, true
}

// HasArchParams returns a boolean if a field has been set.
func (o *GetPolCollectionForecastsRequest) HasArchParams() bool {
	if o != nil && !IsNil(o.ArchParams) {
		return true
	}

	return false
}

// SetArchParams gets a reference to the given GetEthCollectionForecastsRequestArchParams and assigns it to the ArchParams field.
func (o *GetPolCollectionForecastsRequest) SetArchParams(v GetEthCollectionForecastsRequestArchParams) {
	o.ArchParams = &v
}

func (o GetPolCollectionForecastsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPolCollectionForecastsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_address"] = o.CollectionAddress
	if !IsNil(o.Percentiles) {
		toSerialize["percentiles"] = o.Percentiles
	}
	if !IsNil(o.Voltype) {
		toSerialize["voltype"] = o.Voltype
	}
	if !IsNil(o.Horizon) {
		toSerialize["horizon"] = o.Horizon
	}
	if !IsNil(o.Frequency) {
		toSerialize["frequency"] = o.Frequency
	}
	if !IsNil(o.Dist) {
		toSerialize["dist"] = o.Dist
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !IsNil(o.ReturnPriceForecasts) {
		toSerialize["return_price_forecasts"] = o.ReturnPriceForecasts
	}
	if !IsNil(o.Alpha) {
		toSerialize["alpha"] = o.Alpha
	}
	if !IsNil(o.ReptCurr) {
		toSerialize["rept_curr"] = o.ReptCurr
	}
	if !IsNil(o.ExcludeWash) {
		toSerialize["exclude_wash"] = o.ExcludeWash
	}
	if !IsNil(o.ArchParams) {
		toSerialize["arch_params"] = o.ArchParams
	}
	return toSerialize, nil
}

type NullableGetPolCollectionForecastsRequest struct {
	value *GetPolCollectionForecastsRequest
	isSet bool
}

func (v NullableGetPolCollectionForecastsRequest) Get() *GetPolCollectionForecastsRequest {
	return v.value
}

func (v *NullableGetPolCollectionForecastsRequest) Set(val *GetPolCollectionForecastsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPolCollectionForecastsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPolCollectionForecastsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPolCollectionForecastsRequest(val *GetPolCollectionForecastsRequest) *NullableGetPolCollectionForecastsRequest {
	return &NullableGetPolCollectionForecastsRequest{value: val, isSet: true}
}

func (v NullableGetPolCollectionForecastsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPolCollectionForecastsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


