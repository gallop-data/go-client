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

// checks if the GetEthCustomTokenRiskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEthCustomTokenRiskRequest{}

// GetEthCustomTokenRiskRequest struct for GetEthCustomTokenRiskRequest
type GetEthCustomTokenRiskRequest struct {
	// The contract address of the token collection.
	CollectionAddress string `json:"collection_address"`
	// The id(s) for the token(s).
	TokenId []string `json:"token_id"`
	// The holding period to evaluate risk for, e.g. `12M`
	HoldingPeriod string `json:"holding_period"`
	// Return distribution assumed.
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

// NewGetEthCustomTokenRiskRequest instantiates a new GetEthCustomTokenRiskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEthCustomTokenRiskRequest(collectionAddress string, tokenId []string, holdingPeriod string) *GetEthCustomTokenRiskRequest {
	this := GetEthCustomTokenRiskRequest{}
	this.CollectionAddress = collectionAddress
	this.TokenId = tokenId
	this.HoldingPeriod = holdingPeriod
	return &this
}

// NewGetEthCustomTokenRiskRequestWithDefaults instantiates a new GetEthCustomTokenRiskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEthCustomTokenRiskRequestWithDefaults() *GetEthCustomTokenRiskRequest {
	this := GetEthCustomTokenRiskRequest{}
	return &this
}

// GetCollectionAddress returns the CollectionAddress field value
func (o *GetEthCustomTokenRiskRequest) GetCollectionAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value
// and a boolean to check if the value has been set.
func (o *GetEthCustomTokenRiskRequest) GetCollectionAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionAddress, true
}

// SetCollectionAddress sets field value
func (o *GetEthCustomTokenRiskRequest) SetCollectionAddress(v string) {
	o.CollectionAddress = v
}

// GetTokenId returns the TokenId field value
func (o *GetEthCustomTokenRiskRequest) GetTokenId() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *GetEthCustomTokenRiskRequest) GetTokenIdOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenId, true
}

// SetTokenId sets field value
func (o *GetEthCustomTokenRiskRequest) SetTokenId(v []string) {
	o.TokenId = v
}

// GetHoldingPeriod returns the HoldingPeriod field value
func (o *GetEthCustomTokenRiskRequest) GetHoldingPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HoldingPeriod
}

// GetHoldingPeriodOk returns a tuple with the HoldingPeriod field value
// and a boolean to check if the value has been set.
func (o *GetEthCustomTokenRiskRequest) GetHoldingPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HoldingPeriod, true
}

// SetHoldingPeriod sets field value
func (o *GetEthCustomTokenRiskRequest) SetHoldingPeriod(v string) {
	o.HoldingPeriod = v
}

// GetDist returns the Dist field value if set, zero value otherwise.
func (o *GetEthCustomTokenRiskRequest) GetDist() string {
	if o == nil || isNil(o.Dist) {
		var ret string
		return ret
	}
	return *o.Dist
}

// GetDistOk returns a tuple with the Dist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomTokenRiskRequest) GetDistOk() (*string, bool) {
	if o == nil || isNil(o.Dist) {
		return nil, false
	}
	return o.Dist, true
}

// HasDist returns a boolean if a field has been set.
func (o *GetEthCustomTokenRiskRequest) HasDist() bool {
	if o != nil && !isNil(o.Dist) {
		return true
	}

	return false
}

// SetDist gets a reference to the given string and assigns it to the Dist field.
func (o *GetEthCustomTokenRiskRequest) SetDist(v string) {
	o.Dist = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *GetEthCustomTokenRiskRequest) GetStartDate() string {
	if o == nil || isNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomTokenRiskRequest) GetStartDateOk() (*string, bool) {
	if o == nil || isNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *GetEthCustomTokenRiskRequest) HasStartDate() bool {
	if o != nil && !isNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *GetEthCustomTokenRiskRequest) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GetEthCustomTokenRiskRequest) GetEndDate() string {
	if o == nil || isNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomTokenRiskRequest) GetEndDateOk() (*string, bool) {
	if o == nil || isNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GetEthCustomTokenRiskRequest) HasEndDate() bool {
	if o != nil && !isNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *GetEthCustomTokenRiskRequest) SetEndDate(v string) {
	o.EndDate = &v
}

// GetRiskFreeRate returns the RiskFreeRate field value if set, zero value otherwise.
func (o *GetEthCustomTokenRiskRequest) GetRiskFreeRate() float32 {
	if o == nil || isNil(o.RiskFreeRate) {
		var ret float32
		return ret
	}
	return *o.RiskFreeRate
}

// GetRiskFreeRateOk returns a tuple with the RiskFreeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomTokenRiskRequest) GetRiskFreeRateOk() (*float32, bool) {
	if o == nil || isNil(o.RiskFreeRate) {
		return nil, false
	}
	return o.RiskFreeRate, true
}

// HasRiskFreeRate returns a boolean if a field has been set.
func (o *GetEthCustomTokenRiskRequest) HasRiskFreeRate() bool {
	if o != nil && !isNil(o.RiskFreeRate) {
		return true
	}

	return false
}

