# KubernetesAppGroupNamespaceRestore

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoreApplications** | **bool** | Restore Applications of the Namespace | [optional] [default to true]
**Namespaces** | [**[]KubernetesNamespaceRestoreItem**](KubernetesNamespaceRestoreItem.md) |  | [optional] [default to null]
**InPlace** | **bool** | Run In-Place restore job | [optional] [default to null]
**DestinationCluster** | [***IdName**](IdName.md) |  | [optional] [default to null]
**AccessNode** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Overwrite** | **bool** | Overwrite if already present | [optional] [default to false]
**TimeRange** | [***TimeRange**](TimeRange.md) |  | [optional] [default to null]
**Modifier** | **string** | Specify list of resource modifier YAML as a string | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

