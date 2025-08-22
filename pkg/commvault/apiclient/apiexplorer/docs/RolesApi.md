# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewRole**](RolesApi.md#CreateNewRole) | **Post** /V4/Role | Create a Role
[**DeleteRoles**](RolesApi.md#DeleteRoles) | **Delete** /V4/Role/{roleId} | Delete a role
[**GetListOfRoles**](RolesApi.md#GetListOfRoles) | **Get** /V4/Role | Get a list of all the roles
[**GetPermissionResponse**](RolesApi.md#GetPermissionResponse) | **Get** /V4/Permissions | Get Permissions
[**GetRoleDetails**](RolesApi.md#GetRoleDetails) | **Get** /V4/Role/{roleId} | Get a details of the role whose role id has been provided
[**ModifyRole**](RolesApi.md#ModifyRole) | **Put** /V4/Role/{roleId} | Update the Role

# **CreateNewRole**
> IdNameGuid CreateNewRole(ctx, optional)
Create a Role

Create a new role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RolesApiCreateNewRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiCreateNewRoleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateRole**](CreateRole.md)|  | 

### Return type

[**IdNameGuid**](IdNameGUID.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRoles**
> GenericResp DeleteRoles(ctx, roleId)
Delete a role

Used to delete a role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **int32**| Role Id | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListOfRoles**
> RoleListResponse GetListOfRoles(ctx, )
Get a list of all the roles

Get a list of all the roles

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RoleListResponse**](RoleListResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPermissionResponse**
> PermissionListResp GetPermissionResponse(ctx, )
Get Permissions

Get a list of categories and permissions in each category

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PermissionListResp**](PermissionListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleDetails**
> Role GetRoleDetails(ctx, roleId, optional)
Get a details of the role whose role id has been provided

Get a details of the role whose role id has been provided

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **int32**| Role Id | 
 **optional** | ***RolesApiGetRoleDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiGetRoleDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showInheritedAssociation** | **optional.Bool**| Shows inherited security associations | 

### Return type

[**Role**](Role.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyRole**
> GenericResp ModifyRole(ctx, roleId, optional)
Update the Role

Modify the properties of an existing role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **int32**| Role Id | 
 **optional** | ***RolesApiModifyRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiModifyRoleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateRole**](UpdateRole.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

