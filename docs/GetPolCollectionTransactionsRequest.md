# GetPolCollectionTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The contract address of the collection. | 
**Page** | Pointer to **int32** | The pagination cursor. | [optional] 
**PageSize** | Pointer to **int32** | The number of records returned per page. | [optional] 
**StartBlockNumber** | Pointer to **int32** | The oldest block number to return. | [optional] 
**StartDate** | Pointer to **string** | The earliest block timestamp. | [optional] 
**EndDate** | Pointer to **string** | The latest block timestamp. | [optional] 

## Methods

### NewGetPolCollectionTransactionsRequest

`func NewGetPolCollectionTransactionsRequest(collectionAddress string, ) *GetPolCollectionTransactionsRequest`

NewGetPolCollectionTransactionsRequest instantiates a new GetPolCollectionTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPolCollectionTransactionsRequestWithDefaults

`func NewGetPolCollectionTransactionsRequestWithDefaults() *GetPolCollectionTransactionsRequest`

NewGetPolCollectionTransactionsRequestWithDefaults instantiates a new GetPolCollectionTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetPolCollectionTransactionsRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetPolCollectionTransactionsRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetPolCollectionTransactionsRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetPage

`func (o *GetPolCollectionTransactionsRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetPolCollectionTransactionsRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetPolCollectionTransactionsRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetPolCollectionTransactionsRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *GetPolCollectionTransactionsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetPolCollectionTransactionsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetPolCollectionTransactionsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetPolCollectionTransactionsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetStartBlockNumber

`func (o *GetPolCollectionTransactionsRequest) GetStartBlockNumber() int32`

GetStartBlockNumber returns the StartBlockNumber field if non-nil, zero value otherwise.

### GetStartBlockNumberOk

`func (o *GetPolCollectionTransactionsRequest) GetStartBlockNumberOk() (*int32, bool)`

GetStartBlockNumberOk returns a tuple with the StartBlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartBlockNumber

`func (o *GetPolCollectionTransactionsRequest) SetStartBlockNumber(v int32)`

SetStartBlockNumber sets StartBlockNumber field to given value.

### HasStartBlockNumber

`func (o *GetPolCollectionTransactionsRequest) HasStartBlockNumber() bool`

HasStartBlockNumber returns a boolean if a field has been set.

### GetStartDate

`func (o *GetPolCollectionTransactionsRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetPolCollectionTransactionsRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetPolCollectionTransactionsRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetPolCollectionTransactionsRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetPolCollectionTransactionsRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetPolCollectionTransactionsRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetPolCollectionTransactionsRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetPolCollectionTransactionsRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


