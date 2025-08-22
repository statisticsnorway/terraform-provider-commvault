# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerGroups**](ServerGroupApi.md#CreateServerGroups) | **Post** /V4/ServerGroup | Create a Server Group
[**DeleteServerGroup**](ServerGroupApi.md#DeleteServerGroup) | **Delete** /V4/ServerGroup/{serverGroupId} | Delete a ServerGroup
[**DeleteServerGroupDCPlan**](ServerGroupApi.md#DeleteServerGroupDCPlan) | **Delete** /V4/ServerGroup/{serverGroupId}/DCPlan | Remove associated DC plan from server group
[**DisableBackupServerGroup**](ServerGroupApi.md#DisableBackupServerGroup) | **Put** /V4/ServerGroup/{serverGroupId}/Backup/Action/Disable | Disable backups for ServerGroup
[**DisableDataAgingServerGroup**](ServerGroupApi.md#DisableDataAgingServerGroup) | **Put** /V4/ServerGroup/{serverGroupId}/DataAging/Action/Disable | Disable dataaging for ServerGroup
[**DisableRestoreServerGroup**](ServerGroupApi.md#DisableRestoreServerGroup) | **Put** /V4/ServerGroup/{serverGroupId}/Restore/Action/Disable | Disable Restore for servergroup
[**EnableBackupServerGroup**](ServerGroupApi.md#EnableBackupServerGroup) | **Put** /V4/ServerGroup/{serverGroupId}/Backup/Action/Enable | Enable Backup for ServerGroup
[**EnableDataAgingServerGroup**](ServerGroupApi.md#EnableDataAgingServerGroup) | **Put** /V4/ServerGroup/{serverGroupId}/DataAging/Action/Enable | Enable Data aging for Server Group
[**EnableRestoreServerGroup**](ServerGroupApi.md#EnableRestoreServerGroup) | **Put** /V4/ServerGroup/{serverGroupId}/Restore/Action/Enable | Enable restore for ServerGroup
[**GetServerGroupDCPlan**](ServerGroupApi.md#GetServerGroupDCPlan) | **Get** /V4/ServerGroup/{serverGroupId}/DCPlan | Get ServerGroup DC Plan
[**GetServerGroupIdDetails**](ServerGroupApi.md#GetServerGroupIdDetails) | **Get** /V4/ServerGroup/{serverGroupId} | Get Server Group Details
[**GetServerGroups**](ServerGroupApi.md#GetServerGroups) | **Get** /V4/ServerGroup | Get Server Groups
[**PutServerGroupDCPlan**](ServerGroupApi.md#PutServerGroupDCPlan) | **Put** /V4/ServerGroup/{serverGroupId}/DCPlan | Associate specific plan to servergroup
[**RefreshServerGroupAssociation**](ServerGroupApi.md#RefreshServerGroupAssociation) | **Put** /V4/ServerGroup/{serverGroupId}/Refresh | Refresh the server group associations
[**ServerGroupsCreationPreview**](ServerGroupApi.md#ServerGroupsCreationPreview) | **Post** /V4/ServerGroup/Preview | Server Group Preview
[**UpdateServerGroupAssociation**](ServerGroupApi.md#UpdateServerGroupAssociation) | **Put** /V4/ServerGroup/{serverGroupId} | Update the server group

# **CreateServerGroups**
> CreateServerGroupResp CreateServerGroups(ctx, optional)
Create a Server Group

Create Server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServerGroupApiCreateServerGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerGroupApiCreateServerGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateServerGroup**](CreateServerGroup.md)|  | 

### Return type

[**CreateServerGroupResp**](CreateServerGroupResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServerGroup**
> GenericResp DeleteServerGroup(ctx, serverGroupId)
Delete a ServerGroup

Used to delete a serverGroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverGroupId** | **int32**| Id of the serverGroup to delete | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServerGroupDCPlan**
> GenericResp DeleteServerGroupDCPlan(ctx, serverGroupId, optional)
Remove associated DC plan from server group

Api to remove associated DC plan from server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverGroupId** | **string**|  | 
 **optional** | ***ServerGroupApiDeleteServerGroupDCPlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerGroupApiDeleteServerGroupDCPlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeChildAssociations** | **optional.Bool**| Optional parameter. To be set only in case, if caller want to leave child association as-is and decouple them from server group | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableBackupServerGroup**
> GenericResp DisableBackupServerGroup(ctx, serverGroupId, optional)
Disable backups for ServerGroup

Used to disable backup property for a server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverGroupId** | **int32**| Id of the serverGroup to modify | 
 **optional** | ***ServerGroupApiDisableBackupServerGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerGroupApiDisableBackupServerGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableAfterADelay** | **optional.Int32**| Provide UTC time in unix format. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableDataAgingServerGroup**
> GenericResp DisableDataAgingServerGroup(ctx, serverGroupId, optional)
Disable dataaging for ServerGroup

Used to disable data aging property for a server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverGroupId** | **int32**| Id of the serverGroup to modify | 
 **optional** | ***ServerGroupApiDisableDataAgingServerGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerGroupApiDisableDataAgingServerGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableAfterADelay** | **optional.Int32**| Provide UTC time in unix format. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableRestoreServerGroup**
> GenericResp DisableRestoreServerGroup(ctx, serverGroupId, optional)
Disable Restore for servergroup

Used to disable restore property for a server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverGroupId** | **int32**| Id of the serverGroup to modify | 
 **optional** | ***ServerGroupApiDisableRestoreServerGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerGroupApiDisableRestoreServerGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableAfterADelay** | **optional.Int32**| Provide UTC time in unix format. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableBackupServerGroup**
> GenericResp EnableBackupServerGroup(ctx, serverGroupId)
Enable Backup for ServerGroup

Used to enable backup property for a server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverGroupId** | **int32**| Id of the serverGroup to modify | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableDataAgingServerGroup**
> GenericResp EnableDataAgingServerGroup(ctx, serverGroupId)
Enable Data aging for Server Group

Used to enable data aging property for a server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverGroupId** | **int32**| Id of the serverGroup to modify | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableRestoreServerGroup**
> GenericResp EnableRestoreServerGroup(ctx, serverGroupId)
Enable restore for ServerGroup

Used to enable restore property for a server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverGroupId** | **int32**| Id of the serverGroup to modify | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerGroupDCPlan**
> IdName GetServerGroupDCPlan(ctx, serverGroupId)
Get ServerGroup DC Plan

API to fetch DC plan associated to server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverGroupId** | **string**|  | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerGroupIdDetails**
> ServerGroup GetServerGroupIdDetails(ctx, serverGroupId, optional)
Get Server Group Details

Get details of a serverGroup based on id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverGroupId** | **int32**| Id of the serverGroupId whose details have to be fetched | 
 **optional** | ***ServerGroupApiGetServerGroupIdDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerGroupApiGetServerGroupIdDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | [**optional.Interface of Mode**](.md)|  | 

### Return type

[**ServerGroup**](ServerGroup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerGroups**
> GetServerGroupsResp GetServerGroups(ctx, )
Get Server Groups

Get All Server Groups

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetServerGroupsResp**](GetServerGroupsResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutServerGroupDCPlan**
> GenericResp PutServerGroupDCPlan(ctx, serverGroupId, optional)
Associate specific plan to servergroup

API to associate data classification plan to servergroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverGroupId** | **string**|  | 
 **optional** | ***ServerGroupApiPutServerGroupDCPlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerGroupApiPutServerGroupDCPlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdName**](IdName.md)| DC Plan id and name | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshServerGroupAssociation**
> GenericResponse RefreshServerGroupAssociation(ctx, serverGroupId)
Refresh the server group associations

Used to immediately refresh automatic server associations for a server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverGroupId** | **int32**| Id of the serverGroup whose associations has to be refreshed | 

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServerGroupsCreationPreview**
> CreateServerGroupPreviewResp ServerGroupsCreationPreview(ctx, optional)
Server Group Preview

Get a preview of servers affected with create server group operation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServerGroupApiServerGroupsCreationPreviewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerGroupApiServerGroupsCreationPreviewOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AutomaticServerGroupAssociationDetails**](AutomaticServerGroupAssociationDetails.md)|  | 

### Return type

[**CreateServerGroupPreviewResp**](CreateServerGroupPreviewResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServerGroupAssociation**
> GenericResp UpdateServerGroupAssociation(ctx, serverGroupId, optional)
Update the server group

Used to update server associations for a server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverGroupId** | **int32**| Id of the serverGroupId whose details have to be fetched | 
 **optional** | ***ServerGroupApiUpdateServerGroupAssociationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerGroupApiUpdateServerGroupAssociationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateServerGroup**](UpdateServerGroup.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

