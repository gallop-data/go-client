# GetEthEnsLookupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | The pagination cursor. | [optional] 
**PageSize** | Pointer to **int32** | The number of records returned per page. | [optional] 
**WalletAddress** | Pointer to **string** | The wallet address to query. | [optional] 
**Name** | Pointer to **string** | The name to query. | [optional] 

## Methods

### NewGetEthEnsLookupRequest

`func NewGetEthEnsLookupRequest() *GetEthEnsLookupRequest`

NewGetEthEnsLookupRequest instantiates a new GetEthEnsLookupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthEnsLookupRequestWithDefaults

`func NewGetEthEnsLookupRequestWithDefaults() *GetEthEnsLookupRequest`

NewGetEthEnsLookupRequestWithDefaults instantiates a new GetEthEnsLookupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetEthEnsLookupRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetEthEnsLookupRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetEthEnsLookupRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetEthEnsLookupRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *GetEthEnsLookupRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetEthEnsLookupRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetEthEnsLookupRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetEthEnsLookupRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetWalletAddress

`func (o *GetEthEnsLookupRequest) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *GetEthEnsLookupRequest) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *GetEthEnsLookupRequest) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.

### HasWalletAddress

`func (o *GetEthEnsLookupRequest) HasWalletAddress() bool`

HasWalletAddress returns a boolean if a field has been set.

### GetName

`func (o *GetEthEnsLookupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetEthEnsLookupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetEthEnsLookupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetEthEnsLookupRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


