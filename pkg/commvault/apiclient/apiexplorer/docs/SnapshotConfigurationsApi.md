# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditClientSnapConfigs**](SnapshotConfigurationsApi.md#EditClientSnapConfigs) | **Put** /V4/StorageArrays/{arrayId}/Client/{clientId}/Snap/Configs | Edit Client Snap Configs for Storage Arrays
[**EditCopySnapConfigs**](SnapshotConfigurationsApi.md#EditCopySnapConfigs) | **Put** /V4/StorageArrays/Copy/{copyId}/Arrays/{arrayId}/Snap/Configs | Edit Storage Array Snap Configs at Copy Level
[**EditSnapConfigs**](SnapshotConfigurationsApi.md#EditSnapConfigs) | **Put** /V4/StorageArrays/{arrayId}/Snap/Configs | Edit Snap Configs for Storage Array at Array levels
[**EditSubclientSnapConfigs**](SnapshotConfigurationsApi.md#EditSubclientSnapConfigs) | **Put** /V4/StorageArrays/{arrayId}/Subclient/{subclientId}/Snap/Configs | Edit Subclient Snap Configs for Storage Arrays
[**GetClientSnapConfigs**](SnapshotConfigurationsApi.md#GetClientSnapConfigs) | **Get** /V4/StorageArrays/{arrayId}/Client/{clientId}/Snap/Configs | Get Snap Configs for Storage Array at Client
[**GetCopyConfigs**](SnapshotConfigurationsApi.md#GetCopyConfigs) | **Get** /V4/StorageArrays/Copy/{copyId}/Arrays/{arrayId}/Snap/Configs | Get Storage Array Snap Configs at Copy Level
[**GetSnapConfig**](SnapshotConfigurationsApi.md#GetSnapConfig) | **Get** /V4/StorageArrays/Vendors/{vendorId} | Get Snap Configs for storage Array vendor
[**GetSubclientSnapConfigs**](SnapshotConfigurationsApi.md#GetSubclientSnapConfigs) | **Get** /V4/StorageArrays/{arrayId}/Subclient/{subclientId}/Snap/Configs | Get Snap Configs for Storage Array at Subclient
[**GetVendors**](SnapshotConfigurationsApi.md#GetVendors) | **Get** /V4/StorageArrays/Vendors | List Snapshot Vendors

# **EditClientSnapConfigs**
> GenericResp EditClientSnapConfigs(ctx, arrayId, clientId, optional)
Edit Client Snap Configs for Storage Arrays

API to edit Snap Configurations at Client Level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 
  **clientId** | **int32**|  | 
 **optional** | ***SnapshotConfigurationsApiEditClientSnapConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotConfigurationsApiEditClientSnapConfigsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SnapConfigOverrideEditReq**](SnapConfigOverrideEditReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditCopySnapConfigs**
> GenericResp EditCopySnapConfigs(ctx, copyId, arrayId, optional)
Edit Storage Array Snap Configs at Copy Level

API to edit snap configurations for storage array at copy level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **copyId** | **int32**|  | 
  **arrayId** | **int32**|  | 
 **optional** | ***SnapshotConfigurationsApiEditCopySnapConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotConfigurationsApiEditCopySnapConfigsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SnapConfigOverrideEditReq**](SnapConfigOverrideEditReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditSnapConfigs**
> GenericResp EditSnapConfigs(ctx, arrayId, optional)
Edit Snap Configs for Storage Array at Array levels

API to edit Snap configs at a Array level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 
 **optional** | ***SnapshotConfigurationsApiEditSnapConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotConfigurationsApiEditSnapConfigsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SnapConfigOverrideEditReq**](SnapConfigOverrideEditReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditSubclientSnapConfigs**
> GenericResp EditSubclientSnapConfigs(ctx, arrayId, subclientId, optional)
Edit Subclient Snap Configs for Storage Arrays

API to edit Snap Configurations at Subclient level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 
  **subclientId** | **int32**|  | 
 **optional** | ***SnapshotConfigurationsApiEditSubclientSnapConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotConfigurationsApiEditSubclientSnapConfigsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SnapConfigOverrideEditReq**](SnapConfigOverrideEditReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientSnapConfigs**
> SnapConfigOverrideResp GetClientSnapConfigs(ctx, arrayId, clientId)
Get Snap Configs for Storage Array at Client

API to fetch Snap configs for Storage Arrays at client level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 
  **clientId** | **int32**|  | 

### Return type

[**SnapConfigOverrideResp**](SnapConfigOverrideResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCopyConfigs**
> SnapConfigOverrideResp GetCopyConfigs(ctx, copyId, arrayId)
Get Storage Array Snap Configs at Copy Level

API to fetch snap configs for storage arrays at copy level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **copyId** | **int32**|  | 
  **arrayId** | **int32**|  | 

### Return type

[**SnapConfigOverrideResp**](SnapConfigOverrideResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSnapConfig**
> GetSnapConfig GetSnapConfig(ctx, vendorId)
Get Snap Configs for storage Array vendor

get snap config for vendor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorId** | **int32**|  | 

### Return type

[**GetSnapConfig**](GetSnapConfig.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubclientSnapConfigs**
> SnapConfigOverrideResp GetSubclientSnapConfigs(ctx, arrayId, subclientId)
Get Snap Configs for Storage Array at Subclient

API to fetch Snap configs for Storage Arrays at subclient level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **arrayId** | **int32**|  | 
  **subclientId** | **int32**|  | 

### Return type

[**SnapConfigOverrideResp**](SnapConfigOverrideResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVendors**
> SnapVendors GetVendors(ctx, )
List Snapshot Vendors

Get all snap vendors

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SnapVendors**](SnapVendors.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

