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

// checks if the GetPolDefaultTokenRiskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPolDefaultTokenRiskRequest{}

// GetPolDefaultTokenRiskRequest struct for GetPolDefaultTokenRiskRequest
type GetPolDefaultTokenRiskRequest struct {
	// The contract address of the token collection.
	CollectionAddress string `json:"collection_address"`
	// The id(s) for the token(s).
	TokenId []string `json:"token_id"`
	// The holding period to evaluate risk for, e.g. '12M'
	HoldingPeriod string `json:"holding_period"`
	// The currency to report results in
	ReptCurr *string `json:"rept_curr,omitempty"`
	// If true, report drawdown volatility (based on negative returns only).
	Drawdown *bool `json:"drawdown,omitempty"`
}

// NewGetPolDefaultTokenRiskRequest instantiates a new GetPolDefaultTokenRiskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPolDefaultTokenRiskRequest(collectionAddress string, tokenId []string, holdingPeriod string) *GetPolDefaultTokenRiskRequest {
	this := GetPolDefaultTokenRiskRequest{}
	this.CollectionAddress = collectionAddress
	this.TokenId = tokenId
	this.HoldingPeriod = holdingPeriod
	return &this
}

// NewGetPolDefaultTokenRiskRequestWithDefaults instantiates a new GetPolDefaultTokenRiskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPolDefaultTokenRiskRequestWithDefaults() *GetPolDefaultTokenRiskRequest {
	this := GetPolDefaultTokenRiskRequest{}
	return &this
}

// GetCollectionAddress returns the CollectionAddress field value
func (o *GetPolDefaultTokenRiskRequest) GetCollectionAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value
// and a boolean to check if the value has been set.
func (o *GetPolDefaultTokenRiskRequest) GetCollectionAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionAddress, true
}

// SetCollectionAddress sets field value
func (o *GetPolDefaultTokenRiskRequest) SetCollectionAddress(v string) {
	o.CollectionAddress = v
}

// GetTokenId returns the TokenId field value
func (o *GetPolDefaultTokenRiskRequest) GetTokenId() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *GetPolDefaultTokenRiskRequest) GetTokenIdOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenId, true
}

// SetTokenId sets field value
func (o *GetPolDefaultTokenRiskRequest) SetTokenId(v []string) {
	o.TokenId = v
}

// GetHoldingPeriod returns the HoldingPeriod field value
func (o *GetPolDefaultTokenRiskRequest) GetHoldingPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HoldingPeriod
}

// GetHoldingPeriodOk returns a tuple with the HoldingPeriod field value
// and a boolean to check if the value has been set.
func (o *GetPolDefaultTokenRiskRequest) GetHoldingPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HoldingPeriod, true
}

// SetHoldingPeriod sets field value
func (o *GetPolDefaultTokenRiskRequest) SetHoldingPeriod(v string) {
	o.HoldingPeriod = v
}

// GetReptCurr returns the ReptCurr field value if set, zero value otherwise.
func (o *GetPolDefaultTokenRiskRequest) GetReptCurr() string {
	if o == nil || isNil(o.ReptCurr) {
		var ret string
		return ret
	}
	return *o.ReptCurr
}

// GetReptCurrOk returns a tuple with the ReptCurr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolDefaultTokenRiskRequest) GetReptCurrOk() (*string, bool) {
	if o == nil || isNil(o.ReptCurr) {
		return nil, false
	}
	return o.ReptCurr, true
}

// HasReptCurr returns a boolean if a field has been set.
func (o *GetPolDefaultTokenRiskRequest) HasReptCurr() bool {
	if o != nil && !isNil(o.ReptCurr) {
		return true
	}

	return false
}

// SetReptCurr gets a reference to the given string and assigns it to the ReptCurr field.
func (o *GetPolDefaultTokenRiskRequest) SetReptCurr(v string) {
	o.ReptCurr = &v
}

// GetDrawdown returns the Drawdown field value if set, zero value otherwise.
func (o *GetPolDefaultTokenRiskRequest) GetDrawdown() bool {
	if o == nil || isNil(o.Drawdown) {
		var ret bool
		return ret
	}
	return *o.Drawdown
}

// GetDrawdownOk returns a tuple with the Drawdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolDefaultTokenRiskRequest) GetDrawdownOk() (*bool, bool) {
	if o == nil || isNil(o.Drawdown) {
		return nil, false
	}
	return o.Drawdown, true
}

// HasDrawdown returns a boolean if a field has been set.
func (o *GetPolDefaultTokenRiskRequest) HasDrawdown() bool {
	if o != nil && !isNil(o.Drawdown) {
		return true
	}

	return false
}

// SetDrawdown gets a reference to the given bool and assigns it to the Drawdown field.
func (o *GetPolDefaultTokenRiskRequest) SetDrawdown(v bool) {
	o.Drawdown = &v
}

func (o GetPolDefaultTokenRiskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPolDefaultTokenRiskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_address"] = o.CollectionAddress
	toSerialize["token_id"] = o.TokenId
	toSerialize["holding_period"] = o.HoldingPeriod
	if !isNil(o.ReptCurr) {
		toSerialize["rept_curr"] = o.ReptCurr
	}
	if !isNil(o.Drawdown) {
		toSerialize["drawdown"] = o.Drawdown
	}
	return toSerialize, nil
}

type NullableGetPolDefaultTokenRiskRequest struct {
	value *GetPolDefaultTokenRiskRequest
	isSet bool
}

func (v NullableGetPolDefaultTokenRiskRequest) Get() *GetPolDefaultTokenRiskRequest {
	return v.value
}

func (v *NullableGetPolDefaultTokenRiskRequest) Set(val *GetPolDefaultTokenRiskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPolDefaultTokenRiskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPolDefaultTokenRiskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPolDefaultTokenRiskRequest(val *GetPolDefaultTokenRiskRequest) *NullableGetPolDefaultTokenRiskRequest {
	return &NullableGetPolDefaultTokenRiskRequest{value: val, isSet: true}
}

func (v NullableGetPolDefaultTokenRiskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPolDefaultTokenRiskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


