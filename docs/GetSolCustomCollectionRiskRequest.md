# GetSolCustomCollectionRiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionTag** | **string** | The Gallop tag to identify the collection. | 
**HoldingPeriod** | **string** | The holding period to evaluate risk for, e.g. &#x60;12M&#x60; | 
**Percentiles** | Pointer to **[]int32** | The collection percentile(s) | [optional] 
**Amount** | Pointer to **int32** | The amount of tokens in your portfolio | [optional] 
**Dist** | Pointer to **string** | The distribution assumed to calculate parametric risk for | [optional] 
**StartDate** | Pointer to **string** | The start date to pull data for calculations | [optional] 
**EndDate** | Pointer to **string** | The end date to pull data for calculations | [optional] 
**WinsOutliers** | Pointer to **bool** | Whether to winsorize time series outliers prior to calculating risk | [optional] 
**RiskFreeRate** | Pointer to **float32** | The rate of return for an asset deemed risk free in the contemplated holding period | [optional] 
**Frequency** | Pointer to **string** | The interval at which to calculate returns to base the forecasts upon, e.g. &#x60;1D&#x60; for daily, &#x60;1M&#x60; for monthly etc. | [optional] 
**ReptCurr** | Pointer to **string** | The currency to report results in | [optional] 
**ExcludeWash** | Pointer to **bool** | Exclude suspected wash transactions? | [optional] 
**Drawdown** | Pointer to **bool** | If true, report drawdown volatility (based on negative returns only). | [optional] 
**ArcParams** | Pointer to [**GetEthCustomCollectionRiskRequestArcParams**](GetEthCustomCollectionRiskRequestArcParams.md) |  | [optional] 
**GarParams** | Pointer to [**GetEthCustomCollectionRiskRequestGarParams**](GetEthCustomCollectionRiskRequestGarParams.md) |  | [optional] 
**HarParams** | Pointer to [**GetEthCustomCollectionRiskRequestHarParams**](GetEthCustomCollectionRiskRequestHarParams.md) |  | [optional] 

## Methods

### NewGetSolCustomCollectionRiskRequest

`func NewGetSolCustomCollectionRiskRequest(collectionTag string, holdingPeriod string, ) *GetSolCustomCollectionRiskRequest`

NewGetSolCustomCollectionRiskRequest instantiates a new GetSolCustomCollectionRiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSolCustomCollectionRiskRequestWithDefaults

`func NewGetSolCustomCollectionRiskRequestWithDefaults() *GetSolCustomCollectionRiskRequest`

NewGetSolCustomCollectionRiskRequestWithDefaults instantiates a new GetSolCustomCollectionRiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionTag

`func (o *GetSolCustomCollectionRiskRequest) GetCollectionTag() string`

GetCollectionTag returns the CollectionTag field if non-nil, zero value otherwise.

### GetCollectionTagOk

`func (o *GetSolCustomCollectionRiskRequest) GetCollectionTagOk() (*string, bool)`

GetCollectionTagOk returns a tuple with the CollectionTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTag

`func (o *GetSolCustomCollectionRiskRequest) SetCollectionTag(v string)`

SetCollectionTag sets CollectionTag field to given value.


### GetHoldingPeriod

`func (o *GetSolCustomCollectionRiskRequest) GetHoldingPeriod() string`

GetHoldingPeriod returns the HoldingPeriod field if non-nil, zero value otherwise.

### GetHoldingPeriodOk

`func (o *GetSolCustomCollectionRiskRequest) GetHoldingPeriodOk() (*string, bool)`

GetHoldingPeriodOk returns a tuple with the HoldingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingPeriod

`func (o *GetSolCustomCollectionRiskRequest) SetHoldingPeriod(v string)`

SetHoldingPeriod sets HoldingPeriod field to given value.


### GetPercentiles

`func (o *GetSolCustomCollectionRiskRequest) GetPercentiles() []int32`

GetPercentiles returns the Percentiles field if non-nil, zero value otherwise.

### GetPercentilesOk

`func (o *GetSolCustomCollectionRiskRequest) GetPercentilesOk() (*[]int32, bool)`

GetPercentilesOk returns a tuple with the Percentiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentiles

`func (o *GetSolCustomCollectionRiskRequest) SetPercentiles(v []int32)`

SetPercentiles sets Percentiles field to given value.

### HasPercentiles

`func (o *GetSolCustomCollectionRiskRequest) HasPercentiles() bool`

HasPercentiles returns a boolean if a field has been set.

### GetAmount

`func (o *GetSolCustomCollectionRiskRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetSolCustomCollectionRiskRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetSolCustomCollectionRiskRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetSolCustomCollectionRiskRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDist

`func (o *GetSolCustomCollectionRiskRequest) GetDist() string`

GetDist returns the Dist field if non-nil, zero value otherwise.

### GetDistOk

`func (o *GetSolCustomCollectionRiskRequest) GetDistOk() (*string, bool)`

GetDistOk returns a tuple with the Dist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDist

`func (o *GetSolCustomCollectionRiskRequest) SetDist(v string)`

SetDist sets Dist field to given value.

### HasDist

`func (o *GetSolCustomCollectionRiskRequest) HasDist() bool`

HasDist returns a boolean if a field has been set.

### GetStartDate

`func (o *GetSolCustomCollectionRiskRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetSolCustomCollectionRiskRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetSolCustomCollectionRiskRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetSolCustomCollectionRiskRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetSolCustomCollectionRiskRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetSolCustomCollectionRiskRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetSolCustomCollectionRiskRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetSolCustomCollectionRiskRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetWinsOutliers

`func (o *GetSolCustomCollectionRiskRequest) GetWinsOutliers() bool`

GetWinsOutliers returns the WinsOutliers field if non-nil, zero value otherwise.

### GetWinsOutliersOk

`func (o *GetSolCustomCollectionRiskRequest) GetWinsOutliersOk() (*bool, bool)`

GetWinsOutliersOk returns a tuple with the WinsOutliers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsOutliers

`func (o *GetSolCustomCollectionRiskRequest) SetWinsOutliers(v bool)`

SetWinsOutliers sets WinsOutliers field to given value.

### HasWinsOutliers

`func (o *GetSolCustomCollectionRiskRequest) HasWinsOutliers() bool`

HasWinsOutliers returns a boolean if a field has been set.

### GetRiskFreeRate

`func (o *GetSolCustomCollectionRiskRequest) GetRiskFreeRate() float32`

GetRiskFreeRate returns the RiskFreeRate field if non-nil, zero value otherwise.

### GetRiskFreeRateOk

`func (o *GetSolCustomCollectionRiskRequest) GetRiskFreeRateOk() (*float32, bool)`

GetRiskFreeRateOk returns a tuple with the RiskFreeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskFreeRate

`func (o *GetSolCustomCollectionRiskRequest) SetRiskFreeRate(v float32)`

SetRiskFreeRate sets RiskFreeRate field to given value.

### HasRiskFreeRate

`func (o *GetSolCustomCollectionRiskRequest) HasRiskFreeRate() bool`

HasRiskFreeRate returns a boolean if a field has been set.

### GetFrequency

`func (o *GetSolCustomCollectionRiskRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetSolCustomCollectionRiskRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetSolCustomCollectionRiskRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetSolCustomCollectionRiskRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetReptCurr

`func (o *GetSolCustomCollectionRiskRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetSolCustomCollectionRiskRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetSolCustomCollectionRiskRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetSolCustomCollectionRiskRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetExcludeWash

`func (o *GetSolCustomCollectionRiskRequest) GetExcludeWash() bool`

GetExcludeWash returns the ExcludeWash field if non-nil, zero value otherwise.

### GetExcludeWashOk

`func (o *GetSolCustomCollectionRiskRequest) GetExcludeWashOk() (*bool, bool)`

GetExcludeWashOk returns a tuple with the ExcludeWash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeWash

`func (o *GetSolCustomCollectionRiskRequest) SetExcludeWash(v bool)`

SetExcludeWash sets ExcludeWash field to given value.

### HasExcludeWash

`func (o *GetSolCustomCollectionRiskRequest) HasExcludeWash() bool`

HasExcludeWash returns a boolean if a field has been set.

### GetDrawdown

`func (o *GetSolCustomCollectionRiskRequest) GetDrawdown() bool`

GetDrawdown returns the Drawdown field if non-nil, zero value otherwise.

### GetDrawdownOk

`func (o *GetSolCustomCollectionRiskRequest) GetDrawdownOk() (*bool, bool)`

GetDrawdownOk returns a tuple with the Drawdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawdown

`func (o *GetSolCustomCollectionRiskRequest) SetDrawdown(v bool)`

SetDrawdown sets Drawdown field to given value.

### HasDrawdown

`func (o *GetSolCustomCollectionRiskRequest) HasDrawdown() bool`

HasDrawdown returns a boolean if a field has been set.

### GetArcParams

`func (o *GetSolCustomCollectionRiskRequest) GetArcParams() GetEthCustomCollectionRiskRequestArcParams`

GetArcParams returns the ArcParams field if non-nil, zero value otherwise.

### GetArcParamsOk

`func (o *GetSolCustomCollectionRiskRequest) GetArcParamsOk() (*GetEthCustomCollectionRiskRequestArcParams, bool)`

GetArcParamsOk returns a tuple with the ArcParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcParams

`func (o *GetSolCustomCollectionRiskRequest) SetArcParams(v GetEthCustomCollectionRiskRequestArcParams)`

SetArcParams sets ArcParams field to given value.

### HasArcParams

`func (o *GetSolCustomCollectionRiskRequest) HasArcParams() bool`

HasArcParams returns a boolean if a field has been set.

### GetGarParams

`func (o *GetSolCustomCollectionRiskRequest) GetGarParams() GetEthCustomCollectionRiskRequestGarParams`

GetGarParams returns the GarParams field if non-nil, zero value otherwise.

### GetGarParamsOk

`func (o *GetSolCustomCollectionRiskRequest) GetGarParamsOk() (*GetEthCustomCollectionRiskRequestGarParams, bool)`

GetGarParamsOk returns a tuple with the GarParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarParams

`func (o *GetSolCustomCollectionRiskRequest) SetGarParams(v GetEthCustomCollectionRiskRequestGarParams)`

SetGarParams sets GarParams field to given value.

### HasGarParams

`func (o *GetSolCustomCollectionRiskRequest) HasGarParams() bool`

HasGarParams returns a boolean if a field has been set.

### GetHarParams

`func (o *GetSolCustomCollectionRiskRequest) GetHarParams() GetEthCustomCollectionRiskRequestHarParams`

GetHarParams returns the HarParams field if non-nil, zero value otherwise.

### GetHarParamsOk

`func (o *GetSolCustomCollectionRiskRequest) GetHarParamsOk() (*GetEthCustomCollectionRiskRequestHarParams, bool)`

GetHarParamsOk returns a tuple with the HarParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarParams

`func (o *GetSolCustomCollectionRiskRequest) SetHarParams(v GetEthCustomCollectionRiskRequestHarParams)`

SetHarParams sets HarParams field to given value.

### HasHarParams

`func (o *GetSolCustomCollectionRiskRequest) HasHarParams() bool`

HasHarParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


