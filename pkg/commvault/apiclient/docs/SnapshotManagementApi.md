# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSnapClients**](SnapshotManagementApi.md#GetSnapClients) | **Get** /V4/StorageArrays/Client/Snap/Enabled | List Intellisnap Enabled Clients
[**GetSnapSubclients**](SnapshotManagementApi.md#GetSnapSubclients) | **Get** /V4/StorageArrays/Client/{clientId}/subclient/Snap/Enabled | List Intellisnap Enabled Subclients
[**GetSubclientsForSnapEngine**](SnapshotManagementApi.md#GetSubclientsForSnapEngine) | **Get** /V4/Snaps/Subclients | Get Subclients for Snap Engine

# **GetSnapClients**
> []IntelliSnapClientsList GetSnapClients(ctx, )
List Intellisnap Enabled Clients

API to get all IntelliSnap enabled Clients

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]IntelliSnapClientsList**](IntelliSnapClientsList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSnapSubclients**
> []IntelliSnapSubclients GetSnapSubclients(ctx, clientId)
List Intellisnap Enabled Subclients

API to get all IntelliSnap enabled subclients for particular client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **int32**|  | 

### Return type

[**[]IntelliSnapSubclients**](IntelliSnapSubclients.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubclientsForSnapEngine**
> GetSubclientForSnapEngine GetSubclientsForSnapEngine(ctx, )
Get Subclients for Snap Engine

Get subclients configured with a particular snap engine

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetSubclientForSnapEngine**](GetSubclientForSnapEngine.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

