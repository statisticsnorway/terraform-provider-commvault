# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTeamsContent**](TeamsApi.md#AddTeamsContent) | **Put** /V4/Office365/Teams/{appId}/Content | Add content to Office 365 Teams app
[**AddTeamsCustomCategory**](TeamsApi.md#AddTeamsCustomCategory) | **Put** /V4/Office365/Teams/{appId}/Content/CustomCategory | Add Office 365 Teams Custom Category
[**AddTeamsList**](TeamsApi.md#AddTeamsList) | **Put** /V4/Office365/Teams/{appId}/Teams | Add Office 365 Teams List
[**AddTeamsUsersList**](TeamsApi.md#AddTeamsUsersList) | **Put** /V4/Office365/Teams/{appId}/Users | Add Office 365 Teams Users List
[**CreateTeamsClient**](TeamsApi.md#CreateTeamsClient) | **Post** /V4/Office365/Teams | Create Office 365 Teams app
[**GetTeamsClient**](TeamsApi.md#GetTeamsClient) | **Get** /V4/Office365/Teams | Get Office 365 Teams apps
[**GetTeamsContent**](TeamsApi.md#GetTeamsContent) | **Get** /V4/Office365/Teams/{appId}/Content | Get Office 365 Teams Content
[**GetTeamsList**](TeamsApi.md#GetTeamsList) | **Get** /V4/Office365/Teams/{appId}/Teams | Get Office 365 Teams List
[**GetTeamsUsersList**](TeamsApi.md#GetTeamsUsersList) | **Get** /V4/Office365/Teams/{appId}/Users | Get Office 365 Teams Users List
[**PerformTeamsBackup**](TeamsApi.md#PerformTeamsBackup) | **Post** /V4/Office365/Teams/{appId}/Backup | Perform Office 365 Teams Backup
[**PerformTeamsRestore**](TeamsApi.md#PerformTeamsRestore) | **Post** /V4/Office365/Teams/{appId}/Restore | Perform Office 365 Teams Restore

# **AddTeamsContent**
> GenericResp AddTeamsContent(ctx, appId, optional)
Add content to Office 365 Teams app

Add content to Office 365 Teams app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault teams app id | 
 **optional** | ***TeamsApiAddTeamsContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiAddTeamsContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TeamsContentReq**](TeamsContentReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTeamsCustomCategory**
> GenericResp AddTeamsCustomCategory(ctx, appId, optional)
Add Office 365 Teams Custom Category

Add Office 365 Teams Custom Category

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault Teams app id | 
 **optional** | ***TeamsApiAddTeamsCustomCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiAddTeamsCustomCategoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Office365TeamsCustomCategoryReq**](Office365TeamsCustomCategoryReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTeamsList**
> GenericResp AddTeamsList(ctx, appId, optional)
Add Office 365 Teams List

Add Office 365 Teams List

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault teams app id | 
 **optional** | ***TeamsApiAddTeamsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiAddTeamsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Office365TeamsReq**](Office365TeamsReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTeamsUsersList**
> GenericResp AddTeamsUsersList(ctx, appId, optional)
Add Office 365 Teams Users List

Add Office 365 Teams Users List

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault teams app id | 
 **optional** | ***TeamsApiAddTeamsUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiAddTeamsUsersListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Office365TeamsUsersReq**](Office365TeamsUsersReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTeamsClient**
> IdName CreateTeamsClient(ctx, optional)
Create Office 365 Teams app

Create Office 365 Teams app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamsApiCreateTeamsClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiCreateTeamsClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateTeamsClient**](CreateTeamsClient.md)|  | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamsClient**
> Office365ClientsResp GetTeamsClient(ctx, optional)
Get Office 365 Teams apps

Get list of Office 365 Teams apps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamsApiGetTeamsClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiGetTeamsClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 
 **searchKey** | **optional.String**|  | 

### Return type

[**Office365ClientsResp**](Office365ClientsResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamsContent**
> Office365TeamsContentResp GetTeamsContent(ctx, appId, optional)
Get Office 365 Teams Content

Get Office 365 Teams Content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault teams app id | 
 **optional** | ***TeamsApiGetTeamsContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiGetTeamsContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 
 **searchKey** | **optional.String**|  | 

### Return type

[**Office365TeamsContentResp**](Office365TeamsContentResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamsList**
> Office365TeamsResp GetTeamsList(ctx, appId, optional)
Get Office 365 Teams List

Get Office 365 Teams List

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault teams app id | 
 **optional** | ***TeamsApiGetTeamsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiGetTeamsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **propertyLevel** | **optional.String**| Optional parameter. Set it to BasicProperties or AllProperties | 
 **pageNumber** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 
 **searchKey** | **optional.String**|  | 

### Return type

[**Office365TeamsResp**](Office365TeamsResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamsUsersList**
> Office365TeamsUsersResp GetTeamsUsersList(ctx, appId, optional)
Get Office 365 Teams Users List

Get Office 365 Teams Users List

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault teams app id | 
 **optional** | ***TeamsApiGetTeamsUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiGetTeamsUsersListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **propertyLevel** | **optional.String**| Optional parameter. Set it to BasicProperties or AllProperties | 
 **pageNumber** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 
 **searchKey** | **optional.String**|  | 

### Return type

[**Office365TeamsUsersResp**](Office365TeamsUsersResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PerformTeamsBackup**
> CreateTaskRespforBackup PerformTeamsBackup(ctx, appId, optional)
Perform Office 365 Teams Backup

Perform Office 365 Teams Backup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault Teams app id | 
 **optional** | ***TeamsApiPerformTeamsBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiPerformTeamsBackupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Office365TeamsBackupReq**](Office365TeamsBackupReq.md)|  | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PerformTeamsRestore**
> CreateTaskRespforBackup PerformTeamsRestore(ctx, appId, optional)
Perform Office 365 Teams Restore

Perform Office 365 Teams Restore

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault Teams app id | 
 **optional** | ***TeamsApiPerformTeamsRestoreOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiPerformTeamsRestoreOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Office365TeamsRestoreReq**](Office365TeamsRestoreReq.md)|  | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

