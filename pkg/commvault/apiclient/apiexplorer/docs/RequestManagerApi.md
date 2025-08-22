# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRequestManagerRequest**](RequestManagerApi.md#DeleteRequestManagerRequest) | **Delete** /V4/RequestManager/Request/{requestId} | 
[**GetProjectsForRequest**](RequestManagerApi.md#GetProjectsForRequest) | **Get** /V4/RequestManager/Request/{requestId}/Projects | 
[**GetRequestDetails**](RequestManagerApi.md#GetRequestDetails) | **Get** /V4/RequestManager/Request/{requestId} | Retrieve details of an existing request
[**RMConfigureRequestOperation**](RequestManagerApi.md#RMConfigureRequestOperation) | **Put** /V4/RequestManager/Request/{requestId}/Configure | RequestManager - Configure Request Operation
[**RMCreateRequestOperation**](RequestManagerApi.md#RMCreateRequestOperation) | **Post** /V4/RequestManager/Request | RequestManager - Create Request Operation
[**RequestManagerRequestList**](RequestManagerApi.md#RequestManagerRequestList) | **Get** /V4/RequestManager/Request | 

# **DeleteRequestManagerRequest**
> GenericResp DeleteRequestManagerRequest(ctx, requestId)


Deleting an existing request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **int32**| Unique identifier for the request | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectsForRequest**
> RmProjectsList GetProjectsForRequest(ctx, requestId)


Retrieve list of projects for the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **int32**| Unique request id | 

### Return type

[**RmProjectsList**](RMProjectsList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRequestDetails**
> RmRequestDetails GetRequestDetails(ctx, requestId)
Retrieve details of an existing request

Retrieve details of an existing request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **int32**| Unique identifier for the request | 

### Return type

[**RmRequestDetails**](RMRequestDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RMConfigureRequestOperation**
> GenericResp RMConfigureRequestOperation(ctx, requestId, optional)
RequestManager - Configure Request Operation

Configure a created request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **int32**|  | 
 **optional** | ***RequestManagerApiRMConfigureRequestOperationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RequestManagerApiRMConfigureRequestOperationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RmConfigureRequest**](RmConfigureRequest.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RMCreateRequestOperation**
> IdName RMCreateRequestOperation(ctx, optional)
RequestManager - Create Request Operation

Creating a request for request manager

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RequestManagerApiRMCreateRequestOperationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RequestManagerApiRMCreateRequestOperationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RmCreateRequest**](RmCreateRequest.md)|  | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestManagerRequestList**
> RmRequestList RequestManagerRequestList(ctx, optional)


Retrieves the list of Requests

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RequestManagerApiRequestManagerRequestListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RequestManagerApiRequestManagerRequestListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdFrom** | **optional.String**| Source application of the request | 
 **sourceEntityType** | **optional.String**| Entity type of the source from which data is gathered for the request | 
 **sourceEntityId** | **optional.Int32**| Entity id of the source from which data is gathered for the request | 

### Return type

[**RmRequestList**](RMRequestList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

