# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAggregateOnEntity**](GlobalSearchApi.md#GetAggregateOnEntity) | **Get** /V4/{globalSearchEntity}/aggregate | Get Aggregate count of Entity

# **GetAggregateOnEntity**
> GlobalEntityAggregationValueResp GetAggregateOnEntity(ctx, globalSearchEntity, func_, optional)
Get Aggregate count of Entity

Get aggregate count of entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **globalSearchEntity** | **string**| name of global search entity | 
  **func_** | **string**| aggregation function to be applied. | 
 **optional** | ***GlobalSearchApiGetAggregateOnEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalSearchApiGetAggregateOnEntityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fq** | **optional.String**| Filter criteria to filter out global search entities | 
 **groupBy** | **optional.String**| comma-separated list of entity response attributes based on which aggregate results would be grouped | 
 **aggregateOn** | **optional.String**| attribute on which aggregation function would be applied. | 

### Return type

[**GlobalEntityAggregationValueResp**](GlobalEntityAggregationValueResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

