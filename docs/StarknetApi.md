# \StarknetApi

All URIs are relative to *https://api.prod.gallop.run/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSknMarketplaceData**](StarknetApi.md#GetSknMarketplaceData) | **Post** /data/skn/getMarketplaceData | Gallop Marketplace Data
[**GetSknMarketplaceFloorPrice**](StarknetApi.md#GetSknMarketplaceFloorPrice) | **Post** /data/skn/getMarketplaceFloorPrice | Marketplace Floor Price by Collection



## GetSknMarketplaceData

> GetSknMarketplaceData(ctx).GetSknMarketplaceDataRequest(getSknMarketplaceDataRequest).Execute()

Gallop Marketplace Data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    getSknMarketplaceDataRequest := *openapiclient.NewGetSknMarketplaceDataRequest([]string{"0x04acd4b2a59eae7196f6a5c26ead8cb5f9d7ad3d911096418a23357bb2eac075"}) // GetSknMarketplaceDataRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StarknetApi.GetSknMarketplaceData(context.Background()).GetSknMarketplaceDataRequest(getSknMarketplaceDataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StarknetApi.GetSknMarketplaceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSknMarketplaceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSknMarketplaceDataRequest** | [**GetSknMarketplaceDataRequest**](GetSknMarketplaceDataRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSknMarketplaceFloorPrice

> GetSknMarketplaceFloorPrice(ctx).GetSknMarketplaceFloorPriceRequest(getSknMarketplaceFloorPriceRequest).Execute()

Marketplace Floor Price by Collection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    getSknMarketplaceFloorPriceRequest := *openapiclient.NewGetSknMarketplaceFloorPriceRequest() // GetSknMarketplaceFloorPriceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StarknetApi.GetSknMarketplaceFloorPrice(context.Background()).GetSknMarketplaceFloorPriceRequest(getSknMarketplaceFloorPriceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StarknetApi.GetSknMarketplaceFloorPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSknMarketplaceFloorPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSknMarketplaceFloorPriceRequest** | [**GetSknMarketplaceFloorPriceRequest**](GetSknMarketplaceFloorPriceRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

