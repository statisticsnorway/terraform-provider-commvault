# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGlobalServerPlan**](GlobalServerPlanApi.md#CreateGlobalServerPlan) | **Post** /V4/Global/ServerPlan | Create Global Server Plan

# **CreateGlobalServerPlan**
> PlanResp CreateGlobalServerPlan(ctx, optional)
Create Global Server Plan

Create a Global Server Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GlobalServerPlanApiCreateGlobalServerPlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalServerPlanApiCreateGlobalServerPlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateGlobalServerPlan**](CreateGlobalServerPlan.md)|  | 

### Return type

[**PlanResp**](PlanResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

