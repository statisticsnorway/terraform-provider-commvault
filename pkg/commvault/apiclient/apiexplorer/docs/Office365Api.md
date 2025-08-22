# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOffice365AzureApp**](Office365Api.md#AddOffice365AzureApp) | **Put** /V4/Office365/{appId}/AzureApp | Add/Update Office365 Azure Apps
[**DeleteOffice365AzureApp**](Office365Api.md#DeleteOffice365AzureApp) | **Delete** /V4/Office365/{appId}/AzureApp | Delete Office365 Azure Apps
[**GetOffice365AzureApp**](Office365Api.md#GetOffice365AzureApp) | **Get** /V4/Office365/{appId}/AzureApp | Get Office 365 Azure apps
[**GetOffice365Client**](Office365Api.md#GetOffice365Client) | **Get** /V4/Office365 | Get Office 365 apps
[**GetOffice365PlanDashboard**](Office365Api.md#GetOffice365PlanDashboard) | **Get** /V4/Office365/Plans | Get Office365 Plans for Dashboard

# **AddOffice365AzureApp**
> UpdateOffice365AzureAppResp AddOffice365AzureApp(ctx, appId, optional)
Add/Update Office365 Azure Apps

Add list of Office365 Azure Apps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**|  | 
 **optional** | ***Office365ApiAddOffice365AzureAppOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Office365ApiAddOffice365AzureAppOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateOffice365AzureAppReq**](UpdateOffice365AzureAppReq.md)|  | 

### Return type

[**UpdateOffice365AzureAppResp**](UpdateOffice365AzureAppResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOffice365AzureApp**
> UpdateOffice365AzureAppResp DeleteOffice365AzureApp(ctx, appId, optional)
Delete Office365 Azure Apps

Delete list of Office365 Azure Apps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**|  | 
 **optional** | ***Office365ApiDeleteOffice365AzureAppOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Office365ApiDeleteOffice365AzureAppOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DeleteOffice365AzureAppReq**](DeleteOffice365AzureAppReq.md)|  | 

### Return type

[**UpdateOffice365AzureAppResp**](UpdateOffice365AzureAppResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOffice365AzureApp**
> GetOffice365AzureAppResp GetOffice365AzureApp(ctx, appId)
Get Office 365 Azure apps

Get list of Office 365 Azure apps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**|  | 

### Return type

[**GetOffice365AzureAppResp**](GetOffice365AzureAppResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOffice365Client**
> Office365AllClientsResp GetOffice365Client(ctx, optional)
Get Office 365 apps

Get list of Office 365 apps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Office365ApiGetOffice365ClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Office365ApiGetOffice365ClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 
 **searchKey** | **optional.String**|  | 

### Return type

[**Office365AllClientsResp**](Office365AllClientsResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOffice365PlanDashboard**
> Office365PlanSummaryListResp GetOffice365PlanDashboard(ctx, )
Get Office365 Plans for Dashboard

Returns all Office365 active plans with number of entities associated to the plan per workload

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Office365PlanSummaryListResp**](Office365PlanSummaryListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

