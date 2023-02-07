# GetSolMarketplaceFloorPriceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | The pagination cursor. | [optional] 
**PageSize** | Pointer to **int32** | The number of records returned per page. | [optional] 
**CollectionTag** | Pointer to **[]string** | Array of Gallop collection tags | [optional] 

## Methods

### NewGetSolMarketplaceFloorPriceRequest

`func NewGetSolMarketplaceFloorPriceRequest() *GetSolMarketplaceFloorPriceRequest`

NewGetSolMarketplaceFloorPriceRequest instantiates a new GetSolMarketplaceFloorPriceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSolMarketplaceFloorPriceRequestWithDefaults

`func NewGetSolMarketplaceFloorPriceRequestWithDefaults() *GetSolMarketplaceFloorPriceRequest`

NewGetSolMarketplaceFloorPriceRequestWithDefaults instantiates a new GetSolMarketplaceFloorPriceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetSolMarketplaceFloorPriceRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetSolMarketplaceFloorPriceRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetSolMarketplaceFloorPriceRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetSolMarketplaceFloorPriceRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *GetSolMarketplaceFloorPriceRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetSolMarketplaceFloorPriceRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetSolMarketplaceFloorPriceRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetSolMarketplaceFloorPriceRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetCollectionTag

`func (o *GetSolMarketplaceFloorPriceRequest) GetCollectionTag() []string`

GetCollectionTag returns the CollectionTag field if non-nil, zero value otherwise.

### GetCollectionTagOk

`func (o *GetSolMarketplaceFloorPriceRequest) GetCollectionTagOk() (*[]string, bool)`

GetCollectionTagOk returns a tuple with the CollectionTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTag

`func (o *GetSolMarketplaceFloorPriceRequest) SetCollectionTag(v []string)`

SetCollectionTag sets CollectionTag field to given value.

### HasCollectionTag

`func (o *GetSolMarketplaceFloorPriceRequest) HasCollectionTag() bool`

HasCollectionTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


