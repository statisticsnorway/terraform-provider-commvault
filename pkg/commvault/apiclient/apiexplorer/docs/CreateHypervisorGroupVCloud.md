# CreateHypervisorGroupVCloud

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HypervisorType** | **string** |  | [default to null]
**VOrganizationName** | **string** | Organization name of VMware Cloud Director | [optional] [default to null]
**VOrganizationGuid** | **string** | Organization Guid of VMware Cloud Director | [optional] [default to null]
**VCloudDirector** | **string** | VMware Cloud Director Name | [optional] [default to null]
**HostName** | **string** | VMware Cloud Director hostname or organization guid if org client | [optional] [default to null]
**UserName** | **string** | VMware Cloud Director username | [optional] [default to null]
**Password** | **string** | VMware Cloud Director password | [optional] [default to null]
**CreateOrgAccount** | **bool** | True if organization account to be used | [optional] [default to false]
**AutoCompany** | **string** | indicates the mode for company association | [optional] [default to null]
**CompanyId** | **int32** | company id is required if use existing mode is selected for org client | [optional] [default to null]
**CompanyName** | **string** | company name is required for org client | [optional] [default to null]
**Name** | **string** | The name of the hypervisor group being created | [default to null]
**SkipCredentialValidation** | **bool** | if credential validation has to be skipped. | [optional] [default to false]
**AccessNodes** | [**[]AccessNodeModel**](accessNodeModel.md) |  | [optional] [default to null]
**Credentials** | [***IdName**](IdName.md) |  | [optional] [default to null]
**EtcdProtection** | [***EtcdProtectionItem**](EtcdProtectionItem.md) |  | [optional] [default to null]
**PlanEntity** | [***IdName**](IdName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

