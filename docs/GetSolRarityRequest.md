# GetSolRarityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionTag** | **string** | The Gallop tag for the Solana collection. Please see sol/getCollections endpoint. | 
**MintAddress** | Pointer to **[]string** | A list of token addresses. | [optional] 
**Weights** | Pointer to **map[string]interface{}** | Dict containing trait keys and weight values. | [optional] 

## Methods

### NewGetSolRarityRequest

`func NewGetSolRarityRequest(collectionTag string, ) *GetSolRarityRequest`

NewGetSolRarityRequest instantiates a new GetSolRarityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSolRarityRequestWithDefaults

`func NewGetSolRarityRequestWithDefaults() *GetSolRarityRequest`

NewGetSolRarityRequestWithDefaults instantiates a new GetSolRarityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionTag

`func (o *GetSolRarityRequest) GetCollectionTag() string`

GetCollectionTag returns the CollectionTag field if non-nil, zero value otherwise.

### GetCollectionTagOk

`func (o *GetSolRarityRequest) GetCollectionTagOk() (*string, bool)`

GetCollectionTagOk returns a tuple with the CollectionTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTag

`func (o *GetSolRarityRequest) SetCollectionTag(v string)`

SetCollectionTag sets CollectionTag field to given value.


### GetMintAddress

`func (o *GetSolRarityRequest) GetMintAddress() []string`

GetMintAddress returns the MintAddress field if non-nil, zero value otherwise.

### GetMintAddressOk

`func (o *GetSolRarityRequest) GetMintAddressOk() (*[]string, bool)`

GetMintAddressOk returns a tuple with the MintAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintAddress

`func (o *GetSolRarityRequest) SetMintAddress(v []string)`

SetMintAddress sets MintAddress field to given value.

### HasMintAddress

`func (o *GetSolRarityRequest) HasMintAddress() bool`

HasMintAddress returns a boolean if a field has been set.

### GetWeights

`func (o *GetSolRarityRequest) GetWeights() map[string]interface{}`

GetWeights returns the Weights field if non-nil, zero value otherwise.

### GetWeightsOk

`func (o *GetSolRarityRequest) GetWeightsOk() (*map[string]interface{}, bool)`

GetWeightsOk returns a tuple with the Weights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeights

`func (o *GetSolRarityRequest) SetWeights(v map[string]interface{})`

SetWeights sets Weights field to given value.

### HasWeights

`func (o *GetSolRarityRequest) HasWeights() bool`

HasWeights returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


