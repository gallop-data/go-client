# GetPolDefaultTokenRiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The contract address of the token collection. | 
**TokenId** | **[]string** | The id(s) for the token(s). | 
**HoldingPeriod** | **string** | The holding period to evaluate risk for, e.g. &#39;12M&#39; | 
**ReptCurr** | Pointer to **string** | The currency to report results in | [optional] 
**Drawdown** | Pointer to **bool** | If true, report drawdown volatility (based on negative returns only). | [optional] 

## Methods

### NewGetPolDefaultTokenRiskRequest

`func NewGetPolDefaultTokenRiskRequest(collectionAddress string, tokenId []string, holdingPeriod string, ) *GetPolDefaultTokenRiskRequest`

NewGetPolDefaultTokenRiskRequest instantiates a new GetPolDefaultTokenRiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPolDefaultTokenRiskRequestWithDefaults

`func NewGetPolDefaultTokenRiskRequestWithDefaults() *GetPolDefaultTokenRiskRequest`

NewGetPolDefaultTokenRiskRequestWithDefaults instantiates a new GetPolDefaultTokenRiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetPolDefaultTokenRiskRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetPolDefaultTokenRiskRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetPolDefaultTokenRiskRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetTokenId

`func (o *GetPolDefaultTokenRiskRequest) GetTokenId() []string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetPolDefaultTokenRiskRequest) GetTokenIdOk() (*[]string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetPolDefaultTokenRiskRequest) SetTokenId(v []string)`

SetTokenId sets TokenId field to given value.


### GetHoldingPeriod

`func (o *GetPolDefaultTokenRiskRequest) GetHoldingPeriod() string`

GetHoldingPeriod returns the HoldingPeriod field if non-nil, zero value otherwise.

### GetHoldingPeriodOk

`func (o *GetPolDefaultTokenRiskRequest) GetHoldingPeriodOk() (*string, bool)`

GetHoldingPeriodOk returns a tuple with the HoldingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingPeriod

`func (o *GetPolDefaultTokenRiskRequest) SetHoldingPeriod(v string)`

SetHoldingPeriod sets HoldingPeriod field to given value.


### GetReptCurr

`func (o *GetPolDefaultTokenRiskRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetPolDefaultTokenRiskRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetPolDefaultTokenRiskRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetPolDefaultTokenRiskRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetDrawdown

`func (o *GetPolDefaultTokenRiskRequest) GetDrawdown() bool`

GetDrawdown returns the Drawdown field if non-nil, zero value otherwise.

### GetDrawdownOk

`func (o *GetPolDefaultTokenRiskRequest) GetDrawdownOk() (*bool, bool)`

GetDrawdownOk returns a tuple with the Drawdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawdown

`func (o *GetPolDefaultTokenRiskRequest) SetDrawdown(v bool)`

SetDrawdown sets Drawdown field to given value.

### HasDrawdown

`func (o *GetPolDefaultTokenRiskRequest) HasDrawdown() bool`

HasDrawdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


