# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLabVMs**](VMProvisioningApi.md#GetLabVMs) | **Get** /VirtualLab/VirtualMachines | List of VMs in Lab
[**GetVMDetailsByGUID**](VMProvisioningApi.md#GetVMDetailsByGUID) | **Get** /VMProvisioning/VirtualMachine/{vmGUID} | VM Details By GUID
[**GetVirtualLabs**](VMProvisioningApi.md#GetVirtualLabs) | **Get** /VirtualLabs | List of Virtual Labs
[**Getvirtualmachinesnapshots**](VMProvisioningApi.md#Getvirtualmachinesnapshots) | **Get** /VMProvisioning/VirtualMachine/{vmGUID}/Snapshots | Get Snapshots for Virtual Machine
[**PostVMProvisioningVMAction**](VMProvisioningApi.md#PostVMProvisioningVMAction) | **Post** /VMProvisioning/VirtualMachines/Action | Perform Virtual Machine Operations
[**PostVmActionDelete**](VMProvisioningApi.md#PostVmActionDelete) | **Post** /V4/VM/{vmGUID}/Action/Delete | Delete the provisioned virtual machine
[**PostVmActionRefresh**](VMProvisioningApi.md#PostVmActionRefresh) | **Post** /V4/VM/{vmGUID}/Action/Refresh | Refresh the provisioned virtual machine
[**PostVmActionRenew**](VMProvisioningApi.md#PostVmActionRenew) | **Post** /V4/VM/{vmGUID}/Action/Renew | Renew provisioned virtual machine

# **GetLabVMs**
> VmList GetLabVMs(ctx, optional)
List of VMs in Lab

List of virtual machines created under lab

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VMProvisioningApiGetLabVMsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMProvisioningApiGetLabVMsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Float64**| Lab Id | 
 **name** | **optional.String**| Lab Name | 

### Return type

[**VmList**](VMList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVMDetailsByGUID**
> VmDetails GetVMDetailsByGUID(ctx, vmGUID)
VM Details By GUID

Get Virtual Machine Details by GUID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmGUID** | **string**| GUID of Virtual Machine | 

### Return type

[**VmDetails**](VMDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVirtualLabs**
> VmLabList GetVirtualLabs(ctx, )
List of Virtual Labs

Get List of Virtual Labs

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**VmLabList**](VMLabList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getvirtualmachinesnapshots**
> VmSnapshotResp Getvirtualmachinesnapshots(ctx, vmGUID)
Get Snapshots for Virtual Machine

Get snapshots for virtual machine

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmGUID** | **string**|  | 

### Return type

[**VmSnapshotResp**](VMSnapshotResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVMProvisioningVMAction**
> GenericResp PostVMProvisioningVMAction(ctx, optional)
Perform Virtual Machine Operations

Perform VM operations :- PowerOn, PowerOff, Delete, Reset, Edit Owners, Edit Creators

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VMProvisioningApiPostVMProvisioningVMActionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMProvisioningApiPostVMProvisioningVMActionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of VmActionReq**](VmActionReq.md)| VMs list request body | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVmActionDelete**
> GenericResp PostVmActionDelete(ctx, vmGUID)
Delete the provisioned virtual machine

Deletes the VM.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmGUID** | **string**| GUID of the Provisioned VM | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVmActionRefresh**
> GenericResp PostVmActionRefresh(ctx, vmGUID)
Refresh the provisioned virtual machine

Refreshes the VM.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmGUID** | **string**| GUID of the Provisioned VM | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVmActionRenew**
> GenericResp PostVmActionRenew(ctx, vmGUID, optional)
Renew provisioned virtual machine

Renew the VM with the provided timestamp.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmGUID** | **string**| GUID of the Provisioned VM | 
 **optional** | ***VMProvisioningApiPostVmActionRenewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMProvisioningApiPostVmActionRenewOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ActionRenewBody**](ActionRenewBody.md)| The renewal timestamp has to be provided in unix format. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

