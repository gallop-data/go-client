# GetEthCollectionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | The pagination cursor. | [optional] 
**PageSize** | Pointer to **int32** | The number of records returned per page. | [optional] 
**CollectionName** | Pointer to **string** | The name of the collection searched. | [optional] 
**Traded** | Pointer to **bool** | Only return collections that have traded. | [optional] 
**CreatedAfter** | Pointer to **string** | Only return collections recorded after this day [YYYY-MM-DD] | [optional] 
**SortBy** | Pointer to **string** | The value to sort by. Defaults to created_at | [optional] 

## Methods

### NewGetEthCollectionsRequest

`func NewGetEthCollectionsRequest() *GetEthCollectionsRequest`

NewGetEthCollectionsRequest instantiates a new GetEthCollectionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthCollectionsRequestWithDefaults

`func NewGetEthCollectionsRequestWithDefaults() *GetEthCollectionsRequest`

NewGetEthCollectionsRequestWithDefaults instantiates a new GetEthCollectionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetEthCollectionsRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetEthCollectionsRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetEthCollectionsRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetEthCollectionsRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *GetEthCollectionsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetEthCollectionsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetEthCollectionsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetEthCollectionsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetCollectionName

`func (o *GetEthCollectionsRequest) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *GetEthCollectionsRequest) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *GetEthCollectionsRequest) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.

### HasCollectionName

`func (o *GetEthCollectionsRequest) HasCollectionName() bool`

HasCollectionName returns a boolean if a field has been set.

### GetTraded

`func (o *GetEthCollectionsRequest) GetTraded() bool`

GetTraded returns the Traded field if non-nil, zero value otherwise.

### GetTradedOk

`func (o *GetEthCollectionsRequest) GetTradedOk() (*bool, bool)`

GetTradedOk returns a tuple with the Traded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraded

`func (o *GetEthCollectionsRequest) SetTraded(v bool)`

SetTraded sets Traded field to given value.

### HasTraded

`func (o *GetEthCollectionsRequest) HasTraded() bool`

HasTraded returns a boolean if a field has been set.

### GetCreatedAfter

`func (o *GetEthCollectionsRequest) GetCreatedAfter() string`

GetCreatedAfter returns the CreatedAfter field if non-nil, zero value otherwise.

### GetCreatedAfterOk

`func (o *GetEthCollectionsRequest) GetCreatedAfterOk() (*string, bool)`

GetCreatedAfterOk returns a tuple with the CreatedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAfter

`func (o *GetEthCollectionsRequest) SetCreatedAfter(v string)`

SetCreatedAfter sets CreatedAfter field to given value.

### HasCreatedAfter

`func (o *GetEthCollectionsRequest) HasCreatedAfter() bool`

HasCreatedAfter returns a boolean if a field has been set.

### GetSortBy

`func (o *GetEthCollectionsRequest) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *GetEthCollectionsRequest) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *GetEthCollectionsRequest) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *GetEthCollectionsRequest) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


