# GetSolCollectionPriceDiffRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionTag** | **string** | The Gallop tag to identify the collection. | 
**StartDate** | Pointer to **string** | The start date to pull data for calculations | [optional] 
**EndDate** | Pointer to **string** | The end date to pull data for calculations | [optional] 
**ExcludeWash** | Pointer to **bool** | Exclude suspected wash transactions? | [optional] 

## Methods

### NewGetSolCollectionPriceDiffRequest

`func NewGetSolCollectionPriceDiffRequest(collectionTag string, ) *GetSolCollectionPriceDiffRequest`

NewGetSolCollectionPriceDiffRequest instantiates a new GetSolCollectionPriceDiffRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSolCollectionPriceDiffRequestWithDefaults

`func NewGetSolCollectionPriceDiffRequestWithDefaults() *GetSolCollectionPriceDiffRequest`

NewGetSolCollectionPriceDiffRequestWithDefaults instantiates a new GetSolCollectionPriceDiffRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionTag

`func (o *GetSolCollectionPriceDiffRequest) GetCollectionTag() string`

GetCollectionTag returns the CollectionTag field if non-nil, zero value otherwise.

### GetCollectionTagOk

`func (o *GetSolCollectionPriceDiffRequest) GetCollectionTagOk() (*string, bool)`

GetCollectionTagOk returns a tuple with the CollectionTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTag

`func (o *GetSolCollectionPriceDiffRequest) SetCollectionTag(v string)`

SetCollectionTag sets CollectionTag field to given value.


### GetStartDate

`func (o *GetSolCollectionPriceDiffRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetSolCollectionPriceDiffRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetSolCollectionPriceDiffRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetSolCollectionPriceDiffRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetSolCollectionPriceDiffRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetSolCollectionPriceDiffRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetSolCollectionPriceDiffRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetSolCollectionPriceDiffRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetExcludeWash

`func (o *GetSolCollectionPriceDiffRequest) GetExcludeWash() bool`

GetExcludeWash returns the ExcludeWash field if non-nil, zero value otherwise.

### GetExcludeWashOk

`func (o *GetSolCollectionPriceDiffRequest) GetExcludeWashOk() (*bool, bool)`

GetExcludeWashOk returns a tuple with the ExcludeWash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeWash

`func (o *GetSolCollectionPriceDiffRequest) SetExcludeWash(v bool)`

SetExcludeWash sets ExcludeWash field to given value.

### HasExcludeWash

`func (o *GetSolCollectionPriceDiffRequest) HasExcludeWash() bool`

HasExcludeWash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


