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

// checks if the GetSolCollectionForecastsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSolCollectionForecastsRequest{}

// GetSolCollectionForecastsRequest struct for GetSolCollectionForecastsRequest
type GetSolCollectionForecastsRequest struct {
	// The Gallop tag to identify the collection.
	CollectionTag string `json:"collection_tag"`
	// The collection percentile(s)
	Percentiles []int32 `json:"percentiles,omitempty"`
	// Type of statistical forecasting model to be calculated as a 3-char string, e.g. 'arc' for ARCH
	Voltype *string `json:"voltype,omitempty"`
	// The forecast horizon (i.e. the number of periods to forecast out)
	Horizon *int32 `json:"horizon,omitempty"`
	// The interval at which to calculate returns to base the forecasts upon, e.g. `1D` for daily, `1M` for monthly etc.
	Frequency *string `json:"frequency,omitempty"`
	// The distribution assumed to calculate parametric risk for
	Dist *string `json:"dist,omitempty"`
	// The start date to pull data for calculations
	StartDate *string `json:"start_date,omitempty"`
	// The end date to pull data for calculations
	EndDate *string `json:"end_date,omitempty"`
	// Set to True, returns confidencve intervals at alpha significance for price on top of forecasts for returns and volatilities
	ReturnPriceForecasts *bool `json:"return_price_forecasts,omitempty"`
	// The significance level, e.g. 0.05 for 95% confidence
	Alpha *float32 `json:"alpha,omitempty"`
	// The currency to report results in
	ReptCurr *string `json:"rept_curr,omitempty"`
	// Exclude suspected wash transactions?
	ExcludeWash *bool `json:"exclude_wash,omitempty"`
	ArchParams *GetSolCollectionForecastsRequestArchParams `json:"arch_params,omitempty"`
}

// NewGetSolCollectionForecastsRequest instantiates a new GetSolCollectionForecastsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolCollectionForecastsRequest(collectionTag string) *GetSolCollectionForecastsRequest {
	this := GetSolCollectionForecastsRequest{}
	this.CollectionTag = collectionTag
	return &this
}

// NewGetSolCollectionForecastsRequestWithDefaults instantiates a new GetSolCollectionForecastsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolCollectionForecastsRequestWithDefaults() *GetSolCollectionForecastsRequest {
	this := GetSolCollectionForecastsRequest{}
	return &this
}

// GetCollectionTag returns the CollectionTag field value
func (o *GetSolCollectionForecastsRequest) GetCollectionTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionTag
}

// GetCollectionTagOk returns a tuple with the CollectionTag field value
// and a boolean to check if the value has been set.
func (o *GetSolCollectionForecastsRequest) GetCollectionTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionTag, true
}

// SetCollectionTag sets field value
func (o *GetSolCollectionForecastsRequest) SetCollectionTag(v string) {
	o.CollectionTag = v
}

// GetPercentiles returns the Percentiles field value if set, zero value otherwise.
func (o *GetSolCollectionForecastsRequest) GetPercentiles() []int32 {
	if o == nil || isNil(o.Percentiles) {
		var ret []int32
		return ret
	}
	return o.Percentiles
}

// GetPercentilesOk returns a tuple with the Percentiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolCollectionForecastsRequest) GetPercentilesOk() ([]int32, bool) {
	if o == nil || isNil(o.Percentiles) {
		return nil, false
	}
	return o.Percentiles, true
}

// HasPercentiles returns a boolean if a field has been set.
func (o *GetSolCollectionForecastsRequest) HasPercentiles() bool {
	if o != nil && !isNil(o.Percentiles) {
		return true
	}

	return false
}

// SetPercentiles gets a reference to the given []int32 and assigns it to the Percentiles field.
func (o *GetSolCollectionForecastsRequest) SetPercentiles(v []int32) {
	o.Percentiles = v
}

// GetVoltype returns the Voltype field value if set, zero value otherwise.
func (o *GetSolCollectionForecastsRequest) GetVoltype() string {
	if o == nil || isNil(o.Voltype) {
		var ret string
		return ret
	}
	return *o.Voltype
}

// GetVoltypeOk returns a tuple with the Voltype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolCollectionForecastsRequest) GetVoltypeOk() (*string, bool) {
	if o == nil || isNil(o.Voltype) {
		return nil, false
	}
	return o.Voltype, true
}

// HasVoltype returns a boolean if a field has been set.
func (o *GetSolCollectionForecastsRequest) HasVoltype() bool {
	if o != nil && !isNil(o.Voltype) {
		return true
	}

	return false
}

// SetVoltype gets a reference to the given string and assigns it to the Voltype field.
func (o *GetSolCollectionForecastsRequest) SetVoltype(v string) {
	o.Voltype = &v
}

// GetHorizon returns the Horizon field value if set, zero value otherwise.
func (o *GetSolCollectionForecastsRequest) GetHorizon() int32 {
	if o == nil || isNil(o.Horizon) {
		var ret int32
		return ret
	}
	return *o.Horizon
}

