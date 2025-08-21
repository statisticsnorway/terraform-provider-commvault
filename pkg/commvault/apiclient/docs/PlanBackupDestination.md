# PlanBackupDestination

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanBackupDestination** | [***IdName**](IdName.md) |  | [optional] [default to null]
**NetAppCloudTarget** | **bool** | Only for snap copy. Tells if the snap copy supports SVM Mapping to NetApp cloud targets only. | [optional] [default to null]
**IsImmutableSnapCopy** | **bool** | Only for snap copy. Tells if the snap copy has immutable option enabled. | [optional] [default to null]
**IsDefault** | **bool** | Is this a default backup destination? | [optional] [default to null]
**IsCopyInMaintenanceMode** | **bool** | Is this copy in Maintenance mode? | [optional] [default to null]
**IsCopyPromotionRequestSubmitted** | **bool** | Is copy promotion request submitted for this copy? | [optional] [default to null]
**IsConfiguredForReplication** | **bool** | Used if the copy is used for replication group | [optional] [default to null]
**IsSnapCopy** | **bool** | Is this a snap copy? | [optional] [default to null]
**IsMirrorCopy** | **bool** | Is this a mirror copy? | [optional] [default to null]
**IsAirgapCopy** | **bool** | Set to true if this copy uses Airgap storage | [optional] [default to null]
**IsSourceBackupCopy** | **bool** | Is this the source snap copy for backup copy operations? | [optional] [default to null]
**CopyType** | [***PlanCopyType**](PlanCopyType.md) |  | [optional] [default to null]
**CopyTypeName** | [***PlanCopyTypeName**](PlanCopyTypeName.md) |  | [optional] [default to null]
**CopyPrecedence** | **int32** | Order of backup destinaion copy created in storage policy | [optional] [default to null]
**StoragePool** | [***StoragePool**](StoragePool.md) |  | [optional] [default to null]
**StorageType** | [***StorageType**](StorageType.md) |  | [optional] [default to null]
**SourceCopy** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Region** | [***IdNameDisplayName**](IdNameDisplayName.md) |  | [optional] [default to null]
**BackupsToCopy** | [***PlanFullBackupType**](PlanFullBackupType.md) |  | [optional] [default to null]
**FullBackupTypesToCopy** | [***PlanFullBackupTypeToCopy**](PlanFullBackupTypeToCopy.md) |  | [optional] [default to null]
**BackupStartTime** | **int32** | Backup start time in number of seconds. The time is provided in unix time format. | [optional] [default to null]
**EnableDataAging** | **bool** | Tells if this copy has data aging enabled | [optional] [default to null]
**OverrideRetentionSettings** | **bool** | Tells if this copy should use storage pool retention period days or the retention defined for this copy. Set as true to use retention defined on this copy. | [optional] [default to null]
**RetentionRuleType** | [***RetentionRuleTypes**](RetentionRuleTypes.md) |  | [optional] [default to null]
**RetentionPeriodDays** | **int32** | Retention period in days | [optional] [default to null]
**SnapRecoveryPoints** | **int32** | Number of snap recovery points for snap copy for retention | [optional] [default to null]
**UseExtendedRetentionRules** | **bool** | Should extended retention rules be used | [optional] [default to null]
**ExtendedRetentionRules** | [***ExtendedRetentionRules**](ExtendedRetentionRules.md) |  | [optional] [default to null]
**Mappings** | [**[]SnapshotCopyMapping**](SnapshotCopyMapping.md) |  | [optional] [default to null]
**StorageTemplateTags** | [**[]IdNameValue**](IdNameValue.md) | It is used in Global config template plan creation. Needs in plan creation on global commcell | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

