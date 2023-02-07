# GetEthCustomCollectionRiskRequestArcParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mean** | Pointer to **string** | Estimator for the location model of the time series, e.g: &#x60;Zero&#x60;, &#x60;Constant&#x60;, &#x60;ARX&#x60;, ... . | [optional] 
**Lags** | Pointer to **int32** | Number of time time period lags considered. Note that the time period is set by the &#x60;frequency&#x60; parameter, so a value of 2 will assume 2-day lags if &#x60;frequency&#x3D;&#x3D;&#39;1D&#39;&#x60;. | [optional] 
**Vol** | Pointer to **string** | Estimator used for the volatility process of the time series, e.g: &#x60;Constant&#x60;, &#x60;ARCH&#x60;, &#x60;GARCH&#x60;, ...  | [optional] 
**P** | Pointer to **int32** | Order of the symmetric innovation(s). | [optional] 
**Dist** | Pointer to **string** | Return distribution assumed. | [optional] 

## Methods

### NewGetEthCustomCollectionRiskRequestArcParams

`func NewGetEthCustomCollectionRiskRequestArcParams() *GetEthCustomCollectionRiskRequestArcParams`

NewGetEthCustomCollectionRiskRequestArcParams instantiates a new GetEthCustomCollectionRiskRequestArcParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthCustomCollectionRiskRequestArcParamsWithDefaults

`func NewGetEthCustomCollectionRiskRequestArcParamsWithDefaults() *GetEthCustomCollectionRiskRequestArcParams`

NewGetEthCustomCollectionRiskRequestArcParamsWithDefaults instantiates a new GetEthCustomCollectionRiskRequestArcParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMean

`func (o *GetEthCustomCollectionRiskRequestArcParams) GetMean() string`

GetMean returns the Mean field if non-nil, zero value otherwise.

### GetMeanOk

`func (o *GetEthCustomCollectionRiskRequestArcParams) GetMeanOk() (*string, bool)`

GetMeanOk returns a tuple with the Mean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMean

`func (o *GetEthCustomCollectionRiskRequestArcParams) SetMean(v string)`

SetMean sets Mean field to given value.

### HasMean

`func (o *GetEthCustomCollectionRiskRequestArcParams) HasMean() bool`

HasMean returns a boolean if a field has been set.

### GetLags

`func (o *GetEthCustomCollectionRiskRequestArcParams) GetLags() int32`

GetLags returns the Lags field if non-nil, zero value otherwise.

### GetLagsOk

`func (o *GetEthCustomCollectionRiskRequestArcParams) GetLagsOk() (*int32, bool)`

GetLagsOk returns a tuple with the Lags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLags

`func (o *GetEthCustomCollectionRiskRequestArcParams) SetLags(v int32)`

SetLags sets Lags field to given value.

### HasLags

`func (o *GetEthCustomCollectionRiskRequestArcParams) HasLags() bool`

HasLags returns a boolean if a field has been set.

### GetVol

`func (o *GetEthCustomCollectionRiskRequestArcParams) GetVol() string`

GetVol returns the Vol field if non-nil, zero value otherwise.

### GetVolOk

`func (o *GetEthCustomCollectionRiskRequestArcParams) GetVolOk() (*string, bool)`

GetVolOk returns a tuple with the Vol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol

`func (o *GetEthCustomCollectionRiskRequestArcParams) SetVol(v string)`

SetVol sets Vol field to given value.

### HasVol

`func (o *GetEthCustomCollectionRiskRequestArcParams) HasVol() bool`

HasVol returns a boolean if a field has been set.

### GetP

`func (o *GetEthCustomCollectionRiskRequestArcParams) GetP() int32`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *GetEthCustomCollectionRiskRequestArcParams) GetPOk() (*int32, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *GetEthCustomCollectionRiskRequestArcParams) SetP(v int32)`

SetP sets P field to given value.

### HasP

`func (o *GetEthCustomCollectionRiskRequestArcParams) HasP() bool`

HasP returns a boolean if a field has been set.

### GetDist

`func (o *GetEthCustomCollectionRiskRequestArcParams) GetDist() string`

GetDist returns the Dist field if non-nil, zero value otherwise.

### GetDistOk

`func (o *GetEthCustomCollectionRiskRequestArcParams) GetDistOk() (*string, bool)`

GetDistOk returns a tuple with the Dist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDist

`func (o *GetEthCustomCollectionRiskRequestArcParams) SetDist(v string)`

SetDist sets Dist field to given value.

### HasDist

`func (o *GetEthCustomCollectionRiskRequestArcParams) HasDist() bool`

HasDist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


