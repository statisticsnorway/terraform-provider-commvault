# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArchiveBackupDestination**](ArchivePlanApi.md#CreateArchiveBackupDestination) | **Post** /V4/ArchivePlan/{planId}/BackupDestination | 
[**CreateArchivePlan**](ArchivePlanApi.md#CreateArchivePlan) | **Post** /V4/ArchivePlan | Create Archive Plan
[**DeleteArchiveBackupDestination**](ArchivePlanApi.md#DeleteArchiveBackupDestination) | **Delete** /V4/ArchivePlan/{planId}/BackupDestination/{BackupDestinationId} | 
[**DeleteArchivePlanById**](ArchivePlanApi.md#DeleteArchivePlanById) | **Delete** /V4/ArchivePlan/{planId} | Delete existing archive plan
[**DeleteStorageRegionForArchivePlan**](ArchivePlanApi.md#DeleteStorageRegionForArchivePlan) | **Delete** /V4/ArchivePlan/{planId}/storageRegion/{regionList} | 
[**GetArchiveBackupDestinations**](ArchivePlanApi.md#GetArchiveBackupDestinations) | **Get** /V4/ArchivePlan/{planId}/BackupDestinations | 
[**GetArchivePlanById**](ArchivePlanApi.md#GetArchivePlanById) | **Get** /V4/ArchivePlan/{planId} | Get Archive Plan details
[**ModifyArchiveBackupDestination**](ArchivePlanApi.md#ModifyArchiveBackupDestination) | **Put** /V4/ArchivePlan/{planId}/BackupDestination/{BackupDestinationId} | 
[**ModifyArchivePlanById**](ArchivePlanApi.md#ModifyArchivePlanById) | **Put** /V4/ArchivePlan/{planId} | Modify existing archive plan details

# **CreateArchiveBackupDestination**
> PlanBackupDestinationResp CreateArchiveBackupDestination(ctx, planId, optional)


Create a Backup Destination for a Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan to modify | 
 **optional** | ***ArchivePlanApiCreateArchiveBackupDestinationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArchivePlanApiCreateArchiveBackupDestinationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateArchivePlanBackupDestinations**](CreateArchivePlanBackupDestinations.md)|  | 

### Return type

[**PlanBackupDestinationResp**](PlanBackupDestinationResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateArchivePlan**
> PlanResp CreateArchivePlan(ctx, optional)
Create Archive Plan

Create Archive Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ArchivePlanApiCreateArchivePlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArchivePlanApiCreateArchivePlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ArchivePlan**](ArchivePlan.md)|  | 

### Return type

[**PlanResp**](PlanResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteArchiveBackupDestination**
> GenericResp DeleteArchiveBackupDestination(ctx, planId, backupDestinationId)


Delete Backup Destination for a Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan to modify | 
  **backupDestinationId** | **int32**| Id of the backupDestination to be deleted | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteArchivePlanById**
> GenericResp DeleteArchivePlanById(ctx, planId)
Delete existing archive plan

Delete existing archive plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the plan | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStorageRegionForArchivePlan**
> GenericResp DeleteStorageRegionForArchivePlan(ctx, planId, regionList, optional)


Delete Storage Region for a archive Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Archive Plan to modify | 
  **regionList** | **string**| List of region names/ids to be deleted. If region ids are passed, set isRegionIdList&#x3D;true | 
 **optional** | ***ArchivePlanApiDeleteStorageRegionForArchivePlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArchivePlanApiDeleteStorageRegionForArchivePlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isRegionIdList** | **optional.Bool**| Is regionList a list of region ids | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArchiveBackupDestinations**
> ArchivePlanBackupDestinations GetArchiveBackupDestinations(ctx, planId)


Get Backup Destinations for a Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Archive Plan to retrieve backup destinations | 

### Return type

[**ArchivePlanBackupDestinations**](ArchivePlanBackupDestinations.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArchivePlanById**
> ArchivePlanDetails GetArchivePlanById(ctx, planId)
Get Archive Plan details

Get Archive Plan details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the plan | 

### Return type

[**ArchivePlanDetails**](ArchivePlanDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyArchiveBackupDestination**
> PlanBackupDestinationResp ModifyArchiveBackupDestination(ctx, planId, backupDestinationId, optional)


Modify Backup Destination for a Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan to modify | 
  **backupDestinationId** | **int32**| Id of the backupDestination to be modified | 
 **optional** | ***ArchivePlanApiModifyArchiveBackupDestinationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArchivePlanApiModifyArchiveBackupDestinationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpdateArchivePlanBackupDestination**](UpdateArchivePlanBackupDestination.md)|  | 

### Return type

[**PlanBackupDestinationResp**](PlanBackupDestinationResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyArchivePlanById**
> GenericResp ModifyArchivePlanById(ctx, planId, optional)
Modify existing archive plan details

Modify existing archive plan details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the plan | 
 **optional** | ***ArchivePlanApiModifyArchivePlanByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArchivePlanApiModifyArchivePlanByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateArchivePlan**](UpdateArchivePlan.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

