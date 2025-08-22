# PlanPattern

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinuteFrequency** | **int32** | Frequency in minutes | [optional] [default to null]
**HourlyFrequency** | **int32** | Frequency in hours per day. Precedence is hourly, daily, weekly and monthly, in that order. | [optional] [default to null]
**DailyFrequency** | **int32** | Frequency in days | [optional] [default to null]
**WeeklyFrequency** | **int32** | Frequency in weeks. Only for incremental backup frequency. | [optional] [default to null]
**WeeklyFrequencyDays** | [**[]DayOfTheWeek**](DayOfTheWeek.md) | Days of the week for weekly frequency | [optional] [default to null]
**MonthlyFrequency** | **int32** | Frequency in months. Only for incremental backup frequency. | [optional] [default to null]
**MonthlyFrequencyWeekOfMonth** | [***WeekOfTheMonth**](WeekOfTheMonth.md) |  | [optional] [default to null]
**MonthlyFrequencyDayOfWeek** | [***DayOfWeek**](DayOfWeek.md) |  | [optional] [default to null]
**MonthlyFrequencyDayOfMonth** | **int32** | Only available for incremental backup frequency | [optional] [default to null]
**YearlyFrequencyWeekOfMonth** | [***WeekOfTheMonth**](WeekOfTheMonth.md) |  | [optional] [default to null]
**YearlyFrequencyDayOfWeek** | [***DayOfWeek**](DayOfWeek.md) |  | [optional] [default to null]
**YearlyFrequencyDayOfMonth** | **int32** | Only available for incremental backup frequency | [optional] [default to null]
**YearlyFrequencyMonthOfYear** | [***MonthOfYear**](MonthOfYear.md) |  | [optional] [default to null]
**StartTime** | **int32** | Time in seconds from the beginning of the day when the backup starts. This is a mandatory field for daily, weekly, monthly, yearly frequencies | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

