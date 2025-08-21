# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMediaAgentRole**](ServersApi.md#AddMediaAgentRole) | **Put** /V4/Servers/{serverId}/Role/MediaAgent | Add Media Agent License for a Client
[**ChangeServerDomain**](ServersApi.md#ChangeServerDomain) | **Put** /V4/Servers/Domain | Change domain for servers
[**DisableBackupServer**](ServersApi.md#DisableBackupServer) | **Put** /V4/Server/{serverId}/Backup/Action/Disable | Disable Backups for a server
[**DisableDataAgingServer**](ServersApi.md#DisableDataAgingServer) | **Put** /V4/Server/{serverId}/DataAging/Action/Disable | Disable Data Aging for a Server
[**DisableRestoreServer**](ServersApi.md#DisableRestoreServer) | **Put** /V4/Server/{serverId}/Restore/Action/Disable | Disable Restores for a Server
[**EnableBackupServer**](ServersApi.md#EnableBackupServer) | **Put** /V4/Server/{serverId}/Backup/Action/Enable | Enable backup property for a server
[**EnableDataAgingServer**](ServersApi.md#EnableDataAgingServer) | **Put** /V4/Server/{serverId}/DataAging/Action/Enable | Enable Data Aging for a Server
[**EnableRestoreServer**](ServersApi.md#EnableRestoreServer) | **Put** /V4/Server/{serverId}/Restore/Action/Enable | Enable Restores for a Server
[**GetPlatformUpgradeOptions**](ServersApi.md#GetPlatformUpgradeOptions) | **Get** /V4/PlatformUpgrade/Options | Get Platform Upgrade Options
[**GetServerDomain**](ServersApi.md#GetServerDomain) | **Get** /V4/Servers/Domain | Get domain for servers
[**GetServers**](ServersApi.md#GetServers) | **Get** /V4/Servers | Get Servers
[**RetireClientPackages**](ServersApi.md#RetireClientPackages) | **Put** /V4/Client/{clientId}/Retire | Retire ClientPackages
[**RetireServerGroup**](ServersApi.md#RetireServerGroup) | **Put** /V4/ServerGroup/{serverGroupId}/Retire | Retire Servers of a Server group
[**RetireServers**](ServersApi.md#RetireServers) | **Put** /V4/Servers/Retire | Retire Multiple Servers

# **AddMediaAgentRole**
> GenericResp AddMediaAgentRole(ctx, serverId)
Add Media Agent License for a Client

Put call to consume ma license for a client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| serverId of the client for which the MA role needs to be added | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeServerDomain**
> GenericResp ChangeServerDomain(ctx, optional)
Change domain for servers

Change domain for list of servers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServersApiChangeServerDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServersApiChangeServerDomainOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ChangeServerDomainReq**](ChangeServerDomainReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableBackupServer**
> GenericResp DisableBackupServer(ctx, serverId, optional)
Disable Backups for a server

Used to disable backup property for server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| Id of the Server whose details have to be modified | 
 **optional** | ***ServersApiDisableBackupServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServersApiDisableBackupServerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableAfterADelay** | **optional.Int32**| Provide UTC time in unix format. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableDataAgingServer**
> GenericResp DisableDataAgingServer(ctx, serverId, optional)
Disable Data Aging for a Server

Used to disable Data Aging property for server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| Id of the Server whose details have to be modified | 
 **optional** | ***ServersApiDisableDataAgingServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServersApiDisableDataAgingServerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableAfterADelay** | **optional.Int32**| Provide UTC time in unix format. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableRestoreServer**
> GenericResp DisableRestoreServer(ctx, serverId, optional)
Disable Restores for a Server

Used to disable restore property for server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| Id of the Server whose details have to be modified | 
 **optional** | ***ServersApiDisableRestoreServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServersApiDisableRestoreServerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableAfterADelay** | **optional.Int32**| Provide UTC time in unix format. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableBackupServer**
> GenericResp EnableBackupServer(ctx, serverId)
Enable backup property for a server

Used to enable backup property for Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| Id of the Server whose details have to be modified | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableDataAgingServer**
> GenericResp EnableDataAgingServer(ctx, serverId)
Enable Data Aging for a Server

Used to enable Data Aging property for Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| Id of the Server whose details have to be modified | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableRestoreServer**
> GenericResp EnableRestoreServer(ctx, serverId)
Enable Restores for a Server

Used to enable restore property for Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| Id of the Server whose details have to be modified | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlatformUpgradeOptions**
> GetSetWindowsMssqlPatchingMsg GetPlatformUpgradeOptions(ctx, optional)
Get Platform Upgrade Options

API to get eligibility for install job option at different entity level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServersApiGetPlatformUpgradeOptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServersApiGetPlatformUpgradeOptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **option** | **optional.Int32**| Specify the entity level for which to fetch install jobs options. Accepted values [ 2 &#x3D; Client, 3 &#x3D; Client Group, 4 &#x3D; Commcell] | 
 **clientId** | **optional.Int32**| Id of the client or client group for which to fetch install job options | 

### Return type

[**GetSetWindowsMssqlPatchingMsg**](GetSetWindowsMSSQLPatchingMsg.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerDomain**
> GetServerDomainResp GetServerDomain(ctx, )
Get domain for servers

Get list of servers and respective domains

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetServerDomainResp**](GetServerDomainResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServers**
> ServersListResp GetServers(ctx, optional)
Get Servers

This end point returns the list of servers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServersApiGetServersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServersApiGetServersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showOnlyInfrastructureMachines** | **optional.Int32**| Returns only infrastructure machines if value is 1 if 0, it returns all servers, default value is 1 | [default to 1]

### Return type

[**ServersListResp**](ServersListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetireClientPackages**
> RetireClientResponse RetireClientPackages(ctx, clientId, optional)
Retire ClientPackages

Selectively uninstall packages from given client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **int32**|  | 
 **optional** | ***ServersApiRetireClientPackagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServersApiRetireClientPackagesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RetireClientRequest**](RetireClientRequest.md)|  | 

### Return type

[**RetireClientResponse**](RetireClientResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetireServerGroup**
> RetireServerGroupResp RetireServerGroup(ctx, serverGroupId)
Retire Servers of a Server group

Retire servers in given server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverGroupId** | **int32**| Id of the serverGroup in which servers need to be retired | 

### Return type

[**RetireServerGroupResp**](RetireServerGroupResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetireServers**
> RetireServersResponse RetireServers(ctx, optional)
Retire Multiple Servers

Retire multiple servers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServersApiRetireServersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServersApiRetireServersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RetireServersRequest**](RetireServersRequest.md)|  | 

### Return type

[**RetireServersResponse**](RetireServersResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

