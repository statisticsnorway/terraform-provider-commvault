# ReplicationConfigurationOci

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceVM** | **string** | Source VM instance name | [optional] [default to null]
**SourceVMGuid** | **string** | Source VM instance guid | [optional] [default to null]
**DestinationVM** | **string** | Destination VM instance name | [optional] [default to null]
**AvailabilityDomain** | **string** | Destination VM instance availability domain | [optional] [default to null]
**StagingBucket** | **string** | Staged bucket name for restore operation | [optional] [default to null]
**Region** | **string** | Instance region | [optional] [default to null]
**CreatePublicIP** | **bool** | Public/External IP of the destination VM instance | [optional] [default to null]
**Shape** | **string** | Destination VM instance shape | [optional] [default to null]
**OverrideReplicationOptions** | [***OverrideReplicationOptionsOci**](OverrideReplicationOptionsOCI.md) |  | [optional] [default to null]
**Compartment** | [***NameValue**](NameValue.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

