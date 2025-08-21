# CreateDiskStorage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Disk Storage to be created. | [default to null]
**Storage** | [**[]Path**](Path.md) | A list of backup locations can be provided for the storage pool being created. | [default to null]
**EnableDeduplication** | **bool** | enables or disables deduplication | [default to null]
**DeduplicationDBStorage** | [**[]DedupePath**](DedupePath.md) | A list of dedupe locations can be provided for the storage pool being created. This provides an efficient way to save/store data by eliminating duplicate blocks of data during backups. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

