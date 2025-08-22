# PlanSchedule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleId** | **int32** | Id of the schedule if available, required for modifying, deleting schedule | [optional] [default to null]
**ScheduleName** | **string** | Name of the schedule, for modify | [optional] [default to null]
**PolicyId** | **int32** | Schedule policy Id to which the schedule belongs | [optional] [default to null]
**ScheduleOperation** | **string** | Operation being performed on schedule | [optional] [default to SCHEDULE_OPERATION.MODIFY]
**VmOperationType** | **string** | Type of DR operation (only applicable for Failover groups) | [optional] [default to null]
**BackupType** | **string** | Schedule Backup level | [default to null]
**ForDatabasesOnly** | **bool** | Boolean to indicate if schedule is for database agents | [optional] [default to false]
**IsRetentionBasedSyntheticFull** | **bool** | Boolean to indicate if synthetic full schedule is based on retention rules | [optional] [default to false]
**IsAuxCopySchedule** | **bool** | Boolean to indicate if schedule is aux copy schedule. | [optional] [default to false]
**SchedulePattern** | [***SchedulePattern**](SchedulePattern.md) |  | [default to null]
**ScheduleOption** | [***ScheduleOption**](ScheduleOption.md) |  | [optional] [default to null]
**AdditionalInfo** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

