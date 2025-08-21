# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSubclientsForReplication**](SnapshotOperationsApi.md#GetSubclientsForReplication) | **Get** /V4/Snaps/snapEngine/{snapEngineId}/Plan/{planId}/Subclients | Get subclients for replication
[**InstantCloneDelete**](SnapshotOperationsApi.md#InstantCloneDelete) | **Post** /V4/StorageArrays/{arrayId}/Snaps/InstantClone/Delete | Delete Instant Clone Snap
[**InstantCloneForceDelete**](SnapshotOperationsApi.md#InstantCloneForceDelete) | **Post** /V4/StorageArrays/{arrayId}/Snaps/InstantClone/Delete/Force | Force Delete Instant Clone Snapshot
[**ListSnapshots**](SnapshotOperationsApi.md#ListSnapshots) | **Get** /V4/StorageArrays/{arrayId}/Snaps | List Snapshots of Storage Array
[**MultiNodeForceUnmount**](SnapshotOperationsApi.md#MultiNodeForceUnmount) | **Post** /V4/StorageArrays/{arrayId}/Snaps/MultiNode/Unmount/Force | Force Unmount Multi Node Snapshots of Storage Array
[**MultiNodeUnmount**](SnapshotOperationsApi.md#MultiNodeUnmount) | **Post** /V4/StorageArrays/{arrayId}/Snaps/MultiNode/Unmount | Unmount Multi Node Snapshots of Storage Array
[**SnapByJobs**](SnapshotOperationsApi.md#SnapByJobs) | **Get** /V4/StorageArrays/Job/{jobId}/Snaps | Get Snaps for a Job
[**SnapDelete**](SnapshotOperationsApi.md#SnapDelete) | **Post** /V4/StorageArrays/{arrayId}/Snaps/Delete | Delete Storage Array Snapshot
[**SnapForceDelete**](SnapshotOperationsApi.md#SnapForceDelete) | **Post** /V4/StorageArrays/{arrayId}/Snaps/Delete/Force | Delete Snapshot Force
[**SnapForceUnmount**](SnapshotOperationsApi.md#SnapForceUnmount) | **Post** /V4/StorageArrays/{arrayId}/Snaps/Unmount/Force | Force Unmount Snapshot
[**SnapMount**](SnapshotOperationsApi.md#SnapMount) | **Post** /V4/StorageArrays/{arrayId}/Snaps/Mount | Mount Storage Array Snapshot
[**SnapReconWorkflow**](SnapshotOperationsApi.md#SnapReconWorkflow) | **Put** /V4/StorageArrays/{arrayId}/Snaps/Recon | Reconcile Snapshots of Storage Array
[**SnapRevert**](SnapshotOperationsApi.md#SnapRevert) | **Post** /V4/StorageArrays/{arrayId}/Snaps/Revert | Revert Snapshot of Storage Array
[**SnapUnmount**](SnapshotOperationsApi.md#SnapUnmount) | **Post** /V4/StorageArrays/{arrayId}/Snaps/Unmount | Unmount Snapshot of Storage Array

# **GetSubclientsForReplication**
> GetSubclientForSnapEngine GetSubclientsForReplication(ctx, planId, snapEngineId, optional)
Get subclients for replication

Get subclients for requested snapEngineId for requested planId and copyId.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| plan Id | 
  **snapEngineId** | **int32**| Snap Engine Id | 
 **optional** | ***SnapshotOperationsApiGetSubclientsForReplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotOperationsApiGetSubclientsForReplicationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **copyId** | **optional.Int32**| Copy Id | 

### Return type

[**GetSubclientForSnapEngine**](GetSubclientForSnapEngine.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstantCloneDelete**
> SnapReconResp InstantCloneDelete(ctx, arrayId, optional)
Delete Instant Clone Snap

API for snap instant clone delete operation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 
 **optional** | ***SnapshotOperationsApiInstantCloneDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotOperationsApiInstantCloneDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SnapNodeReq**](SnapNodeReq.md)|  | 

### Return type

[**SnapReconResp**](SnapReconResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstantCloneForceDelete**
> SnapReconResp InstantCloneForceDelete(ctx, arrayId, optional)
Force Delete Instant Clone Snapshot

API for snap instant clone force delete operation. Please be aware, triggering this API will not delete the clones from the storage array, so ensure that you delete clones on the array before using this operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 
 **optional** | ***SnapshotOperationsApiInstantCloneForceDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotOperationsApiInstantCloneForceDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SnapNodeReq**](SnapNodeReq.md)|  | 

### Return type

[**SnapReconResp**](SnapReconResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSnapshots**
> InlineResponse20023 ListSnapshots(ctx, arrayId)
List Snapshots of Storage Array

Listing Snapshot of a particular array

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MultiNodeForceUnmount**
> SnapReconResp MultiNodeForceUnmount(ctx, arrayId, optional)
Force Unmount Multi Node Snapshots of Storage Array

API for snap multi node force unmount operation. Please be aware, triggering this API will not physically unmount the clones from the storage array or MA, so ensure that you unmapped the clones on the array and cleanup the mount host and LVM before using this operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 
 **optional** | ***SnapshotOperationsApiMultiNodeForceUnmountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotOperationsApiMultiNodeForceUnmountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SnapNodeReq**](SnapNodeReq.md)|  | 

### Return type

[**SnapReconResp**](SnapReconResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MultiNodeUnmount**
> SnapReconResp MultiNodeUnmount(ctx, arrayId, optional)
Unmount Multi Node Snapshots of Storage Array

API for snap multi node unmount operation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 
 **optional** | ***SnapshotOperationsApiMultiNodeUnmountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotOperationsApiMultiNodeUnmountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SnapNodeReq**](SnapNodeReq.md)|  | 

### Return type

[**SnapReconResp**](SnapReconResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapByJobs**
> InlineResponse20023 SnapByJobs(ctx, jobId, optional)
Get Snaps for a Job

API to list snaps for a particular Job Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **int32**|  | 
 **optional** | ***SnapshotOperationsApiSnapByJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotOperationsApiSnapByJobsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commcellId** | **optional.Int32**|  | [default to 2]

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapDelete**
> SnapReconResp SnapDelete(ctx, arrayId, optional)
Delete Storage Array Snapshot

Deleting an unmounted or created snapshot from Storage Arrays

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 
 **optional** | ***SnapshotOperationsApiSnapDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotOperationsApiSnapDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DeleteSnapReq**](DeleteSnapReq.md)|  | 

### Return type

[**SnapReconResp**](SnapReconResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapForceDelete**
> SnapReconResp SnapForceDelete(ctx, arrayId, optional)
Delete Snapshot Force

Force deleting an unmounted or created snapshot from Storage Arrays. Please be aware, triggering this API will not delete the snapshots from the storage array, so ensure that you delete snapshots on the array before using this operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 
 **optional** | ***SnapshotOperationsApiSnapForceDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotOperationsApiSnapForceDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SnapForceReq**](SnapForceReq.md)|  | 

### Return type

[**SnapReconResp**](SnapReconResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapForceUnmount**
> SnapReconResp SnapForceUnmount(ctx, arrayId, optional)
Force Unmount Snapshot

Force unmounting a mounted Snapshot from Storage Array. Please be aware, triggering this API will not physically unmount the snapshots from the storage array or MA, so ensure that you unmapped the snapshots on the array and cleanup the mount host and LVM before using this operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 
 **optional** | ***SnapshotOperationsApiSnapForceUnmountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotOperationsApiSnapForceUnmountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SnapForceReq**](SnapForceReq.md)|  | 

### Return type

[**SnapReconResp**](SnapReconResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapMount**
> SnapReconResp SnapMount(ctx, arrayId, optional)
Mount Storage Array Snapshot

Mounting a snapshot from the Storage Array

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 
 **optional** | ***SnapshotOperationsApiSnapMountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotOperationsApiSnapMountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MountReq**](MountReq.md)|  | 

### Return type

[**SnapReconResp**](SnapReconResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapReconWorkflow**
> SnapReconResp SnapReconWorkflow(ctx, arrayId)
Reconcile Snapshots of Storage Array

Starts a workflow snap reconcile for array

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **string**|  | 

### Return type

[**SnapReconResp**](SnapReconResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapRevert**
> SnapReconResp SnapRevert(ctx, arrayId, optional)
Revert Snapshot of Storage Array

Reverting a created or unmounted snapshot from the storage array

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 
 **optional** | ***SnapshotOperationsApiSnapRevertOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotOperationsApiSnapRevertOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SnapRevertReq**](SnapRevertReq.md)|  | 

### Return type

[**SnapReconResp**](SnapReconResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapUnmount**
> SnapReconResp SnapUnmount(ctx, arrayId, optional)
Unmount Snapshot of Storage Array

Unmounting a mounted Snapshot from Storage Array

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 
 **optional** | ***SnapshotOperationsApiSnapUnmountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotOperationsApiSnapUnmountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UnmountSnapReq**](UnmountSnapReq.md)|  | 

### Return type

[**SnapReconResp**](SnapReconResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

