# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllowBackupOpOnHyperScaleXStoragePool**](HyperScaleStorageApi.md#AllowBackupOpOnHyperScaleXStoragePool) | **Put** /V4/Storage/HyperScale/{storagePoolId}/Extend | Allow Backup operation On HyperScale X StoragePool
[**CreateHyperScaleStorage**](HyperScaleStorageApi.md#CreateHyperScaleStorage) | **Post** /V4/Storage/HyperScale | Create HyperScaleStorage
[**CreateNodeforHyperScaleStorage**](HyperScaleStorageApi.md#CreateNodeforHyperScaleStorage) | **Post** /V4/Storage/HyperScale/{hyperScaleStorageId}/Nodes | Create Node for Hyper ScaleStorage
[**DeleteHyperScaleStorageById**](HyperScaleStorageApi.md#DeleteHyperScaleStorageById) | **Delete** /V4/Storage/HyperScale/{hyperScaleStorageId} | Delete HyperScaleStorage ById
[**GetHyperScaleStorageById**](HyperScaleStorageApi.md#GetHyperScaleStorageById) | **Get** /V4/Storage/HyperScale/{hyperScaleStorageId} | Get details of a specific HyperScale Storage
[**GetHyperScaleStorageList**](HyperScaleStorageApi.md#GetHyperScaleStorageList) | **Get** /V4/Storage/HyperScale | Get All HyperScale Storage
[**GetNodeDetailsOfHyperScaleStorage**](HyperScaleStorageApi.md#GetNodeDetailsOfHyperScaleStorage) | **Get** /V4/Storage/HyperScale/{hyperScaleStorageId}/Node/{nodeId} | Get NodeDetails Of HyperScaleStorage
[**ModifyHyperScaleStorageById**](HyperScaleStorageApi.md#ModifyHyperScaleStorageById) | **Put** /V4/Storage/HyperScale/{hyperScaleStorageId} | Modify HyperScaleStorage ById
[**RefreshNodeOfHyperScaleStorage**](HyperScaleStorageApi.md#RefreshNodeOfHyperScaleStorage) | **Put** /V4/Storage/HyperScale/{hyperScaleStorageId}/Node/{nodeId}/Refresh | Refresh Node Of HyperScaleStorage

# **AllowBackupOpOnHyperScaleXStoragePool**
> GenericResponse AllowBackupOpOnHyperScaleXStoragePool(ctx, storagePoolId)
Allow Backup operation On HyperScale X StoragePool

This operation is applicable only for HyperScale X Storage Pool. When the storage utilization is above 85%, it will be marked read only on next Tuesday. As long as utilization is below 90%, this API can be used to allow backup operation on the storage for an additional week.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the HyperScale X storage pool | 

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHyperScaleStorage**
> HyperScaleStorageResp CreateHyperScaleStorage(ctx, optional)
Create HyperScaleStorage

Create a HyperScale Storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HyperScaleStorageApiCreateHyperScaleStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HyperScaleStorageApiCreateHyperScaleStorageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of HyperScaleStorage**](HyperScaleStorage.md)|  | 

### Return type

[**HyperScaleStorageResp**](HyperScaleStorageResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNodeforHyperScaleStorage**
> GenericResp CreateNodeforHyperScaleStorage(ctx, hyperScaleStorageId, optional)
Create Node for Hyper ScaleStorage

Add nodes for a specific hyperscale storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hyperScaleStorageId** | **int32**| Id of hyperscale storage | 
 **optional** | ***HyperScaleStorageApiCreateNodeforHyperScaleStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HyperScaleStorageApiCreateNodeforHyperScaleStorageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of HyperScaleStorageIdNodesBody**](HyperScaleStorageIdNodesBody.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHyperScaleStorageById**
> GenericResp DeleteHyperScaleStorageById(ctx, hyperScaleStorageId)
Delete HyperScaleStorage ById

Delete existing hyperscale storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hyperScaleStorageId** | **int32**| Id of hyperscale storage | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHyperScaleStorageById**
> HyperScaleStorageDetails GetHyperScaleStorageById(ctx, hyperScaleStorageId, optional)
Get details of a specific HyperScale Storage

Get details of a specific HyperScale Storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hyperScaleStorageId** | **int32**| Id of hyperscale storage | 
 **optional** | ***HyperScaleStorageApiGetHyperScaleStorageByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HyperScaleStorageApiGetHyperScaleStorageByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showInheritedAssociation** | **optional.String**| Set to true if want to show inherited security associations | 

### Return type

[**HyperScaleStorageDetails**](HyperScaleStorageDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHyperScaleStorageList**
> InlineResponse2005 GetHyperScaleStorageList(ctx, )
Get All HyperScale Storage

Get All HyperScale Storage

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeDetailsOfHyperScaleStorage**
> HyperScaleNodeDetails GetNodeDetailsOfHyperScaleStorage(ctx, hyperScaleStorageId, nodeId)
Get NodeDetails Of HyperScaleStorage

Get details of a specific node of a hyperscale storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hyperScaleStorageId** | **int32**| Id of hyperscale storage | 
  **nodeId** | **int32**| Id of node | 

### Return type

[**HyperScaleNodeDetails**](HyperScaleNodeDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyHyperScaleStorageById**
> GenericResp ModifyHyperScaleStorageById(ctx, hyperScaleStorageId, optional)
Modify HyperScaleStorage ById

Modify details like name, encryption, security of a specific hyperscale storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hyperScaleStorageId** | **int32**| Id of hyperscale storage | 
 **optional** | ***HyperScaleStorageApiModifyHyperScaleStorageByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HyperScaleStorageApiModifyHyperScaleStorageByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateHyperScaleStorage**](UpdateHyperScaleStorage.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshNodeOfHyperScaleStorage**
> GenericResp RefreshNodeOfHyperScaleStorage(ctx, hyperScaleStorageId, nodeId)
Refresh Node Of HyperScaleStorage

Refresh a specific node of HyperScale

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hyperScaleStorageId** | **int32**| Id of hyperscale storage | 
  **nodeId** | **int32**| Id of node | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

