# GetEthTokenSummaryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The contract address of the token collection. | 
**TokenId** | **[]string** | The id for the token. | 
**StartDate** | Pointer to **string** | The start date to pull data for calculations | [optional] 
**EndDate** | Pointer to **string** | The end date to pull data for calculations | [optional] 
**ReptCurr** | Pointer to **string** | The currency to report results in | [optional] 
**ExcludeWash** | Pointer to **bool** | Exclude suspected wash transactions? | [optional] 

## Methods

### NewGetEthTokenSummaryRequest

`func NewGetEthTokenSummaryRequest(collectionAddress string, tokenId []string, ) *GetEthTokenSummaryRequest`

NewGetEthTokenSummaryRequest instantiates a new GetEthTokenSummaryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthTokenSummaryRequestWithDefaults

`func NewGetEthTokenSummaryRequestWithDefaults() *GetEthTokenSummaryRequest`

NewGetEthTokenSummaryRequestWithDefaults instantiates a new GetEthTokenSummaryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetEthTokenSummaryRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetEthTokenSummaryRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetEthTokenSummaryRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetTokenId

`func (o *GetEthTokenSummaryRequest) GetTokenId() []string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetEthTokenSummaryRequest) GetTokenIdOk() (*[]string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetEthTokenSummaryRequest) SetTokenId(v []string)`

SetTokenId sets TokenId field to given value.


### GetStartDate

`func (o *GetEthTokenSummaryRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetEthTokenSummaryRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetEthTokenSummaryRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetEthTokenSummaryRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetEthTokenSummaryRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetEthTokenSummaryRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetEthTokenSummaryRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetEthTokenSummaryRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetReptCurr

`func (o *GetEthTokenSummaryRequest) GetReptCurr() string`

GetReptCurr returns the ReptCurr field if non-nil, zero value otherwise.

### GetReptCurrOk

`func (o *GetEthTokenSummaryRequest) GetReptCurrOk() (*string, bool)`

GetReptCurrOk returns a tuple with the ReptCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptCurr

`func (o *GetEthTokenSummaryRequest) SetReptCurr(v string)`

SetReptCurr sets ReptCurr field to given value.

### HasReptCurr

`func (o *GetEthTokenSummaryRequest) HasReptCurr() bool`

HasReptCurr returns a boolean if a field has been set.

### GetExcludeWash

`func (o *GetEthTokenSummaryRequest) GetExcludeWash() bool`

GetExcludeWash returns the ExcludeWash field if non-nil, zero value otherwise.

### GetExcludeWashOk

`func (o *GetEthTokenSummaryRequest) GetExcludeWashOk() (*bool, bool)`

GetExcludeWashOk returns a tuple with the ExcludeWash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeWash

`func (o *GetEthTokenSummaryRequest) SetExcludeWash(v bool)`

SetExcludeWash sets ExcludeWash field to given value.

### HasExcludeWash

`func (o *GetEthTokenSummaryRequest) HasExcludeWash() bool`

HasExcludeWash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


