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

// checks if the GetSolTokenForecastsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSolTokenForecastsRequest{}

// GetSolTokenForecastsRequest struct for GetSolTokenForecastsRequest
type GetSolTokenForecastsRequest struct {
	// A token mint address or list of token addresses.
	MintAddress []string `json:"mint_address"`
	// The numerical id for the token. Provide either id or mint address.
	TokenId *string `json:"token_id,omitempty"`
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
	ArchParams *GetSolTokenForecastsRequestArchParams `json:"arch_params,omitempty"`
}

// NewGetSolTokenForecastsRequest instantiates a new GetSolTokenForecastsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolTokenForecastsRequest(mintAddress []string) *GetSolTokenForecastsRequest {
	this := GetSolTokenForecastsRequest{}
	this.MintAddress = mintAddress
	return &this
}

// NewGetSolTokenForecastsRequestWithDefaults instantiates a new GetSolTokenForecastsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolTokenForecastsRequestWithDefaults() *GetSolTokenForecastsRequest {
	this := GetSolTokenForecastsRequest{}
	return &this
}

// GetMintAddress returns the MintAddress field value
func (o *GetSolTokenForecastsRequest) GetMintAddress() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MintAddress
}

// GetMintAddressOk returns a tuple with the MintAddress field value
// and a boolean to check if the value has been set.
func (o *GetSolTokenForecastsRequest) GetMintAddressOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MintAddress, true
}

// SetMintAddress sets field value
func (o *GetSolTokenForecastsRequest) SetMintAddress(v []string) {
	o.MintAddress = v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *GetSolTokenForecastsRequest) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolTokenForecastsRequest) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *GetSolTokenForecastsRequest) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *GetSolTokenForecastsRequest) SetTokenId(v string) {
	o.TokenId = &v
}

// GetVoltype returns the Voltype field value if set, zero value otherwise.
func (o *GetSolTokenForecastsRequest) GetVoltype() string {
	if o == nil || IsNil(o.Voltype) {
		var ret string
		return ret
	}
	return *o.Voltype
}

// GetVoltypeOk returns a tuple with the Voltype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolTokenForecastsRequest) GetVoltypeOk() (*string, bool) {
	if o == nil || IsNil(o.Voltype) {
		return nil, false
	}
	return o.Voltype, true
}

// HasVoltype returns a boolean if a field has been set.
func (o *GetSolTokenForecastsRequest) HasVoltype() bool {
	if o != nil && !IsNil(o.Voltype) {
		return true
	}

	return false
}

// SetVoltype gets a reference to the given string and assigns it to the Voltype field.
func (o *GetSolTokenForecastsRequest) SetVoltype(v string) {
	o.Voltype = &v
}

// GetHorizon returns the Horizon field value if set, zero value otherwise.
func (o *GetSolTokenForecastsRequest) GetHorizon() int32 {
	if o == nil || IsNil(o.Horizon) {
		var ret int32
		return ret
	}
	return *o.Horizon
}

// GetHorizonOk returns a tuple with the Horizon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolTokenForecastsRequest) GetHorizonOk() (*int32, bool) {
	if o == nil || IsNil(o.Horizon) {
		return nil, false
	}
	return o.Horizon, true
}

// HasHorizon returns a boolean if a field has been set.
func (o *GetSolTokenForecastsRequest) HasHorizon() bool {
	if o != nil && !IsNil(o.Horizon) {
		return true
	}

	return false
}

// SetHorizon gets a reference to the given int32 and assigns it to the Horizon field.
func (o *GetSolTokenForecastsRequest) SetHorizon(v int32) {
	o.Horizon = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *GetSolTokenForecastsRequest) GetFrequency() string {
	if o == nil || IsNil(o.Frequency) {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolTokenForecastsRequest) GetFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *GetSolTokenForecastsRequest) HasFrequency() bool {
	if o != nil && !IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *GetSolTokenForecastsRequest) SetFrequency(v string) {
	o.Frequency = &v
}

// GetDist returns the Dist field value if set, zero value otherwise.
func (o *GetSolTokenForecastsRequest) GetDist() string {
	if o == nil || IsNil(o.Dist) {
		var ret string
		return ret
	}
	return *o.Dist
}

// GetDistOk returns a tuple with the Dist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolTokenForecastsRequest) GetDistOk() (*string, bool) {
	if o == nil || IsNil(o.Dist) {
		return nil, false
	}
	return o.Dist, true
}

// HasDist returns a boolean if a field has been set.
func (o *GetSolTokenForecastsRequest) HasDist() bool {
	if o != nil && !IsNil(o.Dist) {
		return true
	}

	return false
}

// SetDist gets a reference to the given string and assigns it to the Dist field.
func (o *GetSolTokenForecastsRequest) SetDist(v string) {
	o.Dist = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *GetSolTokenForecastsRequest) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolTokenForecastsRequest) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *GetSolTokenForecastsRequest) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *GetSolTokenForecastsRequest) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GetSolTokenForecastsRequest) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolTokenForecastsRequest) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GetSolTokenForecastsRequest) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *GetSolTokenForecastsRequest) SetEndDate(v string) {
	o.EndDate = &v
}

// GetReturnPriceForecasts returns the ReturnPriceForecasts field value if set, zero value otherwise.
func (o *GetSolTokenForecastsRequest) GetReturnPriceForecasts() bool {
	if o == nil || IsNil(o.ReturnPriceForecasts) {
		var ret bool
		return ret
	}
	return *o.ReturnPriceForecasts
}

