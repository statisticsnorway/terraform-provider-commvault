# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLocalAccessPath**](LocalStorageApi.md#AddLocalAccessPath) | **Post** /V4/Storage/Local/{storagePoolId}/BackupLocation/{backupLocationId}/AccessPath | Add Media Agent to Local Storage Path
[**CreateLocalBackupLocation**](LocalStorageApi.md#CreateLocalBackupLocation) | **Post** /V4/Storage/Local/{storagePoolId}/BackupLocation | Create a backup location for local storage
[**CreateLocalStorage**](LocalStorageApi.md#CreateLocalStorage) | **Post** /V4/Storage/Local | Create a local storage pool
[**DeleteLocalAccessPath**](LocalStorageApi.md#DeleteLocalAccessPath) | **Delete** /V4/Storage/Local/{storagePoolId}/BackupLocation/{backupLocationId}/AccessPath/{accessPathId} | Delete Media Agent to Local Storage Path
[**DeleteLocalBackupLocation**](LocalStorageApi.md#DeleteLocalBackupLocation) | **Delete** /V4/Storage/Local/{storagePoolId}/BackupLocation/{backupLocationId} | Delete Backup Location of local storage
[**DeleteLocalStorage**](LocalStorageApi.md#DeleteLocalStorage) | **Delete** /V4/Storage/Local/{storagePoolId} | Delete Local Storage
[**GetLocalBackupLocationDetails**](LocalStorageApi.md#GetLocalBackupLocationDetails) | **Get** /V4/Storage/Local/{storagePoolId}/BackupLocation/{backupLocationId} | Get Backup location details of local storage
[**GetLocalStorageDetails**](LocalStorageApi.md#GetLocalStorageDetails) | **Get** /V4/Storage/Local/{storagePoolId} | Get Local Storage Details
[**GetLocalStorages**](LocalStorageApi.md#GetLocalStorages) | **Get** /V4/Storage/Local | Get All Local Storage
[**LocalGetEligibleMediaAgentsForAccessPath**](LocalStorageApi.md#LocalGetEligibleMediaAgentsForAccessPath) | **Get** /V4/Storage/Local/{storagePoolId}/BackupLocation/{backupLocationId}/AccessPath/MediaAgents | fetch media agents which can be added as access paths for local storage
[**ModifyLocalAccessPath**](LocalStorageApi.md#ModifyLocalAccessPath) | **Put** /V4/Storage/Local/{storagePoolId}/BackupLocation/{backupLocationId}/AccessPath/{accessPathId} | Modify a local storage access path
[**ModifyLocalBackupLocation**](LocalStorageApi.md#ModifyLocalBackupLocation) | **Put** /V4/Storage/Local/{storagePoolId}/BackupLocation/{backupLocationId} | Update Backup Location properties of local storage
[**ModifyLocalStorage**](LocalStorageApi.md#ModifyLocalStorage) | **Put** /V4/Storage/Local/{storagePoolId} | Update Local Storage Properties

# **AddLocalAccessPath**
> GenericResp AddLocalAccessPath(ctx, storagePoolId, backupLocationId, optional)
Add Media Agent to Local Storage Path

Used to add a media agent to a local storage pool access path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the local storage pool whose details have to be fetched to add a new access path | 
  **backupLocationId** | **int32**| Id of the backup location whose details have to be fetched to add a new access path | 
 **optional** | ***LocalStorageApiAddLocalAccessPathOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalStorageApiAddLocalAccessPathOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AddLocalAccessPath**](AddLocalAccessPath.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLocalBackupLocation**
> IdName CreateLocalBackupLocation(ctx, storagePoolId, optional)
Create a backup location for local storage

Create a new backup location for local storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the local storage pool to update | 
 **optional** | ***LocalStorageApiCreateLocalBackupLocationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalStorageApiCreateLocalBackupLocationOpts struct
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

# **CreateLocalStorage**
> IdName CreateLocalStorage(ctx, optional)
Create a local storage pool

Create a new Local storage pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocalStorageApiCreateLocalStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalStorageApiCreateLocalStorageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateLocalStorage**](CreateLocalStorage.md)|  | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLocalAccessPath**
> GenericResp DeleteLocalAccessPath(ctx, storagePoolId, backupLocationId, accessPathId)
Delete Media Agent to Local Storage Path

Used to delete a media agent to a local storage access path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the local storage pool whose access path has to be deleted | 
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

# **DeleteLocalBackupLocation**
> GenericResp DeleteLocalBackupLocation(ctx, storagePoolId, backupLocationId, optional)
Delete Backup Location of local storage

Delete an existing mount path of local storage pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the local storage pool to whose backup location has to be deleted. | 
  **backupLocationId** | **int32**| Id of the backup location to delete | 
 **optional** | ***LocalStorageApiDeleteLocalBackupLocationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalStorageApiDeleteLocalBackupLocationOpts struct
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

# **DeleteLocalStorage**
> GenericResp DeleteLocalStorage(ctx, storagePoolId)
Delete Local Storage

Used to delete a local storage pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the local storage pool to delete | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLocalBackupLocationDetails**
> BackupLocationDetails GetLocalBackupLocationDetails(ctx, storagePoolId, backupLocationId)
Get Backup location details of local storage

Used to fetch mount path details of the local storage pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the local storage pool whose details have to be fetched | 
  **backupLocationId** | **int32**| Id of the backup location whose details have to be fetched | 

### Return type

[**BackupLocationDetails**](BackupLocationDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLocalStorageDetails**
> LocalStorage GetLocalStorageDetails(ctx, storagePoolId, optional)
Get Local Storage Details

Get details of a local storage pool based on id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the local storage pool whose details have to be fetched | 
 **optional** | ***LocalStorageApiGetLocalStorageDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalStorageApiGetLocalStorageDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showInheritedAssociation** | **optional.Bool**| Show inherited security association | 

### Return type

[**LocalStorage**](LocalStorage.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLocalStorages**
> StorageListResponse GetLocalStorages(ctx, )
Get All Local Storage

Get a list of Local storage pools

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

# **LocalGetEligibleMediaAgentsForAccessPath**
> MediaAgentList LocalGetEligibleMediaAgentsForAccessPath(ctx, storagePoolId, backupLocationId)
fetch media agents which can be added as access paths for local storage

Used to fetch available media agents which can be added as access paths for local storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the local storage pool whose media agent has to be shared | 
  **backupLocationId** | **int32**| Id of the backup location of which media agent has to be shared | 

### Return type

[**MediaAgentList**](MediaAgentList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyLocalAccessPath**
> GenericResp ModifyLocalAccessPath(ctx, storagePoolId, backupLocationId, accessPathId, optional)
Modify a local storage access path

Used to modify a local storage access path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the Local storage pool whose access path has to be modified | 
  **backupLocationId** | **int32**| Id of the mount path whose access path has to be modified | 
  **accessPathId** | **int32**| Id of the mount path whose access path has to be modified | 
 **optional** | ***LocalStorageApiModifyLocalAccessPathOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalStorageApiModifyLocalAccessPathOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of UpdateLocalAccessPath**](UpdateLocalAccessPath.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyLocalBackupLocation**
> GenericResp ModifyLocalBackupLocation(ctx, storagePoolId, backupLocationId, optional)
Update Backup Location properties of local storage

Modify the properties of an existing mount path of local storage pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the local storage pool to update | 
  **backupLocationId** | **int32**| Id of the backup location to update | 
 **optional** | ***LocalStorageApiModifyLocalBackupLocationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalStorageApiModifyLocalBackupLocationOpts struct
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

# **ModifyLocalStorage**
> GenericResp ModifyLocalStorage(ctx, storagePoolId, optional)
Update Local Storage Properties

Modify the properties of an existing local storage pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the local storage to update | 
 **optional** | ***LocalStorageApiModifyLocalStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalStorageApiModifyLocalStorageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateLocalStorage**](UpdateLocalStorage.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

