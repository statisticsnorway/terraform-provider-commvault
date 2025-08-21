# AwsInstanceInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **string** | AWS instance ID | [optional] [default to null]
**InstanceName** | **string** | AWS instance name to be set after restore. Defaults to source instance name. | [optional] [default to Defaults to source instance name]
**Zone** | **string** | AWS Availability zone. Defaults to source instance availability zone. | [optional] [default to Defaults to source instance availability zone.]
**Region** | **string** |  | [optional] [default to null]
**InstanceSize** | **string** | Instance size to be after restore. Defaults to source instance size | [optional] [default to Defaults to source instance size]
**VolumeType** | **string** |  | [optional] [default to null]
**EncryptionKey** | [***VsaEncryptionOption**](VSAEncryptionOption.md) |  | [optional] [default to null]
**Nics** | [**[]AwsNetworkInterfaceInfo**](AWSNetworkInterfaceInfo.md) | AWS instance network interface list | [optional] [default to null]
**SecurityGroups** | [**[]VmSecurityGroupInfo**](VMSecurityGroupInfo.md) | List of AWS security groups to be set for the instance. | [optional] [default to null]
**VmTags** | [**[]NameValue**](NameValue.md) | AWS VM tag list | [optional] [default to null]
**RestoreSourceVMTags** | **bool** | Restore source VM tags. | [optional] [default to true]
**KeyPairs** | [**[]VmKeyPairInfo**](VMKeyPairInfo.md) | List of key pairs | [optional] [default to null]
**RestoreSourceNetworkConfig** | **bool** | If set to true, we will use the source network configuration during restore | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

