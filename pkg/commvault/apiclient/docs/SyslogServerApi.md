# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigureSyslogStatus**](SyslogServerApi.md#ConfigureSyslogStatus) | **Post** /V4/syslogServer | Configure Syslog Server
[**GetSyslogStatus**](SyslogServerApi.md#GetSyslogStatus) | **Get** /V4/syslogServer | Get Syslog Server Status

# **ConfigureSyslogStatus**
> GenericResp ConfigureSyslogStatus(ctx, optional)
Configure Syslog Server

This endpoint configures a syslog server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SyslogServerApiConfigureSyslogStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SyslogServerApiConfigureSyslogStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SyslogConfigure**](SyslogConfigure.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSyslogStatus**
> SyslogStatus GetSyslogStatus(ctx, )
Get Syslog Server Status

This endpoint returns the details of a syslog server.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SyslogStatus**](SyslogStatus.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

