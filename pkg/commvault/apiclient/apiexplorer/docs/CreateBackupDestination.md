# CreateBackupDestination

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of backup destination | [optional] [default to null]
**StoragePool** | [***IdName**](IdName.md) |  | [default to null]
**StorageType** | [***StorageType**](StorageType.md) |  | [optional] [default to null]
**NetAppCloudTarget** | **bool** | Only for snap copy. Enabling this changes SVM Mapping  to NetApp cloud targets only. | [optional] [default to null]
**IsSnapCopy** | **bool** | Is this a snap copy? If isMirrorCopy is not set, then default is Vault/Replica. | [optional] [default to null]
**IsMirrorCopy** | **bool** | Is this a mirror copy? Only considered when isSnapCopy is true. | [optional] [default to null]
**Region** | [***IdName**](IdName.md) |  | [optional] [default to null]
**SourceCopy** | [***IdName**](IdName.md) |  | [optional] [default to null]
**BackupsToCopy** | [***PlanFullBackupType**](PlanFullBackupType.md) |  | [optional] [default to null]
**FullBackupTypesToCopy** | [***PlanFullBackupTypeToCopy**](PlanFullBackupTypeToCopy.md) |  | [optional] [default to null]
**BackupStartTime** | **int32** | Backup start time in seconds. The time is provided in unix time format. | [optional] [default to null]
**OptimizeForInstantClone** | **bool** | Flag to specify if primary storage is copy data management enabled. | [optional] [default to null]
**OverrideRetentionSettings** | **bool** | Tells if this copy should use storage pool retention period days or the retention defined for this copy. Set as true to use retention defined on this copy. | [optional] [default to null]
**RetentionRuleType** | [***RetentionRuleTypes**](RetentionRuleTypes.md) |  | [optional] [default to null]
**RetentionPeriodDays** | **int32** | Retention period in days. -1 can be specified for infinite retention. If this and snapRecoveryPoints both are not specified, this takes  precedence. | [optional] [default to 30]
**SnapRecoveryPoints** | **int32** | Number of snap recovery points for snap copy for retention. Can be specified instead of retention period in Days for snap copy. | [optional] [default to null]
**UseExtendedRetentionRules** | **bool** | Use extended retention rules | [optional] [default to null]
**ExtendedRetentionRules** | [***ExtendedRetentionRules**](ExtendedRetentionRules.md) |  | [optional] [default to null]
**Mappings** | [**[]SnapshotCopyMapping**](SnapshotCopyMapping.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

