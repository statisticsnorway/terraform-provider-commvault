# VirtualMachineSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The display name of the VM. | [optional] [default to null]
**Vendor** | **string** | The hypervisor where the VM is located. | [optional] [default to null]
**Hypervisor** | [***IdName**](IdName.md) |  | [optional] [default to null]
**VmGroup** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Os** | **string** | The operating system version for the VM or instance. | [optional] [default to null]
**Host** | **string** | The host name for the computer where the source virtual machine or instance resides. | [optional] [default to null]
**VmSize** | **int32** | The total space allocated(in bytes) for the virtual machine. | [optional] [default to null]
**Status** | **string** | The status of the backup for the VM, instance, or  container. | [optional] [default to null]
**UUID** | **string** | The globally unique identifier for the VM client. | [optional] [default to null]
**CommcellName** | **string** | Name of the commcell the VM belongs | [optional] [default to null]
**LastBackup** | [***LastBackup**](lastBackup.md) |  | [optional] [default to null]
**ApplicationSize** | **int32** | The amount of data being protected for the VM client(in bytes) | [optional] [default to null]
**Plan** | [***PlanIdNameType**](PlanIdNameType.md) |  | [optional] [default to null]
**SLA** | [***VmSla**](VMSla.md) |  | [optional] [default to null]
**Company** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Tags** | [**[]IdNameValue**](IdNameValue.md) |  | [optional] [default to null]
**Commcell** | [***CommcellInfo**](CommcellInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

