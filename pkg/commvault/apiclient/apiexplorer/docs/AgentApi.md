# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAgent**](AgentApi.md#DeleteAgent) | **Delete** /V4/Server/{serverId}/Agent/{agentId} | delete an agent from a server
[**DisableBackupAgent**](AgentApi.md#DisableBackupAgent) | **Put** /V4/Server/{serverId}/Agent/{agentId}/Backup/Action/Disable | Disable Backup Property for an Agent
[**DisableRestoreAgent**](AgentApi.md#DisableRestoreAgent) | **Put** /V4/Server/{serverId}/Agent/{agentId}/Restore/Action/Disable | Disable restore property for an agent
[**EnableBackupAgent**](AgentApi.md#EnableBackupAgent) | **Put** /V4/Server/{serverId}/Agent/{agentId}/Backup/Action/Enable | Enable Backup Property for an Agent
[**EnableRestoreAgent**](AgentApi.md#EnableRestoreAgent) | **Put** /V4/Server/{serverId}/Agent/{agentId}/Restore/Action/Enable | Enable restore property for an agent
[**GetConfiguredAgents**](AgentApi.md#GetConfiguredAgents) | **Get** /V4/CommCell/Agents | Gets a list of all the agents configured for a user

# **DeleteAgent**
> GenericResp DeleteAgent(ctx, serverId, agentId)
delete an agent from a server

Used to delete an agent from a server It is expected that the agent has been deconfigured before performing delete operation. But internally if the agent is not deconfigured, then we force deconfigure it to proceed with delete operation Examples of supported agentIds are: 33-File System, 106-Virtual Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| Id of the server to modify | 
  **agentId** | **int32**| Id of the agent to be modified | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableBackupAgent**
> GenericResp DisableBackupAgent(ctx, serverId, agentId, optional)
Disable Backup Property for an Agent

Used to disable backup property for an agent. Types of available agentIds are: 41-Active Directory, 21-AIX File System, 64-Big Data Apps, 134-Cloud Apps, 37-DB2, 103-DB2 MultiNode, 62-DB2 on Unix, 128-Documentum, 90-Domino Mailbox Archiver, 91-DPM, 67-Exchange Compliance Archiver, 53-Exchange Database, 45-Exchange Mailbox, 54-Exchange Mailbox (Classic), 56-Exchange Mailbox Archiver, 82-Exchange PF Archiver, 35-Exchange Public Folder, 73-File Share Archiver, 33-File System, 74-FreeBSD, 71-GroupWise DB, 17-HP-UX Files System, 65-Image Level, 75-Image Level On Unix, 76-Image Level ProxyHost, 87-Image Level ProxyHost on Unix, 3-Informix Database, 29-Linux File System, 89-MS SharePoint Archiver, 104-MySQL, 13-NAS, 83-Netware File Archiver, 12-Netware File System, 10-Novell Directory Services, 124-Object Link, 131-Object Store, 86-OES File System on Linux, 22-Oracle, 80-Oracle RAC, 130-Other External Agent, 125-PostgreSQL, 38-Proxy Client File System, 87-ProxyHost on Unix, 61-SAP for Oracle, 135-SAP HANA, 78-SharePoint Server, 20-Solaris 64bit File System, 19-Solaris File System, 81-SQL Server, 5-Sybase Database, 66-Unix File Archiver, 36-Unix Tru64 64-bit File System, 106-Virtual Server, 58- Windows File Archiver

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| Id of the server to modify | 
  **agentId** | **int32**| Id of the agent to be modified | 
 **optional** | ***AgentApiDisableBackupAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentApiDisableBackupAgentOpts struct
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

# **DisableRestoreAgent**
> GenericResp DisableRestoreAgent(ctx, serverId, agentId, optional)
Disable restore property for an agent

Used to disable restore property for an agent Types of available agentIds are: 41-Active Directory, 21-AIX File System, 64-Big Data Apps, 134-Cloud Apps, 37-DB2, 103-DB2 MultiNode, 62-DB2 on Unix, 128-Documentum, 90-Domino Mailbox Archiver, 91-DPM, 67-Exchange Compliance Archiver, 53-Exchange Database, 45-Exchange Mailbox, 54-Exchange Mailbox (Classic), 56-Exchange Mailbox Archiver, 82-Exchange PF Archiver, 35-Exchange Public Folder, 73-File Share Archiver, 33-File System, 74-FreeBSD, 71-GroupWise DB, 17-HP-UX Files System, 65-Image Level, 75-Image Level On Unix, 76-Image Level ProxyHost, 87-Image Level ProxyHost on Unix, 3-Informix Database, 29-Linux File System, 89-MS SharePoint Archiver, 104-MySQL, 13-NAS, 83-Netware File Archiver, 12-Netware File System, 10-Novell Directory Services, 124-Object Link, 131-Object Store, 86-OES File System on Linux, 22-Oracle, 80-Oracle RAC, 130-Other External Agent, 125-PostgreSQL, 38-Proxy Client File System, 87-ProxyHost on Unix, 61-SAP for Oracle, 135-SAP HANA, 78-SharePoint Server, 20-Solaris 64bit File System, 19-Solaris File System, 81-SQL Server, 5-Sybase Database, 66-Unix File Archiver, 36-Unix Tru64 64-bit File System, 106-Virtual Server, 58- Windows File Archiver

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| Id of the server to modify | 
  **agentId** | **int32**| Id of the agent to be modified | 
 **optional** | ***AgentApiDisableRestoreAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentApiDisableRestoreAgentOpts struct
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

# **EnableBackupAgent**
> GenericResp EnableBackupAgent(ctx, serverId, agentId)
Enable Backup Property for an Agent

Used to enable backup property for an agent Types of available agentIds are: 41-Active Directory, 21-AIX File System, 64-Big Data Apps, 134-Cloud Apps, 37-DB2, 103-DB2 MultiNode, 62-DB2 on Unix, 128-Documentum, 90-Domino Mailbox Archiver, 91-DPM, 67-Exchange Compliance Archiver, 53-Exchange Database, 45-Exchange Mailbox, 54-Exchange Mailbox (Classic), 56-Exchange Mailbox Archiver, 82-Exchange PF Archiver, 35-Exchange Public Folder, 73-File Share Archiver, 33-File System, 74-FreeBSD, 71-GroupWise DB, 17-HP-UX Files System, 65-Image Level, 75-Image Level On Unix, 76-Image Level ProxyHost, 87-Image Level ProxyHost on Unix, 3-Informix Database, 29-Linux File System, 89-MS SharePoint Archiver, 104-MySQL, 13-NAS, 83-Netware File Archiver, 12-Netware File System, 10-Novell Directory Services, 124-Object Link, 131-Object Store, 86-OES File System on Linux, 22-Oracle, 80-Oracle RAC, 130-Other External Agent, 125-PostgreSQL, 38-Proxy Client File System, 87-ProxyHost on Unix, 61-SAP for Oracle, 135-SAP HANA, 78-SharePoint Server, 20-Solaris 64bit File System, 19-Solaris File System, 81-SQL Server, 5-Sybase Database, 66-Unix File Archiver, 36-Unix Tru64 64-bit File System, 106-Virtual Server, 58- Windows File Archiver

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| Id of the server to modify | 
  **agentId** | **int32**| Id of the agent to be modified | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableRestoreAgent**
> GenericResp EnableRestoreAgent(ctx, serverId, agentId)
Enable restore property for an agent

Used to enable restore property for an agent Types of available agentIds are: 41-Active Directory, 21-AIX File System, 64-Big Data Apps, 134-Cloud Apps, 37-DB2, 103-DB2 MultiNode, 62-DB2 on Unix, 128-Documentum, 90-Domino Mailbox Archiver, 91-DPM, 67-Exchange Compliance Archiver, 53-Exchange Database, 45-Exchange Mailbox, 54-Exchange Mailbox (Classic), 56-Exchange Mailbox Archiver, 82-Exchange PF Archiver, 35-Exchange Public Folder, 73-File Share Archiver, 33-File System, 74-FreeBSD, 71-GroupWise DB, 17-HP-UX Files System, 65-Image Level, 75-Image Level On Unix, 76-Image Level ProxyHost, 87-Image Level ProxyHost on Unix, 3-Informix Database, 29-Linux File System, 89-MS SharePoint Archiver, 104-MySQL, 13-NAS, 83-Netware File Archiver, 12-Netware File System, 10-Novell Directory Services, 124-Object Link, 131-Object Store, 86-OES File System on Linux, 22-Oracle, 80-Oracle RAC, 130-Other External Agent, 125-PostgreSQL, 38-Proxy Client File System, 87-ProxyHost on Unix, 61-SAP for Oracle, 135-SAP HANA, 78-SharePoint Server, 20-Solaris 64bit File System, 19-Solaris File System, 81-SQL Server, 5-Sybase Database, 66-Unix File Archiver, 36-Unix Tru64 64-bit File System, 106-Virtual Server, 58- Windows File Archiver

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| Id of the server to modify | 
  **agentId** | **int32**| Id of the agent to be modified | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfiguredAgents**
> GetAllApplicableAgentResp GetConfiguredAgents(ctx, )
Gets a list of all the agents configured for a user

To get a list of all the agents configured for a user

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetAllApplicableAgentResp**](GetAllApplicableAgentResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

