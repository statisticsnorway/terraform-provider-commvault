# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrowseKubernetesNamespaces**](KubernetesApi.md#BrowseKubernetesNamespaces) | **Get** /V5/Kubernetes/Cluster/{clusterId}/Content/Namespace | Browse all namespaces of a Kubernetes cluster
[**BrowseKubernetesStorageClass**](KubernetesApi.md#BrowseKubernetesStorageClass) | **Get** /V5/Kubernetes/Cluster/{clusterId}/Content/StorageClass | Browse all storage classes of a Kubernetes cluster
[**BrowseNamespaceApplications**](KubernetesApi.md#BrowseNamespaceApplications) | **Get** /V5/Kubernetes/Cluster/{clusterId}/Content/Namespace/{nameSpace}/Applications | Content browse applications in a namespace
[**BrowseNamespaceLabels**](KubernetesApi.md#BrowseNamespaceLabels) | **Get** /V5/Kubernetes/Cluster/{clusterId}/Content/Namespace/{nameSpace}/Labels | Content browse labels in a namespace
[**BrowseNamespaceVolumes**](KubernetesApi.md#BrowseNamespaceVolumes) | **Get** /V5/Kubernetes/Cluster/{clusterId}/Content/Namespace/{nameSpace}/Volumes | Content browse volumes in a namespace
[**CreateKubernetesApplicationGroup**](KubernetesApi.md#CreateKubernetesApplicationGroup) | **Post** /V5/Kubernetes/ApplicationGroup | Create a new Kubernetes application group
[**CreateKubernetesClusterOp**](KubernetesApi.md#CreateKubernetesClusterOp) | **Post** /V5/Kubernetes/Cluster | Create a new Kubernetes cluster client
[**DeleteDeconfiguredKubernetesCluster**](KubernetesApi.md#DeleteDeconfiguredKubernetesCluster) | **Delete** /V5/Kubernetes/Cluster/{clusterId} | Delete a deconfigured Kubernetes cluster client
[**DeleteKubernetesAppGroup**](KubernetesApi.md#DeleteKubernetesAppGroup) | **Delete** /V5/Kubernetes/ApplicationGroup/{applicationGroupId} | Delete a Kubernetes application group
[**DoApplicationBrowse**](KubernetesApi.md#DoApplicationBrowse) | **Get** /V5/Kubernetes/ApplicationGroup/{applicationGroupId}/Browse/Application | Browse backed up Applications of the application group
[**DoApplicationRestore**](KubernetesApi.md#DoApplicationRestore) | **Post** /V5/Kubernetes/ApplicationGroup/{applicationGroupId}/Restore/Application | Full Application Restore - Restore backed up Applications of the application group
[**DoNamespaceBrowse**](KubernetesApi.md#DoNamespaceBrowse) | **Get** /V5/Kubernetes/ApplicationGroup/{applicationGroupId}/Browse/Namespace | Browse backed up Namespaces of the application group
[**DoNamespaceRestore**](KubernetesApi.md#DoNamespaceRestore) | **Post** /V5/Kubernetes/ApplicationGroup/{applicationGroupId}/Restore/Namespace | Namespace Level Restore - Restore backed up Namespaces of the application group
[**GetAllKubernetesAppGroups**](KubernetesApi.md#GetAllKubernetesAppGroups) | **Get** /V5/Kubernetes/ApplicationGroup | Get all Kubernetes application groups
[**GetAllKubernetesApps**](KubernetesApi.md#GetAllKubernetesApps) | **Get** /V5/Kubernetes/Application | Get all Kubernetes applications
[**GetAllKubernetesClustersOp**](KubernetesApi.md#GetAllKubernetesClustersOp) | **Get** /V5/Kubernetes/Cluster | Get all Kubernetes clusters
[**GetApplicationGroupDetails**](KubernetesApi.md#GetApplicationGroupDetails) | **Get** /V5/Kubernetes/ApplicationGroup/{applicationGroupId} | Get details of a Kubernetes application group
[**GetKubernetesAppDetails**](KubernetesApi.md#GetKubernetesAppDetails) | **Get** /V5/Kubernetes/Application/{appGUID} | Get Kubernetes applications details
[**GetKubernetesClusterDetails**](KubernetesApi.md#GetKubernetesClusterDetails) | **Get** /V5/Kubernetes/Cluster/{clusterId} | Get details of a Kubernetes Cluster client
[**PreviewApplicationGroup**](KubernetesApi.md#PreviewApplicationGroup) | **Post** /V5/Kubernetes/ApplicationGroup/Preview | Preview application group content for discovery of applications
[**ReconfigureKubernetesCluster**](KubernetesApi.md#ReconfigureKubernetesCluster) | **Put** /V5/Kubernetes/Cluster/{clusterId}/Reconfigure | Reconfigure a retired Kubernetes cluster client
[**RetireKubernetesCluster**](KubernetesApi.md#RetireKubernetesCluster) | **Delete** /V5/Kubernetes/Cluster/{clusterId}/Retire | Retire and release license of a Kubernetes clueter client
[**RunApplicationBackup**](KubernetesApi.md#RunApplicationBackup) | **Post** /V5/Kubernetes/Application/{appGUID}/Backup | Run backup for an application
[**RunApplicationGroupBackup**](KubernetesApi.md#RunApplicationGroupBackup) | **Post** /V5/Kubernetes/ApplicationGroup/{applicationGroupId}/Backup | Run backup for an application group
[**UpdateKubernetesAppGroupOp**](KubernetesApi.md#UpdateKubernetesAppGroupOp) | **Put** /V5/Kubernetes/ApplicationGroup/{applicationGroupId} | Update properties of a Kubernetes application group
[**UpdateKubernetesProperties**](KubernetesApi.md#UpdateKubernetesProperties) | **Put** /V5/Kubernetes/Cluster/{clusterId} | Update properties of a Kubernetes cluster client
[**V4KubernetesClusterIdModifierGet**](KubernetesApi.md#V4KubernetesClusterIdModifierGet) | **Get** /V4/Kubernetes/{clusterId}/Modifier | Fetch all the modifiers
[**V4KubernetesClusterIdModifierModifierNameDelete**](KubernetesApi.md#V4KubernetesClusterIdModifierModifierNameDelete) | **Delete** /V4/Kubernetes/{clusterId}/Modifier/{ModifierName} | Delete a Modifier
[**V4KubernetesClusterIdModifierModifierNameGet**](KubernetesApi.md#V4KubernetesClusterIdModifierModifierNameGet) | **Get** /V4/Kubernetes/{clusterId}/Modifier/{ModifierName} | Fetch(Read) Modfier details
[**V4KubernetesClusterIdModifierModifierNamePut**](KubernetesApi.md#V4KubernetesClusterIdModifierModifierNamePut) | **Put** /V4/Kubernetes/{clusterId}/Modifier/{ModifierName} | Edit a Restore modifier
[**V4KubernetesClusterIdModifierPost**](KubernetesApi.md#V4KubernetesClusterIdModifierPost) | **Post** /V4/Kubernetes/{clusterId}/Modifier | Create a Modifier
[**V4KubernetesClusterIdModifierTestPost**](KubernetesApi.md#V4KubernetesClusterIdModifierTestPost) | **Post** /V4/Kubernetes/{clusterId}/Modifier/Test | Test a Yaml paylod against a restore modfier

# **BrowseKubernetesNamespaces**
> KubernetesBrowseClusterResp BrowseKubernetesNamespaces(ctx, clusterId)
Browse all namespaces of a Kubernetes cluster

API to get all namespaces of a Kubernetes cluster with clusterId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| clusterId is the ID of the Kubernetes cluster client | 

### Return type

[**KubernetesBrowseClusterResp**](KubernetesBrowseClusterResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseKubernetesStorageClass**
> KubernetesBrowseClusterResp BrowseKubernetesStorageClass(ctx, clusterId)
Browse all storage classes of a Kubernetes cluster

API to get all StorageClasses of a Kubernetes cluster with clusterId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| clusterId is the ID of the Kubernetes cluster client | 

### Return type

[**KubernetesBrowseClusterResp**](KubernetesBrowseClusterResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseNamespaceApplications**
> KubernetesBrowseClusterResp BrowseNamespaceApplications(ctx, clusterId, nameSpace)
Content browse applications in a namespace

API to browse applications in a namespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| clusterId is the ID of the Kubernetes cluster client | 
  **nameSpace** | **string**| Name of the namespace to browse for content | 

### Return type

[**KubernetesBrowseClusterResp**](KubernetesBrowseClusterResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseNamespaceLabels**
> KubernetesBrowseClusterResp BrowseNamespaceLabels(ctx, clusterId, nameSpace)
Content browse labels in a namespace

API to browse labels in a namespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| clusterId is the ID of the Kubernetes cluster client | 
  **nameSpace** | **string**| Name of the namespace to browse for content | 

### Return type

[**KubernetesBrowseClusterResp**](KubernetesBrowseClusterResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BrowseNamespaceVolumes**
> KubernetesBrowseClusterResp BrowseNamespaceVolumes(ctx, clusterId, nameSpace)
Content browse volumes in a namespace

API to browse volumes in a namespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| clusterId is the ID of the Kubernetes cluster client | 
  **nameSpace** | **string**| Name of the namespace to browse for content | 

### Return type

[**KubernetesBrowseClusterResp**](KubernetesBrowseClusterResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateKubernetesApplicationGroup**
> IdName CreateKubernetesApplicationGroup(ctx, optional)
Create a new Kubernetes application group

API to create new Kubernetes application group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***KubernetesApiCreateKubernetesApplicationGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiCreateKubernetesApplicationGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateApplicationGroupRequest**](CreateApplicationGroupRequest.md)|  | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateKubernetesClusterOp**
> CreateKubernetesClusterResponse CreateKubernetesClusterOp(ctx, optional)
Create a new Kubernetes cluster client

API to create new Kubernetes cluster client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***KubernetesApiCreateKubernetesClusterOpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiCreateKubernetesClusterOpOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateKubernetesClusterRequest**](CreateKubernetesClusterRequest.md)|  | 

### Return type

[**CreateKubernetesClusterResponse**](CreateKubernetesClusterResponse.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDeconfiguredKubernetesCluster**
> GenericResp DeleteDeconfiguredKubernetesCluster(ctx, clusterId)
Delete a deconfigured Kubernetes cluster client

API to delete a deconfigured Kubernetes cluster with clusterId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| clusterId is the ID of the Kubernetes cluster client | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteKubernetesAppGroup**
> GenericResp DeleteKubernetesAppGroup(ctx, applicationGroupId)
Delete a Kubernetes application group

API to delete a Kubernetes application group with applicationGroupId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationGroupId** | **int32**| applicationGroupId is the ID of the Kubernetes application group | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoApplicationBrowse**
> KubernetesAppGroupApplicationBrowseResp DoApplicationBrowse(ctx, applicationGroupId, optional)
Browse backed up Applications of the application group

API to browse backed up applications for an application group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationGroupId** | **int32**| ID of the application group to browse for | 
 **optional** | ***KubernetesApiDoApplicationBrowseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiDoApplicationBrowseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number for number of results in pagination | 
 **limit** | **optional.Int32**| Page limit for number of results in pagination | 
 **fromTime** | **optional.Int32**| Browse from a specific time (in epochs) | 
 **toTime** | **optional.Int32**| Browse till a specific time (in epochs) | 

### Return type

[**KubernetesAppGroupApplicationBrowseResp**](KubernetesAppGroupApplicationBrowseResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoApplicationRestore**
> CreateTaskRespforBackup DoApplicationRestore(ctx, applicationGroupId, optional)
Full Application Restore - Restore backed up Applications of the application group

API to run Full Application Restore for an application group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationGroupId** | **int32**| ID of the application group to restore from | 
 **optional** | ***KubernetesApiDoApplicationRestoreOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiDoApplicationRestoreOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of KubernetesAppGroupApplicationRestore**](KubernetesAppGroupApplicationRestore.md)|  | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoNamespaceBrowse**
> KubernetesAppGroupNamespaceBrowseResp DoNamespaceBrowse(ctx, applicationGroupId, optional)
Browse backed up Namespaces of the application group

API to browse backed up namespaces for an application group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationGroupId** | **int32**| ID of the application group to browse for | 
 **optional** | ***KubernetesApiDoNamespaceBrowseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiDoNamespaceBrowseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number for number of results in pagination | 
 **limit** | **optional.Int32**| Page limit for number of results in pagination | 
 **fromTime** | **optional.Int32**| Browse from a specific time (in epochs) | 
 **toTime** | **optional.Int32**| Browse till a specific time (in epochs) | 

### Return type

[**KubernetesAppGroupNamespaceBrowseResp**](KubernetesAppGroupNamespaceBrowseResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoNamespaceRestore**
> CreateTaskRespforBackup DoNamespaceRestore(ctx, applicationGroupId, optional)
Namespace Level Restore - Restore backed up Namespaces of the application group

API to run Namespace Level Restore for an application group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationGroupId** | **int32**| ID of the application group to restore from | 
 **optional** | ***KubernetesApiDoNamespaceRestoreOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiDoNamespaceRestoreOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of KubernetesAppGroupNamespaceRestore**](KubernetesAppGroupNamespaceRestore.md)|  | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllKubernetesAppGroups**
> GetApplicationGroupsList GetAllKubernetesAppGroups(ctx, optional)
Get all Kubernetes application groups

API to get all Kubernetes application groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***KubernetesApiGetAllKubernetesAppGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiGetAllKubernetesAppGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **optional.Int32**| clusterId is the ID of the Kubernetes cluster client to filter Application Groups for | 

### Return type

[**GetApplicationGroupsList**](GetApplicationGroupsList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllKubernetesApps**
> GetApplicationsList GetAllKubernetesApps(ctx, )
Get all Kubernetes applications

API to get all Kubernetes applications

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetApplicationsList**](GetApplicationsList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllKubernetesClustersOp**
> GetClusterResp GetAllKubernetesClustersOp(ctx, )
Get all Kubernetes clusters

API to get all Kubernetes clusters

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetClusterResp**](GetClusterResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationGroupDetails**
> GetApplicationGroupDetails GetApplicationGroupDetails(ctx, applicationGroupId)
Get details of a Kubernetes application group

API to get details of a Kubernetes application group with applicationGroupId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationGroupId** | **int32**| applicationGroupId is the ID of the Kubernetes application group | 

### Return type

[**GetApplicationGroupDetails**](GetApplicationGroupDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKubernetesAppDetails**
> KubernetesApplicationDetails GetKubernetesAppDetails(ctx, appGUID)
Get Kubernetes applications details

API to get details of a Kubernetes application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appGUID** | **string**| GUID of the Application to get details | 

### Return type

[**KubernetesApplicationDetails**](KubernetesApplicationDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKubernetesClusterDetails**
> GetClusterDetailsResp GetKubernetesClusterDetails(ctx, clusterId)
Get details of a Kubernetes Cluster client

API to get details of a Kubernetes clusters with clusterId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**|  | 

### Return type

[**GetClusterDetailsResp**](GetClusterDetailsResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PreviewApplicationGroup**
> KubernetesApplicationGroupPreviewResp PreviewApplicationGroup(ctx, clusterId, optional)
Preview application group content for discovery of applications

API to preview application content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| clusterId of the Kubernetes client to do a Preview | 
 **optional** | ***KubernetesApiPreviewApplicationGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiPreviewApplicationGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of KubernetesApplicationGroupPreviewReq**](KubernetesApplicationGroupPreviewReq.md)|  | 

### Return type

[**KubernetesApplicationGroupPreviewResp**](KubernetesApplicationGroupPreviewResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReconfigureKubernetesCluster**
> GenericRespWithWarning ReconfigureKubernetesCluster(ctx, clusterId)
Reconfigure a retired Kubernetes cluster client

API to reconfigure a Kubernetes cluster with clusterId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| clusterId is the ID of the Kubernetes cluster client | 

### Return type

[**GenericRespWithWarning**](GenericRespWithWarning.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetireKubernetesCluster**
> GenericResp RetireKubernetesCluster(ctx, clusterId)
Retire and release license of a Kubernetes clueter client

API to retire a Kubernetes cluster with clusterId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| clusterId is the ID of the Kubernetes cluster client | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunApplicationBackup**
> CreateTaskRespforBackup RunApplicationBackup(ctx, appGUID, optional)
Run backup for an application

API to run backup for an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appGUID** | **string**| GUID of the Application to run backup for | 
 **optional** | ***KubernetesApiRunApplicationBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiRunApplicationBackupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backupLevel** | **optional.String**| Backup level , Default :Incremental | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunApplicationGroupBackup**
> CreateTaskRespforBackup RunApplicationGroupBackup(ctx, applicationGroupId, optional)
Run backup for an application group

API to run backup for an application group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationGroupId** | **int32**| applicationGroupId is the ID of the Kubernetes application group | 
 **optional** | ***KubernetesApiRunApplicationGroupBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiRunApplicationGroupBackupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backupLevel** | **optional.String**| Backup level , Default :Incremental | 

### Return type

[**CreateTaskRespforBackup**](CreateTaskRespforBackup.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateKubernetesAppGroupOp**
> GenericResp UpdateKubernetesAppGroupOp(ctx, applicationGroupId, optional)
Update properties of a Kubernetes application group

API to update properties of a Kubernetes application group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationGroupId** | **int32**| applicationGroupId is the ID of the Kubernetes application group | 
 **optional** | ***KubernetesApiUpdateKubernetesAppGroupOpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiUpdateKubernetesAppGroupOpOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateApplicationGroupRequest**](UpdateApplicationGroupRequest.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateKubernetesProperties**
> GenericResp UpdateKubernetesProperties(ctx, clusterId, optional)
Update properties of a Kubernetes cluster client

API to update the properties of a Kubernetes cluster with clusterId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| clusterId is the ID of the Kubernetes cluster client | 
 **optional** | ***KubernetesApiUpdateKubernetesPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiUpdateKubernetesPropertiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateKubernetesClusterRequest**](UpdateKubernetesClusterRequest.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V4KubernetesClusterIdModifierGet**
> K8sRestoreModifierApiResp V4KubernetesClusterIdModifierGet(ctx, clusterId)
Fetch all the modifiers

Fetch all the modifiers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| Id of the cluster whose modifier has to be listed. | 

### Return type

[**K8sRestoreModifierApiResp**](K8sRestoreModifierAPIResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V4KubernetesClusterIdModifierModifierNameDelete**
> GenericResp V4KubernetesClusterIdModifierModifierNameDelete(ctx, clusterId, modifierName)
Delete a Modifier

Delete a Modifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| Id of the cluster whose modifier has to be deleted. | 
  **modifierName** | **string**| Name of the modifier to be deleted. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V4KubernetesClusterIdModifierModifierNameGet**
> K8sRestoreModifierApiResp V4KubernetesClusterIdModifierModifierNameGet(ctx, clusterId, modifierName)
Fetch(Read) Modfier details

Fetch(Read) Modfier details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| Id of the cluster whose modifier has to be read. | 
  **modifierName** | **string**| Name of the modifier to be read. | 

### Return type

[**K8sRestoreModifierApiResp**](K8sRestoreModifierAPIResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V4KubernetesClusterIdModifierModifierNamePut**
> GenericResp V4KubernetesClusterIdModifierModifierNamePut(ctx, clusterId, modifierName, optional)
Edit a Restore modifier

Edit a Restore modifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| Id of the cluster whose modifier has to be modified. | 
  **modifierName** | **string**| Name of the modifier to be modified. | 
 **optional** | ***KubernetesApiV4KubernetesClusterIdModifierModifierNamePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiV4KubernetesClusterIdModifierModifierNamePutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of K8sRestoreModifierApiReq**](K8sRestoreModifierApiReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V4KubernetesClusterIdModifierPost**
> GenericResp V4KubernetesClusterIdModifierPost(ctx, clusterId, optional)
Create a Modifier

Create a Modifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| Id of the cluster whose modifier has to be created. | 
 **optional** | ***KubernetesApiV4KubernetesClusterIdModifierPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiV4KubernetesClusterIdModifierPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of K8sRestoreModifierApiReq**](K8sRestoreModifierApiReq.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V4KubernetesClusterIdModifierTestPost**
> K8sRestoreModifierApiResp V4KubernetesClusterIdModifierTestPost(ctx, clusterId, optional)
Test a Yaml paylod against a restore modfier

Test a Yaml paylod against a restore modfier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| Id of the cluster whose modifier needs to be tested | 
 **optional** | ***KubernetesApiV4KubernetesClusterIdModifierTestPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiV4KubernetesClusterIdModifierTestPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TestK8sModifierReq**](TestK8sModifierReq.md)|  | 

### Return type

[**K8sRestoreModifierApiResp**](K8sRestoreModifierAPIResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

