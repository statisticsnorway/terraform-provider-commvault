# ArchivePlanBackupDestination

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanBackupDestination** | [***IdName**](IdName.md) |  | [optional] [default to null]
**IsDefault** | **bool** | Is this a default backup destination? | [optional] [default to null]
**CopyType** | [***PlanCopyType**](PlanCopyType.md) |  | [optional] [default to null]
**CopyTypeName** | [***PlanCopyTypeName**](PlanCopyTypeName.md) |  | [optional] [default to null]
**CopyPrecedence** | **int32** | Order of copy created in storage policy | [optional] [default to null]
**StoragePool** | [***StoragePool**](StoragePool.md) |  | [optional] [default to null]
**StorageType** | [***StorageType**](StorageType.md) |  | [optional] [default to null]
**SourceCopy** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Region** | [***IdNameDisplayName**](IdNameDisplayName.md) |  | [optional] [default to null]
**EnableDataAging** | **bool** | Tells if this copy has data aging enabled | [optional] [default to null]
**RetentionPeriodDays** | **int32** | Retention period in days | [optional] [default to null]
**OverrideRetentionSettings** | **bool** | Tells if this copy should use storage pool retention period days or the retention defined for this copy. Set as true to use retention defined on this copy. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

