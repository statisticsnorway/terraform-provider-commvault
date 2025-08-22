# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrowseHypervisorInventory**](HypervisorOperationsApi.md#BrowseHypervisorInventory) | **Get** /V4/Hypervisor/{hypervisorId}/{InventoryEntityName}/Browse | Browse the Inventory of your hypervisor by hypervisor ID
[**CreateHypervisor**](HypervisorOperationsApi.md#CreateHypervisor) | **Post** /V4/Hypervisor | Create a Hypervisor
[**DeleteHypervisor**](HypervisorOperationsApi.md#DeleteHypervisor) | **Delete** /V4/Hypervisor/{hypervisorId} | Delete a Hypervisor
[**GetHypervisorFilter**](HypervisorOperationsApi.md#GetHypervisorFilter) | **Get** /v4/hypervisor/{hypervisorId}/filter | Get hypervisor filters
[**GetHypervisors**](HypervisorOperationsApi.md#GetHypervisors) | **Get** /V4/Hypervisor/{hypervisorId} | Get Hypervisor Details
[**ListHypervisors**](HypervisorOperationsApi.md#ListHypervisors) | **Get** /V4/Hypervisor | Get all Hypervisor
[**ReConfigureHypervisor**](HypervisorOperationsApi.md#ReConfigureHypervisor) | **Post** /V4/Hypervisor/{hypervisorId}/Reconfigure | REconfigures and Renew License for  the hypervisor client
[**SetHypervisorAccessNode**](HypervisorOperationsApi.md#SetHypervisorAccessNode) | **Put** /V4/Hypervisor/{hypervisorId}/AccessNode | Change the Access Node for Hypervisor
[**UpdateHypervisor**](HypervisorOperationsApi.md#UpdateHypervisor) | **Put** /V4/Hypervisor/{hypervisorId} | Update the Hypervisor details
[**UpdateHypervisorCredentials**](HypervisorOperationsApi.md#UpdateHypervisorCredentials) | **Put** /V4/Hypervisor/{HypervisorId}/Credentials | Update Hypervisor Credentials
[**UpdateHypervisorFilter**](HypervisorOperationsApi.md#UpdateHypervisorFilter) | **Put** /v4/hypervisor/{hypervisorId}/filter | Update hypervisor filters

# **BrowseHypervisorInventory**
> InlineResponse20013 BrowseHypervisorInventory(ctx, hypervisorId, inventoryEntityName, optional)
Browse the Inventory of your hypervisor by hypervisor ID

Browse the Inventory of your hypervisor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hypervisorId** | **string**| Hypervisor client ID to browse | 
  **inventoryEntityName** | **string**| Name of the inventory entity that needs to be browsed like ESX Host name in VCenter | 
 **optional** | ***HypervisorOperationsApiBrowseHypervisorInventoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HypervisorOperationsApiBrowseHypervisorInventoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inventoryType** | [**optional.Interface of VmContentType**](.md)| Type of resource to browse in inventory | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHypervisor**
> InlineResponse20012 CreateHypervisor(ctx, optional)
Create a Hypervisor

Create Hypervisor for that particular type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HypervisorOperationsApiCreateHypervisorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HypervisorOperationsApiCreateHypervisorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of V4HypervisorBody**](V4HypervisorBody.md)|  | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHypervisor**
> GenericResp DeleteHypervisor(ctx, hypervisorId)
Delete a Hypervisor

delete an existing Hypervisor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hypervisorId** | **int32**| Id of the Hypervisor to delete | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHypervisorFilter**
> GetHypervisorFilterResp GetHypervisorFilter(ctx, hypervisorId, optional)
Get hypervisor filters

Get hypervisor filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hypervisorId** | **string**|  | 
 **optional** | ***HypervisorOperationsApiGetHypervisorFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HypervisorOperationsApiGetHypervisorFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backupsetId** | **optional.String**| if not specified the default backupset will be fetched | 

### Return type

[**GetHypervisorFilterResp**](GetHypervisorFilterResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHypervisors**
> HypervisorPropertiesResp GetHypervisors(ctx, hypervisorId)
Get Hypervisor Details

Get the details of  HYpervisor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hypervisorId** | **int32**| Id of the HYpervisor to get | 

### Return type

[**HypervisorPropertiesResp**](HypervisorPropertiesResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHypervisors**
> ListHypervisors ListHypervisors(ctx, )
Get all Hypervisor

Get the details of all Hypervisor

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListHypervisors**](ListHypervisors.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReConfigureHypervisor**
> GenericResp ReConfigureHypervisor(ctx, hypervisorId)
REconfigures and Renew License for  the hypervisor client

REconfigures and Renew License for  the hypervisor client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hypervisorId** | **string**| Hypervisor client ID to reconfigure | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetHypervisorAccessNode**
> SetHypervisorAccessNode(ctx, hypervisorId, optional)
Change the Access Node for Hypervisor

Endpoint to Change the Access Node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hypervisorId** | **int32**| Hypervisor ID to update the Access Node | 
 **optional** | ***HypervisorOperationsApiSetHypervisorAccessNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HypervisorOperationsApiSetHypervisorAccessNodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of HypervisorIdAccessNodeBody**](HypervisorIdAccessNodeBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHypervisor**
> HypervisorPropertiesResp UpdateHypervisor(ctx, hypervisorId, optional)
Update the Hypervisor details

Updates the Hypervisors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hypervisorId** | **int32**| Id of the Hypervisor to update | 
 **optional** | ***HypervisorOperationsApiUpdateHypervisorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HypervisorOperationsApiUpdateHypervisorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of HypervisorHypervisorIdBody**](HypervisorHypervisorIdBody.md)|  | 

### Return type

[**HypervisorPropertiesResp**](HypervisorPropertiesResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHypervisorCredentials**
> CreateHypervisorResp UpdateHypervisorCredentials(ctx, hypervisorId, optional)
Update Hypervisor Credentials

Update Hypervisor's credentials information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hypervisorId** | **string**|  | 
 **optional** | ***HypervisorOperationsApiUpdateHypervisorCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HypervisorOperationsApiUpdateHypervisorCredentialsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of HypervisorIdCredentialsBody**](HypervisorIdCredentialsBody.md)|  | 

### Return type

[**CreateHypervisorResp**](CreateHypervisorResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHypervisorFilter**
> GenericResp UpdateHypervisorFilter(ctx, hypervisorId, optional)
Update hypervisor filters

Update hypervisor filter details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hypervisorId** | **string**|  | 
 **optional** | ***HypervisorOperationsApiUpdateHypervisorFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HypervisorOperationsApiUpdateHypervisorFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateHypervisorFilterReq**](UpdateHypervisorFilterReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

