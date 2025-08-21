# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCRDCloudAccounts**](CloudAppsApi.md#GetCRDCloudAccounts) | **Get** /V4/CRD/CloudAccounts | Get cloud accounts for Cloud Resource Discovery operation
[**TestCloudConnection**](CloudAppsApi.md#TestCloudConnection) | **Post** /CloudApps/TestCloudConnection | Verify Cloud Resource Connectivity

# **GetCRDCloudAccounts**
> GetCrdCloudAccountsResponse GetCRDCloudAccounts(ctx, vendor, optional)
Get cloud accounts for Cloud Resource Discovery operation

This endpoint is used to return the list of cloud accounts for Cloud Resource Discovery operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendor** | **int32**| Type of the cloud vendor. Cloud accounts for the specified cloud vendor will be returned. Supported Vendor type: 3-Azure  | 
 **optional** | ***CloudAppsApiGetCRDCloudAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudAppsApiGetCRDCloudAccountsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetApp** | **optional.Int32**| App type for which the Cloud Resorce Discovery instance is being configured. Supported App type: 134-Cloud Apps  | 
 **targetInstance** | **optional.Int32**| Cloud apps instance type for which the Cloud Resorce Discovery instance is being configured. Supported Cloud apps instance type: 6-Azure Blob  | 

### Return type

[**GetCrdCloudAccountsResponse**](GetCRDCloudAccountsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestCloudConnection**
> GenericResp TestCloudConnection(ctx, optional)
Verify Cloud Resource Connectivity

Verifies cloud resource connectivity from a list of backup gateways to given hostname and port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CloudAppsApiTestCloudConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudAppsApiTestCloudConnectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TestCloudConnectionReq**](TestCloudConnectionReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

