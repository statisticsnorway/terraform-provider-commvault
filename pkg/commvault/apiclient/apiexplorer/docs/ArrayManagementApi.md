# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArray**](ArrayManagementApi.md#CreateArray) | **Post** /V4/StorageArrays | Add Storage Array
[**DeleteArray**](ArrayManagementApi.md#DeleteArray) | **Delete** /V4/StorageArrays/{arrayId} | Deleting a Storage Array
[**EditArray**](ArrayManagementApi.md#EditArray) | **Put** /V4/StorageArrays/{arrayId} | Edit Storage Array
[**GetArrayDetails**](ArrayManagementApi.md#GetArrayDetails) | **Get** /V4/StorageArrays/{arrayId} | Get Storage Array Details
[**GetArrays**](ArrayManagementApi.md#GetArrays) | **Get** /V4/StorageArrays | List All Storage Arrays
[**GetClientLvlArrays**](ArrayManagementApi.md#GetClientLvlArrays) | **Get** /V4/StorageArrays/Client/{clientId}/Arrays | List Storage Arrays of a Client
[**GetCopytLvlArrays**](ArrayManagementApi.md#GetCopytLvlArrays) | **Get** /V4/StorageArrays/Copy/{copyId}/Arrays | Get Storage Arrays for Copy
[**GetSubclientLvlArrays**](ArrayManagementApi.md#GetSubclientLvlArrays) | **Get** /V4/StorageArrays/Client/{clientId}/Subclient/{subclientId}/Arrays | List Storage Arrays of a Subclient

# **CreateArray**
> GenericResp CreateArray(ctx, optional)
Add Storage Array

Creating a Storage Array

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ArrayManagementApiCreateArrayOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArrayManagementApiCreateArrayOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateArray**](CreateArray.md)| Request to create an Array | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteArray**
> GenericResp DeleteArray(ctx, arrayId)
Deleting a Storage Array

Deleting a Storage Array

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditArray**
> GenericResp EditArray(ctx, arrayId, optional)
Edit Storage Array

Editing Array details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 
 **optional** | ***ArrayManagementApiEditArrayOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArrayManagementApiEditArrayOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EditArray**](EditArray.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArrayDetails**
> ArrayDetails GetArrayDetails(ctx, arrayId)
Get Storage Array Details

Getting Array Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 

### Return type

[**ArrayDetails**](ArrayDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArrays**
> GetArraysResp GetArrays(ctx, )
List All Storage Arrays

Get all storage arrays.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetArraysResp**](GetArraysResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientLvlArrays**
> GetArraysResp GetClientLvlArrays(ctx, clientId)
List Storage Arrays of a Client

API to get all arrays at client level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **int32**|  | 

### Return type

[**GetArraysResp**](GetArraysResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCopytLvlArrays**
> GetArraysResp GetCopytLvlArrays(ctx, copyId)
Get Storage Arrays for Copy

API to get arrays list at copy level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **copyId** | **int64**|  | 

### Return type

[**GetArraysResp**](GetArraysResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubclientLvlArrays**
> GetArraysResp GetSubclientLvlArrays(ctx, clientId, subclientId)
List Storage Arrays of a Subclient

API to get all arrays at Subclient Level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **int32**|  | 
  **subclientId** | **int32**|  | 

### Return type

[**GetArraysResp**](GetArraysResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

