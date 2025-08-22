# UpdateHypervisorGroupVCloud

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HypervisorType** | **string** |  | [default to null]
**VOrganizationName** | **string** | Organization name of VCloud Director | [optional] [default to null]
**VOrganizationGuid** | **string** | Organization Guid of VCloud Director | [optional] [default to null]
**HostName** | **string** | VCloud Director Name | [optional] [default to null]
**VcenterHostName** | **string** | Vcenter hostname | [optional] [default to null]
**VcenterUserName** | **string** | Vcenter username | [optional] [default to null]
**VcenterPassword** | **string** | Vcenter password | [optional] [default to null]
**CreateOrgAccount** | **bool** | True if organization account to be used | [optional] [default to false]
**UserName** | **string** | vCloud username | [optional] [default to null]
**Password** | **string** | VCloud base 64 encoded password | [optional] [default to null]
**AccessNodes** | [**[]AccessNodeModel**](accessNodeModel.md) |  | [optional] [default to null]
**FbrUnixMediaAgent** | [***IdName**](IdName.md) |  | [optional] [default to null]
**ActivityControl** | [***ActivityControlOptions**](ActivityControlOptions.md) |  | [optional] [default to null]
**Security** | [***VmHypervisorSecurityProp**](VMHypervisorSecurityProp.md) |  | [optional] [default to null]
**NewName** | **string** | The name of the hypervisor that has to be changed | [optional] [default to null]
**Settings** | [***HypervisorSettings**](hypervisorSettings.md) |  | [optional] [default to null]
**Credentials** | [***IdName**](IdName.md) |  | [optional] [default to null]
**SkipCredentialValidation** | **bool** | if credential validation has to be skipped. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

