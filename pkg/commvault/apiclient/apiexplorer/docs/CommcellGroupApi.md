# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommcellGroups**](CommcellGroupApi.md#CreateCommcellGroups) | **Post** /V4/CommcellGroup | Create a Commcell Group
[**DeleteCommcellGroup**](CommcellGroupApi.md#DeleteCommcellGroup) | **Delete** /V4/CommcellGroup/{commcellGroupId} | Delete a CommcellGroup
[**GetCommcellGroupIdDetails**](CommcellGroupApi.md#GetCommcellGroupIdDetails) | **Get** /V4/CommcellGroup/{commcellGroupId} | Get Commcell Group Details
[**GetCommcellGroups**](CommcellGroupApi.md#GetCommcellGroups) | **Get** /V4/CommcellGroup | Get Commcell Groups
[**GetCommcellGroups_0**](CommcellGroupApi.md#GetCommcellGroups_0) | **Get** /CommcellGroup | Get Commcell Groups
[**UpdateCommcellGroupAssociation**](CommcellGroupApi.md#UpdateCommcellGroupAssociation) | **Put** /V4/CommcellGroup/{commcellGroupId} | Update the commcell group

# **CreateCommcellGroups**
> CreateCommcellGroupResp CreateCommcellGroups(ctx, optional)
Create a Commcell Group

Create Commcell group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommcellGroupApiCreateCommcellGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommcellGroupApiCreateCommcellGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateCommcellGroupReq**](CreateCommcellGroupReq.md)|  | 

### Return type

[**CreateCommcellGroupResp**](CreateCommcellGroupResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCommcellGroup**
> GenericResp DeleteCommcellGroup(ctx, commcellGroupId)
Delete a CommcellGroup

Used to delete a commcellGroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commcellGroupId** | **int32**| Id of the commcellGroup to delete | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommcellGroupIdDetails**
> CommcellGroupInfo GetCommcellGroupIdDetails(ctx, commcellGroupId, optional)
Get Commcell Group Details

Get details of a commcellGroup based on id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commcellGroupId** | **int32**| Id of the commcellGroupId whose details have to be fetched | 
 **optional** | ***CommcellGroupApiGetCommcellGroupIdDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommcellGroupApiGetCommcellGroupIdDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | [**optional.Interface of Mode**](.md)|  | 

### Return type

[**CommcellGroupInfo**](CommcellGroupInfo.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommcellGroups**
> GetCommcellGroupsResp GetCommcellGroups(ctx, )
Get Commcell Groups

Get All Commcell Groups

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetCommcellGroupsResp**](GetCommcellGroupsResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommcellGroups_0**
> GetCommcellGroupsResp GetCommcellGroups_0(ctx, )
Get Commcell Groups

Get Commcell Groups info

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetCommcellGroupsResp**](GetCommcellGroupsResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCommcellGroupAssociation**
> GenericResp UpdateCommcellGroupAssociation(ctx, commcellGroupId, optional)
Update the commcell group

Used to update commcell associations for a server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commcellGroupId** | **int32**| Id of the commcellGroupId whose details have to be fetched | 
 **optional** | ***CommcellGroupApiUpdateCommcellGroupAssociationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommcellGroupApiUpdateCommcellGroupAssociationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateCommcellGroupReq**](UpdateCommcellGroupReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

