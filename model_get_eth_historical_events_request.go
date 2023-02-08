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

// checks if the GetEthHistoricalEventsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEthHistoricalEventsRequest{}

// GetEthHistoricalEventsRequest struct for GetEthHistoricalEventsRequest
type GetEthHistoricalEventsRequest struct {
	// The contract address of a collection.
	CollectionAddress string `json:"collection_address"`
	// The id for the token.
	TokenId *string `json:"token_id,omitempty"`
	// The pagination cursor.
	Page *int32 `json:"page,omitempty"`
	// The number of records returned per page.
	PageSize *int32 `json:"page_size,omitempty"`
	// The type of event: list, transfer, offer, mint, sale, cancel_list or cancel_offer
	EventType *string `json:"event_type,omitempty"`
}

// NewGetEthHistoricalEventsRequest instantiates a new GetEthHistoricalEventsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEthHistoricalEventsRequest(collectionAddress string) *GetEthHistoricalEventsRequest {
	this := GetEthHistoricalEventsRequest{}
	this.CollectionAddress = collectionAddress
	return &this
}

// NewGetEthHistoricalEventsRequestWithDefaults instantiates a new GetEthHistoricalEventsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEthHistoricalEventsRequestWithDefaults() *GetEthHistoricalEventsRequest {
	this := GetEthHistoricalEventsRequest{}
	return &this
}

// GetCollectionAddress returns the CollectionAddress field value
func (o *GetEthHistoricalEventsRequest) GetCollectionAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value
// and a boolean to check if the value has been set.
func (o *GetEthHistoricalEventsRequest) GetCollectionAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionAddress, true
}

// SetCollectionAddress sets field value
func (o *GetEthHistoricalEventsRequest) SetCollectionAddress(v string) {
	o.CollectionAddress = v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *GetEthHistoricalEventsRequest) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthHistoricalEventsRequest) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *GetEthHistoricalEventsRequest) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *GetEthHistoricalEventsRequest) SetTokenId(v string) {
	o.TokenId = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetEthHistoricalEventsRequest) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthHistoricalEventsRequest) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetEthHistoricalEventsRequest) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *GetEthHistoricalEventsRequest) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *GetEthHistoricalEventsRequest) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthHistoricalEventsRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *GetEthHistoricalEventsRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *GetEthHistoricalEventsRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *GetEthHistoricalEventsRequest) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthHistoricalEventsRequest) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *GetEthHistoricalEventsRequest) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *GetEthHistoricalEventsRequest) SetEventType(v string) {
	o.EventType = &v
}

func (o GetEthHistoricalEventsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEthHistoricalEventsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_address"] = o.CollectionAddress
	if !IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.EventType) {
		toSerialize["event_type"] = o.EventType
	}
	return toSerialize, nil
}

type NullableGetEthHistoricalEventsRequest struct {
	value *GetEthHistoricalEventsRequest
	isSet bool
}

func (v NullableGetEthHistoricalEventsRequest) Get() *GetEthHistoricalEventsRequest {
	return v.value
}

func (v *NullableGetEthHistoricalEventsRequest) Set(val *GetEthHistoricalEventsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEthHistoricalEventsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEthHistoricalEventsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEthHistoricalEventsRequest(val *GetEthHistoricalEventsRequest) *NullableGetEthHistoricalEventsRequest {
	return &NullableGetEthHistoricalEventsRequest{value: val, isSet: true}
}

func (v NullableGetEthHistoricalEventsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEthHistoricalEventsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


