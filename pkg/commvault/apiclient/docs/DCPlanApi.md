# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDCPlan**](DCPlanApi.md#CreateDCPlan) | **Post** /V4/DCPlan | Create a Data Classification Plan
[**DeleteDCPlan**](DCPlanApi.md#DeleteDCPlan) | **Delete** /V4/DCPlan/{planId} | Delete Data Classification Plan
[**GetDCPlanById**](DCPlanApi.md#GetDCPlanById) | **Get** /V4/DCPlan/{planId} | Get Data Classification Plan By ID
[**ModifyDCPlan**](DCPlanApi.md#ModifyDCPlan) | **Put** /V4/DCPlan/{planId} | Modify Data Classification Plan

# **CreateDCPlan**
> PlanResp CreateDCPlan(ctx, optional)
Create a Data Classification Plan

Create a Data Classification Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DCPlanApiCreateDCPlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DCPlanApiCreateDCPlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateDcPlan**](CreateDcPlan.md)|  | 

### Return type

[**PlanResp**](PlanResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDCPlan**
> GenericResp DeleteDCPlan(ctx, planId)
Delete Data Classification Plan

Used to delete an existing data classification plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the plan to delete | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDCPlanById**
> DcPlanDetails GetDCPlanById(ctx, planId)
Get Data Classification Plan By ID

Get Data Classification Plan details 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the plan to fetch details | 

### Return type

[**DcPlanDetails**](DCPlanDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyDCPlan**
> GenericResp ModifyDCPlan(ctx, planId, optional)
Modify Data Classification Plan

Used to modify an existing data classification plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan to update | 
 **optional** | ***DCPlanApiModifyDCPlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DCPlanApiModifyDCPlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateDcPlan**](UpdateDcPlan.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

