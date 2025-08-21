# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBulkLaptopOwnerMappingActionDelete**](LaptopOwnerMappingApi.md#DeleteBulkLaptopOwnerMappingActionDelete) | **Put** /V4/LaptopOwnerMapping/action/delete | delete multiple laptopowner mappings
[**GetLaptopOwnerMapping**](LaptopOwnerMappingApi.md#GetLaptopOwnerMapping) | **Get** /V4/LaptopOwnerMapping | Get LaptopOwner Mapping
[**GetV4LaptopOwnerMappingPreview**](LaptopOwnerMappingApi.md#GetV4LaptopOwnerMappingPreview) | **Get** /V4/LaptopOwnerMapping/Preview | 
[**PostLaptopOwnerMapping**](LaptopOwnerMappingApi.md#PostLaptopOwnerMapping) | **Post** /V4/LaptopOwnerMapping | Create laptop owner mappings
[**PutLaptopOwnerMapping**](LaptopOwnerMappingApi.md#PutLaptopOwnerMapping) | **Put** /V4/LaptopOwnerMapping | Update existing laptop owner mapping
[**PutLaptopOwnerMappingActionAssign**](LaptopOwnerMappingApi.md#PutLaptopOwnerMappingActionAssign) | **Put** /V4/LaptopOwnerMapping/action/Assign | Put LaptopOwner Mapping Action Assign

# **DeleteBulkLaptopOwnerMappingActionDelete**
> GenericResp DeleteBulkLaptopOwnerMappingActionDelete(ctx, optional)
delete multiple laptopowner mappings

API is used to delete multiple laptopowner mappings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LaptopOwnerMappingApiDeleteBulkLaptopOwnerMappingActionDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaptopOwnerMappingApiDeleteBulkLaptopOwnerMappingActionDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DeleteLaptopOwnerMapping**](DeleteLaptopOwnerMapping.md)| List of laptop owner mappings ids which needs to be deleted | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLaptopOwnerMapping**
> LaptopOwnerMapping GetLaptopOwnerMapping(ctx, optional)
Get LaptopOwner Mapping

API to get the list of laptop owner mappings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LaptopOwnerMappingApiGetLaptopOwnerMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaptopOwnerMappingApiGetLaptopOwnerMappingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **optional.Int32**| Id of the company for which the laptop owner mapping needs to be returned | 

### Return type

[**LaptopOwnerMapping**](LaptopOwnerMapping.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV4LaptopOwnerMappingPreview**
> LaptopOwnerMapping GetV4LaptopOwnerMappingPreview(ctx, )


Gives list of valid laptop owner mapping in the response

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LaptopOwnerMapping**](LaptopOwnerMapping.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostLaptopOwnerMapping**
> GenericResp PostLaptopOwnerMapping(ctx, optional)
Create laptop owner mappings

API to create laptop owner mappings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LaptopOwnerMappingApiPostLaptopOwnerMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaptopOwnerMappingApiPostLaptopOwnerMappingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateLaptopOwnerMapping**](CreateLaptopOwnerMapping.md)| List of laptop owner mappings to be added | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutLaptopOwnerMapping**
> GenericResp PutLaptopOwnerMapping(ctx, optional)
Update existing laptop owner mapping

API to update existing laptop owner mapping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LaptopOwnerMappingApiPutLaptopOwnerMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaptopOwnerMappingApiPutLaptopOwnerMappingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LaptopOwnerMapping**](LaptopOwnerMapping.md)| List of laptop owner mappings to update. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutLaptopOwnerMappingActionAssign**
> GenericResp PutLaptopOwnerMappingActionAssign(ctx, optional)
Put LaptopOwner Mapping Action Assign

API to assign users specified in device owner mapping as owners of laptops presented in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LaptopOwnerMappingApiPutLaptopOwnerMappingActionAssignOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaptopOwnerMappingApiPutLaptopOwnerMappingActionAssignOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AssignLaptopOwnerMappingReq**](AssignLaptopOwnerMappingReq.md)| Assign owners to laptops in the request. | 
 **companyId** | **optional.**| Id of the company that this operation will be performed for | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

