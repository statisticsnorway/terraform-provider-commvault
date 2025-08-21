# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnomalousAudits**](AnomalousConditionsApi.md#GetAnomalousAudits) | **Get** /V4/AnomalousAudits | gets anomalous audits
[**GetAnomalousConditions**](AnomalousConditionsApi.md#GetAnomalousConditions) | **Get** /V4/AnomalousConditions | get AnomalousConditions

# **GetAnomalousAudits**
> AnomalousAuditsResp GetAnomalousAudits(ctx, optional)
gets anomalous audits

Get anomalous audits in the specified time

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AnomalousConditionsApiGetAnomalousAuditsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnomalousConditionsApiGetAnomalousAuditsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromTime** | **optional.Int32**| Anomalous audits from a specific time (in epochs) | 
 **toTime** | **optional.Int32**| Anomalous audits till a specific time (in epochs) | 

### Return type

[**AnomalousAuditsResp**](AnomalousAuditsResp.md)

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

