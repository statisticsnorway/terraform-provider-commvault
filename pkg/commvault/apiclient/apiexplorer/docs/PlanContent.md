# PlanContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WindowsIncludedPaths** | **[]string** | Paths to include for Windows | [optional] [default to null]
**WindowsExcludedPaths** | **[]string** | Paths to exclude for Windows | [optional] [default to null]
**WindowsFilterToExcludePaths** | **[]string** | Paths that are exception to excluded paths for Windows | [optional] [default to null]
**BackupSystemState** | **bool** | Do you want to back up the system state? Applicable only for Windows | [optional] [default to null]
**BackupSystemStateOnlyWithFullBackup** | **bool** | Do you want to back up system state only with full backup? Applicable only if the value of backupSystemState is true | [optional] [default to false]
**UseVSSForSystemState** | **bool** | Do you want to back up system state with VSS? Applicable only if the value of backupSystemState is true | [optional] [default to null]
**WindowsNumberOfDataReaders** | [***PlanContentDataReaders**](PlanContentDataReaders.md) |  | [optional] [default to null]
**MacIncludedPaths** | **[]string** | Paths to include for Mac | [optional] [default to null]
**MacExcludedPaths** | **[]string** | Paths to exclude for Mac | [optional] [default to null]
**MacFilterToExcludePaths** | **[]string** | Paths that are exception to excluded paths for Mac | [optional] [default to null]
**UnixNumberOfDataReaders** | [***PlanContentDataReaders**](PlanContentDataReaders.md) |  | [optional] [default to null]
**UnixIncludedPaths** | **[]string** | Paths to include for UNIX | [optional] [default to null]
**UnixExcludedPaths** | **[]string** | Paths to exclude for UNIX | [optional] [default to null]
**UnixFilterToExcludePaths** | **[]string** | Paths that are exception to excluded paths for Unix | [optional] [default to null]
**MacNumberOfDataReaders** | [***PlanContentDataReaders**](PlanContentDataReaders.md) |  | [optional] [default to null]
**ForceUpdateProperties** | **bool** | Do you want to sync properties on associated subclients even if properties are overriden at subclient level? | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

