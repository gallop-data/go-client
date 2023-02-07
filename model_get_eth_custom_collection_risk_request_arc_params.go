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

// checks if the GetEthCustomCollectionRiskRequestArcParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEthCustomCollectionRiskRequestArcParams{}

// GetEthCustomCollectionRiskRequestArcParams JSON containing options for the ARCH model.
type GetEthCustomCollectionRiskRequestArcParams struct {
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

// NewGetEthCustomCollectionRiskRequestArcParams instantiates a new GetEthCustomCollectionRiskRequestArcParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEthCustomCollectionRiskRequestArcParams() *GetEthCustomCollectionRiskRequestArcParams {
	this := GetEthCustomCollectionRiskRequestArcParams{}
	return &this
}

// NewGetEthCustomCollectionRiskRequestArcParamsWithDefaults instantiates a new GetEthCustomCollectionRiskRequestArcParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEthCustomCollectionRiskRequestArcParamsWithDefaults() *GetEthCustomCollectionRiskRequestArcParams {
	this := GetEthCustomCollectionRiskRequestArcParams{}
	return &this
}

// GetMean returns the Mean field value if set, zero value otherwise.
func (o *GetEthCustomCollectionRiskRequestArcParams) GetMean() string {
	if o == nil || isNil(o.Mean) {
		var ret string
		return ret
	}
	return *o.Mean
}

// GetMeanOk returns a tuple with the Mean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomCollectionRiskRequestArcParams) GetMeanOk() (*string, bool) {
	if o == nil || isNil(o.Mean) {
		return nil, false
	}
	return o.Mean, true
}

// HasMean returns a boolean if a field has been set.
func (o *GetEthCustomCollectionRiskRequestArcParams) HasMean() bool {
	if o != nil && !isNil(o.Mean) {
		return true
	}

	return false
}

// SetMean gets a reference to the given string and assigns it to the Mean field.
func (o *GetEthCustomCollectionRiskRequestArcParams) SetMean(v string) {
	o.Mean = &v
}

// GetLags returns the Lags field value if set, zero value otherwise.
func (o *GetEthCustomCollectionRiskRequestArcParams) GetLags() int32 {
	if o == nil || isNil(o.Lags) {
		var ret int32
		return ret
	}
	return *o.Lags
}

// GetLagsOk returns a tuple with the Lags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomCollectionRiskRequestArcParams) GetLagsOk() (*int32, bool) {
	if o == nil || isNil(o.Lags) {
		return nil, false
	}
	return o.Lags, true
}

// HasLags returns a boolean if a field has been set.
func (o *GetEthCustomCollectionRiskRequestArcParams) HasLags() bool {
	if o != nil && !isNil(o.Lags) {
		return true
	}

	return false
}

// SetLags gets a reference to the given int32 and assigns it to the Lags field.
func (o *GetEthCustomCollectionRiskRequestArcParams) SetLags(v int32) {
	o.Lags = &v
}

// GetVol returns the Vol field value if set, zero value otherwise.
func (o *GetEthCustomCollectionRiskRequestArcParams) GetVol() string {
	if o == nil || isNil(o.Vol) {
		var ret string
		return ret
	}
	return *o.Vol
}

// GetVolOk returns a tuple with the Vol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomCollectionRiskRequestArcParams) GetVolOk() (*string, bool) {
	if o == nil || isNil(o.Vol) {
		return nil, false
	}
	return o.Vol, true
}

// HasVol returns a boolean if a field has been set.
func (o *GetEthCustomCollectionRiskRequestArcParams) HasVol() bool {
	if o != nil && !isNil(o.Vol) {
		return true
	}

	return false
}

// SetVol gets a reference to the given string and assigns it to the Vol field.
func (o *GetEthCustomCollectionRiskRequestArcParams) SetVol(v string) {
	o.Vol = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *GetEthCustomCollectionRiskRequestArcParams) GetP() int32 {
	if o == nil || isNil(o.P) {
		var ret int32
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomCollectionRiskRequestArcParams) GetPOk() (*int32, bool) {
	if o == nil || isNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *GetEthCustomCollectionRiskRequestArcParams) HasP() bool {
	if o != nil && !isNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given int32 and assigns it to the P field.
func (o *GetEthCustomCollectionRiskRequestArcParams) SetP(v int32) {
	o.P = &v
}

// GetDist returns the Dist field value if set, zero value otherwise.
func (o *GetEthCustomCollectionRiskRequestArcParams) GetDist() string {
	if o == nil || isNil(o.Dist) {
		var ret string
		return ret
	}
	return *o.Dist
}

// GetDistOk returns a tuple with the Dist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthCustomCollectionRiskRequestArcParams) GetDistOk() (*string, bool) {
	if o == nil || isNil(o.Dist) {
		return nil, false
	}
	return o.Dist, true
}

// HasDist returns a boolean if a field has been set.
func (o *GetEthCustomCollectionRiskRequestArcParams) HasDist() bool {
	if o != nil && !isNil(o.Dist) {
		return true
	}

	return false
}

// SetDist gets a reference to the given string and assigns it to the Dist field.
func (o *GetEthCustomCollectionRiskRequestArcParams) SetDist(v string) {
	o.Dist = &v
}

func (o GetEthCustomCollectionRiskRequestArcParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEthCustomCollectionRiskRequestArcParams) ToMap() (map[string]interface{}, error) {
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

type NullableGetEthCustomCollectionRiskRequestArcParams struct {
	value *GetEthCustomCollectionRiskRequestArcParams
	isSet bool
}

func (v NullableGetEthCustomCollectionRiskRequestArcParams) Get() *GetEthCustomCollectionRiskRequestArcParams {
	return v.value
}

func (v *NullableGetEthCustomCollectionRiskRequestArcParams) Set(val *GetEthCustomCollectionRiskRequestArcParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEthCustomCollectionRiskRequestArcParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEthCustomCollectionRiskRequestArcParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEthCustomCollectionRiskRequestArcParams(val *GetEthCustomCollectionRiskRequestArcParams) *NullableGetEthCustomCollectionRiskRequestArcParams {
	return &NullableGetEthCustomCollectionRiskRequestArcParams{value: val, isSet: true}
}

func (v NullableGetEthCustomCollectionRiskRequestArcParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEthCustomCollectionRiskRequestArcParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


