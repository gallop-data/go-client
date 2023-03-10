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

// checks if the GetPolCustomCollectionRiskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPolCustomCollectionRiskRequest{}

// GetPolCustomCollectionRiskRequest struct for GetPolCustomCollectionRiskRequest
type GetPolCustomCollectionRiskRequest struct {
	// The contract address of the token collection.
	CollectionAddress string `json:"collection_address"`
	// The holding period to evaluate risk for, e.g. `12M`
	HoldingPeriod string `json:"holding_period"`
	// The collection percentile(s)
	Percentiles []int32 `json:"percentiles,omitempty"`
	// The amount of tokens in your portfolio
	Amount *int32 `json:"amount,omitempty"`
	// The distribution assumed to calculate parametric risk for
	Dist *string `json:"dist,omitempty"`
	// The start date to pull data for calculations
	StartDate *string `json:"start_date,omitempty"`
	// The end date to pull data for calculations
	EndDate *string `json:"end_date,omitempty"`
	// The rate of return for an asset deemed risk free in the contemplated holding period
	RiskFreeRate *float32 `json:"risk_free_rate,omitempty"`
	// Whether to winsorize time series outliers prior to calculating risk
	WinsOutliers *bool `json:"wins_outliers,omitempty"`
	// The interval at which to calculate returns to base the forecasts upon, e.g. `1D` for daily, `1M` for monthly etc.
	Frequency *string `json:"frequency,omitempty"`
	// The currency to report results in
	ReptCurr *string `json:"rept_curr,omitempty"`
	// Exclude suspected wash transactions?
	ExcludeWash *bool `json:"exclude_wash,omitempty"`
	// If true, report drawdown volatility (based on negative returns only).
	Drawdown *bool `json:"drawdown,omitempty"`
	ArcParams *GetEthCustomCollectionRiskRequestArcParams `json:"arc_params,omitempty"`
	GarParams *GetEthCustomCollectionRiskRequestGarParams `json:"gar_params,omitempty"`
	HarParams *GetEthCustomCollectionRiskRequestHarParams `json:"har_params,omitempty"`
}

// NewGetPolCustomCollectionRiskRequest instantiates a new GetPolCustomCollectionRiskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPolCustomCollectionRiskRequest(collectionAddress string, holdingPeriod string) *GetPolCustomCollectionRiskRequest {
	this := GetPolCustomCollectionRiskRequest{}
	this.CollectionAddress = collectionAddress
	this.HoldingPeriod = holdingPeriod
	return &this
}

// NewGetPolCustomCollectionRiskRequestWithDefaults instantiates a new GetPolCustomCollectionRiskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPolCustomCollectionRiskRequestWithDefaults() *GetPolCustomCollectionRiskRequest {
	this := GetPolCustomCollectionRiskRequest{}
	return &this
}

// GetCollectionAddress returns the CollectionAddress field value
func (o *GetPolCustomCollectionRiskRequest) GetCollectionAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value
// and a boolean to check if the value has been set.
func (o *GetPolCustomCollectionRiskRequest) GetCollectionAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionAddress, true
}

// SetCollectionAddress sets field value
func (o *GetPolCustomCollectionRiskRequest) SetCollectionAddress(v string) {
	o.CollectionAddress = v
}

// GetHoldingPeriod returns the HoldingPeriod field value
func (o *GetPolCustomCollectionRiskRequest) GetHoldingPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HoldingPeriod
}

// GetHoldingPeriodOk returns a tuple with the HoldingPeriod field value
// and a boolean to check if the value has been set.
func (o *GetPolCustomCollectionRiskRequest) GetHoldingPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HoldingPeriod, true
}

// SetHoldingPeriod sets field value
func (o *GetPolCustomCollectionRiskRequest) SetHoldingPeriod(v string) {
	o.HoldingPeriod = v
}

// GetPercentiles returns the Percentiles field value if set, zero value otherwise.
func (o *GetPolCustomCollectionRiskRequest) GetPercentiles() []int32 {
	if o == nil || IsNil(o.Percentiles) {
		var ret []int32
		return ret
	}
	return o.Percentiles
}

// GetPercentilesOk returns a tuple with the Percentiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCustomCollectionRiskRequest) GetPercentilesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Percentiles) {
		return nil, false
	}
	return o.Percentiles, true
}

