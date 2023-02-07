# GetEthCollectionPriceDiffRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The Ethereum contract address to identify the collection. | 
**StartDate** | Pointer to **string** | The start date to pull data for calculations | [optional] 
**EndDate** | Pointer to **string** | The end date to pull data for calculations | [optional] 
**ExcludeWash** | Pointer to **bool** | Exclude suspected wash transactions? | [optional] 

## Methods

### NewGetEthCollectionPriceDiffRequest

`func NewGetEthCollectionPriceDiffRequest(collectionAddress string, ) *GetEthCollectionPriceDiffRequest`

NewGetEthCollectionPriceDiffRequest instantiates a new GetEthCollectionPriceDiffRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthCollectionPriceDiffRequestWithDefaults

`func NewGetEthCollectionPriceDiffRequestWithDefaults() *GetEthCollectionPriceDiffRequest`

NewGetEthCollectionPriceDiffRequestWithDefaults instantiates a new GetEthCollectionPriceDiffRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetEthCollectionPriceDiffRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetEthCollectionPriceDiffRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetEthCollectionPriceDiffRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetStartDate

`func (o *GetEthCollectionPriceDiffRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetEthCollectionPriceDiffRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetEthCollectionPriceDiffRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetEthCollectionPriceDiffRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetEthCollectionPriceDiffRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetEthCollectionPriceDiffRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetEthCollectionPriceDiffRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetEthCollectionPriceDiffRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetExcludeWash

`func (o *GetEthCollectionPriceDiffRequest) GetExcludeWash() bool`

GetExcludeWash returns the ExcludeWash field if non-nil, zero value otherwise.

### GetExcludeWashOk

`func (o *GetEthCollectionPriceDiffRequest) GetExcludeWashOk() (*bool, bool)`

GetExcludeWashOk returns a tuple with the ExcludeWash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeWash

`func (o *GetEthCollectionPriceDiffRequest) SetExcludeWash(v bool)`

SetExcludeWash sets ExcludeWash field to given value.

### HasExcludeWash

`func (o *GetEthCollectionPriceDiffRequest) HasExcludeWash() bool`

HasExcludeWash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


