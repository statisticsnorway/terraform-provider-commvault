# CreateKubernetesClusterRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessNodes** | [**[]IdNameType**](IdNameType.md) |  | [default to null]
**Name** | **string** | Name of the Kubernetes Cluster | [default to null]
**ServiceType** | [***KubernetesServiceTypes**](KubernetesServiceTypes.md) |  | [optional] [default to null]
**EtcdProtection** | [***EtcdProtectionItem**](EtcdProtectionItem.md) |  | [optional] [default to null]
**ApiServer** | **string** | API Server Endpoint of the cluster | [default to null]
**ServiceAccount** | **string** | Name of the Service Account to authenticate with the cluster | [default to null]
**ServiceToken** | **string** | Secret token to authenticate with the cluster | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

