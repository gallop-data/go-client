# GetEthWalletNFTsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletAddress** | **string** | The wallet address to search. | 
**Page** | Pointer to **int32** | The pagination cursor. | [optional] 
**PageSize** | Pointer to **int32** | The number of records returned per page. | [optional] 

## Methods

### NewGetEthWalletNFTsRequest

`func NewGetEthWalletNFTsRequest(walletAddress string, ) *GetEthWalletNFTsRequest`

NewGetEthWalletNFTsRequest instantiates a new GetEthWalletNFTsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthWalletNFTsRequestWithDefaults

`func NewGetEthWalletNFTsRequestWithDefaults() *GetEthWalletNFTsRequest`

NewGetEthWalletNFTsRequestWithDefaults instantiates a new GetEthWalletNFTsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletAddress

`func (o *GetEthWalletNFTsRequest) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *GetEthWalletNFTsRequest) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *GetEthWalletNFTsRequest) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.


### GetPage

`func (o *GetEthWalletNFTsRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetEthWalletNFTsRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetEthWalletNFTsRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetEthWalletNFTsRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *GetEthWalletNFTsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetEthWalletNFTsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetEthWalletNFTsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetEthWalletNFTsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


