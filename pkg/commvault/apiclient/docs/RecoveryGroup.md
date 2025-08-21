# RecoveryGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Id of the recovery group | [optional] [default to null]
**Name** | **string** | Name of the recovery group | [optional] [default to null]
**Target** | [***IdName**](IdName.md) |  | [optional] [default to null]
**PostRecoveryActions** | [**[]DrOperationScript**](DROperationScript.md) |  | [optional] [default to null]
**DelayBetweenPriorityMachines** | **int32** | The delay between machines in different priorities in minutes | [optional] [default to null]
**ContinueOnFailure** | **bool** | Set to true to continue to the next priority machines on failure | [optional] [default to false]
**Action** | [***RecoveryAction**](RecoveryAction.md) |  | [optional] [default to null]
**RecoveryPointDetails** | [***RecoveryGroupRpDetails**](RecoveryGroupRPDetails.md) |  | [optional] [default to null]
**RecoveryPoint** | **int32** | Timestamp for group restore in case of disaster | [optional] [default to 0]
**RecoveryExpirationOptions** | [***RecoveryExpirationOptions**](RecoveryExpirationOptions.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

