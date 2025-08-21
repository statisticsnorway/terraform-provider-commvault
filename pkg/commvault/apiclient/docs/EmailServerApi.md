# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigureEmailServer**](EmailServerApi.md#ConfigureEmailServer) | **Post** /V4/EmailServer | Configure Email Server
[**GetEmailServer**](EmailServerApi.md#GetEmailServer) | **Get** /V4/EmailServer | Get Email Server
[**SendTestMailReq**](EmailServerApi.md#SendTestMailReq) | **Post** /V4/EmailServer/Action/Test | Send Test Mail Req
[**UpdateEmailServer**](EmailServerApi.md#UpdateEmailServer) | **Put** /V4/EmailServer | Update Email Server

# **ConfigureEmailServer**
> GenericResp ConfigureEmailServer(ctx, optional)
Configure Email Server

Configure SMTP server settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmailServerApiConfigureEmailServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EmailServerApiConfigureEmailServerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ConfigureSmtpServerReq**](ConfigureSmtpServerReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailServer**
> GetEmailServerDetails GetEmailServer(ctx, )
Get Email Server

Retrieves SMTP server details

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetEmailServerDetails**](GetEmailServerDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTestMailReq**
> GenericResp SendTestMailReq(ctx, optional)
Send Test Mail Req

To test Email settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmailServerApiSendTestMailReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EmailServerApiSendTestMailReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SendTestMailReq**](SendTestMailReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEmailServer**
> GenericResp UpdateEmailServer(ctx, optional)
Update Email Server

Update the SMTP server settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmailServerApiUpdateEmailServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EmailServerApiUpdateEmailServerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UpdateEmailServerReq**](UpdateEmailServerReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

