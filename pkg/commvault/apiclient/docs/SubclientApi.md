# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableBackupSubclient**](SubclientApi.md#DisableBackupSubclient) | **Put** /V4/Subclient/{subclientId}/Backup/Action/Disable | 
[**EnableBackupSubclient**](SubclientApi.md#EnableBackupSubclient) | **Put** /V4/Subclient/{subclientId}/Backup/Action/Enable | 

# **DisableBackupSubclient**
> GenericResp DisableBackupSubclient(ctx, subclientId, optional)


Used to disable backup property for a subclient

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subclientId** | **int32**| Id of the subclient to modify | 
 **optional** | ***SubclientApiDisableBackupSubclientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubclientApiDisableBackupSubclientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableAfterADelay** | **optional.Int32**| Provide UTC time in unix format. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableBackupSubclient**
> GenericResp EnableBackupSubclient(ctx, subclientId)


Used to enable backup property for an subclient

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subclientId** | **int32**| Id of the subclient to modify | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

