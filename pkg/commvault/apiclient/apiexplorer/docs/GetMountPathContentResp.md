# GetMountPathContentResp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountpathName** | **string** | Name of mountpath | [optional] [default to null]
**TotalSizeOnMedia** | **int64** | Total size of the data stored on the mountpath. This is the size of the data after deduplication and compression, if employed. This amount also includes metadata information, so in some cases, it could be larger than the actual size of the backed up data. | [optional] [default to null]
**TotalDataWritten** | **int64** | Total amount of data written on mountpath | [optional] [default to null]
**IsRequiredByAuxiliaryCopy** | **bool** | This indicates if any of the jobs on this mountpath is required for auxiliary copy. | [optional] [default to null]
**RetainDataUntil** | **int64** | Maximum DateTime upto which the jobs on this mountpath is retained | [optional] [default to null]
**IsSingleInstanced** | **bool** | Indicates if mountpath has any DDB references | [optional] [default to null]
**HasSubclientLogCacheDump** | **bool** | Indicates if mountpath has any subclient lLog cache dump | [optional] [default to null]
**JobInfoList** | [**[]MountPathJobInfo**](MountPathJobInfo.md) | Jobs on mountpath | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

