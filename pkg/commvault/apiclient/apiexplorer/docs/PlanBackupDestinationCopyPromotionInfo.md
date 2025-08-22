# PlanBackupDestinationCopyPromotionInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Immediate** | **bool** | Tells if the copy has to be promoted immediately. | [optional] [default to null]
**Hours** | **int32** | Tells the number of hours to wait before failing/killing jobs for copy promotion | [optional] [default to null]
**WaitForRunningJobs** | **bool** | Tells if copy promotion needs to wait for running jobs | [optional] [default to null]
**WaitForSynchronize** | **bool** | Tells if copy promotion needs to wait for synchronization of copies. | [optional] [default to null]
**ForceKill** | **bool** | Tells if we need to force kill jobs and promote copy after waiting. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

