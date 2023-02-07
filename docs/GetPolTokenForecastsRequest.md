# GetPolTokenForecastsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The contract address of the token collection. | 
**TokenId** | **[]string** | The id(s) for the token(s). | 
**Voltype** | Pointer to **string** | Type of statistical forecasting model to be calculated as a 3-char string, e.g. &#x60;arc&#x60; for ARCH | [optional] 
**Horizon** | Pointer to **int32** | The forecast horizon (i.e. the number of periods to forecast out) | [optional] 
**Frequency** | Pointer to **string** | The interval at which to calculate returns to base the forecasts upon, e.g. &#x60;1D&#x60; for daily, &#x60;1M&#x60; for monthly etc. | [optional] 
**Dist** | Pointer to **string** | The distribution assumed to calculate parametric risk for. | [optional] 
**StartDate** | Pointer to **string** | The start date to pull data for calculations | [optional] 
**EndDate** | Pointer to **string** | The end date to pull data for calculations | [optional] 
**ReturnPriceForecasts** | Pointer to **bool** | Set to True, returns confidencve intervals at alpha significance for price on top of forecasts for returns and volatilities | [optional] 
**Alpha** | Pointer to **float32** | The significance level, e.g. 0.05 for 95% confidence | [optional] 
**ReptCurr** | Pointer to **string** | The currency to report results in | [optional] 
**ExcludeWash** | Pointer to **bool** | Exclude suspected wash transactions? | [optional] 
**ArchParams** | Pointer to [**GetEthCollectionForecastsRequestArchParams**](GetEthCollectionForecastsRequestArchParams.md) |  | [optional] 

## Methods

### NewGetPolTokenForecastsRequest

`func NewGetPolTokenForecastsRequest(collectionAddress string, tokenId []string, ) *GetPolTokenForecastsRequest`

NewGetPolTokenForecastsRequest instantiates a new GetPolTokenForecastsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPolTokenForecastsRequestWithDefaults

`func NewGetPolTokenForecastsRequestWithDefaults() *GetPolTokenForecastsRequest`

NewGetPolTokenForecastsRequestWithDefaults instantiates a new GetPolTokenForecastsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetPolTokenForecastsRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetPolTokenForecastsRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetPolTokenForecastsRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetTokenId

`func (o *GetPolTokenForecastsRequest) GetTokenId() []string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetPolTokenForecastsRequest) GetTokenIdOk() (*[]string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetPolTokenForecastsRequest) SetTokenId(v []string)`

SetTokenId sets TokenId field to given value.


### GetVoltype

`func (o *GetPolTokenForecastsRequest) GetVoltype() string`

GetVoltype returns the Voltype field if non-nil, zero value otherwise.

### GetVoltypeOk

`func (o *GetPolTokenForecastsRequest) GetVoltypeOk() (*string, bool)`

GetVoltypeOk returns a tuple with the Voltype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltype

`func (o *GetPolTokenForecastsRequest) SetVoltype(v string)`

SetVoltype sets Voltype field to given value.

### HasVoltype

`func (o *GetPolTokenForecastsRequest) HasVoltype() bool`

HasVoltype returns a boolean if a field has been set.

### GetHorizon

`func (o *GetPolTokenForecastsRequest) GetHorizon() int32`

GetHorizon returns the Horizon field if non-nil, zero value otherwise.

### GetHorizonOk

`func (o *GetPolTokenForecastsRequest) GetHorizonOk() (*int32, bool)`

GetHorizonOk returns a tuple with the Horizon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizon

`func (o *GetPolTokenForecastsRequest) SetHorizon(v int32)`

SetHorizon sets Horizon field to given value.

### HasHorizon

`func (o *GetPolTokenForecastsRequest) HasHorizon() bool`

HasHorizon returns a boolean if a field has been set.

### GetFrequency

