# ReplicationMonitorDetailRecoveryOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoverypointStore** | [***IdName**](IdName.md) |  | [optional] [default to null]
**CCRPInterval** | **int32** | Time interval between crash consistent recovery points in seconds | [optional] [default to null]
**MergeRecoveryPointsOlderThan** | **int32** | Time which should be satisfied to merge recovery points in seconds | [optional] [default to null]
**RetainRecoveryPointsFor** | **int32** | Recovery points retention time in seconds | [optional] [default to null]
**RPIntervalAfterRetention** | **int32** | Recovery point interval at the end of retention time in seconds | [optional] [default to null]
**PruneAndMergeOffPeakOnly** | **bool** | Boolean which determines Prune and Merge Recovery Points during off peak time only. | [optional] [default to null]
**ACRPInterval** | **int32** | Gives information about application consistent recovery point interval in seconds | [optional] [default to null]
**SwitchToLatestIfStoreOfflineFor** | **int32** | Gives information about switching to  latest recovery point store to latest if it is offline for time  in seconds | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

