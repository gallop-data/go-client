# GetEthTokensRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The contract address of the token collection. | 
**TokenId** | Pointer to **[]string** | A list of token ids. | [optional] 
**Page** | Pointer to **int32** | The pagination cursor. | [optional] 
**PageSize** | Pointer to **int32** | The number of records returned per page. | [optional] 

## Methods

### NewGetEthTokensRequest

`func NewGetEthTokensRequest(collectionAddress string, ) *GetEthTokensRequest`

NewGetEthTokensRequest instantiates a new GetEthTokensRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthTokensRequestWithDefaults

`func NewGetEthTokensRequestWithDefaults() *GetEthTokensRequest`

NewGetEthTokensRequestWithDefaults instantiates a new GetEthTokensRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetEthTokensRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetEthTokensRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetEthTokensRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetTokenId

`func (o *GetEthTokensRequest) GetTokenId() []string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetEthTokensRequest) GetTokenIdOk() (*[]string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetEthTokensRequest) SetTokenId(v []string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *GetEthTokensRequest) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetPage

`func (o *GetEthTokensRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetEthTokensRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetEthTokensRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetEthTokensRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *GetEthTokensRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetEthTokensRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetEthTokensRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetEthTokensRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


