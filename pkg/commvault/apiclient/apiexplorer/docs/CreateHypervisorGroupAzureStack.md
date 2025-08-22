# CreateHypervisorGroupAzureStack

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HypervisorType** | **string** |  | [default to null]
**ResourceManagerURL** | **string** | resource Manager URL for Azure Stack | [optional] [default to null]
**TenantId** | **string** | Tenant id of Azure login Application | [default to null]
**SubscriptionId** | **string** | subscription id of Azure  | [default to null]
**ApplicationId** | **string** | Application id of Azure login Application | [default to null]
**ApplicationPassword** | **string** | Application Password of Azure login Application | [default to null]
**Name** | **string** | The name of the hypervisor group being created | [default to null]
**SkipCredentialValidation** | **bool** | if credential validation has to be skipped. | [optional] [default to false]
**AccessNodes** | [**[]AccessNodeModel**](accessNodeModel.md) |  | [optional] [default to null]
**Credentials** | [***IdName**](IdName.md) |  | [optional] [default to null]
**EtcdProtection** | [***EtcdProtectionItem**](EtcdProtectionItem.md) |  | [optional] [default to null]
**PlanEntity** | [***IdName**](IdName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

