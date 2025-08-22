# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletagValueForTagId**](EntityTagApi.md#DeletagValueForTagId) | **Delete** /V4/Tags/{tagId}/TagValues/{tagValue} | Delete given tagValue for tagId
[**DeleteEntityTag**](EntityTagApi.md#DeleteEntityTag) | **Delete** /V4/EntityTags/{tagId} | Delete tag from the entity tag list
[**GetTagAssociatedEntities**](EntityTagApi.md#GetTagAssociatedEntities) | **Get** /V4/Tags/AssociatedEntities | Get list of tags and count of its associated entities
[**GetTagValuesForTagId**](EntityTagApi.md#GetTagValuesForTagId) | **Get** /V4/Tags/{tagId}/Values | Get list of values for given tagId
[**PutValuesForTagId**](EntityTagApi.md#PutValuesForTagId) | **Put** /V4/Tags/{tagId}/Values | Modify values for give tagId

# **DeletagValueForTagId**
> GenericResp DeletagValueForTagId(ctx, tagId, tagValue)
Delete given tagValue for tagId

Delete given tagValue for tagId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagId** | **int32**| tag id | 
  **tagValue** | **string**| tag value to delete | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEntityTag**
> GenericResp DeleteEntityTag(ctx, tagId)
Delete tag from the entity tag list

Used to delete tag from the entity tag list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagId** | **int32**| Id of the tag to delete | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagAssociatedEntities**
> TagAssociatedEntitiesResponse GetTagAssociatedEntities(ctx, optional)
Get list of tags and count of its associated entities

Returns list of tags for the logged in user's default entity tag set along with count of associated entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EntityTagApiGetTagAssociatedEntitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntityTagApiGetTagAssociatedEntitiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagId** | **optional.Int32**| Id of the tag whose associated entities details needs to be returned | 

### Return type

[**TagAssociatedEntitiesResponse**](TagAssociatedEntitiesResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagValuesForTagId**
> TagValueResponse GetTagValuesForTagId(ctx, tagId)
Get list of values for given tagId

Gives list of values for given tagId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagId** | **int32**| Id of the tag name whose value is requested | 

### Return type

[**TagValueResponse**](TagValueResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutValuesForTagId**
> GenericResp PutValuesForTagId(ctx, tagId, optional)
Modify values for give tagId

Modify values for give tagId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagId** | **int32**| Id of the tag name whose value is requested | 
 **optional** | ***EntityTagApiPutValuesForTagIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntityTagApiPutValuesForTagIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateTagValueRequest**](UpdateTagValueRequest.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

