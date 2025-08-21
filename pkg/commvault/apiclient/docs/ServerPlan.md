# ServerPlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**BackupContent** | [***PlanContent**](PlanContent.md) |  | [optional] [default to null]
**Rpo** | [***ServerPlanRpo**](ServerPlanRPO.md) |  | [optional] [default to null]
**RegionsConfigured** | **bool** | Specifies if the destinations are associated to regions | [optional] [default to null]
**BackupDestinations** | [**[]PlanBackupDestination**](PlanBackupDestination.md) | Backup destinations for the plan | [optional] [default to null]
**BackupDestinationIds** | **[]int32** | Primary Backup Destination Ids associated to this plan. | [optional] [default to null]
**DatabaseOptions** | [***ServerPlanDatabaseOptionsInfo**](ServerPlanDatabaseOptionsInfo.md) |  | [optional] [default to null]
**SnapshotOptions** | [***PlanSnapshotOptions**](PlanSnapshotOptions.md) |  | [optional] [default to null]
**Settings** | [***ServerPlanSettings**](ServerPlanSettings.md) |  | [optional] [default to null]
**AssociatedEntities** | [**[]IdNameCount**](IdNameCount.md) |  | [optional] [default to null]
**AllowPlanOverride** | **bool** | Is deriving and overriding the plan allowed | [optional] [default to null]
**OverrideRestrictions** | [***PlanOverrideSettings**](PlanOverrideSettings.md) |  | [optional] [default to null]
**InheritSettings** | [***ServerPlanInheritSettings**](ServerPlanInheritSettings.md) |  | [optional] [default to null]
**ParentInheritSettings** | [***ServerPlanInheritSettings**](ServerPlanInheritSettings.md) |  | [optional] [default to null]
**Permissions** | [**[]IdName**](IdName.md) |  | [optional] [default to null]
**AdditionalProperties** | [***PlanAdditionalProperties**](PlanAdditionalProperties.md) |  | [optional] [default to null]
**Workload** | [***PlanWorkloads**](PlanWorkloads.md) |  | [optional] [default to null]
**GlobalConfigInfo** | [***GlobalConfigInfo**](GlobalConfigInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

