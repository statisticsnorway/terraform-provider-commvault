# UpdateBlackoutWindow

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewName** | **string** | Refers to the newName given to the blackout Window. | [optional] [default to null]
**BackupOperations** | **[]string** | Refers to backup types to include in the blackout window | [optional] [default to null]
**Company** | [***IdName**](IdName.md) |  | [optional] [default to null]
**AllDays** | [**[]DaysAndTimes**](DaysAndTimes.md) | Days of the week along with the time on which the black out window will be in effect. | [optional] [default to null]
**Weeks** | [**[]WeekOfMonth**](WeekOfMonth.md) | Refers to the weeks of the month that the blackout window will be in effect. | [optional] [default to null]
**BetweenDates** | [***StartEnd**](StartEnd.md) |  | [optional] [default to null]
**DoNotSubmitJob** | **bool** | Allows or Denies submitting a job when the blackout window is in effect. If allowed, the job is submitted and resumed once the blackout window ends. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

