# GoogleCloudVmInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **string** | Instance ID of the google cloud virtual machine to be restored. Ex: 123456789123456789. | [optional] [default to null]
**Name** | **string** | The new name of the instance after restore. Defaults to source name after restore. | [optional] [default to Defaults to source VM name]
**ProjectId** | **string** | GCP Project ID | [optional] [default to null]
**Zone** | **string** | GCP zones. Ex: asia-east1-a. Defaults to source instance zone after restore. For reference: cloud.google.com/compute/docs/regions-zones#available | [optional] [default to Defaults to source instance zone after restore. For reference: cloud.google.com/compute/docs/regions-zones#available']
**Region** | **string** | GCP Region. Ex: asia-east1. Defaults to source instance region after restore. For reference: cloud.google.com/compute/docs/regions-zones#available | [optional] [default to Defaults to source instance region after restore. For reference: cloud.google.com/compute/docs/regions-zones#available]
**MachineType** | **string** | GCP Machine Family. Ex: c2d-highmem-8. Defaults to --Auto Select--. For Reference: cloud.google.com/compute/docs/machine-resource | [optional] [default to --Auto Select--]
**CreatePublicIP** | **bool** | Flag to create external IP after restore. Defaults to false. | [optional] [default to false]
**NetworkInterfaces** | [**[]GoogleCloudNetworkInterfaceInfo**](GoogleCloudNetworkInterfaceInfo.md) | Google Cloud VM network interfaces list | [optional] [default to null]
**CustomMetadata** | [**[]NameValue**](NameValue.md) | GCP custom metadata key-value pairs | [optional] [default to null]
**NodeAffinities** | [**[]NameValue**](NameValue.md) | Node affinity for the virtual machine to provision virtual machine as sole tenanat VM. Defaults to no node affinity. Format: &#x27;[{name: compute.googleapis.com/node-group-name, value: {Node Group Name} }, {name: compute.googleapis.com/node-name, value: {Node Name} }]&#x27; | [optional] [default to null]
**ServiceAccount** | [***ServiceAccount**](ServiceAccount.md) |  | [optional] [default to null]
**EncryptionKey** | **string** | Customer Managed Encryption Key | [optional] [default to null]
**KeyProtectionLevel** | **string** | Protection level of encryption key | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

