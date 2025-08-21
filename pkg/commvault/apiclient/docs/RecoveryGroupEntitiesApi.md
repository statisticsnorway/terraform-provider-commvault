# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddorUpdateEntitySettings**](RecoveryGroupEntitiesApi.md#AddorUpdateEntitySettings) | **Put** /V4/EntitySettings | Add or Update Entity Settings
[**CreateRecoveryEntities**](RecoveryGroupEntitiesApi.md#CreateRecoveryEntities) | **Post** /RecoveryGroup/{recoveryGroupId}/Entity | Create entities in a recovery group
[**DeleteRecoveryEntities**](RecoveryGroupEntitiesApi.md#DeleteRecoveryEntities) | **Delete** /RecoveryGroup/{recoveryGroupId}/Entity | Delete recovery entities for a group
[**GetRecoveryEntity**](RecoveryGroupEntitiesApi.md#GetRecoveryEntity) | **Get** /RecoveryEntity/{recoveryEntityId} | Get the details of a recovery entity
[**UpdateRecoveryEntities**](RecoveryGroupEntitiesApi.md#UpdateRecoveryEntities) | **Put** /RecoveryGroup/{recoveryGroupId}/Entity | Update recovery entities in a recovery group
[**UpdateRecoveryEntity**](RecoveryGroupEntitiesApi.md#UpdateRecoveryEntity) | **Put** /RecoveryGroup/{recoveryGroupId}/Entity/{entityId} | Update recovery entity in a recovery group

# **AddorUpdateEntitySettings**
> GenericResp AddorUpdateEntitySettings(ctx, optional)
Add or Update Entity Settings

This API allows you to add or update settings for linked entities, such as servers or server groups. Use this endpoint to modify the default behavior of these entities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecoveryGroupEntitiesApiAddorUpdateEntitySettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecoveryGroupEntitiesApiAddorUpdateEntitySettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ModifyEntitySettings**](ModifyEntitySettings.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRecoveryEntities**
> GenericResp CreateRecoveryEntities(ctx, recoveryGroupId, optional)
Create entities in a recovery group

Create entities in the recovery group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recoveryGroupId** | **int32**| Recovery group Id | 
 **optional** | ***RecoveryGroupEntitiesApiCreateRecoveryEntitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecoveryGroupEntitiesApiCreateRecoveryEntitiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AddUpdateRecoveryEntityRequest**](AddUpdateRecoveryEntityRequest.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRecoveryEntities**
> GenericResp DeleteRecoveryEntities(ctx, recoveryGroupId, optional)
Delete recovery entities for a group

Delete recovery entities for a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recoveryGroupId** | **int32**| Recovery group Id | 
 **optional** | ***RecoveryGroupEntitiesApiDeleteRecoveryEntitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecoveryGroupEntitiesApiDeleteRecoveryEntitiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DeleteRecoveryEntityReq**](DeleteRecoveryEntityReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRecoveryEntity**
> RecoveryEntity GetRecoveryEntity(ctx, recoveryEntityId)
Get the details of a recovery entity

Get the details of a recovery entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recoveryEntityId** | **int32**| Recovery entity Id | 

### Return type

[**RecoveryEntity**](RecoveryEntity.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRecoveryEntities**
> GenericResp UpdateRecoveryEntities(ctx, recoveryGroupId, optional)
Update recovery entities in a recovery group

Updating recovery entities in the recovery group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recoveryGroupId** | **int32**| Recovery group Id | 
 **optional** | ***RecoveryGroupEntitiesApiUpdateRecoveryEntitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecoveryGroupEntitiesApiUpdateRecoveryEntitiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AddUpdateRecoveryEntityRequest**](AddUpdateRecoveryEntityRequest.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRecoveryEntity**
> GenericResp UpdateRecoveryEntity(ctx, recoveryGroupId, entityId, optional)
Update recovery entity in a recovery group

Updating the entity in the recovery group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recoveryGroupId** | **int32**| Recovery group Id | 
  **entityId** | **int32**| Recovery entity Id | 
 **optional** | ***RecoveryGroupEntitiesApiUpdateRecoveryEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecoveryGroupEntitiesApiUpdateRecoveryEntityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of RecoveryEntity**](RecoveryEntity.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

