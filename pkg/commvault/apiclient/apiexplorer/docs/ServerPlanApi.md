# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupServerPlan**](ServerPlanApi.md#BackupServerPlan) | **Post** /V4/ServerPlan/{planId}/Backup | Run backups on server plan associated entities
[**BackupServerPlan_0**](ServerPlanApi.md#BackupServerPlan_0) | **Post** /V4/ServerPlan/{planId}/RPO/Run | Run backups on server plan associated entities based on existing schedule
[**CreateBackupDestination**](ServerPlanApi.md#CreateBackupDestination) | **Post** /V4/ServerPlan/{planId}/BackupDestination | Create Backup Destination for Server Plan
[**CreateServerPlan**](ServerPlanApi.md#CreateServerPlan) | **Post** /V4/ServerPlan | Create Server Plan
[**CreateServerPlanBackupContent**](ServerPlanApi.md#CreateServerPlanBackupContent) | **Post** /V4/ServerPlan/{planId}/BackupContent | Create Server Plan Backup Content
[**CreateServerPlanRPO**](ServerPlanApi.md#CreateServerPlanRPO) | **Post** /V4/ServerPlan/{planId}/RPO | Create Server Plan RPO
[**DeleteBackupDestination**](ServerPlanApi.md#DeleteBackupDestination) | **Delete** /V4/ServerPlan/{planId}/BackupDestination/{BackupDestinationId} | Delete ServerPlan BackupDestination
[**DeletePlan**](ServerPlanApi.md#DeletePlan) | **Delete** /V4/ServerPlan/{planId} | Delete a Server Plan
[**DeleteStorageRegion**](ServerPlanApi.md#DeleteStorageRegion) | **Delete** /V4/ServerPlan/{planId}/storageRegion/{regionList} | Delete Server Plan Storage Region
[**DeleteStorageTagRegion**](ServerPlanApi.md#DeleteStorageTagRegion) | **Delete** /V4/Global/ServerPlan/{planId}/storageRegion/{regionId} | Delete Global Server Plan Storage Region
[**DisableBackupOnPlan**](ServerPlanApi.md#DisableBackupOnPlan) | **Put** /V4/ServerPlan/{planId}/Backup/Disable | Disable Backup Schedule policies on serverplan
[**EnableBackupOnPlan**](ServerPlanApi.md#EnableBackupOnPlan) | **Put** /V4/ServerPlan/{planId}/Backup/Enable | Enable backup Schedule policies on server plan
[**GetBackupDestinationDetails**](ServerPlanApi.md#GetBackupDestinationDetails) | **Get** /V5/ServerPlan/{planId}/BackupDestination/{BackupDestinationId} | Get Backup Destination details of ServerPlan
[**GetBackupDestinations**](ServerPlanApi.md#GetBackupDestinations) | **Get** /V4/ServerPlan/{planId}/BackupDestinations | Get Backup Destination of Server Plan
[**GetPlanById**](ServerPlanApi.md#GetPlanById) | **Get** /V4/ServerPlan/{planId} | Get Server Plan Details
[**GetServerPlanRPO**](ServerPlanApi.md#GetServerPlanRPO) | **Get** /V4/ServerPlan/{planId}/RPO | Get Server Plan RPO
[**GetSourceCopiesForGivenServerPlanRegion**](ServerPlanApi.md#GetSourceCopiesForGivenServerPlanRegion) | **Get** /V5/ServerPlan/{planId}/BackupDestination/SourceCopy/Eligible | Gives a list of eligible source copies for given region id
[**ModifyBackupDestination**](ServerPlanApi.md#ModifyBackupDestination) | **Put** /V4/ServerPlan/{planId}/BackupDestination/{BackupDestinationId} | Modify ServerPlan BackupDestination
[**ModifyBackupDestinationDetails**](ServerPlanApi.md#ModifyBackupDestinationDetails) | **Put** /V5/ServerPlan/{planId}/BackupDestination/{BackupDestinationId} | Update Backup Destination details of ServerPlan
[**ModifyPlan**](ServerPlanApi.md#ModifyPlan) | **Put** /V4/ServerPlan/{planId} | Update the server plan
[**UpdateServerPlanRPO**](ServerPlanApi.md#UpdateServerPlanRPO) | **Put** /V4/ServerPlan/{planId}/RPO | Update Server Plan RPO
[**V4ServerPlanPlanIdRetentionSynthfullPut**](ServerPlanApi.md#V4ServerPlanPlanIdRetentionSynthfullPut) | **Put** /V4/ServerPlan/{planId}/Retention/Synthfull | Update the server plan retention based synthetic full

# **BackupServerPlan**
> CreateTaskRespforBackup BackupServerPlan(ctx, planId, backupLevel)
Run backups on server plan associated entities

Run backups on server plan associated entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Server Plan | 
  **backupLevel** | **string**| Backup level of jobs | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BackupServerPlan_0**
> CreateTaskRespforBackup BackupServerPlan_0(ctx, planId, optional)
Run backups on server plan associated entities based on existing schedule

Run backups on server plan associated entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan to modify | 
 **optional** | ***ServerPlanApiBackupServerPlan_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerPlanApiBackupServerPlan_1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ServerPlanBackupRequest**](ServerPlanBackupRequest.md)|  | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBackupDestination**
> PlanBackupDestinationResp CreateBackupDestination(ctx, planId, optional)
Create Backup Destination for Server Plan

Create a Backup Destination for a Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan to modify | 
 **optional** | ***ServerPlanApiCreateBackupDestinationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerPlanApiCreateBackupDestinationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreatePlanBackupDestinations**](CreatePlanBackupDestinations.md)|  | 

### Return type

[**PlanBackupDestinationResp**](PlanBackupDestinationResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateServerPlan**
> PlanResp CreateServerPlan(ctx, optional)
Create Server Plan

Create a Server Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServerPlanApiCreateServerPlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerPlanApiCreateServerPlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateServerPlan**](CreateServerPlan.md)|  | 

### Return type

[**PlanResp**](PlanResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateServerPlanBackupContent**
> CreateServerPlanBackupContentResponse CreateServerPlanBackupContent(ctx, planId, optional)
Create Server Plan Backup Content

API to create backup content on server plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan to modify | 
 **optional** | ***ServerPlanApiCreateServerPlanBackupContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerPlanApiCreateServerPlanBackupContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateServerPlanBackupContent**](CreateServerPlanBackupContent.md)|  | 

### Return type

[**CreateServerPlanBackupContentResponse**](CreateServerPlanBackupContentResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateServerPlanRPO**
> CreateServerPlanRpoResponse CreateServerPlanRPO(ctx, planId, optional)
Create Server Plan RPO

API to create RPO schedules on server plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Server Plan | 
 **optional** | ***ServerPlanApiCreateServerPlanRPOOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerPlanApiCreateServerPlanRPOOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateServerPlanRpo**](CreateServerPlanRpo.md)|  | 

### Return type

[**CreateServerPlanRpoResponse**](CreateServerPlanRPOResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBackupDestination**
> GenericResp DeleteBackupDestination(ctx, planId, backupDestinationId)
Delete ServerPlan BackupDestination

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

# **DeletePlan**
> GenericResp DeletePlan(ctx, planId)
Delete a Server Plan

Used to delete an existing server plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the plan to delete | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStorageRegion**
> GenericResp DeleteStorageRegion(ctx, planId, regionList, optional)
Delete Server Plan Storage Region

Delete Storage Region for an Elastic Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan to modify | 
  **regionList** | **string**| List of region names/ids to be deleted. If region ids are passed, set isRegionIdList&#x3D;true | 
 **optional** | ***ServerPlanApiDeleteStorageRegionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerPlanApiDeleteStorageRegionOpts struct
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

# **DeleteStorageTagRegion**
> GenericResp DeleteStorageTagRegion(ctx, planId, regionId)
Delete Global Server Plan Storage Region

Delete Storage Region for a Global Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan to modify | 
  **regionId** | **int32**| Region id to be deleted. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableBackupOnPlan**
> GenericResp DisableBackupOnPlan(ctx, planId)
Disable Backup Schedule policies on serverplan

API to Disable backup schedule policies on server plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Server Plan | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableBackupOnPlan**
> GenericResp EnableBackupOnPlan(ctx, planId)
Enable backup Schedule policies on server plan

API to Enable backup schedule policies on server plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Server Plan | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupDestinationDetails**
> PlanBackupDestinationDetails GetBackupDestinationDetails(ctx, planId, backupDestinationId)
Get Backup Destination details of ServerPlan

Get specific backup destination details for a server plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the plan to retrieve backup destination | 
  **backupDestinationId** | **int32**| Id of the backup destination | 

### Return type

[**PlanBackupDestinationDetails**](PlanBackupDestinationDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupDestinations**
> PlanBackupDestinations GetBackupDestinations(ctx, planId)
Get Backup Destination of Server Plan

Get Backup Destinations for a Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan to retrieve backup destinations | 

### Return type

[**PlanBackupDestinations**](PlanBackupDestinations.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlanById**
> ServerPlan GetPlanById(ctx, planId)
Get Server Plan Details

Get Plan details 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the plan to fetch details | 

### Return type

[**ServerPlan**](ServerPlan.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerPlanRPO**
> PlanSchedules GetServerPlanRPO(ctx, planId)
Get Server Plan RPO

API to fetch schedules responsible for server plan RPO

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Server Plan | 

### Return type

[**PlanSchedules**](PlanSchedules.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourceCopiesForGivenServerPlanRegion**
> PlanSourceCopies GetSourceCopiesForGivenServerPlanRegion(ctx, planId, optional)
Gives a list of eligible source copies for given region id

Gives a list of eligible source copies for given region id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of plan | 
 **optional** | ***ServerPlanApiGetSourceCopiesForGivenServerPlanRegionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerPlanApiGetSourceCopiesForGivenServerPlanRegionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regionId** | **optional.Int32**| Region Id against which we want to check eligible source copies. Skip if no region present. | 
 **forSnapCopy** | **optional.Bool**| Get list of source copy eligible for snap copy in given region. | 

### Return type

[**PlanSourceCopies**](PlanSourceCopies.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyBackupDestination**
> PlanBackupDestinationResp ModifyBackupDestination(ctx, planId, backupDestinationId, optional)
Modify ServerPlan BackupDestination

Modify Backup Destination for a Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan to modify | 
  **backupDestinationId** | **int32**| Id of the backupDestination to be modified | 
 **optional** | ***ServerPlanApiModifyBackupDestinationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerPlanApiModifyBackupDestinationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpdatePlanBackupDestination**](UpdatePlanBackupDestination.md)|  | 

### Return type

[**PlanBackupDestinationResp**](PlanBackupDestinationResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyBackupDestinationDetails**
> PlanBackupDestinationResp ModifyBackupDestinationDetails(ctx, planId, backupDestinationId, optional)
Update Backup Destination details of ServerPlan

Modify Backup Destination details for a Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan to modify | 
  **backupDestinationId** | **int32**| Id of the backupDestination to be modified | 
 **optional** | ***ServerPlanApiModifyBackupDestinationDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerPlanApiModifyBackupDestinationDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpdatePlanBackupDestinationDetails**](UpdatePlanBackupDestinationDetails.md)|  | 

### Return type

[**PlanBackupDestinationResp**](PlanBackupDestinationResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyPlan**
> GenericResp ModifyPlan(ctx, planId, optional)
Update the server plan

Used to modify an exsiting server plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan to update | 
 **optional** | ***ServerPlanApiModifyPlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerPlanApiModifyPlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateServerPlan**](UpdateServerPlan.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServerPlanRPO**
> GenericResp UpdateServerPlanRPO(ctx, planId, optional)
Update Server Plan RPO

API to modify RPO schedules on server plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Server Plan | 
 **optional** | ***ServerPlanApiUpdateServerPlanRPOOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerPlanApiUpdateServerPlanRPOOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ServerPlanUpdateSchedules**](ServerPlanUpdateSchedules.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V4ServerPlanPlanIdRetentionSynthfullPut**
> GenericResp V4ServerPlanPlanIdRetentionSynthfullPut(ctx, planId)
Update the server plan retention based synthetic full

Used to update synthetic full schedule on plan based on retention settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan to update | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

