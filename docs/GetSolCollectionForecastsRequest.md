# GetSolCollectionForecastsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionTag** | **string** | The Gallop tag to identify the collection. | 
**Percentiles** | Pointer to **[]int32** | The collection percentile(s) | [optional] 
**Voltype** | Pointer to **string** | Type of statistical forecasting model to be calculated as a 3-char string, e.g. &#39;arc&#39; for ARCH | [optional] 
**Horizon** | Pointer to **int32** | The forecast horizon (i.e. the number of periods to forecast out) | [optional] 
**Frequency** | Pointer to **string** | The interval at which to calculate returns to base the forecasts upon, e.g. &#x60;1D&#x60; for daily, &#x60;1M&#x60; for monthly etc. | [optional] 
**Dist** | Pointer to **string** | The distribution assumed to calculate parametric risk for | [optional] 
**StartDate** | Pointer to **string** | The start date to pull data for calculations | [optional] 
**EndDate** | Pointer to **string** | The end date to pull data for calculations | [optional] 
**ReturnPriceForecasts** | Pointer to **bool** | Set to True, returns confidencve intervals at alpha significance for price on top of forecasts for returns and volatilities | [optional] 
**Alpha** | Pointer to **float32** | The significance level, e.g. 0.05 for 95% confidence | [optional] 
**ReptCurr** | Pointer to **string** | The currency to report results in | [optional] 
**ExcludeWash** | Pointer to **bool** | Exclude suspected wash transactions? | [optional] 
**ArchParams** | Pointer to [**GetSolCollectionForecastsRequestArchParams**](GetSolCollectionForecastsRequestArchParams.md) |  | [optional] 

## Methods

### NewGetSolCollectionForecastsRequest

`func NewGetSolCollectionForecastsRequest(collectionTag string, ) *GetSolCollectionForecastsRequest`

NewGetSolCollectionForecastsRequest instantiates a new GetSolCollectionForecastsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSolCollectionForecastsRequestWithDefaults

`func NewGetSolCollectionForecastsRequestWithDefaults() *GetSolCollectionForecastsRequest`

NewGetSolCollectionForecastsRequestWithDefaults instantiates a new GetSolCollectionForecastsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionTag

`func (o *GetSolCollectionForecastsRequest) GetCollectionTag() string`

GetCollectionTag returns the CollectionTag field if non-nil, zero value otherwise.

### GetCollectionTagOk

`func (o *GetSolCollectionForecastsRequest) GetCollectionTagOk() (*string, bool)`

GetCollectionTagOk returns a tuple with the CollectionTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTag

`func (o *GetSolCollectionForecastsRequest) SetCollectionTag(v string)`

SetCollectionTag sets CollectionTag field to given value.


### GetPercentiles

`func (o *GetSolCollectionForecastsRequest) GetPercentiles() []int32`

GetPercentiles returns the Percentiles field if non-nil, zero value otherwise.

### GetPercentilesOk

`func (o *GetSolCollectionForecastsRequest) GetPercentilesOk() (*[]int32, bool)`

GetPercentilesOk returns a tuple with the Percentiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentiles

`func (o *GetSolCollectionForecastsRequest) SetPercentiles(v []int32)`

SetPercentiles sets Percentiles field to given value.

### HasPercentiles

`func (o *GetSolCollectionForecastsRequest) HasPercentiles() bool`

HasPercentiles returns a boolean if a field has been set.

### GetVoltype

`func (o *GetSolCollectionForecastsRequest) GetVoltype() string`

GetVoltype returns the Voltype field if non-nil, zero value otherwise.

### GetVoltypeOk

`func (o *GetSolCollectionForecastsRequest) GetVoltypeOk() (*string, bool)`

GetVoltypeOk returns a tuple with the Voltype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltype

`func (o *GetSolCollectionForecastsRequest) SetVoltype(v string)`

SetVoltype sets Voltype field to given value.

### HasVoltype

`func (o *GetSolCollectionForecastsRequest) HasVoltype() bool`

HasVoltype returns a boolean if a field has been set.

### GetHorizon

`func (o *GetSolCollectionForecastsRequest) GetHorizon() int32`

GetHorizon returns the Horizon field if non-nil, zero value otherwise.

### GetHorizonOk

`func (o *GetSolCollectionForecastsRequest) GetHorizonOk() (*int32, bool)`

GetHorizonOk returns a tuple with the Horizon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizon

