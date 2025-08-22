# DestinationOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationHost** | **string** | Destination host for the VM to deploy | [optional] [default to null]
**DataStore** | **string** | Datastore for the destination VM to store the disks and its config files. In-case of Microsoft Hyper-V, datastore refers to the destination folder for restore when default folder is not set. | [optional] [default to null]
**DataStoreClusterName** | **string** | Datastore cluster for the destination VM to store disks and config files in the datastore associated with the datastore cluster. | [optional] [default to null]
**ResourcePoolPath** | **string** | Resource pool for the destination VM | [optional] [default to null]
**VmFolder** | **string** | Folder path where the destination VM will be located | [optional] [default to null]
**IamRole** | [***IdName**](IdName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

