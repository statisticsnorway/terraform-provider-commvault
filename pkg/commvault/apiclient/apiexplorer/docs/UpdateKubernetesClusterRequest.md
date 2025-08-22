# UpdateKubernetesClusterRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessNodes** | [**[]IdNameType**](IdNameType.md) |  | [optional] [default to null]
**Name** | **string** | Name of the Kubernetes Cluster | [optional] [default to null]
**ServiceType** | [***KubernetesServiceTypes**](KubernetesServiceTypes.md) |  | [optional] [default to null]
**EtcdProtection** | [***GetEtcdProtectionItem**](GetEtcdProtectionItem.md) |  | [optional] [default to null]
**Tags** | [**[]NameValue**](NameValue.md) | Modify or add tags on the cluster | [optional] [default to null]
**Region** | [***IdName**](IdName.md) |  | [optional] [default to null]
**ActivityControl** | [***EditClusterActivityControlOptions**](EditClusterActivityControlOptions.md) |  | [optional] [default to null]
**Options** | [***EditClusterAdvancedOptionsInfo**](EditClusterAdvancedOptionsInfo.md) |  | [optional] [default to null]
**ApiServer** | **string** | API Server Endpoint of the cluster | [optional] [default to null]
**ServiceAccount** | **string** | Name of the Service Account to authenticate with the cluster | [optional] [default to null]
**ServiceToken** | **string** | Secret token to authenticate with the cluster | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