// GetHorizonOk returns a tuple with the Horizon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolCollectionForecastsRequest) GetHorizonOk() (*int32, bool) {
	if o == nil || isNil(o.Horizon) {
		return nil, false
	}
	return o.Horizon, true
}

// HasHorizon returns a boolean if a field has been set.
func (o *GetSolCollectionForecastsRequest) HasHorizon() bool {
	if o != nil && !isNil(o.Horizon) {
		return true
	}

	return false
}

// SetHorizon gets a reference to the given int32 and assigns it to the Horizon field.
func (o *GetSolCollectionForecastsRequest) SetHorizon(v int32) {
	o.Horizon = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *GetSolCollectionForecastsRequest) GetFrequency() string {
	if o == nil || isNil(o.Frequency) {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolCollectionForecastsRequest) GetFrequencyOk() (*string, bool) {
	if o == nil || isNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *GetSolCollectionForecastsRequest) HasFrequency() bool {
	if o != nil && !isNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *GetSolCollectionForecastsRequest) SetFrequency(v string) {
	o.Frequency = &v
}

// GetDist returns the Dist field value if set, zero value otherwise.
func (o *GetSolCollectionForecastsRequest) GetDist() string {
	if o == nil || isNil(o.Dist) {
		var ret string
		return ret
	}
	return *o.Dist
}

// GetDistOk returns a tuple with the Dist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolCollectionForecastsRequest) GetDistOk() (*string, bool) {
	if o == nil || isNil(o.Dist) {
		return nil, false
	}
	return o.Dist, true
}

// HasDist returns a boolean if a field has been set.
func (o *GetSolCollectionForecastsRequest) HasDist() bool {
	if o != nil && !isNil(o.Dist) {
		return true
	}

	return false
}

// SetDist gets a reference to the given string and assigns it to the Dist field.
func (o *GetSolCollectionForecastsRequest) SetDist(v string) {
	o.Dist = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *GetSolCollectionForecastsRequest) GetStartDate() string {
	if o == nil || isNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolCollectionForecastsRequest) GetStartDateOk() (*string, bool) {
	if o == nil || isNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *GetSolCollectionForecastsRequest) HasStartDate() bool {
	if o != nil && !isNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *GetSolCollectionForecastsRequest) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GetSolCollectionForecastsRequest) GetEndDate() string {
	if o == nil || isNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolCollectionForecastsRequest) GetEndDateOk() (*string, bool) {
	if o == nil || isNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GetSolCollectionForecastsRequest) HasEndDate() bool {
	if o != nil && !isNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *GetSolCollectionForecastsRequest) SetEndDate(v string) {
	o.EndDate = &v
}

// GetReturnPriceForecasts returns the ReturnPriceForecasts field value if set, zero value otherwise.
func (o *GetSolCollectionForecastsRequest) GetReturnPriceForecasts() bool {
	if o == nil || isNil(o.ReturnPriceForecasts) {
		var ret bool
		return ret
	}
	return *o.ReturnPriceForecasts
}

// GetReturnPriceForecastsOk returns a tuple with the ReturnPriceForecasts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolCollectionForecastsRequest) GetReturnPriceForecastsOk() (*bool, bool) {
	if o == nil || isNil(o.ReturnPriceForecasts) {
		return nil, false
	}
	return o.ReturnPriceForecasts, true
}

// HasReturnPriceForecasts returns a boolean if a field has been set.
func (o *GetSolCollectionForecastsRequest) HasReturnPriceForecasts() bool {
	if o != nil && !isNil(o.ReturnPriceForecasts) {
		return true
	}

	return false
}

// SetReturnPriceForecasts gets a reference to the given bool and assigns it to the ReturnPriceForecasts field.
func (o *GetSolCollectionForecastsRequest) SetReturnPriceForecasts(v bool) {
	o.ReturnPriceForecasts = &v
}

// GetAlpha returns the Alpha field value if set, zero value otherwise.
func (o *GetSolCollectionForecastsRequest) GetAlpha() float32 {
	if o == nil || isNil(o.Alpha) {
		var ret float32
		return ret
	}
	return *o.Alpha
}

// GetAlphaOk returns a tuple with the Alpha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolCollectionForecastsRequest) GetAlphaOk() (*float32, bool) {
	if o == nil || isNil(o.Alpha) {
		return nil, false
	}
	return o.Alpha, true
}

// HasAlpha returns a boolean if a field has been set.
func (o *GetSolCollectionForecastsRequest) HasAlpha() bool {
	if o != nil && !isNil(o.Alpha) {
		return true
	}

	return false
}

// SetAlpha gets a reference to the given float32 and assigns it to the Alpha field.
func (o *GetSolCollectionForecastsRequest) SetAlpha(v float32) {
	o.Alpha = &v
}

// GetReptCurr returns the ReptCurr field value if set, zero value otherwise.
func (o *GetSolCollectionForecastsRequest) GetReptCurr() string {
	if o == nil || isNil(o.ReptCurr) {
		var ret string
		return ret
	}
	return *o.ReptCurr
}

