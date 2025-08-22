# PlanBackupDestinationDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**PlanBackupDestination** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Configurable** | **bool** | Flag to specify whether backup destination is configurable (True when it belongs to a base plan or a derived plan with overriden storage pool) | [optional] [default to null]
**Region** | [***IdNameDisplayName**](IdNameDisplayName.md) |  | [optional] [default to null]
**General** | [***PlanBackupDestinationGeneralInfo**](PlanBackupDestinationGeneralInfo.md) |  | [optional] [default to null]
**BackupType** | [***PlanBackupDestinationBackupTypeInfo**](PlanBackupDestinationBackupTypeInfo.md) |  | [optional] [default to null]
**RetentionRules** | [***PlanBackupDestinationRetentionRuleInfo**](PlanBackupDestinationRetentionRuleInfo.md) |  | [optional] [default to null]
**Mappings** | [**[]SnapshotCopyMapping**](SnapshotCopyMapping.md) |  | [optional] [default to null]
**GlobalConfigInfo** | [***GlobalConfigInfo**](GlobalConfigInfo.md) |  | [optional] [default to null]
**SnapOptions** | [***PlanBackupDestinationSnapOptions**](PlanBackupDestinationSnapOptions.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

