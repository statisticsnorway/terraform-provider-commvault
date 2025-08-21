# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatelockConfiguration**](TapeStorageApi.md#CreatelockConfiguration) | **Put** /V4/MMConfiguration/Lock | 
[**CreatelunlockConfiguration**](TapeStorageApi.md#CreatelunlockConfiguration) | **Put** /V4/MMConfiguration/Unlock | 
[**FetchTapeDetails**](TapeStorageApi.md#FetchTapeDetails) | **Get** /V4/Storage/Tape/{libraryId} | 
[**FetchTapeMediaDetails**](TapeStorageApi.md#FetchTapeMediaDetails) | **Get** /V4/Storage/Tape/{libraryId}/Media | 
[**GetInfoForDiscoverMedia**](TapeStorageApi.md#GetInfoForDiscoverMedia) | **Get** /V4/Storage/Tape/{libraryId}/MediaType | Get Info for Discover Media Operation
[**GetMediaListForLibrary**](TapeStorageApi.md#GetMediaListForLibrary) | **Get** /V4/Library/{libraryId}/Media | Get Media List for a Library
[**GetSlotListForLibrary**](TapeStorageApi.md#GetSlotListForLibrary) | **Get** /V4/Library/{libraryId}/Slot | Get Slot List for a Library
[**GetSpareGroupDetails**](TapeStorageApi.md#GetSpareGroupDetails) | **Get** /V4/Storage/Tape/{libraryId}/SpareGroups | List of Spare Group Details
[**GetSpareGroupProperties**](TapeStorageApi.md#GetSpareGroupProperties) | **Get** /V4/Storage/Tape/{libraryId}/SpareGroup/{spareGroupId} | Fetches properties of a Spare group
[**GetVaultTackerActionList**](TapeStorageApi.md#GetVaultTackerActionList) | **Get** /V4/VaultTrackerAction | Get Vault Tracker Action List
[**GetVaultTrackerMediaDetails**](TapeStorageApi.md#GetVaultTrackerMediaDetails) | **Get** /V4/VaultTracker/Media | Get Vault Tracker Media Details either by actionId or by policyId or by containerId or by locationId
[**PerformTapeMediaOperation**](TapeStorageApi.md#PerformTapeMediaOperation) | **Post** /V4/Storage/Tape/{libraryId}/Media | Perform operation on Tape media
[**ReturnDrivePoolDetails**](TapeStorageApi.md#ReturnDrivePoolDetails) | **Get** /V4/Storage/Tape/{libraryId}/Drive/{driveId} | 
[**ReturnListOfLocations**](TapeStorageApi.md#ReturnListOfLocations) | **Get** /V4/Storage/Tape/Locations | 
[**ReturnListOfTapes**](TapeStorageApi.md#ReturnListOfTapes) | **Get** /V4/Storage/Tape | Get Tape Storage
[**VaultTrackerOperation**](TapeStorageApi.md#VaultTrackerOperation) | **Post** /V4/VaultTrackerAction/Operation | Operations for VaultTracker Actions

# **CreatelockConfiguration**
> GenericRespWithWarning CreatelockConfiguration(ctx, )


To lock a configuration

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GenericRespWithWarning**](GenericRespWithWarning.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatelunlockConfiguration**
> GenericRespWithWarning CreatelunlockConfiguration(ctx, )


To unlock a configuration

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GenericRespWithWarning**](GenericRespWithWarning.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchTapeDetails**
> TapeDetails FetchTapeDetails(ctx, libraryId)


Fetches tape details based on the tape library Id provided

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **libraryId** | **int32**| Id of the library to view the data | 

### Return type

[**TapeDetails**](TapeDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchTapeMediaDetails**
> TapeMediaDetails FetchTapeMediaDetails(ctx, libraryId)


Fetch tape media details based on the tape library Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **libraryId** | **int32**| Id of the library to view the data | 

### Return type

[**TapeMediaDetails**](TapeMediaDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInfoForDiscoverMedia**
> DiscoverMediaInfoResp GetInfoForDiscoverMedia(ctx, libraryId)
Get Info for Discover Media Operation

API to fetch all required details for performing Discover Media operation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **libraryId** | **int32**| Library ID of the Tape Storage | 

### Return type

[**DiscoverMediaInfoResp**](DiscoverMediaInfoResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMediaListForLibrary**
> GetMediaListResp GetMediaListForLibrary(ctx, libraryId, filterMediaType, isExported)
Get Media List for a Library

Get the list of media for given library.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **libraryId** | **int32**| Id of Library | 
  **filterMediaType** | [**FilterMediaType**](.md)| Filter media list to given FilterMediaType | 
  **isExported** | **bool**| List \&quot;Media In Library\&quot; only if isExported is set to false otherwise list \&quot;Exported Media\&quot; | 

### Return type

[**GetMediaListResp**](GetMediaListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSlotListForLibrary**
> GetSlotListResp GetSlotListForLibrary(ctx, libraryId, slotType)
Get Slot List for a Library

Get the list of slot for given library.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **libraryId** | **int32**| Id of Library | 
  **slotType** | [**SlotTypes**](.md)| Filter slot list to given SlotTypes | 

### Return type

[**GetSlotListResp**](GetSlotListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpareGroupDetails**
> GetSpareGroupListResp GetSpareGroupDetails(ctx, libraryId, spareGroupType)
List of Spare Group Details

API to fetch information of Spare Groups of a particular type for a Storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **libraryId** | **int32**| Library ID of the Tape Storage | 
  **spareGroupType** | [**SpareGroupTypes**](.md)| Get all spare groups of given SpareGroupType | 

### Return type

[**GetSpareGroupListResp**](GetSpareGroupListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpareGroupProperties**
> GetSpareGroupProperties GetSpareGroupProperties(ctx, spareGroupId, libraryId)
Fetches properties of a Spare group

API to fetch properties of a Spare Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spareGroupId** | **int32**| Id of the Spare group whose properties need to be fetched. spareGroupId can be fetched from GET V4/Storage/Tape/{libraryId}/MediaType | 
  **libraryId** | **int32**| Library ID of the Tape Storage | 

### Return type

[**GetSpareGroupProperties**](GetSpareGroupProperties.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVaultTackerActionList**
> GetVaultTrackerActionListResp GetVaultTackerActionList(ctx, optional)
Get Vault Tracker Action List

Get the list of vault tracker actions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TapeStorageApiGetVaultTackerActionListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TapeStorageApiGetVaultTackerActionListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **optional.String**| Time period selection for which to fetch actions. Accepted values [lastHour, last24Hours, lastWeek, lastMonth, last3Months, last6Months, lastYear, custom]. When custom is selected Actions are filtered based on values provided in other params. | [default to lastHour]
 **libraryId** | **optional.Int32**| Filter by libraryId | [default to 0]
 **policyId** | **optional.Int32**| Filter by policyId | [default to 0]
 **startTime** | **optional.Int64**| Start time of the time range. | 
 **endTime** | **optional.Int64**| End time of the time range. | 

### Return type

[**GetVaultTrackerActionListResp**](GetVaultTrackerActionListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVaultTrackerMediaDetails**
> GetVaultTrackerMediaDetailsResp GetVaultTrackerMediaDetails(ctx, optional)
Get Vault Tracker Media Details either by actionId or by policyId or by containerId or by locationId

Get the details of vault tracker media.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TapeStorageApiGetVaultTrackerMediaDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TapeStorageApiGetVaultTrackerMediaDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actionId** | **optional.Int32**| List Media by actionId | [default to 0]
 **policyId** | **optional.Int32**| List Media by policyId | [default to 0]
 **containerId** | **optional.Int32**| List Media by containerId | [default to 0]
 **locationId** | **optional.Int32**| List Media by locationId | [default to 0]

### Return type

[**GetVaultTrackerMediaDetailsResp**](GetVaultTrackerMediaDetailsResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PerformTapeMediaOperation**
> GenericResponse PerformTapeMediaOperation(ctx, optional)
Perform operation on Tape media

API to perform various media level operations on a Tape Storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TapeStorageApiPerformTapeMediaOperationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TapeStorageApiPerformTapeMediaOperationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TapeMediaOperationsReq**](TapeMediaOperationsReq.md)|  | 

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReturnDrivePoolDetails**
> DrivePoolDetailsResp ReturnDrivePoolDetails(ctx, libraryId, driveId)


This endpoint is used to return the details of Drive Pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **libraryId** | **int32**| Id of the Tape of which the drive details has to be displayed | 
  **driveId** | **int32**| Id of the Drive of which the drive details has to be displayed | 

### Return type

[**DrivePoolDetailsResp**](DrivePoolDetailsResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReturnListOfLocations**
> LocationListResp ReturnListOfLocations(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LocationListResp**](LocationListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReturnListOfTapes**
> TapeListResp ReturnListOfTapes(ctx, )
Get Tape Storage

This endpoint is used to return the list of tape storages.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TapeListResp**](TapeListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultTrackerOperation**
> GenericResponse VaultTrackerOperation(ctx, optional)
Operations for VaultTracker Actions

Perform VaultTracker Operations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TapeStorageApiVaultTrackerOperationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TapeStorageApiVaultTrackerOperationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of VaultTrackerOperationReq**](VaultTrackerOperationReq.md)|  | 

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

