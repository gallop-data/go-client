# GetEthRarityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The Ethereum contract address to identify the collection. | 
**Weights** | Pointer to **map[string]interface{}** | Dict containing trait keys and weight values. | [optional] 
**TokenId** | Pointer to **[]string** | An array of token ids. | [optional] 

## Methods

### NewGetEthRarityRequest

`func NewGetEthRarityRequest(collectionAddress string, ) *GetEthRarityRequest`

NewGetEthRarityRequest instantiates a new GetEthRarityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthRarityRequestWithDefaults

`func NewGetEthRarityRequestWithDefaults() *GetEthRarityRequest`

NewGetEthRarityRequestWithDefaults instantiates a new GetEthRarityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetEthRarityRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetEthRarityRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetEthRarityRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetWeights

`func (o *GetEthRarityRequest) GetWeights() map[string]interface{}`

GetWeights returns the Weights field if non-nil, zero value otherwise.

### GetWeightsOk

`func (o *GetEthRarityRequest) GetWeightsOk() (*map[string]interface{}, bool)`

GetWeightsOk returns a tuple with the Weights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeights

`func (o *GetEthRarityRequest) SetWeights(v map[string]interface{})`

SetWeights sets Weights field to given value.

### HasWeights

`func (o *GetEthRarityRequest) HasWeights() bool`

HasWeights returns a boolean if a field has been set.

### GetTokenId

`func (o *GetEthRarityRequest) GetTokenId() []string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetEthRarityRequest) GetTokenIdOk() (*[]string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetEthRarityRequest) SetTokenId(v []string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *GetEthRarityRequest) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


