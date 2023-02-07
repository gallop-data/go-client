# GetPolCustomTokenRiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The contract address of the token collection. | 
**TokenId** | **[]string** | The id(s) for the token(s). | 
**HoldingPeriod** | **string** | The holding period to evaluate risk for, e.g. &#x60;12M&#x60; | 
**Dist** | Pointer to **string** | Return distribution assumed. | [optional] 
**StartDate** | Pointer to **string** | The start date to pull data for calculations | [optional] 
**EndDate** | Pointer to **string** | The end date to pull data for calculations | [optional] 
**RiskFreeRate** | Pointer to **float32** | The rate of return for an asset deemed risk free in the contemplated holding period | [optional] 
**WinsOutliers** | Pointer to **bool** | Whether to winsorize time series outliers prior to calculating risk | [optional] 
**Frequency** | Pointer to **string** | The interval at which to calculate returns to base the forecasts upon, e.g. &#x60;1D&#x60; for daily, &#x60;1M&#x60; for monthly etc. | [optional] 
**ReptCurr** | Pointer to **string** | The currency to report results in | [optional] 
**ExcludeWash** | Pointer to **bool** | Exclude suspected wash transactions? | [optional] 
**Drawdown** | Pointer to **bool** | If true, report drawdown volatility (based on negative returns only). | [optional] 
**ArcParams** | Pointer to [**GetEthCustomCollectionRiskRequestArcParams**](GetEthCustomCollectionRiskRequestArcParams.md) |  | [optional] 
**GarParams** | Pointer to [**GetEthCustomCollectionRiskRequestGarParams**](GetEthCustomCollectionRiskRequestGarParams.md) |  | [optional] 
**HarParams** | Pointer to [**GetEthCustomCollectionRiskRequestHarParams**](GetEthCustomCollectionRiskRequestHarParams.md) |  | [optional] 

## Methods

### NewGetPolCustomTokenRiskRequest

`func NewGetPolCustomTokenRiskRequest(collectionAddress string, tokenId []string, holdingPeriod string, ) *GetPolCustomTokenRiskRequest`

NewGetPolCustomTokenRiskRequest instantiates a new GetPolCustomTokenRiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPolCustomTokenRiskRequestWithDefaults

`func NewGetPolCustomTokenRiskRequestWithDefaults() *GetPolCustomTokenRiskRequest`

NewGetPolCustomTokenRiskRequestWithDefaults instantiates a new GetPolCustomTokenRiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetPolCustomTokenRiskRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetPolCustomTokenRiskRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetPolCustomTokenRiskRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetTokenId

`func (o *GetPolCustomTokenRiskRequest) GetTokenId() []string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetPolCustomTokenRiskRequest) GetTokenIdOk() (*[]string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetPolCustomTokenRiskRequest) SetTokenId(v []string)`

SetTokenId sets TokenId field to given value.


### GetHoldingPeriod

`func (o *GetPolCustomTokenRiskRequest) GetHoldingPeriod() string`

GetHoldingPeriod returns the HoldingPeriod field if non-nil, zero value otherwise.

### GetHoldingPeriodOk

`func (o *GetPolCustomTokenRiskRequest) GetHoldingPeriodOk() (*string, bool)`

GetHoldingPeriodOk returns a tuple with the HoldingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingPeriod

`func (o *GetPolCustomTokenRiskRequest) SetHoldingPeriod(v string)`

SetHoldingPeriod sets HoldingPeriod field to given value.


### GetDist

`func (o *GetPolCustomTokenRiskRequest) GetDist() string`

GetDist returns the Dist field if non-nil, zero value otherwise.

### GetDistOk

`func (o *GetPolCustomTokenRiskRequest) GetDistOk() (*string, bool)`

GetDistOk returns a tuple with the Dist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDist

`func (o *GetPolCustomTokenRiskRequest) SetDist(v string)`

SetDist sets Dist field to given value.

### HasDist

`func (o *GetPolCustomTokenRiskRequest) HasDist() bool`

HasDist returns a boolean if a field has been set.

### GetStartDate

`func (o *GetPolCustomTokenRiskRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetPolCustomTokenRiskRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetPolCustomTokenRiskRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetPolCustomTokenRiskRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetPolCustomTokenRiskRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetPolCustomTokenRiskRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetPolCustomTokenRiskRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetPolCustomTokenRiskRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetRiskFreeRate

`func (o *GetPolCustomTokenRiskRequest) GetRiskFreeRate() float32`

GetRiskFreeRate returns the RiskFreeRate field if non-nil, zero value otherwise.

### GetRiskFreeRateOk

`func (o *GetPolCustomTokenRiskRequest) GetRiskFreeRateOk() (*float32, bool)`

GetRiskFreeRateOk returns a tuple with the RiskFreeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskFreeRate

`func (o *GetPolCustomTokenRiskRequest) SetRiskFreeRate(v float32)`

SetRiskFreeRate sets RiskFreeRate field to given value.

### HasRiskFreeRate

`func (o *GetPolCustomTokenRiskRequest) HasRiskFreeRate() bool`

HasRiskFreeRate returns a boolean if a field has been set.

### GetWinsOutliers

`func (o *GetPolCustomTokenRiskRequest) GetWinsOutliers() bool`

GetWinsOutliers returns the WinsOutliers field if non-nil, zero value otherwise.

### GetWinsOutliersOk

`func (o *GetPolCustomTokenRiskRequest) GetWinsOutliersOk() (*bool, bool)`

GetWinsOutliersOk returns a tuple with the WinsOutliers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsOutliers

`func (o *GetPolCustomTokenRiskRequest) SetWinsOutliers(v bool)`

SetWinsOutliers sets WinsOutliers field to given value.

### HasWinsOutliers

`func (o *GetPolCustomTokenRiskRequest) HasWinsOutliers() bool`

HasWinsOutliers returns a boolean if a field has been set.

### GetFrequency

`func (o *GetPolCustomTokenRiskRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetPolCustomTokenRiskRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetPolCustomTokenRiskRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetPolCustomTokenRiskRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetReptCurr

`func (o *GetPolCustomTokenRiskRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetPolCustomTokenRiskRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetPolCustomTokenRiskRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetPolCustomTokenRiskRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetExcludeWash

`func (o *GetPolCustomTokenRiskRequest) GetExcludeWash() bool`

GetExcludeWash returns the ExcludeWash field if non-nil, zero value otherwise.

### GetExcludeWashOk

`func (o *GetPolCustomTokenRiskRequest) GetExcludeWashOk() (*bool, bool)`

GetExcludeWashOk returns a tuple with the ExcludeWash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeWash

`func (o *GetPolCustomTokenRiskRequest) SetExcludeWash(v bool)`

SetExcludeWash sets ExcludeWash field to given value.

### HasExcludeWash

`func (o *GetPolCustomTokenRiskRequest) HasExcludeWash() bool`

HasExcludeWash returns a boolean if a field has been set.

### GetDrawdown

`func (o *GetPolCustomTokenRiskRequest) GetDrawdown() bool`

GetDrawdown returns the Drawdown field if non-nil, zero value otherwise.

### GetDrawdownOk

`func (o *GetPolCustomTokenRiskRequest) GetDrawdownOk() (*bool, bool)`

GetDrawdownOk returns a tuple with the Drawdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawdown

`func (o *GetPolCustomTokenRiskRequest) SetDrawdown(v bool)`

SetDrawdown sets Drawdown field to given value.

### HasDrawdown

`func (o *GetPolCustomTokenRiskRequest) HasDrawdown() bool`

HasDrawdown returns a boolean if a field has been set.

### GetArcParams

`func (o *GetPolCustomTokenRiskRequest) GetArcParams() GetEthCustomCollectionRiskRequestArcParams`

GetArcParams returns the ArcParams field if non-nil, zero value otherwise.

### GetArcParamsOk

`func (o *GetPolCustomTokenRiskRequest) GetArcParamsOk() (*GetEthCustomCollectionRiskRequestArcParams, bool)`

GetArcParamsOk returns a tuple with the ArcParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcParams

`func (o *GetPolCustomTokenRiskRequest) SetArcParams(v GetEthCustomCollectionRiskRequestArcParams)`

SetArcParams sets ArcParams field to given value.

### HasArcParams

`func (o *GetPolCustomTokenRiskRequest) HasArcParams() bool`

HasArcParams returns a boolean if a field has been set.

### GetGarParams

`func (o *GetPolCustomTokenRiskRequest) GetGarParams() GetEthCustomCollectionRiskRequestGarParams`

GetGarParams returns the GarParams field if non-nil, zero value otherwise.

### GetGarParamsOk

`func (o *GetPolCustomTokenRiskRequest) GetGarParamsOk() (*GetEthCustomCollectionRiskRequestGarParams, bool)`

GetGarParamsOk returns a tuple with the GarParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarParams

`func (o *GetPolCustomTokenRiskRequest) SetGarParams(v GetEthCustomCollectionRiskRequestGarParams)`

SetGarParams sets GarParams field to given value.

### HasGarParams

`func (o *GetPolCustomTokenRiskRequest) HasGarParams() bool`

HasGarParams returns a boolean if a field has been set.

### GetHarParams

`func (o *GetPolCustomTokenRiskRequest) GetHarParams() GetEthCustomCollectionRiskRequestHarParams`

GetHarParams returns the HarParams field if non-nil, zero value otherwise.

### GetHarParamsOk

`func (o *GetPolCustomTokenRiskRequest) GetHarParamsOk() (*GetEthCustomCollectionRiskRequestHarParams, bool)`

GetHarParamsOk returns a tuple with the HarParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarParams

`func (o *GetPolCustomTokenRiskRequest) SetHarParams(v GetEthCustomCollectionRiskRequestHarParams)`

SetHarParams sets HarParams field to given value.

### HasHarParams

`func (o *GetPolCustomTokenRiskRequest) HasHarParams() bool`

HasHarParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


