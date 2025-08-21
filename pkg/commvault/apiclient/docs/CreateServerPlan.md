# CreateServerPlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanName** | **string** | Name of the new plan | [default to null]
**BackupDestinations** | [**[]CreatePlanBackupDestination**](CreatePlanBackupDestination.md) | Backup destinations for the plan. Specify where you want to store your backup data. | [optional] [default to null]
**BackupDestinationIds** | **[]int32** | Primary Backup Destination Ids (which were created before plan creation). This is only considered when backupDestinations array object is not defined. | [optional] [default to null]
**Rpo** | [***ServerBackupPlanRpo**](ServerBackupPlanRPO.md) |  | [optional] [default to null]
**FilesystemAddon** | **bool** | flag to enable backup content association for applicable file system workload. | [optional] [default to null]
**BackupContent** | [***PlanContent**](PlanContent.md) |  | [optional] [default to null]
**SnapshotOptions** | [***CreatePlanSnapshotOptions**](CreatePlanSnapshotOptions.md) |  | [optional] [default to null]
**DatabaseOptions** | [***ServerPlanDatabaseOptions**](ServerPlanDatabaseOptions.md) |  | [optional] [default to null]
**Workload** | [***PlanWorkloads**](PlanWorkloads.md) |  | [optional] [default to null]
**AllowPlanOverride** | **bool** | Flag to enable overriding of plan. Plan cannot be overriden by default. | [optional] [default to null]
**OverrideRestrictions** | [***PlanOverrideSettings**](PlanOverrideSettings.md) |  | [optional] [default to null]
**ParentPlan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**AdditionalProperties** | [***PlanAdditionalProperties**](PlanAdditionalProperties.md) |  | [optional] [default to null]
**Settings** | [***ServerPlanSettings**](ServerPlanSettings.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

