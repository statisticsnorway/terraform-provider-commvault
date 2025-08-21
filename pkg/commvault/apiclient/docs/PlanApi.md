# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackupDestinationWithoutPlanInfo**](PlanApi.md#CreateBackupDestinationWithoutPlanInfo) | **Post** /V4/Plan/BackupDestinations | Create backup destination (without Plan)
[**DeleteBackupDestinationWithoutPlanInfo**](PlanApi.md#DeleteBackupDestinationWithoutPlanInfo) | **Delete** /V4/Plan/BackupDestination/{BackupDestinationId} | Delete backup destination
[**GetAssocMAsForPlan**](PlanApi.md#GetAssocMAsForPlan) | **Get** /V4/Plan/{planId}/MediaAgents | Get associated media agents detail for a Plan
[**GetBackupDestinationClientAssociations**](PlanApi.md#GetBackupDestinationClientAssociations) | **Get** /V4/Plan/{planId}/BackupDestination/{backupDestinationId}/Associations | Get all the associated entities of backupdestination.
[**GetBackupDestinationDetailsWithoutPlanInfo**](PlanApi.md#GetBackupDestinationDetailsWithoutPlanInfo) | **Get** /V4/Plan/BackupDestination/{BackupDestinationId} | Get details of a backup destination
[**GetClientsForBackupDestination**](PlanApi.md#GetClientsForBackupDestination) | **Get** /V4/Plan/BackupDestination/{backupDestinationId}/Clients | Get clients associated with BackupDestination
[**GetJobsOnBackupDestination**](PlanApi.md#GetJobsOnBackupDestination) | **Get** /V4/Plan/BackupDestination/{backupDestinationId}/Jobs | Get Jobs for a Backup Destination
[**GetPlanIds**](PlanApi.md#GetPlanIds) | **Get** /V4/PlanIds | Get List of Plans
[**GetPlanSummary**](PlanApi.md#GetPlanSummary) | **Get** /V4/Plan/Summary | Get Summary of a Plan
[**GetSinglePlanSummary**](PlanApi.md#GetSinglePlanSummary) | **Get** /v4/Plan/Summary/{planId} | Get Single Plan Summary
[**ModifyBackupDestinationWithoutPlanInfo**](PlanApi.md#ModifyBackupDestinationWithoutPlanInfo) | **Put** /V4/Plan/BackupDestination/{BackupDestinationId} | Modify Backup Destination (Without a Plan)
[**PickUnpickJobsForCatalog**](PlanApi.md#PickUnpickJobsForCatalog) | **Put** /V4/Plan/{planId}/BackupDestination/{backupDestinationId}/SnapCatalog/Jobs | pick or unpick jobs for snapshot catalog
[**RunAuxCopyForPlan**](PlanApi.md#RunAuxCopyForPlan) | **Post** /V4/Plan/{planId}/BackupDestination/{backupDestinationId}/AuxCopy/Run | Run aux copy for backup destination(s) of a plan
[**RunBackupCopy**](PlanApi.md#RunBackupCopy) | **Post** /V4/Plan/{planId}/BackupDestination/{backupDestionationId}/BackupCopy/Run | RunBackupCopyForPlan
[**RunJobOperationOnPlanBackupDestination**](PlanApi.md#RunJobOperationOnPlanBackupDestination) | **Post** /V4/Plan/BackupDestination/JobOperations | Job Operations
[**RunSnapCatalog**](PlanApi.md#RunSnapCatalog) | **Post** /V4/Plan/{planId}/BackupDestination/{backupDestionationId}/SnapCatalog/Run | RunSnapCatalogForPlan
[**UpdateBackupDestinationAssociations**](PlanApi.md#UpdateBackupDestinationAssociations) | **Put** /V4/Plan/{planId}/BackupDestination/{backupDestinationId}/Associations | Update all the included/excluded entities of backupdestination.

# **CreateBackupDestinationWithoutPlanInfo**
> CreatePlanBackupDestinationResp CreateBackupDestinationWithoutPlanInfo(ctx, optional)
Create backup destination (without Plan)

Create backup destination(s) for a plan before creating the plan. The primary backup destination can be then associated with the plan to associate all secondary copies also. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlanApiCreateBackupDestinationWithoutPlanInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanApiCreateBackupDestinationWithoutPlanInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateBackupDestinations**](CreateBackupDestinations.md)|  | 

### Return type

[**CreatePlanBackupDestinationResp**](CreatePlanBackupDestinationResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBackupDestinationWithoutPlanInfo**
> GenericResp DeleteBackupDestinationWithoutPlanInfo(ctx, backupDestinationId)
Delete backup destination

Delete Backup Destination. If trying to delete primary backup destination which is not associated with any plan, it will delete all secondary copies also.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupDestinationId** | **int32**| Id of the backupDestination to be deleted | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAssocMAsForPlan**
> AssociatedMasForPlanResp GetAssocMAsForPlan(ctx, planId, optional)
Get associated media agents detail for a Plan

Get list of associated media agents detail for a Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Plan id of the plan | 
 **optional** | ***PlanApiGetAssocMAsForPlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanApiGetAssocMAsForPlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **packageName** | [**optional.Interface of PackageName**](.md)| Provide the name of the package that needs to be checked if installed on the Media Agents. isPackageInstalled will be populated accordingly. By default it checks for Media agent package. Refer to packageName list. | 

### Return type

[**AssociatedMasForPlanResp**](AssociatedMAsForPlanResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupDestinationClientAssociations**
> BackupDestinationAssociationsResp GetBackupDestinationClientAssociations(ctx, planId, backupDestinationId, optional)
Get all the associated entities of backupdestination.

Get all the associated server/server groups and included/excluded entities of backupdestination. It will also return summary of all the entities that will be auxcopied for backupdestination. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**|  | 
  **backupDestinationId** | **int32**|  | 
 **optional** | ***PlanApiGetBackupDestinationClientAssociationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanApiGetBackupDestinationClientAssociationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **entityType** | **optional.Int32**| It is to further filter and get all the child entities of given entityType/entityId associated with backupdestination. | 
 **entityId** | **optional.Int32**| It is to further filter and get all the child entities of given entityType/entityId associated with backupdestination. | 

### Return type

[**BackupDestinationAssociationsResp**](BackupDestinationAssociationsResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupDestinationDetailsWithoutPlanInfo**
> BackupDestination GetBackupDestinationDetailsWithoutPlanInfo(ctx, backupDestinationId)
Get details of a backup destination

Fetch details of a backup destination. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupDestinationId** | **int32**| Id of the backupDestination for which details will be fetched. | 

### Return type

[**BackupDestination**](BackupDestination.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientsForBackupDestination**
> ClientListForStorageResp GetClientsForBackupDestination(ctx, backupDestinationId)
Get clients associated with BackupDestination

Get the list of clients associated with the BackupDestination.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupDestinationId** | **int32**| Id of the BackupDestination. | 

### Return type

[**ClientListForStorageResp**](ClientListForStorageResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobsOnBackupDestination**
> GetJobsOnStorageResp GetJobsOnBackupDestination(ctx, backupDestinationId, optional)
Get Jobs for a Backup Destination

Get the list of Jobs for selected backupDestinationId.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupDestinationId** | **int32**| Id of the BackupDestination to fetch Job List. | 
 **optional** | ***PlanApiGetJobsOnBackupDestinationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanApiGetJobsOnBackupDestinationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **optional.String**| Time period selection for which to fetch jobs. Accepted values [last24Hours, lastWeek, lastMonth, last3Months, custom]. When custom is selected Jobs are filtered based on values provided in other params. | [default to last24Hours]
 **clients** | **optional.String**| Comma separated Client Ids to filter the jobs based on clients associated | 
 **agedData** | **optional.Int32**| Aged Jobs selection. Accepted values [ 0 to exclude aged jobs, 1 to show only aged jobs, 2 to include aged jobs]. | 
 **backupLvl** | **optional.Int32**| Job backup type filter. Accepted values [1&#x3D;Full, 2&#x3D;Incremental, 4&#x3D;Differential, 8&#x3D;All, 64&#x3D;Synthetic full]. | 
 **copyState** | **optional.Int32**| Filter jobs by data status. Accepted values [0  &#x3D; show all, 1  &#x3D; show available, 4  &#x3D; show to be copied, 8  &#x3D; show not to be copied, 16 &#x3D; show extended retained]. | 
 **startTime** | **optional.Int64**| Start time of the time range. | 
 **endTime** | **optional.Int64**| End time of the time range. | 
 **minAppSizeMB** | **optional.Int64**| Minimum size of job in Megabytes for application size range filter | 
 **maxAppSizeMB** | **optional.Int64**| Maximum size of job in Megabytes for application size range filter | 

### Return type

[**GetJobsOnStorageResp**](GetJobsOnStorageResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlanIds**
> EntityDetails GetPlanIds(ctx, )
Get List of Plans

Get All Plans as Name Id Pairs

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EntityDetails**](EntityDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlanSummary**
> PlanSummaryListResp GetPlanSummary(ctx, )
Get Summary of a Plan

Get Summary of a Plan

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PlanSummaryListResp**](PlanSummaryListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSinglePlanSummary**
> SinglePlanSummary GetSinglePlanSummary(ctx, planId)
Get Single Plan Summary

Api to fetch summary of a plan.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the plan to get summary of | 

### Return type

[**SinglePlanSummary**](SinglePlanSummary.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyBackupDestinationWithoutPlanInfo**
> PlanBackupDestinationResp ModifyBackupDestinationWithoutPlanInfo(ctx, backupDestinationId, optional)
Modify Backup Destination (Without a Plan)

Modify Backup Destination. Only primary backup destination will be considered for region updation. If the primary backup destination is not associated with a plan, only then the region associaion can be modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupDestinationId** | **int32**| Id of the backupDestination to be modified | 
 **optional** | ***PlanApiModifyBackupDestinationWithoutPlanInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanApiModifyBackupDestinationWithoutPlanInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateBackupDestination**](UpdateBackupDestination.md)|  | 

### Return type

[**PlanBackupDestinationResp**](PlanBackupDestinationResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PickUnpickJobsForCatalog**
> GenericResp PickUnpickJobsForCatalog(ctx, planId, backupDestinationId, optional)
pick or unpick jobs for snapshot catalog

Send job details to pick or unpick jobs for snapshot catalog

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| plan Id | 
  **backupDestinationId** | **int32**| Backup Destination Id | 
 **optional** | ***PlanApiPickUnpickJobsForCatalogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanApiPickUnpickJobsForCatalogOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of JobsForSnapCatalog**](JobsForSnapCatalog.md)| Job details for unpicking jobs for snapshot catalog. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunAuxCopyForPlan**
> CreateTaskRespforBackup RunAuxCopyForPlan(ctx, planId, backupDestinationId, optional)
Run aux copy for backup destination(s) of a plan

Run aux copy job for backup destination(s) of a plan.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan | 
  **backupDestinationId** | **int32**| Id of Backup Destination. | 
 **optional** | ***PlanApiRunAuxCopyForPlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanApiRunAuxCopyForPlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeAllCopies** | **optional.Bool**| Primary backupDestinationId can be passed in backupDestinationId Path param when includeAllCopies is set to true, this will trigger the aux copy on all supported backup destinations of a region | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunBackupCopy**
> CreateTaskRespforBackup RunBackupCopy(ctx, planId, backupDestionationId)
RunBackupCopyForPlan

Run Backup copy job for a backupdestination

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**|  | 
  **backupDestionationId** | **int32**|  | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunJobOperationOnPlanBackupDestination**
> JobOperationOnCopyResp RunJobOperationOnPlanBackupDestination(ctx, optional)
Job Operations

Run different job operations for a plan backup destination

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlanApiRunJobOperationOnPlanBackupDestinationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanApiRunJobOperationOnPlanBackupDestinationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of JobOperationOnCopyReq**](JobOperationOnCopyReq.md)|  | 

### Return type

[**JobOperationOnCopyResp**](JobOperationOnCopyResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunSnapCatalog**
> CreateTaskRespforBackup RunSnapCatalog(ctx, planId, backupDestionationId)
RunSnapCatalogForPlan

Run deferred snap catalog job for a backupdestination

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**|  | 
  **backupDestionationId** | **int32**|  | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBackupDestinationAssociations**
> GenericResp UpdateBackupDestinationAssociations(ctx, planId, backupDestinationId, optional)
Update all the included/excluded entities of backupdestination.

Update all the included/excluded entities of backupdestination.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**|  | 
  **backupDestinationId** | **int32**|  | 
 **optional** | ***PlanApiUpdateBackupDestinationAssociationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanApiUpdateBackupDestinationAssociationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpdateBackupDestinationAssociationsReq**](UpdateBackupDestinationAssociationsReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

