# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHybridFileStoreShareStatus**](HybridFileStoresApi.md#GetHybridFileStoreShareStatus) | **Get** /v4/HFSShare/{HFSShareId}/Status | Get Hybrid File Store Share Status
[**GetHybridFileStores**](HybridFileStoresApi.md#GetHybridFileStores) | **Get** /V4/HybridFileStores | Get Hybrid File Stores

# **GetHybridFileStoreShareStatus**
> HfsShareStatusResp GetHybridFileStoreShareStatus(ctx, hFSShareId)
Get Hybrid File Store Share Status

API to get Hybrid File Store Share Status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hFSShareId** | **int32**| Id of the HFS Share to fetch its status | 

### Return type

[**HfsShareStatusResp**](HFSShareStatusResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHybridFileStores**
> HybridFileStoresListResp GetHybridFileStores(ctx, )
Get Hybrid File Stores

Get list of hybrid file stores

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HybridFileStoresListResp**](HybridFileStoresListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

