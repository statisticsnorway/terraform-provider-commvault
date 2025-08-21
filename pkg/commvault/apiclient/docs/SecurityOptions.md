# SecurityOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]IdNameFullName**](IdNameFullName.md) | Users that have access to the recovery target | [optional] [default to null]
**UserGroups** | [**[]IdName**](IdName.md) | User groups that have access to the recovery target | [optional] [default to null]
**SecurityGroups** | [**[]SecurityGroup**](SecurityGroup.md) | List of security groups that provide controlled access of the DR VM | [optional] [default to null]
**TestSecurityGroups** | [**[]SecurityGroup**](SecurityGroup.md) | List of security groups that provide controlled access of the VM in cloud for test failover | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

