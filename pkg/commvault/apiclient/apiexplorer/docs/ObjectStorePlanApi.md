# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObjectStorePlan**](ObjectStorePlanApi.md#CreateObjectStorePlan) | **Post** /V4/ObjectStorePlan | Create ObjectStore Plan
[**DeleteObjectStorePlan**](ObjectStorePlanApi.md#DeleteObjectStorePlan) | **Delete** /V4/ObjectStorePlan/{planId} | Delete Object Store Plan
[**GetObjectStorePlanId**](ObjectStorePlanApi.md#GetObjectStorePlanId) | **Get** /V4/ObjectStorePlan/{planId} | Get Object Store Plan Details
[**ModifyObjectStorePlan**](ObjectStorePlanApi.md#ModifyObjectStorePlan) | **Put** /V4/ObjectStorePlan/{planId} | Update the object store plan

# **CreateObjectStorePlan**
> PlanResp CreateObjectStorePlan(ctx, optional)
Create ObjectStore Plan

Create a ObjectStore Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ObjectStorePlanApiCreateObjectStorePlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ObjectStorePlanApiCreateObjectStorePlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateObjectStorePlan**](CreateObjectStorePlan.md)|  | 

### Return type

[**PlanResp**](PlanResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteObjectStorePlan**
> GenericResp DeleteObjectStorePlan(ctx, planId)
Delete Object Store Plan

Used to delete an existing server plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the object store plan to delete | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectStorePlanId**
> ObjectStorePlan GetObjectStorePlanId(ctx, planId)
Get Object Store Plan Details

Get Object Store Plan details 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the object store plan to fetch details | 

### Return type

[**ObjectStorePlan**](ObjectStorePlan.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyObjectStorePlan**
> GenericResp ModifyObjectStorePlan(ctx, planId, optional)
Update the object store plan

Used to modify an exsiting ObjectStore plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan to update | 
 **optional** | ***ObjectStorePlanApiModifyObjectStorePlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ObjectStorePlanApiModifyObjectStorePlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateObjectStorePlan**](UpdateObjectStorePlan.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

