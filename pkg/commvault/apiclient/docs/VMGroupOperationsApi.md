# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVMGroup**](VMGroupOperationsApi.md#CreateVMGroup) | **Post** /V4/VMGroup | Create a VM Group
[**DeleteVMGroup**](VMGroupOperationsApi.md#DeleteVMGroup) | **Delete** /V4/VmGroup/{VmGroupId} | 
[**GetAllVMGroups**](VMGroupOperationsApi.md#GetAllVMGroups) | **Get** /V4/VMGroups | Get VM Groups
[**GetVMGroup**](VMGroupOperationsApi.md#GetVMGroup) | **Get** /V4/VmGroup/{VmGroupId} | Get VM Group Details
[**UpdateVMGroup**](VMGroupOperationsApi.md#UpdateVMGroup) | **Put** /V4/VmGroup/{VmGroupId} | Modify VM Group Details
[**VMGroupBackup**](VMGroupOperationsApi.md#VMGroupBackup) | **Post** /V4/VmGroup/{VmGroupId}/backup | Backup VM Group
[**VmGroupPreview**](VMGroupOperationsApi.md#VmGroupPreview) | **Post** /V4/VmGroup/Preview | Preview VM Group

# **CreateVMGroup**
> CreateVmGroupResp CreateVMGroup(ctx, optional)
Create a VM Group

Create a VM Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VMGroupOperationsApiCreateVMGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMGroupOperationsApiCreateVMGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreatevmGroupReq**](CreatevmGroupReq.md)|  | 

### Return type

[**CreateVmGroupResp**](CreateVMGroupResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVMGroup**
> VmgroupdeleteResponse DeleteVMGroup(ctx, vmGroupId)


delete an existing vm group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmGroupId** | **int32**| Id of the vmgroup to delete | 

### Return type

[**VmgroupdeleteResponse**](vmgroupdeleteResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllVMGroups**
> InlineResponse20010 GetAllVMGroups(ctx, optional)
Get VM Groups

Get the details of all vmGroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VMGroupOperationsApiGetAllVMGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMGroupOperationsApiGetAllVMGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hypervisorId** | **optional.Int32**| Id of the hypervisor to list vm groups | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVMGroup**
> GetVmGroupResp GetVMGroup(ctx, vmGroupId)
Get VM Group Details

Get the details of  vmGroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmGroupId** | **int32**| Id of the VmGroup to get | 

### Return type

[**GetVmGroupResp**](GetVMGroupResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVMGroup**
> GenericRespWithWarning UpdateVMGroup(ctx, vmGroupId, optional)
Modify VM Group Details

Updates the VM Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmGroupId** | **int32**| Id of the VMgroup to update | 
 **optional** | ***VMGroupOperationsApiUpdateVMGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMGroupOperationsApiUpdateVMGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdatevmGroupReq**](UpdatevmGroupReq.md)|  | 

### Return type

[**GenericRespWithWarning**](GenericRespWithWarning.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VMGroupBackup**
> CreateTaskRespforBackup VMGroupBackup(ctx, vmGroupId, optional)
Backup VM Group

To Backup the virtual machines in vmgroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmGroupId** | **int32**| Id of the VMgroup to backup | 
 **optional** | ***VMGroupOperationsApiVMGroupBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMGroupOperationsApiVMGroupBackupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backupLevel** | **optional.String**| Backup level , Default :Incremental | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmGroupPreview**
> VmGroupPreviewResp VmGroupPreview(ctx, optional)
Preview VM Group

Preview of the vm to be protected in VMGroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VMGroupOperationsApiVmGroupPreviewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMGroupOperationsApiVmGroupPreviewOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of VmGroupPreviewReq**](VmGroupPreviewReq.md)|  | 
 **hypervisorId** | **optional.**| Hypervisor Id of VMGroup | 
 **vmgroupId** | **optional.**| Id of VMGroup | 

### Return type

[**VmGroupPreviewResp**](vmGroupPreviewResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

