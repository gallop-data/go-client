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

// checks if the GetEthCollectionForecastsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEthCollectionForecastsRequest{}

// GetEthCollectionForecastsRequest struct for GetEthCollectionForecastsRequest
type GetEthCollectionForecastsRequest struct {
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

// NewGetEthCollectionForecastsRequest instantiates a new GetEthCollectionForecastsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEthCollectionForecastsRequest(collectionAddress string) *GetEthCollectionForecastsRequest {
	this := GetEthCollectionForecastsRequest{}
	this.CollectionAddress = collectionAddress
	return &this
}

// NewGetEthCollectionForecastsRequestWithDefaults instantiates a new GetEthCollectionForecastsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEthCollectionForecastsRequestWithDefaults() *GetEthCollectionForecastsRequest {
	this := GetEthCollectionForecastsRequest{}
	return &this
}

// GetCollectionAddress returns the CollectionAddress field value
func (o *GetEthCollectionForecastsRequest) GetCollectionAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value
// and a boolean to check if the value has been set.
func (o *GetEthCollectionForecastsRequest) GetCollectionAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionAddress, true
}

// SetCollectionAddress sets field value
func (o *GetEthCollectionForecastsRequest) SetCollectionAddress(v string) {
	o.CollectionAddress = v
}

// GetPercentiles returns the Percentiles field value if set, zero value otherwise.
func (o *GetEthCollectionForecastsRequest) GetPercentiles() []int32 {
	if o == nil || isNil(o.Percentiles) {
		var ret []int32
		return ret
	}
	return o.Percentiles
}

// GetPercentilesOk returns a tuple with the Percentiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionForecastsRequest) GetPercentilesOk() ([]int32, bool) {
	if o == nil || isNil(o.Percentiles) {
		return nil, false
	}
	return o.Percentiles, true
}

// HasPercentiles returns a boolean if a field has been set.
func (o *GetEthCollectionForecastsRequest) HasPercentiles() bool {
	if o != nil && !isNil(o.Percentiles) {
		return true
	}

	return false
}

// SetPercentiles gets a reference to the given []int32 and assigns it to the Percentiles field.
func (o *GetEthCollectionForecastsRequest) SetPercentiles(v []int32) {
	o.Percentiles = v
}

// GetVoltype returns the Voltype field value if set, zero value otherwise.
func (o *GetEthCollectionForecastsRequest) GetVoltype() string {
	if o == nil || isNil(o.Voltype) {
		var ret string
		return ret
	}
	return *o.Voltype
}

// GetVoltypeOk returns a tuple with the Voltype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionForecastsRequest) GetVoltypeOk() (*string, bool) {
	if o == nil || isNil(o.Voltype) {
		return nil, false
	}
	return o.Voltype, true
}

// HasVoltype returns a boolean if a field has been set.
func (o *GetEthCollectionForecastsRequest) HasVoltype() bool {
	if o != nil && !isNil(o.Voltype) {
		return true
	}

	return false
}

// SetVoltype gets a reference to the given string and assigns it to the Voltype field.
func (o *GetEthCollectionForecastsRequest) SetVoltype(v string) {
	o.Voltype = &v
}

// GetHorizon returns the Horizon field value if set, zero value otherwise.
func (o *GetEthCollectionForecastsRequest) GetHorizon() int32 {
	if o == nil || isNil(o.Horizon) {
		var ret int32
		return ret
	}
	return *o.Horizon
}

// GetHorizonOk returns a tuple with the Horizon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionForecastsRequest) GetHorizonOk() (*int32, bool) {
	if o == nil || isNil(o.Horizon) {
		return nil, false
	}
	return o.Horizon, true
}

// HasHorizon returns a boolean if a field has been set.
func (o *GetEthCollectionForecastsRequest) HasHorizon() bool {
	if o != nil && !isNil(o.Horizon) {
		return true
	}

	return false
}

