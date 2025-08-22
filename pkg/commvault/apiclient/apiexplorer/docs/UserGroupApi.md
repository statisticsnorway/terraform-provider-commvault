# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserGroup**](UserGroupApi.md#CreateUserGroup) | **Post** /V4/UserGroup | Create User group
[**DeleteMultipleUserGroups**](UserGroupApi.md#DeleteMultipleUserGroups) | **Put** /V4/UserGroup/action/delete | Delete user groups
[**DeleteUserGroup**](UserGroupApi.md#DeleteUserGroup) | **Delete** /V4/UserGroup/{userGroupId} | Delete User Group
[**GetUserGroupDetails**](UserGroupApi.md#GetUserGroupDetails) | **Get** /V4/UserGroup/{userGroupId} | Get User Group Details
[**GetUserGroups**](UserGroupApi.md#GetUserGroups) | **Get** /V4/UserGroup | Get UserGroups
[**ModifyUserGroup**](UserGroupApi.md#ModifyUserGroup) | **Put** /V4/UserGroup/{userGroupId} | Update User Group Details

# **CreateUserGroup**
> IdNameGuid CreateUserGroup(ctx, optional)
Create User group

Create a new user-group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserGroupApiCreateUserGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserGroupApiCreateUserGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateUserGroup**](CreateUserGroup.md)|  | 

### Return type

[**IdNameGuid**](IdNameGUID.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMultipleUserGroups**
> GenericResp DeleteMultipleUserGroups(ctx, optional)
Delete user groups

Delete multiple userGroups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserGroupApiDeleteMultipleUserGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserGroupApiDeleteMultipleUserGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DeleteMultipleUserGroups**](DeleteMultipleUserGroups.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserGroup**
> GenericResp DeleteUserGroup(ctx, userGroupId, optional)
Delete User Group

Used to delete a user-group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userGroupId** | **int32**| Id of the user-group to delete | 
 **optional** | ***UserGroupApiDeleteUserGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserGroupApiDeleteUserGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipOwnerShipTransfer** | **optional.Bool**| If user group properties needn&#x27;t be transferred to other user or usergroup | 
 **transferToUser** | **optional.Int32**| If user group properties needs to be transferred to other user | 
 **transfertoUserGroup** | **optional.Int32**| If user group properties needs to be transferred to other userGroup | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserGroupDetails**
> UserGroup GetUserGroupDetails(ctx, userGroupId, optional)
Get User Group Details

Get details of a user-group based on id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userGroupId** | **int32**| Id of the user-group whose details have to be fetched | 
 **optional** | ***UserGroupApiGetUserGroupDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserGroupApiGetUserGroupDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | [**optional.Interface of Mode**](.md)|  | 

### Return type

[**UserGroup**](UserGroup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserGroups**
> UserGroupListResponse GetUserGroups(ctx, optional)
Get UserGroups

Get a list of existing user groups.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserGroupApiGetUserGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserGroupApiGetUserGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mode** | [**optional.Interface of Mode**](.md)|  | 

### Return type

[**UserGroupListResponse**](UserGroupListResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyUserGroup**
> GenericResp ModifyUserGroup(ctx, userGroupId, optional)
Update User Group Details

Modify the properties of an existing user-group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userGroupId** | **int32**| Id of the user-group to update | 
 **optional** | ***UserGroupApiModifyUserGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserGroupApiModifyUserGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateUserGroup**](UpdateUserGroup.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

