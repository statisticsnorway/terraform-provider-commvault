# VmGroupDetailsSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HypervisorName** | **string** |  | [optional] [default to null]
**LastBackupTime** | **int32** | Last Backup time of the VMGroup | [optional] [default to null]
**LastBackupSize** | **int64** | Last Backup Job Size | [optional] [default to null]
**NextBackupTime** | **int32** | Next Backup Job start time | [optional] [default to null]
**IsDefaultVMGroup** | **bool** | True if subclient is default subclient | [optional] [default to false]
**BackupActivityStatus** | **string** | Current backup activity status | [optional] [default to null]
**Plan** | [***PlanIdNameType**](PlanIdNameType.md) |  | [optional] [default to null]
**Region** | [***RegionInfo**](RegionInfo.md) |  | [optional] [default to null]
**ReplicationGroup** | [***IdName**](IdName.md) |  | [optional] [default to null]
**TimeZone** | [***IdName**](IdName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

