# GetEthMarketplaceDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **[]string** | Array of collection addresses | 
**SubCollectionTags** | Pointer to **[]string** | Array of sub collections (e.g. Art Blocks) | [optional] 

## Methods

### NewGetEthMarketplaceDataRequest

`func NewGetEthMarketplaceDataRequest(collectionAddress []string, ) *GetEthMarketplaceDataRequest`

NewGetEthMarketplaceDataRequest instantiates a new GetEthMarketplaceDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthMarketplaceDataRequestWithDefaults

`func NewGetEthMarketplaceDataRequestWithDefaults() *GetEthMarketplaceDataRequest`

NewGetEthMarketplaceDataRequestWithDefaults instantiates a new GetEthMarketplaceDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetEthMarketplaceDataRequest) GetCollectionAddress() []string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetEthMarketplaceDataRequest) GetCollectionAddressOk() (*[]string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetEthMarketplaceDataRequest) SetCollectionAddress(v []string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetSubCollectionTags

`func (o *GetEthMarketplaceDataRequest) GetSubCollectionTags() []string`

GetSubCollectionTags returns the SubCollectionTags field if non-nil, zero value otherwise.

### GetSubCollectionTagsOk

`func (o *GetEthMarketplaceDataRequest) GetSubCollectionTagsOk() (*[]string, bool)`

GetSubCollectionTagsOk returns a tuple with the SubCollectionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCollectionTags

`func (o *GetEthMarketplaceDataRequest) SetSubCollectionTags(v []string)`

SetSubCollectionTags sets SubCollectionTags field to given value.

### HasSubCollectionTags

`func (o *GetEthMarketplaceDataRequest) HasSubCollectionTags() bool`

HasSubCollectionTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


