# OverrideReplicationOptionsVmWareCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceVM** | [***NameGuid**](NameGUID.md) |  | [default to null]
**VmDisplayName** | **string** | Enter the display name for the destination VM | [default to null]
**DestinationHost** | **string** | Name of host where the destination VM will reside | [default to null]
**Datastore** | **string** | The datastore to be used for virtual machine disks | [default to null]
**ResourcePool** | **string** | The resource pool that will be used on the destination VM | [default to null]
**VmFolder** | **string** | If the destination VM has to reside inside a folder, specify the name here | [optional] [default to null]
**NetworkSettings** | [**[]NetworkSetting**](NetworkSetting.md) | Mapping between a source network and a destination network | [optional] [default to null]
**IpAddressSettings** | [**[]IpAddressSettingVmWareCreate**](IpAddressSettingVMWareCreate.md) | A mapping of IP between the source and the destination VM with an option to configure either a dynamic or a static IP. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

