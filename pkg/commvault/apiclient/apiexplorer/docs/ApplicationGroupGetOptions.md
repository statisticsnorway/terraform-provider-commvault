# ApplicationGroupGetOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupStreams** | **int32** | Define number of parallel data readers | [optional] [default to 5]
**JobStartTime** | **int32** | Define the backup job start time in epochs | [optional] [default to null]
**OnSnapFailureFallbackToLiveVolume** | **bool** | Define setting to enable fallback to live volume backup in case of snap failure | [optional] [default to false]
**ScheduleWorkerToConfigNamespace** | **bool** | Define setting to enable scheduling worker Pods to CV Namespace for CSI-Snapshot enabled backups | [optional] [default to false]
**WorkerResources** | [***ApplicationGroupWorkerResourcesOptions**](ApplicationGroupWorkerResourcesOptions.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

