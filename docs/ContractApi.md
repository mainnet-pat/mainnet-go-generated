# \ContractApi

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEscrow**](ContractApi.md#CreateEscrow) | **Post** /contract/escrow/create | Create an escrow contract
[**EscrowFn**](ContractApi.md#EscrowFn) | **Post** /contract/escrow/call | Finalize an escrow contract
[**EscrowUtxos**](ContractApi.md#EscrowUtxos) | **Post** /contract/escrow/utxos | List specific utxos in a contract



## CreateEscrow

> ContractResponse CreateEscrow(ctx, escrowRequest)

Create an escrow contract

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**escrowRequest** | [**EscrowRequest**](EscrowRequest.md)| Request a new escrow contract | 

### Return type

[**ContractResponse**](ContractResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EscrowFn

> ContractFnResponse EscrowFn(ctx, contractFnRequest)

Finalize an escrow contract

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractFnRequest** | [**ContractFnRequest**](ContractFnRequest.md)| null | 

### Return type

[**ContractFnResponse**](ContractFnResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EscrowUtxos

> UtxoResponse EscrowUtxos(ctx, contract)

List specific utxos in a contract

Returns all UTXOs that can be spent by the  contract. Both confirmed and unconfirmed UTXOs are included. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | [**Contract**](Contract.md)|  | 

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

