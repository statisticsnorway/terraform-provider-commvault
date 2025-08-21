# FailoverGroupDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Failover group Id | [optional] [default to null]
**Name** | **string** | Failover group name | [optional] [default to null]
**SourceHypervisor** | [***FailoverGroupHypervisor**](FailoverGroupHypervisor.md) |  | [optional] [default to null]
**DestinationHypervisor** | [***FailoverGroupHypervisor**](FailoverGroupHypervisor.md) |  | [optional] [default to null]
**Source** | [***FailoverGroupSource**](FailoverGroupSource.md) |  | [optional] [default to null]
**IsApprovalRequired** | **bool** | Whether a user approval is required for DR operations | [optional] [default to null]
**UsersToNotify** | [**[]IdNameDisplayNameCompany**](IdNameDisplayNameCompany.md) | Users to notify/for approval of DR operations | [optional] [default to null]
**OperationType** | [***FailoverGroupOperationType**](FailoverGroupOperationType.md) |  | [optional] [default to null]
**Subclient** | [***IdName**](IdName.md) |  | [optional] [default to null]
**StoragePolicy** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Copy** | [***IdName**](IdName.md) |  | [optional] [default to null]
**PriorityIntervalInMinutes** | **int32** | The interval between DR operations of different priorities | [optional] [default to null]
**ContinueDROnFailure** | **bool** | Whether to continue to next priority on DR job failure | [optional] [default to null]
**RecoveryTarget** | [***IdName**](IdName.md) |  | [optional] [default to null]
**AutomaticFailover** | [***FailoverGroupAutomaticFailover**](FailoverGroupAutomaticFailover.md) |  | [optional] [default to null]
**PlannedFailoverSchedules** | [**[]PlanSchedule**](PlanSchedule.md) | The list of all planned failover schedules | [optional] [default to null]
**TestBootSchedules** | [**[]PlanSchedule**](PlanSchedule.md) | The list of all test boot schedules | [optional] [default to null]
**EsxServerMappings** | [**[]EsxServerMapping**](ESXServerMapping.md) | The list of all ESX server mappings | [optional] [default to null]
**ReplicationPairs** | [**[]FailoverGroupReplicationPair**](FailoverGroupReplicationPair.md) | List of all replication pairs that are part of the failover group | [optional] [default to null]
**ArrayReplicationPairs** | [**[]FailoverGroupArrayReplicationPair**](FailoverGroupArrayReplicationPair.md) | List of all array replication pairs for failover group | [optional] [default to null]
**Script** | [***ReplicationGroupScript**](ReplicationGroupScript.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

