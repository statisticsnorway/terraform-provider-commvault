# ReplicationConfigurationAzure

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceVM** | **string** | Source VM | [optional] [default to null]
**DestinationVM** | **string** | Destination VM | [optional] [default to null]
**ResourceGroup** | **string** | Resource group in which the converted virtual machine should be created. | [optional] [default to null]
**StorageAccount** | **string** | Azure Standard or Premium general-purpose storage account. | [optional] [default to null]
**VmSize** | **string** | VM size | [optional] [default to null]
**CreatePublicIP** | **bool** | Create a public IP. | [optional] [default to null]
**RestoreAsManagedVm** | **bool** | Restore the VM as a managed disk for the destination VMs. | [optional] [default to null]
**OverrideReplicationOptions** | [***OverrideReplicationOptionsAzure**](OverrideReplicationOptionsAzure.md) |  | [optional] [default to null]
**DiskType** | **string** | Type of the disk | [optional] [default to null]
**SourceVMGuid** | **string** | GUID of the source VM | [optional] [default to null]
**Region** | **string** | region of the VM | [optional] [default to null]
**TestFailoverVmSize** | **string** | Vm size to be used during the test failover operation | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

