# UpdateArchivePlanBackupDestination

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewName** | **string** | New name for backup destination | [optional] [default to null]
**SourceCopy** | [***IdName**](IdName.md) |  | [optional] [default to null]
**EnableDataAging** | **bool** | Tells if this copy has data aging enabled | [optional] [default to null]
**RetentionPeriodDays** | **int32** | Retention period in days. -1 can be specified for infinite retention. | [optional] [default to null]
**OverrideRetentionSettings** | **bool** | Tells if this copy should use storage pool retention period days or the retention defined for this copy. Set as true to use retention defined on this copy. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

