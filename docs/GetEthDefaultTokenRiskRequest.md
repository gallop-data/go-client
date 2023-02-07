# GetEthDefaultTokenRiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The contract address of the token collection. | 
**TokenId** | **[]string** | The id(s) for the token(s). | 
**HoldingPeriod** | **string** | The holding period to evaluate risk for, e.g. &#39;12M&#39; | 
**ReptCurr** | Pointer to **string** | The currency to report results in | [optional] 
**Drawdown** | Pointer to **bool** | If true, report drawdown volatility (based on negative returns only). | [optional] 

## Methods

### NewGetEthDefaultTokenRiskRequest

`func NewGetEthDefaultTokenRiskRequest(collectionAddress string, tokenId []string, holdingPeriod string, ) *GetEthDefaultTokenRiskRequest`

NewGetEthDefaultTokenRiskRequest instantiates a new GetEthDefaultTokenRiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthDefaultTokenRiskRequestWithDefaults

`func NewGetEthDefaultTokenRiskRequestWithDefaults() *GetEthDefaultTokenRiskRequest`

NewGetEthDefaultTokenRiskRequestWithDefaults instantiates a new GetEthDefaultTokenRiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetEthDefaultTokenRiskRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetEthDefaultTokenRiskRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetEthDefaultTokenRiskRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetTokenId

`func (o *GetEthDefaultTokenRiskRequest) GetTokenId() []string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetEthDefaultTokenRiskRequest) GetTokenIdOk() (*[]string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetEthDefaultTokenRiskRequest) SetTokenId(v []string)`

SetTokenId sets TokenId field to given value.


### GetHoldingPeriod

`func (o *GetEthDefaultTokenRiskRequest) GetHoldingPeriod() string`

GetHoldingPeriod returns the HoldingPeriod field if non-nil, zero value otherwise.

### GetHoldingPeriodOk

`func (o *GetEthDefaultTokenRiskRequest) GetHoldingPeriodOk() (*string, bool)`

GetHoldingPeriodOk returns a tuple with the HoldingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingPeriod

`func (o *GetEthDefaultTokenRiskRequest) SetHoldingPeriod(v string)`

SetHoldingPeriod sets HoldingPeriod field to given value.


### GetReptCurr

`func (o *GetEthDefaultTokenRiskRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetEthDefaultTokenRiskRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetEthDefaultTokenRiskRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetEthDefaultTokenRiskRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetDrawdown

`func (o *GetEthDefaultTokenRiskRequest) GetDrawdown() bool`

GetDrawdown returns the Drawdown field if non-nil, zero value otherwise.

### GetDrawdownOk

`func (o *GetEthDefaultTokenRiskRequest) GetDrawdownOk() (*bool, bool)`

GetDrawdownOk returns a tuple with the Drawdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawdown

`func (o *GetEthDefaultTokenRiskRequest) SetDrawdown(v bool)`

SetDrawdown sets Drawdown field to given value.

### HasDrawdown

`func (o *GetEthDefaultTokenRiskRequest) HasDrawdown() bool`

HasDrawdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


