# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKeyManagementServers**](KeyManagementServerApi.md#DeleteKeyManagementServers) | **Delete** /V4/KeyManagementServers/{kmsId} | Delete Key Management Server
[**GetKeyManagementServers**](KeyManagementServerApi.md#GetKeyManagementServers) | **Get** /V4/KeyManagementServers | Get Key Management Servers

# **DeleteKeyManagementServers**
> GenericResp DeleteKeyManagementServers(ctx, kmsId)
Delete Key Management Server

Delete key management server based on Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **kmsId** | **int32**| Id of Key Management Server | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKeyManagementServers**
> InlineResponse20020 GetKeyManagementServers(ctx, )
Get Key Management Servers

Get key management servers

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

