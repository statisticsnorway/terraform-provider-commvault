# ReplicationGroupAdvOptionsVmWareCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidateDestinationVM** | **bool** | Validates that the destination VM is bootable by powering it on and then powering off | [optional] [default to true]
**DiskProvisioning** | **string** | The disk provisioning type for the destination VM. | [optional] [default to DISK_PROVISIONING.ORIGINAL]
**TransportMode** | **string** | The transport mode based on environment. | [optional] [default to TRANSPORT_MODE.AUTO]
**UnconditionalOverwrite** | **bool** | This will replace the instance at the destination if the instance with the same name already exists. | [optional] [default to false]
**SnapshotsToRetain** | **int32** | Number of snapshots to retain on destination VM. This is only applicable if snap engine is provided | [optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

