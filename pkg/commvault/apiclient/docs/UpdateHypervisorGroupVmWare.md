# UpdateHypervisorGroupVmWare

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HypervisorType** | **string** |  | [default to null]
**VcenterHostName** | **string** | Vcenter hostname | [optional] [default to null]
**UserName** | **string** | Vcenter username | [optional] [default to null]
**Password** | **string** | Vcenter password | [optional] [default to null]
**AccessNodes** | [**[]AccessNodeModel**](accessNodeModel.md) |  | [optional] [default to null]
**FbrUnixMediaAgent** | [***IdName**](IdName.md) |  | [optional] [default to null]
**ActivityControl** | [***ActivityControlOptions**](ActivityControlOptions.md) |  | [optional] [default to null]
**Security** | [***VmHypervisorSecurityProp**](VMHypervisorSecurityProp.md) |  | [optional] [default to null]
**NewName** | **string** | The name of the hypervisor that has to be changed | [optional] [default to null]
**Settings** | [***HypervisorSettings**](hypervisorSettings.md) |  | [optional] [default to null]
**Credentials** | [***IdName**](IdName.md) |  | [optional] [default to null]
**SkipCredentialValidation** | **bool** | if credential validation has to be skipped. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

