# GetEthDefaultCollectionRiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The contract address of the token collection. | 
**HoldingPeriod** | **string** | The holding period to evaluate risk for, e.g. &#39;12M&#39; | 
**ReptCurr** | Pointer to **string** | The currency to report results in | [optional] 
**Amount** | Pointer to **int32** | The amount of tokens in your portfolio | [optional] 
**Drawdown** | Pointer to **bool** | If true, report drawdown volatility (based on negative returns only). | [optional] 

## Methods

### NewGetEthDefaultCollectionRiskRequest

`func NewGetEthDefaultCollectionRiskRequest(collectionAddress string, holdingPeriod string, ) *GetEthDefaultCollectionRiskRequest`

NewGetEthDefaultCollectionRiskRequest instantiates a new GetEthDefaultCollectionRiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthDefaultCollectionRiskRequestWithDefaults

`func NewGetEthDefaultCollectionRiskRequestWithDefaults() *GetEthDefaultCollectionRiskRequest`

NewGetEthDefaultCollectionRiskRequestWithDefaults instantiates a new GetEthDefaultCollectionRiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetEthDefaultCollectionRiskRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetEthDefaultCollectionRiskRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetEthDefaultCollectionRiskRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetHoldingPeriod

`func (o *GetEthDefaultCollectionRiskRequest) GetHoldingPeriod() string`

GetHoldingPeriod returns the HoldingPeriod field if non-nil, zero value otherwise.

### GetHoldingPeriodOk

`func (o *GetEthDefaultCollectionRiskRequest) GetHoldingPeriodOk() (*string, bool)`

GetHoldingPeriodOk returns a tuple with the HoldingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingPeriod

`func (o *GetEthDefaultCollectionRiskRequest) SetHoldingPeriod(v string)`

SetHoldingPeriod sets HoldingPeriod field to given value.


### GetReptCurr

`func (o *GetEthDefaultCollectionRiskRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetEthDefaultCollectionRiskRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetEthDefaultCollectionRiskRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetEthDefaultCollectionRiskRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetAmount

`func (o *GetEthDefaultCollectionRiskRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetEthDefaultCollectionRiskRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetEthDefaultCollectionRiskRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetEthDefaultCollectionRiskRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDrawdown

`func (o *GetEthDefaultCollectionRiskRequest) GetDrawdown() bool`

GetDrawdown returns the Drawdown field if non-nil, zero value otherwise.

### GetDrawdownOk

`func (o *GetEthDefaultCollectionRiskRequest) GetDrawdownOk() (*bool, bool)`

GetDrawdownOk returns a tuple with the Drawdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawdown

`func (o *GetEthDefaultCollectionRiskRequest) SetDrawdown(v bool)`

SetDrawdown sets Drawdown field to given value.

### HasDrawdown

`func (o *GetEthDefaultCollectionRiskRequest) HasDrawdown() bool`

HasDrawdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


