# GetObjectStorePlanBackupDestination

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanBackupDestination** | [***IdName**](IdName.md) |  | [optional] [default to null]
**BackupDestinationName** | **string** | Backup destination name | [optional] [default to null]
**StoragePool** | [***StoragePool**](StoragePool.md) |  | [optional] [default to null]
**StorageType** | [***StorageType**](StorageType.md) |  | [optional] [default to null]
**RetentionPeriodDays** | **int32** | Retention period in days. -1 can be specified for infinite retention. | [optional] [default to null]
**EnableDataAging** | **bool** | Tells if this copy has data aging enabled | [optional] [default to null]
**CopyPrecedence** | **int32** | Order of backup destination copy created in storage policy | [optional] [default to null]
**BackupsToCopy** | [***PlanFullBackupType**](PlanFullBackupType.md) |  | [optional] [default to null]
**IsDefault** | **bool** | Is this a default backup destination? | [optional] [default to null]
**IsSnapCopy** | **bool** | Is this a snap copy? | [optional] [default to null]
**UseExtendedRetentionRules** | **bool** | Use extended retention rules | [optional] [default to null]
**ExtendedRetentionRules** | [***ExtendedRetentionRules**](ExtendedRetentionRules.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

