# UpdateHypervisorGroupAzureStack

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HypervisorType** | **string** |  | [default to null]
**TenantId** | **string** | Tenant id of Azure Stack login Application | [optional] [default to null]
**ServerName** | **string** | Client Name to Update | [optional] [default to null]
**SubscriptionId** | **string** | subscription id of Azure Stack | [optional] [default to null]
**ApplicationId** | **string** | Application id of Azure login Application | [optional] [default to null]
**ApplicationPassword** | **string** | Application Password of Azure login Application | [optional] [default to null]
**ResourceManagerURL** | **string** | Resource manager URL for Azure Stack client | [optional] [default to null]
**UseManagedIdentity** | **bool** | Use managed identities for Azure stack authentication | [optional] [default to null]
**AccessNodes** | [**[]AccessNodeModel**](accessNodeModel.md) |  | [optional] [default to null]
**FbrUnixMediaAgent** | [***IdName**](IdName.md) |  | [optional] [default to null]
**ActivityControl** | [***ActivityControlOptions**](ActivityControlOptions.md) |  | [optional] [default to null]
**Security** | [***VmHypervisorSecurityProp**](VMHypervisorSecurityProp.md) |  | [optional] [default to null]
**NewName** | **string** | The name of the hypervisor that has to be changed | [optional] [default to null]
**Settings** | [***HypervisorSettings**](hypervisorSettings.md) |  | [optional] [default to null]
**Credentials** | [***IdName**](IdName.md) |  | [optional] [default to null]
**SkipCredentialValidation** | **bool** | if credential validation has to be skipped. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

