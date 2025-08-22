# HypervisorListResp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | This gives the id of the Hypervisor. | [optional] [default to null]
**Name** | **string** | This give the name of the Hypervisor. | [optional] [default to null]
**DisplayName** | **string** | This gives the name of the Hypervisor as it shown on the admin console or GUI. | [optional] [default to null]
**HostName** | **string** | This give the host name of the Hypervisor. | [optional] [default to null]
**HypervisorType** | [***VsVendor**](VSVendor.md) |  | [optional] [default to null]
**Status** | [***RetireClientPhase**](RetireClientPhase.md) |  | [optional] [default to null]
**Version** | **string** | VMware Vcenter Version | [optional] [default to null]
**Instance** | [***IdName**](IdName.md) |  | [optional] [default to null]
**ClientActivityControl** | [***ActivityControlOptionsProp**](ActivityControlOptionsProp.md) |  | [optional] [default to null]
**Tags** | [**[]IdNameValue**](IdNameValue.md) | Array of all the tags associated with hypervisor. | [optional] [default to null]
**Commcell** | [***CommcellInfo**](CommcellInfo.md) |  | [optional] [default to null]
**ManageSnapshot** | [***IdName**](IdName.md) |  | [optional] [default to null]
**IsManagedIdentity** | **bool** | True if hypervisor is configured using Azure Managed Service Identity. | [optional] [default to null]
**UseHostedInfrastructure** | **bool** | Use Metallic hosted infrastructure | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

