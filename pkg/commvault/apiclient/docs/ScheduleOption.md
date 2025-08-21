# ScheduleOption

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseDiskCacheForLogBackups** | **bool** | Used to enable disk caching feature on databases for automatic schedules on server plan | [optional] [default to null]
**CommitFrequencyInHours** | **int32** | Commit frequency in hours for disk cache backups from automatic schedules | [optional] [default to null]
**DaysBetweenAutoConvert** | **int32** | Number of days between auto conversion of backup level applicable for databases on incremental and differential schedules of server plan | [optional] [default to null]
**O365ItemSelectionOption** | **string** | item backup option for O365 V2 backup jobs | [optional] [default to null]
**JobRunningTimeInMins** | **int32** | total job running time in minutes | [optional] [default to null]
**MinBackupIntervalInMins** | **int32** | The min number of mins to check for file activity on automatic schedule. | [optional] [default to 15]
**LogFilesThreshold** | **int32** | The min number of archived log files before a backup job should start | [optional] [default to 50]
**LogsDiskUtilizationPercent** | **int32** | The min log destination disk threshold percentage | [optional] [default to 80]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

