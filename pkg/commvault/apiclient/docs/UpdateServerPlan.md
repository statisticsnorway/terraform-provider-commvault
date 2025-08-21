# UpdateServerPlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewName** | **string** | New plan name to update | [optional] [default to null]
**BackupContent** | [***PlanContent**](PlanContent.md) |  | [optional] [default to null]
**FilesystemAddon** | **bool** | flag to enable backup content association for applicable file system workload. | [optional] [default to null]
**Rpo** | [***ServerPlanUpdateRpo**](ServerPlanUpdateRPO.md) |  | [optional] [default to null]
**DatabaseOptions** | [***ServerPlanDatabaseOptionsInfo**](ServerPlanDatabaseOptionsInfo.md) |  | [optional] [default to null]
**RegionToConfigure** | [***IdName**](IdName.md) |  | [optional] [default to null]
**BackupDestinationIds** | **[]int32** | Primary Backup Destination Ids (which were created before plan creation). | [optional] [default to null]
**SnapshotOptions** | [***PlanSnapshotOptions**](PlanSnapshotOptions.md) |  | [optional] [default to null]
**Workload** | [***PlanWorkloads**](PlanWorkloads.md) |  | [optional] [default to null]
**AllowPlanOverride** | **bool** | Flag to enable overriding of plan. Once enabled, cannot be disabled. | [optional] [default to null]
**OverrideRestrictions** | [***PlanOverrideSettings**](PlanOverrideSettings.md) |  | [optional] [default to null]
**OverrideInheritSettings** | [***PlanOverrideInheritSetting**](PlanOverrideInheritSetting.md) |  | [optional] [default to null]
**Settings** | [***ServerPlanSettings**](ServerPlanSettings.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

