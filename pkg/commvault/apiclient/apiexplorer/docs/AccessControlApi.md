# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessControl**](AccessControlApi.md#GetAccessControl) | **Get** /V4/AccessControl | Get Access Control Details
[**PutAccessControl**](AccessControlApi.md#PutAccessControl) | **Put** /V4/AccessControl | Modify Access Control

# **GetAccessControl**
> AccessControl GetAccessControl(ctx, )
Get Access Control Details

Gets owner permissions and laptop ownership details

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AccessControl**](AccessControl.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAccessControl**
> GenericResp PutAccessControl(ctx, optional)
Modify Access Control

Updates owner permissions and/or automatic laptop ownership assignment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccessControlApiPutAccessControlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccessControlApiPutAccessControlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AccessControl**](AccessControl.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

