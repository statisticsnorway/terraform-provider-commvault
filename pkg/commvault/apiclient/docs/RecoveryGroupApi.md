# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddorUpdateGlobalSettings**](RecoveryGroupApi.md#AddorUpdateGlobalSettings) | **Put** /V4/GlobalSettings | Add or Update Global Settings
[**CreateRecoveryGroup**](RecoveryGroupApi.md#CreateRecoveryGroup) | **Post** /RecoveryGroup | Create Recovery Group
[**DeleteRecoveryGroup**](RecoveryGroupApi.md#DeleteRecoveryGroup) | **Delete** /RecoveryGroup/{recoveryGroupId} | Delete recovery group
[**GetLicenseUsage**](RecoveryGroupApi.md#GetLicenseUsage) | **Get** /LicenseUsage | Get license usage details for a license type
[**GetRecoveryGroup**](RecoveryGroupApi.md#GetRecoveryGroup) | **Get** /RecoveryGroup/{recoveryGroupId} | Get details for recovery group
[**GetRecoveryGroups**](RecoveryGroupApi.md#GetRecoveryGroups) | **Get** /RecoveryGroups | Get the list of Recovery groups
[**RecoveryGroupCleanupRecovery**](RecoveryGroupApi.md#RecoveryGroupCleanupRecovery) | **Post** /RecoveryGroup/CleanupRecovery | Launching cleanup recovery action for a recovery group
[**RecoveryGroupRecover**](RecoveryGroupApi.md#RecoveryGroupRecover) | **Post** /RecoveryGroup/{recoveryGroupId}/Recover | Launching recovery action for a recovery group
[**UpdateRecoveryGroup**](RecoveryGroupApi.md#UpdateRecoveryGroup) | **Put** /RecoveryGroup/{recoveryGroupId} | Update Recovery Group

# **AddorUpdateGlobalSettings**
> GenericResp AddorUpdateGlobalSettings(ctx, optional)
Add or Update Global Settings

This API allows you to add new global settings or update existing ones, overriding the default configurations to customize the system's behavior.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecoveryGroupApiAddorUpdateGlobalSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecoveryGroupApiAddorUpdateGlobalSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ModifyGlobalSettings**](ModifyGlobalSettings.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRecoveryGroup**
> CreateRecoveryGroupResp CreateRecoveryGroup(ctx, optional)
Create Recovery Group

Create Recovery group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecoveryGroupApiCreateRecoveryGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecoveryGroupApiCreateRecoveryGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RecoveryGroup**](RecoveryGroup.md)|  | 

### Return type

[**CreateRecoveryGroupResp**](CreateRecoveryGroupResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRecoveryGroup**
> CreateRecoveryGroupResp DeleteRecoveryGroup(ctx, recoveryGroupId)
Delete recovery group

Delete recovery group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recoveryGroupId** | **int32**| Recovery group Id | 

### Return type

[**CreateRecoveryGroupResp**](CreateRecoveryGroupResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicenseUsage**
> LicenseUsageDetails GetLicenseUsage(ctx, licenseType)
Get license usage details for a license type

Get license usage details for a license type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **licenseType** | [**LicenseType**](.md)| License type for which usage details are required | 

### Return type

[**LicenseUsageDetails**](LicenseUsageDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRecoveryGroup**
> GetRecoveryGroupDetailsResp GetRecoveryGroup(ctx, recoveryGroupId, optional)
Get details for recovery group

Get details for recovery group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recoveryGroupId** | **int32**| Recovery group Id | 
 **optional** | ***RecoveryGroupApiGetRecoveryGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecoveryGroupApiGetRecoveryGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getEntityDetails** | **optional.Bool**| Get the entity details with recovery group details | [default to false]

### Return type

[**GetRecoveryGroupDetailsResp**](GetRecoveryGroupDetailsResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRecoveryGroups**
> RecoveryGroupListResp GetRecoveryGroups(ctx, )
Get the list of Recovery groups

Get the list of Recovery groups

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RecoveryGroupListResp**](RecoveryGroupListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecoveryGroupCleanupRecovery**
> JobIdResponse RecoveryGroupCleanupRecovery(ctx, optional)
Launching cleanup recovery action for a recovery group

Cleanup recovery action will delete the recovered entities for a recovery group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecoveryGroupApiRecoveryGroupCleanupRecoveryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecoveryGroupApiRecoveryGroupCleanupRecoveryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RecoveryActionRequest**](RecoveryActionRequest.md)|  | 

### Return type

[**JobIdResponse**](JobIdResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecoveryGroupRecover**
> JobIdResponse RecoveryGroupRecover(ctx, recoveryGroupId, optional)
Launching recovery action for a recovery group

Launching recovery action for a recovery group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recoveryGroupId** | **string**| Recovery group Id | 
 **optional** | ***RecoveryGroupApiRecoveryGroupRecoverOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecoveryGroupApiRecoveryGroupRecoverOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RecoveryActionRequest**](RecoveryActionRequest.md)|  | 

### Return type

[**JobIdResponse**](JobIdResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRecoveryGroup**
> CreateRecoveryGroupResp UpdateRecoveryGroup(ctx, recoveryGroupId, optional)
Update Recovery Group

Update Recovery Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recoveryGroupId** | **int32**| Recovery group Id | 
 **optional** | ***RecoveryGroupApiUpdateRecoveryGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecoveryGroupApiUpdateRecoveryGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RecoveryGroup**](RecoveryGroup.md)|  | 

### Return type

[**CreateRecoveryGroupResp**](CreateRecoveryGroupResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

