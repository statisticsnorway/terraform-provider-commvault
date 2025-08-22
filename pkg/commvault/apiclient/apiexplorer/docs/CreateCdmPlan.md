# CreateCdmPlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanName** | **string** | Name of the new plan | [default to null]
**BackupDestinations** | [**[]CreatePlanBackupDestination**](CreatePlanBackupDestination.md) | Copy destinations for the plan. Specify where you want to store your data. | [default to null]
**Rpo** | [***ServerBackupPlanRpo**](ServerBackupPlanRPO.md) |  | [optional] [default to null]
**SnapshotOptions** | [***CreatePlanSnapshotOptions**](CreatePlanSnapshotOptions.md) |  | [optional] [default to null]
**DatabaseOptions** | [***ServerPlanDatabaseOptions**](ServerPlanDatabaseOptions.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

