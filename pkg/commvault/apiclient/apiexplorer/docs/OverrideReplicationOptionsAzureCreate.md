# OverrideReplicationOptionsAzureCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceVM** | [***NameGuid**](NameGUID.md) |  | [default to null]
**VmDisplayName** | **string** | Display name of destination VM | [default to null]
**ResourceGroup** | **string** | The resource group to be used for the destination VM | [default to null]
**Region** | **string** | The name of the region where the destination VM will reside | [default to null]
**StorageAccount** | **string** | The name of the storage account to be used on the destination VM | [default to null]
**VmSizeId** | **string** | The id of the vm size to be applied to the destination VM. Default value is Auto | [optional] [default to null]
**DiskTypeId** | **string** | The id of the disk type to be used for the destination VM. Default value is Auto | [optional] [default to null]
**VirtualNetwork** | [**[]NetworkSubnet**](NetworkSubnet.md) | The network to be used on the destination VM. Default value is Auto | [optional] [default to null]
**SecurityGroup** | [***SecurityGroup**](SecurityGroup.md) |  | [optional] [default to null]
**CreatePublicIP** | **bool** | Creates a public IP on the network | [optional] [default to null]
**RestoreAsManagedVM** | **bool** | Creates the destination as a managed VM | [optional] [default to null]
**PrivateIPaddress** | **string** | Private IP address of the network | [optional] [default to null]
**PublicIPaddress** | **string** | Public IP address of the network | [optional] [default to null]
**PublicIPaddressID** | **string** | Public IP address id of the network | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