`func (o *GetSolCollectionForecastsRequest) SetHorizon(v int32)`

SetHorizon sets Horizon field to given value.

### HasHorizon

`func (o *GetSolCollectionForecastsRequest) HasHorizon() bool`

HasHorizon returns a boolean if a field has been set.

### GetFrequency

`func (o *GetSolCollectionForecastsRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetSolCollectionForecastsRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetSolCollectionForecastsRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetSolCollectionForecastsRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetDist

`func (o *GetSolCollectionForecastsRequest) GetDist() string`

GetDist returns the Dist field if non-nil, zero value otherwise.

### GetDistOk

`func (o *GetSolCollectionForecastsRequest) GetDistOk() (*string, bool)`

GetDistOk returns a tuple with the Dist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDist

`func (o *GetSolCollectionForecastsRequest) SetDist(v string)`

SetDist sets Dist field to given value.

### HasDist

`func (o *GetSolCollectionForecastsRequest) HasDist() bool`

HasDist returns a boolean if a field has been set.

### GetStartDate

`func (o *GetSolCollectionForecastsRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetSolCollectionForecastsRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetSolCollectionForecastsRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetSolCollectionForecastsRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetSolCollectionForecastsRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetSolCollectionForecastsRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetSolCollectionForecastsRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetSolCollectionForecastsRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetReturnPriceForecasts

`func (o *GetSolCollectionForecastsRequest) GetReturnPriceForecasts() bool`

GetReturnPriceForecasts returns the ReturnPriceForecasts field if non-nil, zero value otherwise.

### GetReturnPriceForecastsOk

`func (o *GetSolCollectionForecastsRequest) GetReturnPriceForecastsOk() (*bool, bool)`

GetReturnPriceForecastsOk returns a tuple with the ReturnPriceForecasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPriceForecasts

`func (o *GetSolCollectionForecastsRequest) SetReturnPriceForecasts(v bool)`

SetReturnPriceForecasts sets ReturnPriceForecasts field to given value.

### HasReturnPriceForecasts

`func (o *GetSolCollectionForecastsRequest) HasReturnPriceForecasts() bool`

HasReturnPriceForecasts returns a boolean if a field has been set.

### GetAlpha

`func (o *GetSolCollectionForecastsRequest) GetAlpha() float32`

GetAlpha returns the Alpha field if non-nil, zero value otherwise.

### GetAlphaOk

`func (o *GetSolCollectionForecastsRequest) GetAlphaOk() (*float32, bool)`

GetAlphaOk returns a tuple with the Alpha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpha

`func (o *GetSolCollectionForecastsRequest) SetAlpha(v float32)`

SetAlpha sets Alpha field to given value.

### HasAlpha

`func (o *GetSolCollectionForecastsRequest) HasAlpha() bool`

HasAlpha returns a boolean if a field has been set.

### GetReptCurr

`func (o *GetSolCollectionForecastsRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetSolCollectionForecastsRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetSolCollectionForecastsRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetSolCollectionForecastsRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetExcludeWash

`func (o *GetSolCollectionForecastsRequest) GetExcludeWash() bool`

GetExcludeWash returns the ExcludeWash field if non-nil, zero value otherwise.

### GetExcludeWashOk

`func (o *GetSolCollectionForecastsRequest) GetExcludeWashOk() (*bool, bool)`

GetExcludeWashOk returns a tuple with the ExcludeWash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeWash

`func (o *GetSolCollectionForecastsRequest) SetExcludeWash(v bool)`

SetExcludeWash sets ExcludeWash field to given value.

### HasExcludeWash

`func (o *GetSolCollectionForecastsRequest) HasExcludeWash() bool`

HasExcludeWash returns a boolean if a field has been set.

### GetArchParams

`func (o *GetSolCollectionForecastsRequest) GetArchParams() GetSolCollectionForecastsRequestArchParams`

GetArchParams returns the ArchParams field if non-nil, zero value otherwise.

### GetArchParamsOk

`func (o *GetSolCollectionForecastsRequest) GetArchParamsOk() (*GetSolCollectionForecastsRequestArchParams, bool)`

GetArchParamsOk returns a tuple with the ArchParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchParams

`func (o *GetSolCollectionForecastsRequest) SetArchParams(v GetSolCollectionForecastsRequestArchParams)`

SetArchParams sets ArchParams field to given value.

### HasArchParams

`func (o *GetSolCollectionForecastsRequest) HasArchParams() bool`

HasArchParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


