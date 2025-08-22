# OverrideReplicationOptionsVmWare

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmDisplayName** | **string** | Enter the display name for the destination VM | [optional] [default to null]
**DestinationHost** | **string** | Host for the destination VM | [optional] [default to null]
**Datastore** | **string** | Select a datastore to be used for virtual machine disks | [optional] [default to null]
**ResourcePool** | **string** | Select a resource pool for the destination VM | [optional] [default to null]
**VmFolder** | **string** | VM folder replication | [optional] [default to null]
**NetworkSettings** | [**[]NetworkSetting**](NetworkSetting.md) | Mapping between a source network and a destination network | [optional] [default to null]
**IpAddressSettings** | [**[]IpAddressSetting**](IpAddressSetting.md) | Customize IP address settings | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

