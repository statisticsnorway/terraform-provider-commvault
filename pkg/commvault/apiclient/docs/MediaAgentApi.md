# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMediaAgent**](MediaAgentApi.md#DeleteMediaAgent) | **Delete** /V4/mediaAgent/{mediaAgentId} | Delete Media Agent
[**GetMediaAgentDDBDisks**](MediaAgentApi.md#GetMediaAgentDDBDisks) | **Get** /V4/mediaAgent/{mediaAgentId}/DDB/Disks | Get MediaAgent DDBDisks
[**GetMediaAgentDetails**](MediaAgentApi.md#GetMediaAgentDetails) | **Get** /V4/mediaAgent/{mediaAgentId} | Get Media Agent Details
[**GetMediaAgents**](MediaAgentApi.md#GetMediaAgents) | **Get** /V4/mediaAgent | Get MediaAgents
[**GetMediaAgentsForDDB**](MediaAgentApi.md#GetMediaAgentsForDDB) | **Get** /V4/DDB/MediaAgents | Get MediaAgents for DDB
[**InstallMediaAgent**](MediaAgentApi.md#InstallMediaAgent) | **Post** /V4/mediaAgent | Create a Media Agent
[**MediaAgentDDBDiskMgmt**](MediaAgentApi.md#MediaAgentDDBDiskMgmt) | **Put** /V4/mediaAgent/{mediaAgentId}/DDB/Disks | Perform DDB disk management operations on MediaAgent
[**ModifyMediaAgent**](MediaAgentApi.md#ModifyMediaAgent) | **Put** /V4/mediaAgent/{mediaAgentId} | Modify Media Agent Details

# **DeleteMediaAgent**
> GenericResp DeleteMediaAgent(ctx, mediaAgentId)
Delete Media Agent

Used to delete a media agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mediaAgentId** | **int32**| Id of the media agent to delete | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMediaAgentDDBDisks**
> MaddbDiskMgmtResp GetMediaAgentDDBDisks(ctx, mediaAgentId)
Get MediaAgent DDBDisks

Fetch the list of DDB disks hosted on the MediaAgent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mediaAgentId** | **int32**| Id of the MediaAgent | 

### Return type

[**MaddbDiskMgmtResp**](MADDBDiskMgmtResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMediaAgentDetails**
> MediaAgent GetMediaAgentDetails(ctx, mediaAgentId, optional)
Get Media Agent Details

Get details of a media agent based on id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mediaAgentId** | **int32**| Id of the Media Agent whose details have to be fetched | 
 **optional** | ***MediaAgentApiGetMediaAgentDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MediaAgentApiGetMediaAgentDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showInheritedAssociation** | **optional.Bool**| Show inherited security association | 

### Return type

[**MediaAgent**](MediaAgent.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMediaAgents**
> MediaAgentListResponse GetMediaAgents(ctx, )
Get MediaAgents

Get All Media Agents

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MediaAgentListResponse**](MediaAgentListResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMediaAgentsForDDB**
> MediaAgentForDdbListResponse GetMediaAgentsForDDB(ctx, optional)
Get MediaAgents for DDB

Get All Media Agents for DDB

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MediaAgentApiGetMediaAgentsForDDBOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MediaAgentApiGetMediaAgentsForDDBOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetchDDBDisks** | **optional.Bool**| If set to true then fetch the list of DDB disks hosted on each MediaAgent | [default to false]

### Return type

[**MediaAgentForDdbListResponse**](MediaAgentForDDBListResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstallMediaAgent**
> JobId InstallMediaAgent(ctx, optional)
Create a Media Agent

To install MediaAgent package on a server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MediaAgentApiInstallMediaAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MediaAgentApiInstallMediaAgentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InstallMediaAgent**](InstallMediaAgent.md)| Server information where you want to install MediaAgent package. | 

### Return type

[**JobId**](JobId.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MediaAgentDDBDiskMgmt**
> MaddbDiskMgmtResp MediaAgentDDBDiskMgmt(ctx, mediaAgentId, optional)
Perform DDB disk management operations on MediaAgent

Perform DDB disk management operations on MediaAgent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mediaAgentId** | **int32**| Id of the MediaAgent | 
 **optional** | ***MediaAgentApiMediaAgentDDBDiskMgmtOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MediaAgentApiMediaAgentDDBDiskMgmtOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MaddbDiskMgmtReq**](MaddbDiskMgmtReq.md)|  | 

### Return type

[**MaddbDiskMgmtResp**](MADDBDiskMgmtResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyMediaAgent**
> GenericResp ModifyMediaAgent(ctx, mediaAgentId, optional)
Modify Media Agent Details

Modify the properties of an existing media agent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mediaAgentId** | **int32**| Id of the mediaAgent to update | 
 **optional** | ***MediaAgentApiModifyMediaAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MediaAgentApiModifyMediaAgentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateMediaAgent**](UpdateMediaAgent.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

