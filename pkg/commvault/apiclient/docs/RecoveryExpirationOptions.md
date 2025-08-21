# RecoveryExpirationOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableExpirationOption** | **bool** | If true, cleans up recovered VMs after the set daysToExpire value | [optional] [default to true]
**DaysToExpire** | **int32** | Number of days after which recovered VMs are automatically cleaned up | [optional] [default to 7]
**IsRescuedCommServe** | **bool** | Whether the CommServe is rescued or not | [optional] [default to null]
**ExpirationTime** | **int32** | Timestamp when the clean up of all recovered VMs happens on the rescued CommServe. It happens 4 hours before the CommServe expiration time | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

