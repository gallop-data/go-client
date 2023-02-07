# GetEthHistoricalTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The contract address of the collection. | 
**TokenId** | Pointer to **string** | The id for the token. | [optional] 
**Page** | Pointer to **int32** | The pagination cursor. | [optional] 

## Methods

### NewGetEthHistoricalTransactionsRequest

`func NewGetEthHistoricalTransactionsRequest(collectionAddress string, ) *GetEthHistoricalTransactionsRequest`

NewGetEthHistoricalTransactionsRequest instantiates a new GetEthHistoricalTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthHistoricalTransactionsRequestWithDefaults

`func NewGetEthHistoricalTransactionsRequestWithDefaults() *GetEthHistoricalTransactionsRequest`

NewGetEthHistoricalTransactionsRequestWithDefaults instantiates a new GetEthHistoricalTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetEthHistoricalTransactionsRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetEthHistoricalTransactionsRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetEthHistoricalTransactionsRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetTokenId

`func (o *GetEthHistoricalTransactionsRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetEthHistoricalTransactionsRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetEthHistoricalTransactionsRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *GetEthHistoricalTransactionsRequest) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetPage

`func (o *GetEthHistoricalTransactionsRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetEthHistoricalTransactionsRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetEthHistoricalTransactionsRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetEthHistoricalTransactionsRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


