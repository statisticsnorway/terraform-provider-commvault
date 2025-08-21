# V4BlackoutWindowSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | blackout window Id | [optional] [default to null]
**Name** | **string** | blackout window name | [optional] [default to null]
**Company** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Days** | [**[]DayOfTheWeek**](DayOfTheWeek.md) | Days of the week on which the black out window will be in effect. | [optional] [default to null]
**BetweenDates** | [***StartEnd**](StartEnd.md) |  | [optional] [default to null]
**Time** | [**[]StartEnd**](StartEnd.md) | Refers to the time between which the blackout window will be in effect. It has to be provided in seconds. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

