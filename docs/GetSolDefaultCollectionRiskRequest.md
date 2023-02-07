# GetSolDefaultCollectionRiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionTag** | **string** | The Gallop tag to identify the collection. | 
**HoldingPeriod** | **string** | The holding period to evaluate risk for, e.g. &#39;12M&#39; | 
**Amount** | Pointer to **int32** | The amount of tokens in your portfolio | [optional] 
**ReptCurr** | Pointer to **string** | The currency to report results in | [optional] 
**Drawdown** | Pointer to **bool** | If true, report drawdown volatility (based on negative returns only). | [optional] 

## Methods

### NewGetSolDefaultCollectionRiskRequest

`func NewGetSolDefaultCollectionRiskRequest(collectionTag string, holdingPeriod string, ) *GetSolDefaultCollectionRiskRequest`

NewGetSolDefaultCollectionRiskRequest instantiates a new GetSolDefaultCollectionRiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSolDefaultCollectionRiskRequestWithDefaults

`func NewGetSolDefaultCollectionRiskRequestWithDefaults() *GetSolDefaultCollectionRiskRequest`

NewGetSolDefaultCollectionRiskRequestWithDefaults instantiates a new GetSolDefaultCollectionRiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionTag

`func (o *GetSolDefaultCollectionRiskRequest) GetCollectionTag() string`

GetCollectionTag returns the CollectionTag field if non-nil, zero value otherwise.

### GetCollectionTagOk

`func (o *GetSolDefaultCollectionRiskRequest) GetCollectionTagOk() (*string, bool)`

GetCollectionTagOk returns a tuple with the CollectionTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTag

`func (o *GetSolDefaultCollectionRiskRequest) SetCollectionTag(v string)`

SetCollectionTag sets CollectionTag field to given value.


### GetHoldingPeriod

`func (o *GetSolDefaultCollectionRiskRequest) GetHoldingPeriod() string`

GetHoldingPeriod returns the HoldingPeriod field if non-nil, zero value otherwise.

### GetHoldingPeriodOk

`func (o *GetSolDefaultCollectionRiskRequest) GetHoldingPeriodOk() (*string, bool)`

GetHoldingPeriodOk returns a tuple with the HoldingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingPeriod

`func (o *GetSolDefaultCollectionRiskRequest) SetHoldingPeriod(v string)`

SetHoldingPeriod sets HoldingPeriod field to given value.


### GetAmount

`func (o *GetSolDefaultCollectionRiskRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetSolDefaultCollectionRiskRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetSolDefaultCollectionRiskRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetSolDefaultCollectionRiskRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetReptCurr

`func (o *GetSolDefaultCollectionRiskRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetSolDefaultCollectionRiskRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetSolDefaultCollectionRiskRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetSolDefaultCollectionRiskRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetDrawdown

`func (o *GetSolDefaultCollectionRiskRequest) GetDrawdown() bool`

GetDrawdown returns the Drawdown field if non-nil, zero value otherwise.

### GetDrawdownOk

`func (o *GetSolDefaultCollectionRiskRequest) GetDrawdownOk() (*bool, bool)`

GetDrawdownOk returns a tuple with the Drawdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawdown

`func (o *GetSolDefaultCollectionRiskRequest) SetDrawdown(v bool)`

SetDrawdown sets Drawdown field to given value.

### HasDrawdown

`func (o *GetSolDefaultCollectionRiskRequest) HasDrawdown() bool`

HasDrawdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


