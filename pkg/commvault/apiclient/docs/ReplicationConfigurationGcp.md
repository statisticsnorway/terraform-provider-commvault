# ReplicationConfigurationGcp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceVM** | **string** | Source VM instance name | [optional] [default to null]
**SourceVMGuid** | **string** | Source VM instance guid | [optional] [default to null]
**DestinationVM** | **string** | Destination VM instance name | [optional] [default to null]
**PrimaryZone** | **string** | Primary zone of the destination VM instance | [optional] [default to null]
**Datacenter** | **string** | Datacenter of the destination VM instance | [optional] [default to null]
**ProjectId** | **string** | Gcp project id | [optional] [default to null]
**CreatePublicIP** | **bool** | Public/External IP of the destination VM instance | [optional] [default to null]
**MachineType** | **string** | Destination VM instance machine type | [optional] [default to null]
**SecondaryZone** | **string** | Secondary zone of the destination VM instance | [optional] [default to null]
**OverrideReplicationOptions** | [***OverrideReplicationOptionsGcp**](OverrideReplicationOptionsGCP.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

