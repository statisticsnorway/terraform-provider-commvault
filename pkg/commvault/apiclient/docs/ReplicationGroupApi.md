# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReplicationGroup**](ReplicationGroupApi.md#CreateReplicationGroup) | **Post** /V4/ReplicationGroup | Create replication group
[**GetReplicationGroupDetails**](ReplicationGroupApi.md#GetReplicationGroupDetails) | **Get** /V4/ReplicationGroup/{replicationGroupId} | Get replication group details
[**GetReplicationGroups**](ReplicationGroupApi.md#GetReplicationGroups) | **Get** /V4/ReplicationGroup | Get list of replication groups
[**ModifyReplicationGroup**](ReplicationGroupApi.md#ModifyReplicationGroup) | **Put** /V4/ReplicationGroup/{replicationGroupId} | Modify ReplicationGroup

# **CreateReplicationGroup**
> IdName CreateReplicationGroup(ctx, optional)
Create replication group

Create a replication group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReplicationGroupApiCreateReplicationGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationGroupApiCreateReplicationGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of V4ReplicationGroupBody**](V4ReplicationGroupBody.md)| Provide the options for one of the vendors to create a replication group | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReplicationGroupDetails**
> InlineResponse2006 GetReplicationGroupDetails(ctx, replicationGroupId, optional)
Get replication group details

Get details of a replication group based on replicationGroupId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **replicationGroupId** | **string**|  | 
 **optional** | ***ReplicationGroupApiGetReplicationGroupDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationGroupApiGetReplicationGroupDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **viewReplicationOptions** | **optional.Bool**| Set to true if want to show replication options for a continuous replication group | 
 **overrideReplicationOptions** | **optional.Bool**| Set to true if you want to see override replication options for a periodic replication group | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReplicationGroups**
> ReplicationGroupListResp GetReplicationGroups(ctx, )
Get list of replication groups

Get all replication groups

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ReplicationGroupListResp**](ReplicationGroupListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyReplicationGroup**
> GenericResp ModifyReplicationGroup(ctx, replicationGroupId, optional)
Modify ReplicationGroup

Modify the properties of an existing replication group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **replicationGroupId** | **string**|  | 
 **optional** | ***ReplicationGroupApiModifyReplicationGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationGroupApiModifyReplicationGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateReplicationGroup**](UpdateReplicationGroup.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

