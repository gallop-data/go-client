# GetSolTokenAppraisalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MintAddress** | **[]string** | List of mint addresses of tokens to appraise | 
**ReptCurr** | Pointer to **string** | The currency to report results in | [optional] 
**Frequency** | Pointer to **string** | The interval at which to calculate intermediate results and forecasts. | [optional] 
**Horizon** | Pointer to **int32** | The forecast horizon (i.e. the number of periods to forecast out). Defaults to zero which only returns nowcasts. | [optional] 
**Alpha** | Pointer to **float32** | The significance level for the liquidation estimate, e.g. 0.05 for 95% confidence | [optional] 
**ExcludeWash** | Pointer to **bool** | Exclude suspected wash transactions? | [optional] 

## Methods

### NewGetSolTokenAppraisalRequest

`func NewGetSolTokenAppraisalRequest(mintAddress []string, ) *GetSolTokenAppraisalRequest`

NewGetSolTokenAppraisalRequest instantiates a new GetSolTokenAppraisalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSolTokenAppraisalRequestWithDefaults

`func NewGetSolTokenAppraisalRequestWithDefaults() *GetSolTokenAppraisalRequest`

NewGetSolTokenAppraisalRequestWithDefaults instantiates a new GetSolTokenAppraisalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMintAddress

`func (o *GetSolTokenAppraisalRequest) GetMintAddress() []string`

GetMintAddress returns the MintAddress field if non-nil, zero value otherwise.

### GetMintAddressOk

`func (o *GetSolTokenAppraisalRequest) GetMintAddressOk() (*[]string, bool)`

GetMintAddressOk returns a tuple with the MintAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintAddress

`func (o *GetSolTokenAppraisalRequest) SetMintAddress(v []string)`

SetMintAddress sets MintAddress field to given value.


### GetReptCurr

`func (o *GetSolTokenAppraisalRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetSolTokenAppraisalRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetSolTokenAppraisalRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetSolTokenAppraisalRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetFrequency

`func (o *GetSolTokenAppraisalRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetSolTokenAppraisalRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetSolTokenAppraisalRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetSolTokenAppraisalRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetHorizon

`func (o *GetSolTokenAppraisalRequest) GetHorizon() int32`

GetHorizon returns the Horizon field if non-nil, zero value otherwise.

### GetHorizonOk

`func (o *GetSolTokenAppraisalRequest) GetHorizonOk() (*int32, bool)`

GetHorizonOk returns a tuple with the Horizon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizon

`func (o *GetSolTokenAppraisalRequest) SetHorizon(v int32)`

SetHorizon sets Horizon field to given value.

### HasHorizon

`func (o *GetSolTokenAppraisalRequest) HasHorizon() bool`

HasHorizon returns a boolean if a field has been set.

### GetAlpha

`func (o *GetSolTokenAppraisalRequest) GetAlpha() float32`

GetAlpha returns the Alpha field if non-nil, zero value otherwise.

### GetAlphaOk

`func (o *GetSolTokenAppraisalRequest) GetAlphaOk() (*float32, bool)`

GetAlphaOk returns a tuple with the Alpha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpha

`func (o *GetSolTokenAppraisalRequest) SetAlpha(v float32)`

SetAlpha sets Alpha field to given value.

### HasAlpha

`func (o *GetSolTokenAppraisalRequest) HasAlpha() bool`

HasAlpha returns a boolean if a field has been set.

### GetExcludeWash

`func (o *GetSolTokenAppraisalRequest) GetExcludeWash() bool`

GetExcludeWash returns the ExcludeWash field if non-nil, zero value otherwise.

### GetExcludeWashOk

`func (o *GetSolTokenAppraisalRequest) GetExcludeWashOk() (*bool, bool)`

GetExcludeWashOk returns a tuple with the ExcludeWash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeWash

`func (o *GetSolTokenAppraisalRequest) SetExcludeWash(v bool)`

SetExcludeWash sets ExcludeWash field to given value.

### HasExcludeWash

`func (o *GetSolTokenAppraisalRequest) HasExcludeWash() bool`

HasExcludeWash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


