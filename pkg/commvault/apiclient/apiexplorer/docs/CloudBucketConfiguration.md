# CloudBucketConfiguration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | When true, means mount path is enabled | [optional] [default to null]
**DisableBackupLocationForFutureBackups** | **bool** | When true, prevents new data writes to backup location by changing number of writers to zero | [optional] [default to null]
**PrepareForRetirement** | **bool** | When true, the deduplicated blocks in the mount path will not be referenced when there are multiple mount paths in the library. | [optional] [default to null]
**StorageAcceleratorCredentials** | [***IdName**](IdName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

