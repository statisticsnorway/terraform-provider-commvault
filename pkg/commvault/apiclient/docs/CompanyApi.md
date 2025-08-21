# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCompany**](CompanyApi.md#CreateCompany) | **Post** /V4/company | Create a Company
[**DeleteCompany**](CompanyApi.md#DeleteCompany) | **Delete** /V4/company/{companyId} | Delete a company
[**GetAssociationsForCompany**](CompanyApi.md#GetAssociationsForCompany) | **Get** /V4/Company/{companyId}/associations | API to get list of association for a company
[**GetCompanies**](CompanyApi.md#GetCompanies) | **Get** /V4/company | Get Companies
[**GetCompanyDetails**](CompanyApi.md#GetCompanyDetails) | **Get** /V4/company/{companyId} | Get Company Details
[**GetCompanyLaptopAdmins**](CompanyApi.md#GetCompanyLaptopAdmins) | **Get** /V4/Company/{companyId}/LaptopAdmins | Get a company&#x27;s laptop admins.
[**GetServersForCompany**](CompanyApi.md#GetServersForCompany) | **Get** /V4/Company/{companyId}/Servers |  API to get list of servers with no packages for a company
[**ModifyCompany**](CompanyApi.md#ModifyCompany) | **Put** /V4/company/{companyId} | Update the properties of existing company
[**UpdateCompanyLaptopAdmins**](CompanyApi.md#UpdateCompanyLaptopAdmins) | **Put** /V4/Company/{companyId}/LaptopAdmins | 
[**UpdateCompanyTimezoneForServers**](CompanyApi.md#UpdateCompanyTimezoneForServers) | **Put** /V4/Company/{companyId}/Servers/Timezone | Update the company Timezone for servers

# **CreateCompany**
> IdName CreateCompany(ctx, optional)
Create a Company

Create a Company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CompanyApiCreateCompanyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CompanyApiCreateCompanyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateCompany**](CreateCompany.md)|  | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCompany**
> GenericResp DeleteCompany(ctx, companyId)
Delete a company

Used to delete a company which has been deactivated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | **int32**| Id of the company to delete | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAssociationsForCompany**
> GetAssociationsForCompanyResp GetAssociationsForCompany(ctx, companyId)
API to get list of association for a company

API to get list of association for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | **int32**| Id of the Company whose associations have to be fetched | 

### Return type

[**GetAssociationsForCompanyResp**](GetAssociationsForCompanyResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCompanies**
> CompanyListResponse GetCompanies(ctx, optional)
Get Companies

Get All Companies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CompanyApiGetCompaniesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CompanyApiGetCompaniesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeDeletedCompanies** | **optional.Bool**| If true, companies marked for deletion are included in the response | 

### Return type

[**CompanyListResponse**](CompanyListResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCompanyDetails**
> Company GetCompanyDetails(ctx, companyId, optional)
Get Company Details

Get details of a company based on id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | **int32**| Id of the Company whose details have to be fetched | 
 **optional** | ***CompanyApiGetCompanyDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CompanyApiGetCompanyDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showInheritedAssociation** | **optional.Bool**| Show inherited security association | 

### Return type

[**Company**](Company.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCompanyLaptopAdmins**
> GetLaptopAdminsResp GetCompanyLaptopAdmins(ctx, companyId)
Get a company's laptop admins.

Get a company's laptop admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | **string**|  | 

### Return type

[**GetLaptopAdminsResp**](GetLaptopAdminsResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServersForCompany**
> GetPseudoClientsForCompanyResp GetServersForCompany(ctx, companyId)
 API to get list of servers with no packages for a company

API to get list of servers with no packages for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | **int32**| Id of the Company whose servers(with no packages) have to be fetched | 

### Return type

[**GetPseudoClientsForCompanyResp**](GetPseudoClientsForCompanyResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyCompany**
> GenericResp ModifyCompany(ctx, companyId, optional)
Update the properties of existing company

Modify the properties of an existing company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | **int32**| Id of the company to update | 
 **optional** | ***CompanyApiModifyCompanyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CompanyApiModifyCompanyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateCompany**](UpdateCompany.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCompanyLaptopAdmins**
> UpdateLaptopAdminsResp UpdateCompanyLaptopAdmins(ctx, companyId, optional)


Update a company's laptop admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | **string**|  | 
 **optional** | ***CompanyApiUpdateCompanyLaptopAdminsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CompanyApiUpdateCompanyLaptopAdminsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateLaptopAdminsReq**](UpdateLaptopAdminsReq.md)| List of users or user groups and operation type | 

### Return type

[**UpdateLaptopAdminsResp**](UpdateLaptopAdminsResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCompanyTimezoneForServers**
> GenericResp UpdateCompanyTimezoneForServers(ctx, companyId, optional)
Update the company Timezone for servers

Assigns company's timezone as the timezone for servers with no packages installed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | **int32**| Id of the company whose timezone is used to update timezone of servers (with no packages) | 
 **optional** | ***CompanyApiUpdateCompanyTimezoneForServersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CompanyApiUpdateCompanyTimezoneForServersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateTimezoneForClientsReq**](UpdateTimezoneForClientsReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

