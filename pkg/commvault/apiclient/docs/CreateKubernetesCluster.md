# CreateKubernetesCluster

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HypervisorType** | **string** |  | [default to null]
**Endpointurl** | **string** | Endpoint url to connect | [default to null]
**ServiceName** | **string** | Service Name to connect in case authentication mode is service account | [optional] [default to null]
**SecretKey** | **string** | SecretKey to connect in case authentication mode is service account | [optional] [default to null]
**UserName** | **string** | Username to connect in case authentication mode is Username and password | [optional] [default to null]
**Password** | **string** | Username to connect in case authentication mode is Username and password | [optional] [default to null]
**K8ServiceType** | **string** |  | [optional] [default to null]
**CloudStorage** | [***KubernetesCloudStorage**](KubernetesCloudStorage.md) |  | [optional] [default to null]
**Name** | **string** | The name of the hypervisor group being created | [default to null]
**SkipCredentialValidation** | **bool** | if credential validation has to be skipped. | [optional] [default to false]
**AccessNodes** | [**[]AccessNodeModel**](accessNodeModel.md) |  | [optional] [default to null]
**Credentials** | [***IdName**](IdName.md) |  | [optional] [default to null]
**EtcdProtection** | [***EtcdProtectionItem**](EtcdProtectionItem.md) |  | [optional] [default to null]
**PlanEntity** | [***IdName**](IdName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