// GetReturnPriceForecastsOk returns a tuple with the ReturnPriceForecasts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolTokenForecastsRequest) GetReturnPriceForecastsOk() (*bool, bool) {
	if o == nil || IsNil(o.ReturnPriceForecasts) {
		return nil, false
	}
	return o.ReturnPriceForecasts, true
}

// HasReturnPriceForecasts returns a boolean if a field has been set.
func (o *GetSolTokenForecastsRequest) HasReturnPriceForecasts() bool {
	if o != nil && !IsNil(o.ReturnPriceForecasts) {
		return true
	}

	return false
}

// SetReturnPriceForecasts gets a reference to the given bool and assigns it to the ReturnPriceForecasts field.
func (o *GetSolTokenForecastsRequest) SetReturnPriceForecasts(v bool) {
	o.ReturnPriceForecasts = &v
}

// GetAlpha returns the Alpha field value if set, zero value otherwise.
func (o *GetSolTokenForecastsRequest) GetAlpha() float32 {
	if o == nil || IsNil(o.Alpha) {
		var ret float32
		return ret
	}
	return *o.Alpha
}

// GetAlphaOk returns a tuple with the Alpha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolTokenForecastsRequest) GetAlphaOk() (*float32, bool) {
	if o == nil || IsNil(o.Alpha) {
		return nil, false
	}
	return o.Alpha, true
}

// HasAlpha returns a boolean if a field has been set.
func (o *GetSolTokenForecastsRequest) HasAlpha() bool {
	if o != nil && !IsNil(o.Alpha) {
		return true
	}

	return false
}

// SetAlpha gets a reference to the given float32 and assigns it to the Alpha field.
func (o *GetSolTokenForecastsRequest) SetAlpha(v float32) {
	o.Alpha = &v
}

// GetReptCurr returns the ReptCurr field value if set, zero value otherwise.
func (o *GetSolTokenForecastsRequest) GetReptCurr() string {
	if o == nil || IsNil(o.ReptCurr) {
		var ret string
		return ret
	}
	return *o.ReptCurr
}

// GetReptCurrOk returns a tuple with the ReptCurr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolTokenForecastsRequest) GetReptCurrOk() (*string, bool) {
	if o == nil || IsNil(o.ReptCurr) {
		return nil, false
	}
	return o.ReptCurr, true
}

// HasReptCurr returns a boolean if a field has been set.
func (o *GetSolTokenForecastsRequest) HasReptCurr() bool {
	if o != nil && !IsNil(o.ReptCurr) {
		return true
	}

	return false
}

// SetReptCurr gets a reference to the given string and assigns it to the ReptCurr field.
func (o *GetSolTokenForecastsRequest) SetReptCurr(v string) {
	o.ReptCurr = &v
}

// GetExcludeWash returns the ExcludeWash field value if set, zero value otherwise.
func (o *GetSolTokenForecastsRequest) GetExcludeWash() bool {
	if o == nil || IsNil(o.ExcludeWash) {
		var ret bool
		return ret
	}
	return *o.ExcludeWash
}

// GetExcludeWashOk returns a tuple with the ExcludeWash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolTokenForecastsRequest) GetExcludeWashOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeWash) {
		return nil, false
	}
	return o.ExcludeWash, true
}

// HasExcludeWash returns a boolean if a field has been set.
func (o *GetSolTokenForecastsRequest) HasExcludeWash() bool {
	if o != nil && !IsNil(o.ExcludeWash) {
		return true
	}

	return false
}

// SetExcludeWash gets a reference to the given bool and assigns it to the ExcludeWash field.
func (o *GetSolTokenForecastsRequest) SetExcludeWash(v bool) {
	o.ExcludeWash = &v
}

// GetArchParams returns the ArchParams field value if set, zero value otherwise.
func (o *GetSolTokenForecastsRequest) GetArchParams() GetSolTokenForecastsRequestArchParams {
	if o == nil || IsNil(o.ArchParams) {
		var ret GetSolTokenForecastsRequestArchParams
		return ret
	}
	return *o.ArchParams
}

// GetArchParamsOk returns a tuple with the ArchParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolTokenForecastsRequest) GetArchParamsOk() (*GetSolTokenForecastsRequestArchParams, bool) {
	if o == nil || IsNil(o.ArchParams) {
		return nil, false
	}
	return o.ArchParams, true
}

// HasArchParams returns a boolean if a field has been set.
func (o *GetSolTokenForecastsRequest) HasArchParams() bool {
	if o != nil && !IsNil(o.ArchParams) {
		return true
	}

	return false
}

// SetArchParams gets a reference to the given GetSolTokenForecastsRequestArchParams and assigns it to the ArchParams field.
func (o *GetSolTokenForecastsRequest) SetArchParams(v GetSolTokenForecastsRequestArchParams) {
	o.ArchParams = &v
}

func (o GetSolTokenForecastsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSolTokenForecastsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mint_address"] = o.MintAddress
	if !IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
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

type NullableGetSolTokenForecastsRequest struct {
	value *GetSolTokenForecastsRequest
	isSet bool
}

func (v NullableGetSolTokenForecastsRequest) Get() *GetSolTokenForecastsRequest {
	return v.value
}

func (v *NullableGetSolTokenForecastsRequest) Set(val *GetSolTokenForecastsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolTokenForecastsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolTokenForecastsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolTokenForecastsRequest(val *GetSolTokenForecastsRequest) *NullableGetSolTokenForecastsRequest {
	return &NullableGetSolTokenForecastsRequest{value: val, isSet: true}
}

func (v NullableGetSolTokenForecastsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSolTokenForecastsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


