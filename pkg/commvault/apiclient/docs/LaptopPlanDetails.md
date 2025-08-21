# LaptopPlanDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**General** | [***LaptopPlanGeneralInfo**](LaptopPlanGeneralInfo.md) |  | [optional] [default to null]
**AllowedFeatures** | [***AllowedFeatures**](AllowedFeatures.md) |  | [optional] [default to null]
**Rpo** | [***LaptopPlanRpo**](LaptopPlanRPO.md) |  | [optional] [default to null]
**BackupContent** | [***LaptopPlanBackupContent**](LaptopPlanBackupContent.md) |  | [optional] [default to null]
**RegionsConfigured** | **bool** | Specifies if the destinations are associated to regions | [optional] [default to null]
**BackupDestinations** | [**[]PlanBackupDestination**](PlanBackupDestination.md) |  | [optional] [default to null]
**Retention** | [***LaptopPlanRetention**](LaptopPlanRetention.md) |  | [optional] [default to null]
**AssociatedUsersAndUserGroups** | [**[]PlanUserOrGroups**](PlanUserOrGroups.md) |  | [optional] [default to null]
**Alerts** | [**[]IdName**](IdName.md) |  | [optional] [default to null]
**AllowPlanOverride** | **bool** | Flag to enable overriding of plan | [optional] [default to null]
**OverrideRestrictions** | [***LaptopPlanOverrideSettings**](LaptopPlanOverrideSettings.md) |  | [optional] [default to null]
**InheritSettings** | [***LaptopPlanInheritSettings**](LaptopPlanInheritSettings.md) |  | [optional] [default to null]
**OfflineLaptops** | [***AutoRetireDevices**](AutoRetireDevices.md) |  | [optional] [default to null]
**NetworkResources** | [***LaptopPlanNetworkResources**](LaptopPlanNetworkResources.md) |  | [optional] [default to null]
**Permissions** | [**[]IdName**](IdName.md) |  | [optional] [default to null]
**AdditionalProperties** | [***PlanAdditionalProperties**](PlanAdditionalProperties.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

