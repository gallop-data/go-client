# GetSolDefaultTokenRiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MintAddress** | **[]string** | A token mint address or list of mint addresses. | 
**HoldingPeriod** | **string** | The holding period to evaluate risk for, e.g. &#39;12M&#39; | 
**ReptCurr** | Pointer to **string** | The currency to report results in | [optional] 
**Drawdown** | Pointer to **bool** | If true, report drawdown volatility (based on negative returns only). | [optional] 

## Methods

### NewGetSolDefaultTokenRiskRequest

`func NewGetSolDefaultTokenRiskRequest(mintAddress []string, holdingPeriod string, ) *GetSolDefaultTokenRiskRequest`

NewGetSolDefaultTokenRiskRequest instantiates a new GetSolDefaultTokenRiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSolDefaultTokenRiskRequestWithDefaults

`func NewGetSolDefaultTokenRiskRequestWithDefaults() *GetSolDefaultTokenRiskRequest`

NewGetSolDefaultTokenRiskRequestWithDefaults instantiates a new GetSolDefaultTokenRiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMintAddress

`func (o *GetSolDefaultTokenRiskRequest) GetMintAddress() []string`

GetMintAddress returns the MintAddress field if non-nil, zero value otherwise.

### GetMintAddressOk

`func (o *GetSolDefaultTokenRiskRequest) GetMintAddressOk() (*[]string, bool)`

GetMintAddressOk returns a tuple with the MintAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintAddress

`func (o *GetSolDefaultTokenRiskRequest) SetMintAddress(v []string)`

SetMintAddress sets MintAddress field to given value.


### GetHoldingPeriod

`func (o *GetSolDefaultTokenRiskRequest) GetHoldingPeriod() string`

GetHoldingPeriod returns the HoldingPeriod field if non-nil, zero value otherwise.

### GetHoldingPeriodOk

`func (o *GetSolDefaultTokenRiskRequest) GetHoldingPeriodOk() (*string, bool)`

GetHoldingPeriodOk returns a tuple with the HoldingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingPeriod

`func (o *GetSolDefaultTokenRiskRequest) SetHoldingPeriod(v string)`

SetHoldingPeriod sets HoldingPeriod field to given value.


### GetReptCurr

`func (o *GetSolDefaultTokenRiskRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetSolDefaultTokenRiskRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetSolDefaultTokenRiskRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetSolDefaultTokenRiskRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetDrawdown

`func (o *GetSolDefaultTokenRiskRequest) GetDrawdown() bool`

GetDrawdown returns the Drawdown field if non-nil, zero value otherwise.

### GetDrawdownOk

`func (o *GetSolDefaultTokenRiskRequest) GetDrawdownOk() (*bool, bool)`

GetDrawdownOk returns a tuple with the Drawdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawdown

`func (o *GetSolDefaultTokenRiskRequest) SetDrawdown(v bool)`

SetDrawdown sets Drawdown field to given value.

### HasDrawdown

`func (o *GetSolDefaultTokenRiskRequest) HasDrawdown() bool`

HasDrawdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


