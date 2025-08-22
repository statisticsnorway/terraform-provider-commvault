# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableBackupInstance**](InstancesApi.md#DisableBackupInstance) | **Put** /V4/Instance/{instanceId}/Backup/Action/Disable | Disable Backup for a Database Instance
[**DisableRestoreInstance**](InstancesApi.md#DisableRestoreInstance) | **Put** /V4/Instance/{instanceId}/Restore/Action/Disable | Disable Restore for an Instance
[**EnableBackupInstance**](InstancesApi.md#EnableBackupInstance) | **Put** /V4/Instance/{instanceId}/Backup/Action/Enable | Enable Backups for an Instance
[**EnableRestoreInstance**](InstancesApi.md#EnableRestoreInstance) | **Put** /V4/Instance/{instanceId}/Restore/Action/Enable | Enable restore for an Instance
[**ExcludeSLAInstance**](InstancesApi.md#ExcludeSLAInstance) | **Put** /V4/Instance/{instanceId}/SLA/Exclude | Exclude Instance from the SLA
[**GetInstances**](InstancesApi.md#GetInstances) | **Get** /V4/Instances | Get DB Instances
[**IncludeSLAInstance**](InstancesApi.md#IncludeSLAInstance) | **Put** /V4/Instance/{instanceId}/SLA/Include | Include Instance in SLA

# **DisableBackupInstance**
> GenericResp DisableBackupInstance(ctx, instanceId, optional)
Disable Backup for a Database Instance

Used to disable backup property for an instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instanceId** | **int32**| Id of the instance to modify | 
 **optional** | ***InstancesApiDisableBackupInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstancesApiDisableBackupInstanceOpts struct
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

# **DisableRestoreInstance**
> GenericResp DisableRestoreInstance(ctx, instanceId, optional)
Disable Restore for an Instance

Used to disable restore property for an instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instanceId** | **int32**| Id of the instance to modify | 
 **optional** | ***InstancesApiDisableRestoreInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstancesApiDisableRestoreInstanceOpts struct
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

# **EnableBackupInstance**
> GenericResp EnableBackupInstance(ctx, instanceId)
Enable Backups for an Instance

Used to enable backup property for an instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instanceId** | **int32**| Id of the instance to modify | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableRestoreInstance**
> GenericResp EnableRestoreInstance(ctx, instanceId)
Enable restore for an Instance

Used to enable restore property for an instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instanceId** | **int32**| Id of the instance to modify | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExcludeSLAInstance**
> GenericResp ExcludeSLAInstance(ctx, instanceId)
Exclude Instance from the SLA

Used to exclude instance from SLA. Applicable for Salesforce & Office365

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instanceId** | **int32**| Id of the instance to modify | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInstances**
> InstanceListResp GetInstances(ctx, )
Get DB Instances

This endpoint is used to return the list of instances.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InstanceListResp**](InstanceListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IncludeSLAInstance**
> GenericResp IncludeSLAInstance(ctx, instanceId)
Include Instance in SLA

Used to include instance in SLA. Applicable for Salesforce

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instanceId** | **int32**| Id of the instance to modify | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

