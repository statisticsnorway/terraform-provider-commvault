# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLDAP**](IdentityServersApi.md#CreateLDAP) | **Post** /V4/LDAP | Create LDAP/Active directory for user authentication
[**CreateSAMLApp**](IdentityServersApi.md#CreateSAMLApp) | **Post** /V4/SAML | Create SAML App
[**DeleteADLDAP**](IdentityServersApi.md#DeleteADLDAP) | **Delete** /V4/LDAP/{domainId} | Delete AD/LDAP
[**GetADLDAPDetails**](IdentityServersApi.md#GetADLDAPDetails) | **Get** /V4/LDAP/{domainId} | Get AD/LDAP Details
[**GetIdentityServers**](IdentityServersApi.md#GetIdentityServers) | **Get** /V4/IdentityServers | Get Identity Servers
[**GetSAMLApp**](IdentityServersApi.md#GetSAMLApp) | **Get** /V4/SAML/{name} | Get SAML Server
[**UpdateADLDAP**](IdentityServersApi.md#UpdateADLDAP) | **Put** /V4/LDAP/{domainId} | Update AD/LDAP
[**UpdateSAMLApp**](IdentityServersApi.md#UpdateSAMLApp) | **Put** /V4/SAML/{name} | Modify SAML App Details

# **CreateLDAP**
> GenericResp CreateLDAP(ctx, optional)
Create LDAP/Active directory for user authentication

Create LDAP/Active directory for user authentication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IdentityServersApiCreateLDAPOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IdentityServersApiCreateLDAPOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of V4LdapBody**](V4LdapBody.md)| LDAP body request | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSAMLApp**
> GenericResp CreateSAMLApp(ctx, optional)
Create SAML App

Creates SAML app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IdentityServersApiCreateSAMLAppOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IdentityServersApiCreateSAMLAppOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SamlReq**](SamlReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteADLDAP**
> GenericResp DeleteADLDAP(ctx, domainId, optional)
Delete AD/LDAP

Delete an AD/LDAP domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **int32**| ID of the AD/LDAP domain | 
 **optional** | ***IdentityServersApiDeleteADLDAPOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IdentityServersApiDeleteADLDAPOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferToUserId** | **optional.Int32**| Optionally transfer the ownership to user | 
 **transferToUserGroupId** | **optional.Int32**| Optionally transfer the ownership to user group | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetADLDAPDetails**
> AdldapDetails GetADLDAPDetails(ctx, domainId)
Get AD/LDAP Details

Get detail of the AD/LDAP domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **int32**| ID of the AD/LDAP domain | 

### Return type

[**AdldapDetails**](ADLDAPDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdentityServers**
> IdentityServersListResp GetIdentityServers(ctx, )
Get Identity Servers

Get identity servers list

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IdentityServersListResp**](IdentityServersListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSAMLApp**
> Saml GetSAMLApp(ctx, name)
Get SAML Server

Gets details of SAML app based on SAML app name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**Saml**](SAML.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateADLDAP**
> GenericResp UpdateADLDAP(ctx, domainId, optional)
Update AD/LDAP

Update an AD/LDAP domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **int32**| ID of the AD/LDAP domain | 
 **optional** | ***IdentityServersApiUpdateADLDAPOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IdentityServersApiUpdateADLDAPOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateAdldap**](UpdateAdldap.md)| Properties of the domain to be updated | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSAMLApp**
> GenericResp UpdateSAMLApp(ctx, name, optional)
Modify SAML App Details

Updates details of existing SAML app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
 **optional** | ***IdentityServersApiUpdateSAMLAppOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IdentityServersApiUpdateSAMLAppOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SamlUpdate**](SamlUpdate.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