`func (o *GetPolTokenForecastsRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetPolTokenForecastsRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetPolTokenForecastsRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetPolTokenForecastsRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetDist

`func (o *GetPolTokenForecastsRequest) GetDist() string`

GetDist returns the Dist field if non-nil, zero value otherwise.

### GetDistOk

`func (o *GetPolTokenForecastsRequest) GetDistOk() (*string, bool)`

GetDistOk returns a tuple with the Dist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDist

`func (o *GetPolTokenForecastsRequest) SetDist(v string)`

SetDist sets Dist field to given value.

### HasDist

`func (o *GetPolTokenForecastsRequest) HasDist() bool`

HasDist returns a boolean if a field has been set.

### GetStartDate

`func (o *GetPolTokenForecastsRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetPolTokenForecastsRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetPolTokenForecastsRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetPolTokenForecastsRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetPolTokenForecastsRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetPolTokenForecastsRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetPolTokenForecastsRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetPolTokenForecastsRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetReturnPriceForecasts

`func (o *GetPolTokenForecastsRequest) GetReturnPriceForecasts() bool`

GetReturnPriceForecasts returns the ReturnPriceForecasts field if non-nil, zero value otherwise.

### GetReturnPriceForecastsOk

`func (o *GetPolTokenForecastsRequest) GetReturnPriceForecastsOk() (*bool, bool)`

GetReturnPriceForecastsOk returns a tuple with the ReturnPriceForecasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPriceForecasts

`func (o *GetPolTokenForecastsRequest) SetReturnPriceForecasts(v bool)`

SetReturnPriceForecasts sets ReturnPriceForecasts field to given value.

### HasReturnPriceForecasts

`func (o *GetPolTokenForecastsRequest) HasReturnPriceForecasts() bool`

HasReturnPriceForecasts returns a boolean if a field has been set.

### GetAlpha

`func (o *GetPolTokenForecastsRequest) GetAlpha() float32`

GetAlpha returns the Alpha field if non-nil, zero value otherwise.

### GetAlphaOk

`func (o *GetPolTokenForecastsRequest) GetAlphaOk() (*float32, bool)`

GetAlphaOk returns a tuple with the Alpha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpha

`func (o *GetPolTokenForecastsRequest) SetAlpha(v float32)`

SetAlpha sets Alpha field to given value.

### HasAlpha

`func (o *GetPolTokenForecastsRequest) HasAlpha() bool`

HasAlpha returns a boolean if a field has been set.

### GetReptCurr

`func (o *GetPolTokenForecastsRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetPolTokenForecastsRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetPolTokenForecastsRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetPolTokenForecastsRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetExcludeWash

`func (o *GetPolTokenForecastsRequest) GetExcludeWash() bool`

GetExcludeWash returns the ExcludeWash field if non-nil, zero value otherwise.

### GetExcludeWashOk

`func (o *GetPolTokenForecastsRequest) GetExcludeWashOk() (*bool, bool)`

GetExcludeWashOk returns a tuple with the ExcludeWash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeWash

`func (o *GetPolTokenForecastsRequest) SetExcludeWash(v bool)`

SetExcludeWash sets ExcludeWash field to given value.

### HasExcludeWash

`func (o *GetPolTokenForecastsRequest) HasExcludeWash() bool`

HasExcludeWash returns a boolean if a field has been set.

### GetArchParams

`func (o *GetPolTokenForecastsRequest) GetArchParams() GetEthCollectionForecastsRequestArchParams`

GetArchParams returns the ArchParams field if non-nil, zero value otherwise.

### GetArchParamsOk

`func (o *GetPolTokenForecastsRequest) GetArchParamsOk() (*GetEthCollectionForecastsRequestArchParams, bool)`

GetArchParamsOk returns a tuple with the ArchParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchParams

`func (o *GetPolTokenForecastsRequest) SetArchParams(v GetEthCollectionForecastsRequestArchParams)`

SetArchParams sets ArchParams field to given value.

### HasArchParams

`func (o *GetPolTokenForecastsRequest) HasArchParams() bool`

HasArchParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