// HasPercentiles returns a boolean if a field has been set.
func (o *GetPolCustomCollectionRiskRequest) HasPercentiles() bool {
	if o != nil && !IsNil(o.Percentiles) {
		return true
	}

	return false
}

// SetPercentiles gets a reference to the given []int32 and assigns it to the Percentiles field.
func (o *GetPolCustomCollectionRiskRequest) SetPercentiles(v []int32) {
	o.Percentiles = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetPolCustomCollectionRiskRequest) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCustomCollectionRiskRequest) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetPolCustomCollectionRiskRequest) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *GetPolCustomCollectionRiskRequest) SetAmount(v int32) {
	o.Amount = &v
}

// GetDist returns the Dist field value if set, zero value otherwise.
func (o *GetPolCustomCollectionRiskRequest) GetDist() string {
	if o == nil || IsNil(o.Dist) {
		var ret string
		return ret
	}
	return *o.Dist
}

// GetDistOk returns a tuple with the Dist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCustomCollectionRiskRequest) GetDistOk() (*string, bool) {
	if o == nil || IsNil(o.Dist) {
		return nil, false
	}
	return o.Dist, true
}

// HasDist returns a boolean if a field has been set.
func (o *GetPolCustomCollectionRiskRequest) HasDist() bool {
	if o != nil && !IsNil(o.Dist) {
		return true
	}

	return false
}

// SetDist gets a reference to the given string and assigns it to the Dist field.
func (o *GetPolCustomCollectionRiskRequest) SetDist(v string) {
	o.Dist = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *GetPolCustomCollectionRiskRequest) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCustomCollectionRiskRequest) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *GetPolCustomCollectionRiskRequest) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *GetPolCustomCollectionRiskRequest) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GetPolCustomCollectionRiskRequest) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCustomCollectionRiskRequest) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GetPolCustomCollectionRiskRequest) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *GetPolCustomCollectionRiskRequest) SetEndDate(v string) {
	o.EndDate = &v
}

// GetRiskFreeRate returns the RiskFreeRate field value if set, zero value otherwise.
func (o *GetPolCustomCollectionRiskRequest) GetRiskFreeRate() float32 {
	if o == nil || IsNil(o.RiskFreeRate) {
		var ret float32
		return ret
	}
	return *o.RiskFreeRate
}

// GetRiskFreeRateOk returns a tuple with the RiskFreeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCustomCollectionRiskRequest) GetRiskFreeRateOk() (*float32, bool) {
	if o == nil || IsNil(o.RiskFreeRate) {
		return nil, false
	}
	return o.RiskFreeRate, true
}

// HasRiskFreeRate returns a boolean if a field has been set.
func (o *GetPolCustomCollectionRiskRequest) HasRiskFreeRate() bool {
	if o != nil && !IsNil(o.RiskFreeRate) {
		return true
	}

	return false
}

// SetRiskFreeRate gets a reference to the given float32 and assigns it to the RiskFreeRate field.
func (o *GetPolCustomCollectionRiskRequest) SetRiskFreeRate(v float32) {
	o.RiskFreeRate = &v
}

// GetWinsOutliers returns the WinsOutliers field value if set, zero value otherwise.
func (o *GetPolCustomCollectionRiskRequest) GetWinsOutliers() bool {
	if o == nil || IsNil(o.WinsOutliers) {
		var ret bool
		return ret
	}
	return *o.WinsOutliers
}

// GetWinsOutliersOk returns a tuple with the WinsOutliers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCustomCollectionRiskRequest) GetWinsOutliersOk() (*bool, bool) {
	if o == nil || IsNil(o.WinsOutliers) {
		return nil, false
	}
	return o.WinsOutliers, true
}

// HasWinsOutliers returns a boolean if a field has been set.
func (o *GetPolCustomCollectionRiskRequest) HasWinsOutliers() bool {
	if o != nil && !IsNil(o.WinsOutliers) {
		return true
	}

	return false
}

// SetWinsOutliers gets a reference to the given bool and assigns it to the WinsOutliers field.
func (o *GetPolCustomCollectionRiskRequest) SetWinsOutliers(v bool) {
	o.WinsOutliers = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *GetPolCustomCollectionRiskRequest) GetFrequency() string {
	if o == nil || IsNil(o.Frequency) {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCustomCollectionRiskRequest) GetFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *GetPolCustomCollectionRiskRequest) HasFrequency() bool {
	if o != nil && !IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *GetPolCustomCollectionRiskRequest) SetFrequency(v string) {
	o.Frequency = &v
}

