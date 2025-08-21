# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOneDriveCustomCategory**](OneDriveApi.md#AddOneDriveCustomCategory) | **Put** /V4/Office365/OneDrive/{appId}/Content/CustomCategory | Add Office 365 OneDrive Custom Category
[**AddOnedriveContent**](OneDriveApi.md#AddOnedriveContent) | **Put** /V4/Office365/OneDrive/{appId}/Content | Add content to Office 365 OneDrive for Business app
[**AddOnedriveUsersList**](OneDriveApi.md#AddOnedriveUsersList) | **Put** /V4/Office365/Onedrive/{appId}/Users | Add Office 365 Onedrive Users List
[**CreateOnedriveClient**](OneDriveApi.md#CreateOnedriveClient) | **Post** /V4/Office365/OneDrive | Create Office 365 Onedrive for Business app
[**GetOnedriveClient**](OneDriveApi.md#GetOnedriveClient) | **Get** /V4/Office365/OneDrive | Get Office 365 Onedrive for Business apps
[**GetOnedriveContent**](OneDriveApi.md#GetOnedriveContent) | **Get** /V4/Office365/OneDrive/{appId}/Content | Get Office 365 Onedrive for Business Content
[**GetOnedriveUsers**](OneDriveApi.md#GetOnedriveUsers) | **Get** /V4/Office365/Onedrive/{appId}/Users | Get Office 365 Onedrive for Business Users
[**PerformOnedriveBackup**](OneDriveApi.md#PerformOnedriveBackup) | **Post** /V4/Office365/OneDrive/{appId}/Backup | Perform Office 365 Onedrive Backup
[**PerformOnedriveRestore**](OneDriveApi.md#PerformOnedriveRestore) | **Post** /V4/Office365/OneDrive/{appId}/Restore | Perform Office 365 Onedrive Restore

# **AddOneDriveCustomCategory**
> GenericResp AddOneDriveCustomCategory(ctx, appId, optional)
Add Office 365 OneDrive Custom Category

Add Office 365 OneDrive Custom Category

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault OneDrive app id | 
 **optional** | ***OneDriveApiAddOneDriveCustomCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OneDriveApiAddOneDriveCustomCategoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Office365OnedriveCustomCategoryReq**](Office365OnedriveCustomCategoryReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddOnedriveContent**
> GenericResp AddOnedriveContent(ctx, appId, optional)
Add content to Office 365 OneDrive for Business app

Add content to Office 365 OneDrive for Business app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault onedrive app id | 
 **optional** | ***OneDriveApiAddOnedriveContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OneDriveApiAddOnedriveContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OnedriveContentReq**](OnedriveContentReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddOnedriveUsersList**
> GenericResp AddOnedriveUsersList(ctx, appId, optional)
Add Office 365 Onedrive Users List

Add Office 365 Onedrive Users List

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault onedrive app id | 
 **optional** | ***OneDriveApiAddOnedriveUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OneDriveApiAddOnedriveUsersListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Office365OnedriveUsersReq**](Office365OnedriveUsersReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOnedriveClient**
> IdName CreateOnedriveClient(ctx, optional)
Create Office 365 Onedrive for Business app

Create Office 365 Onedrive for Business app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OneDriveApiCreateOnedriveClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OneDriveApiCreateOnedriveClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateOnedriveClient**](CreateOnedriveClient.md)|  | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOnedriveClient**
> Office365ClientsResp GetOnedriveClient(ctx, optional)
Get Office 365 Onedrive for Business apps

Get list of Office 365 Onedrive for Business apps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OneDriveApiGetOnedriveClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OneDriveApiGetOnedriveClientOpts struct
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

# **GetOnedriveContent**
> Office365OnedriveContentResp GetOnedriveContent(ctx, appId, optional)
Get Office 365 Onedrive for Business Content

Get Office 365 Onedrive for Business Content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault onedrive app id | 
 **optional** | ***OneDriveApiGetOnedriveContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OneDriveApiGetOnedriveContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 
 **searchKey** | **optional.String**|  | 

### Return type

[**Office365OnedriveContentResp**](Office365OnedriveContentResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOnedriveUsers**
> Office365OnedriveUsersResp GetOnedriveUsers(ctx, appId, optional)
Get Office 365 Onedrive for Business Users

Get Office 365 Onedrive for Business Users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault onedrive app id | 
 **optional** | ***OneDriveApiGetOnedriveUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OneDriveApiGetOnedriveUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **propertyLevel** | **optional.String**| Optional parameter. Set it to BasicProperties or AllProperties | 
 **pageNumber** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 
 **searchKey** | **optional.String**|  | 

### Return type

[**Office365OnedriveUsersResp**](Office365OnedriveUsersResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PerformOnedriveBackup**
> CreateTaskRespforBackup PerformOnedriveBackup(ctx, appId, optional)
Perform Office 365 Onedrive Backup

Perform Office 365 Onedrive Backup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault Onedrive app id | 
 **optional** | ***OneDriveApiPerformOnedriveBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OneDriveApiPerformOnedriveBackupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Office365OnedriveBackupReq**](Office365OnedriveBackupReq.md)|  | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PerformOnedriveRestore**
> CreateTaskRespforBackup PerformOnedriveRestore(ctx, appId, optional)
Perform Office 365 Onedrive Restore

Perform Office 365 Onedrive Restore

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault Onedrive app id | 
 **optional** | ***OneDriveApiPerformOnedriveRestoreOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OneDriveApiPerformOnedriveRestoreOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Office365OnedriveRestoreReq**](Office365OnedriveRestoreReq.md)|  | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

