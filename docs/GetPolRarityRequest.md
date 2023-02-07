# GetPolRarityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The Polygon contract address to identify the collection. | 
**Weights** | Pointer to **map[string]interface{}** | Dict containing trait keys and weight values. | [optional] 
**TokenId** | Pointer to **[]string** | An array of token ids. | [optional] 

## Methods

### NewGetPolRarityRequest

`func NewGetPolRarityRequest(collectionAddress string, ) *GetPolRarityRequest`

NewGetPolRarityRequest instantiates a new GetPolRarityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPolRarityRequestWithDefaults

`func NewGetPolRarityRequestWithDefaults() *GetPolRarityRequest`

NewGetPolRarityRequestWithDefaults instantiates a new GetPolRarityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetPolRarityRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetPolRarityRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetPolRarityRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetWeights

`func (o *GetPolRarityRequest) GetWeights() map[string]interface{}`

GetWeights returns the Weights field if non-nil, zero value otherwise.

### GetWeightsOk

`func (o *GetPolRarityRequest) GetWeightsOk() (*map[string]interface{}, bool)`

GetWeightsOk returns a tuple with the Weights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeights

`func (o *GetPolRarityRequest) SetWeights(v map[string]interface{})`

SetWeights sets Weights field to given value.

### HasWeights

`func (o *GetPolRarityRequest) HasWeights() bool`

HasWeights returns a boolean if a field has been set.

### GetTokenId

`func (o *GetPolRarityRequest) GetTokenId() []string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetPolRarityRequest) GetTokenIdOk() (*[]string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetPolRarityRequest) SetTokenId(v []string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *GetPolRarityRequest) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


