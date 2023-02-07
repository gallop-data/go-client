# GetPolMarketplaceFloorPriceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | The pagination cursor. | [optional] 
**PageSize** | Pointer to **int32** | The number of records returned per page. | [optional] 
**CollectionAddress** | Pointer to **[]string** | Array of collection addresses | [optional] 

## Methods

### NewGetPolMarketplaceFloorPriceRequest

`func NewGetPolMarketplaceFloorPriceRequest() *GetPolMarketplaceFloorPriceRequest`

NewGetPolMarketplaceFloorPriceRequest instantiates a new GetPolMarketplaceFloorPriceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPolMarketplaceFloorPriceRequestWithDefaults

`func NewGetPolMarketplaceFloorPriceRequestWithDefaults() *GetPolMarketplaceFloorPriceRequest`

NewGetPolMarketplaceFloorPriceRequestWithDefaults instantiates a new GetPolMarketplaceFloorPriceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetPolMarketplaceFloorPriceRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetPolMarketplaceFloorPriceRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetPolMarketplaceFloorPriceRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetPolMarketplaceFloorPriceRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *GetPolMarketplaceFloorPriceRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetPolMarketplaceFloorPriceRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetPolMarketplaceFloorPriceRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetPolMarketplaceFloorPriceRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetCollectionAddress

`func (o *GetPolMarketplaceFloorPriceRequest) GetCollectionAddress() []string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetPolMarketplaceFloorPriceRequest) GetCollectionAddressOk() (*[]string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetPolMarketplaceFloorPriceRequest) SetCollectionAddress(v []string)`

SetCollectionAddress sets CollectionAddress field to given value.

### HasCollectionAddress

`func (o *GetPolMarketplaceFloorPriceRequest) HasCollectionAddress() bool`

HasCollectionAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


