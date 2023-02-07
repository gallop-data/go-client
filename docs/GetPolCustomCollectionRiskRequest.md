# GetPolCustomCollectionRiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The contract address of the token collection. | 
**HoldingPeriod** | **string** | The holding period to evaluate risk for, e.g. &#x60;12M&#x60; | 
**Percentiles** | Pointer to **[]int32** | The collection percentile(s) | [optional] 
**Amount** | Pointer to **int32** | The amount of tokens in your portfolio | [optional] 
**Dist** | Pointer to **string** | The distribution assumed to calculate parametric risk for | [optional] 
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

### NewGetPolCustomCollectionRiskRequest

`func NewGetPolCustomCollectionRiskRequest(collectionAddress string, holdingPeriod string, ) *GetPolCustomCollectionRiskRequest`

NewGetPolCustomCollectionRiskRequest instantiates a new GetPolCustomCollectionRiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPolCustomCollectionRiskRequestWithDefaults

`func NewGetPolCustomCollectionRiskRequestWithDefaults() *GetPolCustomCollectionRiskRequest`

NewGetPolCustomCollectionRiskRequestWithDefaults instantiates a new GetPolCustomCollectionRiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetPolCustomCollectionRiskRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetPolCustomCollectionRiskRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetPolCustomCollectionRiskRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetHoldingPeriod

`func (o *GetPolCustomCollectionRiskRequest) GetHoldingPeriod() string`

GetHoldingPeriod returns the HoldingPeriod field if non-nil, zero value otherwise.

### GetHoldingPeriodOk

`func (o *GetPolCustomCollectionRiskRequest) GetHoldingPeriodOk() (*string, bool)`

GetHoldingPeriodOk returns a tuple with the HoldingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingPeriod

`func (o *GetPolCustomCollectionRiskRequest) SetHoldingPeriod(v string)`

SetHoldingPeriod sets HoldingPeriod field to given value.


### GetPercentiles

`func (o *GetPolCustomCollectionRiskRequest) GetPercentiles() []int32`

GetPercentiles returns the Percentiles field if non-nil, zero value otherwise.

### GetPercentilesOk

`func (o *GetPolCustomCollectionRiskRequest) GetPercentilesOk() (*[]int32, bool)`

GetPercentilesOk returns a tuple with the Percentiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentiles

`func (o *GetPolCustomCollectionRiskRequest) SetPercentiles(v []int32)`

SetPercentiles sets Percentiles field to given value.

### HasPercentiles

`func (o *GetPolCustomCollectionRiskRequest) HasPercentiles() bool`

HasPercentiles returns a boolean if a field has been set.

### GetAmount

`func (o *GetPolCustomCollectionRiskRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetPolCustomCollectionRiskRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetPolCustomCollectionRiskRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetPolCustomCollectionRiskRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDist

`func (o *GetPolCustomCollectionRiskRequest) GetDist() string`

GetDist returns the Dist field if non-nil, zero value otherwise.

### GetDistOk

`func (o *GetPolCustomCollectionRiskRequest) GetDistOk() (*string, bool)`

GetDistOk returns a tuple with the Dist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDist

`func (o *GetPolCustomCollectionRiskRequest) SetDist(v string)`

SetDist sets Dist field to given value.

### HasDist

`func (o *GetPolCustomCollectionRiskRequest) HasDist() bool`

HasDist returns a boolean if a field has been set.

### GetStartDate

`func (o *GetPolCustomCollectionRiskRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetPolCustomCollectionRiskRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetPolCustomCollectionRiskRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetPolCustomCollectionRiskRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetPolCustomCollectionRiskRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetPolCustomCollectionRiskRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetPolCustomCollectionRiskRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetPolCustomCollectionRiskRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetRiskFreeRate

`func (o *GetPolCustomCollectionRiskRequest) GetRiskFreeRate() float32`

GetRiskFreeRate returns the RiskFreeRate field if non-nil, zero value otherwise.

### GetRiskFreeRateOk

`func (o *GetPolCustomCollectionRiskRequest) GetRiskFreeRateOk() (*float32, bool)`

GetRiskFreeRateOk returns a tuple with the RiskFreeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskFreeRate

`func (o *GetPolCustomCollectionRiskRequest) SetRiskFreeRate(v float32)`

SetRiskFreeRate sets RiskFreeRate field to given value.

### HasRiskFreeRate

`func (o *GetPolCustomCollectionRiskRequest) HasRiskFreeRate() bool`

HasRiskFreeRate returns a boolean if a field has been set.

### GetWinsOutliers

`func (o *GetPolCustomCollectionRiskRequest) GetWinsOutliers() bool`

GetWinsOutliers returns the WinsOutliers field if non-nil, zero value otherwise.

### GetWinsOutliersOk

`func (o *GetPolCustomCollectionRiskRequest) GetWinsOutliersOk() (*bool, bool)`

GetWinsOutliersOk returns a tuple with the WinsOutliers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsOutliers

`func (o *GetPolCustomCollectionRiskRequest) SetWinsOutliers(v bool)`

SetWinsOutliers sets WinsOutliers field to given value.

### HasWinsOutliers

`func (o *GetPolCustomCollectionRiskRequest) HasWinsOutliers() bool`

HasWinsOutliers returns a boolean if a field has been set.

### GetFrequency

`func (o *GetPolCustomCollectionRiskRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetPolCustomCollectionRiskRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetPolCustomCollectionRiskRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetPolCustomCollectionRiskRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetReptCurr

`func (o *GetPolCustomCollectionRiskRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetPolCustomCollectionRiskRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetPolCustomCollectionRiskRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetPolCustomCollectionRiskRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetExcludeWash

`func (o *GetPolCustomCollectionRiskRequest) GetExcludeWash() bool`

GetExcludeWash returns the ExcludeWash field if non-nil, zero value otherwise.

### GetExcludeWashOk

`func (o *GetPolCustomCollectionRiskRequest) GetExcludeWashOk() (*bool, bool)`

GetExcludeWashOk returns a tuple with the ExcludeWash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeWash

`func (o *GetPolCustomCollectionRiskRequest) SetExcludeWash(v bool)`

SetExcludeWash sets ExcludeWash field to given value.

### HasExcludeWash

`func (o *GetPolCustomCollectionRiskRequest) HasExcludeWash() bool`

HasExcludeWash returns a boolean if a field has been set.

### GetDrawdown

`func (o *GetPolCustomCollectionRiskRequest) GetDrawdown() bool`

GetDrawdown returns the Drawdown field if non-nil, zero value otherwise.

### GetDrawdownOk

`func (o *GetPolCustomCollectionRiskRequest) GetDrawdownOk() (*bool, bool)`

GetDrawdownOk returns a tuple with the Drawdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawdown

`func (o *GetPolCustomCollectionRiskRequest) SetDrawdown(v bool)`

SetDrawdown sets Drawdown field to given value.

### HasDrawdown

`func (o *GetPolCustomCollectionRiskRequest) HasDrawdown() bool`

HasDrawdown returns a boolean if a field has been set.

### GetArcParams

`func (o *GetPolCustomCollectionRiskRequest) GetArcParams() GetEthCustomCollectionRiskRequestArcParams`

GetArcParams returns the ArcParams field if non-nil, zero value otherwise.

### GetArcParamsOk

`func (o *GetPolCustomCollectionRiskRequest) GetArcParamsOk() (*GetEthCustomCollectionRiskRequestArcParams, bool)`

GetArcParamsOk returns a tuple with the ArcParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcParams

`func (o *GetPolCustomCollectionRiskRequest) SetArcParams(v GetEthCustomCollectionRiskRequestArcParams)`

SetArcParams sets ArcParams field to given value.

### HasArcParams

`func (o *GetPolCustomCollectionRiskRequest) HasArcParams() bool`

HasArcParams returns a boolean if a field has been set.

### GetGarParams

`func (o *GetPolCustomCollectionRiskRequest) GetGarParams() GetEthCustomCollectionRiskRequestGarParams`

GetGarParams returns the GarParams field if non-nil, zero value otherwise.

### GetGarParamsOk

`func (o *GetPolCustomCollectionRiskRequest) GetGarParamsOk() (*GetEthCustomCollectionRiskRequestGarParams, bool)`

GetGarParamsOk returns a tuple with the GarParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarParams

`func (o *GetPolCustomCollectionRiskRequest) SetGarParams(v GetEthCustomCollectionRiskRequestGarParams)`

SetGarParams sets GarParams field to given value.

### HasGarParams

`func (o *GetPolCustomCollectionRiskRequest) HasGarParams() bool`

HasGarParams returns a boolean if a field has been set.

### GetHarParams

`func (o *GetPolCustomCollectionRiskRequest) GetHarParams() GetEthCustomCollectionRiskRequestHarParams`

GetHarParams returns the HarParams field if non-nil, zero value otherwise.

### GetHarParamsOk

`func (o *GetPolCustomCollectionRiskRequest) GetHarParamsOk() (*GetEthCustomCollectionRiskRequestHarParams, bool)`

GetHarParamsOk returns a tuple with the HarParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarParams

`func (o *GetPolCustomCollectionRiskRequest) SetHarParams(v GetEthCustomCollectionRiskRequestHarParams)`

SetHarParams sets HarParams field to given value.

### HasHarParams

`func (o *GetPolCustomCollectionRiskRequest) HasHarParams() bool`

HasHarParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