// GetReptCurrOk returns a tuple with the ReptCurr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolCollectionForecastsRequest) GetReptCurrOk() (*string, bool) {
	if o == nil || isNil(o.ReptCurr) {
		return nil, false
	}
	return o.ReptCurr, true
}

// HasReptCurr returns a boolean if a field has been set.
func (o *GetSolCollectionForecastsRequest) HasReptCurr() bool {
	if o != nil && !isNil(o.ReptCurr) {
		return true
	}

	return false
}

// SetReptCurr gets a reference to the given string and assigns it to the ReptCurr field.
func (o *GetSolCollectionForecastsRequest) SetReptCurr(v string) {
	o.ReptCurr = &v
}

// GetExcludeWash returns the ExcludeWash field value if set, zero value otherwise.
func (o *GetSolCollectionForecastsRequest) GetExcludeWash() bool {
	if o == nil || isNil(o.ExcludeWash) {
		var ret bool
		return ret
	}
	return *o.ExcludeWash
}

// GetExcludeWashOk returns a tuple with the ExcludeWash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolCollectionForecastsRequest) GetExcludeWashOk() (*bool, bool) {
	if o == nil || isNil(o.ExcludeWash) {
		return nil, false
	}
	return o.ExcludeWash, true
}

// HasExcludeWash returns a boolean if a field has been set.
func (o *GetSolCollectionForecastsRequest) HasExcludeWash() bool {
	if o != nil && !isNil(o.ExcludeWash) {
		return true
	}

	return false
}

// SetExcludeWash gets a reference to the given bool and assigns it to the ExcludeWash field.
func (o *GetSolCollectionForecastsRequest) SetExcludeWash(v bool) {
	o.ExcludeWash = &v
}

// GetArchParams returns the ArchParams field value if set, zero value otherwise.
func (o *GetSolCollectionForecastsRequest) GetArchParams() GetSolCollectionForecastsRequestArchParams {
	if o == nil || isNil(o.ArchParams) {
		var ret GetSolCollectionForecastsRequestArchParams
		return ret
	}
	return *o.ArchParams
}

// GetArchParamsOk returns a tuple with the ArchParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolCollectionForecastsRequest) GetArchParamsOk() (*GetSolCollectionForecastsRequestArchParams, bool) {
	if o == nil || isNil(o.ArchParams) {
		return nil, false
	}
	return o.ArchParams, true
}

// HasArchParams returns a boolean if a field has been set.
func (o *GetSolCollectionForecastsRequest) HasArchParams() bool {
	if o != nil && !isNil(o.ArchParams) {
		return true
	}

	return false
}

// SetArchParams gets a reference to the given GetSolCollectionForecastsRequestArchParams and assigns it to the ArchParams field.
func (o *GetSolCollectionForecastsRequest) SetArchParams(v GetSolCollectionForecastsRequestArchParams) {
	o.ArchParams = &v
}

func (o GetSolCollectionForecastsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSolCollectionForecastsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_tag"] = o.CollectionTag
	if !isNil(o.Percentiles) {
		toSerialize["percentiles"] = o.Percentiles
	}
	if !isNil(o.Voltype) {
		toSerialize["voltype"] = o.Voltype
	}
	if !isNil(o.Horizon) {
		toSerialize["horizon"] = o.Horizon
	}
	if !isNil(o.Frequency) {
		toSerialize["frequency"] = o.Frequency
	}
	if !isNil(o.Dist) {
		toSerialize["dist"] = o.Dist
	}
	if !isNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !isNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !isNil(o.ReturnPriceForecasts) {
		toSerialize["return_price_forecasts"] = o.ReturnPriceForecasts
	}
	if !isNil(o.Alpha) {
		toSerialize["alpha"] = o.Alpha
	}
	if !isNil(o.ReptCurr) {
		toSerialize["rept_curr"] = o.ReptCurr
	}
	if !isNil(o.ExcludeWash) {
		toSerialize["exclude_wash"] = o.ExcludeWash
	}
	if !isNil(o.ArchParams) {
		toSerialize["arch_params"] = o.ArchParams
	}
	return toSerialize, nil
}

type NullableGetSolCollectionForecastsRequest struct {
	value *GetSolCollectionForecastsRequest
	isSet bool
}

func (v NullableGetSolCollectionForecastsRequest) Get() *GetSolCollectionForecastsRequest {
	return v.value
}

func (v *NullableGetSolCollectionForecastsRequest) Set(val *GetSolCollectionForecastsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolCollectionForecastsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolCollectionForecastsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolCollectionForecastsRequest(val *GetSolCollectionForecastsRequest) *NullableGetSolCollectionForecastsRequest {
	return &NullableGetSolCollectionForecastsRequest{value: val, isSet: true}
}

func (v NullableGetSolCollectionForecastsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSolCollectionForecastsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

