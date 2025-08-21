# SchedulePattern

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleFrequencyType** | **string** | schedule frequency type | [default to null]
**Frequency** | **int32** | Frequency of the schedule based on schedule frequency type eg. for Hours, value 2 is 2 hours, for Minutes, 30 is 30 minutes, for Daily, 2 is 2 days. for Monthly 2 is it repeats every 2 months | [optional] [default to null]
**WeeklyDays** | [**[]DayOfTheWeek**](DayOfTheWeek.md) | Days of the week for weekly frequency | [optional] [default to null]
**DayOfMonth** | **int32** | Day on which to run the schedule, applicable for monthly, yearly | [optional] [default to null]
**WeekOfMonth** | [***WeekOfTheMonth**](WeekOfTheMonth.md) |  | [optional] [default to null]
**DayOfWeek** | [***DayOfWeek**](DayOfWeek.md) |  | [optional] [default to null]
**MonthOfYear** | [***MonthOfYear**](MonthOfYear.md) |  | [optional] [default to null]
**RepeatIntervalInMinutes** | **int32** | How often in minutes in a day the schedule runs, applicable for daily, weekly, monthly and yearly frequency types. | [optional] [default to null]
**RepeatUntilTime** | **int32** | Until what time to repeat the schedule in a day, requires repeatIntervalInMinutes | [optional] [default to null]
**Timezone** | [***IdName**](IdName.md) |  | [optional] [default to null]
**StartTime** | **int32** | start time of schedule in seconds | [optional] [default to 75600]
**StartDate** | **int32** | start date of schedule in epoch format | [optional] [default to null]
**EndDate** | **int32** | Schedule end date in epoch format | [optional] [default to 0]
**NoOfTimes** | **int32** | The number of times you want the schedule to run. | [optional] [default to 0]
**Exceptions** | [**[]ScheduleRunException**](ScheduleRunException.md) | Exceptions to when a schedule should not run, either in dates or week of month and days | [optional] [default to null]
**DaysBetweenSyntheticFulls** | **int32** | No of days between two synthetic full jobs | [optional] [default to null]
**MaxBackupIntervalInMins** | **int32** | The number of mins to force a backup on automatic schedule. | [optional] [default to 240]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

