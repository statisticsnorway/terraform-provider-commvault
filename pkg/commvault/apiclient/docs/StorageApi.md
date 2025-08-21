# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateMAWithStoragePool**](StorageApi.md#AssociateMAWithStoragePool) | **Put** /V4/StoragePool/{storagePoolId}/MediaAgent | Associate DDB or storage MediaAgent with the Storage Pool
[**DDBSpaceReclaimOnStoragePool**](StorageApi.md#DDBSpaceReclaimOnStoragePool) | **Put** /V4/StoragePool/{storagePoolId}/DDB/SpaceReclaim | DDB Space Reclamation operation on Storage pool
[**DDBVerificationOnStoragePool**](StorageApi.md#DDBVerificationOnStoragePool) | **Put** /V4/StoragePool/{storagePoolId}/DDB/Verify | DDB Verification operation on Storage pool
[**DisableMaintenanceForStoragePool**](StorageApi.md#DisableMaintenanceForStoragePool) | **Put** /V4/StoragePool/{StoragePoolId}/maintenance/Disable | Disable maintenance for a storage pool
[**EnableAirgapOnStoragePool**](StorageApi.md#EnableAirgapOnStoragePool) | **Put** /V4/StoragePool/{storagePoolId}/Airgap | Enable Airgap on StoragePool
[**EnableMaintenanceForStoragePool**](StorageApi.md#EnableMaintenanceForStoragePool) | **Put** /V4/StoragePool/{StoragePoolId}/maintenance/Enable | Enable maintenance for a storage pool
[**GetMountPathContent**](StorageApi.md#GetMountPathContent) | **Get** /V4/MountPath/Content | Get mount path content
[**GetRunningJobsForStoragePool**](StorageApi.md#GetRunningJobsForStoragePool) | **Get** /V4/StoragePool/{storagePoolId}/runningJobs | Get running jobs for a storage pool
[**GetStorageAssociatedCopies**](StorageApi.md#GetStorageAssociatedCopies) | **Get** /V4/Storage/{storagePoolId}/associatedCopies | 

# **AssociateMAWithStoragePool**
> GenericResp AssociateMAWithStoragePool(ctx, storagePoolId, mediaAgentId, action)
Associate DDB or storage MediaAgent with the Storage Pool

API to manage MediaAgents associated with the Storage Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the Storage Pool | 
  **mediaAgentId** | **int32**| Id of the MediaAgent | 
  **action** | **string**| Action to performed on the MediaAgent | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DDBSpaceReclaimOnStoragePool**
> ArchiveCheckOperationResp DDBSpaceReclaimOnStoragePool(ctx, storagePoolId, dDBVerificationLevel)
DDB Space Reclamation operation on Storage pool

API to run DDB Space Reclamation jobs on SIDB Stores associated storage pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the Storage Pool | 
  **dDBVerificationLevel** | **string**| Deduplication Database and Disk Data Verification Level | 

### Return type

[**ArchiveCheckOperationResp**](ArchiveCheckOperationResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DDBVerificationOnStoragePool**
> ArchiveCheckOperationResp DDBVerificationOnStoragePool(ctx, storagePoolId, dDBVerificationLevel)
DDB Verification operation on Storage pool

API to run DDB verification jobs on SIDB Stores associated storage pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the Storage Pool | 
  **dDBVerificationLevel** | **string**| Deduplication Database and Disk Data Verification Level | 

### Return type

[**ArchiveCheckOperationResp**](ArchiveCheckOperationResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableMaintenanceForStoragePool**
> GenericResponse DisableMaintenanceForStoragePool(ctx, storagePoolId)
Disable maintenance for a storage pool

API is used to disable maintenance for a storage pool. It will take the libraries associated to storage pool out maintenance mode and if the storage pool is HyperScale, then same would be applied to the nodes associated to the storage pool. Storage pool will be marked availabe for upcoming jobs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of storage pool | 

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableAirgapOnStoragePool**
> GenericResp EnableAirgapOnStoragePool(ctx, storagePoolId)
Enable Airgap on StoragePool

This API enables Airgap on StoragePool and all its associated MediaAgents. This action is irreversible, once enabled Airgap cannot be disabled on the StoragePool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the StoragePool | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableMaintenanceForStoragePool**
> GenericResponse EnableMaintenanceForStoragePool(ctx, storagePoolId, optional)
Enable maintenance for a storage pool

API is used to enable maintenance for a storage pool. It will put the libraries associated to storage pool in maintenance mode and if the storage pool is HyperScale, then same would be applied to the nodes associated to the storage pool. When a storage pool is in maintenance, it will not be used by any upcoming jobs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of storage pool | 
 **optional** | ***StorageApiEnableMaintenanceForStoragePoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageApiEnableMaintenanceForStoragePoolOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suspendRunningJobs** | **optional.Bool**| Suspend the jobs running to the storage pool | [default to true]

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMountPathContent**
> GetMountPathContentResp GetMountPathContent(ctx, optional)
Get mount path content

API to get mountpath content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageApiGetMountPathContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageApiGetMountPathContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mountPathId** | **optional.Int32**| Mountpath Id | 
 **deviceId** | **optional.Int32**| Device Id | 

### Return type

[**GetMountPathContentResp**](GetMountPathContentResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRunningJobsForStoragePool**
> GetRunningJobsIdList GetRunningJobsForStoragePool(ctx, storagePoolId, optional)
Get running jobs for a storage pool

API to get jobs running to a storage pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of storage pool | 
 **optional** | ***StorageApiGetRunningJobsForStoragePoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageApiGetRunningJobsForStoragePoolOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobType** | **optional.String**| Type of job | [default to All]
 **ignoreSuspendedJobs** | **optional.Bool**| To include/exlude suspended jobs in the response | [default to true]

### Return type

[**GetRunningJobsIdList**](GetRunningJobsIdList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStorageAssociatedCopies**
> StorageAssociatedCopies GetStorageAssociatedCopies(ctx, storagePoolId)


Get associated copies and plan which are utilizing this storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the storage pool whose associated copies have to be fetched | 

### Return type

[**StorageAssociatedCopies**](StorageAssociatedCopies.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

