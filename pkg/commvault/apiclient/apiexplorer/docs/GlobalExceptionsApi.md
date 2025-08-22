# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGlobalExceptions**](GlobalExceptionsApi.md#GetGlobalExceptions) | **Get** /V4/GlobalExceptions | Get GlobalExceptions
[**UpdateGlobalExceptions**](GlobalExceptionsApi.md#UpdateGlobalExceptions) | **Post** /V4/GlobalExceptions | Update GlobalExceptions

# **GetGlobalExceptions**
> GetGlobalExceptions GetGlobalExceptions(ctx, )
Get GlobalExceptions

Get list of global execptions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetGlobalExceptions**](GetGlobalExceptions.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGlobalExceptions**
> GenericResp UpdateGlobalExceptions(ctx, optional)
Update GlobalExceptions

Add/Edit/Delete global execptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GlobalExceptionsApiUpdateGlobalExceptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalExceptionsApiUpdateGlobalExceptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SetGlobalExceptions**](SetGlobalExceptions.md)| Request body to be passed to the POST API | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

