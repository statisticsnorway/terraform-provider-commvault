# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableAllJobActivityCommcell**](CommcellApi.md#DisableAllJobActivityCommcell) | **Put** /V4/Commcell/AllJobActivity/Action/Disable | Disable All Job Activity at Commcell Level
[**DisableAuxillaryCopyCommcell**](CommcellApi.md#DisableAuxillaryCopyCommcell) | **Put** /V4/Commcell/AuxillaryCopy/Action/Disable | Disable Auxillary copy Jobs at Commcell Level
[**DisableBackupCommcell**](CommcellApi.md#DisableBackupCommcell) | **Put** /V4/Commcell/Backup/Action/Disable | Disable Backup at Commcell Level
[**DisableContentIndexingCommcell**](CommcellApi.md#DisableContentIndexingCommcell) | **Put** /V4/Commcell/ContentIndexing/Action/Disable | Disable Context Indexing Jobs at Commcell Level
[**DisableDDBCommcell**](CommcellApi.md#DisableDDBCommcell) | **Put** /V4/Commcell/DDB/Action/Disable | Disable DDB Backups at Commcell Level
[**DisableDataAgingCommcell**](CommcellApi.md#DisableDataAgingCommcell) | **Put** /V4/Commcell/DataAging/Action/Disable | Disable Data Aging at Commcell Level
[**DisableDataVerificationCommcell**](CommcellApi.md#DisableDataVerificationCommcell) | **Put** /V4/Commcell/DataVerification/Action/Disable | Disable Data Verification at Commcell Level
[**DisableRestoreCommcell**](CommcellApi.md#DisableRestoreCommcell) | **Put** /V4/Commcell/Restore/Action/Disable | Disable Restore at Commcell Level
[**DisableSchedulerCommcell**](CommcellApi.md#DisableSchedulerCommcell) | **Put** /V4/Commcell/Scheduler/Action/Disable | Disable Scheduler at Commcell Level
[**DownloadOrCopySoftware**](CommcellApi.md#DownloadOrCopySoftware) | **Put** /V4/DownloadSoftware | Download or Copy Software
[**EnableAllJobActivity**](CommcellApi.md#EnableAllJobActivity) | **Put** /V4/Commcell/AllJobActivity/Action/Enable | Enable All Job Activity at Commcell Level
[**EnableAuxillaryCopyCommcell**](CommcellApi.md#EnableAuxillaryCopyCommcell) | **Put** /V4/Commcell/AuxillaryCopy/Action/Enable | Enable Auxillary Copy Jobs at Commcell Level
[**EnableBackupCommcell**](CommcellApi.md#EnableBackupCommcell) | **Put** /V4/Commcell/Backup/Action/Enable | Enable Backup at Commcell Level
[**EnableContentIndexingCommcell**](CommcellApi.md#EnableContentIndexingCommcell) | **Put** /V4/Commcell/ContentIndexing/Action/Enable | Enable Content Indexing Jobs at Commcell Level
[**EnableDDBCommcell**](CommcellApi.md#EnableDDBCommcell) | **Put** /V4/Commcell/DDB/Action/Enable | Enable DDB Backups at Commcell Level
[**EnableDataAgingCommcell**](CommcellApi.md#EnableDataAgingCommcell) | **Put** /V4/Commcell/DataAging/Action/Enable | Enable Data Aging at commcell level
[**EnableDataVerificationCommcell**](CommcellApi.md#EnableDataVerificationCommcell) | **Put** /V4/Commcell/DataVerification/Action/Enable | Enable Data Verification at Commcell Level
[**EnableRestoreCommcell**](CommcellApi.md#EnableRestoreCommcell) | **Put** /V4/Commcell/Restore/Action/Enable | Enable Restore at Commcell Level
[**EnableSchedulerCommcell**](CommcellApi.md#EnableSchedulerCommcell) | **Put** /V4/Commcell/Scheduler/Action/Enable | Enable Scheduler at Commcell Level
[**RepairSoftware**](CommcellApi.md#RepairSoftware) | **Put** /V4/RepairSoftware | Repair software
[**UpgradeSoftware**](CommcellApi.md#UpgradeSoftware) | **Put** /V4/UpgradeSoftware | Upgrade client and client group software

# **DisableAllJobActivityCommcell**
> GenericResp DisableAllJobActivityCommcell(ctx, optional)
Disable All Job Activity at Commcell Level

Used to disable all job activity property for commcell

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommcellApiDisableAllJobActivityCommcellOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommcellApiDisableAllJobActivityCommcellOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableAfterADelay** | **optional.Int32**| Provide UTC time in unix format. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableAuxillaryCopyCommcell**
> GenericResp DisableAuxillaryCopyCommcell(ctx, optional)
Disable Auxillary copy Jobs at Commcell Level

Used to disable auxillary copy property for commcell

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommcellApiDisableAuxillaryCopyCommcellOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommcellApiDisableAuxillaryCopyCommcellOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableAfterADelay** | **optional.Int32**| Provide UTC time in unix format. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableBackupCommcell**
> GenericResp DisableBackupCommcell(ctx, optional)
Disable Backup at Commcell Level

Used to disable backup property for commcell

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommcellApiDisableBackupCommcellOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommcellApiDisableBackupCommcellOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableAfterADelay** | **optional.Int32**| Provide UTC time in unix format. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableContentIndexingCommcell**
> GenericResp DisableContentIndexingCommcell(ctx, optional)
Disable Context Indexing Jobs at Commcell Level

Used to disable content indexing property for commcell

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommcellApiDisableContentIndexingCommcellOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommcellApiDisableContentIndexingCommcellOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableAfterADelay** | **optional.Int32**| Provide UTC time in unix format. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableDDBCommcell**
> GenericResp DisableDDBCommcell(ctx, optional)
Disable DDB Backups at Commcell Level

Used to disable DDB property for commcell

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommcellApiDisableDDBCommcellOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommcellApiDisableDDBCommcellOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableAfterADelay** | **optional.Int32**| Provide UTC time in unix format. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableDataAgingCommcell**
> GenericResp DisableDataAgingCommcell(ctx, optional)
Disable Data Aging at Commcell Level

Used to disable data aging property for commcell

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommcellApiDisableDataAgingCommcellOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommcellApiDisableDataAgingCommcellOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableAfterADelay** | **optional.Int32**| Provide UTC time in unix format. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableDataVerificationCommcell**
> GenericResp DisableDataVerificationCommcell(ctx, optional)
Disable Data Verification at Commcell Level

Used to disable data verification property for commcell

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommcellApiDisableDataVerificationCommcellOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommcellApiDisableDataVerificationCommcellOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableAfterADelay** | **optional.Int32**| Provide UTC time in unix format. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableRestoreCommcell**
> GenericResp DisableRestoreCommcell(ctx, optional)
Disable Restore at Commcell Level

Used to disable restore property for commcell

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommcellApiDisableRestoreCommcellOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommcellApiDisableRestoreCommcellOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableAfterADelay** | **optional.Int32**| Provide UTC time in unix format. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableSchedulerCommcell**
> GenericResp DisableSchedulerCommcell(ctx, optional)
Disable Scheduler at Commcell Level

Used to disable scheduler property for commcell

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommcellApiDisableSchedulerCommcellOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommcellApiDisableSchedulerCommcellOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableAfterADelay** | **optional.Int32**| Provide UTC time in unix format. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadOrCopySoftware**
> JobIdResponse DownloadOrCopySoftware(ctx, optional)
Download or Copy Software

Download or Copy Software

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommcellApiDownloadOrCopySoftwareOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommcellApiDownloadOrCopySoftwareOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DownloadOrCopySoftware**](DownloadOrCopySoftware.md)|  | 

### Return type

[**JobIdResponse**](JobIdResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableAllJobActivity**
> GenericResp EnableAllJobActivity(ctx, )
Enable All Job Activity at Commcell Level

Used to enable all job activity property for commcell

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableAuxillaryCopyCommcell**
> GenericResp EnableAuxillaryCopyCommcell(ctx, )
Enable Auxillary Copy Jobs at Commcell Level

Used to enable auxillary copy property for commcell

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableBackupCommcell**
> GenericResp EnableBackupCommcell(ctx, )
Enable Backup at Commcell Level

Used to enable backup property for commcell

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableContentIndexingCommcell**
> GenericResp EnableContentIndexingCommcell(ctx, )
Enable Content Indexing Jobs at Commcell Level

Used to enable content indexing property for commcell

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableDDBCommcell**
> GenericResp EnableDDBCommcell(ctx, )
Enable DDB Backups at Commcell Level

Used to enable DDB property for commcell

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableDataAgingCommcell**
> GenericResp EnableDataAgingCommcell(ctx, )
Enable Data Aging at commcell level

Used to enable data aging property for commcell

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableDataVerificationCommcell**
> GenericResp EnableDataVerificationCommcell(ctx, )
Enable Data Verification at Commcell Level

Used to enable data verification property for commcell

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableRestoreCommcell**
> GenericResp EnableRestoreCommcell(ctx, )
Enable Restore at Commcell Level

Used to enable restore property for commcell

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableSchedulerCommcell**
> GenericResp EnableSchedulerCommcell(ctx, )
Enable Scheduler at Commcell Level

Used to enable scheduler property for commcell

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepairSoftware**
> JobIdResponse RepairSoftware(ctx, optional)
Repair software

Repair software for client and clients for the given client group(s).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommcellApiRepairSoftwareOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommcellApiRepairSoftwareOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RepairSoftware**](RepairSoftware.md)|  | 

### Return type

[**JobIdResponse**](JobIdResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeSoftware**
> JobIdResponse UpgradeSoftware(ctx, optional)
Upgrade client and client group software

Upgrade software for client and client group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommcellApiUpgradeSoftwareOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommcellApiUpgradeSoftwareOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UpgradeSoftware**](UpgradeSoftware.md)|  | 

### Return type

[**JobIdResponse**](JobIdResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

