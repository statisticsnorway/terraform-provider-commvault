# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceMeshClient**](ServiceMeshApi.md#CreateServiceMeshClient) | **Post** /V4/ServiceMesh | Create Service Mesh
[**DeleteServiceMeshClient**](ServiceMeshApi.md#DeleteServiceMeshClient) | **Delete** /V4/ServiceMesh/{clientId} | Delete Service Mesh Client
[**GetServiceMeshDetails**](ServiceMeshApi.md#GetServiceMeshDetails) | **Get** /V4/ServiceMesh | Get Service Mesh details
[**UpdateServiceMeshClientProperties**](ServiceMeshApi.md#UpdateServiceMeshClientProperties) | **Put** /V4/ServiceMesh/{clientId} | Update properties of Service mesh client

# **CreateServiceMeshClient**
> IdName CreateServiceMeshClient(ctx, optional)
Create Service Mesh

Create a Service mesh client representation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServiceMeshApiCreateServiceMeshClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceMeshApiCreateServiceMeshClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateServiceMeshReq**](CreateServiceMeshReq.md)|  | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceMeshClient**
> GenericResp DeleteServiceMeshClient(ctx, )
Delete Service Mesh Client

Delete Service Mesh client representation from the commcell

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceMeshDetails**
> GetServiceMeshDetailsResp GetServiceMeshDetails(ctx, )
Get Service Mesh details

Gets the Service mesh and it's related properties configured for this Commcell

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetServiceMeshDetailsResp**](GetServiceMeshDetailsResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceMeshClientProperties**
> GenericResp UpdateServiceMeshClientProperties(ctx, optional)
Update properties of Service mesh client

Update properties of a Service mesh client representation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServiceMeshApiUpdateServiceMeshClientPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceMeshApiUpdateServiceMeshClientPropertiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UpdateServiceMeshReq**](UpdateServiceMeshReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

