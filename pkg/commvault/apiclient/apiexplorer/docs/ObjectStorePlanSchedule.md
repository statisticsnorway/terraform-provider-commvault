# ObjectStorePlanSchedule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleId** | **int32** | Id of the schedule if available, required for modifying, deleting schedule | [optional] [default to null]
**ScheduleName** | **string** | Name of the schedule, for modify | [optional] [default to null]
**PolicyId** | **int32** | Schedule policy Id to which the schedule belongs | [optional] [default to null]
**ScheduleOperation** | **string** | Operation being performed on schedule | [optional] [default to SCHEDULE_OPERATION.MODIFY]
**SchedulePattern** | [***SchedulePattern**](SchedulePattern.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