// SetHorizon gets a reference to the given int32 and assigns it to the Horizon field.
func (o *GetEthCollectionForecastsRequest) SetHorizon(v int32) {
	o.Horizon = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *GetEthCollectionForecastsRequest) GetFrequency() string {
	if o == nil || isNil(o.Frequency) {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionForecastsRequest) GetFrequencyOk() (*string, bool) {
	if o == nil || isNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *GetEthCollectionForecastsRequest) HasFrequency() bool {
	if o != nil && !isNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *GetEthCollectionForecastsRequest) SetFrequency(v string) {
	o.Frequency = &v
}

// GetDist returns the Dist field value if set, zero value otherwise.
func (o *GetEthCollectionForecastsRequest) GetDist() string {
	if o == nil || isNil(o.Dist) {
		var ret string
		return ret
	}
	return *o.Dist
}

// GetDistOk returns a tuple with the Dist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionForecastsRequest) GetDistOk() (*string, bool) {
	if o == nil || isNil(o.Dist) {
		return nil, false
	}
	return o.Dist, true
}

// HasDist returns a boolean if a field has been set.
func (o *GetEthCollectionForecastsRequest) HasDist() bool {
	if o != nil && !isNil(o.Dist) {
		return true
	}

	return false
}

// SetDist gets a reference to the given string and assigns it to the Dist field.
func (o *GetEthCollectionForecastsRequest) SetDist(v string) {
	o.Dist = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *GetEthCollectionForecastsRequest) GetStartDate() string {
	if o == nil || isNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionForecastsRequest) GetStartDateOk() (*string, bool) {
	if o == nil || isNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *GetEthCollectionForecastsRequest) HasStartDate() bool {
	if o != nil && !isNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *GetEthCollectionForecastsRequest) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GetEthCollectionForecastsRequest) GetEndDate() string {
	if o == nil || isNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionForecastsRequest) GetEndDateOk() (*string, bool) {
	if o == nil || isNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GetEthCollectionForecastsRequest) HasEndDate() bool {
	if o != nil && !isNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *GetEthCollectionForecastsRequest) SetEndDate(v string) {
	o.EndDate = &v
}

// GetReturnPriceForecasts returns the ReturnPriceForecasts field value if set, zero value otherwise.
func (o *GetEthCollectionForecastsRequest) GetReturnPriceForecasts() bool {
	if o == nil || isNil(o.ReturnPriceForecasts) {
		var ret bool
		return ret
	}
	return *o.ReturnPriceForecasts
}

// GetReturnPriceForecastsOk returns a tuple with the ReturnPriceForecasts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionForecastsRequest) GetReturnPriceForecastsOk() (*bool, bool) {
	if o == nil || isNil(o.ReturnPriceForecasts) {
		return nil, false
	}
	return o.ReturnPriceForecasts, true
}

// HasReturnPriceForecasts returns a boolean if a field has been set.
func (o *GetEthCollectionForecastsRequest) HasReturnPriceForecasts() bool {
	if o != nil && !isNil(o.ReturnPriceForecasts) {
		return true
	}

	return false
}

// SetReturnPriceForecasts gets a reference to the given bool and assigns it to the ReturnPriceForecasts field.
func (o *GetEthCollectionForecastsRequest) SetReturnPriceForecasts(v bool) {
	o.ReturnPriceForecasts = &v
}

// GetAlpha returns the Alpha field value if set, zero value otherwise.
func (o *GetEthCollectionForecastsRequest) GetAlpha() float32 {
	if o == nil || isNil(o.Alpha) {
		var ret float32
		return ret
	}
	return *o.Alpha
}

// GetAlphaOk returns a tuple with the Alpha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionForecastsRequest) GetAlphaOk() (*float32, bool) {
	if o == nil || isNil(o.Alpha) {
		return nil, false
	}
	return o.Alpha, true
}

// HasAlpha returns a boolean if a field has been set.
func (o *GetEthCollectionForecastsRequest) HasAlpha() bool {
	if o != nil && !isNil(o.Alpha) {
		return true
	}

	return false
}

// SetAlpha gets a reference to the given float32 and assigns it to the Alpha field.
func (o *GetEthCollectionForecastsRequest) SetAlpha(v float32) {
	o.Alpha = &v
}

