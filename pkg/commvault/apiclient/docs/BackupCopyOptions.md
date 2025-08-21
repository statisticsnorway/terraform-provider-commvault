# BackupCopyOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupTypeToCopy** | [***BackupCopyFullBackupType**](BackupCopyFullBackupType.md) |  | [optional] [default to null]
**BackupFullToCopy** | [***PlanFullBackupTypeToCopy**](PlanFullBackupTypeToCopy.md) |  | [optional] [default to null]
**StartTime** | **int32** | Snapshots to be copied from a particular time in unix time format. By default, 0 means since the inception of the snap copy. | [optional] [default to 0]
**EnableAlert** | **bool** | Flag to enable backup copy fallen behind alert | [optional] [default to null]
**AlertInHours** | **int32** | Alert to throw when backup copy falls behind in hours | [optional] [default to 48]
**Action** | **string** | Which type of action should be followed if backup copy falls behind | [optional] [default to null]
**SkipAfterThresholdDays** | **int32** | The allowable delay in days before a backup copy job is considered overdue | [optional] [default to 14]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

