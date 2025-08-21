# LaptopPlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanName** | **string** |  | [default to null]
**AllowedFeatures** | [***LaptopPlanAllowedFeatures**](LaptopPlanAllowedFeatures.md) |  | [optional] [default to null]
**BackupContent** | [***LaptopPlanBackupContent**](LaptopPlanBackupContent.md) |  | [optional] [default to null]
**StorageAndSchedule** | [***LaptopPlanStorageAndSchedule**](LaptopPlanStorageAndSchedule.md) |  | [optional] [default to null]
**Retention** | [***LaptopPlanRetention**](LaptopPlanRetention.md) |  | [optional] [default to null]
**NetworkResources** | [***LaptopPlanNetworkResources**](LaptopPlanNetworkResources.md) |  | [optional] [default to null]
**Alerts** | [**[]IdName**](IdName.md) |  | [optional] [default to null]
**AllowPlanOverride** | **bool** | Flag to enable overriding of plan. Plan cannot be overriden by default. | [optional] [default to false]
**OverrideRestrictions** | [***LaptopPlanOverrideSettings**](LaptopPlanOverrideSettings.md) |  | [optional] [default to null]
**InviteUsersOrGroups** | [**[]PlanUserOrGroups**](PlanUserOrGroups.md) | The users and user groups who should install the end-user Endpoint package on their devices. | [optional] [default to null]
**ParentPlan** | [***IdName**](IdName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

