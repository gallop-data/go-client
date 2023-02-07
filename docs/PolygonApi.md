# \PolygonApi

All URIs are relative to *https://api.prod.gallop.run/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPolCollectionForecasts**](PolygonApi.md#GetPolCollectionForecasts) | **Post** /insights/pol/getCollectionForecasts | Price Forecast by Collection
[**GetPolCollectionOwners**](PolygonApi.md#GetPolCollectionOwners) | **Post** /data/pol/getCollectionOwners | Wallet Owners by Collection
[**GetPolCollectionPriceDiff**](PolygonApi.md#GetPolCollectionPriceDiff) | **Post** /analytics/pol/getCollectionPriceDiff | Price Differentiation by Trait
[**GetPolCollectionSalesOHLC**](PolygonApi.md#GetPolCollectionSalesOHLC) | **Post** /analytics/pol/getCollectionSalesOHLC | Collection Sales Price Candlesticks
[**GetPolCollectionSummary**](PolygonApi.md#GetPolCollectionSummary) | **Post** /analytics/pol/getCollectionSummary | Summary Statistics by Collection
[**GetPolCollectionTraits**](PolygonApi.md#GetPolCollectionTraits) | **Post** /data/pol/getCollectionTraits | Traits by Collection
[**GetPolCollectionTransactions**](PolygonApi.md#GetPolCollectionTransactions) | **Post** /data/pol/getCollectionTransactions | Transactions by Collection
[**GetPolCollections**](PolygonApi.md#GetPolCollections) | **Post** /data/pol/getCollections | Aggregated Collections Supported by Gallop
[**GetPolCustomCollectionRisk**](PolygonApi.md#GetPolCustomCollectionRisk) | **Post** /insights/pol/getCustomCollectionRisk | Custom Volatility &amp; Risk Metrics by Collection
[**GetPolCustomTokenRisk**](PolygonApi.md#GetPolCustomTokenRisk) | **Post** /insights/pol/getCustomTokenRisk | Custom Volatility &amp; Risk Metrics by Token
[**GetPolDefaultCollectionRisk**](PolygonApi.md#GetPolDefaultCollectionRisk) | **Post** /insights/pol/getDefaultCollectionRisk | Default Volatility &amp; Risk Metrics by Collection
[**GetPolDefaultTokenRisk**](PolygonApi.md#GetPolDefaultTokenRisk) | **Post** /insights/pol/getDefaultTokenRisk | Default Volatility &amp; Risk Metrics by Token
[**GetPolHistoricalTransactions**](PolygonApi.md#GetPolHistoricalTransactions) | **Post** /data/pol/getHistoricalTransactions | Historical Transactions by Collection
[**GetPolLeaderBoard**](PolygonApi.md#GetPolLeaderBoard) | **Post** /analytics/pol/getLeaderBoard | Polygon Leaderboard by Collection
[**GetPolMarketplaceData**](PolygonApi.md#GetPolMarketplaceData) | **Post** /data/pol/getMarketplaceData | Collection Summary by Marketplace
[**GetPolMarketplaceFloorPrice**](PolygonApi.md#GetPolMarketplaceFloorPrice) | **Post** /data/pol/getMarketplaceFloorPrice | Marketplace Floor Price by Collection
[**GetPolRarity**](PolygonApi.md#GetPolRarity) | **Post** /analytics/pol/getRarity | Token Rarity by Collection
[**GetPolTokenAppraisal**](PolygonApi.md#GetPolTokenAppraisal) | **Post** /insights/pol/getTokenAppraisal | Liquidation &amp; Appraisal Estimate by Token
[**GetPolTokenForecasts**](PolygonApi.md#GetPolTokenForecasts) | **Post** /insights/pol/getTokenForecasts | Price Forecast by Token
[**GetPolTokenSummary**](PolygonApi.md#GetPolTokenSummary) | **Post** /analytics/pol/getTokenSummary | Summary Statistics by Token
[**GetPolTokenTransactions**](PolygonApi.md#GetPolTokenTransactions) | **Post** /data/pol/getTokenTransactions | Transactions by Token
[**GetPolTokens**](PolygonApi.md#GetPolTokens) | **Post** /data/pol/getTokens | Tokens by Collection
[**GetPolWalletLabels**](PolygonApi.md#GetPolWalletLabels) | **Post** /insights/pol/getWalletLabels | Wallet Activity Labels
[**GetPolWalletNFTs**](PolygonApi.md#GetPolWalletNFTs) | **Post** /data/pol/getWalletNFTs | Tokens Owned by Wallet
[**GetPolWalletTransactions**](PolygonApi.md#GetPolWalletTransactions) | **Post** /data/pol/getWalletTransactions | Historical Token Transactions by Wallet
[**GetPolWashTrade**](PolygonApi.md#GetPolWashTrade) | **Post** /analytics/pol/getWashTrade | Wash Trades by Transaction
[**GetPolWashTransactions**](PolygonApi.md#GetPolWashTransactions) | **Post** /analytics/pol/getWashTransactions | Wash Trades by Collection



## GetPolCollectionForecasts

> GetPolCollectionForecasts(ctx).GetPolCollectionForecastsRequest(getPolCollectionForecastsRequest).Execute()

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
    getPolCollectionForecastsRequest := *openapiclient.NewGetPolCollectionForecastsRequest("0xfbe3ab0cbfbd17d06bdd73aa3f55aaf038720f59") // GetPolCollectionForecastsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolCollectionForecasts(context.Background()).GetPolCollectionForecastsRequest(getPolCollectionForecastsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolCollectionForecasts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolCollectionForecastsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolCollectionForecastsRequest** | [**GetPolCollectionForecastsRequest**](GetPolCollectionForecastsRequest.md) |  | 

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


## GetPolCollectionOwners

> GetPolCollectionOwners(ctx).GetPolCollectionOwnersRequest(getPolCollectionOwnersRequest).Execute()

Wallet Owners by Collection



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
    getPolCollectionOwnersRequest := *openapiclient.NewGetPolCollectionOwnersRequest("0x78306a498516d75a76beeacea04fe772d20fe41b") // GetPolCollectionOwnersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolCollectionOwners(context.Background()).GetPolCollectionOwnersRequest(getPolCollectionOwnersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolCollectionOwners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolCollectionOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolCollectionOwnersRequest** | [**GetPolCollectionOwnersRequest**](GetPolCollectionOwnersRequest.md) |  | 

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


## GetPolCollectionPriceDiff

> GetPolCollectionPriceDiff(ctx).GetPolCollectionPriceDiffRequest(getPolCollectionPriceDiffRequest).Execute()

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
    getPolCollectionPriceDiffRequest := *openapiclient.NewGetPolCollectionPriceDiffRequest("0xad59ecb77033391e847cc96829b08beb83151088") // GetPolCollectionPriceDiffRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolCollectionPriceDiff(context.Background()).GetPolCollectionPriceDiffRequest(getPolCollectionPriceDiffRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolCollectionPriceDiff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolCollectionPriceDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolCollectionPriceDiffRequest** | [**GetPolCollectionPriceDiffRequest**](GetPolCollectionPriceDiffRequest.md) |  | 

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


## GetPolCollectionSalesOHLC

> GetPolCollectionSalesOHLC(ctx).GetPolCollectionSalesOHLCRequest(getPolCollectionSalesOHLCRequest).Execute()

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
    getPolCollectionSalesOHLCRequest := *openapiclient.NewGetPolCollectionSalesOHLCRequest("0xfbe3ab0cbfbd17d06bdd73aa3f55aaf038720f59") // GetPolCollectionSalesOHLCRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolCollectionSalesOHLC(context.Background()).GetPolCollectionSalesOHLCRequest(getPolCollectionSalesOHLCRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolCollectionSalesOHLC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolCollectionSalesOHLCRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolCollectionSalesOHLCRequest** | [**GetPolCollectionSalesOHLCRequest**](GetPolCollectionSalesOHLCRequest.md) |  | 

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


## GetPolCollectionSummary

> GetPolCollectionSummary(ctx).GetPolCollectionSummaryRequest(getPolCollectionSummaryRequest).Execute()

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
    getPolCollectionSummaryRequest := *openapiclient.NewGetPolCollectionSummaryRequest("0xfbe3ab0cbfbd17d06bdd73aa3f55aaf038720f59") // GetPolCollectionSummaryRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolCollectionSummary(context.Background()).GetPolCollectionSummaryRequest(getPolCollectionSummaryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolCollectionSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolCollectionSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolCollectionSummaryRequest** | [**GetPolCollectionSummaryRequest**](GetPolCollectionSummaryRequest.md) |  | 

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


## GetPolCollectionTraits

> GetPolCollectionTraits(ctx).GetPolCollectionTraitsRequest(getPolCollectionTraitsRequest).Execute()

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
    getPolCollectionTraitsRequest := *openapiclient.NewGetPolCollectionTraitsRequest() // GetPolCollectionTraitsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolCollectionTraits(context.Background()).GetPolCollectionTraitsRequest(getPolCollectionTraitsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolCollectionTraits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolCollectionTraitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolCollectionTraitsRequest** | [**GetPolCollectionTraitsRequest**](GetPolCollectionTraitsRequest.md) |  | 

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


## GetPolCollectionTransactions

> GetPolCollectionTransactions(ctx).GetPolCollectionTransactionsRequest(getPolCollectionTransactionsRequest).Execute()

Transactions by Collection



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
    getPolCollectionTransactionsRequest := *openapiclient.NewGetPolCollectionTransactionsRequest("0x0833566d5c6926ba123f6587fb16ff4aa8d7680b") // GetPolCollectionTransactionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolCollectionTransactions(context.Background()).GetPolCollectionTransactionsRequest(getPolCollectionTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolCollectionTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolCollectionTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolCollectionTransactionsRequest** | [**GetPolCollectionTransactionsRequest**](GetPolCollectionTransactionsRequest.md) |  | 

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


## GetPolCollections

> GetPolCollections(ctx).GetPolCollectionsRequest(getPolCollectionsRequest).Execute()

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
    getPolCollectionsRequest := *openapiclient.NewGetPolCollectionsRequest() // GetPolCollectionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolCollections(context.Background()).GetPolCollectionsRequest(getPolCollectionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolCollectionsRequest** | [**GetPolCollectionsRequest**](GetPolCollectionsRequest.md) |  | 

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


## GetPolCustomCollectionRisk

> GetPolCustomCollectionRisk(ctx).GetPolCustomCollectionRiskRequest(getPolCustomCollectionRiskRequest).Execute()

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
    getPolCustomCollectionRiskRequest := *openapiclient.NewGetPolCustomCollectionRiskRequest("0xfbe3ab0cbfbd17d06bdd73aa3f55aaf038720f59", "6M") // GetPolCustomCollectionRiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolCustomCollectionRisk(context.Background()).GetPolCustomCollectionRiskRequest(getPolCustomCollectionRiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolCustomCollectionRisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolCustomCollectionRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolCustomCollectionRiskRequest** | [**GetPolCustomCollectionRiskRequest**](GetPolCustomCollectionRiskRequest.md) |  | 

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


## GetPolCustomTokenRisk

> GetPolCustomTokenRisk(ctx).GetPolCustomTokenRiskRequest(getPolCustomTokenRiskRequest).Execute()

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
    getPolCustomTokenRiskRequest := *openapiclient.NewGetPolCustomTokenRiskRequest("0xfbe3ab0cbfbd17d06bdd73aa3f55aaf038720f59", []string{"303"}, "6M") // GetPolCustomTokenRiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolCustomTokenRisk(context.Background()).GetPolCustomTokenRiskRequest(getPolCustomTokenRiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolCustomTokenRisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolCustomTokenRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolCustomTokenRiskRequest** | [**GetPolCustomTokenRiskRequest**](GetPolCustomTokenRiskRequest.md) |  | 

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


## GetPolDefaultCollectionRisk

> GetPolDefaultCollectionRisk(ctx).GetPolDefaultCollectionRiskRequest(getPolDefaultCollectionRiskRequest).Execute()

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
    getPolDefaultCollectionRiskRequest := *openapiclient.NewGetPolDefaultCollectionRiskRequest("0xfbe3ab0cbfbd17d06bdd73aa3f55aaf038720f59", "6M") // GetPolDefaultCollectionRiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolDefaultCollectionRisk(context.Background()).GetPolDefaultCollectionRiskRequest(getPolDefaultCollectionRiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolDefaultCollectionRisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolDefaultCollectionRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolDefaultCollectionRiskRequest** | [**GetPolDefaultCollectionRiskRequest**](GetPolDefaultCollectionRiskRequest.md) |  | 

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


## GetPolDefaultTokenRisk

> GetPolDefaultTokenRisk(ctx).GetPolDefaultTokenRiskRequest(getPolDefaultTokenRiskRequest).Execute()

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
    getPolDefaultTokenRiskRequest := *openapiclient.NewGetPolDefaultTokenRiskRequest("0xfbe3ab0cbfbd17d06bdd73aa3f55aaf038720f59", []string{"6000"}, "6M") // GetPolDefaultTokenRiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolDefaultTokenRisk(context.Background()).GetPolDefaultTokenRiskRequest(getPolDefaultTokenRiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolDefaultTokenRisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolDefaultTokenRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolDefaultTokenRiskRequest** | [**GetPolDefaultTokenRiskRequest**](GetPolDefaultTokenRiskRequest.md) |  | 

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


## GetPolHistoricalTransactions

> GetPolHistoricalTransactions(ctx).GetPolHistoricalTransactionsRequest(getPolHistoricalTransactionsRequest).Execute()

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
    getPolHistoricalTransactionsRequest := *openapiclient.NewGetPolHistoricalTransactionsRequest("0x78306a498516d75a76beeacea04fe772d20fe41b") // GetPolHistoricalTransactionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolHistoricalTransactions(context.Background()).GetPolHistoricalTransactionsRequest(getPolHistoricalTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolHistoricalTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolHistoricalTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolHistoricalTransactionsRequest** | [**GetPolHistoricalTransactionsRequest**](GetPolHistoricalTransactionsRequest.md) |  | 

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


## GetPolLeaderBoard

> GetPolLeaderBoard(ctx).GetEthLeaderBoardRequest(getEthLeaderBoardRequest).Execute()

Polygon Leaderboard by Collection



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
    getEthLeaderBoardRequest := *openapiclient.NewGetEthLeaderBoardRequest("one_day", "eth_volume") // GetEthLeaderBoardRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolLeaderBoard(context.Background()).GetEthLeaderBoardRequest(getEthLeaderBoardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolLeaderBoard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolLeaderBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthLeaderBoardRequest** | [**GetEthLeaderBoardRequest**](GetEthLeaderBoardRequest.md) |  | 

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


## GetPolMarketplaceData

> GetPolMarketplaceData(ctx).GetPolMarketplaceDataRequest(getPolMarketplaceDataRequest).Execute()

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
    getPolMarketplaceDataRequest := *openapiclient.NewGetPolMarketplaceDataRequest([]string{"0x5cc5b05a8a13e3fbdb0bb9fccd98d38e50f90c38"}) // GetPolMarketplaceDataRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolMarketplaceData(context.Background()).GetPolMarketplaceDataRequest(getPolMarketplaceDataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolMarketplaceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolMarketplaceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolMarketplaceDataRequest** | [**GetPolMarketplaceDataRequest**](GetPolMarketplaceDataRequest.md) |  | 

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


## GetPolMarketplaceFloorPrice

> GetPolMarketplaceFloorPrice(ctx).GetPolMarketplaceFloorPriceRequest(getPolMarketplaceFloorPriceRequest).Execute()

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
    getPolMarketplaceFloorPriceRequest := *openapiclient.NewGetPolMarketplaceFloorPriceRequest() // GetPolMarketplaceFloorPriceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolMarketplaceFloorPrice(context.Background()).GetPolMarketplaceFloorPriceRequest(getPolMarketplaceFloorPriceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolMarketplaceFloorPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolMarketplaceFloorPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolMarketplaceFloorPriceRequest** | [**GetPolMarketplaceFloorPriceRequest**](GetPolMarketplaceFloorPriceRequest.md) |  | 

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


## GetPolRarity

> GetPolRarity(ctx).GetPolRarityRequest(getPolRarityRequest).Execute()

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
    getPolRarityRequest := *openapiclient.NewGetPolRarityRequest("0xfbe3ab0cbfbd17d06bdd73aa3f55aaf038720f59") // GetPolRarityRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolRarity(context.Background()).GetPolRarityRequest(getPolRarityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolRarity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolRarityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolRarityRequest** | [**GetPolRarityRequest**](GetPolRarityRequest.md) |  | 

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


## GetPolTokenAppraisal

> GetPolTokenAppraisal(ctx).GetPolTokenAppraisalRequest(getPolTokenAppraisalRequest).Execute()

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
    getPolTokenAppraisalRequest := *openapiclient.NewGetPolTokenAppraisalRequest("0xad59ecb77033391e847cc96829b08beb83151088", []string{"5715"}) // GetPolTokenAppraisalRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolTokenAppraisal(context.Background()).GetPolTokenAppraisalRequest(getPolTokenAppraisalRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolTokenAppraisal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolTokenAppraisalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolTokenAppraisalRequest** | [**GetPolTokenAppraisalRequest**](GetPolTokenAppraisalRequest.md) |  | 

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


## GetPolTokenForecasts

> GetPolTokenForecasts(ctx).GetPolTokenForecastsRequest(getPolTokenForecastsRequest).Execute()

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
    getPolTokenForecastsRequest := *openapiclient.NewGetPolTokenForecastsRequest("0xfbe3ab0cbfbd17d06bdd73aa3f55aaf038720f59", []string{"1722"}) // GetPolTokenForecastsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolTokenForecasts(context.Background()).GetPolTokenForecastsRequest(getPolTokenForecastsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolTokenForecasts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolTokenForecastsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolTokenForecastsRequest** | [**GetPolTokenForecastsRequest**](GetPolTokenForecastsRequest.md) |  | 

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


## GetPolTokenSummary

> GetPolTokenSummary(ctx).GetPolTokenSummaryRequest(getPolTokenSummaryRequest).Execute()

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
    getPolTokenSummaryRequest := *openapiclient.NewGetPolTokenSummaryRequest("0xfbe3ab0cbfbd17d06bdd73aa3f55aaf038720f59", []string{"1608"}) // GetPolTokenSummaryRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolTokenSummary(context.Background()).GetPolTokenSummaryRequest(getPolTokenSummaryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolTokenSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolTokenSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolTokenSummaryRequest** | [**GetPolTokenSummaryRequest**](GetPolTokenSummaryRequest.md) |  | 

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


## GetPolTokenTransactions

> GetPolTokenTransactions(ctx).GetPolTokenTransactionsRequest(getPolTokenTransactionsRequest).Execute()

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
    getPolTokenTransactionsRequest := *openapiclient.NewGetPolTokenTransactionsRequest("0x9d305a42A3975Ee4c1C57555BeD5919889DCE63F", "993") // GetPolTokenTransactionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolTokenTransactions(context.Background()).GetPolTokenTransactionsRequest(getPolTokenTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolTokenTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolTokenTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolTokenTransactionsRequest** | [**GetPolTokenTransactionsRequest**](GetPolTokenTransactionsRequest.md) |  | 

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


## GetPolTokens

> GetPolTokens(ctx).GetPolTokensRequest(getPolTokensRequest).Execute()

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
    getPolTokensRequest := *openapiclient.NewGetPolTokensRequest("0x0004be0357533839ef78038b991fa24ace95544b") // GetPolTokensRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolTokens(context.Background()).GetPolTokensRequest(getPolTokensRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolTokensRequest** | [**GetPolTokensRequest**](GetPolTokensRequest.md) |  | 

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


## GetPolWalletLabels

> GetPolWalletLabels(ctx).GetEthWalletLabelsRequest(getEthWalletLabelsRequest).Execute()

Wallet Activity Labels



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
    getEthWalletLabelsRequest := *openapiclient.NewGetEthWalletLabelsRequest("0xcf561ea02950b819b0999ab3c3b43243d53e9b51") // GetEthWalletLabelsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolWalletLabels(context.Background()).GetEthWalletLabelsRequest(getEthWalletLabelsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolWalletLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolWalletLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthWalletLabelsRequest** | [**GetEthWalletLabelsRequest**](GetEthWalletLabelsRequest.md) |  | 

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


## GetPolWalletNFTs

> GetPolWalletNFTs(ctx).GetPolWalletNFTsRequest(getPolWalletNFTsRequest).Execute()

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
    getPolWalletNFTsRequest := *openapiclient.NewGetPolWalletNFTsRequest("0xab0cda4cc21207fd9433356afe9428b6fac8a8a5") // GetPolWalletNFTsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolWalletNFTs(context.Background()).GetPolWalletNFTsRequest(getPolWalletNFTsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolWalletNFTs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolWalletNFTsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolWalletNFTsRequest** | [**GetPolWalletNFTsRequest**](GetPolWalletNFTsRequest.md) |  | 

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


## GetPolWalletTransactions

> GetPolWalletTransactions(ctx).GetPolWalletTransactionsRequest(getPolWalletTransactionsRequest).Execute()

Historical Token Transactions by Wallet



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
    getPolWalletTransactionsRequest := *openapiclient.NewGetPolWalletTransactionsRequest() // GetPolWalletTransactionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolWalletTransactions(context.Background()).GetPolWalletTransactionsRequest(getPolWalletTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolWalletTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolWalletTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolWalletTransactionsRequest** | [**GetPolWalletTransactionsRequest**](GetPolWalletTransactionsRequest.md) |  | 

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


## GetPolWashTrade

> GetPolWashTrade(ctx).GetPolWashTradeRequest(getPolWashTradeRequest).Execute()

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
    getPolWashTradeRequest := *openapiclient.NewGetPolWashTradeRequest("0x40692e18987ef5237cb0f9f0eaf8898e044a140b94d1cfc2a66dc9eb39033c1a") // GetPolWashTradeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolWashTrade(context.Background()).GetPolWashTradeRequest(getPolWashTradeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolWashTrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolWashTradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolWashTradeRequest** | [**GetPolWashTradeRequest**](GetPolWashTradeRequest.md) |  | 

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


## GetPolWashTransactions

> GetPolWashTransactions(ctx).GetPolWashTransactionsRequest(getPolWashTransactionsRequest).Execute()

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
    getPolWashTransactionsRequest := *openapiclient.NewGetPolWashTransactionsRequest("0xfbe3ab0cbfbd17d06bdd73aa3f55aaf038720f59") // GetPolWashTransactionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolygonApi.GetPolWashTransactions(context.Background()).GetPolWashTransactionsRequest(getPolWashTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolygonApi.GetPolWashTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolWashTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPolWashTransactionsRequest** | [**GetPolWashTransactionsRequest**](GetPolWashTransactionsRequest.md) |  | 

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

