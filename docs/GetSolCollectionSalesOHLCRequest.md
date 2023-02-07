# GetSolCollectionSalesOHLCRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionTag** | **string** | The Gallop tag to identify the collection. | 
**Frequency** | Pointer to **string** | The interval at which to return OHLC, e.g. &#x60;1D&#x60; for daily, &#x60;1M&#x60; for monthly etc. | [optional] 
**GroupBy** | Pointer to **string** | An attribute of the NFT to group results by. | [optional] 
**Volume** | Pointer to **bool** | If &#39;true&#39;, response dicts contain OHLCV | [optional] 
**ReptCurr** | Pointer to **string** | The currency to report results in | [optional] 
**StartDate** | Pointer to **string** | The start date to pull data for calculations | [optional] 
**EndDate** | Pointer to **string** | The end date to pull data for calculations | [optional] 
**ExcludeWash** | Pointer to **bool** | Exclude suspected wash transactions? | [optional] 

## Methods

### NewGetSolCollectionSalesOHLCRequest

`func NewGetSolCollectionSalesOHLCRequest(collectionTag string, ) *GetSolCollectionSalesOHLCRequest`

NewGetSolCollectionSalesOHLCRequest instantiates a new GetSolCollectionSalesOHLCRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSolCollectionSalesOHLCRequestWithDefaults

`func NewGetSolCollectionSalesOHLCRequestWithDefaults() *GetSolCollectionSalesOHLCRequest`

NewGetSolCollectionSalesOHLCRequestWithDefaults instantiates a new GetSolCollectionSalesOHLCRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionTag

`func (o *GetSolCollectionSalesOHLCRequest) GetCollectionTag() string`

GetCollectionTag returns the CollectionTag field if non-nil, zero value otherwise.

### GetCollectionTagOk

`func (o *GetSolCollectionSalesOHLCRequest) GetCollectionTagOk() (*string, bool)`

GetCollectionTagOk returns a tuple with the CollectionTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTag

`func (o *GetSolCollectionSalesOHLCRequest) SetCollectionTag(v string)`

SetCollectionTag sets CollectionTag field to given value.


### GetFrequency

`func (o *GetSolCollectionSalesOHLCRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetSolCollectionSalesOHLCRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetSolCollectionSalesOHLCRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetSolCollectionSalesOHLCRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetGroupBy

`func (o *GetSolCollectionSalesOHLCRequest) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *GetSolCollectionSalesOHLCRequest) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *GetSolCollectionSalesOHLCRequest) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *GetSolCollectionSalesOHLCRequest) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetVolume

`func (o *GetSolCollectionSalesOHLCRequest) GetVolume() bool`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *GetSolCollectionSalesOHLCRequest) GetVolumeOk() (*bool, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *GetSolCollectionSalesOHLCRequest) SetVolume(v bool)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *GetSolCollectionSalesOHLCRequest) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetReptCurr

`func (o *GetSolCollectionSalesOHLCRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetSolCollectionSalesOHLCRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetSolCollectionSalesOHLCRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetSolCollectionSalesOHLCRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetStartDate

`func (o *GetSolCollectionSalesOHLCRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetSolCollectionSalesOHLCRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetSolCollectionSalesOHLCRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetSolCollectionSalesOHLCRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetSolCollectionSalesOHLCRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetSolCollectionSalesOHLCRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetSolCollectionSalesOHLCRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetSolCollectionSalesOHLCRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetExcludeWash

`func (o *GetSolCollectionSalesOHLCRequest) GetExcludeWash() bool`

GetExcludeWash returns the ExcludeWash field if non-nil, zero value otherwise.

### GetExcludeWashOk

`func (o *GetSolCollectionSalesOHLCRequest) GetExcludeWashOk() (*bool, bool)`

GetExcludeWashOk returns a tuple with the ExcludeWash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeWash

`func (o *GetSolCollectionSalesOHLCRequest) SetExcludeWash(v bool)`

SetExcludeWash sets ExcludeWash field to given value.

### HasExcludeWash

`func (o *GetSolCollectionSalesOHLCRequest) HasExcludeWash() bool`

HasExcludeWash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