// GetReptCurr returns the ReptCurr field value if set, zero value otherwise.
func (o *GetEthCollectionForecastsRequest) GetReptCurr() string {
	if o == nil || isNil(o.ReptCurr) {
		var ret string
		return ret
	}
	return *o.ReptCurr
}

// GetReptCurrOk returns a tuple with the ReptCurr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionForecastsRequest) GetReptCurrOk() (*string, bool) {
	if o == nil || isNil(o.ReptCurr) {
		return nil, false
	}
	return o.ReptCurr, true
}

// HasReptCurr returns a boolean if a field has been set.
func (o *GetEthCollectionForecastsRequest) HasReptCurr() bool {
	if o != nil && !isNil(o.ReptCurr) {
		return true
	}

	return false
}

// SetReptCurr gets a reference to the given string and assigns it to the ReptCurr field.
func (o *GetEthCollectionForecastsRequest) SetReptCurr(v string) {
	o.ReptCurr = &v
}

// GetExcludeWash returns the ExcludeWash field value if set, zero value otherwise.
func (o *GetEthCollectionForecastsRequest) GetExcludeWash() bool {
	if o == nil || isNil(o.ExcludeWash) {
		var ret bool
		return ret
	}
	return *o.ExcludeWash
}

// GetExcludeWashOk returns a tuple with the ExcludeWash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionForecastsRequest) GetExcludeWashOk() (*bool, bool) {
	if o == nil || isNil(o.ExcludeWash) {
		return nil, false
	}
	return o.ExcludeWash, true
}

// HasExcludeWash returns a boolean if a field has been set.
func (o *GetEthCollectionForecastsRequest) HasExcludeWash() bool {
	if o != nil && !isNil(o.ExcludeWash) {
		return true
	}

	return false
}

// SetExcludeWash gets a reference to the given bool and assigns it to the ExcludeWash field.
func (o *GetEthCollectionForecastsRequest) SetExcludeWash(v bool) {
	o.ExcludeWash = &v
}

// GetArchParams returns the ArchParams field value if set, zero value otherwise.
func (o *GetEthCollectionForecastsRequest) GetArchParams() GetEthCollectionForecastsRequestArchParams {
	if o == nil || isNil(o.ArchParams) {
		var ret GetEthCollectionForecastsRequestArchParams
		return ret
	}
	return *o.ArchParams
}

// GetArchParamsOk returns a tuple with the ArchParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCollectionForecastsRequest) GetArchParamsOk() (*GetEthCollectionForecastsRequestArchParams, bool) {
	if o == nil || isNil(o.ArchParams) {
		return nil, false
	}
	return o.ArchParams, true
}

// HasArchParams returns a boolean if a field has been set.
func (o *GetEthCollectionForecastsRequest) HasArchParams() bool {
	if o != nil && !isNil(o.ArchParams) {
		return true
	}

	return false
}

// SetArchParams gets a reference to the given GetEthCollectionForecastsRequestArchParams and assigns it to the ArchParams field.
func (o *GetEthCollectionForecastsRequest) SetArchParams(v GetEthCollectionForecastsRequestArchParams) {
	o.ArchParams = &v
}

func (o GetEthCollectionForecastsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEthCollectionForecastsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_address"] = o.CollectionAddress
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

type NullableGetEthCollectionForecastsRequest struct {
	value *GetEthCollectionForecastsRequest
	isSet bool
}

func (v NullableGetEthCollectionForecastsRequest) Get() *GetEthCollectionForecastsRequest {
	return v.value
}

func (v *NullableGetEthCollectionForecastsRequest) Set(val *GetEthCollectionForecastsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEthCollectionForecastsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEthCollectionForecastsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEthCollectionForecastsRequest(val *GetEthCollectionForecastsRequest) *NullableGetEthCollectionForecastsRequest {
	return &NullableGetEthCollectionForecastsRequest{value: val, isSet: true}
}

func (v NullableGetEthCollectionForecastsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEthCollectionForecastsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


