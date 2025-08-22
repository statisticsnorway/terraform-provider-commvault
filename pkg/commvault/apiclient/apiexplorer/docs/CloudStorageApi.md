# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetEligibleMediaAgentsForAccessPath**](CloudStorageApi.md#CloudGetEligibleMediaAgentsForAccessPath) | **Get** /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}/AccessPath/MediaAgents | 
[**CreateAccessPathForBucketOfCloudStorage**](CloudStorageApi.md#CreateAccessPathForBucketOfCloudStorage) | **Post** /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}/AccessPath | 
[**CreateBucketforCloudStorage**](CloudStorageApi.md#CreateBucketforCloudStorage) | **Post** /V4/Storage/Cloud/{cloudStorageId}/Bucket | Create a new bucket for a specific cloud storage
[**CreateCloudStorage**](CloudStorageApi.md#CreateCloudStorage) | **Post** /V4/Storage/Cloud | Create a Cloud Storage
[**CreateCloudStorageMetaDataCache**](CloudStorageApi.md#CreateCloudStorageMetaDataCache) | **Post** /V4/Storage/Cloud/{cloudStorageId}/MetadataCache | 
[**DeleteAccessPathForBucketOfCloudStorage**](CloudStorageApi.md#DeleteAccessPathForBucketOfCloudStorage) | **Delete** /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}/AccessPath/{accessPathId} | 
[**DeleteBucketOfCloudStorage**](CloudStorageApi.md#DeleteBucketOfCloudStorage) | **Delete** /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId} | 
[**DeleteCloudStorageById**](CloudStorageApi.md#DeleteCloudStorageById) | **Delete** /V4/Storage/Cloud/{cloudStorageId} | 
[**DeleteCloudStorageMetadataCacheById**](CloudStorageApi.md#DeleteCloudStorageMetadataCacheById) | **Delete** /V4/Storage/Cloud/{cloudStorageId}/MetadataCache/{metadataCacheId} | 
[**GetBucketDetailsOfCloudStorage**](CloudStorageApi.md#GetBucketDetailsOfCloudStorage) | **Get** /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId} | 
[**GetCloudStorageById**](CloudStorageApi.md#GetCloudStorageById) | **Get** /V4/Storage/Cloud/{cloudStorageId} | 
[**GetCloudStorageList**](CloudStorageApi.md#GetCloudStorageList) | **Get** /V4/Storage/Cloud | Get All Cloud Storage
[**GetCloudStorageMetaDataCacheById**](CloudStorageApi.md#GetCloudStorageMetaDataCacheById) | **Get** /V4/Storage/Cloud/{cloudStorageId}/MetadataCache/{metadataCacheId} | 
[**ModifyAccessPathOfBucketOfCloudStorage**](CloudStorageApi.md#ModifyAccessPathOfBucketOfCloudStorage) | **Put** /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}/AccessPath/{accessPathId} | 
[**ModifyBucketOfCloudStorage**](CloudStorageApi.md#ModifyBucketOfCloudStorage) | **Put** /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId} | 
[**ModifyCloudStorageById**](CloudStorageApi.md#ModifyCloudStorageById) | **Put** /V4/Storage/Cloud/{cloudStorageId} | 
[**ModifyCloudStorageMetaDataCacheById**](CloudStorageApi.md#ModifyCloudStorageMetaDataCacheById) | **Put** /V4/Storage/Cloud/{cloudStorageId}/MetadataCache/{metadataCacheId} | 

# **CloudGetEligibleMediaAgentsForAccessPath**
> MediaAgentList CloudGetEligibleMediaAgentsForAccessPath(ctx, cloudStorageId, bucketId)


Used to fetch available media agents which can be added as access paths for cloud storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudStorageId** | **int32**| Id of the cloud storage pool for which media agent has to be added | 
  **bucketId** | **int32**| Id of the access path of which media agent has to be shared | 

### Return type

[**MediaAgentList**](MediaAgentList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAccessPathForBucketOfCloudStorage**
> GenericResp CreateAccessPathForBucketOfCloudStorage(ctx, cloudStorageId, bucketId, optional)


Add a new Access path (mediaAgent) to a specific bucket of a specific cloud storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudStorageId** | **int32**| Id of cloud Storage | 
  **bucketId** | **int32**| Id of Bucket | 
 **optional** | ***CloudStorageApiCreateAccessPathForBucketOfCloudStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudStorageApiCreateAccessPathForBucketOfCloudStorageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of BucketIdAccessPathBody**](BucketIdAccessPathBody.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBucketforCloudStorage**
> InlineResponse2003 CreateBucketforCloudStorage(ctx, cloudStorageId, optional)
Create a new bucket for a specific cloud storage

Create a new bucket for a specific cloud storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudStorageId** | **int32**|  | 
 **optional** | ***CloudStorageApiCreateBucketforCloudStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudStorageApiCreateBucketforCloudStorageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CloudStorageIdBucketBody**](CloudStorageIdBucketBody.md)|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCloudStorage**
> CloudStorageResp CreateCloudStorage(ctx, optional)
Create a Cloud Storage

Create a Cloud Storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CloudStorageApiCreateCloudStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudStorageApiCreateCloudStorageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of StorageCloudBody**](StorageCloudBody.md)|  | 

### Return type

[**CloudStorageResp**](CloudStorageResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCloudStorageMetaDataCache**
> GenericResp CreateCloudStorageMetaDataCache(ctx, cloudStorageId, optional)


Add metadata cache paths to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudStorageId** | **int32**| Id of cloud Storage | 
 **optional** | ***CloudStorageApiCreateCloudStorageMetaDataCacheOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudStorageApiCreateCloudStorageMetaDataCacheOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateMetadataCacheConfigurations**](CreateMetadataCacheConfigurations.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccessPathForBucketOfCloudStorage**
> GenericResp DeleteAccessPathForBucketOfCloudStorage(ctx, cloudStorageId, bucketId, accessPathId)


Disassociate a MediaAgent (cloud Access Path) from a bucket of a cloud storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudStorageId** | **int32**| Id of cloud Storage | 
  **bucketId** | **int32**| Id of Bucket | 
  **accessPathId** | **int32**| Id of access path (can be fetched from GET Bucket Details API) | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBucketOfCloudStorage**
> GenericResp DeleteBucketOfCloudStorage(ctx, cloudStorageId, bucketId, optional)


Delete the specified bucket of the cloud storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudStorageId** | **int32**| Id of cloud Storage | 
  **bucketId** | **int32**| Id of Bucket | 
 **optional** | ***CloudStorageApiDeleteBucketOfCloudStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudStorageApiDeleteBucketOfCloudStorageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceDelete** | **optional.Bool**| Force deletes the bucket when physical deletion of data not possible. It will do DB cleanup only. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudStorageById**
> GenericResp DeleteCloudStorageById(ctx, cloudStorageId)


Delete existing cloud storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudStorageId** | **int32**| Id of cloud Storage | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudStorageMetadataCacheById**
> GenericResp DeleteCloudStorageMetadataCacheById(ctx, cloudStorageId, metadataCacheId)


Delete metadata cache of an existing cloud storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudStorageId** | **int32**| Id of cloud Storage | 
  **metadataCacheId** | **int32**| Id of metadata cache | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBucketDetailsOfCloudStorage**
> InlineResponse2004 GetBucketDetailsOfCloudStorage(ctx, cloudStorageId, bucketId)


Get details of a specific bucket of a cloud storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudStorageId** | **int32**| Id of cloud Storage | 
  **bucketId** | **int32**| Id of Bucket | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudStorageById**
> InlineResponse2002 GetCloudStorageById(ctx, cloudStorageId, optional)


Get Cloud Storage Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudStorageId** | **int32**| Id of cloud Storage | 
 **optional** | ***CloudStorageApiGetCloudStorageByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudStorageApiGetCloudStorageByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showInheritedAssociation** | **optional.String**| Set to true if want to show inherited security associations | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudStorageList**
> InlineResponse2001 GetCloudStorageList(ctx, optional)
Get All Cloud Storage

Get All Cloud Storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CloudStorageApiGetCloudStorageListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudStorageApiGetCloudStorageListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storageSubType** | **optional.String**| Filter cloud storage list to given subtype. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudStorageMetaDataCacheById**
> MetadataCacheConfiguration GetCloudStorageMetaDataCacheById(ctx, cloudStorageId, metadataCacheId)


Get details of metadata Cache of a cloud storage based on Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudStorageId** | **int32**| Id of cloud Storage | 
  **metadataCacheId** | **int32**| Id of metadata cache | 

### Return type

[**MetadataCacheConfiguration**](MetadataCacheConfiguration.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyAccessPathOfBucketOfCloudStorage**
> GenericResp ModifyAccessPathOfBucketOfCloudStorage(ctx, cloudStorageId, bucketId, accessPathId, optional)


Modify access path details of specific bucket of a specific cloud storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudStorageId** | **int32**| Id of cloud Storage | 
  **bucketId** | **int32**| Id of Bucket | 
  **accessPathId** | **int32**| Id of access path (can be fetched from GET Bucket Details API) | 
 **optional** | ***CloudStorageApiModifyAccessPathOfBucketOfCloudStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudStorageApiModifyAccessPathOfBucketOfCloudStorageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of CloudStorageAdvanced**](CloudStorageAdvanced.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyBucketOfCloudStorage**
> GenericResp ModifyBucketOfCloudStorage(ctx, cloudStorageId, bucketId, optional)


Modify configuration of a specific bucket of a specific cloud storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudStorageId** | **int32**| Id of cloud Storage | 
  **bucketId** | **int32**| Id of Bucket | 
 **optional** | ***CloudStorageApiModifyBucketOfCloudStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudStorageApiModifyBucketOfCloudStorageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of BucketBucketIdBody**](BucketBucketIdBody.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyCloudStorageById**
> GenericResp ModifyCloudStorageById(ctx, cloudStorageId, optional)


Modify details like name, encryption, security of a specific cloud storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudStorageId** | **int32**| Id of cloud Storage | 
 **optional** | ***CloudStorageApiModifyCloudStorageByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudStorageApiModifyCloudStorageByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateCloudStorage**](UpdateCloudStorage.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyCloudStorageMetaDataCacheById**
> GenericResp ModifyCloudStorageMetaDataCacheById(ctx, cloudStorageId, metadataCacheId, optional)


Modify details like credentials and path for metadata Cache of a cloud storage based on Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudStorageId** | **int32**| Id of cloud Storage | 
  **metadataCacheId** | **int32**| Id of metadata cache | 
 **optional** | ***CloudStorageApiModifyCloudStorageMetaDataCacheByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudStorageApiModifyCloudStorageMetaDataCacheByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpdateMetadataCacheConfiguration**](UpdateMetadataCacheConfiguration.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

