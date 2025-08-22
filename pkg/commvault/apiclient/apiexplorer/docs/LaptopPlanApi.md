# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLaptopBackupDestination**](LaptopPlanApi.md#CreateLaptopBackupDestination) | **Post** /V4/LaptopPlan/{planId}/BackupDestination | Create a LaptopPlan BackupDestination
[**CreateLaptopPlan**](LaptopPlanApi.md#CreateLaptopPlan) | **Post** /V4/LaptopPlan | Create Laptop Plan
[**DeleteLaptopBackupDestination**](LaptopPlanApi.md#DeleteLaptopBackupDestination) | **Delete** /V4/LaptopPlan/{planId}/BackupDestination/{BackupDestinationId} | Delete LaptopPlan BackupDestination
[**DeleteLaptopPlanById**](LaptopPlanApi.md#DeleteLaptopPlanById) | **Delete** /V4/LaptopPlan/{planId} | Delete existing laptop plan
[**DeleteStorageRegionForLaptopPlan**](LaptopPlanApi.md#DeleteStorageRegionForLaptopPlan) | **Delete** /V4/LaptopPlan/{planId}/storageRegion/{regionList} | Delete Storage Region of LaptopPlan
[**GetLaptopBackupDestinations**](LaptopPlanApi.md#GetLaptopBackupDestinations) | **Get** /V4/LaptopPlan/{planId}/BackupDestinations | Get LaptopPlan BackupDestinations
[**GetLaptopPlanById**](LaptopPlanApi.md#GetLaptopPlanById) | **Get** /V4/LaptopPlan/{planId} | Get Laptop Plan details
[**ModifyLaptopBackupDestination**](LaptopPlanApi.md#ModifyLaptopBackupDestination) | **Put** /V4/LaptopPlan/{planId}/BackupDestination/{BackupDestinationId} | Update LaptopPlan BackupDestination
[**ModifyLaptopPlanById**](LaptopPlanApi.md#ModifyLaptopPlanById) | **Put** /V4/LaptopPlan/{planId} |  Modify existing laptop plan details

# **CreateLaptopBackupDestination**
> PlanBackupDestinationResp CreateLaptopBackupDestination(ctx, planId, optional)
Create a LaptopPlan BackupDestination

Create a Backup Destination for a Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan to modify | 
 **optional** | ***LaptopPlanApiCreateLaptopBackupDestinationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaptopPlanApiCreateLaptopBackupDestinationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreatePlanBackupDestinations**](CreatePlanBackupDestinations.md)|  | 

### Return type

[**PlanBackupDestinationResp**](PlanBackupDestinationResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLaptopPlan**
> InlineResponse200 CreateLaptopPlan(ctx, optional)
Create Laptop Plan

Create Laptop Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LaptopPlanApiCreateLaptopPlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaptopPlanApiCreateLaptopPlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of V4LaptopPlanBody**](V4LaptopPlanBody.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLaptopBackupDestination**
> GenericResp DeleteLaptopBackupDestination(ctx, planId, backupDestinationId)
Delete LaptopPlan BackupDestination

Delete Backup Destination for a Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan to modify | 
  **backupDestinationId** | **int32**| Id of the backupDestination to be deleted | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLaptopPlanById**
> GenericResp DeleteLaptopPlanById(ctx, planId)
Delete existing laptop plan

Delete existing laptop plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the plan | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStorageRegionForLaptopPlan**
> GenericResp DeleteStorageRegionForLaptopPlan(ctx, planId, regionList, optional)
Delete Storage Region of LaptopPlan

Delete Storage Region for a laptop Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Laptop Plan to modify | 
  **regionList** | **string**| List of region names/ids to be deleted. If region ids are passed, set isRegionIdList&#x3D;true | 
 **optional** | ***LaptopPlanApiDeleteStorageRegionForLaptopPlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaptopPlanApiDeleteStorageRegionForLaptopPlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isRegionIdList** | **optional.Bool**| Is regionList a list of region ids | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLaptopBackupDestinations**
> PlanBackupDestinations GetLaptopBackupDestinations(ctx, planId)
Get LaptopPlan BackupDestinations

Get Backup Destinations for a Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Laptop Plan to retrieve backup destinations | 

### Return type

[**PlanBackupDestinations**](PlanBackupDestinations.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLaptopPlanById**
> LaptopPlanDetails GetLaptopPlanById(ctx, planId)
Get Laptop Plan details

Get Laptop Plan details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the plan | 

### Return type

[**LaptopPlanDetails**](LaptopPlanDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyLaptopBackupDestination**
> PlanBackupDestinationResp ModifyLaptopBackupDestination(ctx, planId, backupDestinationId, optional)
Update LaptopPlan BackupDestination

Modify Backup Destination for a Plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the Plan to modify | 
  **backupDestinationId** | **int32**| Id of the backupDestination to be modified | 
 **optional** | ***LaptopPlanApiModifyLaptopBackupDestinationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaptopPlanApiModifyLaptopBackupDestinationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpdatePlanBackupDestination**](UpdatePlanBackupDestination.md)|  | 

### Return type

[**PlanBackupDestinationResp**](PlanBackupDestinationResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyLaptopPlanById**
> GenericResp ModifyLaptopPlanById(ctx, planId, optional)
 Modify existing laptop plan details

Modify existing laptop plan details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planId** | **int32**| Id of the plan | 
 **optional** | ***LaptopPlanApiModifyLaptopPlanByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaptopPlanApiModifyLaptopPlanByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateLaptopPlan**](UpdateLaptopPlan.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

