# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSchedulePolicies**](ScheduleApi.md#GetSchedulePolicies) | **Post** /V4/SchedulePolicy/list | Endpoint to get schedule policies
[**Getschedules**](ScheduleApi.md#Getschedules) | **Post** /V4/Schedule/list | Endpoint to get schedules based on filters
[**ModifySchedulePattern**](ScheduleApi.md#ModifySchedulePattern) | **Put** /V4/SchedulePolicy/{schedulePolicyId}/Schedule/{scheduleId}/Pattern | Modify Schedule Pattern of Schedule Policy

# **GetSchedulePolicies**
> TaskListResponse GetSchedulePolicies(ctx, optional)
Endpoint to get schedule policies

API to fetch non plan schedule policies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScheduleApiGetSchedulePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScheduleApiGetSchedulePoliciesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TaskListRequest**](TaskListRequest.md)|  | 

### Return type

[**TaskListResponse**](TaskListResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getschedules**
> TaskListResponse Getschedules(ctx, optional)
Endpoint to get schedules based on filters

API to get non plan schedules based on various filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScheduleApiGetschedulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScheduleApiGetschedulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TaskListRequest**](TaskListRequest.md)| Filter options for the result set | 

### Return type

[**TaskListResponse**](TaskListResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifySchedulePattern**
> GenericResp ModifySchedulePattern(ctx, schedulePolicyId, scheduleId, optional)
Modify Schedule Pattern of Schedule Policy

API to update pattern for schedule in schedule policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **schedulePolicyId** | **string**|  | 
  **scheduleId** | **string**|  | 
 **optional** | ***ScheduleApiModifySchedulePatternOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScheduleApiModifySchedulePatternOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PlanPattern**](PlanPattern.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

