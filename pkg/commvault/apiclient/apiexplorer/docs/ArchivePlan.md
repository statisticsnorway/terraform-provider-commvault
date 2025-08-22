# ArchivePlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanName** | **string** |  | [default to null]
**BackupDestinations** | [**[]CreateArchivePlanBackupDestination**](CreateArchivePlanBackupDestination.md) |  | [optional] [default to null]
**Rpo** | [***ArchivePlanRpo**](ArchivePlanRPO.md) |  | [optional] [default to null]
**ArchivingRules** | [***ArchivePlanArchivingRules**](ArchivePlanArchivingRules.md) |  | [optional] [default to null]
**AllowPlanOverride** | **bool** | Flag to enable overriding of plan. Plan cannot be overriden by default. | [optional] [default to false]
**OverrideRestrictions** | [***ArchivePlanOverrideSettings**](ArchivePlanOverrideSettings.md) |  | [optional] [default to null]
**ParentPlan** | [***IdName**](IdName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

