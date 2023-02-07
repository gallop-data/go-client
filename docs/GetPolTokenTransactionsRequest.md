# GetPolTokenTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The contract address the token belongs to. | 
**TokenId** | **string** | The token id. | 
**Page** | Pointer to **int32** | The pagination cursor. | [optional] 
**PageSize** | Pointer to **int32** | The number of records returned per page. | [optional] 
**StartDate** | Pointer to **string** | The earliest block timestamp. | [optional] 
**StartBlockNumber** | Pointer to **int32** | The oldest block number to return. | [optional] 
**EndDate** | Pointer to **string** | The latest block timestamp. | [optional] 

## Methods

### NewGetPolTokenTransactionsRequest

`func NewGetPolTokenTransactionsRequest(collectionAddress string, tokenId string, ) *GetPolTokenTransactionsRequest`

NewGetPolTokenTransactionsRequest instantiates a new GetPolTokenTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPolTokenTransactionsRequestWithDefaults

`func NewGetPolTokenTransactionsRequestWithDefaults() *GetPolTokenTransactionsRequest`

NewGetPolTokenTransactionsRequestWithDefaults instantiates a new GetPolTokenTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetPolTokenTransactionsRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetPolTokenTransactionsRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetPolTokenTransactionsRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetTokenId

`func (o *GetPolTokenTransactionsRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetPolTokenTransactionsRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetPolTokenTransactionsRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetPage

`func (o *GetPolTokenTransactionsRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetPolTokenTransactionsRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetPolTokenTransactionsRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetPolTokenTransactionsRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *GetPolTokenTransactionsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetPolTokenTransactionsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetPolTokenTransactionsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetPolTokenTransactionsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetStartDate

`func (o *GetPolTokenTransactionsRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetPolTokenTransactionsRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetPolTokenTransactionsRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetPolTokenTransactionsRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartBlockNumber

`func (o *GetPolTokenTransactionsRequest) GetStartBlockNumber() int32`

GetStartBlockNumber returns the StartBlockNumber field if non-nil, zero value otherwise.

### GetStartBlockNumberOk

`func (o *GetPolTokenTransactionsRequest) GetStartBlockNumberOk() (*int32, bool)`

GetStartBlockNumberOk returns a tuple with the StartBlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartBlockNumber

`func (o *GetPolTokenTransactionsRequest) SetStartBlockNumber(v int32)`

SetStartBlockNumber sets StartBlockNumber field to given value.

### HasStartBlockNumber

`func (o *GetPolTokenTransactionsRequest) HasStartBlockNumber() bool`

HasStartBlockNumber returns a boolean if a field has been set.

### GetEndDate

`func (o *GetPolTokenTransactionsRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetPolTokenTransactionsRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetPolTokenTransactionsRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetPolTokenTransactionsRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


