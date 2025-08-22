# OverrideReplicationOptionsAmazonCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceVM** | [***NameGuid**](NameGUID.md) |  | [default to null]
**InstanceName** | **string** | Display name for the destination instance | [default to null]
**RegionName** | **string** | The name of the AWS region where the destination instance will reside | [default to null]
**AvailabilityZone** | **string** | The name of AWS zone where the destination instance will reside | [default to null]
**VolumeType** | **string** | ID of the volume type that will be used on the destination instance. Default value is Auto | [optional] [default to null]
**EncryptionKey** | **string** | The ID of the encryption key that will be used to encrypt the data of the desination instance. Default value is Auto | [optional] [default to null]
**Network** | [***NetworkVpcSubnet**](NetworkVPCSubnet.md) |  | [optional] [default to null]
**SecurityGroups** | [**[]SecurityGroup**](SecurityGroup.md) | The security group to be applied to the destination instance that acts as a virtual firewall. Default value is Auto | [optional] [default to null]
**InstanceType** | **string** | The available CPU cores and memory to be used on the destination instance. Default value is Auto | [optional] [default to null]
**GuestCredentials** | [***GuestCredentialsCreate**](GuestCredentialsCreate.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

