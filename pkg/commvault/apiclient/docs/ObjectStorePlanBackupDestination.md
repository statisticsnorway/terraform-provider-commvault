# ObjectStorePlanBackupDestination

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupDestinationName** | **string** | Backup destination name | [default to null]
**StoragePool** | [***IdName**](IdName.md) |  | [optional] [default to null]
**StorageType** | [***StorageType**](StorageType.md) |  | [default to null]
**RetentionPeriodDays** | **int32** | Retention period in days. -1 can be specified for infinite retention. | [default to null]
**UseExtendedRetentionRules** | **bool** | Use extended retention rules | [optional] [default to null]
**ExtendedRetentionRules** | [***ExtendedRetentionRules**](ExtendedRetentionRules.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