// SetRiskFreeRate gets a reference to the given float32 and assigns it to the RiskFreeRate field.
func (o *GetEthCustomTokenRiskRequest) SetRiskFreeRate(v float32) {
	o.RiskFreeRate = &v
}

// GetWinsOutliers returns the WinsOutliers field value if set, zero value otherwise.
func (o *GetEthCustomTokenRiskRequest) GetWinsOutliers() bool {
	if o == nil || isNil(o.WinsOutliers) {
		var ret bool
		return ret
	}
	return *o.WinsOutliers
}

// GetWinsOutliersOk returns a tuple with the WinsOutliers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomTokenRiskRequest) GetWinsOutliersOk() (*bool, bool) {
	if o == nil || isNil(o.WinsOutliers) {
		return nil, false
	}
	return o.WinsOutliers, true
}

// HasWinsOutliers returns a boolean if a field has been set.
func (o *GetEthCustomTokenRiskRequest) HasWinsOutliers() bool {
	if o != nil && !isNil(o.WinsOutliers) {
		return true
	}

	return false
}

// SetWinsOutliers gets a reference to the given bool and assigns it to the WinsOutliers field.
func (o *GetEthCustomTokenRiskRequest) SetWinsOutliers(v bool) {
	o.WinsOutliers = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *GetEthCustomTokenRiskRequest) GetFrequency() string {
	if o == nil || isNil(o.Frequency) {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomTokenRiskRequest) GetFrequencyOk() (*string, bool) {
	if o == nil || isNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *GetEthCustomTokenRiskRequest) HasFrequency() bool {
	if o != nil && !isNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *GetEthCustomTokenRiskRequest) SetFrequency(v string) {
	o.Frequency = &v
}

// GetReptCurr returns the ReptCurr field value if set, zero value otherwise.
func (o *GetEthCustomTokenRiskRequest) GetReptCurr() string {
	if o == nil || isNil(o.ReptCurr) {
		var ret string
		return ret
	}
	return *o.ReptCurr
}

// GetReptCurrOk returns a tuple with the ReptCurr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomTokenRiskRequest) GetReptCurrOk() (*string, bool) {
	if o == nil || isNil(o.ReptCurr) {
		return nil, false
	}
	return o.ReptCurr, true
}

// HasReptCurr returns a boolean if a field has been set.
func (o *GetEthCustomTokenRiskRequest) HasReptCurr() bool {
	if o != nil && !isNil(o.ReptCurr) {
		return true
	}

	return false
}

// SetReptCurr gets a reference to the given string and assigns it to the ReptCurr field.
func (o *GetEthCustomTokenRiskRequest) SetReptCurr(v string) {
	o.ReptCurr = &v
}

// GetExcludeWash returns the ExcludeWash field value if set, zero value otherwise.
func (o *GetEthCustomTokenRiskRequest) GetExcludeWash() bool {
	if o == nil || isNil(o.ExcludeWash) {
		var ret bool
		return ret
	}
	return *o.ExcludeWash
}

// GetExcludeWashOk returns a tuple with the ExcludeWash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomTokenRiskRequest) GetExcludeWashOk() (*bool, bool) {
	if o == nil || isNil(o.ExcludeWash) {
		return nil, false
	}
	return o.ExcludeWash, true
}

// HasExcludeWash returns a boolean if a field has been set.
func (o *GetEthCustomTokenRiskRequest) HasExcludeWash() bool {
	if o != nil && !isNil(o.ExcludeWash) {
		return true
	}

	return false
}

// SetExcludeWash gets a reference to the given bool and assigns it to the ExcludeWash field.
func (o *GetEthCustomTokenRiskRequest) SetExcludeWash(v bool) {
	o.ExcludeWash = &v
}

// GetDrawdown returns the Drawdown field value if set, zero value otherwise.
func (o *GetEthCustomTokenRiskRequest) GetDrawdown() bool {
	if o == nil || isNil(o.Drawdown) {
		var ret bool
		return ret
	}
	return *o.Drawdown
}

// GetDrawdownOk returns a tuple with the Drawdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomTokenRiskRequest) GetDrawdownOk() (*bool, bool) {
	if o == nil || isNil(o.Drawdown) {
		return nil, false
	}
	return o.Drawdown, true
}

// HasDrawdown returns a boolean if a field has been set.
func (o *GetEthCustomTokenRiskRequest) HasDrawdown() bool {
	if o != nil && !isNil(o.Drawdown) {
		return true
	}

	return false
}

// SetDrawdown gets a reference to the given bool and assigns it to the Drawdown field.
func (o *GetEthCustomTokenRiskRequest) SetDrawdown(v bool) {
	o.Drawdown = &v
}

// GetArcParams returns the ArcParams field value if set, zero value otherwise.
func (o *GetEthCustomTokenRiskRequest) GetArcParams() GetEthCustomCollectionRiskRequestArcParams {
	if o == nil || isNil(o.ArcParams) {
		var ret GetEthCustomCollectionRiskRequestArcParams
		return ret
	}
	return *o.ArcParams
}

