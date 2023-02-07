# GetEthLeaderBoardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | The pagination cursor. | [optional] 
**PageSize** | Pointer to **int32** | The number of records returned per page. | [optional] 
**Interval** | **string** | The requested time interval | 
**RankingMetric** | **string** | The requested calculation metric | 

## Methods

### NewGetEthLeaderBoardRequest

`func NewGetEthLeaderBoardRequest(interval string, rankingMetric string, ) *GetEthLeaderBoardRequest`

NewGetEthLeaderBoardRequest instantiates a new GetEthLeaderBoardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthLeaderBoardRequestWithDefaults

`func NewGetEthLeaderBoardRequestWithDefaults() *GetEthLeaderBoardRequest`

NewGetEthLeaderBoardRequestWithDefaults instantiates a new GetEthLeaderBoardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetEthLeaderBoardRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetEthLeaderBoardRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetEthLeaderBoardRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetEthLeaderBoardRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *GetEthLeaderBoardRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetEthLeaderBoardRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetEthLeaderBoardRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetEthLeaderBoardRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetInterval

`func (o *GetEthLeaderBoardRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GetEthLeaderBoardRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GetEthLeaderBoardRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetRankingMetric

`func (o *GetEthLeaderBoardRequest) GetRankingMetric() string`

GetRankingMetric returns the RankingMetric field if non-nil, zero value otherwise.

### GetRankingMetricOk

`func (o *GetEthLeaderBoardRequest) GetRankingMetricOk() (*string, bool)`

GetRankingMetricOk returns a tuple with the RankingMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankingMetric

`func (o *GetEthLeaderBoardRequest) SetRankingMetric(v string)`

SetRankingMetric sets RankingMetric field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


