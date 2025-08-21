# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredential**](CredentialManagerApi.md#CreateCredential) | **Post** /V4/Credential | Create Credential
[**DeleteCredential**](CredentialManagerApi.md#DeleteCredential) | **Delete** /V5/Credential/{credentialId} | Delete Credential
[**DeleteCredentialByName**](CredentialManagerApi.md#DeleteCredentialByName) | **Delete** /V4/Credential/{credentialName} | Delete Credential ByName
[**GetCredentialDetails**](CredentialManagerApi.md#GetCredentialDetails) | **Get** /V5/Credential/{credentialId} | Get Credential Details
[**GetCredentialDetailsByName**](CredentialManagerApi.md#GetCredentialDetailsByName) | **Get** /V4/Credential/{credentialName} | Get Credential Details by name
[**GetCredentials**](CredentialManagerApi.md#GetCredentials) | **Get** /V4/Credential | Get list of credentials
[**UpdateCredential**](CredentialManagerApi.md#UpdateCredential) | **Put** /V5/Credential/{credentialId} | Update Credential
[**UpdateCredentialByName**](CredentialManagerApi.md#UpdateCredentialByName) | **Put** /V4/Credential/{credentialName} | Update Credential By Name

# **CreateCredential**
> IdType CreateCredential(ctx, optional)
Create Credential

Create a new credential

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CredentialManagerApiCreateCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialManagerApiCreateCredentialOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of V4CredentialBody**](V4CredentialBody.md)|  | 

### Return type

[**IdType**](IdType.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCredential**
> GenericResp DeleteCredential(ctx, credentialId)
Delete Credential

Delete Credential whose id has been provided

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **credentialId** | **string**|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCredentialByName**
> GenericResp DeleteCredentialByName(ctx, credentialName)
Delete Credential ByName

Delete Credential whose name has been provided

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **credentialName** | **string**|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCredentialDetails**
> InlineResponse20017 GetCredentialDetails(ctx, credentialId)
Get Credential Details

Get details of the credential whose credential id is provided

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **credentialId** | **string**|  | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCredentialDetailsByName**
> InlineResponse20017 GetCredentialDetailsByName(ctx, credentialName)
Get Credential Details by name

Get details of the credential whose credential name is provided

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **credentialName** | **string**|  | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCredentials**
> CredentialManagerListResp GetCredentials(ctx, )
Get list of credentials

List of Credentials

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CredentialManagerListResp**](CredentialManagerListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCredential**
> interface{} UpdateCredential(ctx, credentialId, optional)
Update Credential

Edit credential whose id has been provided by credential owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **credentialId** | **string**|  | 
 **optional** | ***CredentialManagerApiUpdateCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialManagerApiUpdateCredentialOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CredentialCredentialIdBody**](CredentialCredentialIdBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCredentialByName**
> interface{} UpdateCredentialByName(ctx, credentialName, optional)
Update Credential By Name

Edit credential whose name has been provided by credential owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **credentialName** | **string**|  | 
 **optional** | ***CredentialManagerApiUpdateCredentialByNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialManagerApiUpdateCredentialByNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CredentialCredentialNameBody**](CredentialCredentialNameBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

