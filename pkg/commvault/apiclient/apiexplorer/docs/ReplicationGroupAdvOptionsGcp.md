# ReplicationGroupAdvOptionsGcp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnconditionalOverwrite** | **bool** | This will replace the instance at the destination if the instance with the same name already exists. | [optional] [default to null]
**DelayBetweenPriorityMachines** | **int32** | The delay between machines in different priorities. Mention in minutes | [optional] [default to 2]
**ContinueOnFailure** | **bool** | Set to true to continue to the next priority machines on failure. | [optional] [default to false]
**Script** | [***ReplicationGroupScript**](ReplicationGroupScript.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

