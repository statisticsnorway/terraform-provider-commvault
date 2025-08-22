# FailoverGroupArrayReplicationPair

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the array replication pair | [optional] [default to null]
**SourceId** | **int32** | The id of the source machine | [optional] [default to null]
**SourceName** | **string** | The name of the source machine | [optional] [default to null]
**DestinationId** | **int32** | The id of the destination machine | [optional] [default to null]
**DestinationName** | **string** | The name of the destination machine | [optional] [default to null]
**GUID** | **string** | The GUID of the source of array replication pair | [optional] [default to null]
**Priority** | **int32** | The priority of the array replication pair for DR operation in failover group. The lower values signify higher priority | [optional] [default to null]
**IpSettings** | [**[]IpAddressSettingVmWareCreate**](IpAddressSettingVMWareCreate.md) | The settings for IP address mapping | [optional] [default to null]
**Hostname** | **string** | The hostname specified for the destination machine | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

