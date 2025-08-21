# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlackoutWindow**](BlackoutWindowsApi.md#CreateBlackoutWindow) | **Post** /V5/BlackoutWindow | Create a Blackout Window
[**CreateV4BlackoutWindow**](BlackoutWindowsApi.md#CreateV4BlackoutWindow) | **Post** /V4/BlackoutWindow | Create new Blackout Window
[**DeleteBlackoutWindow**](BlackoutWindowsApi.md#DeleteBlackoutWindow) | **Delete** /V4/BlackoutWindow/{blackoutWindowId} | Delete Blackout Window
[**GetBlackoutWindowDetails**](BlackoutWindowsApi.md#GetBlackoutWindowDetails) | **Get** /V5/BlackoutWindow/{blackoutWindowId} | Get Blackout Window Properties
[**GetBlackoutWindows**](BlackoutWindowsApi.md#GetBlackoutWindows) | **Get** /V5/BlackoutWindow | Get Blackout Window
[**GetV4BlackoutWindowDetails**](BlackoutWindowsApi.md#GetV4BlackoutWindowDetails) | **Get** /V4/BlackoutWindow/{blackoutWindowId} | Get Blackout Window Details
[**GetV4BlackoutWindows**](BlackoutWindowsApi.md#GetV4BlackoutWindows) | **Get** /V4/BlackoutWindow | Get Blackout Windows List
[**ModifyBlackoutWindow**](BlackoutWindowsApi.md#ModifyBlackoutWindow) | **Put** /V5/BlackoutWindow/{blackoutWindowId} | Modify Blackout Window properties
[**ModifyV4BlackoutWindow**](BlackoutWindowsApi.md#ModifyV4BlackoutWindow) | **Put** /V4/BlackoutWindow/{blackoutWindowId} | Modify Blackout Window Details

# **CreateBlackoutWindow**
> IdName CreateBlackoutWindow(ctx, optional)
Create a Blackout Window

Create a Blackout Window

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BlackoutWindowsApiCreateBlackoutWindowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlackoutWindowsApiCreateBlackoutWindowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateBlackoutWindow**](CreateBlackoutWindow.md)|  | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateV4BlackoutWindow**
> IdName CreateV4BlackoutWindow(ctx, optional)
Create new Blackout Window

Create a Blackout Window

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BlackoutWindowsApiCreateV4BlackoutWindowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlackoutWindowsApiCreateV4BlackoutWindowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of V4CreateBlackoutWindow**](V4CreateBlackoutWindow.md)|  | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBlackoutWindow**
> GenericResp DeleteBlackoutWindow(ctx, blackoutWindowId)
Delete Blackout Window

Used to delete a Blackout Window

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blackoutWindowId** | **int32**| Id of the Blackout Window to delete | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlackoutWindowDetails**
> BlackoutWindow GetBlackoutWindowDetails(ctx, blackoutWindowId)
Get Blackout Window Properties

Get details of a blackout window based on id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blackoutWindowId** | **int32**| Id of the Blackout Window whose details have to be fetched | 

### Return type

[**BlackoutWindow**](BlackoutWindow.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlackoutWindows**
> BlackoutWindowsListResponse GetBlackoutWindows(ctx, optional)
Get Blackout Window

Get All Blackout Windows

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BlackoutWindowsApiGetBlackoutWindowsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlackoutWindowsApiGetBlackoutWindowsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showOnlyCommcellLevel** | **optional.Bool**| Shows blackout window at commcell level if set to true. | 
 **companyId** | **optional.Int32**| Shows blackout windows associated with the company whose id has been provided. | 
 **serverGroupId** | **optional.String**| Shows blackout windows associated with the server group whose id has been provided. | 
 **commcellId** | **optional.String**| Shows blackout windows associated with the commcell whose id has been provided. | 

### Return type

[**BlackoutWindowsListResponse**](BlackoutWindowsListResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV4BlackoutWindowDetails**
> V4BlackoutWindow GetV4BlackoutWindowDetails(ctx, blackoutWindowId)
Get Blackout Window Details

Get details of a blackout window based on id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blackoutWindowId** | **int32**| Id of the Blackout Window whose details have to be fetched | 

### Return type

[**V4BlackoutWindow**](V4BlackoutWindow.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV4BlackoutWindows**
> V4BlackoutWindowsListResponse GetV4BlackoutWindows(ctx, optional)
Get Blackout Windows List

Get All Blackout Windows

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BlackoutWindowsApiGetV4BlackoutWindowsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlackoutWindowsApiGetV4BlackoutWindowsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showOnlyCommcellLevel** | **optional.Bool**| Shows blackout window at commcell level if set to true. | 
 **companyId** | **optional.Int32**| Shows blackout windows associated with the company whose id has been provided. | 
 **serverGroupId** | **optional.String**| Shows blackout windows associated with the server group whose id has been provided. | 
 **commcellId** | **optional.String**| Shows blackout windows associated with the commcell whose id has been provided. | 

### Return type

[**V4BlackoutWindowsListResponse**](V4BlackoutWindowsListResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyBlackoutWindow**
> GenericResp ModifyBlackoutWindow(ctx, blackoutWindowId, optional)
Modify Blackout Window properties

Modify the properties of an existing Blackout Window

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blackoutWindowId** | **int32**| Id of the Blackout Window to update | 
 **optional** | ***BlackoutWindowsApiModifyBlackoutWindowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlackoutWindowsApiModifyBlackoutWindowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateBlackoutWindow**](UpdateBlackoutWindow.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyV4BlackoutWindow**
> GenericResp ModifyV4BlackoutWindow(ctx, blackoutWindowId, optional)
Modify Blackout Window Details

Modify the properties of an existing Blackout Window

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blackoutWindowId** | **int32**| Id of the Blackout Window to update | 
 **optional** | ***BlackoutWindowsApiModifyV4BlackoutWindowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlackoutWindowsApiModifyV4BlackoutWindowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of V4UpdateBlackoutWindow**](V4UpdateBlackoutWindow.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

