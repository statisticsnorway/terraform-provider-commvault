# ReplicationGroupAdvOptionsVmWare

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidateDestinationVM** | **bool** | Validates that the destination VM is bootable by powering it on and then powering off | [optional] [default to true]
**DiskProvisioning** | **string** | Select the disk provisioning type for the destination VM. The values are case sensitive. | [optional] [default to DISK_PROVISIONING.AUTO]
**TransportMode** | **string** | Choose transport mode based on environment. Values are case sensitive | [optional] [default to TRANSPORT_MODE.AUTO]
**UnconditionalOverwrite** | **bool** | Replace an existing virtual machine with the same name in the target location | [optional] [default to false]
**SnapshotsToRetain** | **int32** | Number of snapshots to retain on destination VM | [optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

