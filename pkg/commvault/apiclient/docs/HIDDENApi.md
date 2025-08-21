# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInventoryAsset**](HIDDENApi.md#AddInventoryAsset) | **Put** /V4/InventoryManager/Inventory/{inventoryId}/Assets | Adding a new asset to the inventory
[**AddorUpdateGlobalSettings**](HIDDENApi.md#AddorUpdateGlobalSettings) | **Put** /V4/GlobalSettings | Add or Update Global Settings
[**AssociateMAWithStoragePool**](HIDDENApi.md#AssociateMAWithStoragePool) | **Put** /V4/StoragePool/{storagePoolId}/MediaAgent | Associate DDB or storage MediaAgent with the Storage Pool
[**CompareXmls**](HIDDENApi.md#CompareXmls) | **Post** /V4/CompareXML | 
[**CreateAzureADClient**](HIDDENApi.md#CreateAzureADClient) | **Post** /v4/ActiveDirectory/AzureAD | Create Azure Active Directory app
[**CreateInventory**](HIDDENApi.md#CreateInventory) | **Post** /V4/InventoryManager/Inventory | Creating a new inventory
[**CreateServiceMeshClient**](HIDDENApi.md#CreateServiceMeshClient) | **Post** /V4/ServiceMesh | Create Service Mesh
[**DeleteBulkLaptopOwnerMappingActionDelete**](HIDDENApi.md#DeleteBulkLaptopOwnerMappingActionDelete) | **Put** /V4/LaptopOwnerMapping/action/delete | delete multiple laptopowner mappings
[**DeleteInventory**](HIDDENApi.md#DeleteInventory) | **Delete** /V4/InventoryManager/Inventory/{inventoryId} | Deleting an inventory
[**DeleteInventoryAsset**](HIDDENApi.md#DeleteInventoryAsset) | **Delete** /V4/InventoryManager/Inventory/{inventoryId}/Assets/{assetId} | DeleteInventoryAsset
[**DeleteRequestManagerRequest**](HIDDENApi.md#DeleteRequestManagerRequest) | **Delete** /V4/RequestManager/Request/{requestId} | 
[**DeleteServiceMeshClient**](HIDDENApi.md#DeleteServiceMeshClient) | **Delete** /V4/ServiceMesh/{clientId} | Delete Service Mesh Client
[**DisableLocalAuthenticationForCompany**](HIDDENApi.md#DisableLocalAuthenticationForCompany) | **Put** /V4/Company/{companyId}/LocalAuthentication/Action/Disable | 
[**EnableAirgapOnStoragePool**](HIDDENApi.md#EnableAirgapOnStoragePool) | **Put** /V4/StoragePool/{storagePoolId}/Airgap | Enable Airgap on StoragePool
[**EnableLocalAuthenticationForCompany**](HIDDENApi.md#EnableLocalAuthenticationForCompany) | **Put** /V4/Company/{companyId}/LocalAuthentication/Action/Enable | 
[**FetchLocalAuthenticationDetilsOfCompany**](HIDDENApi.md#FetchLocalAuthenticationDetilsOfCompany) | **Get** /V4/Company/{companyId}/LocalAuthentication | 
[**GetAnomalousConditions**](HIDDENApi.md#GetAnomalousConditions) | **Get** /V4/AnomalousConditions | get AnomalousConditions
[**GetBigDataList**](HIDDENApi.md#GetBigDataList) | **Get** /V4/BigDataApps | 
[**GetClientInfoForRetire**](HIDDENApi.md#GetClientInfoForRetire) | **Get** /V4/Client/{clientId}/RetireInformation | Get ClientInfo ForRetire
[**GetDefaultInventory**](HIDDENApi.md#GetDefaultInventory) | **Get** /V4/InventoryManager/{indexServerClientId}/Default | GetDefaultInventory
[**GetDistributedStorages**](HIDDENApi.md#GetDistributedStorages) | **Get** /V4/Storage/DistributedStorage | Get Distributed Storages
[**GetDistributedSystems**](HIDDENApi.md#GetDistributedSystems) | **Get** /V4/DistributedSystems | 
[**GetEntitySettings**](HIDDENApi.md#GetEntitySettings) | **Get** /V4/EntitySettings | Get Entity Settings
[**GetIdentityServerAssetJobHistory**](HIDDENApi.md#GetIdentityServerAssetJobHistory) | **Get** /V4/InventoryManager/Inventory/{inventoryId}/Assets/{assetId}/jobs | Retrieve job history for identity server asset
[**GetInventoryAssetDetails**](HIDDENApi.md#GetInventoryAssetDetails) | **Get** /V4/InventoryManager/Inventory/{inventoryId}/Assets/{assetId} | GetInventoryAssetDetails
[**GetInventoryAssetList**](HIDDENApi.md#GetInventoryAssetList) | **Get** /V4/InventoryManager/Inventory/{inventoryId}/Assets | GetInventoryAssetList
[**GetInventoryAssociatedDataSourceList**](HIDDENApi.md#GetInventoryAssociatedDataSourceList) | **Get** /V4/InventoryManager/Inventory/{inventoryId}/DataSources | Get Inventory Associated DataSource List
[**GetInventoryDetails**](HIDDENApi.md#GetInventoryDetails) | **Get** /V4/InventoryManager/Inventory/{inventoryId} | Get inventory details
[**GetInventoryList**](HIDDENApi.md#GetInventoryList) | **Get** /V4/InventoryManager/Inventory | Retrieve the list of inventories
[**GetLaptopOwnerMapping**](HIDDENApi.md#GetLaptopOwnerMapping) | **Get** /V4/LaptopOwnerMapping | Get LaptopOwner Mapping
[**GetMediaAgentDDBDisks**](HIDDENApi.md#GetMediaAgentDDBDisks) | **Get** /V4/mediaAgent/{mediaAgentId}/DDB/Disks | Get MediaAgent DDBDisks
[**GetProjectsForRequest**](HIDDENApi.md#GetProjectsForRequest) | **Get** /V4/RequestManager/Request/{requestId}/Projects | 
[**GetRequestDetails**](HIDDENApi.md#GetRequestDetails) | **Get** /V4/RequestManager/Request/{requestId} | Retrieve details of an existing request
[**GetServiceMeshDetails**](HIDDENApi.md#GetServiceMeshDetails) | **Get** /V4/ServiceMesh | Get Service Mesh details
[**GetUserauthenticationMethods**](HIDDENApi.md#GetUserauthenticationMethods) | **Get** /V4/user/{userId}/authenticationMethods | 
[**GetV4LaptopOwnerMappingPreview**](HIDDENApi.md#GetV4LaptopOwnerMappingPreview) | **Get** /V4/LaptopOwnerMapping/Preview | 
[**InventoryCrawl**](HIDDENApi.md#InventoryCrawl) | **Put** /V4/InventoryManager/Inventory/{inventoryId}/Crawl | Run Backup Copy Job For Plan&#x27;s backupdestination
[**MediaAgentDDBDiskMgmt**](HIDDENApi.md#MediaAgentDDBDiskMgmt) | **Put** /V4/mediaAgent/{mediaAgentId}/DDB/Disks | Perform DDB disk management operations on MediaAgent
[**PostLaptopOwnerMapping**](HIDDENApi.md#PostLaptopOwnerMapping) | **Post** /V4/LaptopOwnerMapping | Create laptop owner mappings
[**PostVmActionDelete**](HIDDENApi.md#PostVmActionDelete) | **Post** /V4/VM/{vmGUID}/Action/Delete | Delete the provisioned virtual machine
[**PostVmActionRefresh**](HIDDENApi.md#PostVmActionRefresh) | **Post** /V4/VM/{vmGUID}/Action/Refresh | Refresh the provisioned virtual machine
[**PostVmActionRenew**](HIDDENApi.md#PostVmActionRenew) | **Post** /V4/VM/{vmGUID}/Action/Renew | Renew provisioned virtual machine
[**PutLaptopOwnerMapping**](HIDDENApi.md#PutLaptopOwnerMapping) | **Put** /V4/LaptopOwnerMapping | Update existing laptop owner mapping
[**PutLaptopOwnerMappingActionAssign**](HIDDENApi.md#PutLaptopOwnerMappingActionAssign) | **Put** /V4/LaptopOwnerMapping/action/Assign | Put LaptopOwner Mapping Action Assign
[**RMConfigureRequestOperation**](HIDDENApi.md#RMConfigureRequestOperation) | **Put** /V4/RequestManager/Request/{requestId}/Configure | RequestManager - Configure Request Operation
[**RMCreateRequestOperation**](HIDDENApi.md#RMCreateRequestOperation) | **Post** /V4/RequestManager/Request | RequestManager - Create Request Operation
[**RequestManagerRequestList**](HIDDENApi.md#RequestManagerRequestList) | **Get** /V4/RequestManager/Request | 
[**UpdateServiceMeshClientProperties**](HIDDENApi.md#UpdateServiceMeshClientProperties) | **Put** /V4/ServiceMesh/{clientId} | Update properties of Service mesh client

# **AddInventoryAsset**
> GenericResp AddInventoryAsset(ctx, inventoryId, optional)
Adding a new asset to the inventory

Adding a new asset to the inventory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inventoryId** | **int32**| Inventory id | 
 **optional** | ***HIDDENApiAddInventoryAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiAddInventoryAssetOpts struct
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

# **AddorUpdateGlobalSettings**
> GenericResp AddorUpdateGlobalSettings(ctx, optional)
Add or Update Global Settings

This API allows you to add new global settings or update existing ones, overriding the default configurations to customize the system's behavior.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HIDDENApiAddorUpdateGlobalSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiAddorUpdateGlobalSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ModifyGlobalSettings**](ModifyGlobalSettings.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssociateMAWithStoragePool**
> GenericResp AssociateMAWithStoragePool(ctx, storagePoolId, mediaAgentId, action)
Associate DDB or storage MediaAgent with the Storage Pool

API to manage MediaAgents associated with the Storage Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the Storage Pool | 
  **mediaAgentId** | **int32**| Id of the MediaAgent | 
  **action** | **string**| Action to performed on the MediaAgent | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompareXmls**
> CompareXmlResp CompareXmls(ctx, optional)


Compare old and new XMl or objects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HIDDENApiCompareXmlsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiCompareXmlsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CompareXmlReq**](CompareXmlReq.md)|  | 

### Return type

[**CompareXmlResp**](CompareXMLResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAzureADClient**
> IdName CreateAzureADClient(ctx, optional)
Create Azure Active Directory app

Create Azure Active Directory app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HIDDENApiCreateAzureADClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiCreateAzureADClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateAzureAdClient**](CreateAzureAdClient.md)|  | 

### Return type

[**IdName**](IdName.md)

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
 **optional** | ***HIDDENApiCreateInventoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiCreateInventoryOpts struct
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

# **CreateServiceMeshClient**
> IdName CreateServiceMeshClient(ctx, optional)
Create Service Mesh

Create a Service mesh client representation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HIDDENApiCreateServiceMeshClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiCreateServiceMeshClientOpts struct
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

# **DeleteBulkLaptopOwnerMappingActionDelete**
> GenericResp DeleteBulkLaptopOwnerMappingActionDelete(ctx, optional)
delete multiple laptopowner mappings

API is used to delete multiple laptopowner mappings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HIDDENApiDeleteBulkLaptopOwnerMappingActionDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiDeleteBulkLaptopOwnerMappingActionDeleteOpts struct
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

# **DisableLocalAuthenticationForCompany**
> GenericResp DisableLocalAuthenticationForCompany(ctx, companyId)


Disable local authentication for the company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | **int32**| Id of the company | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableAirgapOnStoragePool**
> GenericResp EnableAirgapOnStoragePool(ctx, storagePoolId)
Enable Airgap on StoragePool

This API enables Airgap on StoragePool and all its associated MediaAgents. This action is irreversible, once enabled Airgap cannot be disabled on the StoragePool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePoolId** | **int32**| Id of the StoragePool | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableLocalAuthenticationForCompany**
> GenericResp EnableLocalAuthenticationForCompany(ctx, companyId, optional)


Enable local authentication for the company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | **int32**| Id of the company | 
 **optional** | ***HIDDENApiEnableLocalAuthenticationForCompanyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiEnableLocalAuthenticationForCompanyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of LocalAuthenticationDetails**](LocalAuthenticationDetails.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchLocalAuthenticationDetilsOfCompany**
> LocalAuthenticationDetailsResponse FetchLocalAuthenticationDetilsOfCompany(ctx, companyId)


Fetch local authentication details of the company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | **int32**| Id of the company | 

### Return type

[**LocalAuthenticationDetailsResponse**](LocalAuthenticationDetailsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAnomalousConditions**
> InlineResponse20022 GetAnomalousConditions(ctx, fromTime)
get AnomalousConditions

Get various anomalous conditions like events, jobs, offline clients, high CPU and memory loaded clients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fromTime** | **string**| unix time stamp denotes from which the anomalous events should be retrieved | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBigDataList**
> BigDataAppListResp GetBigDataList(ctx, )


This endpoint is used to return the list of big data apps.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BigDataAppListResp**](BigDataAppListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientInfoForRetire**
> ClientInfoForRetire GetClientInfoForRetire(ctx, clientId)
Get ClientInfo ForRetire

API to get client information that retire operation concerns

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **int32**|  | 

### Return type

[**ClientInfoForRetire**](ClientInfoForRetire.md)

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
 **optional** | ***HIDDENApiGetDefaultInventoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiGetDefaultInventoryOpts struct
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

# **GetDistributedStorages**
> DistributedStorageListResp GetDistributedStorages(ctx, )
Get Distributed Storages

Get the list of distributed storages

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DistributedStorageListResp**](DistributedStorageListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDistributedSystems**
> DistributedSystemsListResp GetDistributedSystems(ctx, )


This endpoint is used to return the list of distributed systems.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DistributedSystemsListResp**](DistributedSystemsListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntitySettings**
> EntitySettingsResponse GetEntitySettings(ctx, )
Get Entity Settings

Get list of entity settings used to modify default behaviour for linked entity like servers or server groups

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EntitySettingsResponse**](EntitySettingsResponse.md)

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

# **GetLaptopOwnerMapping**
> LaptopOwnerMapping GetLaptopOwnerMapping(ctx, optional)
Get LaptopOwner Mapping

API to get the list of laptop owner mappings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HIDDENApiGetLaptopOwnerMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiGetLaptopOwnerMappingOpts struct
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

# **GetMediaAgentDDBDisks**
> MaddbDiskMgmtResp GetMediaAgentDDBDisks(ctx, mediaAgentId)
Get MediaAgent DDBDisks

Fetch the list of DDB disks hosted on the MediaAgent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mediaAgentId** | **int32**| Id of the MediaAgent | 

### Return type

[**MaddbDiskMgmtResp**](MADDBDiskMgmtResp.md)

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

# **GetUserauthenticationMethods**
> UserAuthenticationMethods GetUserauthenticationMethods(ctx, userId, optional)


Used to fetch authentication methods of an existing user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| Id of the User whose authentication method details have to be fetched | 
 **optional** | ***HIDDENApiGetUserauthenticationMethodsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiGetUserauthenticationMethodsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fetchAvailableMethodsOnly** | **optional.Bool**| Set as true to fetch the list of available authentication methods for the user | 

### Return type

[**UserAuthenticationMethods**](UserAuthenticationMethods.md)

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

# **MediaAgentDDBDiskMgmt**
> MaddbDiskMgmtResp MediaAgentDDBDiskMgmt(ctx, mediaAgentId, optional)
Perform DDB disk management operations on MediaAgent

Perform DDB disk management operations on MediaAgent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mediaAgentId** | **int32**| Id of the MediaAgent | 
 **optional** | ***HIDDENApiMediaAgentDDBDiskMgmtOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiMediaAgentDDBDiskMgmtOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MaddbDiskMgmtReq**](MaddbDiskMgmtReq.md)|  | 

### Return type

[**MaddbDiskMgmtResp**](MADDBDiskMgmtResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
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
 **optional** | ***HIDDENApiPostLaptopOwnerMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiPostLaptopOwnerMappingOpts struct
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

# **PostVmActionDelete**
> GenericResp PostVmActionDelete(ctx, vmGUID)
Delete the provisioned virtual machine

Deletes the VM.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmGUID** | **string**| GUID of the Provisioned VM | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVmActionRefresh**
> GenericResp PostVmActionRefresh(ctx, vmGUID)
Refresh the provisioned virtual machine

Refreshes the VM.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmGUID** | **string**| GUID of the Provisioned VM | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVmActionRenew**
> GenericResp PostVmActionRenew(ctx, vmGUID, optional)
Renew provisioned virtual machine

Renew the VM with the provided timestamp.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmGUID** | **string**| GUID of the Provisioned VM | 
 **optional** | ***HIDDENApiPostVmActionRenewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiPostVmActionRenewOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ActionRenewBody**](ActionRenewBody.md)| The renewal timestamp has to be provided in unix format. | 

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
 **optional** | ***HIDDENApiPutLaptopOwnerMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiPutLaptopOwnerMappingOpts struct
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
 **optional** | ***HIDDENApiPutLaptopOwnerMappingActionAssignOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiPutLaptopOwnerMappingActionAssignOpts struct
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

# **RMConfigureRequestOperation**
> GenericResp RMConfigureRequestOperation(ctx, requestId, optional)
RequestManager - Configure Request Operation

Configure a created request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **int32**|  | 
 **optional** | ***HIDDENApiRMConfigureRequestOperationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiRMConfigureRequestOperationOpts struct
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
 **optional** | ***HIDDENApiRMCreateRequestOperationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiRMCreateRequestOperationOpts struct
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
 **optional** | ***HIDDENApiRequestManagerRequestListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiRequestManagerRequestListOpts struct
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

# **UpdateServiceMeshClientProperties**
> GenericResp UpdateServiceMeshClientProperties(ctx, optional)
Update properties of Service mesh client

Update properties of a Service mesh client representation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HIDDENApiUpdateServiceMeshClientPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HIDDENApiUpdateServiceMeshClientPropertiesOpts struct
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

