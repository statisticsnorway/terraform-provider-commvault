# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMediaAgent**](DiskStorageApi.md#AddMediaAgent) | **Post** /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}/AccessPath | Add Media Agent to Disk Path
[**CreateBackupLocation**](DiskStorageApi.md#CreateBackupLocation) | **Post** /V4/Storage/Disk/{storagePoolId}/BackupLocation | Create Backup Location
[**CreateDiskStorage**](DiskStorageApi.md#CreateDiskStorage) | **Post** /V4/Storage/Disk | Create a disk storage pool
[**DeleteBackupLocation**](DiskStorageApi.md#DeleteBackupLocation) | **Delete** /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId} | Delete Backup Location
[**DeleteDiskAccessPath**](DiskStorageApi.md#DeleteDiskAccessPath) | **Delete** /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}/AccessPath/{accessPathId} | Delete Media Agent to Disk Path
[**DeleteDiskStorage**](DiskStorageApi.md#DeleteDiskStorage) | **Delete** /V4/Storage/Disk/{storagePoolId} | Delete Disk Storage
[**DiskGetEligibleMediaAgentsForAccessPath**](DiskStorageApi.md#DiskGetEligibleMediaAgentsForAccessPath) | **Get** /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}/AccessPath/MediaAgents | 
[**GetBackupLocationDetails**](DiskStorageApi.md#GetBackupLocationDetails) | **Get** /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId} | Get Backup location details
[**GetDiskStorageDetails**](DiskStorageApi.md#GetDiskStorageDetails) | **Get** /V4/Storage/Disk/{storagePoolId} | Get Disk Storage Details
[**GetDiskStorages**](DiskStorageApi.md#GetDiskStorages) | **Get** /V4/Storage/Disk | Get Disk Storage
[**ModifyBackupLocation**](DiskStorageApi.md#ModifyBackupLocation) | **Put** /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId} | Update Backup Location properties
[**ModifyDiskAccessPath**](DiskStorageApi.md#ModifyDiskAccessPath) | **Put** /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}/AccessPath/{accessPathId} | Modify a disk access path
[**ModifyDiskStorage**](DiskStorageApi.md#ModifyDiskStorage) | **Put** /V4/Storage/Disk/{storagePoolId} | Update Disk Storage Properties

# **AddMediaAgent**
> GenericResp AddMediaAgent(ctx, storagePoolId, backupLocationId, optional)
Add Media Agent to Disk Path

Used to add a media agent to a disk access path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the disk storage pool whose details have to be fetched to add a new access path | 
  **backupLocationId** | **int32**| Id of the backup location whose details have to be fetched to add a new access path | 
 **optional** | ***DiskStorageApiAddMediaAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DiskStorageApiAddMediaAgentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AddDiskAccessPath**](AddDiskAccessPath.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBackupLocation**
> IdName CreateBackupLocation(ctx, storagePoolId, optional)
Create Backup Location

Create a new backup location

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the disk storage to update | 
 **optional** | ***DiskStorageApiCreateBackupLocationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DiskStorageApiCreateBackupLocationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateBackupLocation**](CreateBackupLocation.md)|  | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDiskStorage**
> IdName CreateDiskStorage(ctx, optional)
Create a disk storage pool

Create a new disk storage pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DiskStorageApiCreateDiskStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DiskStorageApiCreateDiskStorageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateDiskStorage**](CreateDiskStorage.md)|  | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBackupLocation**
> GenericResp DeleteBackupLocation(ctx, storagePoolId, backupLocationId, optional)
Delete Backup Location

Modify the properties of an existing mount path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the disk storage to whose backup location has to be deleted | 
  **backupLocationId** | **int32**| Id of the backup location to delete | 
 **optional** | ***DiskStorageApiDeleteBackupLocationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DiskStorageApiDeleteBackupLocationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceDelete** | **optional.Bool**| Force deletes a backup location. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDiskAccessPath**
> GenericResp DeleteDiskAccessPath(ctx, storagePoolId, backupLocationId, accessPathId)
Delete Media Agent to Disk Path

Used to delete a media agent to a disk access path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the disk storage pool whose access path has to be deleted | 
  **backupLocationId** | **int32**| Id of the mount path whose access path has to be deleted | 
  **accessPathId** | **int32**| Id of the mount path whose access path has to be deleted | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDiskStorage**
> GenericResp DeleteDiskStorage(ctx, storagePoolId)
Delete Disk Storage

Used to delete a disk storage pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the disk storage pool to delete | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiskGetEligibleMediaAgentsForAccessPath**
> MediaAgentList DiskGetEligibleMediaAgentsForAccessPath(ctx, storagePoolId, backupLocationId)


Used to fetch available media agents which can be added as access paths for disk storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the disk storage pool whose media agent has to be shared | 
  **backupLocationId** | **int32**| Id of the back up location of which media agent has to be shared | 

### Return type

[**MediaAgentList**](MediaAgentList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupLocationDetails**
> BackupLocationDetails GetBackupLocationDetails(ctx, storagePoolId, backupLocationId)
Get Backup location details

Used to fetch mount path details of the disk storage pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the disk storage pool whose details have to be fetched | 
  **backupLocationId** | **int32**| Id of the backup location whose details have to be fetched | 

### Return type

[**BackupLocationDetails**](BackupLocationDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiskStorageDetails**
> DiskStorage GetDiskStorageDetails(ctx, storagePoolId, optional)
Get Disk Storage Details

Get details of a disk storage pool based on id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the disk storage pool whose details have to be fetched | 
 **optional** | ***DiskStorageApiGetDiskStorageDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DiskStorageApiGetDiskStorageDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showInheritedAssociation** | **optional.Bool**| Show inherited security association | 

### Return type

[**DiskStorage**](DiskStorage.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiskStorages**
> StorageListResponse GetDiskStorages(ctx, )
Get Disk Storage

Get a list of disk storage pools

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**StorageListResponse**](StorageListResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyBackupLocation**
> GenericResp ModifyBackupLocation(ctx, storagePoolId, backupLocationId, optional)
Update Backup Location properties

Modify the properties of an existing mount path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the disk storage to update | 
  **backupLocationId** | **int32**| Id of the backup location to update | 
 **optional** | ***DiskStorageApiModifyBackupLocationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DiskStorageApiModifyBackupLocationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpdateBackupLocation**](UpdateBackupLocation.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyDiskAccessPath**
> GenericResp ModifyDiskAccessPath(ctx, storagePoolId, backupLocationId, accessPathId, optional)
Modify a disk access path

Used to modify a disk access path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the disk storage pool whose access path has to be modified | 
  **backupLocationId** | **int32**| Id of the mount path whose access path has to be modified | 
  **accessPathId** | **int32**| Id of the mount path whose access path has to be modified | 
 **optional** | ***DiskStorageApiModifyDiskAccessPathOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DiskStorageApiModifyDiskAccessPathOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of UpdateDiskAccessPath**](UpdateDiskAccessPath.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyDiskStorage**
> GenericResp ModifyDiskStorage(ctx, storagePoolId, optional)
Update Disk Storage Properties

Modify the properties of an existing disk storage pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the disk storage to update | 
 **optional** | ***DiskStorageApiModifyDiskStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DiskStorageApiModifyDiskStorageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateDiskStorage**](UpdateDiskStorage.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

