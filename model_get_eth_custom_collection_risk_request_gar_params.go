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

// checks if the GetEthCustomCollectionRiskRequestGarParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEthCustomCollectionRiskRequestGarParams{}

// GetEthCustomCollectionRiskRequestGarParams JSON containing options for the GARCH model.
type GetEthCustomCollectionRiskRequestGarParams struct {
	// Estimator for the location model of the time series, e.g: `Zero`, `Constant`, `ARX`, ... .
	Mean *string `json:"mean,omitempty"`
	// Number of time time period lags considered. Note that the time period is set by the `frequency` parameter, so a value of 2 will assume 2-day lags if `frequency=='1D'`.
	Lags *int32 `json:"lags,omitempty"`
	// Estimator used for the volatility process of the time series, e.g: `Constant`, `ARCH`, `GARCH`, ... 
	Vol *string `json:"vol,omitempty"`
	// Order of the symmetric innovation(s).
	P *int32 `json:"p,omitempty"`
	// Return distribution assumed.
	Dist *string `json:"dist,omitempty"`
}

// NewGetEthCustomCollectionRiskRequestGarParams instantiates a new GetEthCustomCollectionRiskRequestGarParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEthCustomCollectionRiskRequestGarParams() *GetEthCustomCollectionRiskRequestGarParams {
	this := GetEthCustomCollectionRiskRequestGarParams{}
	return &this
}

// NewGetEthCustomCollectionRiskRequestGarParamsWithDefaults instantiates a new GetEthCustomCollectionRiskRequestGarParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEthCustomCollectionRiskRequestGarParamsWithDefaults() *GetEthCustomCollectionRiskRequestGarParams {
	this := GetEthCustomCollectionRiskRequestGarParams{}
	return &this
}

// GetMean returns the Mean field value if set, zero value otherwise.
func (o *GetEthCustomCollectionRiskRequestGarParams) GetMean() string {
	if o == nil || isNil(o.Mean) {
		var ret string
		return ret
	}
	return *o.Mean
}

// GetMeanOk returns a tuple with the Mean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomCollectionRiskRequestGarParams) GetMeanOk() (*string, bool) {
	if o == nil || isNil(o.Mean) {
		return nil, false
	}
	return o.Mean, true
}

// HasMean returns a boolean if a field has been set.
func (o *GetEthCustomCollectionRiskRequestGarParams) HasMean() bool {
	if o != nil && !isNil(o.Mean) {
		return true
	}

	return false
}

// SetMean gets a reference to the given string and assigns it to the Mean field.
func (o *GetEthCustomCollectionRiskRequestGarParams) SetMean(v string) {
	o.Mean = &v
}

// GetLags returns the Lags field value if set, zero value otherwise.
func (o *GetEthCustomCollectionRiskRequestGarParams) GetLags() int32 {
	if o == nil || isNil(o.Lags) {
		var ret int32
		return ret
	}
	return *o.Lags
}

// GetLagsOk returns a tuple with the Lags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomCollectionRiskRequestGarParams) GetLagsOk() (*int32, bool) {
	if o == nil || isNil(o.Lags) {
		return nil, false
	}
	return o.Lags, true
}

// HasLags returns a boolean if a field has been set.
func (o *GetEthCustomCollectionRiskRequestGarParams) HasLags() bool {
	if o != nil && !isNil(o.Lags) {
		return true
	}

	return false
}

// SetLags gets a reference to the given int32 and assigns it to the Lags field.
func (o *GetEthCustomCollectionRiskRequestGarParams) SetLags(v int32) {
	o.Lags = &v
}

// GetVol returns the Vol field value if set, zero value otherwise.
func (o *GetEthCustomCollectionRiskRequestGarParams) GetVol() string {
	if o == nil || isNil(o.Vol) {
		var ret string
		return ret
	}
	return *o.Vol
}

// GetVolOk returns a tuple with the Vol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomCollectionRiskRequestGarParams) GetVolOk() (*string, bool) {
	if o == nil || isNil(o.Vol) {
		return nil, false
	}
	return o.Vol, true
}

// HasVol returns a boolean if a field has been set.
func (o *GetEthCustomCollectionRiskRequestGarParams) HasVol() bool {
	if o != nil && !isNil(o.Vol) {
		return true
	}

	return false
}

// SetVol gets a reference to the given string and assigns it to the Vol field.
func (o *GetEthCustomCollectionRiskRequestGarParams) SetVol(v string) {
	o.Vol = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *GetEthCustomCollectionRiskRequestGarParams) GetP() int32 {
	if o == nil || isNil(o.P) {
		var ret int32
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomCollectionRiskRequestGarParams) GetPOk() (*int32, bool) {
	if o == nil || isNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *GetEthCustomCollectionRiskRequestGarParams) HasP() bool {
	if o != nil && !isNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given int32 and assigns it to the P field.
func (o *GetEthCustomCollectionRiskRequestGarParams) SetP(v int32) {
	o.P = &v
}

// GetDist returns the Dist field value if set, zero value otherwise.
func (o *GetEthCustomCollectionRiskRequestGarParams) GetDist() string {
	if o == nil || isNil(o.Dist) {
		var ret string
		return ret
	}
	return *o.Dist
}

// GetDistOk returns a tuple with the Dist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomCollectionRiskRequestGarParams) GetDistOk() (*string, bool) {
	if o == nil || isNil(o.Dist) {
		return nil, false
	}
	return o.Dist, true
}

// HasDist returns a boolean if a field has been set.
func (o *GetEthCustomCollectionRiskRequestGarParams) HasDist() bool {
	if o != nil && !isNil(o.Dist) {
		return true
	}

	return false
}

// SetDist gets a reference to the given string and assigns it to the Dist field.
func (o *GetEthCustomCollectionRiskRequestGarParams) SetDist(v string) {
	o.Dist = &v
}

func (o GetEthCustomCollectionRiskRequestGarParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEthCustomCollectionRiskRequestGarParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mean) {
		toSerialize["mean"] = o.Mean
	}
	if !isNil(o.Lags) {
		toSerialize["lags"] = o.Lags
	}
	if !isNil(o.Vol) {
		toSerialize["vol"] = o.Vol
	}
	if !isNil(o.P) {
		toSerialize["p"] = o.P
	}
	if !isNil(o.Dist) {
		toSerialize["dist"] = o.Dist
	}
	return toSerialize, nil
}

type NullableGetEthCustomCollectionRiskRequestGarParams struct {
	value *GetEthCustomCollectionRiskRequestGarParams
	isSet bool
}

func (v NullableGetEthCustomCollectionRiskRequestGarParams) Get() *GetEthCustomCollectionRiskRequestGarParams {
	return v.value
}

func (v *NullableGetEthCustomCollectionRiskRequestGarParams) Set(val *GetEthCustomCollectionRiskRequestGarParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEthCustomCollectionRiskRequestGarParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEthCustomCollectionRiskRequestGarParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEthCustomCollectionRiskRequestGarParams(val *GetEthCustomCollectionRiskRequestGarParams) *NullableGetEthCustomCollectionRiskRequestGarParams {
	return &NullableGetEthCustomCollectionRiskRequestGarParams{value: val, isSet: true}
}

func (v NullableGetEthCustomCollectionRiskRequestGarParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEthCustomCollectionRiskRequestGarParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


