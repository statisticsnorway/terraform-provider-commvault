# CloudDestinationOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | [***Region**](Region.md) |  | [optional] [default to null]
**KeyPair** | **string** | Name of key pair that associates with the VM for authentication | [optional] [default to null]
**RestoreAsManagedVM** | **bool** | If true, restore creates the destination VM as a managed VM in Azure | [optional] [default to null]
**MachineType** | [***IdName**](IdName.md) |  | [optional] [default to null]
**SoleTenancyNodes** | [**[]NameValue**](NameValue.md) |  | [optional] [default to null]
**Tags** | [**[]NameValue**](NameValue.md) |  | [optional] [default to null]
**AvailabilityZone** | **string** | Availability zone name for the destination instance | [optional] [default to null]
**VmInstanceType** | **string** | Instance type that provides the available cpu cores and memory to the machine | [optional] [default to null]
**EncryptionKey** | [***IdName**](IdName.md) |  | [optional] [default to null]
**VolumeType** | **string** | Volume type of destination VM | [optional] [default to null]
**TestDiskType** | **string** | Disk / Volume type for test failover VM | [optional] [default to null]
**InstanceTypes** | **[]string** |  | [optional] [default to null]
**PublicIP** | **bool** | If true, public IP address are configured for destination VMs | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

