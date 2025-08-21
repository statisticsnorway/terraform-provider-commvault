# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddorUpdateEntitySettings**](AdditionalSettingsApi.md#AddorUpdateEntitySettings) | **Put** /V4/EntitySettings | Add or Update Entity Settings
[**AddorUpdateGlobalSettings**](AdditionalSettingsApi.md#AddorUpdateGlobalSettings) | **Put** /V4/GlobalSettings | Add or Update Global Settings
[**GetEntitySettings**](AdditionalSettingsApi.md#GetEntitySettings) | **Get** /V4/EntitySettings | Get Entity Settings
[**GetGlobalSettings**](AdditionalSettingsApi.md#GetGlobalSettings) | **Get** /V4/GlobalSettings | Get Global Settings
[**ListAddedEntitySettings**](AdditionalSettingsApi.md#ListAddedEntitySettings) | **Get** /V4/AddedEntitySettings | List added entity settings
[**ListAddedGlobalSettings**](AdditionalSettingsApi.md#ListAddedGlobalSettings) | **Get** /V4/AddedGlobalSettings | Get Added Global Settings
[**ListAvailableSettings**](AdditionalSettingsApi.md#ListAvailableSettings) | **Get** /V4/LookupSettings | Get all available settings

# **AddorUpdateEntitySettings**
> GenericResp AddorUpdateEntitySettings(ctx, optional)
Add or Update Entity Settings

This API allows you to add or update settings for linked entities, such as servers or server groups. Use this endpoint to modify the default behavior of these entities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AdditionalSettingsApiAddorUpdateEntitySettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdditionalSettingsApiAddorUpdateEntitySettingsOpts struct
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

# **AddorUpdateGlobalSettings**
> GenericResp AddorUpdateGlobalSettings(ctx, optional)
Add or Update Global Settings

This API allows you to add new global settings or update existing ones, overriding the default configurations to customize the system's behavior.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AdditionalSettingsApiAddorUpdateGlobalSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdditionalSettingsApiAddorUpdateGlobalSettingsOpts struct
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

# **GetEntitySettings**
> EntitySettingsResponse GetEntitySettings(ctx, )
Get Entity Settings

Get list of entity settings used to modify default behaviour for linked entity like servers or server groups

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EntitySettingsResponse**](EntitySettingsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGlobalSettings**
> GlobalSettingsResponse GetGlobalSettings(ctx, )
Get Global Settings

Get list of global settings used to modify system default behaviour

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GlobalSettingsResponse**](GlobalSettingsResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAddedEntitySettings**
> ListAddedEntitySettings ListAddedEntitySettings(ctx, )
List added entity settings

Returns a list of all entity settings added 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListAddedEntitySettings**](ListAddedEntitySettings.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAddedGlobalSettings**
> ListAddedGlobalSettings ListAddedGlobalSettings(ctx, )
Get Added Global Settings

Get list of global settings used to modify system default behaviour

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListAddedGlobalSettings**](ListAddedGlobalSettings.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAvailableSettings**
> ListAdditionalSettings ListAvailableSettings(ctx, )
Get all available settings

List all available settings that can be added to entity or global

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListAdditionalSettings**](ListAdditionalSettings.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

