# OverrideReplicationOptionsAmazon

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuestCreds** | [***GuestCredentialsCreate**](GuestCredentialsCreate.md) |  | [optional] [default to null]
**InstanceName** | **string** | Display name for the destination instance | [optional] [default to null]
**AvailabilityZone** | **string** | Destination zone | [optional] [default to null]
**VolumeType** | **string** | Type of volume used for replication | [optional] [default to null]
**EncryptionKey** | **string** | Encryption key if used | [optional] [default to null]
**Network** | [***AmazonNetwork**](AmazonNetwork.md) |  | [optional] [default to null]
**TestFailoverNetwork** | [***AmazonNetwork**](AmazonNetwork.md) |  | [optional] [default to null]
**AutoSelectSecurityGroup** | [**[]SecurityGroup**](SecurityGroup.md) | Select a security group for the destination instances | [optional] [default to null]
**InstanceType** | **string** | Provides the available CPU cores and memory for the instance. | [optional] [default to null]
**ComputerName** | **string** | Enable required drivers to be installed on the Amazon guest instance | [optional] [default to null]
**UserName** | **string** | Enable required drivers to be installed on the Amazon guest instance | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

