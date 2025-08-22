# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSoftwareCache**](SoftwareCacheApi.md#AddSoftwareCache) | **Post** /V4/SoftwareCache | Create a new software cache
[**DeleteSoftwareCacheServer**](SoftwareCacheApi.md#DeleteSoftwareCacheServer) | **Delete** /V4/SoftwareCache/{clientId} | Delete software cache for server
[**GetQualifiedServersForSoftwareCache**](SoftwareCacheApi.md#GetQualifiedServersForSoftwareCache) | **Get** /V4/SoftwareCache/QualifiedServers | Get Qualified servers for software cache creation
[**GetSoftwareCacheDetailsForServer**](SoftwareCacheApi.md#GetSoftwareCacheDetailsForServer) | **Get** /V4/SoftwareCache/{clientId} | Get software cache details for specific server
[**GetSoftwareCachesDetails**](SoftwareCacheApi.md#GetSoftwareCachesDetails) | **Get** /V4/SoftwareCache | Get all software caches details
[**ModifySoftwareCacheDetails**](SoftwareCacheApi.md#ModifySoftwareCacheDetails) | **Put** /V4/SoftwareCache/{clientId} | Modify software cache details

# **AddSoftwareCache**
> GenericResp AddSoftwareCache(ctx, optional)
Create a new software cache

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoftwareCacheApiAddSoftwareCacheOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwareCacheApiAddSoftwareCacheOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SoftwareCacheDetail**](SoftwareCacheDetail.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSoftwareCacheServer**
> GenericResp DeleteSoftwareCacheServer(ctx, clientId)
Delete software cache for server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **int32**| Software cache client id | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQualifiedServersForSoftwareCache**
> SoftwareCacheDetailListResp GetQualifiedServersForSoftwareCache(ctx, )
Get Qualified servers for software cache creation

Get qualified server list to create a new software cache

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SoftwareCacheDetailListResp**](SoftwareCacheDetailListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSoftwareCacheDetailsForServer**
> SoftwareCacheDetail GetSoftwareCacheDetailsForServer(ctx, clientId)
Get software cache details for specific server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **int32**| Software cache client id | 

### Return type

[**SoftwareCacheDetail**](SoftwareCacheDetail.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSoftwareCachesDetails**
> SoftwareCacheDetailListResp GetSoftwareCachesDetails(ctx, )
Get all software caches details

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SoftwareCacheDetailListResp**](SoftwareCacheDetailListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifySoftwareCacheDetails**
> GenericResp ModifySoftwareCacheDetails(ctx, clientId, optional)
Modify software cache details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **int32**| Software cache client id | 
 **optional** | ***SoftwareCacheApiModifySoftwareCacheDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwareCacheApiModifySoftwareCacheDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ModifySoftwareCacheDetails**](ModifySoftwareCacheDetails.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

