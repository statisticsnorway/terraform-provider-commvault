# UpdatePlanBackupDestination

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewName** | **string** | New name for backup destination | [optional] [default to null]
**SourceCopy** | [***IdName**](IdName.md) |  | [optional] [default to null]
**BackupsToCopy** | [***PlanFullBackupType**](PlanFullBackupType.md) |  | [optional] [default to null]
**BackupStartTime** | **int32** | Specify the Backup start time in seconds. The time is provided in unix time format. | [optional] [default to null]
**EnableDataAging** | **bool** | Tells if this copy has data aging enabled | [optional] [default to null]
**OverrideRetentionSettings** | **bool** | Tells if this copy should use storage pool retention period days or the retention defined for this copy. Set as true to use retention defined on this copy. | [optional] [default to null]
**RetentionRuleType** | [***RetentionRuleTypes**](RetentionRuleTypes.md) |  | [optional] [default to null]
**RetentionPeriodDays** | **int32** | Retention period in days. -1 can be specified for infinite retention. | [optional] [default to null]
**SnapRecoveryPoints** | **int32** | Number of snap recovery points for snap copy for retention. Can be specified instead of retention period in Days for snap copy. | [optional] [default to null]
**UseExtendedRetentionRules** | **bool** | Use extended retention rules. Must specify if updating extended retention rules. | [optional] [default to null]
**ExtendedRetentionRules** | [***ExtendedRetentionRules**](ExtendedRetentionRules.md) |  | [optional] [default to null]
**Mappings** | [**[]SnapshotCopyMapping**](SnapshotCopyMapping.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

