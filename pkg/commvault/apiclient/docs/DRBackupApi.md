# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDRBackupDestination**](DRBackupApi.md#CreateDRBackupDestination) | **Post** /V4/DRBackup/BackupDestinations | Create a backup destination for CommServe DR Backups
[**DeleteDRBackupDestination**](DRBackupApi.md#DeleteDRBackupDestination) | **Delete** /V4/DRBackup/BackupDestination/{BackupDestinationId} | Delete Backup Destination of Commserve DR backup
[**GetDRBackupDestinationDetails**](DRBackupApi.md#GetDRBackupDestinationDetails) | **Get** /V4/DRBackup/BackupDestination/{BackupDestinationId} | Get Backup Destination details of DR backup destination
[**GetDRBackupDestinations**](DRBackupApi.md#GetDRBackupDestinations) | **Get** /V4/DRBackup/BackupDestinations | Get backup destinations for CommServe DR Backups
[**GetDRBackupStoragePolicy**](DRBackupApi.md#GetDRBackupStoragePolicy) | **Get** /V4/DRBackup/StoragePolicy | Get details of DR backup storage policy
[**GetSourceCopiesForDRBackup**](DRBackupApi.md#GetSourceCopiesForDRBackup) | **Get** /V4/DRBackup/BackupDestination/SourceCopy/Eligible | Gives a list of eligible source copies for CommServe DR storage policy
[**ModifyDRBackupDestinationDetails**](DRBackupApi.md#ModifyDRBackupDestinationDetails) | **Put** /V4/DRBackup/BackupDestination/{BackupDestinationId} | Update Backup Destination details of DR backup destination

# **CreateDRBackupDestination**
> PlanBackupDestinationResp CreateDRBackupDestination(ctx, optional)
Create a backup destination for CommServe DR Backups

Create a backup destination for CommServe DR Backups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DRBackupApiCreateDRBackupDestinationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DRBackupApiCreateDRBackupDestinationOpts struct
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

# **DeleteDRBackupDestination**
> GenericResp DeleteDRBackupDestination(ctx, backupDestinationId)
Delete Backup Destination of Commserve DR backup

Delete DR Backup Destination for Commserve DR backup

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

# **GetDRBackupDestinationDetails**
> PlanBackupDestinationDetails GetDRBackupDestinationDetails(ctx, backupDestinationId)
Get Backup Destination details of DR backup destination

Get specific backup destination details for Commserve DR Backups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupDestinationId** | **int32**| Id of the backup destination | 

### Return type

[**PlanBackupDestinationDetails**](PlanBackupDestinationDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDRBackupDestinations**
> PlanBackupDestinations GetDRBackupDestinations(ctx, )
Get backup destinations for CommServe DR Backups

Get backup destinations for CommServe DR Backups

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PlanBackupDestinations**](PlanBackupDestinations.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDRBackupStoragePolicy**
> IdName GetDRBackupStoragePolicy(ctx, )
Get details of DR backup storage policy

Get DR backup storage policy details

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourceCopiesForDRBackup**
> PlanSourceCopies GetSourceCopiesForDRBackup(ctx, )
Gives a list of eligible source copies for CommServe DR storage policy

Gives a list of eligible source copies for CommServe DR storage policy

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PlanSourceCopies**](PlanSourceCopies.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyDRBackupDestinationDetails**
> PlanBackupDestinationResp ModifyDRBackupDestinationDetails(ctx, backupDestinationId, optional)
Update Backup Destination details of DR backup destination

Modify Backup Destination details for CommServe DR backups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupDestinationId** | **int32**| Id of the backupDestination to be modified | 
 **optional** | ***DRBackupApiModifyDRBackupDestinationDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DRBackupApiModifyDRBackupDestinationDetailsOpts struct
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

