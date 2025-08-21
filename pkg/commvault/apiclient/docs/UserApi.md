# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UserApi.md#CreateUser) | **Post** /V4/user | 
[**DeleteMultipleUsers**](UserApi.md#DeleteMultipleUsers) | **Put** /V4/user/action/delete | Delete Multiple Users
[**DeleteUser**](UserApi.md#DeleteUser) | **Delete** /V4/user/{userId} | Delete an User
[**GetUserDetails**](UserApi.md#GetUserDetails) | **Get** /V4/user/{userId} | Get User Details
[**GetUsers**](UserApi.md#GetUsers) | **Get** /V4/user | Get Users
[**ModifyUser**](UserApi.md#ModifyUser) | **Put** /V4/user/{userId} | Update existing user

# **CreateUser**
> CreateUserResp CreateUser(ctx, optional)


Create a User

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiCreateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiCreateUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateUsers**](CreateUsers.md)|  | 

### Return type

[**CreateUserResp**](CreateUserResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMultipleUsers**
> GenericResp DeleteMultipleUsers(ctx, optional)
Delete Multiple Users

Delete multiple users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiDeleteMultipleUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiDeleteMultipleUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DeleteMultipleUsers**](DeleteMultipleUsers.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> GenericResp DeleteUser(ctx, userId, optional)
Delete an User

Used to delete a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| Id of the user to delete | 
 **optional** | ***UserApiDeleteUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiDeleteUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipOwnerShipTransfer** | **optional.Bool**| If user properties needn&#x27;t be transferred to other user or usergroup | 
 **transferToUser** | **optional.Int32**| If user properties needs to be transferred to other user | 
 **transfertoUserGroup** | **optional.Int32**| If user properties needs to be transferred to other userGroup | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserDetails**
> User GetUserDetails(ctx, userId, optional)
Get User Details

Get details of a User based on id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| Id of the User whose details have to be fetched | 
 **optional** | ***UserApiGetUserDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetUserDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | [**optional.Interface of Mode**](.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers**
> UserListResp GetUsers(ctx, optional)
Get Users

Get All Users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiGetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **laptopUsers** | **optional.Bool**| Returns only laptop users when set to true. This param is only effective when EdgeMode header is passed. | 
 **userGroupId** | **optional.Int32**| Returns the list of users associated to the userGroupId provided. | 

### Return type

[**UserListResp**](UserListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyUser**
> GenericResp ModifyUser(ctx, userId, optional)
Update existing user

Used to modify an existing user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| Id of the User to update | 
 **optional** | ***UserApiModifyUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiModifyUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateUser**](UpdateUser.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

