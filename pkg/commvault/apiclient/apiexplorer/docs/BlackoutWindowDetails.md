# BlackoutWindowDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the blackout window operation rule | [optional] [default to null]
**BackupOperations** | [***BackupOperationsList**](backupOperationsList.md) |  | [optional] [default to null]
**DoNotRunJob** | [**[]DoNotRunJobDetails**](doNotRunJobDetails.md) |  | [optional] [default to null]
**BetweenDates** | **bool** | Between these dates, do not run job | [optional] [default to null]
**StartDate** | **int32** | Start date in seconds, during which the operation will not run | [optional] [default to null]
**EndDate** | **int32** | End date in seconds, during which the operation will not run | [optional] [default to null]
**DoNotSubmitJob** | **bool** | To skip a scheduled job | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

