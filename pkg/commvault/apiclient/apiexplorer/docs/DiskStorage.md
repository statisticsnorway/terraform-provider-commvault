# DiskStorage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Id of the Disk Storage Pool | [optional] [default to null]
**Name** | **string** | Name of the Disk Storage Pool | [optional] [default to null]
**General** | [***DiskStorageGeneralInfo**](DiskStorageGeneralInfo.md) |  | [optional] [default to null]
**BackupLocations** | [**[]IdNameStatus**](IdNameStatus.md) | Gives backup location/mount path details for the storage pool. | [optional] [default to null]
**Encryption** | [***Encryption**](Encryption.md) |  | [optional] [default to null]
**Security** | [**[]SecurityAssoc**](SecurityAssoc.md) |  | [optional] [default to null]
**AssociatedPlanList** | [**[]IdName**](IdName.md) | Provides the list of plans associated with the storage pool. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

