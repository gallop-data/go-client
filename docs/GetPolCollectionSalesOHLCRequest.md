# GetPolCollectionSalesOHLCRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The Polygon contract address to identify the collection. | 
**Frequency** | Pointer to **string** | The interval at which to return OHLC, e.g. &#x60;1D&#x60; for daily, &#x60;1M&#x60; for monthly etc. | [optional] 
**GroupBy** | Pointer to **string** | An attribute of the NFT to group results by. | [optional] 
**Volume** | Pointer to **bool** | If &#39;true&#39;, response dicts contain OHLCV | [optional] 
**NTrades** | Pointer to **bool** | If &#39;true&#39;, append number of trades to OHLC(V) | [optional] 
**ReptCurr** | Pointer to **string** | The currency to report results in | [optional] 
**StartDate** | Pointer to **string** | The start date to pull data for calculations | [optional] 
**EndDate** | Pointer to **string** | The end date to pull data for calculations | [optional] 
**ExcludeWash** | Pointer to **bool** | Exclude suspected wash transactions? | [optional] 

## Methods

### NewGetPolCollectionSalesOHLCRequest

`func NewGetPolCollectionSalesOHLCRequest(collectionAddress string, ) *GetPolCollectionSalesOHLCRequest`

NewGetPolCollectionSalesOHLCRequest instantiates a new GetPolCollectionSalesOHLCRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPolCollectionSalesOHLCRequestWithDefaults

`func NewGetPolCollectionSalesOHLCRequestWithDefaults() *GetPolCollectionSalesOHLCRequest`

NewGetPolCollectionSalesOHLCRequestWithDefaults instantiates a new GetPolCollectionSalesOHLCRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetPolCollectionSalesOHLCRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetPolCollectionSalesOHLCRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetPolCollectionSalesOHLCRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetFrequency

`func (o *GetPolCollectionSalesOHLCRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetPolCollectionSalesOHLCRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetPolCollectionSalesOHLCRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetPolCollectionSalesOHLCRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetGroupBy

`func (o *GetPolCollectionSalesOHLCRequest) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *GetPolCollectionSalesOHLCRequest) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *GetPolCollectionSalesOHLCRequest) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *GetPolCollectionSalesOHLCRequest) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetVolume

`func (o *GetPolCollectionSalesOHLCRequest) GetVolume() bool`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *GetPolCollectionSalesOHLCRequest) GetVolumeOk() (*bool, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *GetPolCollectionSalesOHLCRequest) SetVolume(v bool)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *GetPolCollectionSalesOHLCRequest) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetNTrades

`func (o *GetPolCollectionSalesOHLCRequest) GetNTrades() bool`

GetNTrades returns the NTrades field if non-nil, zero value otherwise.

### GetNTradesOk

`func (o *GetPolCollectionSalesOHLCRequest) GetNTradesOk() (*bool, bool)`

GetNTradesOk returns a tuple with the NTrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrades

`func (o *GetPolCollectionSalesOHLCRequest) SetNTrades(v bool)`

SetNTrades sets NTrades field to given value.

### HasNTrades

`func (o *GetPolCollectionSalesOHLCRequest) HasNTrades() bool`

HasNTrades returns a boolean if a field has been set.

### GetReptCurr

`func (o *GetPolCollectionSalesOHLCRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetPolCollectionSalesOHLCRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetPolCollectionSalesOHLCRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetPolCollectionSalesOHLCRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetStartDate

`func (o *GetPolCollectionSalesOHLCRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetPolCollectionSalesOHLCRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetPolCollectionSalesOHLCRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetPolCollectionSalesOHLCRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetPolCollectionSalesOHLCRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetPolCollectionSalesOHLCRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetPolCollectionSalesOHLCRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetPolCollectionSalesOHLCRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetExcludeWash

`func (o *GetPolCollectionSalesOHLCRequest) GetExcludeWash() bool`

GetExcludeWash returns the ExcludeWash field if non-nil, zero value otherwise.

### GetExcludeWashOk

`func (o *GetPolCollectionSalesOHLCRequest) GetExcludeWashOk() (*bool, bool)`

GetExcludeWashOk returns a tuple with the ExcludeWash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeWash

`func (o *GetPolCollectionSalesOHLCRequest) SetExcludeWash(v bool)`

SetExcludeWash sets ExcludeWash field to given value.

### HasExcludeWash

`func (o *GetPolCollectionSalesOHLCRequest) HasExcludeWash() bool`

HasExcludeWash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