// GetArcParamsOk returns a tuple with the ArcParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomTokenRiskRequest) GetArcParamsOk() (*GetEthCustomCollectionRiskRequestArcParams, bool) {
	if o == nil || isNil(o.ArcParams) {
		return nil, false
	}
	return o.ArcParams, true
}

// HasArcParams returns a boolean if a field has been set.
func (o *GetEthCustomTokenRiskRequest) HasArcParams() bool {
	if o != nil && !isNil(o.ArcParams) {
		return true
	}

	return false
}

// SetArcParams gets a reference to the given GetEthCustomCollectionRiskRequestArcParams and assigns it to the ArcParams field.
func (o *GetEthCustomTokenRiskRequest) SetArcParams(v GetEthCustomCollectionRiskRequestArcParams) {
	o.ArcParams = &v
}

// GetGarParams returns the GarParams field value if set, zero value otherwise.
func (o *GetEthCustomTokenRiskRequest) GetGarParams() GetEthCustomCollectionRiskRequestGarParams {
	if o == nil || isNil(o.GarParams) {
		var ret GetEthCustomCollectionRiskRequestGarParams
		return ret
	}
	return *o.GarParams
}

// GetGarParamsOk returns a tuple with the GarParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomTokenRiskRequest) GetGarParamsOk() (*GetEthCustomCollectionRiskRequestGarParams, bool) {
	if o == nil || isNil(o.GarParams) {
		return nil, false
	}
	return o.GarParams, true
}

// HasGarParams returns a boolean if a field has been set.
func (o *GetEthCustomTokenRiskRequest) HasGarParams() bool {
	if o != nil && !isNil(o.GarParams) {
		return true
	}

	return false
}

// SetGarParams gets a reference to the given GetEthCustomCollectionRiskRequestGarParams and assigns it to the GarParams field.
func (o *GetEthCustomTokenRiskRequest) SetGarParams(v GetEthCustomCollectionRiskRequestGarParams) {
	o.GarParams = &v
}

// GetHarParams returns the HarParams field value if set, zero value otherwise.
func (o *GetEthCustomTokenRiskRequest) GetHarParams() GetEthCustomCollectionRiskRequestHarParams {
	if o == nil || isNil(o.HarParams) {
		var ret GetEthCustomCollectionRiskRequestHarParams
		return ret
	}
	return *o.HarParams
}

// GetHarParamsOk returns a tuple with the HarParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomTokenRiskRequest) GetHarParamsOk() (*GetEthCustomCollectionRiskRequestHarParams, bool) {
	if o == nil || isNil(o.HarParams) {
		return nil, false
	}
	return o.HarParams, true
}

// HasHarParams returns a boolean if a field has been set.
func (o *GetEthCustomTokenRiskRequest) HasHarParams() bool {
	if o != nil && !isNil(o.HarParams) {
		return true
	}

	return false
}

// SetHarParams gets a reference to the given GetEthCustomCollectionRiskRequestHarParams and assigns it to the HarParams field.
func (o *GetEthCustomTokenRiskRequest) SetHarParams(v GetEthCustomCollectionRiskRequestHarParams) {
	o.HarParams = &v
}

func (o GetEthCustomTokenRiskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEthCustomTokenRiskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_address"] = o.CollectionAddress
	toSerialize["token_id"] = o.TokenId
	toSerialize["holding_period"] = o.HoldingPeriod
	if !isNil(o.Dist) {
		toSerialize["dist"] = o.Dist
	}
	if !isNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !isNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !isNil(o.RiskFreeRate) {
		toSerialize["risk_free_rate"] = o.RiskFreeRate
	}
	if !isNil(o.WinsOutliers) {
		toSerialize["wins_outliers"] = o.WinsOutliers
	}
	if !isNil(o.Frequency) {
		toSerialize["frequency"] = o.Frequency
	}
	if !isNil(o.ReptCurr) {
		toSerialize["rept_curr"] = o.ReptCurr
	}
	if !isNil(o.ExcludeWash) {
		toSerialize["exclude_wash"] = o.ExcludeWash
	}
	if !isNil(o.Drawdown) {
		toSerialize["drawdown"] = o.Drawdown
	}
	if !isNil(o.ArcParams) {
		toSerialize["arc_params"] = o.ArcParams
	}
	if !isNil(o.GarParams) {
		toSerialize["gar_params"] = o.GarParams
	}
	if !isNil(o.HarParams) {
		toSerialize["har_params"] = o.HarParams
	}
	return toSerialize, nil
}

type NullableGetEthCustomTokenRiskRequest struct {
	value *GetEthCustomTokenRiskRequest
	isSet bool
}

func (v NullableGetEthCustomTokenRiskRequest) Get() *GetEthCustomTokenRiskRequest {
	return v.value
}

func (v *NullableGetEthCustomTokenRiskRequest) Set(val *GetEthCustomTokenRiskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEthCustomTokenRiskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEthCustomTokenRiskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEthCustomTokenRiskRequest(val *GetEthCustomTokenRiskRequest) *NullableGetEthCustomTokenRiskRequest {
	return &NullableGetEthCustomTokenRiskRequest{value: val, isSet: true}
}

func (v NullableGetEthCustomTokenRiskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEthCustomTokenRiskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

