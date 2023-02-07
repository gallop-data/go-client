# GetSolCollectionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | The pagination cursor. | [optional] 
**PageSize** | Pointer to **int32** | The number of records returned per page. | [optional] 
**CollectionName** | Pointer to **string** | The name of the collection searched. | [optional] 

## Methods

### NewGetSolCollectionsRequest

`func NewGetSolCollectionsRequest() *GetSolCollectionsRequest`

NewGetSolCollectionsRequest instantiates a new GetSolCollectionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSolCollectionsRequestWithDefaults

`func NewGetSolCollectionsRequestWithDefaults() *GetSolCollectionsRequest`

NewGetSolCollectionsRequestWithDefaults instantiates a new GetSolCollectionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetSolCollectionsRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetSolCollectionsRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetSolCollectionsRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetSolCollectionsRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *GetSolCollectionsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetSolCollectionsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetSolCollectionsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetSolCollectionsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetCollectionName

`func (o *GetSolCollectionsRequest) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *GetSolCollectionsRequest) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *GetSolCollectionsRequest) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.

### HasCollectionName

`func (o *GetSolCollectionsRequest) HasCollectionName() bool`

HasCollectionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


