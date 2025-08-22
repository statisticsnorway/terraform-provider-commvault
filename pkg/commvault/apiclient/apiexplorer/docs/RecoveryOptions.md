# RecoveryOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryType** | **string** | Type of recovery. Values are case sensitive | [optional] [default to null]
**RecoveryPointStore** | **string** | Select the Recovery Point Store on the destination site where journal entries for each recovery point are stored | [optional] [default to null]
**CrashConsistentRpInterval** | **int32** | Option to create crash-consistent recovery points | [optional] [default to null]
**MergeRecoveryPointsOlderThan** | **int32** | Option to combine older recovery points into larger recovery points | [optional] [default to null]
**RetainRecoveryPointsFor** | **int32** | Option to specify how long journal entries for a recovery point should be retained | [optional] [default to null]
**RecoveryPointInterval** | **int32** | Option to specify the time interval between the oldest recovery points in the RP Store | [optional] [default to null]
**PruneAndMergeRecoveryPoints** | **bool** | Specify whether to transfer updates for the oldest recovery points to destination computer during off-peak hours. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

