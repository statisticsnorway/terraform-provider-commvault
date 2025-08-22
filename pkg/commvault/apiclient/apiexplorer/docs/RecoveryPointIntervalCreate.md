# RecoveryPointIntervalCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrashConsistentRpInterval** | **int32** | Minimum time interval(in seconds) between VM recovery points. Applicable only for Point in time recovery. | [optional] [default to 300]
**ApplicationConsistentRpInterval** | **int32** | Time(in seconds) after which the source VMs are quiesced to create application-consistent recovery points for destination VMs. Applicable only for Point in time recovery. | [optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

