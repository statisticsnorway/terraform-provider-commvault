# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGCPSpannerDatabases**](GCPSpannerApi.md#GetGCPSpannerDatabases) | **Get** /V4/GCPSpanner/databases | Get database list for GCP
[**GetGCPSpannerInstanceList**](GCPSpannerApi.md#GetGCPSpannerInstanceList) | **Get** /V4/GCPSpanner/instances | To get the list of instances for GCP
[**GetGCPSpannerPermissions**](GCPSpannerApi.md#GetGCPSpannerPermissions) | **Get** /V4/GCPSpanner/permissions | Check permission on a project for a cloud account

# **GetGCPSpannerDatabases**
> GcpDatabaseList GetGCPSpannerDatabases(ctx, cloudAccountId, projectName, instanceName)
Get database list for GCP

Get database list for GCP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudAccountId** | **int32**| the id of the node where the browse request is sent | 
  **projectName** | **string**| google cloud project the instance belongs to | 
  **instanceName** | **string**| the google spanner instance name | 

### Return type

[**GcpDatabaseList**](GCPDatabaseList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGCPSpannerInstanceList**
> GcpInstanceList GetGCPSpannerInstanceList(ctx, cloudAccountId)
To get the list of instances for GCP

To get the list of instances for GCP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudAccountId** | **int32**| the google cloud hypervisor account id | 

### Return type

[**GcpInstanceList**](GCPInstanceList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGCPSpannerPermissions**
> StatusModel GetGCPSpannerPermissions(ctx, optional)
Check permission on a project for a cloud account

Get permission status for a cloudAccount on a GCP project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GCPSpannerApiGetGCPSpannerPermissionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GCPSpannerApiGetGCPSpannerPermissionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudAccountId** | **optional.Int32**| the id of the node where the browse request is sent | 
 **projectName** | **optional.String**| google cloud project the instance belongs to | 

### Return type

[**StatusModel**](StatusModel.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

