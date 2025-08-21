# ContinuousPairFailoverRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailoverType** | **string** | Type of failover to be performed | [default to null]
**RecoverFrom** | **string** | Recovery point type | [default to null]
**RecoveryPointTime** | **string** | Applicable if recoverFrom is GRAPH_SELECTION or RECOVERY_POINT. Format- MM/DD/YYY HH:MM:SS AM/PM | [optional] [default to null]
**Hostname** | **string** | Override boot options - Destination host name | [optional] [default to null]
**VmNetworkName** | **string** | Override boot options - Destination network | [optional] [default to null]
**FailoverVmName** | **string** | New Failover VM Name - Applicable only for Test Failover and Clone DR VM | [optional] [default to null]
**VmExpirationTime** | **int64** | Expiration time in seconds for new test failover VM - Applicable only for Test Failover | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

