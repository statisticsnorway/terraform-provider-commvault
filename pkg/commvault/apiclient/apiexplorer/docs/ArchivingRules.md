# ArchivingRules

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartCleaningIfLessThan** | **int32** | When free disk space falls below specified amount (in percentage), start cleaning up the disk | [optional] [default to 50]
**StopCleaningIfupto** | **int32** | When free disk space more than specified amount (in percentage), stop cleaning up the disk | [optional] [default to 80]
**FileAccessTimeOlderThan** | **int32** | To archive files based on the last accessed date of each file within the folder, specify the number of days | [optional] [default to 90]
**FileModifiedTimeOlderThan** | **int32** | To archive files based on the last modified date of each file within the folder, specify the number of days | [optional] [default to 90]
**FileCreatedTimeOlderThan** | **int32** | To archive files based on the time the files were created within the folder, specify the number of days | [optional] [default to 0]
**FileSizeGreaterThan** | **int32** | To archive files based on the size of the file, specify the minimum file size in KB. All files whose size ranges between the minimum and maximum values are archived. | [optional] [default to 1024]
**MaximumFileSize** | **int32** | To archive files based on the size of the file, specify the maximum file size in KB. All files whose size ranges between the minimum and maximum values are archived. | [optional] [default to 0]
**ArchiveReadOnlyFiles** | **bool** | To archive files based on the Read-Only attribute, set to TRUE | [optional] [default to false]
**AfterArchiving** | [***ArchivingRulesFile**](ArchivingRulesFile.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

