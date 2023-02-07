# GetSolHistoricalTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionTag** | **string** | The Gallop slug for the collection. Please see sol/getCollections endpoint. | 
**Page** | Pointer to **int32** | The pagination cursor. | [optional] 

## Methods

### NewGetSolHistoricalTransactionsRequest

`func NewGetSolHistoricalTransactionsRequest(collectionTag string, ) *GetSolHistoricalTransactionsRequest`

NewGetSolHistoricalTransactionsRequest instantiates a new GetSolHistoricalTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSolHistoricalTransactionsRequestWithDefaults

`func NewGetSolHistoricalTransactionsRequestWithDefaults() *GetSolHistoricalTransactionsRequest`

NewGetSolHistoricalTransactionsRequestWithDefaults instantiates a new GetSolHistoricalTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionTag

`func (o *GetSolHistoricalTransactionsRequest) GetCollectionTag() string`

GetCollectionTag returns the CollectionTag field if non-nil, zero value otherwise.

### GetCollectionTagOk

`func (o *GetSolHistoricalTransactionsRequest) GetCollectionTagOk() (*string, bool)`

GetCollectionTagOk returns a tuple with the CollectionTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTag

`func (o *GetSolHistoricalTransactionsRequest) SetCollectionTag(v string)`

SetCollectionTag sets CollectionTag field to given value.


### GetPage

`func (o *GetSolHistoricalTransactionsRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetSolHistoricalTransactionsRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetSolHistoricalTransactionsRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetSolHistoricalTransactionsRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


