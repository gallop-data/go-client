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

// checks if the GetPolTokenTransactionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPolTokenTransactionsRequest{}

// GetPolTokenTransactionsRequest struct for GetPolTokenTransactionsRequest
type GetPolTokenTransactionsRequest struct {
	// The contract address the token belongs to.
	CollectionAddress string `json:"collection_address"`
	// The token id.
	TokenId string `json:"token_id"`
	// The pagination cursor.
	Page *int32 `json:"page,omitempty"`
	// The number of records returned per page.
	PageSize *int32 `json:"page_size,omitempty"`
	// The earliest block timestamp.
	StartDate *string `json:"start_date,omitempty"`
	// The oldest block number to return.
	StartBlockNumber *int32 `json:"start_block_number,omitempty"`
	// The latest block timestamp.
	EndDate *string `json:"end_date,omitempty"`
}

// NewGetPolTokenTransactionsRequest instantiates a new GetPolTokenTransactionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPolTokenTransactionsRequest(collectionAddress string, tokenId string) *GetPolTokenTransactionsRequest {
	this := GetPolTokenTransactionsRequest{}
	this.CollectionAddress = collectionAddress
	this.TokenId = tokenId
	return &this
}

// NewGetPolTokenTransactionsRequestWithDefaults instantiates a new GetPolTokenTransactionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPolTokenTransactionsRequestWithDefaults() *GetPolTokenTransactionsRequest {
	this := GetPolTokenTransactionsRequest{}
	return &this
}

// GetCollectionAddress returns the CollectionAddress field value
func (o *GetPolTokenTransactionsRequest) GetCollectionAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value
// and a boolean to check if the value has been set.
func (o *GetPolTokenTransactionsRequest) GetCollectionAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionAddress, true
}

// SetCollectionAddress sets field value
func (o *GetPolTokenTransactionsRequest) SetCollectionAddress(v string) {
	o.CollectionAddress = v
}

// GetTokenId returns the TokenId field value
func (o *GetPolTokenTransactionsRequest) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *GetPolTokenTransactionsRequest) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *GetPolTokenTransactionsRequest) SetTokenId(v string) {
	o.TokenId = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetPolTokenTransactionsRequest) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolTokenTransactionsRequest) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetPolTokenTransactionsRequest) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *GetPolTokenTransactionsRequest) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *GetPolTokenTransactionsRequest) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolTokenTransactionsRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *GetPolTokenTransactionsRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *GetPolTokenTransactionsRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *GetPolTokenTransactionsRequest) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolTokenTransactionsRequest) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *GetPolTokenTransactionsRequest) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *GetPolTokenTransactionsRequest) SetStartDate(v string) {
	o.StartDate = &v
}

// GetStartBlockNumber returns the StartBlockNumber field value if set, zero value otherwise.
func (o *GetPolTokenTransactionsRequest) GetStartBlockNumber() int32 {
	if o == nil || IsNil(o.StartBlockNumber) {
		var ret int32
		return ret
	}
	return *o.StartBlockNumber
}

// GetStartBlockNumberOk returns a tuple with the StartBlockNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolTokenTransactionsRequest) GetStartBlockNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.StartBlockNumber) {
		return nil, false
	}
	return o.StartBlockNumber, true
}

// HasStartBlockNumber returns a boolean if a field has been set.
func (o *GetPolTokenTransactionsRequest) HasStartBlockNumber() bool {
	if o != nil && !IsNil(o.StartBlockNumber) {
		return true
	}

	return false
}

// SetStartBlockNumber gets a reference to the given int32 and assigns it to the StartBlockNumber field.
func (o *GetPolTokenTransactionsRequest) SetStartBlockNumber(v int32) {
	o.StartBlockNumber = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GetPolTokenTransactionsRequest) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPolTokenTransactionsRequest) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GetPolTokenTransactionsRequest) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *GetPolTokenTransactionsRequest) SetEndDate(v string) {
	o.EndDate = &v
}

func (o GetPolTokenTransactionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPolTokenTransactionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_address"] = o.CollectionAddress
	toSerialize["token_id"] = o.TokenId
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.StartBlockNumber) {
		toSerialize["start_block_number"] = o.StartBlockNumber
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableGetPolTokenTransactionsRequest struct {
	value *GetPolTokenTransactionsRequest
	isSet bool
}

func (v NullableGetPolTokenTransactionsRequest) Get() *GetPolTokenTransactionsRequest {
	return v.value
}

func (v *NullableGetPolTokenTransactionsRequest) Set(val *GetPolTokenTransactionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPolTokenTransactionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPolTokenTransactionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPolTokenTransactionsRequest(val *GetPolTokenTransactionsRequest) *NullableGetPolTokenTransactionsRequest {
	return &NullableGetPolTokenTransactionsRequest{value: val, isSet: true}
}

func (v NullableGetPolTokenTransactionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPolTokenTransactionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


