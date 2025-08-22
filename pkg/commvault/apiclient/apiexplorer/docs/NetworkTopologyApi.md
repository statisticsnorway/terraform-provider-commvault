# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEFirewallTopology**](NetworkTopologyApi.md#DELETEFirewallTopology) | **Delete** /V4/NetworkTopology/{topologyId} | Delete a Network Topology
[**GETFirewallTopology**](NetworkTopologyApi.md#GETFirewallTopology) | **Get** /V4/NetworkTopology | Get Network Topologies
[**GETFirewallTopologyDetails**](NetworkTopologyApi.md#GETFirewallTopologyDetails) | **Get** /V4/NetworkTopology/{topologyId} | Get Network Topology Details
[**POSTFirewallTopology**](NetworkTopologyApi.md#POSTFirewallTopology) | **Post** /V4/NetworkTopology | Create a Network Topology
[**PUTFirewallTopology**](NetworkTopologyApi.md#PUTFirewallTopology) | **Put** /V4/NetworkTopology/{topologyId} | Modify Network Topology details

# **DELETEFirewallTopology**
> GenericResp DELETEFirewallTopology(ctx, topologyId)
Delete a Network Topology

This endpoint is used to delete network topology.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **topologyId** | **string**|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETFirewallTopology**
> InlineResponse20014 GETFirewallTopology(ctx, )
Get Network Topologies

This endpoint is used to return the list of network topology.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETFirewallTopologyDetails**
> InlineResponse20015 GETFirewallTopologyDetails(ctx, topologyId)
Get Network Topology Details

This endpoint is used to return the detail of network topology.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **topologyId** | **string**|  | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTFirewallTopology**
> FirewallTopologyCreateResp POSTFirewallTopology(ctx, optional)
Create a Network Topology

This endpoint is used to create network topology.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworkTopologyApiPOSTFirewallTopologyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkTopologyApiPOSTFirewallTopologyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of FirewallTopologyCreateReq**](FirewallTopologyCreateReq.md)|  | 

### Return type

[**FirewallTopologyCreateResp**](FirewallTopologyCreateResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTFirewallTopology**
> GenericResp PUTFirewallTopology(ctx, topologyId, optional)
Modify Network Topology details

This endpoint is used to edit network topology.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **topologyId** | **string**|  | 
 **optional** | ***NetworkTopologyApiPUTFirewallTopologyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkTopologyApiPUTFirewallTopologyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of FirewallTopologyCreateReq**](FirewallTopologyCreateReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

