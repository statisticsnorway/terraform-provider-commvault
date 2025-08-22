# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRegion**](RegionsApi.md#CreateRegion) | **Post** /V4/Regions | Create Region
[**DeleteRegion**](RegionsApi.md#DeleteRegion) | **Delete** /V4/Regions/{regionId} | Delete Region
[**GetRegionDetails**](RegionsApi.md#GetRegionDetails) | **Get** /V4/Regions/{regionId} | Get Region Details
[**GetRegionForEntities**](RegionsApi.md#GetRegionForEntities) | **Get** /V4/entity/region | Get Region for entities
[**GetRegionForEntity**](RegionsApi.md#GetRegionForEntity) | **Get** /V4/entity/{entityType}/{entityId}/region | Get Region Details of entity
[**GetRegions**](RegionsApi.md#GetRegions) | **Get** /V4/Regions | Get Regions
[**SetRegionForEntity**](RegionsApi.md#SetRegionForEntity) | **Put** /V4/entity/{entityType}/{entityId}/region | Set Region for an entity
[**UpdateRegion**](RegionsApi.md#UpdateRegion) | **Put** /V4/Regions/{regionId} | Update Region

# **CreateRegion**
> IdNameGuid CreateRegion(ctx, optional)
Create Region

Create a region

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RegionsApiCreateRegionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegionsApiCreateRegionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateRegion**](CreateRegion.md)|  | 

### Return type

[**IdNameGuid**](IdNameGUID.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRegion**
> GenericResp DeleteRegion(ctx, regionId)
Delete Region

Delete the region

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regionId** | **string**|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRegionDetails**
> RegionDetails GetRegionDetails(ctx, regionId)
Get Region Details

Get details of the region

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regionId** | **string**|  | 

### Return type

[**RegionDetails**](RegionDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRegionForEntities**
> GetEntityRegionResp GetRegionForEntities(ctx, entityType, entities, optional)
Get Region for entities

Fetch region for multiple entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityType** | [**EntityTypes**](.md)| Entity Type Enum | 
  **entities** | **string**| Comma Seprated Entity Ids | 
 **optional** | ***RegionsApiGetRegionForEntitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegionsApiGetRegionForEntitiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **regionType** | [**optional.Interface of EntityRegionTypes**](.md)| Region Type Enum | 

### Return type

[**GetEntityRegionResp**](GetEntityRegionResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRegionForEntity**
> IdNameDisplayName GetRegionForEntity(ctx, entityType, entityId, optional)
Get Region Details of entity

Api to fetch region details for an entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityType** | [**EntityTypes**](.md)| Type of the entity | 
  **entityId** | **int32**| Unique id for the entity | 
 **optional** | ***RegionsApiGetRegionForEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegionsApiGetRegionForEntityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **calculate** | **optional.Bool**| Flag for Enable/Disable Region Calculation | 
 **entityRegionType** | [**optional.Interface of EntityRegionTypes**](.md)| Region Type Enum | 

### Return type

[**IdNameDisplayName**](IdNameDisplayName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRegions**
> RegionsListResp GetRegions(ctx, optional)
Get Regions

Get list of regions 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RegionsApiGetRegionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegionsApiGetRegionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**optional.Interface of RegionType**](.md)| Region type to be filtered | 

### Return type

[**RegionsListResp**](RegionsListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetRegionForEntity**
> GenericResp SetRegionForEntity(ctx, entityType, entityId, optional)
Set Region for an entity

Api to set region for an entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityType** | [**EntityTypes**](.md)| Type of the entity | 
  **entityId** | **int32**| Unique id for the entity | 
 **optional** | ***RegionsApiSetRegionForEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegionsApiSetRegionForEntityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EntityRegionInfo**](EntityRegionInfo.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRegion**
> GenericResp UpdateRegion(ctx, regionId, optional)
Update Region

Update details of the region

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regionId** | **string**|  | 
 **optional** | ***RegionsApiUpdateRegionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegionsApiUpdateRegionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateRegion**](UpdateRegion.md)| Properties to be updated for the region | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

