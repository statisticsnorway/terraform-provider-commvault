# FailoverGroupReplicationPair

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the replication pair | [optional] [default to null]
**SourceId** | **int32** | The id of the source machine | [optional] [default to null]
**SourceName** | **string** | The name of the source machine | [optional] [default to null]
**DestinationId** | **int32** | The id of the destination machine | [optional] [default to null]
**DestinationName** | **string** | The name of the destination machine | [optional] [default to null]
**GUID** | **string** | The GUID of the source of replication pair | [optional] [default to null]
**Priority** | **int32** | The priority of the replication pair for DR operation in failover group. The lower values signify higher priority | [optional] [default to null]
**CopyPrecedence** | **int32** | The copy precedence of replication pair | [optional] [default to null]
**PreFailoverScript** | [***DrOperationScript**](DROperationScript.md) |  | [optional] [default to null]
**PostFailoverScript** | [***DrOperationScript**](DROperationScript.md) |  | [optional] [default to null]
**PreFailbackScript** | [***DrOperationScript**](DROperationScript.md) |  | [optional] [default to null]
**PostFailbackScript** | [***DrOperationScript**](DROperationScript.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

