# GetSolTokenForecastsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MintAddress** | **[]string** | A token mint address or list of token addresses. | 
**TokenId** | Pointer to **string** | The numerical id for the token. Provide either id or mint address. | [optional] 
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
**ArchParams** | Pointer to [**GetSolTokenForecastsRequestArchParams**](GetSolTokenForecastsRequestArchParams.md) |  | [optional] 

## Methods

### NewGetSolTokenForecastsRequest

`func NewGetSolTokenForecastsRequest(mintAddress []string, ) *GetSolTokenForecastsRequest`

NewGetSolTokenForecastsRequest instantiates a new GetSolTokenForecastsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSolTokenForecastsRequestWithDefaults

`func NewGetSolTokenForecastsRequestWithDefaults() *GetSolTokenForecastsRequest`

NewGetSolTokenForecastsRequestWithDefaults instantiates a new GetSolTokenForecastsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMintAddress

`func (o *GetSolTokenForecastsRequest) GetMintAddress() []string`

GetMintAddress returns the MintAddress field if non-nil, zero value otherwise.

### GetMintAddressOk

`func (o *GetSolTokenForecastsRequest) GetMintAddressOk() (*[]string, bool)`

GetMintAddressOk returns a tuple with the MintAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintAddress

`func (o *GetSolTokenForecastsRequest) SetMintAddress(v []string)`

SetMintAddress sets MintAddress field to given value.


### GetTokenId

`func (o *GetSolTokenForecastsRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetSolTokenForecastsRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetSolTokenForecastsRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *GetSolTokenForecastsRequest) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetVoltype

`func (o *GetSolTokenForecastsRequest) GetVoltype() string`

GetVoltype returns the Voltype field if non-nil, zero value otherwise.

### GetVoltypeOk

`func (o *GetSolTokenForecastsRequest) GetVoltypeOk() (*string, bool)`

GetVoltypeOk returns a tuple with the Voltype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltype

`func (o *GetSolTokenForecastsRequest) SetVoltype(v string)`

SetVoltype sets Voltype field to given value.

### HasVoltype

`func (o *GetSolTokenForecastsRequest) HasVoltype() bool`

HasVoltype returns a boolean if a field has been set.

### GetHorizon

`func (o *GetSolTokenForecastsRequest) GetHorizon() int32`

GetHorizon returns the Horizon field if non-nil, zero value otherwise.

### GetHorizonOk

`func (o *GetSolTokenForecastsRequest) GetHorizonOk() (*int32, bool)`

GetHorizonOk returns a tuple with the Horizon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizon

`func (o *GetSolTokenForecastsRequest) SetHorizon(v int32)`

SetHorizon sets Horizon field to given value.

### HasHorizon

`func (o *GetSolTokenForecastsRequest) HasHorizon() bool`

HasHorizon returns a boolean if a field has been set.

### GetFrequency

`func (o *GetSolTokenForecastsRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetSolTokenForecastsRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetSolTokenForecastsRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetSolTokenForecastsRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetDist

`func (o *GetSolTokenForecastsRequest) GetDist() string`

GetDist returns the Dist field if non-nil, zero value otherwise.

### GetDistOk

`func (o *GetSolTokenForecastsRequest) GetDistOk() (*string, bool)`

GetDistOk returns a tuple with the Dist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDist

`func (o *GetSolTokenForecastsRequest) SetDist(v string)`

SetDist sets Dist field to given value.

### HasDist

`func (o *GetSolTokenForecastsRequest) HasDist() bool`

HasDist returns a boolean if a field has been set.

### GetStartDate

`func (o *GetSolTokenForecastsRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetSolTokenForecastsRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetSolTokenForecastsRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetSolTokenForecastsRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetSolTokenForecastsRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetSolTokenForecastsRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetSolTokenForecastsRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetSolTokenForecastsRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetReturnPriceForecasts

`func (o *GetSolTokenForecastsRequest) GetReturnPriceForecasts() bool`

GetReturnPriceForecasts returns the ReturnPriceForecasts field if non-nil, zero value otherwise.

### GetReturnPriceForecastsOk

`func (o *GetSolTokenForecastsRequest) GetReturnPriceForecastsOk() (*bool, bool)`

GetReturnPriceForecastsOk returns a tuple with the ReturnPriceForecasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPriceForecasts

`func (o *GetSolTokenForecastsRequest) SetReturnPriceForecasts(v bool)`

SetReturnPriceForecasts sets ReturnPriceForecasts field to given value.

### HasReturnPriceForecasts

`func (o *GetSolTokenForecastsRequest) HasReturnPriceForecasts() bool`

HasReturnPriceForecasts returns a boolean if a field has been set.

### GetAlpha

`func (o *GetSolTokenForecastsRequest) GetAlpha() float32`

GetAlpha returns the Alpha field if non-nil, zero value otherwise.

### GetAlphaOk

`func (o *GetSolTokenForecastsRequest) GetAlphaOk() (*float32, bool)`

GetAlphaOk returns a tuple with the Alpha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpha

`func (o *GetSolTokenForecastsRequest) SetAlpha(v float32)`

SetAlpha sets Alpha field to given value.

### HasAlpha

`func (o *GetSolTokenForecastsRequest) HasAlpha() bool`

HasAlpha returns a boolean if a field has been set.

### GetReptCurr

`func (o *GetSolTokenForecastsRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetSolTokenForecastsRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetSolTokenForecastsRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetSolTokenForecastsRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetExcludeWash

`func (o *GetSolTokenForecastsRequest) GetExcludeWash() bool`

GetExcludeWash returns the ExcludeWash field if non-nil, zero value otherwise.

### GetExcludeWashOk

`func (o *GetSolTokenForecastsRequest) GetExcludeWashOk() (*bool, bool)`

GetExcludeWashOk returns a tuple with the ExcludeWash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeWash

`func (o *GetSolTokenForecastsRequest) SetExcludeWash(v bool)`

SetExcludeWash sets ExcludeWash field to given value.

### HasExcludeWash

`func (o *GetSolTokenForecastsRequest) HasExcludeWash() bool`

HasExcludeWash returns a boolean if a field has been set.

### GetArchParams

`func (o *GetSolTokenForecastsRequest) GetArchParams() GetSolTokenForecastsRequestArchParams`

GetArchParams returns the ArchParams field if non-nil, zero value otherwise.

### GetArchParamsOk

`func (o *GetSolTokenForecastsRequest) GetArchParamsOk() (*GetSolTokenForecastsRequestArchParams, bool)`

GetArchParamsOk returns a tuple with the ArchParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchParams

`func (o *GetSolTokenForecastsRequest) SetArchParams(v GetSolTokenForecastsRequestArchParams)`

SetArchParams sets ArchParams field to given value.

### HasArchParams

`func (o *GetSolTokenForecastsRequest) HasArchParams() bool`

HasArchParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


