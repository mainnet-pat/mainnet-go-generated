# \WalletApi

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Balance**](WalletApi.md#Balance) | **Post** /wallet/balance | Get total balance for wallet
[**CreateWallet**](WalletApi.md#CreateWallet) | **Post** /wallet/create | create a new wallet
[**DepositAddress**](WalletApi.md#DepositAddress) | **Post** /wallet/deposit_address | Get a deposit address in cash address format
[**DepositQr**](WalletApi.md#DepositQr) | **Post** /wallet/deposit_qr | Get receiving cash address as a qrcode
[**MaxAmountToSend**](WalletApi.md#MaxAmountToSend) | **Post** /wallet/max_amount_to_send | Get maximum spendable amount
[**Send**](WalletApi.md#Send) | **Post** /wallet/send | Send some amount to a given address
[**SendMax**](WalletApi.md#SendMax) | **Post** /wallet/send_max | Send all available funds to a given address
[**Utxos**](WalletApi.md#Utxos) | **Post** /wallet/utxo | Get detailed information about unspent outputs (utxos)



## Balance

> BalanceResponse Balance(ctx, balanceRequest)

Get total balance for wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**balanceRequest** | [**BalanceRequest**](BalanceRequest.md)| Request for a wallet balance  | 

### Return type

[**BalanceResponse**](BalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWallet

> WalletResponse CreateWallet(ctx, walletRequest)

create a new wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletRequest** | [**WalletRequest**](WalletRequest.md)| Request a new new random wallet | 

### Return type

[**WalletResponse**](WalletResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DepositAddress

> DepositAddressResponse DepositAddress(ctx, serializedWallet)

Get a deposit address in cash address format

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serializedWallet** | [**SerializedWallet**](SerializedWallet.md)| Request for a deposit address given a wallet  | 

### Return type

[**DepositAddressResponse**](DepositAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DepositQr

> ScalableVectorGraphic DepositQr(ctx, serializedWallet)

Get receiving cash address as a qrcode

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serializedWallet** | [**SerializedWallet**](SerializedWallet.md)| Request for a deposit cash address as a Quick Response code (qrcode)  | 

### Return type

[**ScalableVectorGraphic**](ScalableVectorGraphic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaxAmountToSend

> BalanceResponse MaxAmountToSend(ctx, maxAmountToSendRequest)

Get maximum spendable amount

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maxAmountToSendRequest** | [**MaxAmountToSendRequest**](MaxAmountToSendRequest.md)| get amount that will be spend with a spend max request. If a unit type is specified, a numeric value will be returned. | 

### Return type

[**BalanceResponse**](BalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Send

> SendResponse Send(ctx, sendRequest)

Send some amount to a given address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sendRequest** | [**SendRequest**](SendRequest.md)| place a send request | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMax

> SendMaxResponse SendMax(ctx, sendMaxRequest)

Send all available funds to a given address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sendMaxRequest** | [**SendMaxRequest**](SendMaxRequest.md)| Request to all available funds to a given address | 

### Return type

[**SendMaxResponse**](SendMaxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Utxos

> UtxoResponse Utxos(ctx, serializedWallet)

Get detailed information about unspent outputs (utxos)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serializedWallet** | [**SerializedWallet**](SerializedWallet.md)| Request detailed list of unspent transaction outputs  | 

### Return type

[**UtxoResponse**](UtxoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

