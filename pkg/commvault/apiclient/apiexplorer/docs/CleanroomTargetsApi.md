# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRecoveryTarget**](CleanroomTargetsApi.md#DeleteRecoveryTarget) | **Delete** /V4/RecoveryTarget/{recoveryTargetId} | Delete existing Recovery Target
[**GetRecoveryTarget**](CleanroomTargetsApi.md#GetRecoveryTarget) | **Get** /V4/RecoveryTarget/{recoveryTargetId} | Get an existing recovery target details
[**GetRecoveryTargets**](CleanroomTargetsApi.md#GetRecoveryTargets) | **Get** /V4/RecoveryTargets | Get Recovery Target List

# **DeleteRecoveryTarget**
> GenericResp DeleteRecoveryTarget(ctx, recoveryTargetId)
Delete existing Recovery Target

To delete an exisitng recovery target

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recoveryTargetId** | **int32**| id of recovery target | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRecoveryTarget**
> InlineResponse20018 GetRecoveryTarget(ctx, recoveryTargetId)
Get an existing recovery target details

To get recovery target details by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recoveryTargetId** | **int32**| id of recovery target | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRecoveryTargets**
> GetRecoveryTargetListResponse GetRecoveryTargets(ctx, )
Get Recovery Target List

Get the list of recovery targets

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetRecoveryTargetListResponse**](GetRecoveryTargetListResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

