# Go API client for openapi

Data and insights APIs, webooks, and dashboards enabling businesses to launch tokenized products in seconds.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.prod.gallop.run/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*EthereumApi* | [**GetEthCollectionFloorPriceOHLC**](docs/EthereumApi.md#getethcollectionfloorpriceohlc) | **Post** /analytics/eth/getCollectionFloorPriceOHLC | Intraday Marketplace Floor Price by Collection
*EthereumApi* | [**GetEthCollectionForecasts**](docs/EthereumApi.md#getethcollectionforecasts) | **Post** /insights/eth/getCollectionForecasts | Price Forecast by Collection
*EthereumApi* | [**GetEthCollectionListingsOHLC**](docs/EthereumApi.md#getethcollectionlistingsohlc) | **Post** /analytics/eth/getCollectionListingsOHLC | Collection Floor Price and Listings Candlesticks
*EthereumApi* | [**GetEthCollectionOwners**](docs/EthereumApi.md#getethcollectionowners) | **Post** /data/eth/getCollectionOwners | Wallet Owners by Collection
*EthereumApi* | [**GetEthCollectionPriceDiff**](docs/EthereumApi.md#getethcollectionpricediff) | **Post** /analytics/eth/getCollectionPriceDiff | Price Differentiation by Trait
*EthereumApi* | [**GetEthCollectionSalesOHLC**](docs/EthereumApi.md#getethcollectionsalesohlc) | **Post** /analytics/eth/getCollectionSalesOHLC | Collection Sales Price Candlesticks
*EthereumApi* | [**GetEthCollectionSummary**](docs/EthereumApi.md#getethcollectionsummary) | **Post** /analytics/eth/getCollectionSummary | Summary Statistics by Collection
*EthereumApi* | [**GetEthCollectionTransactions**](docs/EthereumApi.md#getethcollectiontransactions) | **Post** /data/eth/getCollectionTransactions | Transactions by Collection
*EthereumApi* | [**GetEthCollections**](docs/EthereumApi.md#getethcollections) | **Post** /data/eth/getCollections | Aggregated Collections Supported by Gallop
*EthereumApi* | [**GetEthCustomCollectionRisk**](docs/EthereumApi.md#getethcustomcollectionrisk) | **Post** /insights/eth/getCustomCollectionRisk | Custom Volatility &amp; Risk Metrics by Collection
*EthereumApi* | [**GetEthCustomTokenRisk**](docs/EthereumApi.md#getethcustomtokenrisk) | **Post** /insights/eth/getCustomTokenRisk | Custom Volatility &amp; Risk Metrics by Token
*EthereumApi* | [**GetEthDefaultCollectionRisk**](docs/EthereumApi.md#getethdefaultcollectionrisk) | **Post** /insights/eth/getDefaultCollectionRisk | Default Volatility &amp; Risk Metrics by Collection
*EthereumApi* | [**GetEthDefaultTokenRisk**](docs/EthereumApi.md#getethdefaulttokenrisk) | **Post** /insights/eth/getDefaultTokenRisk | Default Volatility &amp; Risk Metrics by Token
*EthereumApi* | [**GetEthEnsLookup**](docs/EthereumApi.md#getethenslookup) | **Post** /data/eth/getEnsLookup | ENS Lookup
*EthereumApi* | [**GetEthHistoricalEvents**](docs/EthereumApi.md#getethhistoricalevents) | **Post** /data/eth/getHistoricalEvents | Marketplace Activity by Collection
*EthereumApi* | [**GetEthHistoricalTransactions**](docs/EthereumApi.md#getethhistoricaltransactions) | **Post** /data/eth/getHistoricalTransactions | Historical Transactions by Collection
*EthereumApi* | [**GetEthLeaderBoard**](docs/EthereumApi.md#getethleaderboard) | **Post** /analytics/eth/getLeaderBoard | Ethereum Leaderboard by Collection
*EthereumApi* | [**GetEthMarketplaceData**](docs/EthereumApi.md#getethmarketplacedata) | **Post** /data/eth/getMarketplaceData | Collection Summary by Marketplace
*EthereumApi* | [**GetEthMarketplaceFloorPrice**](docs/EthereumApi.md#getethmarketplacefloorprice) | **Post** /data/eth/getMarketplaceFloorPrice | Marketplace Floor Price by Collection
*EthereumApi* | [**GetEthMarketplaceTraitData**](docs/EthereumApi.md#getethmarketplacetraitdata) | **Post** /data/eth/getMarketplaceTraitData | Collection Listings by Trait &amp; Marketplace
*EthereumApi* | [**GetEthRarity**](docs/EthereumApi.md#getethrarity) | **Post** /analytics/eth/getRarity | Token Rarity by Collection
*EthereumApi* | [**GetEthTokenAppraisal**](docs/EthereumApi.md#getethtokenappraisal) | **Post** /insights/eth/getTokenAppraisal | Liquidation &amp; Appraisal Estimate by Token
*EthereumApi* | [**GetEthTokenForecasts**](docs/EthereumApi.md#getethtokenforecasts) | **Post** /insights/eth/getTokenForecasts | Price Forecast by Token
*EthereumApi* | [**GetEthTokenSummary**](docs/EthereumApi.md#getethtokensummary) | **Post** /analytics/eth/getTokenSummary | Summary Statistics by Token
*EthereumApi* | [**GetEthTokenTransactions**](docs/EthereumApi.md#getethtokentransactions) | **Post** /data/eth/getTokenTransactions | Transactions by Token
*EthereumApi* | [**GetEthTokens**](docs/EthereumApi.md#getethtokens) | **Post** /data/eth/getTokens | Tokens by Collection
*EthereumApi* | [**GetEthWalletLabels**](docs/EthereumApi.md#getethwalletlabels) | **Post** /insights/eth/getWalletLabels | Wallet Activity Labels
*EthereumApi* | [**GetEthWalletNFTs**](docs/EthereumApi.md#getethwalletnfts) | **Post** /data/eth/getWalletNFTs | Tokens Owned by Wallet
*EthereumApi* | [**GetEthWalletTransactions**](docs/EthereumApi.md#getethwallettransactions) | **Post** /data/eth/getWalletTransactions | Historical Token Transactions by Wallet
*EthereumApi* | [**GetEthWalletValuation**](docs/EthereumApi.md#getethwalletvaluation) | **Post** /insights/eth/getWalletValuation | Value All Tokens Owned by Wallet
*EthereumApi* | [**GetEthWashTrade**](docs/EthereumApi.md#getethwashtrade) | **Post** /analytics/eth/getWashTrade | Wash Trades by Transaction
*EthereumApi* | [**GetEthWashTransactions**](docs/EthereumApi.md#getethwashtransactions) | **Post** /analytics/eth/getWashTransactions | Wash Trades by Collection
*PolygonApi* | [**GetPolCollectionForecasts**](docs/PolygonApi.md#getpolcollectionforecasts) | **Post** /insights/pol/getCollectionForecasts | Price Forecast by Collection
*PolygonApi* | [**GetPolCollectionOwners**](docs/PolygonApi.md#getpolcollectionowners) | **Post** /data/pol/getCollectionOwners | Wallet Owners by Collection
*PolygonApi* | [**GetPolCollectionPriceDiff**](docs/PolygonApi.md#getpolcollectionpricediff) | **Post** /analytics/pol/getCollectionPriceDiff | Price Differentiation by Trait
*PolygonApi* | [**GetPolCollectionSalesOHLC**](docs/PolygonApi.md#getpolcollectionsalesohlc) | **Post** /analytics/pol/getCollectionSalesOHLC | Collection Sales Price Candlesticks
*PolygonApi* | [**GetPolCollectionSummary**](docs/PolygonApi.md#getpolcollectionsummary) | **Post** /analytics/pol/getCollectionSummary | Summary Statistics by Collection
*PolygonApi* | [**GetPolCollectionTraits**](docs/PolygonApi.md#getpolcollectiontraits) | **Post** /data/pol/getCollectionTraits | Traits by Collection
*PolygonApi* | [**GetPolCollectionTransactions**](docs/PolygonApi.md#getpolcollectiontransactions) | **Post** /data/pol/getCollectionTransactions | Transactions by Collection
*PolygonApi* | [**GetPolCollections**](docs/PolygonApi.md#getpolcollections) | **Post** /data/pol/getCollections | Aggregated Collections Supported by Gallop
*PolygonApi* | [**GetPolCustomCollectionRisk**](docs/PolygonApi.md#getpolcustomcollectionrisk) | **Post** /insights/pol/getCustomCollectionRisk | Custom Volatility &amp; Risk Metrics by Collection
*PolygonApi* | [**GetPolCustomTokenRisk**](docs/PolygonApi.md#getpolcustomtokenrisk) | **Post** /insights/pol/getCustomTokenRisk | Custom Volatility &amp; Risk Metrics by Token
*PolygonApi* | [**GetPolDefaultCollectionRisk**](docs/PolygonApi.md#getpoldefaultcollectionrisk) | **Post** /insights/pol/getDefaultCollectionRisk | Default Volatility &amp; Risk Metrics by Collection
*PolygonApi* | [**GetPolDefaultTokenRisk**](docs/PolygonApi.md#getpoldefaulttokenrisk) | **Post** /insights/pol/getDefaultTokenRisk | Default Volatility &amp; Risk Metrics by Token
*PolygonApi* | [**GetPolHistoricalTransactions**](docs/PolygonApi.md#getpolhistoricaltransactions) | **Post** /data/pol/getHistoricalTransactions | Historical Transactions by Collection
*PolygonApi* | [**GetPolLeaderBoard**](docs/PolygonApi.md#getpolleaderboard) | **Post** /analytics/pol/getLeaderBoard | Polygon Leaderboard by Collection
*PolygonApi* | [**GetPolMarketplaceData**](docs/PolygonApi.md#getpolmarketplacedata) | **Post** /data/pol/getMarketplaceData | Collection Summary by Marketplace
*PolygonApi* | [**GetPolMarketplaceFloorPrice**](docs/PolygonApi.md#getpolmarketplacefloorprice) | **Post** /data/pol/getMarketplaceFloorPrice | Marketplace Floor Price by Collection
*PolygonApi* | [**GetPolRarity**](docs/PolygonApi.md#getpolrarity) | **Post** /analytics/pol/getRarity | Token Rarity by Collection
*PolygonApi* | [**GetPolTokenAppraisal**](docs/PolygonApi.md#getpoltokenappraisal) | **Post** /insights/pol/getTokenAppraisal | Liquidation &amp; Appraisal Estimate by Token
*PolygonApi* | [**GetPolTokenForecasts**](docs/PolygonApi.md#getpoltokenforecasts) | **Post** /insights/pol/getTokenForecasts | Price Forecast by Token
*PolygonApi* | [**GetPolTokenSummary**](docs/PolygonApi.md#getpoltokensummary) | **Post** /analytics/pol/getTokenSummary | Summary Statistics by Token
*PolygonApi* | [**GetPolTokenTransactions**](docs/PolygonApi.md#getpoltokentransactions) | **Post** /data/pol/getTokenTransactions | Transactions by Token
*PolygonApi* | [**GetPolTokens**](docs/PolygonApi.md#getpoltokens) | **Post** /data/pol/getTokens | Tokens by Collection
*PolygonApi* | [**GetPolWalletLabels**](docs/PolygonApi.md#getpolwalletlabels) | **Post** /insights/pol/getWalletLabels | Wallet Activity Labels
*PolygonApi* | [**GetPolWalletNFTs**](docs/PolygonApi.md#getpolwalletnfts) | **Post** /data/pol/getWalletNFTs | Tokens Owned by Wallet
*PolygonApi* | [**GetPolWalletTransactions**](docs/PolygonApi.md#getpolwallettransactions) | **Post** /data/pol/getWalletTransactions | Historical Token Transactions by Wallet
*PolygonApi* | [**GetPolWashTrade**](docs/PolygonApi.md#getpolwashtrade) | **Post** /analytics/pol/getWashTrade | Wash Trades by Transaction
*PolygonApi* | [**GetPolWashTransactions**](docs/PolygonApi.md#getpolwashtransactions) | **Post** /analytics/pol/getWashTransactions | Wash Trades by Collection
*SolanaApi* | [**GetSolAccountNFTs**](docs/SolanaApi.md#getsolaccountnfts) | **Post** /data/sol/getAccountNFTs | Tokens Owned by Wallet
*SolanaApi* | [**GetSolCollectionForecasts**](docs/SolanaApi.md#getsolcollectionforecasts) | **Post** /insights/sol/getCollectionForecasts | Price Forecast by Collection
*SolanaApi* | [**GetSolCollectionPriceDiff**](docs/SolanaApi.md#getsolcollectionpricediff) | **Post** /analytics/sol/getCollectionPriceDiff | Price Differentiation by Trait
*SolanaApi* | [**GetSolCollectionSalesOHLC**](docs/SolanaApi.md#getsolcollectionsalesohlc) | **Post** /analytics/sol/getCollectionSalesOHLC | Collection Sales Price Candlesticks
*SolanaApi* | [**GetSolCollectionSummary**](docs/SolanaApi.md#getsolcollectionsummary) | **Post** /analytics/sol/getCollectionSummary | Summary Statistics by Collection
*SolanaApi* | [**GetSolCollectionTraits**](docs/SolanaApi.md#getsolcollectiontraits) | **Post** /data/sol/getCollectionTraits | Traits by Collection
*SolanaApi* | [**GetSolCollectionTransactions**](docs/SolanaApi.md#getsolcollectiontransactions) | **Post** /data/sol/getCollectionTransactions | Transactions by Collections
*SolanaApi* | [**GetSolCollections**](docs/SolanaApi.md#getsolcollections) | **Post** /data/sol/getCollections | Aggregated Collections Supported by Gallop
*SolanaApi* | [**GetSolCustomCollectionRisk**](docs/SolanaApi.md#getsolcustomcollectionrisk) | **Post** /insights/sol/getCustomCollectionRisk | Custom Volatility &amp; Risk Metrics by Collection
*SolanaApi* | [**GetSolCustomTokenRisk**](docs/SolanaApi.md#getsolcustomtokenrisk) | **Post** /insights/sol/getCustomTokenRisk | Custom Volatility &amp; Risk Metrics by Token
*SolanaApi* | [**GetSolDefaultCollectionRisk**](docs/SolanaApi.md#getsoldefaultcollectionrisk) | **Post** /insights/sol/getDefaultCollectionRisk | Default Volatility &amp; Risk Metrics by Collection
*SolanaApi* | [**GetSolDefaultTokenRisk**](docs/SolanaApi.md#getsoldefaulttokenrisk) | **Post** /insights/sol/getDefaultTokenRisk | Default Volatility &amp; Risk Metrics by Token
*SolanaApi* | [**GetSolHistoricalTransactions**](docs/SolanaApi.md#getsolhistoricaltransactions) | **Post** /data/sol/getHistoricalTransactions | Historical Transactions by Collection
*SolanaApi* | [**GetSolMarketplaceData**](docs/SolanaApi.md#getsolmarketplacedata) | **Post** /data/sol/getMarketplaceData | Collection Summary by Marketplace
*SolanaApi* | [**GetSolMarketplaceFloorPrice**](docs/SolanaApi.md#getsolmarketplacefloorprice) | **Post** /data/sol/getMarketplaceFloorPrice | Marketplace Floor Price by Collection
*SolanaApi* | [**GetSolMarketplaceTraitData**](docs/SolanaApi.md#getsolmarketplacetraitdata) | **Post** /data/sol/getMarketplaceTraitData | Collection Listings by Trait &amp; Marketplace
*SolanaApi* | [**GetSolNFTAccount**](docs/SolanaApi.md#getsolnftaccount) | **Post** /data/sol/getNFTAccount | Wallet Owners by Token
*SolanaApi* | [**GetSolRarity**](docs/SolanaApi.md#getsolrarity) | **Post** /analytics/sol/getRarity | Token Rarity by Collection
*SolanaApi* | [**GetSolTokenAppraisal**](docs/SolanaApi.md#getsoltokenappraisal) | **Post** /insights/sol/getTokenAppraisal | Liquidation &amp; Appraisal Estimate by Token
*SolanaApi* | [**GetSolTokenForecasts**](docs/SolanaApi.md#getsoltokenforecasts) | **Post** /insights/sol/getTokenForecasts | Price Forecast by Token
*SolanaApi* | [**GetSolTokenSummary**](docs/SolanaApi.md#getsoltokensummary) | **Post** /analytics/sol/getTokenSummary | Summary Statistics by Token
*SolanaApi* | [**GetSolTokenTransactions**](docs/SolanaApi.md#getsoltokentransactions) | **Post** /data/sol/getTokenTransactions | Transactions by Token
*SolanaApi* | [**GetSolTokens**](docs/SolanaApi.md#getsoltokens) | **Post** /data/sol/getTokens | Tokens by Collection
*SolanaApi* | [**GetSolWashTrade**](docs/SolanaApi.md#getsolwashtrade) | **Post** /analytics/sol/getWashTrade | Wash Trades by Transaction
*SolanaApi* | [**GetSolWashTransactions**](docs/SolanaApi.md#getsolwashtransactions) | **Post** /analytics/sol/getWashTransactions | Wash Trades by Collection
*StarknetApi* | [**GetSknMarketplaceData**](docs/StarknetApi.md#getsknmarketplacedata) | **Post** /data/skn/getMarketplaceData | Gallop Marketplace Data
*StarknetApi* | [**GetSknMarketplaceFloorPrice**](docs/StarknetApi.md#getsknmarketplacefloorprice) | **Post** /data/skn/getMarketplaceFloorPrice | Marketplace Floor Price by Collection


## Documentation For Models

 - [GetEthCollectionFloorPriceOHLCRequest](docs/GetEthCollectionFloorPriceOHLCRequest.md)
 - [GetEthCollectionForecastsRequest](docs/GetEthCollectionForecastsRequest.md)
 - [GetEthCollectionForecastsRequestArchParams](docs/GetEthCollectionForecastsRequestArchParams.md)
 - [GetEthCollectionListingsOHLCRequest](docs/GetEthCollectionListingsOHLCRequest.md)
 - [GetEthCollectionOwnersRequest](docs/GetEthCollectionOwnersRequest.md)
 - [GetEthCollectionPriceDiffRequest](docs/GetEthCollectionPriceDiffRequest.md)
 - [GetEthCollectionSalesOHLCRequest](docs/GetEthCollectionSalesOHLCRequest.md)
 - [GetEthCollectionSummaryRequest](docs/GetEthCollectionSummaryRequest.md)
 - [GetEthCollectionTransactionsRequest](docs/GetEthCollectionTransactionsRequest.md)
 - [GetEthCollectionsRequest](docs/GetEthCollectionsRequest.md)
 - [GetEthCustomCollectionRiskRequest](docs/GetEthCustomCollectionRiskRequest.md)
 - [GetEthCustomCollectionRiskRequestArcParams](docs/GetEthCustomCollectionRiskRequestArcParams.md)
 - [GetEthCustomCollectionRiskRequestGarParams](docs/GetEthCustomCollectionRiskRequestGarParams.md)
 - [GetEthCustomCollectionRiskRequestHarParams](docs/GetEthCustomCollectionRiskRequestHarParams.md)
 - [GetEthCustomTokenRiskRequest](docs/GetEthCustomTokenRiskRequest.md)
 - [GetEthDefaultCollectionRiskRequest](docs/GetEthDefaultCollectionRiskRequest.md)
 - [GetEthDefaultTokenRiskRequest](docs/GetEthDefaultTokenRiskRequest.md)
 - [GetEthEnsLookupRequest](docs/GetEthEnsLookupRequest.md)
 - [GetEthHistoricalEventsRequest](docs/GetEthHistoricalEventsRequest.md)
 - [GetEthHistoricalTransactionsRequest](docs/GetEthHistoricalTransactionsRequest.md)
 - [GetEthLeaderBoardRequest](docs/GetEthLeaderBoardRequest.md)
 - [GetEthMarketplaceDataRequest](docs/GetEthMarketplaceDataRequest.md)
 - [GetEthMarketplaceFloorPriceRequest](docs/GetEthMarketplaceFloorPriceRequest.md)
 - [GetEthMarketplaceTraitDataRequest](docs/GetEthMarketplaceTraitDataRequest.md)
 - [GetEthRarityRequest](docs/GetEthRarityRequest.md)
 - [GetEthTokenAppraisalRequest](docs/GetEthTokenAppraisalRequest.md)
 - [GetEthTokenForecastsRequest](docs/GetEthTokenForecastsRequest.md)
 - [GetEthTokenSummaryRequest](docs/GetEthTokenSummaryRequest.md)
 - [GetEthTokenTransactionsRequest](docs/GetEthTokenTransactionsRequest.md)
 - [GetEthTokensRequest](docs/GetEthTokensRequest.md)
 - [GetEthWalletLabelsRequest](docs/GetEthWalletLabelsRequest.md)
 - [GetEthWalletNFTsRequest](docs/GetEthWalletNFTsRequest.md)
 - [GetEthWalletTransactionsRequest](docs/GetEthWalletTransactionsRequest.md)
 - [GetEthWashTradeRequest](docs/GetEthWashTradeRequest.md)
 - [GetEthWashTransactionsRequest](docs/GetEthWashTransactionsRequest.md)
 - [GetPolCollectionForecastsRequest](docs/GetPolCollectionForecastsRequest.md)
 - [GetPolCollectionOwnersRequest](docs/GetPolCollectionOwnersRequest.md)
 - [GetPolCollectionPriceDiffRequest](docs/GetPolCollectionPriceDiffRequest.md)
 - [GetPolCollectionSalesOHLCRequest](docs/GetPolCollectionSalesOHLCRequest.md)
 - [GetPolCollectionSummaryRequest](docs/GetPolCollectionSummaryRequest.md)
 - [GetPolCollectionTraitsRequest](docs/GetPolCollectionTraitsRequest.md)
 - [GetPolCollectionTransactionsRequest](docs/GetPolCollectionTransactionsRequest.md)
 - [GetPolCollectionsRequest](docs/GetPolCollectionsRequest.md)
 - [GetPolCustomCollectionRiskRequest](docs/GetPolCustomCollectionRiskRequest.md)
 - [GetPolCustomTokenRiskRequest](docs/GetPolCustomTokenRiskRequest.md)
 - [GetPolDefaultCollectionRiskRequest](docs/GetPolDefaultCollectionRiskRequest.md)
 - [GetPolDefaultTokenRiskRequest](docs/GetPolDefaultTokenRiskRequest.md)
 - [GetPolHistoricalTransactionsRequest](docs/GetPolHistoricalTransactionsRequest.md)
 - [GetPolMarketplaceDataRequest](docs/GetPolMarketplaceDataRequest.md)
 - [GetPolMarketplaceFloorPriceRequest](docs/GetPolMarketplaceFloorPriceRequest.md)
 - [GetPolRarityRequest](docs/GetPolRarityRequest.md)
 - [GetPolTokenAppraisalRequest](docs/GetPolTokenAppraisalRequest.md)
 - [GetPolTokenForecastsRequest](docs/GetPolTokenForecastsRequest.md)
 - [GetPolTokenSummaryRequest](docs/GetPolTokenSummaryRequest.md)
 - [GetPolTokenTransactionsRequest](docs/GetPolTokenTransactionsRequest.md)
 - [GetPolTokensRequest](docs/GetPolTokensRequest.md)
 - [GetPolWalletNFTsRequest](docs/GetPolWalletNFTsRequest.md)
 - [GetPolWalletTransactionsRequest](docs/GetPolWalletTransactionsRequest.md)
 - [GetPolWashTradeRequest](docs/GetPolWashTradeRequest.md)
 - [GetPolWashTransactionsRequest](docs/GetPolWashTransactionsRequest.md)
 - [GetSknMarketplaceDataRequest](docs/GetSknMarketplaceDataRequest.md)
 - [GetSknMarketplaceFloorPriceRequest](docs/GetSknMarketplaceFloorPriceRequest.md)
 - [GetSolAccountNFTsRequest](docs/GetSolAccountNFTsRequest.md)
 - [GetSolCollectionForecastsRequest](docs/GetSolCollectionForecastsRequest.md)
 - [GetSolCollectionForecastsRequestArchParams](docs/GetSolCollectionForecastsRequestArchParams.md)
 - [GetSolCollectionPriceDiffRequest](docs/GetSolCollectionPriceDiffRequest.md)
 - [GetSolCollectionSalesOHLCRequest](docs/GetSolCollectionSalesOHLCRequest.md)
 - [GetSolCollectionSummaryRequest](docs/GetSolCollectionSummaryRequest.md)
 - [GetSolCollectionTraitsRequest](docs/GetSolCollectionTraitsRequest.md)
 - [GetSolCollectionTransactionsRequest](docs/GetSolCollectionTransactionsRequest.md)
 - [GetSolCollectionsRequest](docs/GetSolCollectionsRequest.md)
 - [GetSolCustomCollectionRiskRequest](docs/GetSolCustomCollectionRiskRequest.md)
 - [GetSolCustomTokenRiskRequest](docs/GetSolCustomTokenRiskRequest.md)
 - [GetSolDefaultCollectionRiskRequest](docs/GetSolDefaultCollectionRiskRequest.md)
 - [GetSolDefaultTokenRiskRequest](docs/GetSolDefaultTokenRiskRequest.md)
 - [GetSolHistoricalTransactionsRequest](docs/GetSolHistoricalTransactionsRequest.md)
 - [GetSolMarketplaceDataRequest](docs/GetSolMarketplaceDataRequest.md)
 - [GetSolMarketplaceFloorPriceRequest](docs/GetSolMarketplaceFloorPriceRequest.md)
 - [GetSolMarketplaceTraitDataRequest](docs/GetSolMarketplaceTraitDataRequest.md)
 - [GetSolNFTAccountRequest](docs/GetSolNFTAccountRequest.md)
 - [GetSolRarityRequest](docs/GetSolRarityRequest.md)
 - [GetSolTokenAppraisalRequest](docs/GetSolTokenAppraisalRequest.md)
 - [GetSolTokenForecastsRequest](docs/GetSolTokenForecastsRequest.md)
 - [GetSolTokenForecastsRequestArchParams](docs/GetSolTokenForecastsRequestArchParams.md)
 - [GetSolTokenSummaryRequest](docs/GetSolTokenSummaryRequest.md)
 - [GetSolTokenTransactionsRequest](docs/GetSolTokenTransactionsRequest.md)
 - [GetSolTokensRequest](docs/GetSolTokensRequest.md)
 - [GetSolWashTradeRequest](docs/GetSolWashTradeRequest.md)
 - [GetSolWashTransactionsRequest](docs/GetSolWashTransactionsRequest.md)


## Documentation For Authorization



### api_key

- **Type**: API key
- **API key parameter name**: x-api-key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: x-api-key and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@higallop.com

