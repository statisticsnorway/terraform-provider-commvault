# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculateTotalCostOfOwnership**](CyberResilienceApi.md#CalculateTotalCostOfOwnership) | **Post** /V4/CyberResilience/TCO | Calculate Total Cost Of Ownership
[**ConfigureCyberResilience**](CyberResilienceApi.md#ConfigureCyberResilience) | **Post** /V4/CyberResilience | Configure Cyber Resilience
[**GetCyberResilienceCredentials**](CyberResilienceApi.md#GetCyberResilienceCredentials) | **Get** /V4/CyberResilience/Credentials | Get Cyber Resilience Credentials
[**GetCyberResilienceSyncStatus**](CyberResilienceApi.md#GetCyberResilienceSyncStatus) | **Get** /V4/CyberResilience/Credentials/Sync/Status | Get Cyber Resilience Sync Status
[**ManageCyberResilienceCredentials**](CyberResilienceApi.md#ManageCyberResilienceCredentials) | **Put** /V4/CyberResilience/Credentials | Manage Cyber Resilience Credentials
[**SyncCyberResilienceCredentials**](CyberResilienceApi.md#SyncCyberResilienceCredentials) | **Put** /V4/CyberResilience/Credentials/Sync | Sync Cyber Resilience Credentials

# **CalculateTotalCostOfOwnership**
> TcoCalculateResponse CalculateTotalCostOfOwnership(ctx, optional)
Calculate Total Cost Of Ownership

Calculate the Total Cost Of Ownership for AZURE or AWS in comparison to Commvault Cloud

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CyberResilienceApiCalculateTotalCostOfOwnershipOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CyberResilienceApiCalculateTotalCostOfOwnershipOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TcoCalculateRequest**](TcoCalculateRequest.md)|  | 

### Return type

[**TcoCalculateResponse**](TCOCalculateResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigureCyberResilience**
> CyberResilienceGenericResponse ConfigureCyberResilience(ctx, optional)
Configure Cyber Resilience

Configures Cyber Resilience in the environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CyberResilienceApiConfigureCyberResilienceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CyberResilienceApiConfigureCyberResilienceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CyberResilienceConfigureRequest**](CyberResilienceConfigureRequest.md)|  | 

### Return type

[**CyberResilienceGenericResponse**](CyberResilienceGenericResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCyberResilienceCredentials**
> CyberResilienceCredentials GetCyberResilienceCredentials(ctx, optional)
Get Cyber Resilience Credentials

Retrieve the list of credentials for Cyber Resilience

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CyberResilienceApiGetCyberResilienceCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CyberResilienceApiGetCyberResilienceCredentialsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **retrieveConfiguredOnly** | **optional.Bool**| If set to true, retrieve only the credentials that are configured | [default to false]

### Return type

[**CyberResilienceCredentials**](CyberResilienceCredentials.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCyberResilienceSyncStatus**
> CyberResilienceSyncStatus GetCyberResilienceSyncStatus(ctx, )
Get Cyber Resilience Sync Status

Retrieve the sync status for configured credentials

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CyberResilienceSyncStatus**](CyberResilienceSyncStatus.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ManageCyberResilienceCredentials**
> CyberResilienceGenericResponse ManageCyberResilienceCredentials(ctx, optional)
Manage Cyber Resilience Credentials

Manage Credentials configured for Cyber Resilience

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CyberResilienceApiManageCyberResilienceCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CyberResilienceApiManageCyberResilienceCredentialsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CyberResilienceManageCredentials**](CyberResilienceManageCredentials.md)|  | 

### Return type

[**CyberResilienceGenericResponse**](CyberResilienceGenericResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncCyberResilienceCredentials**
> CyberResilienceGenericResponse SyncCyberResilienceCredentials(ctx, )
Sync Cyber Resilience Credentials

Sync the Credentials configured for Cyber Resilience with latest information

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CyberResilienceGenericResponse**](CyberResilienceGenericResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

