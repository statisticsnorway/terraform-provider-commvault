# UpdateLaptopPlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewName** | **string** | New plan name to update | [optional] [default to null]
**General** | [***LaptopPlanUpdateGeneralInfo**](LaptopPlanUpdateGeneralInfo.md) |  | [optional] [default to null]
**AllowedFeatures** | [***LaptopPlanAllowedFeatures**](LaptopPlanAllowedFeatures.md) |  | [optional] [default to null]
**Rpo** | [***LaptopPlanUpdateRpo**](LaptopPlanUpdateRPO.md) |  | [optional] [default to null]
**BackupContent** | [***LaptopPlanBackupContent**](LaptopPlanBackupContent.md) |  | [optional] [default to null]
**Retention** | [***LaptopPlanRetention**](LaptopPlanRetention.md) |  | [optional] [default to null]
**RegionToConfigure** | [***IdName**](IdName.md) |  | [optional] [default to null]
**AssociatedUsersAndUserGroups** | [**[]PlanUserOrGroups**](PlanUserOrGroups.md) |  | [optional] [default to null]
**Alerts** | [**[]IdName**](IdName.md) |  | [optional] [default to null]
**AllowPlanOverride** | **bool** | Flag to enable overriding of plan | [optional] [default to null]
**OverrideRestrictions** | [***LaptopPlanOverrideSettings**](LaptopPlanOverrideSettings.md) |  | [optional] [default to null]
**NetworkResources** | [***LaptopPlanNetworkResources**](LaptopPlanNetworkResources.md) |  | [optional] [default to null]
**OfflineLaptops** | [***AutoRetireDevices**](AutoRetireDevices.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

