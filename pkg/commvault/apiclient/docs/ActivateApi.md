# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActivateEntityAuditList**](ActivateApi.md#GetActivateEntityAuditList) | **Get** /V4/Activate/Entity/{entityType}/{entityId}/Audit | Get Activate Entity Audit

# **GetActivateEntityAuditList**
> EntityAuditList GetActivateEntityAuditList(ctx, entityType, entityId)
Get Activate Entity Audit

Retrieve the audit details for activate entities like clients, data sources, requests. Available entity types: 3 - Client 132 - Data source 9515 - Request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityType** | **int32**| Type of the entity | 
  **entityId** | **int32**| Unique id for the entity | 

### Return type

[**EntityAuditList**](EntityAuditList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

