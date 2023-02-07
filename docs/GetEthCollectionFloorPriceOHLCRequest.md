# GetEthCollectionFloorPriceOHLCRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The Ethereum contract address to identify the collection. | 
**Frequency** | Pointer to **string** | The interval at which to return OHLC, e.g. &#x60;1D&#x60; for daily, &#x60;1M&#x60; for monthly etc. | [optional] 
**GroupBy** | Pointer to **string** | An attribute of the NFT to group results by. | [optional] 
**StartDate** | Pointer to **string** | The start date to pull data for calculations | [optional] 
**EndDate** | Pointer to **string** | The end date to pull data for calculations | [optional] 

## Methods

### NewGetEthCollectionFloorPriceOHLCRequest

`func NewGetEthCollectionFloorPriceOHLCRequest(collectionAddress string, ) *GetEthCollectionFloorPriceOHLCRequest`

NewGetEthCollectionFloorPriceOHLCRequest instantiates a new GetEthCollectionFloorPriceOHLCRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthCollectionFloorPriceOHLCRequestWithDefaults

`func NewGetEthCollectionFloorPriceOHLCRequestWithDefaults() *GetEthCollectionFloorPriceOHLCRequest`

NewGetEthCollectionFloorPriceOHLCRequestWithDefaults instantiates a new GetEthCollectionFloorPriceOHLCRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetEthCollectionFloorPriceOHLCRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetEthCollectionFloorPriceOHLCRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetEthCollectionFloorPriceOHLCRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetFrequency

`func (o *GetEthCollectionFloorPriceOHLCRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetEthCollectionFloorPriceOHLCRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetEthCollectionFloorPriceOHLCRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetEthCollectionFloorPriceOHLCRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetGroupBy

`func (o *GetEthCollectionFloorPriceOHLCRequest) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *GetEthCollectionFloorPriceOHLCRequest) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *GetEthCollectionFloorPriceOHLCRequest) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *GetEthCollectionFloorPriceOHLCRequest) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetStartDate

`func (o *GetEthCollectionFloorPriceOHLCRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetEthCollectionFloorPriceOHLCRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetEthCollectionFloorPriceOHLCRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetEthCollectionFloorPriceOHLCRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetEthCollectionFloorPriceOHLCRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetEthCollectionFloorPriceOHLCRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetEthCollectionFloorPriceOHLCRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetEthCollectionFloorPriceOHLCRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


