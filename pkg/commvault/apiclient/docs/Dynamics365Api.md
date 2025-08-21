# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddD365Content**](Dynamics365Api.md#AddD365Content) | **Put** /V4/Dynamics365/{appId}/Content | Add content to Dynamics 365 app
[**CreateD365Client**](Dynamics365Api.md#CreateD365Client) | **Post** /v4/Dynamics365 | Create D365 Client

# **AddD365Content**
> GenericResp AddD365Content(ctx, appId, optional)
Add content to Dynamics 365 app

Add content to Dynamics 365 app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| Commvault dynamics 365 app id | 
 **optional** | ***Dynamics365ApiAddD365ContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Dynamics365ApiAddD365ContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Dynamics365ContentReq**](Dynamics365ContentReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateD365Client**
> IdName CreateD365Client(ctx, optional)
Create D365 Client

Create Dynamics 365 app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Dynamics365ApiCreateD365ClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Dynamics365ApiCreateD365ClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateDynamics365Client**](CreateDynamics365Client.md)|  | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

