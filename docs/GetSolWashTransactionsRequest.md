# GetSolWashTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionTag** | **string** | The Gallop slug for the collection. Please see sol/getCollections endpoint. | 
**MintAddress** | Pointer to **[]string** | An optional list of token addresses. | [optional] 
**Page** | Pointer to **int32** | The pagination cursor. | [optional] 
**PageSize** | Pointer to **int32** | The number of records returned per page. | [optional] 

## Methods

### NewGetSolWashTransactionsRequest

`func NewGetSolWashTransactionsRequest(collectionTag string, ) *GetSolWashTransactionsRequest`

NewGetSolWashTransactionsRequest instantiates a new GetSolWashTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSolWashTransactionsRequestWithDefaults

`func NewGetSolWashTransactionsRequestWithDefaults() *GetSolWashTransactionsRequest`

NewGetSolWashTransactionsRequestWithDefaults instantiates a new GetSolWashTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionTag

`func (o *GetSolWashTransactionsRequest) GetCollectionTag() string`

GetCollectionTag returns the CollectionTag field if non-nil, zero value otherwise.

### GetCollectionTagOk

`func (o *GetSolWashTransactionsRequest) GetCollectionTagOk() (*string, bool)`

GetCollectionTagOk returns a tuple with the CollectionTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTag

`func (o *GetSolWashTransactionsRequest) SetCollectionTag(v string)`

SetCollectionTag sets CollectionTag field to given value.


### GetMintAddress

`func (o *GetSolWashTransactionsRequest) GetMintAddress() []string`

GetMintAddress returns the MintAddress field if non-nil, zero value otherwise.

### GetMintAddressOk

`func (o *GetSolWashTransactionsRequest) GetMintAddressOk() (*[]string, bool)`

GetMintAddressOk returns a tuple with the MintAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintAddress

`func (o *GetSolWashTransactionsRequest) SetMintAddress(v []string)`

SetMintAddress sets MintAddress field to given value.

### HasMintAddress

`func (o *GetSolWashTransactionsRequest) HasMintAddress() bool`

HasMintAddress returns a boolean if a field has been set.

### GetPage

`func (o *GetSolWashTransactionsRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetSolWashTransactionsRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetSolWashTransactionsRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetSolWashTransactionsRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *GetSolWashTransactionsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetSolWashTransactionsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetSolWashTransactionsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetSolWashTransactionsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


