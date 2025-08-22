# PlanSnapshotOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableBackupCopy** | **bool** | Flag to enable backup copy | [optional] [default to null]
**BackupCopyRPOMins** | **int32** | Backup copy RPO in minutes | [optional] [default to null]
**BackupCopyFrequency** | [***BackupFrequencyPattern**](BackupFrequencyPattern.md) |  | [optional] [default to null]
**EnableSnapCatalog** | **bool** | Flag to enable deferred snapshot cataloging | [optional] [default to null]
**BackupCopyOptions** | [***BackupCopyOptions**](BackupCopyOptions.md) |  | [optional] [default to null]
**SyncError** | **string** | Error message to indicate the cause of sync issue (if any) between CommServ DB and Command center with respect to snapshot options. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

