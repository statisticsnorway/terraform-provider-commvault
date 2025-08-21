# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddExchangeContent**](ExchangeApi.md#AddExchangeContent) | **Put** /V4/Office365/Exchange/{appId}/Content | Add Content to Office 365 Exchange Online app
[**AddExchangeCustomCategory**](ExchangeApi.md#AddExchangeCustomCategory) | **Put** /V4/Office365/Exchange/{appId}/Content/CustomCategory | Add Office 365 Exchange Custom Category
[**AddExchangeMailboxesList**](ExchangeApi.md#AddExchangeMailboxesList) | **Put** /V4/Office365/Exchange/{appId}/Mailboxes | Add Office 365 Exchange Online Mailboxes List
[**CreateExchangeClient**](ExchangeApi.md#CreateExchangeClient) | **Post** /V4/Office365/Exchange | Create Office 365 Exchange Online app
[**GetExchangeClient**](ExchangeApi.md#GetExchangeClient) | **Get** /V4/Office365/Exchange | Get Office 365 Exchange Online apps
[**GetExchangeContent**](ExchangeApi.md#GetExchangeContent) | **Get** /V4/Office365/Exchange/{appId}/Content | Get Office 365 Exchange Online Content
[**GetExchangeMailboxes**](ExchangeApi.md#GetExchangeMailboxes) | **Get** /V4/Office365/Exchange/{appId}/Mailboxes | Get Office 365 Exchange Online Mailboxes
[**PerformExchangeBackup**](ExchangeApi.md#PerformExchangeBackup) | **Post** /V4/Office365/Exchange/{appId}/Backup | Perform Office 365 Exchange Backup
[**PerformExchangeRestore**](ExchangeApi.md#PerformExchangeRestore) | **Post** /V4/Office365/Exchange/{appId}/Restore | Perform Office 365 Exchange Restore

# **AddExchangeContent**
> GenericResp AddExchangeContent(ctx, appId, optional)
Add Content to Office 365 Exchange Online app

Add Content to Office 365 Exchange Online app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault exchange app id | 
 **optional** | ***ExchangeApiAddExchangeContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExchangeApiAddExchangeContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ExchangeContentReq**](ExchangeContentReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddExchangeCustomCategory**
> GenericResp AddExchangeCustomCategory(ctx, appId, optional)
Add Office 365 Exchange Custom Category

Add Office 365 Exchange Custom Category

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault Exchange app id | 
 **optional** | ***ExchangeApiAddExchangeCustomCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExchangeApiAddExchangeCustomCategoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Office365ExchangeCustomCategoryReq**](Office365ExchangeCustomCategoryReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddExchangeMailboxesList**
> GenericResp AddExchangeMailboxesList(ctx, appId, optional)
Add Office 365 Exchange Online Mailboxes List

Add Office 365 Exchange Online Mailboxes List

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault exchange app id | 
 **optional** | ***ExchangeApiAddExchangeMailboxesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExchangeApiAddExchangeMailboxesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Office365MailboxesReq**](Office365MailboxesReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateExchangeClient**
> IdName CreateExchangeClient(ctx, optional)
Create Office 365 Exchange Online app

Create Office 365 Exchange Online app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExchangeApiCreateExchangeClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExchangeApiCreateExchangeClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateExchangeClient**](CreateExchangeClient.md)|  | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExchangeClient**
> Office365ClientsResp GetExchangeClient(ctx, optional)
Get Office 365 Exchange Online apps

Get list of Office 365 Exchange Online apps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExchangeApiGetExchangeClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExchangeApiGetExchangeClientOpts struct
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

# **GetExchangeContent**
> Office365ExchangeContentResp GetExchangeContent(ctx, appId, optional)
Get Office 365 Exchange Online Content

Get Office 365 Exchange Online Content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault exchange app id | 
 **optional** | ***ExchangeApiGetExchangeContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExchangeApiGetExchangeContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 
 **searchKey** | **optional.String**|  | 

### Return type

[**Office365ExchangeContentResp**](Office365ExchangeContentResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExchangeMailboxes**
> Office365MailboxesResp GetExchangeMailboxes(ctx, appId, optional)
Get Office 365 Exchange Online Mailboxes

Get Office 365 Exchange Online Mailboxes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault exchange app id | 
 **optional** | ***ExchangeApiGetExchangeMailboxesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExchangeApiGetExchangeMailboxesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **propertyLevel** | **optional.String**| Optional parameter. Set it to BasicProperties or AllProperties | 
 **pageNumber** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 
 **searchKey** | **optional.String**|  | 

### Return type

[**Office365MailboxesResp**](Office365MailboxesResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PerformExchangeBackup**
> CreateTaskRespforBackup PerformExchangeBackup(ctx, appId, optional)
Perform Office 365 Exchange Backup

Perform Office 365 Exchange Backup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault Exchange app id | 
 **optional** | ***ExchangeApiPerformExchangeBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExchangeApiPerformExchangeBackupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Office365ExchangeBackupReq**](Office365ExchangeBackupReq.md)|  | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PerformExchangeRestore**
> CreateTaskRespforBackup PerformExchangeRestore(ctx, appId, optional)
Perform Office 365 Exchange Restore

Perform Office 365 Exchange Restore

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault Exchange app id | 
 **optional** | ***ExchangeApiPerformExchangeRestoreOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExchangeApiPerformExchangeRestoreOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Office365ExchangeRestoreReq**](Office365ExchangeRestoreReq.md)|  | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

