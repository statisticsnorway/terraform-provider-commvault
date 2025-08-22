# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCVFSS3Bucket**](CVFSApi.md#CreateCVFSS3Bucket) | **Post** /V4/CVFS/S3Bucket | Create CVFS S3 Bucket of given properties
[**CreateS3BucketClone**](CVFSApi.md#CreateS3BucketClone) | **Post** /V4/CVFS/S3Bucket/{id}/Clone | Create the bucket clone
[**DeleteCVFSS3Bucket**](CVFSApi.md#DeleteCVFSS3Bucket) | **Delete** /V4/CVFS/S3Bucket/{id} | Delete existing CVFS S3 Bucket
[**DeleteS3BucketClone**](CVFSApi.md#DeleteS3BucketClone) | **Delete** /V4/CVFS/S3Bucket/{id}/Clone/{cloneId} | Delete the S3 Clone
[**GetCVFSS3Bucket**](CVFSApi.md#GetCVFSS3Bucket) | **Get** /V4/CVFS/S3Bucket/{id} | Fetch existing CVFS S3 Bucket details
[**GetListS3BucketClone**](CVFSApi.md#GetListS3BucketClone) | **Get** /V4/CVFS/S3Bucket/{id}/Clone | Fetch the list of clones
[**GetS3AccessKey**](CVFSApi.md#GetS3AccessKey) | **Get** /V4/User/S3AccessKey | Fetch the S3AccessKey for the logged in user
[**GetS3BucketCloneDetails**](CVFSApi.md#GetS3BucketCloneDetails) | **Get** /V4/CVFS/S3Bucket/{id}/Clone/{cloneId} | Fetch the CVFS S3 Clone Details
[**GetUserS3AccessKey**](CVFSApi.md#GetUserS3AccessKey) | **Get** /V4/User/{id}/S3AccessKey | Request the S3 Access Key for the given User Id
[**RotateS3AccessKey**](CVFSApi.md#RotateS3AccessKey) | **Put** /V4/User/S3AccessKey | Rotate the S3AccessKey for the logged in user
[**RotateUserS3AccessKey**](CVFSApi.md#RotateUserS3AccessKey) | **Put** /V4/User/{id}/S3AccessKey | Rotate the S3AccessKey for the given User Id
[**UpdateCVFSS3Bucket**](CVFSApi.md#UpdateCVFSS3Bucket) | **Put** /V4/CVFS/S3Bucket/{id} | Update exsiting CVFS S3 Bucket properties and User Access
[**UpdateS3BucketCloneExpiry**](CVFSApi.md#UpdateS3BucketCloneExpiry) | **Put** /V4/CVFS/S3Bucket/{id}/Clone/{cloneId} | Update the bucket clone properties

# **CreateCVFSS3Bucket**
> CreateCvfss3BucketResp CreateCVFSS3Bucket(ctx, optional)
Create CVFS S3 Bucket of given properties

Used to create CVFS S3 bucket and assigns READ-WRITE permission to the requesting user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CVFSApiCreateCVFSS3BucketOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CVFSApiCreateCVFSS3BucketOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateCvfss3Bucket**](CreateCvfss3Bucket.md)|  | 

### Return type

[**CreateCvfss3BucketResp**](CreateCVFSS3BucketResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateS3BucketClone**
> CreateCvfss3BucketCloneResponse CreateS3BucketClone(ctx, id, optional)
Create the bucket clone

Create the bucket clone of the given bucket id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of the bucket | 
 **optional** | ***CVFSApiCreateS3BucketCloneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CVFSApiCreateS3BucketCloneOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateCvfss3BucketClone**](CreateCvfss3BucketClone.md)|  | 

### Return type

[**CreateCvfss3BucketCloneResponse**](CreateCVFSS3BucketCloneResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCVFSS3Bucket**
> GenericResponse DeleteCVFSS3Bucket(ctx, id)
Delete existing CVFS S3 Bucket

Used to delete an existing CVFS S3 Bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of the bucket | 

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteS3BucketClone**
> GenericResponse DeleteS3BucketClone(ctx, id, cloneId)
Delete the S3 Clone

Delete the clone with the given ID before the expiry time

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of the bucket | 
  **cloneId** | **int32**| ID of the cloned bucket | 

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCVFSS3Bucket**
> GetCvfss3BucketResp GetCVFSS3Bucket(ctx, id)
Fetch existing CVFS S3 Bucket details

Fetches the details of the given CVFS bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of the bucket | 

### Return type

[**GetCvfss3BucketResp**](GetCVFSS3BucketResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListS3BucketClone**
> ListCvfss3BucketCloneResp GetListS3BucketClone(ctx, id)
Fetch the list of clones

Fetches the list of clones for the given bucket id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of the bucket | 

### Return type

[**ListCvfss3BucketCloneResp**](ListCVFSS3BucketCloneResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetS3AccessKey**
> GenericResponse GetS3AccessKey(ctx, )
Fetch the S3AccessKey for the logged in user

Get the S3 Access Key for the logged in user on email

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetS3BucketCloneDetails**
> GetCvfss3BucketCloneDetailsResp GetS3BucketCloneDetails(ctx, id, cloneId)
Fetch the CVFS S3 Clone Details

Fetch the details of the clone of the given ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of the bucket | 
  **cloneId** | **int32**| ID of the cloned bucket | 

### Return type

[**GetCvfss3BucketCloneDetailsResp**](GetCVFSS3BucketCloneDetailsResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserS3AccessKey**
> GenericResponse GetUserS3AccessKey(ctx, id)
Request the S3 Access Key for the given User Id

Get the S3 Access Key for the given user on their email

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of the user | 

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RotateS3AccessKey**
> GenericResponse RotateS3AccessKey(ctx, )
Rotate the S3AccessKey for the logged in user

rotates the S3AccessKey for the logged in user

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RotateUserS3AccessKey**
> GenericResponse RotateUserS3AccessKey(ctx, id)
Rotate the S3AccessKey for the given User Id

Rotate the S3AccessKey for the given User Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of the user | 

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCVFSS3Bucket**
> GenericResponse UpdateCVFSS3Bucket(ctx, id, optional)
Update exsiting CVFS S3 Bucket properties and User Access

Used to update CVFS S3 Bucket properties and User Access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of the bucket | 
 **optional** | ***CVFSApiUpdateCVFSS3BucketOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CVFSApiUpdateCVFSS3BucketOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateCvfss3Bucket**](UpdateCvfss3Bucket.md)|  | 

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateS3BucketCloneExpiry**
> GenericResponse UpdateS3BucketCloneExpiry(ctx, id, cloneId, optional)
Update the bucket clone properties

Update the bucket clone properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of the bucket | 
  **cloneId** | **int32**| ID of the cloned bucket | 
 **optional** | ***CVFSApiUpdateS3BucketCloneExpiryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CVFSApiUpdateS3BucketCloneExpiryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpdateCvfss3CloneReq**](UpdateCvfss3CloneReq.md)|  | 

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

