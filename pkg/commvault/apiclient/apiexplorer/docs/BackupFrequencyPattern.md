# BackupFrequencyPattern

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleFrequencyType** | **string** | schedule frequency type | [optional] [default to SCHEDULE_FREQUENCY_TYPE.DAILY]
**Frequency** | **int32** | Frequency of the schedule based on schedule frequency type eg. for Hours, value 2 is 2 hours, for Minutes, 30 is 30 minutes, for Daily, 2 is 2 days. for Monthly 2 is it repeats every 2 months | [optional] [default to 1]
**WeeklyDays** | [**[]DayOfTheWeek**](DayOfTheWeek.md) | Days of the week for weekly frequency | [optional] [default to null]
**DayOfMonth** | **int32** | Day on which to run the schedule, applicable for monthly, yearly | [optional] [default to null]
**WeekOfMonth** | [***WeekOfTheMonth**](WeekOfTheMonth.md) |  | [optional] [default to null]
**DayOfWeek** | [***DayOfWeek**](DayOfWeek.md) |  | [optional] [default to null]
**MonthOfYear** | [***MonthOfYear**](MonthOfYear.md) |  | [optional] [default to null]
**StartTime** | **int32** | start time of schedule in seconds for daily, weekly, monthly, yearly frequency | [optional] [default to 75600]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