// GetReptCurr returns the ReptCurr field value if set, zero value otherwise.
func (o *GetPolCustomCollectionRiskRequest) GetReptCurr() string {
	if o == nil || IsNil(o.ReptCurr) {
		var ret string
		return ret
	}
	return *o.ReptCurr
}

// GetReptCurrOk returns a tuple with the ReptCurr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCustomCollectionRiskRequest) GetReptCurrOk() (*string, bool) {
	if o == nil || IsNil(o.ReptCurr) {
		return nil, false
	}
	return o.ReptCurr, true
}

// HasReptCurr returns a boolean if a field has been set.
func (o *GetPolCustomCollectionRiskRequest) HasReptCurr() bool {
	if o != nil && !IsNil(o.ReptCurr) {
		return true
	}

	return false
}

// SetReptCurr gets a reference to the given string and assigns it to the ReptCurr field.
func (o *GetPolCustomCollectionRiskRequest) SetReptCurr(v string) {
	o.ReptCurr = &v
}

// GetExcludeWash returns the ExcludeWash field value if set, zero value otherwise.
func (o *GetPolCustomCollectionRiskRequest) GetExcludeWash() bool {
	if o == nil || IsNil(o.ExcludeWash) {
		var ret bool
		return ret
	}
	return *o.ExcludeWash
}

// GetExcludeWashOk returns a tuple with the ExcludeWash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCustomCollectionRiskRequest) GetExcludeWashOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeWash) {
		return nil, false
	}
	return o.ExcludeWash, true
}

// HasExcludeWash returns a boolean if a field has been set.
func (o *GetPolCustomCollectionRiskRequest) HasExcludeWash() bool {
	if o != nil && !IsNil(o.ExcludeWash) {
		return true
	}

	return false
}

// SetExcludeWash gets a reference to the given bool and assigns it to the ExcludeWash field.
func (o *GetPolCustomCollectionRiskRequest) SetExcludeWash(v bool) {
	o.ExcludeWash = &v
}

// GetDrawdown returns the Drawdown field value if set, zero value otherwise.
func (o *GetPolCustomCollectionRiskRequest) GetDrawdown() bool {
	if o == nil || IsNil(o.Drawdown) {
		var ret bool
		return ret
	}
	return *o.Drawdown
}

// GetDrawdownOk returns a tuple with the Drawdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCustomCollectionRiskRequest) GetDrawdownOk() (*bool, bool) {
	if o == nil || IsNil(o.Drawdown) {
		return nil, false
	}
	return o.Drawdown, true
}

// HasDrawdown returns a boolean if a field has been set.
func (o *GetPolCustomCollectionRiskRequest) HasDrawdown() bool {
	if o != nil && !IsNil(o.Drawdown) {
		return true
	}

	return false
}

// SetDrawdown gets a reference to the given bool and assigns it to the Drawdown field.
func (o *GetPolCustomCollectionRiskRequest) SetDrawdown(v bool) {
	o.Drawdown = &v
}

// GetArcParams returns the ArcParams field value if set, zero value otherwise.
func (o *GetPolCustomCollectionRiskRequest) GetArcParams() GetEthCustomCollectionRiskRequestArcParams {
	if o == nil || IsNil(o.ArcParams) {
		var ret GetEthCustomCollectionRiskRequestArcParams
		return ret
	}
	return *o.ArcParams
}

// GetArcParamsOk returns a tuple with the ArcParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCustomCollectionRiskRequest) GetArcParamsOk() (*GetEthCustomCollectionRiskRequestArcParams, bool) {
	if o == nil || IsNil(o.ArcParams) {
		return nil, false
	}
	return o.ArcParams, true
}

// HasArcParams returns a boolean if a field has been set.
func (o *GetPolCustomCollectionRiskRequest) HasArcParams() bool {
	if o != nil && !IsNil(o.ArcParams) {
		return true
	}

	return false
}

// SetArcParams gets a reference to the given GetEthCustomCollectionRiskRequestArcParams and assigns it to the ArcParams field.
func (o *GetPolCustomCollectionRiskRequest) SetArcParams(v GetEthCustomCollectionRiskRequestArcParams) {
	o.ArcParams = &v
}

