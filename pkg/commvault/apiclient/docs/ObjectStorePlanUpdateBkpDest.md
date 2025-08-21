# ObjectStorePlanUpdateBkpDest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanBackupDestinationId** | **int32** | Id of the backup destination to modify/delete | [optional] [default to null]
**PlanBackupDestinationName** | **string** | Backup Destination Name for add/modify operation | [optional] [default to null]
**StoragePool** | [***IdName**](IdName.md) |  | [optional] [default to null]
**StorageType** | [***StorageType**](StorageType.md) |  | [optional] [default to null]
**RetentionPeriodDays** | **int32** | Retention period in days. -1 can be specified for infinite retention. | [optional] [default to null]
**UseExtendedRetentionRules** | **bool** | Use extended retention rules(not applicable to the Primary(Snap) copy) | [optional] [default to null]
**ExtendedRetentionRules** | [***ExtendedRetentionRules**](ExtendedRetentionRules.md) |  | [optional] [default to null]
**EnableDataAging** | **bool** | enable or disable Datat aging on the backup destiantion | [optional] [default to null]
**SourceCopy** | [***IdName**](IdName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

