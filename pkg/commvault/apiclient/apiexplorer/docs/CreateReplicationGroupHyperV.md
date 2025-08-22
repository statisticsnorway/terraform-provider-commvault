# CreateReplicationGroupHyperV

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestVendor** | **string** |  | [default to null]
**AdvancedOptions** | [***ReplicationGroupAdvOptionsHyperV**](ReplicationGroupAdvOptionsHyperV.md) |  | [optional] [default to null]
**OverrideReplicationOptions** | [**[]OverrideReplicationOptionsHyperVCreate**](OverrideReplicationOptionsHyperVCreate.md) | The replication options per instance, if provided, will override the replication options from the target. For the instances not in this list, the options are applied from the target. | [optional] [default to null]
**Name** | **string** | The name of the replication group being created | [default to null]
**SourceHypervisor** | [***IdName**](IdName.md) |  | [default to null]
**Vms** | [**[]NameGuid**](NameGUID.md) | A list of name and GUID of all the virtual machines that have to be replicated | [default to null]
**FrequencyInMinutes** | **int32** | The field denotes the frequency of replication. | [optional] [default to 240]
**RecoveryTarget** | [***IdName**](IdName.md) |  | [default to null]
**Storage** | [**[]StorageCopyCreate**](StorageCopyCreate.md) | The primary and an optional secondary storage that will be used for storing the source VM data for replication. The secondary storage if provided, will be the default source for replication. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

