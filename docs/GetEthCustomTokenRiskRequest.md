# GetEthCustomTokenRiskRequest

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

### NewGetEthCustomTokenRiskRequest

`func NewGetEthCustomTokenRiskRequest(collectionAddress string, tokenId []string, holdingPeriod string, ) *GetEthCustomTokenRiskRequest`

NewGetEthCustomTokenRiskRequest instantiates a new GetEthCustomTokenRiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthCustomTokenRiskRequestWithDefaults

`func NewGetEthCustomTokenRiskRequestWithDefaults() *GetEthCustomTokenRiskRequest`

NewGetEthCustomTokenRiskRequestWithDefaults instantiates a new GetEthCustomTokenRiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetEthCustomTokenRiskRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetEthCustomTokenRiskRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetEthCustomTokenRiskRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetTokenId

`func (o *GetEthCustomTokenRiskRequest) GetTokenId() []string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetEthCustomTokenRiskRequest) GetTokenIdOk() (*[]string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetEthCustomTokenRiskRequest) SetTokenId(v []string)`

SetTokenId sets TokenId field to given value.


### GetHoldingPeriod

`func (o *GetEthCustomTokenRiskRequest) GetHoldingPeriod() string`

GetHoldingPeriod returns the HoldingPeriod field if non-nil, zero value otherwise.

### GetHoldingPeriodOk

`func (o *GetEthCustomTokenRiskRequest) GetHoldingPeriodOk() (*string, bool)`

GetHoldingPeriodOk returns a tuple with the HoldingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingPeriod

`func (o *GetEthCustomTokenRiskRequest) SetHoldingPeriod(v string)`

SetHoldingPeriod sets HoldingPeriod field to given value.


### GetDist

`func (o *GetEthCustomTokenRiskRequest) GetDist() string`

GetDist returns the Dist field if non-nil, zero value otherwise.

### GetDistOk

`func (o *GetEthCustomTokenRiskRequest) GetDistOk() (*string, bool)`

GetDistOk returns a tuple with the Dist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDist

`func (o *GetEthCustomTokenRiskRequest) SetDist(v string)`

SetDist sets Dist field to given value.

### HasDist

`func (o *GetEthCustomTokenRiskRequest) HasDist() bool`

HasDist returns a boolean if a field has been set.

### GetStartDate

`func (o *GetEthCustomTokenRiskRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetEthCustomTokenRiskRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetEthCustomTokenRiskRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetEthCustomTokenRiskRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetEthCustomTokenRiskRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetEthCustomTokenRiskRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetEthCustomTokenRiskRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetEthCustomTokenRiskRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetRiskFreeRate

`func (o *GetEthCustomTokenRiskRequest) GetRiskFreeRate() float32`

GetRiskFreeRate returns the RiskFreeRate field if non-nil, zero value otherwise.

### GetRiskFreeRateOk

`func (o *GetEthCustomTokenRiskRequest) GetRiskFreeRateOk() (*float32, bool)`

GetRiskFreeRateOk returns a tuple with the RiskFreeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskFreeRate

`func (o *GetEthCustomTokenRiskRequest) SetRiskFreeRate(v float32)`

SetRiskFreeRate sets RiskFreeRate field to given value.

### HasRiskFreeRate

`func (o *GetEthCustomTokenRiskRequest) HasRiskFreeRate() bool`

HasRiskFreeRate returns a boolean if a field has been set.

### GetWinsOutliers

`func (o *GetEthCustomTokenRiskRequest) GetWinsOutliers() bool`

GetWinsOutliers returns the WinsOutliers field if non-nil, zero value otherwise.

### GetWinsOutliersOk

`func (o *GetEthCustomTokenRiskRequest) GetWinsOutliersOk() (*bool, bool)`

GetWinsOutliersOk returns a tuple with the WinsOutliers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsOutliers

`func (o *GetEthCustomTokenRiskRequest) SetWinsOutliers(v bool)`

SetWinsOutliers sets WinsOutliers field to given value.

### HasWinsOutliers

`func (o *GetEthCustomTokenRiskRequest) HasWinsOutliers() bool`

HasWinsOutliers returns a boolean if a field has been set.

### GetFrequency

`func (o *GetEthCustomTokenRiskRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetEthCustomTokenRiskRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetEthCustomTokenRiskRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetEthCustomTokenRiskRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetReptCurr

`func (o *GetEthCustomTokenRiskRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetEthCustomTokenRiskRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetEthCustomTokenRiskRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetEthCustomTokenRiskRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetExcludeWash

`func (o *GetEthCustomTokenRiskRequest) GetExcludeWash() bool`

GetExcludeWash returns the ExcludeWash field if non-nil, zero value otherwise.

### GetExcludeWashOk

`func (o *GetEthCustomTokenRiskRequest) GetExcludeWashOk() (*bool, bool)`

GetExcludeWashOk returns a tuple with the ExcludeWash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeWash

`func (o *GetEthCustomTokenRiskRequest) SetExcludeWash(v bool)`

SetExcludeWash sets ExcludeWash field to given value.

### HasExcludeWash

`func (o *GetEthCustomTokenRiskRequest) HasExcludeWash() bool`

HasExcludeWash returns a boolean if a field has been set.

### GetDrawdown

`func (o *GetEthCustomTokenRiskRequest) GetDrawdown() bool`

GetDrawdown returns the Drawdown field if non-nil, zero value otherwise.

### GetDrawdownOk

`func (o *GetEthCustomTokenRiskRequest) GetDrawdownOk() (*bool, bool)`

GetDrawdownOk returns a tuple with the Drawdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawdown

`func (o *GetEthCustomTokenRiskRequest) SetDrawdown(v bool)`

SetDrawdown sets Drawdown field to given value.

### HasDrawdown

`func (o *GetEthCustomTokenRiskRequest) HasDrawdown() bool`

HasDrawdown returns a boolean if a field has been set.

### GetArcParams

`func (o *GetEthCustomTokenRiskRequest) GetArcParams() GetEthCustomCollectionRiskRequestArcParams`

GetArcParams returns the ArcParams field if non-nil, zero value otherwise.

### GetArcParamsOk

`func (o *GetEthCustomTokenRiskRequest) GetArcParamsOk() (*GetEthCustomCollectionRiskRequestArcParams, bool)`

GetArcParamsOk returns a tuple with the ArcParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcParams

`func (o *GetEthCustomTokenRiskRequest) SetArcParams(v GetEthCustomCollectionRiskRequestArcParams)`

SetArcParams sets ArcParams field to given value.

### HasArcParams

`func (o *GetEthCustomTokenRiskRequest) HasArcParams() bool`

HasArcParams returns a boolean if a field has been set.

### GetGarParams

`func (o *GetEthCustomTokenRiskRequest) GetGarParams() GetEthCustomCollectionRiskRequestGarParams`

GetGarParams returns the GarParams field if non-nil, zero value otherwise.

### GetGarParamsOk

`func (o *GetEthCustomTokenRiskRequest) GetGarParamsOk() (*GetEthCustomCollectionRiskRequestGarParams, bool)`

GetGarParamsOk returns a tuple with the GarParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarParams

`func (o *GetEthCustomTokenRiskRequest) SetGarParams(v GetEthCustomCollectionRiskRequestGarParams)`

SetGarParams sets GarParams field to given value.

### HasGarParams

`func (o *GetEthCustomTokenRiskRequest) HasGarParams() bool`

HasGarParams returns a boolean if a field has been set.

### GetHarParams

`func (o *GetEthCustomTokenRiskRequest) GetHarParams() GetEthCustomCollectionRiskRequestHarParams`

GetHarParams returns the HarParams field if non-nil, zero value otherwise.

### GetHarParamsOk

`func (o *GetEthCustomTokenRiskRequest) GetHarParamsOk() (*GetEthCustomCollectionRiskRequestHarParams, bool)`

GetHarParamsOk returns a tuple with the HarParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarParams

`func (o *GetEthCustomTokenRiskRequest) SetHarParams(v GetEthCustomCollectionRiskRequestHarParams)`

SetHarParams sets HarParams field to given value.

### HasHarParams

`func (o *GetEthCustomTokenRiskRequest) HasHarParams() bool`

HasHarParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


