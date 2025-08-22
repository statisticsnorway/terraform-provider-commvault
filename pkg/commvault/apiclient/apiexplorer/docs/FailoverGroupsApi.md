# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFailoverGroupSchedule**](FailoverGroupsApi.md#CreateFailoverGroupSchedule) | **Post** /V4/FailoverGroups/{failoverGroupId}/Schedules | Create failover group DR operation schedule
[**DeleteFailoverGroup**](FailoverGroupsApi.md#DeleteFailoverGroup) | **Delete** /V4/FailoverGroups/{failoverGroupId} | Delete failover group
[**DeleteFailoverGroupSchedule**](FailoverGroupsApi.md#DeleteFailoverGroupSchedule) | **Delete** /V4/FailoverGroups/{failoverGroupId}/Schedules/{scheduleId} | Delete failover group DR operation schedule
[**GetEligibleMachinesFailoverGroup**](FailoverGroupsApi.md#GetEligibleMachinesFailoverGroup) | **Get** /V4/FailoverGroups/EligibleMachines | Get list of eligible machines that can be added to failover groups
[**GetFailoverGroupDetails**](FailoverGroupsApi.md#GetFailoverGroupDetails) | **Get** /V4/FailoverGroups/{failoverGroupId} | Get Failover group details
[**GetTestFailoverMachinesFailoverGroup**](FailoverGroupsApi.md#GetTestFailoverMachinesFailoverGroup) | **Get** /V4/FailoverGroups/{failoverGroupId}/TestFailoverMachines | Get test failover VMs list
[**ListFailoverGroups**](FailoverGroupsApi.md#ListFailoverGroups) | **Get** /V4/FailoverGroups | Get Failover groups
[**ModifyFailoverGroup**](FailoverGroupsApi.md#ModifyFailoverGroup) | **Put** /V4/FailoverGroups/{failoverGroupId} | Modify the failover group
[**ModifyFailoverGroupSchedule**](FailoverGroupsApi.md#ModifyFailoverGroupSchedule) | **Put** /V4/FailoverGroups/{failoverGroupId}/Schedules/{scheduleId} | Modify failover group DR operation schedule
[**PlannedFailoverFailoverGroup**](FailoverGroupsApi.md#PlannedFailoverFailoverGroup) | **Put** /V4/FailoverGroups/{failoverGroupId}/Action/{drOperation} | Perform failover group DR operation

# **CreateFailoverGroupSchedule**
> PlanSchedule CreateFailoverGroupSchedule(ctx, failoverGroupId, optional)
Create failover group DR operation schedule

API to create failover group DR operation schedule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **failoverGroupId** | **int32**| Id of the failover group | 
 **optional** | ***FailoverGroupsApiCreateFailoverGroupScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FailoverGroupsApiCreateFailoverGroupScheduleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PlanSchedule**](PlanSchedule.md)|  | 

### Return type

[**PlanSchedule**](PlanSchedule.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFailoverGroup**
> GenericResp DeleteFailoverGroup(ctx, failoverGroupId)
Delete failover group

API to delete failover group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **failoverGroupId** | **int32**| Id of the failover group | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFailoverGroupSchedule**
> GenericResp DeleteFailoverGroupSchedule(ctx, failoverGroupId, scheduleId)
Delete failover group DR operation schedule

API to delete failover group DR operation schedule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **failoverGroupId** | **int32**| Id of the failover group | 
  **scheduleId** | **int32**| ID of the DR operation schedule | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEligibleMachinesFailoverGroup**
> InlineResponse2009 GetEligibleMachinesFailoverGroup(ctx, sourceEntityId, optional)
Get list of eligible machines that can be added to failover groups

API to get list of machines eligible to be added to failover group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceEntityId** | **int32**| The ID of the source entity. This is the ID for client or client group | 
 **optional** | ***FailoverGroupsApiGetEligibleMachinesFailoverGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FailoverGroupsApiGetEligibleMachinesFailoverGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **failoverGroupSource** | [**optional.Interface of FailoverGroupSource**](.md)| The type of failover group source. Default value is &#x27;REPLICATION&#x27; | 
 **sourceEntityType** | [**optional.Interface of FailoverGroupHypervisorType**](.md)| The type of source entity. It can be &#x27;CLIENT&#x27; or &#x27;CLIENT_ENTITY&#x27;. Default value is &#x27;CLIENT&#x27; | 
 **destinationClientId** | **optional.Int32**| The client ID of the destination | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFailoverGroupDetails**
> FailoverGroupDetails GetFailoverGroupDetails(ctx, failoverGroupId)
Get Failover group details

API to fetch failover group information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **failoverGroupId** | **int32**| Id of the failover group | 

### Return type

[**FailoverGroupDetails**](FailoverGroupDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTestFailoverMachinesFailoverGroup**
> InlineResponse2008 GetTestFailoverMachinesFailoverGroup(ctx, failoverGroupId)
Get test failover VMs list

API to get test failover machines for failover group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **failoverGroupId** | **int32**| Id of the failover group | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFailoverGroups**
> InlineResponse2007 ListFailoverGroups(ctx, )
Get Failover groups

This end point return the list of failover groups

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyFailoverGroup**
> GenericResp ModifyFailoverGroup(ctx, failoverGroupId, optional)
Modify the failover group

API to modify failover group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **failoverGroupId** | **int32**| Id of the failover group | 
 **optional** | ***FailoverGroupsApiModifyFailoverGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FailoverGroupsApiModifyFailoverGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ModifyFailoverGroupRequest**](ModifyFailoverGroupRequest.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyFailoverGroupSchedule**
> GenericResp ModifyFailoverGroupSchedule(ctx, failoverGroupId, scheduleId, optional)
Modify failover group DR operation schedule

API to create failover group DR operation schedule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **failoverGroupId** | **int32**| Id of the failover group | 
  **scheduleId** | **int32**| ID of the DR operation schedule | 
 **optional** | ***FailoverGroupsApiModifyFailoverGroupScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FailoverGroupsApiModifyFailoverGroupScheduleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PlanSchedule**](PlanSchedule.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlannedFailoverFailoverGroup**
> JobId PlannedFailoverFailoverGroup(ctx, failoverGroupId, drOperation, optional)
Perform failover group DR operation

API to perform planned failover for failover group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **failoverGroupId** | **int32**| Id of the failover group | 
  **drOperation** | [**DrOperation**](.md)| Name of DR operation. Case insensitive | 
 **optional** | ***FailoverGroupsApiPlannedFailoverFailoverGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FailoverGroupsApiPlannedFailoverFailoverGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipDisableNetworkAdapter** | **optional.Bool**| Whether to skip disabling network adapter in DR job | 
 **replicationId** | **optional.Int32**| Replication ID of particular VM in group to perform DR job on | 

### Return type

[**JobId**](JobId.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

