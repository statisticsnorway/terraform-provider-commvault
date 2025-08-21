# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSharepointContent**](SharepointApi.md#AddSharepointContent) | **Put** /V4/Office365/SharePoint/{appId}/Content | Add content to Office 365 SharePoint Online app
[**AddSharepointCustomCategory**](SharepointApi.md#AddSharepointCustomCategory) | **Put** /V4/Office365/SharePoint/{appId}/Content/CustomCategory | Add Office 365 Sharepoint Custom Category
[**AddSharepointSitesList**](SharepointApi.md#AddSharepointSitesList) | **Put** /V4/Office365/Sharepoint/{appId}/Sites | Add Office 365 Sharepoint sites
[**CreateSharepointClient**](SharepointApi.md#CreateSharepointClient) | **Post** /V4/Office365/SharePoint | Create Office 365 SharePoint Online
[**GetSharepointClient**](SharepointApi.md#GetSharepointClient) | **Get** /V4/Office365/SharePoint | Get Office 365 Sharepoint Online apps
[**GetSharepointContent**](SharepointApi.md#GetSharepointContent) | **Get** /V4/Office365/SharePoint/{appId}/Content | Get Office 365 Sharepoint Online Content
[**GetSharepointSites**](SharepointApi.md#GetSharepointSites) | **Get** /V4/Office365/Sharepoint/{appId}/Sites | Get Office 365 Sharepoint Online Sites
[**PerformSharepointBackup**](SharepointApi.md#PerformSharepointBackup) | **Post** /V4/Office365/Sharepoint/{appId}/Backup | Perform Office 365 Sharepoint Backup
[**PerformSharepointRestore**](SharepointApi.md#PerformSharepointRestore) | **Post** /V4/Office365/Sharepoint/{appId}/Restore | Perform Office 365 Sharepoint Restore

# **AddSharepointContent**
> GenericResp AddSharepointContent(ctx, appId, optional)
Add content to Office 365 SharePoint Online app

Add content to Office 365 SharePoint Online app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault sharepoint app id | 
 **optional** | ***SharepointApiAddSharepointContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharepointApiAddSharepointContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SharepointContentReq**](SharepointContentReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSharepointCustomCategory**
> GenericResp AddSharepointCustomCategory(ctx, appId, optional)
Add Office 365 Sharepoint Custom Category

Add Office 365 Sharepoint Custom Category

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault Sharepoint app id | 
 **optional** | ***SharepointApiAddSharepointCustomCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharepointApiAddSharepointCustomCategoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Office365SharepointCustomCategoryReq**](Office365SharepointCustomCategoryReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSharepointSitesList**
> GenericResp AddSharepointSitesList(ctx, appId, optional)
Add Office 365 Sharepoint sites

Add sites to Office 365 SharePoint Online app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault sharepoint app id | 
 **optional** | ***SharepointApiAddSharepointSitesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharepointApiAddSharepointSitesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Office365SitesReq**](Office365SitesReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSharepointClient**
> IdName CreateSharepointClient(ctx, optional)
Create Office 365 SharePoint Online

Create Office 365 SharePoint Online app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SharepointApiCreateSharepointClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharepointApiCreateSharepointClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateSharepointClient**](CreateSharepointClient.md)|  | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSharepointClient**
> Office365ClientsResp GetSharepointClient(ctx, optional)
Get Office 365 Sharepoint Online apps

Get list of Office 365 Sharepoint Online apps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SharepointApiGetSharepointClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharepointApiGetSharepointClientOpts struct
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

# **GetSharepointContent**
> Office365SharepointContentResp GetSharepointContent(ctx, appId, optional)
Get Office 365 Sharepoint Online Content

Get Office 365 Sharepoint Online Content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault sharepoint app id | 
 **optional** | ***SharepointApiGetSharepointContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharepointApiGetSharepointContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 
 **searchKey** | **optional.String**|  | 

### Return type

[**Office365SharepointContentResp**](Office365SharepointContentResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSharepointSites**
> Office365SitesResp GetSharepointSites(ctx, appId, optional)
Get Office 365 Sharepoint Online Sites

Get Office 365 Sharepoint Online Sites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault sharepoint app id | 
 **optional** | ***SharepointApiGetSharepointSitesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharepointApiGetSharepointSitesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **propertyLevel** | **optional.String**| Optional parameter. Set it to BasicProperties or AllProperties | 
 **pageNumber** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 
 **searchKey** | **optional.String**|  | 

### Return type

[**Office365SitesResp**](Office365SitesResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PerformSharepointBackup**
> CreateTaskRespforBackup PerformSharepointBackup(ctx, appId, optional)
Perform Office 365 Sharepoint Backup

Perform Office 365 Sharepoint Backup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault Sharepoint app id | 
 **optional** | ***SharepointApiPerformSharepointBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharepointApiPerformSharepointBackupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Office365SharepointBackupReq**](Office365SharepointBackupReq.md)|  | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PerformSharepointRestore**
> CreateTaskRespforBackup PerformSharepointRestore(ctx, appId, optional)
Perform Office 365 Sharepoint Restore

Perform Office 365 Sharepoint Restore

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault Sharepoint app id | 
 **optional** | ***SharepointApiPerformSharepointRestoreOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharepointApiPerformSharepointRestoreOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Office365SharepointRestoreReq**](Office365SharepointRestoreReq.md)|  | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