// GetGarParams returns the GarParams field value if set, zero value otherwise.
func (o *GetPolCustomCollectionRiskRequest) GetGarParams() GetEthCustomCollectionRiskRequestGarParams {
	if o == nil || IsNil(o.GarParams) {
		var ret GetEthCustomCollectionRiskRequestGarParams
		return ret
	}
	return *o.GarParams
}

// GetGarParamsOk returns a tuple with the GarParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCustomCollectionRiskRequest) GetGarParamsOk() (*GetEthCustomCollectionRiskRequestGarParams, bool) {
	if o == nil || IsNil(o.GarParams) {
		return nil, false
	}
	return o.GarParams, true
}

// HasGarParams returns a boolean if a field has been set.
func (o *GetPolCustomCollectionRiskRequest) HasGarParams() bool {
	if o != nil && !IsNil(o.GarParams) {
		return true
	}

	return false
}

// SetGarParams gets a reference to the given GetEthCustomCollectionRiskRequestGarParams and assigns it to the GarParams field.
func (o *GetPolCustomCollectionRiskRequest) SetGarParams(v GetEthCustomCollectionRiskRequestGarParams) {
	o.GarParams = &v
}

// GetHarParams returns the HarParams field value if set, zero value otherwise.
func (o *GetPolCustomCollectionRiskRequest) GetHarParams() GetEthCustomCollectionRiskRequestHarParams {
	if o == nil || IsNil(o.HarParams) {
		var ret GetEthCustomCollectionRiskRequestHarParams
		return ret
	}
	return *o.HarParams
}

// GetHarParamsOk returns a tuple with the HarParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolCustomCollectionRiskRequest) GetHarParamsOk() (*GetEthCustomCollectionRiskRequestHarParams, bool) {
	if o == nil || IsNil(o.HarParams) {
		return nil, false
	}
	return o.HarParams, true
}

// HasHarParams returns a boolean if a field has been set.
func (o *GetPolCustomCollectionRiskRequest) HasHarParams() bool {
	if o != nil && !IsNil(o.HarParams) {
		return true
	}

	return false
}

// SetHarParams gets a reference to the given GetEthCustomCollectionRiskRequestHarParams and assigns it to the HarParams field.
func (o *GetPolCustomCollectionRiskRequest) SetHarParams(v GetEthCustomCollectionRiskRequestHarParams) {
	o.HarParams = &v
}

func (o GetPolCustomCollectionRiskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPolCustomCollectionRiskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_address"] = o.CollectionAddress
	toSerialize["holding_period"] = o.HoldingPeriod
	if !IsNil(o.Percentiles) {
		toSerialize["percentiles"] = o.Percentiles
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
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
	if !IsNil(o.RiskFreeRate) {
		toSerialize["risk_free_rate"] = o.RiskFreeRate
	}
	if !IsNil(o.WinsOutliers) {
		toSerialize["wins_outliers"] = o.WinsOutliers
	}
	if !IsNil(o.Frequency) {
		toSerialize["frequency"] = o.Frequency
	}
	if !IsNil(o.ReptCurr) {
		toSerialize["rept_curr"] = o.ReptCurr
	}
	if !IsNil(o.ExcludeWash) {
		toSerialize["exclude_wash"] = o.ExcludeWash
	}
	if !IsNil(o.Drawdown) {
		toSerialize["drawdown"] = o.Drawdown
	}
	if !IsNil(o.ArcParams) {
		toSerialize["arc_params"] = o.ArcParams
	}
	if !IsNil(o.GarParams) {
		toSerialize["gar_params"] = o.GarParams
	}
	if !IsNil(o.HarParams) {
		toSerialize["har_params"] = o.HarParams
	}
	return toSerialize, nil
}

type NullableGetPolCustomCollectionRiskRequest struct {
	value *GetPolCustomCollectionRiskRequest
	isSet bool
}

func (v NullableGetPolCustomCollectionRiskRequest) Get() *GetPolCustomCollectionRiskRequest {
	return v.value
}

func (v *NullableGetPolCustomCollectionRiskRequest) Set(val *GetPolCustomCollectionRiskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPolCustomCollectionRiskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPolCustomCollectionRiskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPolCustomCollectionRiskRequest(val *GetPolCustomCollectionRiskRequest) *NullableGetPolCustomCollectionRiskRequest {
	return &NullableGetPolCustomCollectionRiskRequest{value: val, isSet: true}
}

func (v NullableGetPolCustomCollectionRiskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPolCustomCollectionRiskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


