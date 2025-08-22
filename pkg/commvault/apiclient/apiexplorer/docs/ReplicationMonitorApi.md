# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetArrayReplicationMonitor**](ReplicationMonitorApi.md#GetArrayReplicationMonitor) | **Get** /V4/ArrayReplicationMonitor | Get replication monitor information for array replication
[**GetReplicationMonitorPairId**](ReplicationMonitorApi.md#GetReplicationMonitorPairId) | **Get** /V4/ReplicationMonitor/{pairId} | Get continuous replication pair details
[**GetTestFailoverMachines**](ReplicationMonitorApi.md#GetTestFailoverMachines) | **Get** /V4/ReplicationMonitor/{replicationPairId}/testFailoverMachines | Get testFailover machines from replication monitor

# **GetArrayReplicationMonitor**
> ArrayReplicationMonitorResp GetArrayReplicationMonitor(ctx, optional)
Get replication monitor information for array replication

Get Array Replication Pair information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReplicationMonitorApiGetArrayReplicationMonitorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationMonitorApiGetArrayReplicationMonitorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **failoverGroupId** | **optional.Int32**| The ID of the failover group associated to array replication pairs | 
 **replicationId** | **optional.Int32**| The ID of the array replication pair | 

### Return type

[**ArrayReplicationMonitorResp**](ArrayReplicationMonitorResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReplicationMonitorPairId**
> ReplicationMonitorDetail GetReplicationMonitorPairId(ctx, pairId, optional)
Get continuous replication pair details

Get Continuous Replication Pair Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pairId** | **int32**|  | 
 **optional** | ***ReplicationMonitorApiGetReplicationMonitorPairIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationMonitorApiGetReplicationMonitorPairIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **additionalProperties** | **optional.Bool**| To get extra meta data details for the api | 

### Return type

[**ReplicationMonitorDetail**](ReplicationMonitorDetail.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTestFailoverMachines**
> InlineResponse20021 GetTestFailoverMachines(ctx, replicationPairId)
Get testFailover machines from replication monitor

View all TestFailover Machines of the replication pair using replicationPairId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **replicationPairId** | **int32**| Replication pair id of the Pair | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

