# \EthereumApi

All URIs are relative to *https://api.prod.gallop.run/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEthCollectionFloorPriceOHLC**](EthereumApi.md#GetEthCollectionFloorPriceOHLC) | **Post** /analytics/eth/getCollectionFloorPriceOHLC | Intraday Marketplace Floor Price by Collection
[**GetEthCollectionForecasts**](EthereumApi.md#GetEthCollectionForecasts) | **Post** /insights/eth/getCollectionForecasts | Price Forecast by Collection
[**GetEthCollectionListingsOHLC**](EthereumApi.md#GetEthCollectionListingsOHLC) | **Post** /analytics/eth/getCollectionListingsOHLC | Collection Floor Price and Listings Candlesticks
[**GetEthCollectionOwners**](EthereumApi.md#GetEthCollectionOwners) | **Post** /data/eth/getCollectionOwners | Wallet Owners by Collection
[**GetEthCollectionPriceDiff**](EthereumApi.md#GetEthCollectionPriceDiff) | **Post** /analytics/eth/getCollectionPriceDiff | Price Differentiation by Trait
[**GetEthCollectionSalesOHLC**](EthereumApi.md#GetEthCollectionSalesOHLC) | **Post** /analytics/eth/getCollectionSalesOHLC | Collection Sales Price Candlesticks
[**GetEthCollectionSummary**](EthereumApi.md#GetEthCollectionSummary) | **Post** /analytics/eth/getCollectionSummary | Summary Statistics by Collection
[**GetEthCollectionTransactions**](EthereumApi.md#GetEthCollectionTransactions) | **Post** /data/eth/getCollectionTransactions | Transactions by Collection
[**GetEthCollections**](EthereumApi.md#GetEthCollections) | **Post** /data/eth/getCollections | Aggregated Collections Supported by Gallop
[**GetEthCustomCollectionRisk**](EthereumApi.md#GetEthCustomCollectionRisk) | **Post** /insights/eth/getCustomCollectionRisk | Custom Volatility &amp; Risk Metrics by Collection
[**GetEthCustomTokenRisk**](EthereumApi.md#GetEthCustomTokenRisk) | **Post** /insights/eth/getCustomTokenRisk | Custom Volatility &amp; Risk Metrics by Token
[**GetEthDefaultCollectionRisk**](EthereumApi.md#GetEthDefaultCollectionRisk) | **Post** /insights/eth/getDefaultCollectionRisk | Default Volatility &amp; Risk Metrics by Collection
[**GetEthDefaultTokenRisk**](EthereumApi.md#GetEthDefaultTokenRisk) | **Post** /insights/eth/getDefaultTokenRisk | Default Volatility &amp; Risk Metrics by Token
[**GetEthEnsLookup**](EthereumApi.md#GetEthEnsLookup) | **Post** /data/eth/getEnsLookup | ENS Lookup
[**GetEthHistoricalEvents**](EthereumApi.md#GetEthHistoricalEvents) | **Post** /data/eth/getHistoricalEvents | Marketplace Activity by Collection
[**GetEthHistoricalTransactions**](EthereumApi.md#GetEthHistoricalTransactions) | **Post** /data/eth/getHistoricalTransactions | Historical Transactions by Collection
[**GetEthLeaderBoard**](EthereumApi.md#GetEthLeaderBoard) | **Post** /analytics/eth/getLeaderBoard | Ethereum Leaderboard by Collection
[**GetEthLiveListings**](EthereumApi.md#GetEthLiveListings) | **Post** /data/eth/getLiveListings | Collection live listings
[**GetEthLiveOffers**](EthereumApi.md#GetEthLiveOffers) | **Post** /data/eth/getLiveOffers | Collection live offers
[**GetEthMarketplaceData**](EthereumApi.md#GetEthMarketplaceData) | **Post** /data/eth/getMarketplaceData | Collection Summary by Marketplace
[**GetEthMarketplaceFloorPrice**](EthereumApi.md#GetEthMarketplaceFloorPrice) | **Post** /data/eth/getMarketplaceFloorPrice | Marketplace Floor Price by Collection
[**GetEthMarketplaceTraitData**](EthereumApi.md#GetEthMarketplaceTraitData) | **Post** /data/eth/getMarketplaceTraitData | Collection Listings by Trait &amp; Marketplace
[**GetEthRarity**](EthereumApi.md#GetEthRarity) | **Post** /analytics/eth/getRarity | Token Rarity by Collection
[**GetEthTokenAppraisal**](EthereumApi.md#GetEthTokenAppraisal) | **Post** /insights/eth/getTokenAppraisal | Liquidation &amp; Appraisal Estimate by Token
[**GetEthTokenForecasts**](EthereumApi.md#GetEthTokenForecasts) | **Post** /insights/eth/getTokenForecasts | Price Forecast by Token
[**GetEthTokenSummary**](EthereumApi.md#GetEthTokenSummary) | **Post** /analytics/eth/getTokenSummary | Summary Statistics by Token
[**GetEthTokenTransactions**](EthereumApi.md#GetEthTokenTransactions) | **Post** /data/eth/getTokenTransactions | Transactions by Token
[**GetEthTokens**](EthereumApi.md#GetEthTokens) | **Post** /data/eth/getTokens | Tokens by Collection
[**GetEthWalletLabels**](EthereumApi.md#GetEthWalletLabels) | **Post** /insights/eth/getWalletLabels | Wallet Activity Labels
[**GetEthWalletNFTs**](EthereumApi.md#GetEthWalletNFTs) | **Post** /data/eth/getWalletNFTs | Tokens Owned by Wallet
[**GetEthWalletTransactions**](EthereumApi.md#GetEthWalletTransactions) | **Post** /data/eth/getWalletTransactions | Historical Token Transactions by Wallet
[**GetEthWalletValuation**](EthereumApi.md#GetEthWalletValuation) | **Post** /insights/eth/getWalletValuation | Value All Tokens Owned by Wallet
[**GetEthWashTrade**](EthereumApi.md#GetEthWashTrade) | **Post** /analytics/eth/getWashTrade | Wash Trades by Transaction
[**GetEthWashTransactions**](EthereumApi.md#GetEthWashTransactions) | **Post** /analytics/eth/getWashTransactions | Wash Trades by Collection



## GetEthCollectionFloorPriceOHLC

> GetEthCollectionFloorPriceOHLC(ctx).GetEthCollectionFloorPriceOHLCRequest(getEthCollectionFloorPriceOHLCRequest).Execute()

Intraday Marketplace Floor Price by Collection



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
    getEthCollectionFloorPriceOHLCRequest := *openapiclient.NewGetEthCollectionFloorPriceOHLCRequest("0x0c2e57efddba8c768147d1fdf9176a0a6ebd5d83") // GetEthCollectionFloorPriceOHLCRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthCollectionFloorPriceOHLC(context.Background()).GetEthCollectionFloorPriceOHLCRequest(getEthCollectionFloorPriceOHLCRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthCollectionFloorPriceOHLC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthCollectionFloorPriceOHLCRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthCollectionFloorPriceOHLCRequest** | [**GetEthCollectionFloorPriceOHLCRequest**](GetEthCollectionFloorPriceOHLCRequest.md) |  | 

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


## GetEthCollectionForecasts

> GetEthCollectionForecasts(ctx).GetEthCollectionForecastsRequest(getEthCollectionForecastsRequest).Execute()

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
    getEthCollectionForecastsRequest := *openapiclient.NewGetEthCollectionForecastsRequest("0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb") // GetEthCollectionForecastsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthCollectionForecasts(context.Background()).GetEthCollectionForecastsRequest(getEthCollectionForecastsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthCollectionForecasts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthCollectionForecastsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthCollectionForecastsRequest** | [**GetEthCollectionForecastsRequest**](GetEthCollectionForecastsRequest.md) |  | 

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


## GetEthCollectionListingsOHLC

> GetEthCollectionListingsOHLC(ctx).GetEthCollectionListingsOHLCRequest(getEthCollectionListingsOHLCRequest).Execute()

Collection Floor Price and Listings Candlesticks



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
    getEthCollectionListingsOHLCRequest := *openapiclient.NewGetEthCollectionListingsOHLCRequest("0x7bd29408f11d2bfc23c34f18275bbf23bb716bc7") // GetEthCollectionListingsOHLCRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthCollectionListingsOHLC(context.Background()).GetEthCollectionListingsOHLCRequest(getEthCollectionListingsOHLCRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthCollectionListingsOHLC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthCollectionListingsOHLCRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthCollectionListingsOHLCRequest** | [**GetEthCollectionListingsOHLCRequest**](GetEthCollectionListingsOHLCRequest.md) |  | 

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


## GetEthCollectionOwners

> GetEthCollectionOwners(ctx).GetEthCollectionOwnersRequest(getEthCollectionOwnersRequest).Execute()

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
    getEthCollectionOwnersRequest := *openapiclient.NewGetEthCollectionOwnersRequest("0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb") // GetEthCollectionOwnersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthCollectionOwners(context.Background()).GetEthCollectionOwnersRequest(getEthCollectionOwnersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthCollectionOwners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthCollectionOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthCollectionOwnersRequest** | [**GetEthCollectionOwnersRequest**](GetEthCollectionOwnersRequest.md) |  | 

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


## GetEthCollectionPriceDiff

> GetEthCollectionPriceDiff(ctx).GetEthCollectionPriceDiffRequest(getEthCollectionPriceDiffRequest).Execute()

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
    getEthCollectionPriceDiffRequest := *openapiclient.NewGetEthCollectionPriceDiffRequest("0x7bd29408f11d2bfc23c34f18275bbf23bb716bc7") // GetEthCollectionPriceDiffRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthCollectionPriceDiff(context.Background()).GetEthCollectionPriceDiffRequest(getEthCollectionPriceDiffRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthCollectionPriceDiff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthCollectionPriceDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthCollectionPriceDiffRequest** | [**GetEthCollectionPriceDiffRequest**](GetEthCollectionPriceDiffRequest.md) |  | 

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


## GetEthCollectionSalesOHLC

> GetEthCollectionSalesOHLC(ctx).GetEthCollectionSalesOHLCRequest(getEthCollectionSalesOHLCRequest).Execute()

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
    getEthCollectionSalesOHLCRequest := *openapiclient.NewGetEthCollectionSalesOHLCRequest("0x7bd29408f11d2bfc23c34f18275bbf23bb716bc7") // GetEthCollectionSalesOHLCRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthCollectionSalesOHLC(context.Background()).GetEthCollectionSalesOHLCRequest(getEthCollectionSalesOHLCRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthCollectionSalesOHLC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthCollectionSalesOHLCRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthCollectionSalesOHLCRequest** | [**GetEthCollectionSalesOHLCRequest**](GetEthCollectionSalesOHLCRequest.md) |  | 

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


## GetEthCollectionSummary

> GetEthCollectionSummary(ctx).GetEthCollectionSummaryRequest(getEthCollectionSummaryRequest).Execute()

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
    getEthCollectionSummaryRequest := *openapiclient.NewGetEthCollectionSummaryRequest("0x3f5fb35468e9834a43dca1c160c69eaae78b6360") // GetEthCollectionSummaryRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthCollectionSummary(context.Background()).GetEthCollectionSummaryRequest(getEthCollectionSummaryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthCollectionSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthCollectionSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthCollectionSummaryRequest** | [**GetEthCollectionSummaryRequest**](GetEthCollectionSummaryRequest.md) |  | 

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


## GetEthCollectionTransactions

> GetEthCollectionTransactions(ctx).GetEthCollectionTransactionsRequest(getEthCollectionTransactionsRequest).Execute()

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
    getEthCollectionTransactionsRequest := *openapiclient.NewGetEthCollectionTransactionsRequest("0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb") // GetEthCollectionTransactionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthCollectionTransactions(context.Background()).GetEthCollectionTransactionsRequest(getEthCollectionTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthCollectionTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthCollectionTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthCollectionTransactionsRequest** | [**GetEthCollectionTransactionsRequest**](GetEthCollectionTransactionsRequest.md) |  | 

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


## GetEthCollections

> GetEthCollections(ctx).GetEthCollectionsRequest(getEthCollectionsRequest).Execute()

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
    getEthCollectionsRequest := *openapiclient.NewGetEthCollectionsRequest() // GetEthCollectionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthCollections(context.Background()).GetEthCollectionsRequest(getEthCollectionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthCollectionsRequest** | [**GetEthCollectionsRequest**](GetEthCollectionsRequest.md) |  | 

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


## GetEthCustomCollectionRisk

> GetEthCustomCollectionRisk(ctx).GetEthCustomCollectionRiskRequest(getEthCustomCollectionRiskRequest).Execute()

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
    getEthCustomCollectionRiskRequest := *openapiclient.NewGetEthCustomCollectionRiskRequest("0x3f5fb35468e9834a43dca1c160c69eaae78b6360", "6M") // GetEthCustomCollectionRiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthCustomCollectionRisk(context.Background()).GetEthCustomCollectionRiskRequest(getEthCustomCollectionRiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthCustomCollectionRisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthCustomCollectionRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthCustomCollectionRiskRequest** | [**GetEthCustomCollectionRiskRequest**](GetEthCustomCollectionRiskRequest.md) |  | 

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


## GetEthCustomTokenRisk

> GetEthCustomTokenRisk(ctx).GetEthCustomTokenRiskRequest(getEthCustomTokenRiskRequest).Execute()

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
    getEthCustomTokenRiskRequest := *openapiclient.NewGetEthCustomTokenRiskRequest("0x3f5fb35468e9834a43dca1c160c69eaae78b6360", []string{"303"}, "6M") // GetEthCustomTokenRiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthCustomTokenRisk(context.Background()).GetEthCustomTokenRiskRequest(getEthCustomTokenRiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthCustomTokenRisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthCustomTokenRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthCustomTokenRiskRequest** | [**GetEthCustomTokenRiskRequest**](GetEthCustomTokenRiskRequest.md) |  | 

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


## GetEthDefaultCollectionRisk

> GetEthDefaultCollectionRisk(ctx).GetEthDefaultCollectionRiskRequest(getEthDefaultCollectionRiskRequest).Execute()

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
    getEthDefaultCollectionRiskRequest := *openapiclient.NewGetEthDefaultCollectionRiskRequest("0x3f5fb35468e9834a43dca1c160c69eaae78b6360", "6M") // GetEthDefaultCollectionRiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthDefaultCollectionRisk(context.Background()).GetEthDefaultCollectionRiskRequest(getEthDefaultCollectionRiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthDefaultCollectionRisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthDefaultCollectionRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthDefaultCollectionRiskRequest** | [**GetEthDefaultCollectionRiskRequest**](GetEthDefaultCollectionRiskRequest.md) |  | 

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


## GetEthDefaultTokenRisk

> GetEthDefaultTokenRisk(ctx).GetEthDefaultTokenRiskRequest(getEthDefaultTokenRiskRequest).Execute()

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
    getEthDefaultTokenRiskRequest := *openapiclient.NewGetEthDefaultTokenRiskRequest("0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb", []string{"6000"}, "6M") // GetEthDefaultTokenRiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthDefaultTokenRisk(context.Background()).GetEthDefaultTokenRiskRequest(getEthDefaultTokenRiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthDefaultTokenRisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthDefaultTokenRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthDefaultTokenRiskRequest** | [**GetEthDefaultTokenRiskRequest**](GetEthDefaultTokenRiskRequest.md) |  | 

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


## GetEthEnsLookup

> GetEthEnsLookup(ctx).GetEthEnsLookupRequest(getEthEnsLookupRequest).Execute()

ENS Lookup



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
    getEthEnsLookupRequest := *openapiclient.NewGetEthEnsLookupRequest() // GetEthEnsLookupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthEnsLookup(context.Background()).GetEthEnsLookupRequest(getEthEnsLookupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthEnsLookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthEnsLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthEnsLookupRequest** | [**GetEthEnsLookupRequest**](GetEthEnsLookupRequest.md) |  | 

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


## GetEthHistoricalEvents

> GetEthHistoricalEvents(ctx).GetEthHistoricalEventsRequest(getEthHistoricalEventsRequest).Execute()

Marketplace Activity by Collection



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
    getEthHistoricalEventsRequest := *openapiclient.NewGetEthHistoricalEventsRequest("0x3fe1a4c1481c8351e91b64d5c398b159de07cbc5") // GetEthHistoricalEventsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthHistoricalEvents(context.Background()).GetEthHistoricalEventsRequest(getEthHistoricalEventsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthHistoricalEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthHistoricalEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthHistoricalEventsRequest** | [**GetEthHistoricalEventsRequest**](GetEthHistoricalEventsRequest.md) |  | 

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


## GetEthHistoricalTransactions

> GetEthHistoricalTransactions(ctx).GetEthHistoricalTransactionsRequest(getEthHistoricalTransactionsRequest).Execute()

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
    getEthHistoricalTransactionsRequest := *openapiclient.NewGetEthHistoricalTransactionsRequest("0x3f5fb35468e9834a43dca1c160c69eaae78b6360") // GetEthHistoricalTransactionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthHistoricalTransactions(context.Background()).GetEthHistoricalTransactionsRequest(getEthHistoricalTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthHistoricalTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthHistoricalTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthHistoricalTransactionsRequest** | [**GetEthHistoricalTransactionsRequest**](GetEthHistoricalTransactionsRequest.md) |  | 

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


## GetEthLeaderBoard

> GetEthLeaderBoard(ctx).GetEthLeaderBoardRequest(getEthLeaderBoardRequest).Execute()

Ethereum Leaderboard by Collection



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
    r, err := apiClient.EthereumApi.GetEthLeaderBoard(context.Background()).GetEthLeaderBoardRequest(getEthLeaderBoardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthLeaderBoard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthLeaderBoardRequest struct via the builder pattern


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


## GetEthLiveListings

> GetEthLiveListings(ctx).GetEthLiveListingsRequest(getEthLiveListingsRequest).Execute()

Collection live listings



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
    getEthLiveListingsRequest := *openapiclient.NewGetEthLiveListingsRequest("0x01cc1bd2dece86e54481f1e6aeffc2fcc6915480") // GetEthLiveListingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthLiveListings(context.Background()).GetEthLiveListingsRequest(getEthLiveListingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthLiveListings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthLiveListingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthLiveListingsRequest** | [**GetEthLiveListingsRequest**](GetEthLiveListingsRequest.md) |  | 

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


## GetEthLiveOffers

> GetEthLiveOffers(ctx).GetEthLiveListingsRequest(getEthLiveListingsRequest).Execute()

Collection live offers



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
    getEthLiveListingsRequest := *openapiclient.NewGetEthLiveListingsRequest("0x01cc1bd2dece86e54481f1e6aeffc2fcc6915480") // GetEthLiveListingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthLiveOffers(context.Background()).GetEthLiveListingsRequest(getEthLiveListingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthLiveOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthLiveOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthLiveListingsRequest** | [**GetEthLiveListingsRequest**](GetEthLiveListingsRequest.md) |  | 

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


## GetEthMarketplaceData

> GetEthMarketplaceData(ctx).GetEthMarketplaceDataRequest(getEthMarketplaceDataRequest).Execute()

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
    getEthMarketplaceDataRequest := *openapiclient.NewGetEthMarketplaceDataRequest([]string{"0xd70f41dd5875eee7fa9dd8048567bc932124a8d2"}) // GetEthMarketplaceDataRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthMarketplaceData(context.Background()).GetEthMarketplaceDataRequest(getEthMarketplaceDataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthMarketplaceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthMarketplaceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthMarketplaceDataRequest** | [**GetEthMarketplaceDataRequest**](GetEthMarketplaceDataRequest.md) |  | 

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


## GetEthMarketplaceFloorPrice

> GetEthMarketplaceFloorPrice(ctx).GetEthMarketplaceFloorPriceRequest(getEthMarketplaceFloorPriceRequest).Execute()

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
    getEthMarketplaceFloorPriceRequest := *openapiclient.NewGetEthMarketplaceFloorPriceRequest() // GetEthMarketplaceFloorPriceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthMarketplaceFloorPrice(context.Background()).GetEthMarketplaceFloorPriceRequest(getEthMarketplaceFloorPriceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthMarketplaceFloorPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthMarketplaceFloorPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthMarketplaceFloorPriceRequest** | [**GetEthMarketplaceFloorPriceRequest**](GetEthMarketplaceFloorPriceRequest.md) |  | 

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


## GetEthMarketplaceTraitData

> GetEthMarketplaceTraitData(ctx).GetEthMarketplaceTraitDataRequest(getEthMarketplaceTraitDataRequest).Execute()

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
    getEthMarketplaceTraitDataRequest := *openapiclient.NewGetEthMarketplaceTraitDataRequest("0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d") // GetEthMarketplaceTraitDataRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthMarketplaceTraitData(context.Background()).GetEthMarketplaceTraitDataRequest(getEthMarketplaceTraitDataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthMarketplaceTraitData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthMarketplaceTraitDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthMarketplaceTraitDataRequest** | [**GetEthMarketplaceTraitDataRequest**](GetEthMarketplaceTraitDataRequest.md) |  | 

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


## GetEthRarity

> GetEthRarity(ctx).GetEthRarityRequest(getEthRarityRequest).Execute()

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
    getEthRarityRequest := *openapiclient.NewGetEthRarityRequest("0x716039ab9ce2780e35450b86dc6420f22460c380") // GetEthRarityRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthRarity(context.Background()).GetEthRarityRequest(getEthRarityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthRarity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthRarityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthRarityRequest** | [**GetEthRarityRequest**](GetEthRarityRequest.md) |  | 

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


## GetEthTokenAppraisal

> GetEthTokenAppraisal(ctx).GetEthTokenAppraisalRequest(getEthTokenAppraisalRequest).Execute()

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
    getEthTokenAppraisalRequest := *openapiclient.NewGetEthTokenAppraisalRequest("0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d", []string{"40"}) // GetEthTokenAppraisalRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthTokenAppraisal(context.Background()).GetEthTokenAppraisalRequest(getEthTokenAppraisalRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthTokenAppraisal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthTokenAppraisalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthTokenAppraisalRequest** | [**GetEthTokenAppraisalRequest**](GetEthTokenAppraisalRequest.md) |  | 

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


## GetEthTokenForecasts

> GetEthTokenForecasts(ctx).GetEthTokenForecastsRequest(getEthTokenForecastsRequest).Execute()

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
    getEthTokenForecastsRequest := *openapiclient.NewGetEthTokenForecastsRequest("0x3f5fb35468e9834a43dca1c160c69eaae78b6360", []string{"1722"}) // GetEthTokenForecastsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthTokenForecasts(context.Background()).GetEthTokenForecastsRequest(getEthTokenForecastsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthTokenForecasts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthTokenForecastsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthTokenForecastsRequest** | [**GetEthTokenForecastsRequest**](GetEthTokenForecastsRequest.md) |  | 

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


## GetEthTokenSummary

> GetEthTokenSummary(ctx).GetEthTokenSummaryRequest(getEthTokenSummaryRequest).Execute()

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
    getEthTokenSummaryRequest := *openapiclient.NewGetEthTokenSummaryRequest("0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb", []string{"6000"}) // GetEthTokenSummaryRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthTokenSummary(context.Background()).GetEthTokenSummaryRequest(getEthTokenSummaryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthTokenSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthTokenSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthTokenSummaryRequest** | [**GetEthTokenSummaryRequest**](GetEthTokenSummaryRequest.md) |  | 

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


## GetEthTokenTransactions

> GetEthTokenTransactions(ctx).GetEthTokenTransactionsRequest(getEthTokenTransactionsRequest).Execute()

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
    getEthTokenTransactionsRequest := *openapiclient.NewGetEthTokenTransactionsRequest("0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb", "6000") // GetEthTokenTransactionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthTokenTransactions(context.Background()).GetEthTokenTransactionsRequest(getEthTokenTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthTokenTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthTokenTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthTokenTransactionsRequest** | [**GetEthTokenTransactionsRequest**](GetEthTokenTransactionsRequest.md) |  | 

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


## GetEthTokens

> GetEthTokens(ctx).GetEthTokensRequest(getEthTokensRequest).Execute()

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
    getEthTokensRequest := *openapiclient.NewGetEthTokensRequest("0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb") // GetEthTokensRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthTokens(context.Background()).GetEthTokensRequest(getEthTokensRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthTokensRequest** | [**GetEthTokensRequest**](GetEthTokensRequest.md) |  | 

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


## GetEthWalletLabels

> GetEthWalletLabels(ctx).GetEthWalletLabelsRequest(getEthWalletLabelsRequest).Execute()

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
    r, err := apiClient.EthereumApi.GetEthWalletLabels(context.Background()).GetEthWalletLabelsRequest(getEthWalletLabelsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthWalletLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthWalletLabelsRequest struct via the builder pattern


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


## GetEthWalletNFTs

> GetEthWalletNFTs(ctx).GetEthWalletNFTsRequest(getEthWalletNFTsRequest).Execute()

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
    getEthWalletNFTsRequest := *openapiclient.NewGetEthWalletNFTsRequest("0xab0cda4cc21207fd9433356afe9428b6fac8a8a5") // GetEthWalletNFTsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthWalletNFTs(context.Background()).GetEthWalletNFTsRequest(getEthWalletNFTsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthWalletNFTs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthWalletNFTsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthWalletNFTsRequest** | [**GetEthWalletNFTsRequest**](GetEthWalletNFTsRequest.md) |  | 

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


## GetEthWalletTransactions

> GetEthWalletTransactions(ctx).GetEthWalletTransactionsRequest(getEthWalletTransactionsRequest).Execute()

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
    getEthWalletTransactionsRequest := *openapiclient.NewGetEthWalletTransactionsRequest("0xe724e14c6b7599b710804df390e39928abfed082") // GetEthWalletTransactionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthWalletTransactions(context.Background()).GetEthWalletTransactionsRequest(getEthWalletTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthWalletTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthWalletTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthWalletTransactionsRequest** | [**GetEthWalletTransactionsRequest**](GetEthWalletTransactionsRequest.md) |  | 

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


## GetEthWalletValuation

> GetEthWalletValuation(ctx).GetPolWalletNFTsRequest(getPolWalletNFTsRequest).Execute()

Value All Tokens Owned by Wallet



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
    r, err := apiClient.EthereumApi.GetEthWalletValuation(context.Background()).GetPolWalletNFTsRequest(getPolWalletNFTsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthWalletValuation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthWalletValuationRequest struct via the builder pattern


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


## GetEthWashTrade

> GetEthWashTrade(ctx).GetEthWashTradeRequest(getEthWashTradeRequest).Execute()

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
    getEthWashTradeRequest := *openapiclient.NewGetEthWashTradeRequest("0x8a67b9ec06d01ebd2f8363f4a19cb6b55e3fa409877f18d5716716cce457676d") // GetEthWashTradeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthWashTrade(context.Background()).GetEthWashTradeRequest(getEthWashTradeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthWashTrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthWashTradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthWashTradeRequest** | [**GetEthWashTradeRequest**](GetEthWashTradeRequest.md) |  | 

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


## GetEthWashTransactions

> GetEthWashTransactions(ctx).GetEthWashTransactionsRequest(getEthWashTransactionsRequest).Execute()

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
    getEthWashTransactionsRequest := *openapiclient.NewGetEthWashTransactionsRequest("0x3bf2922f4520a8ba0c2efc3d2a1539678dad5e9d") // GetEthWashTransactionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthereumApi.GetEthWashTransactions(context.Background()).GetEthWashTransactionsRequest(getEthWashTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthereumApi.GetEthWashTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthWashTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEthWashTransactionsRequest** | [**GetEthWashTransactionsRequest**](GetEthWashTransactionsRequest.md) |  | 

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

