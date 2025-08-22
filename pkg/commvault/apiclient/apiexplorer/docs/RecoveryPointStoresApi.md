# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRecoveryPointStores**](RecoveryPointStoresApi.md#DeleteRecoveryPointStores) | **Delete** /V4/RecoveryPointStore/{rpsId} | Delete Recovery Point Stores
[**GetRecoveryPointStoreList**](RecoveryPointStoresApi.md#GetRecoveryPointStoreList) | **Get** /V4/RecoveryPointStores | Get RecoveryPoint StoreList
[**GetRecoveryPointStores**](RecoveryPointStoresApi.md#GetRecoveryPointStores) | **Get** /V4/RecoveryPointStore/{rpsId} | Get Recovery Point Stores

# **DeleteRecoveryPointStores**
> GenericResp DeleteRecoveryPointStores(ctx, rpsId)
Delete Recovery Point Stores

Delete a specific recovery point store

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rpsId** | **string**|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRecoveryPointStoreList**
> RpStoreListResponse GetRecoveryPointStoreList(ctx, )
Get RecoveryPoint StoreList

Get the list of recovery point stores.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RpStoreListResponse**](RPStoreListResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRecoveryPointStores**
> RecoveryPointStoreDetails GetRecoveryPointStores(ctx, rpsId)
Get Recovery Point Stores

Fetch details about recovery point stores or library details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rpsId** | **string**|  | 

### Return type

[**RecoveryPointStoreDetails**](RecoveryPointStoreDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

