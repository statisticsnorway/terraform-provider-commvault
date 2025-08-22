# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessNode**](VirtualMachinesApi.md#AccessNode) | **Post** /V4/AccessNode | Deploy Access Node
[**CreateKubernetescluster**](VirtualMachinesApi.md#CreateKubernetescluster) | **Post** /V4/Kubernetes/cluster | 
[**DeleteVirtualMachine**](VirtualMachinesApi.md#DeleteVirtualMachine) | **Delete** /V4/VirtualMachines/{vmUUID} | Delete a Virtual Machine
[**GetAccessNodes**](VirtualMachinesApi.md#GetAccessNodes) | **Get** /V4/AccessNodes | Get Access Nodes
[**GetAllAPllicationGroups**](VirtualMachinesApi.md#GetAllAPllicationGroups) | **Get** /V4/Kubernetes/ApplicationGroups | Get All Application Groups of Kubernetes
[**GetKubernetesApplication**](VirtualMachinesApi.md#GetKubernetesApplication) | **Get** /V4/Kubernetes/Applications | Get all application in Kubenetes cluster
[**GetVMDetails**](VirtualMachinesApi.md#GetVMDetails) | **Get** /V4/VirtualMachines/{vmUUID} | Get Virtual Machine Details
[**GetVirtualMachines**](VirtualMachinesApi.md#GetVirtualMachines) | **Get** /V4/VirtualMachines | Get Virtual Machines
[**ListKubernetesCluster**](VirtualMachinesApi.md#ListKubernetesCluster) | **Get** /V4/Kubernetes/cluster | 
[**ModifyVirtualMachine**](VirtualMachinesApi.md#ModifyVirtualMachine) | **Put** /V4/VirtualMachines/{vmUUID} | Modify Virtual Machine Details
[**SetHypervisorFBR**](VirtualMachinesApi.md#SetHypervisorFBR) | **Put** /V4/Hypervisor/{hypervisorId}/FBR | 
[**VMGroupRestore**](VirtualMachinesApi.md#VMGroupRestore) | **Post** /V4/VmGroup/{VmGroupId}/restore | VMGroup Restore Operation
[**VMRestore**](VirtualMachinesApi.md#VMRestore) | **Post** /V4/VM/{vmGuid}/restore | VM Restore Operation
[**VirtualMachineBackup**](VirtualMachinesApi.md#VirtualMachineBackup) | **Post** /V4/VirtualMachines/{vmUUID}/backup | Backup VM Group

# **AccessNode**
> CreateTaskRespforBackup AccessNode(ctx, optional)
Deploy Access Node

API to deploy Access Node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VirtualMachinesApiAccessNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualMachinesApiAccessNodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AccessNodeDeployment**](AccessNodeDeployment.md)|  | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateKubernetescluster**
> CreateClusterResp CreateKubernetescluster(ctx, optional)


Create Kubernetes cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VirtualMachinesApiCreateKubernetesclusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualMachinesApiCreateKubernetesclusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateKubernetesCluster**](CreateKubernetesCluster.md)|  | 

### Return type

[**CreateClusterResp**](CreateClusterResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVirtualMachine**
> GenericResp DeleteVirtualMachine(ctx, vmUUID)
Delete a Virtual Machine

Used to delete a virtual machine

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmUUID** | **string**| The vmUUID can be obtained from GET /virtualMachines UUID property | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccessNodes**
> AccessNodesList GetAccessNodes(ctx, optional)
Get Access Nodes

Endpoint to get the list of access nodes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VirtualMachinesApiGetAccessNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualMachinesApiGetAccessNodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vendor** | **optional.String**| Vendor Name to be filtered | 
 **userId** | **optional.Int32**| user id to be filtered | 

### Return type

[**AccessNodesList**](accessNodesList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllAPllicationGroups**
> ApplicationGroupListResp GetAllAPllicationGroups(ctx, optional)
Get All Application Groups of Kubernetes

Get the details of all applicationGroups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VirtualMachinesApiGetAllAPllicationGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualMachinesApiGetAllAPllicationGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **optional.Int32**| Id of the hypervisor to list application groups | 

### Return type

[**ApplicationGroupListResp**](ApplicationGroupListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKubernetesApplication**
> ApplicationListResp GetKubernetesApplication(ctx, )
Get all application in Kubenetes cluster

Get all application in Kubenetes cluster

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApplicationListResp**](ApplicationListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVMDetails**
> InlineResponse20011 GetVMDetails(ctx, vmUUID, optional)
Get Virtual Machine Details

Get details for virtual machine

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmUUID** | **string**| The vmUUID can be obtained from GET /virtualMachines UUID property | 
 **optional** | ***VirtualMachinesApiGetVMDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualMachinesApiGetVMDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showInheritedAssociation** | **optional.Bool**|  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVirtualMachines**
> VirtualMachineListResp GetVirtualMachines(ctx, )
Get Virtual Machines

Get all virtual machines

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**VirtualMachineListResp**](VirtualMachineListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListKubernetesCluster**
> ListClusters ListKubernetesCluster(ctx, )


Get the details of all cluster

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListClusters**](ListClusters.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyVirtualMachine**
> ModifyVirtualMachine(ctx, vmUUID, optional)
Modify Virtual Machine Details

Modify the properties of an existing virtual machine

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmUUID** | **string**| The vmUUID can be obtained from GET /virtualMachines UUID property | 
 **optional** | ***VirtualMachinesApiModifyVirtualMachineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualMachinesApiModifyVirtualMachineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateVmProperties**](UpdateVmProperties.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetHypervisorFBR**
> CreateHypervisorResp SetHypervisorFBR(ctx, hypervisorId, optional)


Endpoint to set FBR

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hypervisorId** | **string**| Linux Media Agent to set as FBR for  linux File based restores | 
 **optional** | ***VirtualMachinesApiSetHypervisorFBROpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualMachinesApiSetHypervisorFBROpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of HypervisorIdFbrBody**](HypervisorIdFbrBody.md)|  | 

### Return type

[**CreateHypervisorResp**](CreateHypervisorResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VMGroupRestore**
> CreateTaskRespforBackup VMGroupRestore(ctx, vmGroupId, optional)
VMGroup Restore Operation

To restore the virtual machines in vmgroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmGroupId** | **int32**| Id of the VMgroup to backup | 
 **optional** | ***VirtualMachinesApiVMGroupRestoreOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualMachinesApiVMGroupRestoreOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of VmGroupRestoreRequest**](VmGroupRestoreRequest.md)|  | 
 **mediaAgentName** | **optional.**| Media agent name | 
 **mediaAgentId** | **optional.**| Media agent id | 
 **fromTime** | **optional.**| Restore window UTC from time. Valid Formats: yyyy-MM-ddTHH:mm:ss or yyyy-MM-ddTHH:mm:sszzz or yyyy-MM-dd | 
 **toTime** | **optional.**| Restore window UTC to time. Valid Formats: yyyy-MM-ddTHH:mm:ss or yyyy-MM-ddTHH:mm:sszzz or yyyy-MM-dd | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VMRestore**
> CreateTaskRespforBackup VMRestore(ctx, vmGuid, optional)
VM Restore Operation

To restore a virtual machine

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmGuid** | **string**| GUID of VM to backup | 
 **optional** | ***VirtualMachinesApiVMRestoreOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualMachinesApiVMRestoreOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of VmGroupRestoreRequest**](VmGroupRestoreRequest.md)|  | 
 **mediaAgentName** | **optional.**| Media agent name | 
 **mediaAgentId** | **optional.**| Media agent id | 
 **fromTime** | **optional.**| Restore window UTC from time. Valid Formats: yyyy-MM-ddTHH:mm:ss or yyyy-MM-ddTHH:mm:sszzz or yyyy-MM-dd | 
 **toTime** | **optional.**| Restore window UTC to time. Valid Formats: yyyy-MM-ddTHH:mm:ss or yyyy-MM-ddTHH:mm:sszzz or yyyy-MM-dd | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualMachineBackup**
> CreateTaskRespforBackup VirtualMachineBackup(ctx, vmUUID, optional)
Backup VM Group

To Backup the virtual machines in vmgroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmUUID** | **string**| Id of the Virtual Machine to backup | 
 **optional** | ***VirtualMachinesApiVirtualMachineBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualMachinesApiVirtualMachineBackupOpts struct
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

