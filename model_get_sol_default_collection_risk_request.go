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

// checks if the GetSolDefaultCollectionRiskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSolDefaultCollectionRiskRequest{}

// GetSolDefaultCollectionRiskRequest struct for GetSolDefaultCollectionRiskRequest
type GetSolDefaultCollectionRiskRequest struct {
	// The Gallop tag to identify the collection.
	CollectionTag string `json:"collection_tag"`
	// The holding period to evaluate risk for, e.g. '12M'
	HoldingPeriod string `json:"holding_period"`
	// The amount of tokens in your portfolio
	Amount *int32 `json:"amount,omitempty"`
	// The currency to report results in
	ReptCurr *string `json:"rept_curr,omitempty"`
	// If true, report drawdown volatility (based on negative returns only).
	Drawdown *bool `json:"drawdown,omitempty"`
}

// NewGetSolDefaultCollectionRiskRequest instantiates a new GetSolDefaultCollectionRiskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolDefaultCollectionRiskRequest(collectionTag string, holdingPeriod string) *GetSolDefaultCollectionRiskRequest {
	this := GetSolDefaultCollectionRiskRequest{}
	this.CollectionTag = collectionTag
	this.HoldingPeriod = holdingPeriod
	return &this
}

// NewGetSolDefaultCollectionRiskRequestWithDefaults instantiates a new GetSolDefaultCollectionRiskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolDefaultCollectionRiskRequestWithDefaults() *GetSolDefaultCollectionRiskRequest {
	this := GetSolDefaultCollectionRiskRequest{}
	return &this
}

// GetCollectionTag returns the CollectionTag field value
func (o *GetSolDefaultCollectionRiskRequest) GetCollectionTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionTag
}

// GetCollectionTagOk returns a tuple with the CollectionTag field value
// and a boolean to check if the value has been set.
func (o *GetSolDefaultCollectionRiskRequest) GetCollectionTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionTag, true
}

// SetCollectionTag sets field value
func (o *GetSolDefaultCollectionRiskRequest) SetCollectionTag(v string) {
	o.CollectionTag = v
}

// GetHoldingPeriod returns the HoldingPeriod field value
func (o *GetSolDefaultCollectionRiskRequest) GetHoldingPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HoldingPeriod
}

// GetHoldingPeriodOk returns a tuple with the HoldingPeriod field value
// and a boolean to check if the value has been set.
func (o *GetSolDefaultCollectionRiskRequest) GetHoldingPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HoldingPeriod, true
}

// SetHoldingPeriod sets field value
func (o *GetSolDefaultCollectionRiskRequest) SetHoldingPeriod(v string) {
	o.HoldingPeriod = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetSolDefaultCollectionRiskRequest) GetAmount() int32 {
	if o == nil || isNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolDefaultCollectionRiskRequest) GetAmountOk() (*int32, bool) {
	if o == nil || isNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetSolDefaultCollectionRiskRequest) HasAmount() bool {
	if o != nil && !isNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *GetSolDefaultCollectionRiskRequest) SetAmount(v int32) {
	o.Amount = &v
}

// GetReptCurr returns the ReptCurr field value if set, zero value otherwise.
func (o *GetSolDefaultCollectionRiskRequest) GetReptCurr() string {
	if o == nil || isNil(o.ReptCurr) {
		var ret string
		return ret
	}
	return *o.ReptCurr
}

// GetReptCurrOk returns a tuple with the ReptCurr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolDefaultCollectionRiskRequest) GetReptCurrOk() (*string, bool) {
	if o == nil || isNil(o.ReptCurr) {
		return nil, false
	}
	return o.ReptCurr, true
}

// HasReptCurr returns a boolean if a field has been set.
func (o *GetSolDefaultCollectionRiskRequest) HasReptCurr() bool {
	if o != nil && !isNil(o.ReptCurr) {
		return true
	}

	return false
}

// SetReptCurr gets a reference to the given string and assigns it to the ReptCurr field.
func (o *GetSolDefaultCollectionRiskRequest) SetReptCurr(v string) {
	o.ReptCurr = &v
}

// GetDrawdown returns the Drawdown field value if set, zero value otherwise.
func (o *GetSolDefaultCollectionRiskRequest) GetDrawdown() bool {
	if o == nil || isNil(o.Drawdown) {
		var ret bool
		return ret
	}
	return *o.Drawdown
}

// GetDrawdownOk returns a tuple with the Drawdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolDefaultCollectionRiskRequest) GetDrawdownOk() (*bool, bool) {
	if o == nil || isNil(o.Drawdown) {
		return nil, false
	}
	return o.Drawdown, true
}

// HasDrawdown returns a boolean if a field has been set.
func (o *GetSolDefaultCollectionRiskRequest) HasDrawdown() bool {
	if o != nil && !isNil(o.Drawdown) {
		return true
	}

	return false
}

// SetDrawdown gets a reference to the given bool and assigns it to the Drawdown field.
func (o *GetSolDefaultCollectionRiskRequest) SetDrawdown(v bool) {
	o.Drawdown = &v
}

func (o GetSolDefaultCollectionRiskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSolDefaultCollectionRiskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_tag"] = o.CollectionTag
	toSerialize["holding_period"] = o.HoldingPeriod
	if !isNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !isNil(o.ReptCurr) {
		toSerialize["rept_curr"] = o.ReptCurr
	}
	if !isNil(o.Drawdown) {
		toSerialize["drawdown"] = o.Drawdown
	}
	return toSerialize, nil
}

type NullableGetSolDefaultCollectionRiskRequest struct {
	value *GetSolDefaultCollectionRiskRequest
	isSet bool
}

func (v NullableGetSolDefaultCollectionRiskRequest) Get() *GetSolDefaultCollectionRiskRequest {
	return v.value
}

func (v *NullableGetSolDefaultCollectionRiskRequest) Set(val *GetSolDefaultCollectionRiskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolDefaultCollectionRiskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolDefaultCollectionRiskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolDefaultCollectionRiskRequest(val *GetSolDefaultCollectionRiskRequest) *NullableGetSolDefaultCollectionRiskRequest {
	return &NullableGetSolDefaultCollectionRiskRequest{value: val, isSet: true}
}

func (v NullableGetSolDefaultCollectionRiskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSolDefaultCollectionRiskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

