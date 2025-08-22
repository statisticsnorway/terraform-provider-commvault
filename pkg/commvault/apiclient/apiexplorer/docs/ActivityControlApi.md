# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableActivityControlForEntity**](ActivityControlApi.md#DisableActivityControlForEntity) | **Put** /V4/Entity/{entityType}/{entityId}/ActivityControl/{activityType}/Action/Disable | Disable Activity Control For Entity
[**EnableActivityControlForEntity**](ActivityControlApi.md#EnableActivityControlForEntity) | **Put** /V4/Entity/{entityType}/{entityId}/ActivityControl/{activityType}/Action/Enable | Enable Activity Control For Entity

# **DisableActivityControlForEntity**
> GenericResp DisableActivityControlForEntity(ctx, entityType, entityId, activityType, optional)
Disable Activity Control For Entity

API to disable activity control for an entity. Available activityType are 1 - Backup, 2 - Restore, 16 - Archieve pruning/ data aging 128 - All Job Activity, 256 - Scheduler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityType** | **int32**| Entity type to enable activity control like subclient, instance, server, server group. Available entityType are 3 - server 4 - Application type 5 - instance 7 - subclient 28 - server group | 
  **entityId** | **int32**| Entity Id to disable activity control like subclient, instance, server, server group. | 
  **activityType** | **int32**| denotes the activity type being considered. | 
 **optional** | ***ActivityControlApiDisableActivityControlForEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActivityControlApiDisableActivityControlForEntityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **optional.Int32**| clientId is mandatory when application entity type is passed. | 
 **enableAfterDelayTime** | **optional.Int32**| epoch timestamp on which activity control should be automatically enabled back | 
 **enableAfterDelayTimeZone** | **optional.Int32**| timezone ID | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableActivityControlForEntity**
> GenericResp EnableActivityControlForEntity(ctx, entityType, entityId, activityType, optional)
Enable Activity Control For Entity

API to enable activity control for an entity Available activityType are 1 - Backup, 2 - Restore, 16 - Archieve pruning/ data aging 128 - All Job Activity, 256 - Scheduler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityType** | **int32**| Entity type to enable activity control like subclient, instance, server, server group. Available entityType are 3 - server 4 - Application type 5 - instance 7 - subclient 28 - server group | 
  **entityId** | **int32**| Entity Id to enable activity control like subclient, instance, server, server group. | 
  **activityType** | **int32**| denotes the activity type being considered | 
 **optional** | ***ActivityControlApiEnableActivityControlForEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActivityControlApiEnableActivityControlForEntityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **optional.Int32**| clientId is mandatory when application entity type is passed | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

