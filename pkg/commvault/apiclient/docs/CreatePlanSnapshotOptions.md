# CreatePlanSnapshotOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapRecoveryPoints** | **int32** | Number of snap recovery points for default snap copy for retention. Can be specified instead of retention period in Days for default snap copy. | [optional] [default to null]
**RetentionPeriodDays** | **int32** | Retention period in days. -1 can be specified for infinite retention. If this and snapRecoveryPoints both are not specified, this takes precedence. | [optional] [default to 30]
**EnableBackupCopy** | **bool** | Flag to enable backup copy | [optional] [default to null]
**BackupCopyRPOMins** | **int32** | Backup copy RPO in minutes | [optional] [default to 240]
**EnableSnapCatalog** | **bool** | Flag to enable deferred snapshot cataloging | [optional] [default to null]
**BackupCopyOptions** | [***BackupCopyOptions**](BackupCopyOptions.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

