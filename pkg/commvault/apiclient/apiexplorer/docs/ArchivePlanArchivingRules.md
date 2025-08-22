# ArchivePlanArchivingRules

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileTimestampMethod** | [***ArchivePlanArchivingTimestamp**](ArchivePlanArchivingTimestamp.md) |  | [optional] [default to null]
**FileTimestamp** | **int32** | To archive files based on the last accessed or modified date of each file within the folder, specify the number of days. Should be supplied with fileTimestampMethod. | [optional] [default to 90]
**FileSize** | **int32** | To archive files based on the size of the file, specify the minimum file size in KB. | [optional] [default to 1]
**AfterArchiving** | [***ArchivingRulesFile**](ArchivingRulesFile.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

