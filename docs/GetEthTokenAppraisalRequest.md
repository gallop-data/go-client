# GetEthTokenAppraisalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The contract address of the token collection. | 
**TokenId** | **[]string** | The id(s) for the token(s). | 
**ReptCurr** | Pointer to **string** | The currency to report results in | [optional] 
**Frequency** | Pointer to **string** | The interval at which to calculate intermediate results and forecasts. | [optional] 
**Horizon** | Pointer to **int32** | The forecast horizon (i.e. the number of periods to forecast out). Defaults to zero which only returns nowcasts. | [optional] 
**Alpha** | Pointer to **float32** | The significance level for the liquidation estimate, e.g. 0.05 for 95% confidence | [optional] 
**ExcludeWash** | Pointer to **bool** | Exclude suspected wash transactions? | [optional] 

## Methods

### NewGetEthTokenAppraisalRequest

`func NewGetEthTokenAppraisalRequest(collectionAddress string, tokenId []string, ) *GetEthTokenAppraisalRequest`

NewGetEthTokenAppraisalRequest instantiates a new GetEthTokenAppraisalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthTokenAppraisalRequestWithDefaults

`func NewGetEthTokenAppraisalRequestWithDefaults() *GetEthTokenAppraisalRequest`

NewGetEthTokenAppraisalRequestWithDefaults instantiates a new GetEthTokenAppraisalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetEthTokenAppraisalRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetEthTokenAppraisalRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetEthTokenAppraisalRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetTokenId

`func (o *GetEthTokenAppraisalRequest) GetTokenId() []string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetEthTokenAppraisalRequest) GetTokenIdOk() (*[]string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetEthTokenAppraisalRequest) SetTokenId(v []string)`

SetTokenId sets TokenId field to given value.


### GetReptCurr

`func (o *GetEthTokenAppraisalRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetEthTokenAppraisalRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetEthTokenAppraisalRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetEthTokenAppraisalRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetFrequency

`func (o *GetEthTokenAppraisalRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetEthTokenAppraisalRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetEthTokenAppraisalRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetEthTokenAppraisalRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetHorizon

`func (o *GetEthTokenAppraisalRequest) GetHorizon() int32`

GetHorizon returns the Horizon field if non-nil, zero value otherwise.

### GetHorizonOk

`func (o *GetEthTokenAppraisalRequest) GetHorizonOk() (*int32, bool)`

GetHorizonOk returns a tuple with the Horizon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizon

`func (o *GetEthTokenAppraisalRequest) SetHorizon(v int32)`

SetHorizon sets Horizon field to given value.

### HasHorizon

`func (o *GetEthTokenAppraisalRequest) HasHorizon() bool`

HasHorizon returns a boolean if a field has been set.

### GetAlpha

`func (o *GetEthTokenAppraisalRequest) GetAlpha() float32`

GetAlpha returns the Alpha field if non-nil, zero value otherwise.

### GetAlphaOk

`func (o *GetEthTokenAppraisalRequest) GetAlphaOk() (*float32, bool)`

GetAlphaOk returns a tuple with the Alpha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpha

`func (o *GetEthTokenAppraisalRequest) SetAlpha(v float32)`

SetAlpha sets Alpha field to given value.

### HasAlpha

`func (o *GetEthTokenAppraisalRequest) HasAlpha() bool`

HasAlpha returns a boolean if a field has been set.

### GetExcludeWash

`func (o *GetEthTokenAppraisalRequest) GetExcludeWash() bool`

GetExcludeWash returns the ExcludeWash field if non-nil, zero value otherwise.

### GetExcludeWashOk

`func (o *GetEthTokenAppraisalRequest) GetExcludeWashOk() (*bool, bool)`

GetExcludeWashOk returns a tuple with the ExcludeWash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeWash

`func (o *GetEthTokenAppraisalRequest) SetExcludeWash(v bool)`

SetExcludeWash sets ExcludeWash field to given value.

### HasExcludeWash

`func (o *GetEthTokenAppraisalRequest) HasExcludeWash() bool`

HasExcludeWash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


