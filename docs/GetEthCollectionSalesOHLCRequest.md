# GetEthCollectionSalesOHLCRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The Ethereum contract address to identify the collection. | 
**Frequency** | Pointer to **string** | The interval at which to return OHLC, e.g. &#x60;1D&#x60; for daily, &#x60;1M&#x60; for monthly etc. | [optional] 
**GroupBy** | Pointer to **string** | An attribute of the NFT to group results by. | [optional] 
**Volume** | Pointer to **bool** | If &#39;true&#39;, response dicts contain OHLCV | [optional] 
**NTrades** | Pointer to **bool** | If &#39;true&#39;, append number of trades to OHLC(V) | [optional] 
**ReptCurr** | Pointer to **string** | The currency to report results in | [optional] 
**StartDate** | Pointer to **string** | The start date to pull data for calculations | [optional] 
**EndDate** | Pointer to **string** | The end date to pull data for calculations | [optional] 
**ExcludeWash** | Pointer to **bool** | Exclude suspected wash transactions? | [optional] 

## Methods

### NewGetEthCollectionSalesOHLCRequest

`func NewGetEthCollectionSalesOHLCRequest(collectionAddress string, ) *GetEthCollectionSalesOHLCRequest`

NewGetEthCollectionSalesOHLCRequest instantiates a new GetEthCollectionSalesOHLCRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthCollectionSalesOHLCRequestWithDefaults

`func NewGetEthCollectionSalesOHLCRequestWithDefaults() *GetEthCollectionSalesOHLCRequest`

NewGetEthCollectionSalesOHLCRequestWithDefaults instantiates a new GetEthCollectionSalesOHLCRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetEthCollectionSalesOHLCRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetEthCollectionSalesOHLCRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetEthCollectionSalesOHLCRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetFrequency

`func (o *GetEthCollectionSalesOHLCRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetEthCollectionSalesOHLCRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetEthCollectionSalesOHLCRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetEthCollectionSalesOHLCRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetGroupBy

`func (o *GetEthCollectionSalesOHLCRequest) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *GetEthCollectionSalesOHLCRequest) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *GetEthCollectionSalesOHLCRequest) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *GetEthCollectionSalesOHLCRequest) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetVolume

`func (o *GetEthCollectionSalesOHLCRequest) GetVolume() bool`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *GetEthCollectionSalesOHLCRequest) GetVolumeOk() (*bool, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *GetEthCollectionSalesOHLCRequest) SetVolume(v bool)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *GetEthCollectionSalesOHLCRequest) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetNTrades

`func (o *GetEthCollectionSalesOHLCRequest) GetNTrades() bool`

GetNTrades returns the NTrades field if non-nil, zero value otherwise.

### GetNTradesOk

`func (o *GetEthCollectionSalesOHLCRequest) GetNTradesOk() (*bool, bool)`

GetNTradesOk returns a tuple with the NTrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrades

`func (o *GetEthCollectionSalesOHLCRequest) SetNTrades(v bool)`

SetNTrades sets NTrades field to given value.

### HasNTrades

`func (o *GetEthCollectionSalesOHLCRequest) HasNTrades() bool`

HasNTrades returns a boolean if a field has been set.

### GetReptCurr

`func (o *GetEthCollectionSalesOHLCRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetEthCollectionSalesOHLCRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetEthCollectionSalesOHLCRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetEthCollectionSalesOHLCRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetStartDate

`func (o *GetEthCollectionSalesOHLCRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetEthCollectionSalesOHLCRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetEthCollectionSalesOHLCRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetEthCollectionSalesOHLCRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetEthCollectionSalesOHLCRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetEthCollectionSalesOHLCRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetEthCollectionSalesOHLCRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetEthCollectionSalesOHLCRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetExcludeWash

`func (o *GetEthCollectionSalesOHLCRequest) GetExcludeWash() bool`

GetExcludeWash returns the ExcludeWash field if non-nil, zero value otherwise.

### GetExcludeWashOk

`func (o *GetEthCollectionSalesOHLCRequest) GetExcludeWashOk() (*bool, bool)`

GetExcludeWashOk returns a tuple with the ExcludeWash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeWash

`func (o *GetEthCollectionSalesOHLCRequest) SetExcludeWash(v bool)`

SetExcludeWash sets ExcludeWash field to given value.

### HasExcludeWash

`func (o *GetEthCollectionSalesOHLCRequest) HasExcludeWash() bool`

HasExcludeWash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


