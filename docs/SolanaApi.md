# \SolanaApi

All URIs are relative to *https://api.prod.gallop.run/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSolAccountNFTs**](SolanaApi.md#GetSolAccountNFTs) | **Post** /data/sol/getAccountNFTs | Tokens Owned by Wallet
[**GetSolCollectionForecasts**](SolanaApi.md#GetSolCollectionForecasts) | **Post** /insights/sol/getCollectionForecasts | Price Forecast by Collection
[**GetSolCollectionPriceDiff**](SolanaApi.md#GetSolCollectionPriceDiff) | **Post** /analytics/sol/getCollectionPriceDiff | Price Differentiation by Trait
[**GetSolCollectionSalesOHLC**](SolanaApi.md#GetSolCollectionSalesOHLC) | **Post** /analytics/sol/getCollectionSalesOHLC | Collection Sales Price Candlesticks
[**GetSolCollectionSummary**](SolanaApi.md#GetSolCollectionSummary) | **Post** /analytics/sol/getCollectionSummary | Summary Statistics by Collection
[**GetSolCollectionTraits**](SolanaApi.md#GetSolCollectionTraits) | **Post** /data/sol/getCollectionTraits | Traits by Collection
[**GetSolCollectionTransactions**](SolanaApi.md#GetSolCollectionTransactions) | **Post** /data/sol/getCollectionTransactions | Transactions by Collections
[**GetSolCollections**](SolanaApi.md#GetSolCollections) | **Post** /data/sol/getCollections | Aggregated Collections Supported by Gallop
[**GetSolCustomCollectionRisk**](SolanaApi.md#GetSolCustomCollectionRisk) | **Post** /insights/sol/getCustomCollectionRisk | Custom Volatility &amp; Risk Metrics by Collection
[**GetSolCustomTokenRisk**](SolanaApi.md#GetSolCustomTokenRisk) | **Post** /insights/sol/getCustomTokenRisk | Custom Volatility &amp; Risk Metrics by Token
[**GetSolDefaultCollectionRisk**](SolanaApi.md#GetSolDefaultCollectionRisk) | **Post** /insights/sol/getDefaultCollectionRisk | Default Volatility &amp; Risk Metrics by Collection
[**GetSolDefaultTokenRisk**](SolanaApi.md#GetSolDefaultTokenRisk) | **Post** /insights/sol/getDefaultTokenRisk | Default Volatility &amp; Risk Metrics by Token
[**GetSolHistoricalTransactions**](SolanaApi.md#GetSolHistoricalTransactions) | **Post** /data/sol/getHistoricalTransactions | Historical Transactions by Collection
[**GetSolMarketplaceData**](SolanaApi.md#GetSolMarketplaceData) | **Post** /data/sol/getMarketplaceData | Collection Summary by Marketplace
[**GetSolMarketplaceFloorPrice**](SolanaApi.md#GetSolMarketplaceFloorPrice) | **Post** /data/sol/getMarketplaceFloorPrice | Marketplace Floor Price by Collection
[**GetSolMarketplaceTraitData**](SolanaApi.md#GetSolMarketplaceTraitData) | **Post** /data/sol/getMarketplaceTraitData | Collection Listings by Trait &amp; Marketplace
[**GetSolNFTAccount**](SolanaApi.md#GetSolNFTAccount) | **Post** /data/sol/getNFTAccount | Wallet Owners by Token
[**GetSolRarity**](SolanaApi.md#GetSolRarity) | **Post** /analytics/sol/getRarity | Token Rarity by Collection
[**GetSolTokenAppraisal**](SolanaApi.md#GetSolTokenAppraisal) | **Post** /insights/sol/getTokenAppraisal | Liquidation &amp; Appraisal Estimate by Token
[**GetSolTokenForecasts**](SolanaApi.md#GetSolTokenForecasts) | **Post** /insights/sol/getTokenForecasts | Price Forecast by Token
[**GetSolTokenSummary**](SolanaApi.md#GetSolTokenSummary) | **Post** /analytics/sol/getTokenSummary | Summary Statistics by Token
[**GetSolTokenTransactions**](SolanaApi.md#GetSolTokenTransactions) | **Post** /data/sol/getTokenTransactions | Transactions by Token
[**GetSolTokens**](SolanaApi.md#GetSolTokens) | **Post** /data/sol/getTokens | Tokens by Collection
[**GetSolWashTrade**](SolanaApi.md#GetSolWashTrade) | **Post** /analytics/sol/getWashTrade | Wash Trades by Transaction
[**GetSolWashTransactions**](SolanaApi.md#GetSolWashTransactions) | **Post** /analytics/sol/getWashTransactions | Wash Trades by Collection



## GetSolAccountNFTs

> GetSolAccountNFTs(ctx).GetSolAccountNFTsRequest(getSolAccountNFTsRequest).Execute()

Tokens Owned by Wallet



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
    getSolAccountNFTsRequest := *openapiclient.NewGetSolAccountNFTsRequest("6CDuoSsn4ZRVvcBMdh9NEKkJebF3AFMdYTnPV4YrfEy6") // GetSolAccountNFTsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolAccountNFTs(context.Background()).GetSolAccountNFTsRequest(getSolAccountNFTsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolAccountNFTs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolAccountNFTsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolAccountNFTsRequest** | [**GetSolAccountNFTsRequest**](GetSolAccountNFTsRequest.md) |  | 

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


## GetSolCollectionForecasts

> GetSolCollectionForecasts(ctx).GetSolCollectionForecastsRequest(getSolCollectionForecastsRequest).Execute()

Price Forecast by Collection



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
    getSolCollectionForecastsRequest := *openapiclient.NewGetSolCollectionForecastsRequest("shadowysupercoder") // GetSolCollectionForecastsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolCollectionForecasts(context.Background()).GetSolCollectionForecastsRequest(getSolCollectionForecastsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolCollectionForecasts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolCollectionForecastsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolCollectionForecastsRequest** | [**GetSolCollectionForecastsRequest**](GetSolCollectionForecastsRequest.md) |  | 

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


## GetSolCollectionPriceDiff

> GetSolCollectionPriceDiff(ctx).GetSolCollectionPriceDiffRequest(getSolCollectionPriceDiffRequest).Execute()

Price Differentiation by Trait



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
    getSolCollectionPriceDiffRequest := *openapiclient.NewGetSolCollectionPriceDiffRequest("shadowysupercoder") // GetSolCollectionPriceDiffRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolCollectionPriceDiff(context.Background()).GetSolCollectionPriceDiffRequest(getSolCollectionPriceDiffRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolCollectionPriceDiff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolCollectionPriceDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolCollectionPriceDiffRequest** | [**GetSolCollectionPriceDiffRequest**](GetSolCollectionPriceDiffRequest.md) |  | 

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


## GetSolCollectionSalesOHLC

> GetSolCollectionSalesOHLC(ctx).GetSolCollectionSalesOHLCRequest(getSolCollectionSalesOHLCRequest).Execute()

Collection Sales Price Candlesticks



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
    getSolCollectionSalesOHLCRequest := *openapiclient.NewGetSolCollectionSalesOHLCRequest("shadowysupercoder") // GetSolCollectionSalesOHLCRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolCollectionSalesOHLC(context.Background()).GetSolCollectionSalesOHLCRequest(getSolCollectionSalesOHLCRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolCollectionSalesOHLC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolCollectionSalesOHLCRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolCollectionSalesOHLCRequest** | [**GetSolCollectionSalesOHLCRequest**](GetSolCollectionSalesOHLCRequest.md) |  | 

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


## GetSolCollectionSummary

> GetSolCollectionSummary(ctx).GetSolCollectionSummaryRequest(getSolCollectionSummaryRequest).Execute()

Summary Statistics by Collection



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
    getSolCollectionSummaryRequest := *openapiclient.NewGetSolCollectionSummaryRequest("shadowysupercoder") // GetSolCollectionSummaryRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolCollectionSummary(context.Background()).GetSolCollectionSummaryRequest(getSolCollectionSummaryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolCollectionSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolCollectionSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolCollectionSummaryRequest** | [**GetSolCollectionSummaryRequest**](GetSolCollectionSummaryRequest.md) |  | 

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


## GetSolCollectionTraits

> GetSolCollectionTraits(ctx).GetSolCollectionTraitsRequest(getSolCollectionTraitsRequest).Execute()

Traits by Collection



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
    getSolCollectionTraitsRequest := *openapiclient.NewGetSolCollectionTraitsRequest() // GetSolCollectionTraitsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolCollectionTraits(context.Background()).GetSolCollectionTraitsRequest(getSolCollectionTraitsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolCollectionTraits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolCollectionTraitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolCollectionTraitsRequest** | [**GetSolCollectionTraitsRequest**](GetSolCollectionTraitsRequest.md) |  | 

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


## GetSolCollectionTransactions

> GetSolCollectionTransactions(ctx).GetSolCollectionTransactionsRequest(getSolCollectionTransactionsRequest).Execute()

Transactions by Collections



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
    getSolCollectionTransactionsRequest := *openapiclient.NewGetSolCollectionTransactionsRequest("degenape") // GetSolCollectionTransactionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolCollectionTransactions(context.Background()).GetSolCollectionTransactionsRequest(getSolCollectionTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolCollectionTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolCollectionTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolCollectionTransactionsRequest** | [**GetSolCollectionTransactionsRequest**](GetSolCollectionTransactionsRequest.md) |  | 

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


## GetSolCollections

> GetSolCollections(ctx).GetSolCollectionsRequest(getSolCollectionsRequest).Execute()

Aggregated Collections Supported by Gallop



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
    getSolCollectionsRequest := *openapiclient.NewGetSolCollectionsRequest() // GetSolCollectionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolCollections(context.Background()).GetSolCollectionsRequest(getSolCollectionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolCollectionsRequest** | [**GetSolCollectionsRequest**](GetSolCollectionsRequest.md) |  | 

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


## GetSolCustomCollectionRisk

> GetSolCustomCollectionRisk(ctx).GetSolCustomCollectionRiskRequest(getSolCustomCollectionRiskRequest).Execute()

Custom Volatility & Risk Metrics by Collection



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
    getSolCustomCollectionRiskRequest := *openapiclient.NewGetSolCustomCollectionRiskRequest("shadowysupercoder", "4M") // GetSolCustomCollectionRiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolCustomCollectionRisk(context.Background()).GetSolCustomCollectionRiskRequest(getSolCustomCollectionRiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolCustomCollectionRisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolCustomCollectionRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolCustomCollectionRiskRequest** | [**GetSolCustomCollectionRiskRequest**](GetSolCustomCollectionRiskRequest.md) |  | 

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


## GetSolCustomTokenRisk

> GetSolCustomTokenRisk(ctx).GetSolCustomTokenRiskRequest(getSolCustomTokenRiskRequest).Execute()

Custom Volatility & Risk Metrics by Token



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
    getSolCustomTokenRiskRequest := *openapiclient.NewGetSolCustomTokenRiskRequest([]string{"DcFTUfcU4WRr4ybkgqpy26Y4g5g4YaCz42dryV6AiHp8"}, "4M") // GetSolCustomTokenRiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolCustomTokenRisk(context.Background()).GetSolCustomTokenRiskRequest(getSolCustomTokenRiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolCustomTokenRisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolCustomTokenRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolCustomTokenRiskRequest** | [**GetSolCustomTokenRiskRequest**](GetSolCustomTokenRiskRequest.md) |  | 

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


## GetSolDefaultCollectionRisk

> GetSolDefaultCollectionRisk(ctx).GetSolDefaultCollectionRiskRequest(getSolDefaultCollectionRiskRequest).Execute()

Default Volatility & Risk Metrics by Collection



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
    getSolDefaultCollectionRiskRequest := *openapiclient.NewGetSolDefaultCollectionRiskRequest("shadowysupercoder", "4M") // GetSolDefaultCollectionRiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolDefaultCollectionRisk(context.Background()).GetSolDefaultCollectionRiskRequest(getSolDefaultCollectionRiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolDefaultCollectionRisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolDefaultCollectionRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolDefaultCollectionRiskRequest** | [**GetSolDefaultCollectionRiskRequest**](GetSolDefaultCollectionRiskRequest.md) |  | 

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


## GetSolDefaultTokenRisk

> GetSolDefaultTokenRisk(ctx).GetSolDefaultTokenRiskRequest(getSolDefaultTokenRiskRequest).Execute()

Default Volatility & Risk Metrics by Token



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
    getSolDefaultTokenRiskRequest := *openapiclient.NewGetSolDefaultTokenRiskRequest([]string{"DcFTUfcU4WRr4ybkgqpy26Y4g5g4YaCz42dryV6AiHp8"}, "4M") // GetSolDefaultTokenRiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolDefaultTokenRisk(context.Background()).GetSolDefaultTokenRiskRequest(getSolDefaultTokenRiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolDefaultTokenRisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolDefaultTokenRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolDefaultTokenRiskRequest** | [**GetSolDefaultTokenRiskRequest**](GetSolDefaultTokenRiskRequest.md) |  | 

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


## GetSolHistoricalTransactions

> GetSolHistoricalTransactions(ctx).GetSolHistoricalTransactionsRequest(getSolHistoricalTransactionsRequest).Execute()

Historical Transactions by Collection



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
    getSolHistoricalTransactionsRequest := *openapiclient.NewGetSolHistoricalTransactionsRequest("degenape") // GetSolHistoricalTransactionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolHistoricalTransactions(context.Background()).GetSolHistoricalTransactionsRequest(getSolHistoricalTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolHistoricalTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolHistoricalTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolHistoricalTransactionsRequest** | [**GetSolHistoricalTransactionsRequest**](GetSolHistoricalTransactionsRequest.md) |  | 

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


## GetSolMarketplaceData

> GetSolMarketplaceData(ctx).GetSolMarketplaceDataRequest(getSolMarketplaceDataRequest).Execute()

Collection Summary by Marketplace



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
    getSolMarketplaceDataRequest := *openapiclient.NewGetSolMarketplaceDataRequest([]string{"downtowngirlz"}) // GetSolMarketplaceDataRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolMarketplaceData(context.Background()).GetSolMarketplaceDataRequest(getSolMarketplaceDataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolMarketplaceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolMarketplaceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolMarketplaceDataRequest** | [**GetSolMarketplaceDataRequest**](GetSolMarketplaceDataRequest.md) |  | 

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


## GetSolMarketplaceFloorPrice

> GetSolMarketplaceFloorPrice(ctx).GetSolMarketplaceFloorPriceRequest(getSolMarketplaceFloorPriceRequest).Execute()

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
    getSolMarketplaceFloorPriceRequest := *openapiclient.NewGetSolMarketplaceFloorPriceRequest() // GetSolMarketplaceFloorPriceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolMarketplaceFloorPrice(context.Background()).GetSolMarketplaceFloorPriceRequest(getSolMarketplaceFloorPriceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolMarketplaceFloorPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolMarketplaceFloorPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolMarketplaceFloorPriceRequest** | [**GetSolMarketplaceFloorPriceRequest**](GetSolMarketplaceFloorPriceRequest.md) |  | 

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


## GetSolMarketplaceTraitData

> GetSolMarketplaceTraitData(ctx).GetSolMarketplaceTraitDataRequest(getSolMarketplaceTraitDataRequest).Execute()

Collection Listings by Trait & Marketplace



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
    getSolMarketplaceTraitDataRequest := *openapiclient.NewGetSolMarketplaceTraitDataRequest("downtowngirlz") // GetSolMarketplaceTraitDataRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolMarketplaceTraitData(context.Background()).GetSolMarketplaceTraitDataRequest(getSolMarketplaceTraitDataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolMarketplaceTraitData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolMarketplaceTraitDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolMarketplaceTraitDataRequest** | [**GetSolMarketplaceTraitDataRequest**](GetSolMarketplaceTraitDataRequest.md) |  | 

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


## GetSolNFTAccount

> GetSolNFTAccount(ctx).GetSolNFTAccountRequest(getSolNFTAccountRequest).Execute()

Wallet Owners by Token



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
    getSolNFTAccountRequest := *openapiclient.NewGetSolNFTAccountRequest("G6oNTZRyrU2tN1YnWREmFaLwEcMA7QDoCscxon4Xrkbp") // GetSolNFTAccountRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolNFTAccount(context.Background()).GetSolNFTAccountRequest(getSolNFTAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolNFTAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolNFTAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolNFTAccountRequest** | [**GetSolNFTAccountRequest**](GetSolNFTAccountRequest.md) |  | 

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


## GetSolRarity

> GetSolRarity(ctx).GetSolRarityRequest(getSolRarityRequest).Execute()

Token Rarity by Collection



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
    getSolRarityRequest := *openapiclient.NewGetSolRarityRequest("portals") // GetSolRarityRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolRarity(context.Background()).GetSolRarityRequest(getSolRarityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolRarity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolRarityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolRarityRequest** | [**GetSolRarityRequest**](GetSolRarityRequest.md) |  | 

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


## GetSolTokenAppraisal

> GetSolTokenAppraisal(ctx).GetSolTokenAppraisalRequest(getSolTokenAppraisalRequest).Execute()

Liquidation & Appraisal Estimate by Token



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
    getSolTokenAppraisalRequest := *openapiclient.NewGetSolTokenAppraisalRequest([]string{"8hBJ4TfzgozfNxqaWLT9hwgMzrbWRKCoeJysjPQnNea5"}) // GetSolTokenAppraisalRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolTokenAppraisal(context.Background()).GetSolTokenAppraisalRequest(getSolTokenAppraisalRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolTokenAppraisal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolTokenAppraisalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolTokenAppraisalRequest** | [**GetSolTokenAppraisalRequest**](GetSolTokenAppraisalRequest.md) |  | 

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


## GetSolTokenForecasts

> GetSolTokenForecasts(ctx).GetSolTokenForecastsRequest(getSolTokenForecastsRequest).Execute()

Price Forecast by Token



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
    getSolTokenForecastsRequest := *openapiclient.NewGetSolTokenForecastsRequest([]string{"DcFTUfcU4WRr4ybkgqpy26Y4g5g4YaCz42dryV6AiHp8"}) // GetSolTokenForecastsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolTokenForecasts(context.Background()).GetSolTokenForecastsRequest(getSolTokenForecastsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolTokenForecasts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolTokenForecastsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolTokenForecastsRequest** | [**GetSolTokenForecastsRequest**](GetSolTokenForecastsRequest.md) |  | 

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


## GetSolTokenSummary

> GetSolTokenSummary(ctx).GetSolTokenSummaryRequest(getSolTokenSummaryRequest).Execute()

Summary Statistics by Token



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
    getSolTokenSummaryRequest := *openapiclient.NewGetSolTokenSummaryRequest([]string{"DcFTUfcU4WRr4ybkgqpy26Y4g5g4YaCz42dryV6AiHp8"}) // GetSolTokenSummaryRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolTokenSummary(context.Background()).GetSolTokenSummaryRequest(getSolTokenSummaryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolTokenSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolTokenSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolTokenSummaryRequest** | [**GetSolTokenSummaryRequest**](GetSolTokenSummaryRequest.md) |  | 

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


## GetSolTokenTransactions

> GetSolTokenTransactions(ctx).GetSolTokenTransactionsRequest(getSolTokenTransactionsRequest).Execute()

Transactions by Token



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
    getSolTokenTransactionsRequest := *openapiclient.NewGetSolTokenTransactionsRequest("5uarH96GLf4APmipQRrsD3SsW4K77zRPtWGwETcyz5GE") // GetSolTokenTransactionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolTokenTransactions(context.Background()).GetSolTokenTransactionsRequest(getSolTokenTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolTokenTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolTokenTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolTokenTransactionsRequest** | [**GetSolTokenTransactionsRequest**](GetSolTokenTransactionsRequest.md) |  | 

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


## GetSolTokens

> GetSolTokens(ctx).GetSolTokensRequest(getSolTokensRequest).Execute()

Tokens by Collection



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
    getSolTokensRequest := *openapiclient.NewGetSolTokensRequest("degenape") // GetSolTokensRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolTokens(context.Background()).GetSolTokensRequest(getSolTokensRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolTokensRequest** | [**GetSolTokensRequest**](GetSolTokensRequest.md) |  | 

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


## GetSolWashTrade

> GetSolWashTrade(ctx).GetSolWashTradeRequest(getSolWashTradeRequest).Execute()

Wash Trades by Transaction



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
    getSolWashTradeRequest := *openapiclient.NewGetSolWashTradeRequest("jAV6y8eKCACX9bwNaiVgh6mqoxiCHrovWSYVXWmZfj5EhHDmXHeu23DwgnvnPQHjudKj9DK5ed73zq2yAkaJtqg") // GetSolWashTradeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolWashTrade(context.Background()).GetSolWashTradeRequest(getSolWashTradeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolWashTrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolWashTradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolWashTradeRequest** | [**GetSolWashTradeRequest**](GetSolWashTradeRequest.md) |  | 

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


## GetSolWashTransactions

> GetSolWashTransactions(ctx).GetSolWashTransactionsRequest(getSolWashTransactionsRequest).Execute()

Wash Trades by Collection



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
    getSolWashTransactionsRequest := *openapiclient.NewGetSolWashTransactionsRequest("degods") // GetSolWashTransactionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SolanaApi.GetSolWashTransactions(context.Background()).GetSolWashTransactionsRequest(getSolWashTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaApi.GetSolWashTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolWashTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSolWashTransactionsRequest** | [**GetSolWashTransactionsRequest**](GetSolWashTransactionsRequest.md) |  | 

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

