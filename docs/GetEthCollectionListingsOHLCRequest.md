# GetEthCollectionListingsOHLCRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The Ethereum contract address to identify the collection. | 
**FloorOnly** | Pointer to **bool** | If &#x60;true&#x60;, report only historical floor prices. Otherwise, report OHFC candlesticks, number of active listings, number of unique owners and the average age of open listings. | [optional] 
**Frequency** | Pointer to **string** | The interval at which to return Floor prices / OHLF, e.g. &#x60;1D&#x60; for daily, &#x60;1M&#x60; for monthly etc. Must be &gt;&#x3D; &#x60;6H&#x60; | [optional] 
**ReptCurr** | Pointer to **string** | The currency to report results in | [optional] 
**ReportStartDate** | Pointer to **string** | The ISO 8601 start date/datetime to return results for | [optional] 
**ReportEndDate** | Pointer to **string** | The ISO 8601 end date/datetime to return results for | [optional] 

## Methods

### NewGetEthCollectionListingsOHLCRequest

`func NewGetEthCollectionListingsOHLCRequest(collectionAddress string, ) *GetEthCollectionListingsOHLCRequest`

NewGetEthCollectionListingsOHLCRequest instantiates a new GetEthCollectionListingsOHLCRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthCollectionListingsOHLCRequestWithDefaults

`func NewGetEthCollectionListingsOHLCRequestWithDefaults() *GetEthCollectionListingsOHLCRequest`

NewGetEthCollectionListingsOHLCRequestWithDefaults instantiates a new GetEthCollectionListingsOHLCRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetEthCollectionListingsOHLCRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetEthCollectionListingsOHLCRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetEthCollectionListingsOHLCRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetFloorOnly

`func (o *GetEthCollectionListingsOHLCRequest) GetFloorOnly() bool`

GetFloorOnly returns the FloorOnly field if non-nil, zero value otherwise.

### GetFloorOnlyOk

`func (o *GetEthCollectionListingsOHLCRequest) GetFloorOnlyOk() (*bool, bool)`

GetFloorOnlyOk returns a tuple with the FloorOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorOnly

`func (o *GetEthCollectionListingsOHLCRequest) SetFloorOnly(v bool)`

SetFloorOnly sets FloorOnly field to given value.

### HasFloorOnly

`func (o *GetEthCollectionListingsOHLCRequest) HasFloorOnly() bool`

HasFloorOnly returns a boolean if a field has been set.

### GetFrequency

`func (o *GetEthCollectionListingsOHLCRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetEthCollectionListingsOHLCRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetEthCollectionListingsOHLCRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetEthCollectionListingsOHLCRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetReptCurr

`func (o *GetEthCollectionListingsOHLCRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetEthCollectionListingsOHLCRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetEthCollectionListingsOHLCRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetEthCollectionListingsOHLCRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetReportStartDate

`func (o *GetEthCollectionListingsOHLCRequest) GetReportStartDate() string`

GetReportStartDate returns the ReportStartDate field if non-nil, zero value otherwise.

### GetReportStartDateOk

`func (o *GetEthCollectionListingsOHLCRequest) GetReportStartDateOk() (*string, bool)`

GetReportStartDateOk returns a tuple with the ReportStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportStartDate

`func (o *GetEthCollectionListingsOHLCRequest) SetReportStartDate(v string)`

SetReportStartDate sets ReportStartDate field to given value.

### HasReportStartDate

`func (o *GetEthCollectionListingsOHLCRequest) HasReportStartDate() bool`

HasReportStartDate returns a boolean if a field has been set.

### GetReportEndDate

`func (o *GetEthCollectionListingsOHLCRequest) GetReportEndDate() string`

GetReportEndDate returns the ReportEndDate field if non-nil, zero value otherwise.

### GetReportEndDateOk

`func (o *GetEthCollectionListingsOHLCRequest) GetReportEndDateOk() (*string, bool)`

GetReportEndDateOk returns a tuple with the ReportEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportEndDate

`func (o *GetEthCollectionListingsOHLCRequest) SetReportEndDate(v string)`

SetReportEndDate sets ReportEndDate field to given value.

### HasReportEndDate

`func (o *GetEthCollectionListingsOHLCRequest) HasReportEndDate() bool`

HasReportEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


