# RecoveryPointRetentionCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetainRecoveryPointsFor** | **int32** | The lengh of time(in seconds) to a retain the recovery point entries. Applicable only for Point in time recovery. | [default to 604800]
**MergeRecoveryPoints** | **bool** | Merge the recovery points beyond the time provided in mergeRecoveryPointsOlderThan | [optional] [default to true]
**MergeRecoveryPointsOlderThan** | **int32** | Applicable only if mergeRecoveryPoints is set to true. Beyond this period(in seconds), older recovery points will be merged into progressively coarser intervals going back in time. Applicable only for Point in time recovery. The value cannot exceed the value of retainRecoveryPointsFor. | [optional] [default to 172800]
**RecoveryPointIntervalAtEndofRetention** | **int32** | Applicable only if mergeRecoveryPoints is set to true.  Time interval(in seconds) between the older recovery points. Applicable only if mergeRecoveryPointsOlderThan and Point in time recovery is selected. The value cannot exceed the value of retainRecoveryPointsFor or 86400 seconds(1 day). | [optional] [default to 21600]
**RpStoreOfflineFor** | **int32** | Time(in seconds) after which the recovery type is switched to &#x27;Latest recovery&#x27; if RP store is offline. Applicable only for Point in time recovery. | [optional] [default to 0]
**PruneAndMergeDuringOffPeak** | **bool** | Transfers the updates of the oldest recovery points to destination VM only during off-peak hours. The peak interval time should be configured in the RP store.  Applicable only for Point in time recovery. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

