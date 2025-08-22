# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInventoryAsset**](InventoryManagerApi.md#AddInventoryAsset) | **Put** /V4/InventoryManager/Inventory/{inventoryId}/Assets | Adding a new asset to the inventory
[**CreateInventory**](InventoryManagerApi.md#CreateInventory) | **Post** /V4/InventoryManager/Inventory | Creating a new inventory
[**DeleteInventory**](InventoryManagerApi.md#DeleteInventory) | **Delete** /V4/InventoryManager/Inventory/{inventoryId} | Deleting an inventory
[**DeleteInventoryAsset**](InventoryManagerApi.md#DeleteInventoryAsset) | **Delete** /V4/InventoryManager/Inventory/{inventoryId}/Assets/{assetId} | DeleteInventoryAsset
[**GetDefaultInventory**](InventoryManagerApi.md#GetDefaultInventory) | **Get** /V4/InventoryManager/{indexServerClientId}/Default | GetDefaultInventory
[**GetIdentityServerAssetJobHistory**](InventoryManagerApi.md#GetIdentityServerAssetJobHistory) | **Get** /V4/InventoryManager/Inventory/{inventoryId}/Assets/{assetId}/jobs | Retrieve job history for identity server asset
[**GetInventoryAssetDetails**](InventoryManagerApi.md#GetInventoryAssetDetails) | **Get** /V4/InventoryManager/Inventory/{inventoryId}/Assets/{assetId} | GetInventoryAssetDetails
[**GetInventoryAssetList**](InventoryManagerApi.md#GetInventoryAssetList) | **Get** /V4/InventoryManager/Inventory/{inventoryId}/Assets | GetInventoryAssetList
[**GetInventoryAssociatedDataSourceList**](InventoryManagerApi.md#GetInventoryAssociatedDataSourceList) | **Get** /V4/InventoryManager/Inventory/{inventoryId}/DataSources | Get Inventory Associated DataSource List
[**GetInventoryDetails**](InventoryManagerApi.md#GetInventoryDetails) | **Get** /V4/InventoryManager/Inventory/{inventoryId} | Get inventory details
[**GetInventoryList**](InventoryManagerApi.md#GetInventoryList) | **Get** /V4/InventoryManager/Inventory | Retrieve the list of inventories
[**InventoryCrawl**](InventoryManagerApi.md#InventoryCrawl) | **Put** /V4/InventoryManager/Inventory/{inventoryId}/Crawl | Run Backup Copy Job For Plan&#x27;s backupdestination

# **AddInventoryAsset**
> GenericResp AddInventoryAsset(ctx, inventoryId, optional)
Adding a new asset to the inventory

Adding a new asset to the inventory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inventoryId** | **int32**| Inventory id | 
 **optional** | ***InventoryManagerApiAddInventoryAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoryManagerApiAddInventoryAssetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of InventoryIdAssetsBody**](InventoryIdAssetsBody.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateInventory**
> IdName CreateInventory(ctx, optional)
Creating a new inventory

Creating a new inventory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InventoryManagerApiCreateInventoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoryManagerApiCreateInventoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InventoryCreateRequest**](InventoryCreateRequest.md)|  | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInventory**
> GenericResp DeleteInventory(ctx, inventoryId)
Deleting an inventory

Deleting an inventory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inventoryId** | **int32**|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInventoryAsset**
> GenericResp DeleteInventoryAsset(ctx, inventoryId, assetId)
DeleteInventoryAsset

Delete the inventory asset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inventoryId** | **int32**|  | 
  **assetId** | **string**| FQDN of the asset | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefaultInventory**
> InventoryDetails GetDefaultInventory(ctx, indexServerClientId, optional)
GetDefaultInventory

Fetch the default inventory associated to the index server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **indexServerClientId** | **int32**| Pseudo client id of the index server | 
 **optional** | ***InventoryManagerApiGetDefaultInventoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoryManagerApiGetDefaultInventoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIfAbsent** | **optional.Bool**| Create the default inventory if it is missing for the index server | 

### Return type

[**InventoryDetails**](InventoryDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdentityServerAssetJobHistory**
> DatasourceJobHistory GetIdentityServerAssetJobHistory(ctx, inventoryId, assetId)
Retrieve job history for identity server asset

Retrieve job history for identity server asset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inventoryId** | **int32**|  | 
  **assetId** | **string**| FQDN of the asset | 

### Return type

[**DatasourceJobHistory**](DatasourceJobHistory.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInventoryAssetDetails**
> InventoryAssetDetails GetInventoryAssetDetails(ctx, inventoryId, assetId)
GetInventoryAssetDetails

Get the details of inventory asset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inventoryId** | **int32**|  | 
  **assetId** | **string**| FQDN of the asset | 

### Return type

[**InventoryAssetDetails**](InventoryAssetDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInventoryAssetList**
> InventoryAssetList GetInventoryAssetList(ctx, inventoryId)
GetInventoryAssetList

Get the list of assets associated to the inventory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inventoryId** | **int32**| Inventory id | 

### Return type

[**InventoryAssetList**](InventoryAssetList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInventoryAssociatedDataSourceList**
> InventoryAssociatedDatasourceList GetInventoryAssociatedDataSourceList(ctx, inventoryId)
Get Inventory Associated DataSource List

Get the list of data sources associated to the inventory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inventoryId** | **int32**| Inventory Id | 

### Return type

[**InventoryAssociatedDatasourceList**](InventoryAssociatedDatasourceList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInventoryDetails**
> InventoryDetails GetInventoryDetails(ctx, inventoryId)
Get inventory details

Get inventory details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inventoryId** | **int32**|  | 

### Return type

[**InventoryDetails**](InventoryDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInventoryList**
> InventoryList GetInventoryList(ctx, )
Retrieve the list of inventories

Retrieve the list of inventories

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InventoryList**](InventoryList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryCrawl**
> GenericResp InventoryCrawl(ctx, inventoryId)
Run Backup Copy Job For Plan's backupdestination

To start the data collection job on inventory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inventoryId** | **int32**|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

