# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAzureADClient**](ActiveDirectoryApi.md#CreateAzureADClient) | **Post** /v4/ActiveDirectory/AzureAD | Create Azure Active Directory app
[**GetActiveDirectoryClientsV2**](ActiveDirectoryApi.md#GetActiveDirectoryClientsV2) | **Get** /V4/ActiveDirectory/Apps | Get Azure and OnPrem Active Directory clients with their file system state

# **CreateAzureADClient**
> IdName CreateAzureADClient(ctx, optional)
Create Azure Active Directory app

Create Azure Active Directory app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ActiveDirectoryApiCreateAzureADClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActiveDirectoryApiCreateAzureADClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateAzureAdClient**](CreateAzureAdClient.md)|  | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActiveDirectoryClientsV2**
> ActiveDirectoryClientsV2Resp GetActiveDirectoryClientsV2(ctx, )
Get Azure and OnPrem Active Directory clients with their file system state

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ActiveDirectoryClientsV2Resp**](ActiveDirectoryClientsV2Resp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

