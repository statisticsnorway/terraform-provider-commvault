# DdbDiskInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskId** | **int32** | Id of the DDB disk | [optional] [default to null]
**DiskPath** | **string** | File path of the DDB disk | [optional] [default to null]
**ClientInfo** | [***IdName**](IdName.md) |  | [optional] [default to null]
**PartitionList** | [**[]DdbSubStoreInfo**](DDBSubStoreInfo.md) | List of DDB partitions hosted on this disk | [optional] [default to null]
**NumOfPartitions** | **int32** | Number of DDB partitions hosted on this disk | [optional] [default to null]
**Status** | **string** | Satus of the disk | [optional] [default to null]
**ConsumedSpaceMB** | **int64** | Amount of space consumed by the DDB partitions hosted on this disk | [optional] [default to null]
**FreeSpaceMB** | **int64** | Available usable free space on the disk path | [optional] [default to null]
**TotalSpaceMB** | **int64** | Total space capacity of the disk path | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

