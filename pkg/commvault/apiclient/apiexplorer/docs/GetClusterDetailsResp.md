# GetClusterDetailsResp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiServer** | **string** | API Server Endpoint of the cluster | [optional] [default to null]
**Version** | **string** | Kubernetes client version | [optional] [default to null]
**AccessNode** | [**[]IdName**](IdName.md) |  | [optional] [default to null]
**Tags** | [**[]IdNameValue**](IdNameValue.md) |  | [optional] [default to null]
**EtcdProtection** | [***GetEtcdProtectionItem**](GetEtcdProtectionItem.md) |  | [optional] [default to null]
**ApplicationCount** | [***KubernetesApplicationCountInfo**](KubernetesApplicationCountInfo.md) |  | [optional] [default to null]
**Instance** | [***IdName**](IdName.md) |  | [optional] [default to null]
**ClientGroups** | [**[]IdName**](IdName.md) |  | [optional] [default to null]
**Region** | [***IdName**](IdName.md) |  | [optional] [default to null]
**CommonProperties** | [***ClusterDetailsCommonProps**](ClusterDetailsCommonProps.md) |  | [optional] [default to null]
**ActivityControl** | [***ClusterActivityControlOptions**](ClusterActivityControlOptions.md) |  | [optional] [default to null]
**Options** | [***EditClusterAdvancedOptionsInfo**](EditClusterAdvancedOptionsInfo.md) |  | [optional] [default to null]
**CloudStorage** | [***KubernetesCloudStorage**](KubernetesCloudStorage.md) |  | [optional] [default to null]
**Id** | **int32** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**DisplayName** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

