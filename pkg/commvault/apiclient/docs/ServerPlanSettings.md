# ServerPlanSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileSearch** | [***PlanFileSearch**](PlanFileSearch.md) |  | [optional] [default to null]
**EnableAdvancedView** | **bool** | Setting to suggest plan has some advanced settings present. Setting is OEM specific and not applicable for all cases. | [optional] [default to null]
**DeviceStreams** | **int32** | For each region, the data to backup is divided into these many streams while writing to backup destination. | [optional] [default to 100]
**UpgradedFromStoragePolicy** | **bool** | Setting to suggest plan was created from PolicyToPlan workflow. If true, editing RPO is not allowed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

