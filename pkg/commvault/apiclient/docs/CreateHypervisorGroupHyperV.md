# CreateHypervisorGroupHyperV

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HypervisorType** | **string** |  | [default to null]
**ServerName** | **string** | Hyper-V hostname  | [default to null]
**UserName** | **string** | Hyper-V userName  | [default to null]
**Password** | **string** | Hyper-V password  | [default to null]
**Name** | **string** | The name of the hypervisor group being created | [default to null]
**SkipCredentialValidation** | **bool** | if credential validation has to be skipped. | [optional] [default to false]
**AccessNodes** | [**[]AccessNodeModel**](accessNodeModel.md) |  | [optional] [default to null]
**Credentials** | [***IdName**](IdName.md) |  | [optional] [default to null]
**EtcdProtection** | [***EtcdProtectionItem**](EtcdProtectionItem.md) |  | [optional] [default to null]
**PlanEntity** | [***IdName**](IdName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

