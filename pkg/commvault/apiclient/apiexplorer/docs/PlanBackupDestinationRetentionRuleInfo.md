# PlanBackupDestinationRetentionRuleInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableDataAging** | **bool** | Tells if this copy has data aging enabled | [optional] [default to null]
**OverrideRetentionSettings** | **bool** | Tells if this copy should use storage pool retention period days or the retention defined for this copy. Set as true to use retention defined on this copy. | [optional] [default to null]
**RetentionRuleType** | [***RetentionRuleTypes**](RetentionRuleTypes.md) |  | [optional] [default to null]
**RetentionPeriodDays** | **int32** | Retention period in days | [optional] [default to null]
**SnapRecoveryPoints** | **int32** | Number of snap recovery points for snap copy for retention | [optional] [default to null]
**UseExtendedRetentionRules** | **bool** | Should extended retention rules be used | [optional] [default to null]
**ExtendedRetentionRules** | [***ExtendedRetentionRules**](ExtendedRetentionRules.md) |  | [optional] [default to null]
**FullBackupTypesToBeRetained** | [***PlanFullBackupTypeToCopy**](PlanFullBackupTypeToCopy.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

