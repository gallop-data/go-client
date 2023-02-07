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

// checks if the GetEthHistoricalTransactionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEthHistoricalTransactionsRequest{}

// GetEthHistoricalTransactionsRequest struct for GetEthHistoricalTransactionsRequest
type GetEthHistoricalTransactionsRequest struct {
	// The contract address of the collection.
	CollectionAddress string `json:"collection_address"`
	// The id for the token.
	TokenId *string `json:"token_id,omitempty"`
	// The pagination cursor.
	Page *int32 `json:"page,omitempty"`
}

// NewGetEthHistoricalTransactionsRequest instantiates a new GetEthHistoricalTransactionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEthHistoricalTransactionsRequest(collectionAddress string) *GetEthHistoricalTransactionsRequest {
	this := GetEthHistoricalTransactionsRequest{}
	this.CollectionAddress = collectionAddress
	return &this
}

// NewGetEthHistoricalTransactionsRequestWithDefaults instantiates a new GetEthHistoricalTransactionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEthHistoricalTransactionsRequestWithDefaults() *GetEthHistoricalTransactionsRequest {
	this := GetEthHistoricalTransactionsRequest{}
	return &this
}

// GetCollectionAddress returns the CollectionAddress field value
func (o *GetEthHistoricalTransactionsRequest) GetCollectionAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value
// and a boolean to check if the value has been set.
func (o *GetEthHistoricalTransactionsRequest) GetCollectionAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionAddress, true
}

// SetCollectionAddress sets field value
func (o *GetEthHistoricalTransactionsRequest) SetCollectionAddress(v string) {
	o.CollectionAddress = v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *GetEthHistoricalTransactionsRequest) GetTokenId() string {
	if o == nil || isNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthHistoricalTransactionsRequest) GetTokenIdOk() (*string, bool) {
	if o == nil || isNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *GetEthHistoricalTransactionsRequest) HasTokenId() bool {
	if o != nil && !isNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *GetEthHistoricalTransactionsRequest) SetTokenId(v string) {
	o.TokenId = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetEthHistoricalTransactionsRequest) GetPage() int32 {
	if o == nil || isNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthHistoricalTransactionsRequest) GetPageOk() (*int32, bool) {
	if o == nil || isNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetEthHistoricalTransactionsRequest) HasPage() bool {
	if o != nil && !isNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *GetEthHistoricalTransactionsRequest) SetPage(v int32) {
	o.Page = &v
}

func (o GetEthHistoricalTransactionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEthHistoricalTransactionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_address"] = o.CollectionAddress
	if !isNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	if !isNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	return toSerialize, nil
}

type NullableGetEthHistoricalTransactionsRequest struct {
	value *GetEthHistoricalTransactionsRequest
	isSet bool
}

func (v NullableGetEthHistoricalTransactionsRequest) Get() *GetEthHistoricalTransactionsRequest {
	return v.value
}

func (v *NullableGetEthHistoricalTransactionsRequest) Set(val *GetEthHistoricalTransactionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEthHistoricalTransactionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEthHistoricalTransactionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEthHistoricalTransactionsRequest(val *GetEthHistoricalTransactionsRequest) *NullableGetEthHistoricalTransactionsRequest {
	return &NullableGetEthHistoricalTransactionsRequest{value: val, isSet: true}
}

func (v NullableGetEthHistoricalTransactionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEthHistoricalTransactionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


