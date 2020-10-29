# \MineApi

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Mine**](MineApi.md#Mine) | **Post** /mine | Mine regtest coins to a specified address



## Mine

> []string Mine(ctx, optional)

Mine regtest coins to a specified address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MineOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MineOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mineRequest** | [**optional.Interface of MineRequest**](MineRequest.md)|  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

